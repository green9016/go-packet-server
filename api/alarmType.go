package api

import (
	"../model"
	"github.com/labstack/echo"
)

func GetAlarmTypes(c echo.Context) error {
	var db = InitDb()
	defer db.Close()
	var alarmTypes []model.AlarmType
	db.Find(&alarmTypes)
	return c.JSON(200, alarmTypes)
}