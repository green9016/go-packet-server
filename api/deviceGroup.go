package api

import (
	"../model"
	"github.com/labstack/echo"
	"log"
)

func GetDeviceGroup(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var deviceGroup model.DeviceGroup
	db.First(&deviceGroup, id)
	return c.JSON(200, deviceGroup)
}


func GetDeviceGroups(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var deviceGroups []model.DeviceGroup
	db.Find(&deviceGroups)
	return c.JSON(200, deviceGroups)
}

func PostDeviceGroup(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var deviceGroup model.DeviceGroup
	c.Bind(&deviceGroup)

	log.Println(deviceGroup)

	db.Create(&deviceGroup)
 	return c.JSON(200, deviceGroup)
}

func UpdateDeviceGroup(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var deviceGroup model.DeviceGroup
	if err := db.Where("id = ?", id).First(&deviceGroup).Error; err != nil {
		c.Error(err)
		log.Println(err)
	 }
	 c.Bind(&deviceGroup)
	 db.Save(&deviceGroup)
	 return c.JSON(200, deviceGroup)

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/deviceGroups/1
}

func DeleteDeviceGroup(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var deviceGroup model.DeviceGroup
	d := db.Where("id = ?", id).Delete(&deviceGroup)
	log.Println(d)
	return c.JSON(200, "Deleted")
}