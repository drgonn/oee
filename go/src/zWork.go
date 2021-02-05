package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
func fetchSingleWork(c *gin.Context){
	var work Work
	workID := c.Param("id")
	db.First(&work, workID)
	if work.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No work found!"})
		return
	}
	_work :=  jsonWork{
		ID : work.ID,
		Start_Time : work.Start_Time,
		End_Time : work.End_Time,
		Seconds : work.Seconds,
		Type : work.Type,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _work})
}

func createWork(c *gin.Context){
	var work jsonWork
	err := c.BindJSON(&work)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case work.Start_Time == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"start_time 参数缺失"})
		return
	case work.End_Time == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"end_time 参数缺失"})
		return
	case work.Seconds == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"seconds 参数缺失"})
		return
	case work.Type == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"type 参数缺失"})
		return
	case work.DeviceID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"DeviceID 参数缺失"})
		return
	}
	var device Device
	deviceID := work.DeviceID
	db.First(&device, deviceID)
	if device.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
	}
	dbWork := Work{
		Start_Time : work.Start_Time,
		End_Time : work.End_Time,
		Seconds : work.Seconds,
		Type : work.Type,
		DeviceID : work.DeviceID,
	}
	db.Save(&dbWork)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbWork.ID, 
	})
}

func updateWork(c *gin.Context){
	var work Work
	workID := c.Param("id")
	db.First(&work, workID)
	if work.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No work found!"})
		return
	}
	var jwork jsonWork
	err := c.BindJSON(&jwork)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jwork.DeviceID != 0 {
		var device Device
		deviceID := jwork.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
			return
		}
	db.Model(&work).Update("deviceID",work.DeviceID)
	}
	if jwork.Start_Time != "" {
		db.Model(&work).Update("Start_Time",jwork.Start_Time)
	}
	if jwork.End_Time != "" {
		db.Model(&work).Update("End_Time",jwork.End_Time)
	}
	if jwork.Seconds != 0 {
		db.Model(&work).Update("Seconds",jwork.Seconds)
	}
	if jwork.Type != 0 {
		db.Model(&work).Update("Type",jwork.Type)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":work.ID, 
	})
}

func deleteWork(c *gin.Context){
	var work Work
	workID := c.Param("id")
	db.First(&work, workID)
	if work.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No work found!"})
		return
	}
	db.Delete(&work)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageWork struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonWork
	}

type rpageWork struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonWork `json:"records"`
	}

func fetchPageWork(c *gin.Context){
	var totalcount int
	var pagecount int
	var work pageWork
	var _jWorks []jsonWork
	Db := db
	itemWorks := make([]Work,0)
	err := c.BindJSON(&work)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case work.Pageindex == 0:
		work.Pageindex = 1
		fallthrough
	case work.Pagesize == 0:
		work.Pagesize = 20
	}
	if work.DeviceID != 0 { 
		var device Device
		deviceID := work.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
		} else {
			Db = Db.Where("deviceID = ?",work.DeviceID)
		}
	}
	if work.Sortfield != "" {
		var build strings.Builder
		build.WriteString(work.Sortfield)
		if work.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemWorks).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Work 表数据查询错误"})
		return
	}

	if work.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % work.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / work.Pagesize + yu
	}
	Db = Db.Limit(work.Pagesize).Offset((work.Pageindex - 1)* work.Pagesize)
	if err := Db.Find(&itemWorks).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Work 表数据查询错误"})
		return
	}

	for _, item := range itemWorks {
		_jWorks = append(_jWorks, jsonWork{
			ID : item.ID,
			Start_Time : item.Start_Time,
			End_Time : item.End_Time,
			Seconds : item.Seconds,
			Type : item.Type,
			DeviceID : item.DeviceID,
		})
	}
	_page := rpageWork{
		Pageindex : work.Pageindex,
		Pagesize : work.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jWorks,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

