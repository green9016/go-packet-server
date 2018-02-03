package server

import (
	"strconv"
	"fmt"
	"time"
	"../model"
	"../api"
)

func ProcessDeviceLogin(sn string, data string) {
	var db = api.InitDb()
	var device model.Device
	gpsData, err := ProcessGPSData(data[15:])
	fmt.Printf("return value=%v\n", gpsData)
	if(err == nil) {
		//save GPS Data
		gpsData.DeviceNumber = sn
		db.Create(&gpsData)
	}

	if(db.Where("device_number = ?", sn).First(&device).RecordNotFound()) {
		device = model.Device{DeviceNumber: sn, LastLoginAt:time.Now()}
		db.Create(&device)
	} else {
		device.LastLoginAt = time.Now() 
		//update gps data
		device.Latitude = gpsData.Latitude
		device.Longitude = gpsData.Longitude
		device.Direction = gpsData.Direction
		device.Speed = gpsData.Speed
		device.Mileage = gpsData.Mileage
		db.Save(&device)
	}

	//save device login data
	db.Create(&model.DeviceLogin{DeviceNumber:sn})
}

func ProcessLocation(sn string, data string) {
	var db = api.InitDb()
	var device model.Device
	gpsData, err := ProcessGPSData(data)
	if(err == nil) {
		//save GPS Data
		gpsData.DeviceNumber = sn
		db.Create(&gpsData)
	}

	if(db.Where("device_number = ?", sn).First(&device).RecordNotFound()) {
		device = model.Device{DeviceNumber: sn, LastLoginAt:time.Now()}
		db.Create(&device)
	} else {
		//update gps data
		device.Latitude = gpsData.Latitude
		device.Longitude = gpsData.Longitude
		device.Direction = gpsData.Direction
		device.Speed = gpsData.Speed
		device.Mileage = gpsData.Mileage
		db.Save(&device)
	}

	//save device login data
	db.Create(&model.DeviceLogin{DeviceNumber:sn})
}

func ProcessAlarm(sn string, data string) {
	var db = api.InitDb()
	var device model.Device
	alarm := data[0:1]

	alarmI, _ := strconv.ParseInt(alarm, 10, 64)
	//save alarm
	db.Create(&model.Alarm {TypeCode:alarmI,DeviceNumber:sn})

	gpsData, err := ProcessGPSData(data[1:])
	fmt.Printf("return value=%v\n", gpsData)
	if(err == nil) {
		//save GPS Data
		gpsData.DeviceNumber = sn
		db.Create(&gpsData)
	}

	if(db.Where("device_number = ?", sn).First(&device).RecordNotFound()) {
		device = model.Device{DeviceNumber: sn, LastLoginAt:time.Now()}
		db.Create(&device)
	} else {
		//update gps data
		device.Latitude = gpsData.Latitude
		device.Longitude = gpsData.Longitude
		device.Direction = gpsData.Direction
		device.Speed = gpsData.Speed
		device.Mileage = gpsData.Mileage
		db.Save(&device)
	}

	//save device login data
	db.Create(&model.DeviceLogin{DeviceNumber:sn})
}

func ProcessGPSData(gpsStr string) (*model.GpsData, error){
	//sample data=080524A2232.9806N11404.9355E000.1101241323.8700000000L000450AC
	var date = gpsStr[:6]
	var avail = gpsStr[6:7]
	fmt.Println("GPS Str", gpsStr, date, avail)
	if(avail == "A") {
		var latitude = gpsStr[7:16]
		var latitudeDir = gpsStr[16:17]
		var longitude = gpsStr[17:27]
		var longitudeDir = gpsStr[27:28]
		var speed = gpsStr[28:33]
		var timeStr = gpsStr[33:39]
		var orientation = gpsStr[39:45]
		var ioState = gpsStr[45:53]
		var milepost = gpsStr[53:54]
		

		//fmt.Println(date, latitude, latitudeDir, longitude, longitudeDir, speed, timeStr, orientation, ioState, milepost, mileage)

		gpsData := &model.GpsData{}

		latitudeF, _ := strconv.ParseFloat(latitude, 64)
		longitudeF, _ := strconv.ParseFloat(longitude, 64)
		speedF, _ := strconv.ParseFloat(speed, 64)
		orientationF, _ := strconv.ParseFloat(orientation, 64)
		
		if latitudeDir == "N" {
			gpsData.Latitude = latitudeF/100
		} else {
			gpsData.Latitude = latitudeF/100 * (-1)
		}

		if longitudeDir == "E" {
			gpsData.Longitude = longitudeF/100
		} else {
			gpsData.Longitude = longitudeF/100 * (-1)
		}

		gpsData.Speed = speedF
		gpsData.Direction = orientationF
		gpsData.Speed = speedF
		gpsData.IoState = ioState

		if milepost == "L" {
			var mileage = gpsStr[54:62]
			mileageF, _ := strconv.ParseFloat(mileage, 64)
			gpsData.Mileage = mileageF
		}

		layout := "20060102150405"
		t, err := time.Parse(layout, "20" + date + timeStr)
		if(err != nil) {
			fmt.Println(err)
		}
		gpsData.GpsTime = t

		fmt.Printf("%v\n", gpsData)
		return gpsData, nil
	} else if(avail == "V") {
		return nil, &errorString{"Wrong data"}
	}

	return nil, nil
}