package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
func fetchSingleMessage(c *gin.Context){
	var message Message
	messageID := c.Param("id")
	db.First(&message, messageID)
	if message.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No message found!"})
		return
	}
	_message :=  jsonMessage{
		ID : message.ID,
		Title : message.Title,
		Type : message.Type,
		Description : message.Description,
		Read : message.Read,
		Ct_Time : message.Ct_Time,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _message})
}

func createMessage(c *gin.Context){
	var message jsonMessage
	err := c.BindJSON(&message)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case message.Title == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"title 参数缺失"})
		return
	case message.Type == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"type 参数缺失"})
		return
	case message.Description == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"description 参数缺失"})
		return
	case message.UserID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"UserID 参数缺失"})
		return
	}
	var queryMessage Message
	db.Where("title = ?", message.Title).First(&queryMessage)
	if queryMessage.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Message title 已经存在，不允许重复"})
		return
	}
	var queryMessage Message
	db.Where("type = ?", message.Type).First(&queryMessage)
	if queryMessage.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Message type 已经存在，不允许重复"})
		return
	}
	var queryMessage Message
	db.Where("description = ?", message.Description).First(&queryMessage)
	if queryMessage.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Message description 已经存在，不允许重复"})
		return
	}
	dbMessage := Message{
		Title : message.Title,
		Type : message.Type,
		Description : message.Description,
	}
	db.Save(&dbMessage)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbMessage.ID, 
	})
}

func updateMessage(c *gin.Context){
	var message Message
	messageID := c.Param("id")
	db.First(&message, messageID)
	if message.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No message found!"})
		return
	}
	var jmessage jsonMessage
	err := c.BindJSON(&jmessage)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jmessage.Title != "" {
		var queryMessage Message
		db.Where("Title = ?", jmessage.Title).First(&queryMessage)
		if queryMessage.ID != 0 && queryMessage.Title != message.Title {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Message Title 已经存在，不允许重复"})
			return
		}
		db.Model(&message).Update("Title",jmessage.Title)
	}
	if jmessage.Type != 0 {
		var queryMessage Message
		db.Where("Type = ?", jmessage.Type).First(&queryMessage)
		if queryMessage.ID != 0 && queryMessage.Type != message.Type {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Message Type 已经存在，不允许重复"})
			return
		}
		db.Model(&message).Update("Type",jmessage.Type)
	}
	if jmessage.Read != false {
		db.Model(&message).Update("Read",jmessage.Read)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":message.ID, 
	})
}

func deleteMessage(c *gin.Context){
	var message Message
	messageID := c.Param("id")
	db.First(&message, messageID)
	if message.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No message found!"})
		return
	}
	db.Delete(&message)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageMessage struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonMessage
	}

type rpageMessage struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonMessage `json:"records"`
	}

func fetchPageMessage(c *gin.Context){
	var totalcount int
	var pagecount int
	var message pageMessage
	var _jMessages []jsonMessage
	Db := db
	itemMessages := make([]Message,0)
	err := c.BindJSON(&message)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case message.Pageindex == 0:
		message.Pageindex = 1
		fallthrough
	case message.Pagesize == 0:
		message.Pagesize = 20
	}
	if message.Sortfield != "" {
		var build strings.Builder
		build.WriteString(message.Sortfield)
		if message.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemMessages).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Message 表数据查询错误"})
		return
	}

	if message.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % message.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / message.Pagesize + yu
	}
	Db = Db.Limit(message.Pagesize).Offset((message.Pageindex - 1)* message.Pagesize)
	if err := Db.Find(&itemMessages).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Message 表数据查询错误"})
		return
	}

	for _, item := range itemMessages {
		_jMessages = append(_jMessages, jsonMessage{
			ID : item.ID,
			Title : item.Title,
			Type : item.Type,
			Description : item.Description,
			Read : item.Read,
			Ct_Time : item.Ct_Time,
			UserID : item.UserID,
		})
	}
	_page := rpageMessage{
		Pageindex : message.Pageindex,
		Pagesize : message.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jMessages,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

