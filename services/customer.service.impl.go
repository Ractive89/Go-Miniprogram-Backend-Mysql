package services

import (
	"dayang/conf"
	"dayang/models"
	"dayang/utils"
	"errors"

	"github.com/parnurzeal/gorequest"
	"gorm.io/gorm"
)

type CustomerServiceImpl struct {
	dayangDB *gorm.DB
}

func NewCustomerService(dayangDB *gorm.DB) CustomerService {
	return &CustomerServiceImpl{
		dayangDB: dayangDB,
	}
}

func (cs *CustomerServiceImpl) GetOpenID(data *string) (string, error) {

	url := utils.AddQuery("https://api.weixin.qq.com/sns/jscode2session",
		map[string]string{
			"grant_type": "authorization_code",
			"appid":      conf.AppID,
			"secret":     conf.AppSecret,
			"js_code":    *data,
		})

	_, body, errs := gorequest.New().Get(url).End()

	if errs != nil {
		return "", errors.New("request OpenID Error")
	}

	return body, nil
}

func (cs *CustomerServiceImpl) CreateCustomer(data *models.Customer) error {

	cs.dayangDB.Create(data)

	return nil
}

func (cs *CustomerServiceImpl) GetCustomerByOpenID(data *string) []models.Customer {

	var dataList []models.Customer
	cs.dayangDB.Where("open_id = ?", *data).Find(&dataList)
	return dataList
}

func (cs *CustomerServiceImpl) GetCustomerByName(data *string) []models.Customer {

	var dataList []models.Customer
	cs.dayangDB.Where("name = ?", *data).Find(&dataList)
	return dataList
}

func (cs *CustomerServiceImpl) GetAll(data *string, pageSize int, pageNum int) ([]models.Customer, int64, error) {
	if *data != conf.UserPassWord {
		return nil, 0, errors.New("request password Error")
	}

	var dataList []models.Customer

	var total int64

	cs.dayangDB.Model(dataList).Count(&total).Limit(pageSize).Offset(pageNum).Find(&dataList)

	if len(dataList) == 0 {
		return nil, 0, errors.New("database empty Error")
	} else {

		return dataList, total, nil
	}
}

func (cs *CustomerServiceImpl) UpdateCustomer(data *models.Customer) error {
	cs.dayangDB.Where("open_id =?", &data.OpenID).Updates(data)
	return nil
}

func (cs *CustomerServiceImpl) DeleteCustomer(data []models.Customer) error {
	cs.dayangDB.Delete(data)
	return nil
}
