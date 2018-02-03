package api

import (
	"../model"
	"github.com/labstack/echo"
	"log"
)

func GetDeviceLogin(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var deviceLogin model.DeviceLogin
	db.First(&deviceLogin, id)
	return c.JSON(200, deviceLogin)
}


func GetDeviceLogins(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var deviceLogins []model.DeviceLogin
	db.Find(&deviceLogins)
	return c.JSON(200, deviceLogins)
}

func PostDeviceLogin(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var deviceLogin model.DeviceLogin
	c.Bind(&deviceLogin)

	log.Println(deviceLogin)

	db.Create(&deviceLogin)
 	return c.JSON(200, deviceLogin)
}

func UpdateDeviceLogin(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var deviceLogin model.DeviceLogin
	if err := db.Where("id = ?", id).First(&deviceLogin).Error; err != nil {
		c.Error(err)
		log.Println(err)
	 }
	 c.Bind(&deviceLogin)
	 db.Save(&deviceLogin)
	 return c.JSON(200, deviceLogin)

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/deviceLogins/1
}

func DeleteDeviceLogin(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var deviceLogin model.DeviceLogin
	d := db.Where("id = ?", id).Delete(&deviceLogin)
	log.Println(d)
	return c.JSON(200, "Deleted")
}