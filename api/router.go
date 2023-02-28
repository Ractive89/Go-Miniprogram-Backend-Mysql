package api

import "github.com/gin-gonic/gin"

func (cc *CustomerController) RegisterRoutes(server *gin.Engine) {
	apiRoute := server.Group("/api/v1")
	apiRoute.POST("/getOpenID", cc.GetOpenID)
	apiRoute.POST("/createCustomer", cc.CreateCustomer)
	apiRoute.GET("/getOpenID/:oepnID", cc.GetCustomerByOpenID)
	apiRoute.GET("/getName/:name", cc.GetCustomerByName)
	apiRoute.GET("/getall", cc.GetAll)
	apiRoute.PATCH("/update", cc.UpdateCustomer)
	apiRoute.DELETE("/delete/:oepnID", cc.DeleteCustomer)
}
