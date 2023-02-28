package api

import (
	"dayang/models"
	"dayang/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomerController struct {
	CustomerService services.CustomerService
	SugerLogger     *zap.SugaredLogger
}

func NewCustomerController(customerService services.CustomerService, sugerLogger *zap.SugaredLogger) CustomerController {
	return CustomerController{
		CustomerService: customerService,
		SugerLogger:     sugerLogger,
	}
}

func (cc *CustomerController) GetOpenID(ctx *gin.Context) {
	var data string

	//获取POST数据
	if err := ctx.ShouldBind(&data); err != nil {
		cc.statusSend(Error, ctx, err.Error())
		return
	}

	// 从获取的数据发送到Wechat得到OpenID
	data, err := cc.CustomerService.GetOpenID(&data)
	if err != nil {
		cc.statusSend(Error, ctx, err.Error())
		return
	}

	// 返回数据
	ctx.JSON(http.StatusOK, data)
	cc.SugerLogger.Infoln("Success", http.StatusOK, "GetOpenID")
}

func (cc *CustomerController) CreateCustomer(ctx *gin.Context) {
	var data models.Customer

	if err := ctx.ShouldBind(&data); err != nil {
		cc.statusSend(Error, ctx, err.Error())
		return
	}

	// 检测是否有重复
	res := cc.CustomerService.GetCustomerByOpenID(&data.OpenID)

	// 无重复添加 有重复返回FUll
	if len(res) == 0 {
		err := cc.CustomerService.CreateCustomer(&data)
		if err != nil {
			cc.statusSend(Error, ctx, err.Error())
			return
		}

		cc.statusSend(Success, ctx, "CreateCustomer")
	} else {
		cc.statusSend(Full, ctx, "CreateCustomer")
	}
}

func (cc *CustomerController) GetCustomerByOpenID(ctx *gin.Context) {
	data := ctx.Param("oepnID")

	res := cc.CustomerService.GetCustomerByOpenID(&data)

	if len(res) == 0 {
		cc.statusSend(NotFound, ctx, "GetCustomerByOpenID")
	} else {
		ctx.JSON(http.StatusOK, res)
		cc.SugerLogger.Infoln("Success", http.StatusOK, "GetCustomerByOpenID")
	}

}

func (cc *CustomerController) GetCustomerByName(ctx *gin.Context) {
	data := ctx.Param("name")

	res := cc.CustomerService.GetCustomerByName(&data)

	if len(res) == 0 {
		cc.statusSend(NotFound, ctx, "GetCustomerByName")
	} else {
		ctx.JSON(http.StatusOK, res)
		cc.SugerLogger.Infoln("Success", http.StatusOK, "GetCustomerByName")
	}

}

func (cc *CustomerController) GetAll(ctx *gin.Context) {
	data := ctx.Query("password")
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	res, total, err := cc.CustomerService.GetAll(&data, pageSize, pageNum)
	if err != nil {
		cc.statusSend(Error, ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": gin.H{
			"customer": res,
			"total":    total,
			"pageNum":  pageNum,
			"pageSize": pageSize,
		},
	})
	cc.SugerLogger.Infoln("Success", http.StatusOK, "GetAll")

}

func (cc *CustomerController) UpdateCustomer(ctx *gin.Context) {
	var data models.Customer

	if err := ctx.ShouldBind(&data); err != nil {
		cc.statusSend(Error, ctx, err.Error())
		return
	}

	res := cc.CustomerService.GetCustomerByOpenID(&data.OpenID)
	if len(res) == 0 {
		cc.statusSend(NotFound, ctx)
	} else {
		cc.CustomerService.UpdateCustomer(&data)
		cc.statusSend(Updated, ctx, "UpdateCustomer")
	}
}

func (cc *CustomerController) DeleteCustomer(ctx *gin.Context) {
	data := ctx.Param("oepnID")

	res := cc.CustomerService.GetCustomerByOpenID(&data)

	if len(res) == 0 {
		cc.statusSend(NotFound, ctx, "DeleteCustomer")
	} else {
		cc.CustomerService.DeleteCustomer(res)
		cc.statusSend(Deleted, ctx, "DeleteCustomer")
	}

}
