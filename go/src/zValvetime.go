package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
func fetchSingleValvetime(c *gin.Context){
	var valvetime Valvetime
	valvetimeID := c.Param("id")
	db.First(&valvetime, valvetimeID)
	if valvetime.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No valvetime found!"})
		return
	}
	_valvetime :=  jsonValvetime{
		ID : valvetime.ID,
		Start_Time : valvetime.Start_Time,
		End_Time : valvetime.End_Time,
		Seconds : valvetime.Seconds,
		Volt : valvetime.Volt,
		Amount : valvetime.Amount,
		Good : valvetime.Good,
		Glue : valvetime.Glue,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _valvetime})
}

func createValvetime(c *gin.Context){
	var valvetime jsonValvetime
	err := c.BindJSON(&valvetime)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case valvetime.Start_Time == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"start_time 参数缺失"})
		return
	case valvetime.End_Time == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"end_time 参数缺失"})
		return
	case valvetime.Seconds == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"seconds 参数缺失"})
		return
	case valvetime.Volt == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"volt 参数缺失"})
		return
	case valvetime.Amount == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"amount 参数缺失"})
		return
	case valvetime.Good == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"good 参数缺失"})
		return
	case valvetime.Glue == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"glue 参数缺失"})
		return
	case valvetime.DeviceID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"DeviceID 参数缺失"})
		return
	case valvetime.ValveID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"ValveID 参数缺失"})
		return
	}
	var device Device
	deviceID := valvetime.DeviceID
	db.First(&device, deviceID)
	if device.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
	}
	var valve Valve
	valveID := valvetime.ValveID
	db.First(&valve, valveID)
	if valve.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Valve found!"})
		return
	}
	dbValvetime := Valvetime{
		Start_Time : valvetime.Start_Time,
		End_Time : valvetime.End_Time,
		Seconds : valvetime.Seconds,
		Volt : valvetime.Volt,
		Amount : valvetime.Amount,
		Good : valvetime.Good,
		Glue : valvetime.Glue,
		DeviceID : valvetime.DeviceID,
		ValveID : valvetime.ValveID,
	}
	db.Save(&dbValvetime)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbValvetime.ID, 
	})
}

func updateValvetime(c *gin.Context){
	var valvetime Valvetime
	valvetimeID := c.Param("id")
	db.First(&valvetime, valvetimeID)
	if valvetime.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No valvetime found!"})
		return
	}
	var jvalvetime jsonValvetime
	err := c.BindJSON(&jvalvetime)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jvalvetime.DeviceID != 0 {
		var device Device
		deviceID := jvalvetime.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
			return
		}
	db.Model(&valvetime).Update("deviceID",valvetime.DeviceID)
	}
	if jvalvetime.ValveID != 0 {
		var valve Valve
		valveID := jvalvetime.ValveID
		db.First(&valve, valveID)
		if valve.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Valve found!"})
			return
		}
	db.Model(&valvetime).Update("valveID",valvetime.ValveID)
	}
	if jvalvetime.Start_Time != "" {
		db.Model(&valvetime).Update("Start_Time",jvalvetime.Start_Time)
	}
	if jvalvetime.End_Time != "" {
		db.Model(&valvetime).Update("End_Time",jvalvetime.End_Time)
	}
	if jvalvetime.Seconds != 0 {
		db.Model(&valvetime).Update("Seconds",jvalvetime.Seconds)
	}
	if jvalvetime.Volt != 0 {
		db.Model(&valvetime).Update("Volt",jvalvetime.Volt)
	}
	if jvalvetime.Amount != 0 {
		db.Model(&valvetime).Update("Amount",jvalvetime.Amount)
	}
	if jvalvetime.Good != 0 {
		db.Model(&valvetime).Update("Good",jvalvetime.Good)
	}
	if jvalvetime.Glue != 0 {
		db.Model(&valvetime).Update("Glue",jvalvetime.Glue)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":valvetime.ID, 
	})
}

func deleteValvetime(c *gin.Context){
	var valvetime Valvetime
	valvetimeID := c.Param("id")
	db.First(&valvetime, valvetimeID)
	if valvetime.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No valvetime found!"})
		return
	}
	db.Delete(&valvetime)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageValvetime struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonValvetime
	}

type rpageValvetime struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonValvetime `json:"records"`
	}

func fetchPageValvetime(c *gin.Context){
	var totalcount int
	var pagecount int
	var valvetime pageValvetime
	var _jValvetimes []jsonValvetime
	Db := db
	itemValvetimes := make([]Valvetime,0)
	err := c.BindJSON(&valvetime)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case valvetime.Pageindex == 0:
		valvetime.Pageindex = 1
		fallthrough
	case valvetime.Pagesize == 0:
		valvetime.Pagesize = 20
	}
	if valvetime.DeviceID != 0 { 
		var device Device
		deviceID := valvetime.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
		} else {
			Db = Db.Where("deviceID = ?",valvetime.DeviceID)
		}
	}
	if valvetime.ValveID != 0 { 
		var valve Valve
		valveID := valvetime.ValveID
		db.First(&valve, valveID)
		if valve.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Valve found!"})
		return
		} else {
			Db = Db.Where("valveID = ?",valvetime.ValveID)
		}
	}
	if valvetime.Sortfield != "" {
		var build strings.Builder
		build.WriteString(valvetime.Sortfield)
		if valvetime.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemValvetimes).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valvetime 表数据查询错误"})
		return
	}

	if valvetime.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % valvetime.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / valvetime.Pagesize + yu
	}
	Db = Db.Limit(valvetime.Pagesize).Offset((valvetime.Pageindex - 1)* valvetime.Pagesize)
	if err := Db.Find(&itemValvetimes).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valvetime 表数据查询错误"})
		return
	}

	for _, item := range itemValvetimes {
		_jValvetimes = append(_jValvetimes, jsonValvetime{
			ID : item.ID,
			Start_Time : item.Start_Time,
			End_Time : item.End_Time,
			Seconds : item.Seconds,
			Volt : item.Volt,
			Amount : item.Amount,
			Good : item.Good,
			Glue : item.Glue,
			DeviceID : item.DeviceID,
			ValveID : item.ValveID,
		})
	}
	_page := rpageValvetime{
		Pageindex : valvetime.Pageindex,
		Pagesize : valvetime.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jValvetimes,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

