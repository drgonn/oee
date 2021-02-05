package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func fetchSingleDevice(c *gin.Context){
	var device Device
	deviceID := c.Param("id")
	db.First(&device, deviceID)
	if device.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No device found!"})
		return
	}
	_device :=  jsonDevice{
		ID : device.ID,
		Sn : device.Sn,
		Name : device.Name,
		Ip : device.Ip,
		Img : device.Img,
		Type : device.Type,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _device})
}

func createDevice(c *gin.Context){
	var device jsonDevice
	err := c.BindJSON(&device)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case device.Sn == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"sn 参数缺失"})
		return
	case device.Name == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"name 参数缺失"})
		return
	case device.Ip == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"ip 参数缺失"})
		return
	case device.Img == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"img 参数缺失"})
		return
	case device.Type == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"type 参数缺失"})
		return
	}
	var queryDevice Device
	db.Where("sn = ?", device.Sn).First(&queryDevice)
	if queryDevice.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device sn 已经存在，不允许重复"})
		return
	}
	var queryDevice Device
	db.Where("name = ?", device.Name).First(&queryDevice)
	if queryDevice.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device name 已经存在，不允许重复"})
		return
	}
	var queryDevice Device
	db.Where("ip = ?", device.Ip).First(&queryDevice)
	if queryDevice.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device ip 已经存在，不允许重复"})
		return
	}
	var queryDevice Device
	db.Where("img = ?", device.Img).First(&queryDevice)
	if queryDevice.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device img 已经存在，不允许重复"})
		return
	}
	var queryDevice Device
	db.Where("type = ?", device.Type).First(&queryDevice)
	if queryDevice.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device type 已经存在，不允许重复"})
		return
	}
	dbDevice := Device{
		Sn : device.Sn,
		Name : device.Name,
		Ip : device.Ip,
		Img : device.Img,
		Type : device.Type,
	}
	db.Save(&dbDevice)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbDevice.ID, 
	})
}

func updateDevice(c *gin.Context){
	var device Device
	deviceID := c.Param("id")
	db.First(&device, deviceID)
	if device.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No device found!"})
		return
	}
	var jdevice jsonDevice
	err := c.BindJSON(&jdevice)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jdevice.Sn != "" {
		var queryDevice Device
		db.Where("Sn = ?", jdevice.Sn).First(&queryDevice)
		if queryDevice.ID != 0 && queryDevice.Sn != device.Sn {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device Sn 已经存在，不允许重复"})
			return
		}
		db.Model(&device).Update("Sn",jdevice.Sn)
	}
	if jdevice.Name != "" {
		var queryDevice Device
		db.Where("Name = ?", jdevice.Name).First(&queryDevice)
		if queryDevice.ID != 0 && queryDevice.Name != device.Name {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device Name 已经存在，不允许重复"})
			return
		}
		db.Model(&device).Update("Name",jdevice.Name)
	}
	if jdevice.Ip != "" {
		var queryDevice Device
		db.Where("Ip = ?", jdevice.Ip).First(&queryDevice)
		if queryDevice.ID != 0 && queryDevice.Ip != device.Ip {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device Ip 已经存在，不允许重复"})
			return
		}
		db.Model(&device).Update("Ip",jdevice.Ip)
	}
	if jdevice.Img != "" {
		var queryDevice Device
		db.Where("Img = ?", jdevice.Img).First(&queryDevice)
		if queryDevice.ID != 0 && queryDevice.Img != device.Img {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device Img 已经存在，不允许重复"})
			return
		}
		db.Model(&device).Update("Img",jdevice.Img)
	}
	if jdevice.Type != "" {
		var queryDevice Device
		db.Where("Type = ?", jdevice.Type).First(&queryDevice)
		if queryDevice.ID != 0 && queryDevice.Type != device.Type {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device Type 已经存在，不允许重复"})
			return
		}
		db.Model(&device).Update("Type",jdevice.Type)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":device.ID, 
	})
}

func deleteDevice(c *gin.Context){
	var device Device
	deviceID := c.Param("id")
	db.First(&device, deviceID)
	if device.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No device found!"})
		return
	}
	var worktime Worktime
	db.Where("deviceID",device.ID).First(&worktime)
	if worktime.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "device还拥有worktime，不能删除"})
		return
	}
	var devicestatu Devicestatu
	db.Where("deviceID",device.ID).First(&devicestatu)
	if devicestatu.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "device还拥有devicestatu，不能删除"})
		return
	}
	var valvetime Valvetime
	db.Where("deviceID",device.ID).First(&valvetime)
	if valvetime.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "device还拥有valvetime，不能删除"})
		return
	}
	var valve Valve
	db.Where("deviceID",device.ID).First(&valve)
	if valve.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "device还拥有valve，不能删除"})
		return
	}
	var bug Bug
	db.Where("deviceID",device.ID).First(&bug)
	if bug.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "device还拥有bug，不能删除"})
		return
	}
	var alarm Alarm
	db.Where("deviceID",device.ID).First(&alarm)
	if alarm.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "device还拥有alarm，不能删除"})
		return
	}
	db.Delete(&device)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageDevice struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonDevice
	}

type rpageDevice struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonDevice `json:"records"`
	}

func fetchPageDevice(c *gin.Context){
	var totalcount int
	var pagecount int
	var device pageDevice
	var _jDevices []jsonDevice
	Db := db
	itemDevices := make([]Device,0)
	err := c.BindJSON(&device)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case device.Pageindex == 0:
		device.Pageindex = 1
		fallthrough
	case device.Pagesize == 0:
		device.Pagesize = 20
	}
	if device.Sortfield != "" {
		var build strings.Builder
		build.WriteString(device.Sortfield)
		if device.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemDevices).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device 表数据查询错误"})
		return
	}

	if device.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % device.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / device.Pagesize + yu
	}
	Db = Db.Limit(device.Pagesize).Offset((device.Pageindex - 1)* device.Pagesize)
	if err := Db.Find(&itemDevices).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Device 表数据查询错误"})
		return
	}

	for _, item := range itemDevices {
		_jDevices = append(_jDevices, jsonDevice{
			ID : item.ID,
			Sn : item.Sn,
			Name : item.Name,
			Ip : item.Ip,
			Img : item.Img,
			Type : item.Type,
		})
	}
	_page := rpageDevice{
		Pageindex : device.Pageindex,
		Pagesize : device.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jDevices,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

