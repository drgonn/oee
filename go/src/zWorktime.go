package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
func fetchSingleWorktime(c *gin.Context){
	var worktime Worktime
	worktimeID := c.Param("id")
	db.First(&worktime, worktimeID)
	if worktime.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No worktime found!"})
		return
	}
	_worktime :=  jsonWorktime{
		ID : worktime.ID,
		Start_Time : worktime.Start_Time,
		End_Time : worktime.End_Time,
		Seconds : worktime.Seconds,
		Type : worktime.Type,
		Amount : worktime.Amount,
		Good : worktime.Good,
		Glue : worktime.Glue,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _worktime})
}

func createWorktime(c *gin.Context){
	var worktime jsonWorktime
	err := c.BindJSON(&worktime)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case worktime.Start_Time == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"start_time 参数缺失"})
		return
	case worktime.End_Time == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"end_time 参数缺失"})
		return
	case worktime.Seconds == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"seconds 参数缺失"})
		return
	case worktime.Type == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"type 参数缺失"})
		return
	case worktime.Amount == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"amount 参数缺失"})
		return
	case worktime.Good == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"good 参数缺失"})
		return
	case worktime.Glue == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"glue 参数缺失"})
		return
	case worktime.DeviceID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"DeviceID 参数缺失"})
		return
	}
	var device Device
	deviceID := worktime.DeviceID
	db.First(&device, deviceID)
	if device.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
	}
	dbWorktime := Worktime{
		Start_Time : worktime.Start_Time,
		End_Time : worktime.End_Time,
		Seconds : worktime.Seconds,
		Type : worktime.Type,
		Amount : worktime.Amount,
		Good : worktime.Good,
		Glue : worktime.Glue,
		DeviceID : worktime.DeviceID,
	}
	db.Save(&dbWorktime)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbWorktime.ID, 
	})
}

func updateWorktime(c *gin.Context){
	var worktime Worktime
	worktimeID := c.Param("id")
	db.First(&worktime, worktimeID)
	if worktime.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No worktime found!"})
		return
	}
	var jworktime jsonWorktime
	err := c.BindJSON(&jworktime)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jworktime.DeviceID != 0 {
		var device Device
		deviceID := jworktime.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
			return
		}
	db.Model(&worktime).Update("deviceID",worktime.DeviceID)
	}
	if jworktime.Start_Time != "" {
		db.Model(&worktime).Update("Start_Time",jworktime.Start_Time)
	}
	if jworktime.End_Time != "" {
		db.Model(&worktime).Update("End_Time",jworktime.End_Time)
	}
	if jworktime.Seconds != 0 {
		db.Model(&worktime).Update("Seconds",jworktime.Seconds)
	}
	if jworktime.Type != 0 {
		db.Model(&worktime).Update("Type",jworktime.Type)
	}
	if jworktime.Amount != 0 {
		db.Model(&worktime).Update("Amount",jworktime.Amount)
	}
	if jworktime.Good != 0 {
		db.Model(&worktime).Update("Good",jworktime.Good)
	}
	if jworktime.Glue != 0 {
		db.Model(&worktime).Update("Glue",jworktime.Glue)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":worktime.ID, 
	})
}

func deleteWorktime(c *gin.Context){
	var worktime Worktime
	worktimeID := c.Param("id")
	db.First(&worktime, worktimeID)
	if worktime.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No worktime found!"})
		return
	}
	db.Delete(&worktime)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageWorktime struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonWorktime
	}

type rpageWorktime struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonWorktime `json:"records"`
	}

func fetchPageWorktime(c *gin.Context){
	var totalcount int
	var pagecount int
	var worktime pageWorktime
	var _jWorktimes []jsonWorktime
	Db := db
	itemWorktimes := make([]Worktime,0)
	err := c.BindJSON(&worktime)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case worktime.Pageindex == 0:
		worktime.Pageindex = 1
		fallthrough
	case worktime.Pagesize == 0:
		worktime.Pagesize = 20
	}
	if worktime.DeviceID != 0 { 
		var device Device
		deviceID := worktime.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
		} else {
			Db = Db.Where("deviceID = ?",worktime.DeviceID)
		}
	}
	if worktime.Sortfield != "" {
		var build strings.Builder
		build.WriteString(worktime.Sortfield)
		if worktime.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemWorktimes).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Worktime 表数据查询错误"})
		return
	}

	if worktime.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % worktime.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / worktime.Pagesize + yu
	}
	Db = Db.Limit(worktime.Pagesize).Offset((worktime.Pageindex - 1)* worktime.Pagesize)
	if err := Db.Find(&itemWorktimes).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Worktime 表数据查询错误"})
		return
	}

	for _, item := range itemWorktimes {
		_jWorktimes = append(_jWorktimes, jsonWorktime{
			ID : item.ID,
			Start_Time : item.Start_Time,
			End_Time : item.End_Time,
			Seconds : item.Seconds,
			Type : item.Type,
			Amount : item.Amount,
			Good : item.Good,
			Glue : item.Glue,
			DeviceID : item.DeviceID,
		})
	}
	_page := rpageWorktime{
		Pageindex : worktime.Pageindex,
		Pagesize : worktime.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jWorktimes,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

