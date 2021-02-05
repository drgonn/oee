package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
func fetchSingleBug(c *gin.Context){
	var bug Bug
	bugID := c.Param("id")
	db.First(&bug, bugID)
	if bug.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No bug found!"})
		return
	}
	_bug :=  jsonBug{
		ID : bug.ID,
		Reason : bug.Reason,
		Start_Time : bug.Start_Time,
		End_Time : bug.End_Time,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _bug})
}

func createBug(c *gin.Context){
	var bug jsonBug
	err := c.BindJSON(&bug)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case bug.Reason == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"reason 参数缺失"})
		return
	case bug.Start_Time == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"start_time 参数缺失"})
		return
	case bug.End_Time == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"end_time 参数缺失"})
		return
	case bug.DeviceID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"DeviceID 参数缺失"})
		return
	case bug.BugtypeID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"BugtypeID 参数缺失"})
		return
	}
	var queryBug Bug
	db.Where("reason = ?", bug.Reason).First(&queryBug)
	if queryBug.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bug reason 已经存在，不允许重复"})
		return
	}
	var device Device
	deviceID := bug.DeviceID
	db.First(&device, deviceID)
	if device.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
	}
	var bugtype Bugtype
	bugtypeID := bug.BugtypeID
	db.First(&bugtype, bugtypeID)
	if bugtype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Bugtype found!"})
		return
	}
	dbBug := Bug{
		Reason : bug.Reason,
		Start_Time : bug.Start_Time,
		End_Time : bug.End_Time,
		DeviceID : bug.DeviceID,
		BugtypeID : bug.BugtypeID,
	}
	db.Save(&dbBug)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbBug.ID, 
	})
}

func updateBug(c *gin.Context){
	var bug Bug
	bugID := c.Param("id")
	db.First(&bug, bugID)
	if bug.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No bug found!"})
		return
	}
	var jbug jsonBug
	err := c.BindJSON(&jbug)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jbug.DeviceID != 0 {
		var device Device
		deviceID := jbug.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
			return
		}
	db.Model(&bug).Update("deviceID",bug.DeviceID)
	}
	if jbug.BugtypeID != 0 {
		var bugtype Bugtype
		bugtypeID := jbug.BugtypeID
		db.First(&bugtype, bugtypeID)
		if bugtype.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Bugtype found!"})
			return
		}
	db.Model(&bug).Update("bugtypeID",bug.BugtypeID)
	}
	if jbug.Reason != "" {
		var queryBug Bug
		db.Where("Reason = ?", jbug.Reason).First(&queryBug)
		if queryBug.ID != 0 && queryBug.Reason != bug.Reason {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bug Reason 已经存在，不允许重复"})
			return
		}
		db.Model(&bug).Update("Reason",jbug.Reason)
	}
	if jbug.Start_Time != "" {
		db.Model(&bug).Update("Start_Time",jbug.Start_Time)
	}
	if jbug.End_Time != "" {
		db.Model(&bug).Update("End_Time",jbug.End_Time)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":bug.ID, 
	})
}

func deleteBug(c *gin.Context){
	var bug Bug
	bugID := c.Param("id")
	db.First(&bug, bugID)
	if bug.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No bug found!"})
		return
	}
	db.Delete(&bug)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageBug struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonBug
	}

type rpageBug struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonBug `json:"records"`
	}

func fetchPageBug(c *gin.Context){
	var totalcount int
	var pagecount int
	var bug pageBug
	var _jBugs []jsonBug
	Db := db
	itemBugs := make([]Bug,0)
	err := c.BindJSON(&bug)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case bug.Pageindex == 0:
		bug.Pageindex = 1
		fallthrough
	case bug.Pagesize == 0:
		bug.Pagesize = 20
	}
	if bug.DeviceID != 0 { 
		var device Device
		deviceID := bug.DeviceID
		db.First(&device, deviceID)
		if device.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Device found!"})
		return
		} else {
			Db = Db.Where("deviceID = ?",bug.DeviceID)
		}
	}
	if bug.BugtypeID != 0 { 
		var bugtype Bugtype
		bugtypeID := bug.BugtypeID
		db.First(&bugtype, bugtypeID)
		if bugtype.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Bugtype found!"})
		return
		} else {
			Db = Db.Where("bugtypeID = ?",bug.BugtypeID)
		}
	}
	if bug.Sortfield != "" {
		var build strings.Builder
		build.WriteString(bug.Sortfield)
		if bug.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemBugs).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bug 表数据查询错误"})
		return
	}

	if bug.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % bug.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / bug.Pagesize + yu
	}
	Db = Db.Limit(bug.Pagesize).Offset((bug.Pageindex - 1)* bug.Pagesize)
	if err := Db.Find(&itemBugs).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bug 表数据查询错误"})
		return
	}

	for _, item := range itemBugs {
		_jBugs = append(_jBugs, jsonBug{
			ID : item.ID,
			Reason : item.Reason,
			Start_Time : item.Start_Time,
			End_Time : item.End_Time,
			DeviceID : item.DeviceID,
			BugtypeID : item.BugtypeID,
		})
	}
	_page := rpageBug{
		Pageindex : bug.Pageindex,
		Pagesize : bug.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jBugs,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

