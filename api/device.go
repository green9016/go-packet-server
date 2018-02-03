package api

import (
	"../model"
	"github.com/labstack/echo"
	"log"
)

func GetDevice(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var device model.Device
	db.First(&device, id)
	return c.JSON(200, device)
}


func GetDevices(c echo.Context) error {
	var db = InitDb()
	defer db.Close()

	var devices []model.Device
	db.Find(&devices)
	return c.JSON(200, devices)
}

func PostDevice(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var device model.Device
	c.Bind(&device)

	log.Println(device)

	db.Create(&device)
 	return c.JSON(200, device)
}

func UpdateDevice(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var device model.Device
	if err := db.Where("id = ?", id).First(&device).Error; err != nil {
		c.Error(err)
		log.Println(err)
	 }
	 c.Bind(&device)
	 db.Save(&device)
	 return c.JSON(200, device)

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/devices/1
}

func DeleteDevice(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var device model.Device
	d := db.Where("id = ?", id).Delete(&device)
	log.Println(d)
	return c.JSON(200, "Deleted")
}