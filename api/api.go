package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartApi() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

    api := e.Group("api")
    {
		//users
        api.POST("/users", PostUser)
        api.GET("/users", GetUsers)
		api.GET("/users/:id", GetUser)
        api.PUT("/users/:id", UpdateUser)
		api.DELETE("/users/:id", DeleteUser)
		
		//groups
        api.POST("/groups", PostGroup)
        api.GET("/groups", GetGroups)
		api.GET("/groups/:id", GetGroup)
        api.PUT("/groups/:id", UpdateGroup)
		api.DELETE("/groups/:id", DeleteGroup)

		
		//user groups
        api.POST("/deviceGroups", PostDeviceGroup)
        api.GET("/deviceGroups", GetDeviceGroups)
		api.GET("/deviceGroups/:id", GetDeviceGroup)
        api.PUT("/deviceGroups/:id", UpdateDeviceGroup)
		api.DELETE("/deviceGroups/:id", DeleteDeviceGroup)

		//devices
        api.POST("/devices", PostDevice)
        api.GET("/devices", GetDevices)
		api.GET("/devices/:id", GetDevice)
        api.PUT("/devices/:id", UpdateDevice)
		api.DELETE("/devices/:id", DeleteDevice)

		//device logins
        api.POST("/deviceLogins", PostDeviceLogin)
        api.GET("/deviceLogins", GetDeviceLogins)
		api.GET("/deviceLogins/:id", GetDeviceLogin)
        api.PUT("/deviceLogins/:id", UpdateDeviceLogin)
		api.DELETE("/deviceLogins/:id", DeleteDeviceLogin)

		//gps datas
        api.POST("/gpsDatas", PostGpsData)
        api.GET("/gpsDatas", GetGpsDatas)
		api.GET("/gpsDatas/:id", GetGpsData)
        api.PUT("/gpsDatas/:id", UpdateGpsData)
		api.DELETE("/gpsDatas/:id", DeleteGpsData)
		
		//temperatures
        api.POST("/temperatures", PostTemperature)
        api.GET("/temperatures", GetTemperatures)
		api.GET("/temperatures/:id", GetTemperature)
        api.PUT("/temperatures/:id", UpdateTemperature)
		api.DELETE("/temperatures/:id", DeleteTemperature)

		//alarms
        api.POST("/alarms", PostAlarm)
        api.GET("/alarms", GetAlarms)
		api.GET("/alarms/:id", GetAlarm)
        api.PUT("/alarms/:id", UpdateAlarm)
		api.DELETE("/alarms/:id", DeleteAlarm)

		//alarm types
		api.GET("/alarmTypes", GetAlarmTypes)		
    }

	e.Logger.Fatal(e.Start(":9090"))
}

