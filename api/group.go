package api

import (
	"../model"
	"github.com/labstack/echo"
	"log"
)

func GetGroup(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var group model.Group
	db.First(&group, id)
	return c.JSON(200, group)
}


func GetGroups(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var groups []model.Group
	db.Find(&groups)
	return c.JSON(200, groups)
}

func PostGroup(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var group model.Group
	c.Bind(&group)

	log.Println(group)

	db.Create(&group)
 	return c.JSON(200, group)
}

func UpdateGroup(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var group model.Group
	if err := db.Where("id = ?", id).First(&group).Error; err != nil {
		c.Error(err)
		log.Println(err)
	 }
	 c.Bind(&group)
	 db.Save(&group)
	 return c.JSON(200, group)

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/groups/1
}

func DeleteGroup(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var group model.Group
	d := db.Where("id = ?", id).Delete(&group)
	log.Println(d)
	return c.JSON(200, "Deleted")
}