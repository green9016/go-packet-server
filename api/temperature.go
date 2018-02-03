package api

import (
	"../model"
	"github.com/labstack/echo"
	"log"
)

func GetTemperature(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var temperature model.Temperature
	db.First(&temperature, id)
	return c.JSON(200, temperature)
}


func GetTemperatures(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var temperatures []model.Temperature
	db.Find(&temperatures)
	return c.JSON(200, temperatures)
}

func PostTemperature(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var temperature model.Temperature
	c.Bind(&temperature)

	log.Println(temperature)

	db.Create(&temperature)
 	return c.JSON(200, temperature)
}

func UpdateTemperature(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var temperature model.Temperature
	if err := db.Where("id = ?", id).First(&temperature).Error; err != nil {
		c.Error(err)
		log.Println(err)
	 }
	 c.Bind(&temperature)
	 db.Save(&temperature)
	 return c.JSON(200, temperature)

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/temperatures/1
}

func DeleteTemperature(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var temperature model.Temperature
	d := db.Where("id = ?", id).Delete(&temperature)
	log.Println(d)
	return c.JSON(200, "Deleted")
}