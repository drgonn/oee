package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
func fetchSingleUserlog(c *gin.Context){
	var userlog Userlog
	userlogID := c.Param("id")
	db.First(&userlog, userlogID)
	if userlog.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No userlog found!"})
		return
	}
	_userlog :=  jsonUserlog{
		ID : userlog.ID,
		Ip : userlog.Ip,
		User_Agent : userlog.User_Agent,
		Msg : userlog.Msg,
		Time : userlog.Time,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _userlog})
}

func createUserlog(c *gin.Context){
	var userlog jsonUserlog
	err := c.BindJSON(&userlog)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case userlog.Ip == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"ip 参数缺失"})
		return
	case userlog.UserID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"UserID 参数缺失"})
		return
	}
	var queryUserlog Userlog
	db.Where("ip = ?", userlog.Ip).First(&queryUserlog)
	if queryUserlog.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Userlog ip 已经存在，不允许重复"})
		return
	}
	var queryUserlog Userlog
	db.Where("user_agent = ?", userlog.User_Agent).First(&queryUserlog)
	if queryUserlog.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Userlog user_agent 已经存在，不允许重复"})
		return
	}
	var queryUserlog Userlog
	db.Where("msg = ?", userlog.Msg).First(&queryUserlog)
	if queryUserlog.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Userlog msg 已经存在，不允许重复"})
		return
	}
	dbUserlog := Userlog{
		Ip : userlog.Ip,
	}
	db.Save(&dbUserlog)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbUserlog.ID, 
	})
}

func updateUserlog(c *gin.Context){
	var userlog Userlog
	userlogID := c.Param("id")
	db.First(&userlog, userlogID)
	if userlog.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No userlog found!"})
		return
	}
	var juserlog jsonUserlog
	err := c.BindJSON(&juserlog)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if juserlog.Ip != "" {
		var queryUserlog Userlog
		db.Where("Ip = ?", juserlog.Ip).First(&queryUserlog)
		if queryUserlog.ID != 0 && queryUserlog.Ip != userlog.Ip {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Userlog Ip 已经存在，不允许重复"})
			return
		}
		db.Model(&userlog).Update("Ip",juserlog.Ip)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":userlog.ID, 
	})
}

func deleteUserlog(c *gin.Context){
	var userlog Userlog
	userlogID := c.Param("id")
	db.First(&userlog, userlogID)
	if userlog.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No userlog found!"})
		return
	}
	db.Delete(&userlog)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageUserlog struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonUserlog
	}

type rpageUserlog struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonUserlog `json:"records"`
	}

func fetchPageUserlog(c *gin.Context){
	var totalcount int
	var pagecount int
	var userlog pageUserlog
	var _jUserlogs []jsonUserlog
	Db := db
	itemUserlogs := make([]Userlog,0)
	err := c.BindJSON(&userlog)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case userlog.Pageindex == 0:
		userlog.Pageindex = 1
		fallthrough
	case userlog.Pagesize == 0:
		userlog.Pagesize = 20
	}
	if userlog.Sortfield != "" {
		var build strings.Builder
		build.WriteString(userlog.Sortfield)
		if userlog.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemUserlogs).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Userlog 表数据查询错误"})
		return
	}

	if userlog.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % userlog.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / userlog.Pagesize + yu
	}
	Db = Db.Limit(userlog.Pagesize).Offset((userlog.Pageindex - 1)* userlog.Pagesize)
	if err := Db.Find(&itemUserlogs).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Userlog 表数据查询错误"})
		return
	}

	for _, item := range itemUserlogs {
		_jUserlogs = append(_jUserlogs, jsonUserlog{
			ID : item.ID,
			Ip : item.Ip,
			User_Agent : item.User_Agent,
			Msg : item.Msg,
			Time : item.Time,
			UserID : item.UserID,
		})
	}
	_page := rpageUserlog{
		Pageindex : userlog.Pageindex,
		Pagesize : userlog.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jUserlogs,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

