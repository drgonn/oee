package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
func fetchSingleDevicestatu(c *gin.Context){
	var devicestatu Devicestatu
	devicestatuID := c.Param("id")
	db.First(&devicestatu, devicestatuID)
	if devicestatu.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No devicestatu found!"})
		return
	}
	_devicestatu :=  jsonDevicestatu{
		ID : devicestatu.ID,
		Start_Time : devicestatu.Start_Time,
		Status : devicestatu.Status,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _devicestatu})
}

func createDevicestatu(c *gin.Context){
	var devicestatu jsonDevicestatu
	err := c.BindJSON(&devicestatu)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case devicestatu.Start_Time == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"start_time 参数缺失"})
		return
	case devicestatu.Status == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"status 参数缺失"})
		return
	case devicestatu.DeviceID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"DeviceID 参数缺失"})
		return
	}
	var device Device
	deviceID := devicestatu.DeviceID
	db.First(&device, deviceID)
	if device.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
	}
	dbDevicestatu := Devicestatu{
		Start_Time : devicestatu.Start_Time,
		Status : devicestatu.Status,
		DeviceID : devicestatu.DeviceID,
	}
	db.Save(&dbDevicestatu)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbDevicestatu.ID, 
	})
}

func updateDevicestatu(c *gin.Context){
	var devicestatu Devicestatu
	devicestatuID := c.Param("id")
	db.First(&devicestatu, devicestatuID)
	if devicestatu.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No devicestatu found!"})
		return
	}
	var jdevicestatu jsonDevicestatu
	err := c.BindJSON(&jdevicestatu)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jdevicestatu.DeviceID != 0 {
		var device Device
		deviceID := jdevicestatu.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
			return
		}
	db.Model(&devicestatu).Update("deviceID",devicestatu.DeviceID)
	}
	if jdevicestatu.Start_Time != "" {
		db.Model(&devicestatu).Update("Start_Time",jdevicestatu.Start_Time)
	}
	if jdevicestatu.Status != 0 {
		db.Model(&devicestatu).Update("Status",jdevicestatu.Status)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":devicestatu.ID, 
	})
}

func deleteDevicestatu(c *gin.Context){
	var devicestatu Devicestatu
	devicestatuID := c.Param("id")
	db.First(&devicestatu, devicestatuID)
	if devicestatu.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No devicestatu found!"})
		return
	}
	db.Delete(&devicestatu)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageDevicestatu struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonDevicestatu
	}

type rpageDevicestatu struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonDevicestatu `json:"records"`
	}

func fetchPageDevicestatu(c *gin.Context){
	var totalcount int
	var pagecount int
	var devicestatu pageDevicestatu
	var _jDevicestatus []jsonDevicestatu
	Db := db
	itemDevicestatus := make([]Devicestatu,0)
	err := c.BindJSON(&devicestatu)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case devicestatu.Pageindex == 0:
		devicestatu.Pageindex = 1
		fallthrough
	case devicestatu.Pagesize == 0:
		devicestatu.Pagesize = 20
	}
	if devicestatu.DeviceID != 0 { 
		var device Device
		deviceID := devicestatu.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
		} else {
			Db = Db.Where("deviceID = ?",devicestatu.DeviceID)
		}
	}
	if devicestatu.Sortfield != "" {
		var build strings.Builder
		build.WriteString(devicestatu.Sortfield)
		if devicestatu.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemDevicestatus).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Devicestatu 表数据查询错误"})
		return
	}

	if devicestatu.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % devicestatu.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / devicestatu.Pagesize + yu
	}
	Db = Db.Limit(devicestatu.Pagesize).Offset((devicestatu.Pageindex - 1)* devicestatu.Pagesize)
	if err := Db.Find(&itemDevicestatus).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Devicestatu 表数据查询错误"})
		return
	}

	for _, item := range itemDevicestatus {
		_jDevicestatus = append(_jDevicestatus, jsonDevicestatu{
			ID : item.ID,
			Start_Time : item.Start_Time,
			Status : item.Status,
			DeviceID : item.DeviceID,
		})
	}
	_page := rpageDevicestatu{
		Pageindex : devicestatu.Pageindex,
		Pagesize : devicestatu.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jDevicestatus,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

