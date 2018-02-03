package api

import (
	"../model"
	"github.com/labstack/echo"
	"log"
	"fmt"
)

func GetGpsData(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var gpsData model.GpsData
	db.First(&gpsData, id)
	return c.JSON(200, gpsData)
}


func GetGpsDatas(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	query := c.Param("deviceNumber")
	fmt.Println("query", query)
	var gpsDatas []model.GpsData
	db.Find(&gpsDatas)
	return c.JSON(200, gpsDatas)
}

func PostGpsData(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var gpsData model.GpsData
	c.Bind(&gpsData)

	log.Println(gpsData)

	db.Create(&gpsData)
 	return c.JSON(200, gpsData)
}

func UpdateGpsData(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var gpsData model.GpsData
	if err := db.Where("id = ?", id).First(&gpsData).Error; err != nil {
		c.Error(err)
		log.Println(err)
	 }
	 c.Bind(&gpsData)
	 db.Save(&gpsData)
	 return c.JSON(200, gpsData)

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/gpsDatas/1
}

func DeleteGpsData(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var gpsData model.GpsData
	d := db.Where("id = ?", id).Delete(&gpsData)
	log.Println(d)
	return c.JSON(200, "Deleted")
}