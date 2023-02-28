# Go-Miniprogram-Backend-Mysql

通过Go-Gin开发微信小程序的后端实现对Mysql的CRUD

## Router

- POST("/getOpenID", cc.GetOpenID)

- POST("/createCustomer", cc.CreateCustomer)

- GET("/getOpenID/:oepnID", cc.GetCustomerByOpenID)

- GET("/getName/:name", cc.GetCustomerByName)

- GET("/getall", cc.GetAll)

- PATCH("/update", cc.UpdateCustomer)

- DELETE("/delete/:oepnID", cc.DeleteCustomer)
