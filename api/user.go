package api

import (
	"../model"
	"github.com/labstack/echo"
	"log"
)

func GetUser(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var user model.User
	db.First(&user, id)
	return c.JSON(200, user)
}


func GetUsers(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var users []model.User
	db.Find(&users)
	return c.JSON(200, users)
}

func PostUser(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	
	var user model.User
	c.Bind(&user)

	log.Println(user)

	db.Create(&user)
 	return c.JSON(200, user)
}

func UpdateUser(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var user model.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.Error(err)
		log.Println(err)
	 }
	 c.Bind(&user)
	 db.Save(&user)
	 return c.JSON(200, user)

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

func DeleteUser(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	id := c.Param("id")
	var user model.User
	d := db.Where("id = ?", id).Delete(&user)
	log.Println(d)
	return c.JSON(200, "Deleted")
}