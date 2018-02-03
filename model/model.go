package model

import (
	"time"
)

type User struct {
	Id        int64  `db:"id" json:"id"`
	Username 	string `db:"username" json:"username"`
	Password  	string `db:"password" json:"password"`
	Fullname  	string `db:"fullname" json:"fullname"`
	Deleted  	byte `db:"deleted" json:"deleted"`
	IsAdmin  	string `db:"is_admin" json:"isAdmin"`
	Email	  	string `db:"email" json:"email"`
}

type Group struct {
	Id        int64  `db:"id" json:"id"`
	Name  	string `db:"name" json:"name"`
}

type DeviceGroup struct {
	Id        int64  `db:"id" json:"id"`
	DeviceNumber  	int64  `db:"device_number" json:"deviceNumber"`
	GroupId  	int64  `db:"group_id" json:"groupId"`
}

type Alarm struct {
	Id        int64  `db:"id" json:"id"`
	TypeCode 	int64 `db:"type_code" json:"typeCode"`
	DeviceNumber  	string `db:"device_number" json:"deviceNumber"`
	CreatedAt  	time.Time `db:"created_at" json:"createdAt"`
}

type AlarmType struct {
	TypeCode        int64  `db:"type_code" json:"typeCode"`
	TypeName  	string `db:"type_name" json:"typeName"`
}

type Device struct {
	Id        int64  `db:"id" json:"id"`
	DeviceNumber  	string `db:"device_number" json:"deviceNumber"`
	Name  	string `db:"name" json:"name"`
	Latitude  	float64 `db:"latitude" json:"latitude"`
	Longitude  	float64 `db:"longitude" json:"longitude"`
	Direction  	float64 `db:"direction" json:"direction"`
	Speed  	float64 `db:"speed" json:"speed"`
	Mileage  	float64 `db:"mileage" json:"mileage"`
	CreatedAt  	time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt  	time.Time `db:"updated_at" json:"updatedAt"`
	LastLoginAt  	time.Time `db:"last_login_at" json:"lastLoginAt"`
	Deleted  	byte `db:"deleted" json:"deleted"`
}

type DeviceLogin struct {
	Id        int64  `db:"id" json:"id"`
	DeviceNumber  	string `db:"device_number" json:"deviceNumber"`
	CreatedAt  	time.Time `db:"created_at" json:"createdAt"`
}

type GpsData struct {
	Id        int64  `db:"id" json:"id"`
	DeviceNumber  	string `db:"device_number" json:"deviceNumber"`
	CreatedAt  	time.Time `db:"created_at" json:"createdAt"`
	GpsTime  	time.Time `db:"gps_time" json:"gpsTime"`
	IoState  	string `db:"io_state" json:"ioState"`
	Latitude  	float64 `db:"latitude" json:"latitude"`
	Longitude  	float64 `db:"longitude" json:"longitude"`
	Direction  	float64 `db:"direction" json:"direction"`
	Speed  	float64 `db:"speed" json:"speed"`
	Mileage  	float64 `db:"mileage" json:"mileage"`
}

type Temperature struct {
	Id        int64  `db:"id" json:"id"`
	DeviceNumber  	string `db:"device_number" json:"deviceNumber"`
	CreatedAt  	time.Time `db:"created_at" json:"createdAt"`
	Temperature1  	float64 `db:"temperature1" json:"temperature1"`
	Temperature2  	float64 `db:"temperature2" json:"temperature2"`
	Temperature3  	float64 `db:"temperature3" json:"temperature3"`
	Temperature4  	float64 `db:"temperature4" json:"temperature4"`
}