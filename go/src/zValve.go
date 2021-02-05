package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func fetchSingleValve(c *gin.Context){
	var valve Valve
	valveID := c.Param("id")
	db.First(&valve, valveID)
	if valve.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No valve found!"})
		return
	}
	_valve :=  jsonValve{
		ID : valve.ID,
		Sn : valve.Sn,
		Name : valve.Name,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _valve})
}

func createValve(c *gin.Context){
	var valve jsonValve
	err := c.BindJSON(&valve)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case valve.Sn == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"sn 参数缺失"})
		return
	case valve.Name == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"name 参数缺失"})
		return
	case valve.ValvetypeID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"ValvetypeID 参数缺失"})
		return
	case valve.DeviceID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"DeviceID 参数缺失"})
		return
	}
	var queryValve Valve
	db.Where("sn = ?", valve.Sn).First(&queryValve)
	if queryValve.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valve sn 已经存在，不允许重复"})
		return
	}
	var queryValve Valve
	db.Where("name = ?", valve.Name).First(&queryValve)
	if queryValve.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valve name 已经存在，不允许重复"})
		return
	}
	var valvetype Valvetype
	valvetypeID := valve.ValvetypeID
	db.First(&valvetype, valvetypeID)
	if valvetype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Valvetype found!"})
		return
	}
	var device Device
	deviceID := valve.DeviceID
	db.First(&device, deviceID)
	if device.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
	}
	dbValve := Valve{
		Sn : valve.Sn,
		Name : valve.Name,
		ValvetypeID : valve.ValvetypeID,
		DeviceID : valve.DeviceID,
	}
	db.Save(&dbValve)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbValve.ID, 
	})
}

func updateValve(c *gin.Context){
	var valve Valve
	valveID := c.Param("id")
	db.First(&valve, valveID)
	if valve.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No valve found!"})
		return
	}
	var jvalve jsonValve
	err := c.BindJSON(&jvalve)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jvalve.ValvetypeID != 0 {
		var valvetype Valvetype
		valvetypeID := jvalve.ValvetypeID
		db.First(&valvetype, valvetypeID)
		if valvetype.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Valvetype found!"})
			return
		}
	db.Model(&valve).Update("valvetypeID",valve.ValvetypeID)
	}
	if jvalve.DeviceID != 0 {
		var device Device
		deviceID := jvalve.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
			return
		}
	db.Model(&valve).Update("deviceID",valve.DeviceID)
	}
	if jvalve.Sn != "" {
		var queryValve Valve
		db.Where("Sn = ?", jvalve.Sn).First(&queryValve)
		if queryValve.ID != 0 && queryValve.Sn != valve.Sn {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valve Sn 已经存在，不允许重复"})
			return
		}
		db.Model(&valve).Update("Sn",jvalve.Sn)
	}
	if jvalve.Name != "" {
		var queryValve Valve
		db.Where("Name = ?", jvalve.Name).First(&queryValve)
		if queryValve.ID != 0 && queryValve.Name != valve.Name {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valve Name 已经存在，不允许重复"})
			return
		}
		db.Model(&valve).Update("Name",jvalve.Name)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":valve.ID, 
	})
}

func deleteValve(c *gin.Context){
	var valve Valve
	valveID := c.Param("id")
	db.First(&valve, valveID)
	if valve.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No valve found!"})
		return
	}
	var valvetime Valvetime
	db.Where("valveID",valve.ID).First(&valvetime)
	if valvetime.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "valve还拥有valvetime，不能删除"})
		return
	}
	var alarm Alarm
	db.Where("valveID",valve.ID).First(&alarm)
	if alarm.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "valve还拥有alarm，不能删除"})
		return
	}
	db.Delete(&valve)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageValve struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonValve
	}

type rpageValve struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonValve `json:"records"`
	}

func fetchPageValve(c *gin.Context){
	var totalcount int
	var pagecount int
	var valve pageValve
	var _jValves []jsonValve
	Db := db
	itemValves := make([]Valve,0)
	err := c.BindJSON(&valve)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case valve.Pageindex == 0:
		valve.Pageindex = 1
		fallthrough
	case valve.Pagesize == 0:
		valve.Pagesize = 20
	}
	if valve.ValvetypeID != 0 { 
		var valvetype Valvetype
		valvetypeID := valve.ValvetypeID
		db.First(&valvetype, valvetypeID)
		if valvetype.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Valvetype found!"})
		return
		} else {
			Db = Db.Where("valvetypeID = ?",valve.ValvetypeID)
		}
	}
	if valve.DeviceID != 0 { 
		var device Device
		deviceID := valve.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
		} else {
			Db = Db.Where("deviceID = ?",valve.DeviceID)
		}
	}
	if valve.Sortfield != "" {
		var build strings.Builder
		build.WriteString(valve.Sortfield)
		if valve.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemValves).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valve 表数据查询错误"})
		return
	}

	if valve.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % valve.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / valve.Pagesize + yu
	}
	Db = Db.Limit(valve.Pagesize).Offset((valve.Pageindex - 1)* valve.Pagesize)
	if err := Db.Find(&itemValves).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valve 表数据查询错误"})
		return
	}

	for _, item := range itemValves {
		_jValves = append(_jValves, jsonValve{
			ID : item.ID,
			Sn : item.Sn,
			Name : item.Name,
			ValvetypeID : item.ValvetypeID,
			DeviceID : item.DeviceID,
		})
	}
	_page := rpageValve{
		Pageindex : valve.Pageindex,
		Pagesize : valve.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jValves,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

