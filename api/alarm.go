package api

import (
	"../model"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func GetAlarm(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var alarm model.Alarm
	db.First(&alarm, id)
	return c.JSON(http.StatusOK, alarm)
}


func GetAlarms(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var alarms []model.Alarm
	db.Find(&alarms)
	return c.JSON(http.StatusOK, alarms)
}

func PostAlarm(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var alarm model.Alarm
	c.Bind(&alarm)

	log.Println(alarm)

	db.Create(&alarm)
 	return c.JSON(http.StatusOK, alarm)
}

func UpdateAlarm(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var alarm model.Alarm
	if err := db.Where("id = ?", id).First(&alarm).Error; err != nil {
		c.Error(err)
		log.Println(err)
	 }
	 c.Bind(&alarm)
	 db.Save(&alarm)
	 return c.JSON(http.StatusOK, alarm)

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/alarms/1
}

func DeleteAlarm(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var alarm model.Alarm
	d := db.Where("id = ?", id).Delete(&alarm)
	log.Println(d)
	return c.JSON(200, "Deleted")
}