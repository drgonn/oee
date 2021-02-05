package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
func fetchSingleAlarm(c *gin.Context){
	var alarm Alarm
	alarmID := c.Param("id")
	db.First(&alarm, alarmID)
	if alarm.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No alarm found!"})
		return
	}
	_alarm :=  jsonAlarm{
		ID : alarm.ID,
		Create_Time : alarm.Create_Time,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _alarm})
}

func createAlarm(c *gin.Context){
	var alarm jsonAlarm
	err := c.BindJSON(&alarm)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case alarm.AlarmtypeID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"AlarmtypeID 参数缺失"})
		return
	case alarm.ValveID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"ValveID 参数缺失"})
		return
	case alarm.DeviceID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"DeviceID 参数缺失"})
		return
	}
	var alarmtype Alarmtype
	alarmtypeID := alarm.AlarmtypeID
	db.First(&alarmtype, alarmtypeID)
	if alarmtype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Alarmtype found!"})
		return
	}
	var valve Valve
	valveID := alarm.ValveID
	db.First(&valve, valveID)
	if valve.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Valve found!"})
		return
	}
	var device Device
	deviceID := alarm.DeviceID
	db.First(&device, deviceID)
	if device.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
	}
	dbAlarm := Alarm{
		Create_Time : time.Now(),
		AlarmtypeID : alarm.AlarmtypeID,
		ValveID : alarm.ValveID,
		DeviceID : alarm.DeviceID,
	}
	db.Save(&dbAlarm)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbAlarm.ID, 
	})
}

func updateAlarm(c *gin.Context){
	var alarm Alarm
	alarmID := c.Param("id")
	db.First(&alarm, alarmID)
	if alarm.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No alarm found!"})
		return
	}
	var jalarm jsonAlarm
	err := c.BindJSON(&jalarm)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jalarm.AlarmtypeID != 0 {
		var alarmtype Alarmtype
		alarmtypeID := jalarm.AlarmtypeID
		db.First(&alarmtype, alarmtypeID)
		if alarmtype.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Alarmtype found!"})
			return
		}
	db.Model(&alarm).Update("alarmtypeID",alarm.AlarmtypeID)
	}
	if jalarm.ValveID != 0 {
		var valve Valve
		valveID := jalarm.ValveID
		db.First(&valve, valveID)
		if valve.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Valve found!"})
			return
		}
	db.Model(&alarm).Update("valveID",alarm.ValveID)
	}
	if jalarm.DeviceID != 0 {
		var device Device
		deviceID := jalarm.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
			return
		}
	db.Model(&alarm).Update("deviceID",alarm.DeviceID)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":alarm.ID, 
	})
}

func deleteAlarm(c *gin.Context){
	var alarm Alarm
	alarmID := c.Param("id")
	db.First(&alarm, alarmID)
	if alarm.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No alarm found!"})
		return
	}
	db.Delete(&alarm)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageAlarm struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonAlarm
	}

type rpageAlarm struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonAlarm `json:"records"`
	}

func fetchPageAlarm(c *gin.Context){
	var totalcount int
	var pagecount int
	var alarm pageAlarm
	var _jAlarms []jsonAlarm
	Db := db
	itemAlarms := make([]Alarm,0)
	err := c.BindJSON(&alarm)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case alarm.Pageindex == 0:
		alarm.Pageindex = 1
		fallthrough
	case alarm.Pagesize == 0:
		alarm.Pagesize = 20
	}
	if alarm.AlarmtypeID != 0 { 
		var alarmtype Alarmtype
		alarmtypeID := alarm.AlarmtypeID
		db.First(&alarmtype, alarmtypeID)
		if alarmtype.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Alarmtype found!"})
		return
		} else {
			Db = Db.Where("alarmtypeID = ?",alarm.AlarmtypeID)
		}
	}
	if alarm.ValveID != 0 { 
		var valve Valve
		valveID := alarm.ValveID
		db.First(&valve, valveID)
		if valve.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Valve found!"})
		return
		} else {
			Db = Db.Where("valveID = ?",alarm.ValveID)
		}
	}
	if alarm.DeviceID != 0 { 
		var device Device
		deviceID := alarm.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
		} else {
			Db = Db.Where("deviceID = ?",alarm.DeviceID)
		}
	}
	if alarm.Sortfield != "" {
		var build strings.Builder
		build.WriteString(alarm.Sortfield)
		if alarm.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemAlarms).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Alarm 表数据查询错误"})
		return
	}

	if alarm.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % alarm.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / alarm.Pagesize + yu
	}
	Db = Db.Limit(alarm.Pagesize).Offset((alarm.Pageindex - 1)* alarm.Pagesize)
	if err := Db.Find(&itemAlarms).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Alarm 表数据查询错误"})
		return
	}

	for _, item := range itemAlarms {
		_jAlarms = append(_jAlarms, jsonAlarm{
			ID : item.ID,
			Create_Time : item.Create_Time,
			AlarmtypeID : item.AlarmtypeID,
			ValveID : item.ValveID,
			DeviceID : item.DeviceID,
		})
	}
	_page := rpageAlarm{
		Pageindex : alarm.Pageindex,
		Pagesize : alarm.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jAlarms,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

