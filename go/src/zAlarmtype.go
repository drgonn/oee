package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func fetchSingleAlarmtype(c *gin.Context){
	var alarmtype Alarmtype
	alarmtypeID := c.Param("id")
	db.First(&alarmtype, alarmtypeID)
	if alarmtype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No alarmtype found!"})
		return
	}
	_alarmtype :=  jsonAlarmtype{
		ID : alarmtype.ID,
		Code : alarmtype.Code,
		Mean : alarmtype.Mean,
		Cause : alarmtype.Cause,
		Solution : alarmtype.Solution,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _alarmtype})
}

func createAlarmtype(c *gin.Context){
	var alarmtype jsonAlarmtype
	err := c.BindJSON(&alarmtype)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case alarmtype.Code == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"code 参数缺失"})
		return
	case alarmtype.Mean == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"mean 参数缺失"})
		return
	case alarmtype.Cause == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"cause 参数缺失"})
		return
	case alarmtype.Solution == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"solution 参数缺失"})
		return
	}
	dbAlarmtype := Alarmtype{
		Code : alarmtype.Code,
		Mean : alarmtype.Mean,
		Cause : alarmtype.Cause,
		Solution : alarmtype.Solution,
	}
	db.Save(&dbAlarmtype)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbAlarmtype.ID, 
	})
}

func updateAlarmtype(c *gin.Context){
	var alarmtype Alarmtype
	alarmtypeID := c.Param("id")
	db.First(&alarmtype, alarmtypeID)
	if alarmtype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No alarmtype found!"})
		return
	}
	var jalarmtype jsonAlarmtype
	err := c.BindJSON(&jalarmtype)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jalarmtype.Code != "" {
		db.Model(&alarmtype).Update("Code",jalarmtype.Code)
	}
	if jalarmtype.Mean != "" {
		db.Model(&alarmtype).Update("Mean",jalarmtype.Mean)
	}
	if jalarmtype.Cause != "" {
		db.Model(&alarmtype).Update("Cause",jalarmtype.Cause)
	}
	if jalarmtype.Solution != "" {
		db.Model(&alarmtype).Update("Solution",jalarmtype.Solution)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":alarmtype.ID, 
	})
}

func deleteAlarmtype(c *gin.Context){
	var alarmtype Alarmtype
	alarmtypeID := c.Param("id")
	db.First(&alarmtype, alarmtypeID)
	if alarmtype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No alarmtype found!"})
		return
	}
	var alarm Alarm
	db.Where("alarmtypeID",alarmtype.ID).First(&alarm)
	if alarm.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "alarmtype还拥有alarm，不能删除"})
		return
	}
	db.Delete(&alarmtype)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageAlarmtype struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonAlarmtype
	}

type rpageAlarmtype struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonAlarmtype `json:"records"`
	}

func fetchPageAlarmtype(c *gin.Context){
	var totalcount int
	var pagecount int
	var alarmtype pageAlarmtype
	var _jAlarmtypes []jsonAlarmtype
	Db := db
	itemAlarmtypes := make([]Alarmtype,0)
	err := c.BindJSON(&alarmtype)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case alarmtype.Pageindex == 0:
		alarmtype.Pageindex = 1
		fallthrough
	case alarmtype.Pagesize == 0:
		alarmtype.Pagesize = 20
	}
	if alarmtype.Sortfield != "" {
		var build strings.Builder
		build.WriteString(alarmtype.Sortfield)
		if alarmtype.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemAlarmtypes).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Alarmtype 表数据查询错误"})
		return
	}

	if alarmtype.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % alarmtype.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / alarmtype.Pagesize + yu
	}
	Db = Db.Limit(alarmtype.Pagesize).Offset((alarmtype.Pageindex - 1)* alarmtype.Pagesize)
	if err := Db.Find(&itemAlarmtypes).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Alarmtype 表数据查询错误"})
		return
	}

	for _, item := range itemAlarmtypes {
		_jAlarmtypes = append(_jAlarmtypes, jsonAlarmtype{
			ID : item.ID,
			Code : item.Code,
			Mean : item.Mean,
			Cause : item.Cause,
			Solution : item.Solution,
		})
	}
	_page := rpageAlarmtype{
		Pageindex : alarmtype.Pageindex,
		Pagesize : alarmtype.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jAlarmtypes,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

