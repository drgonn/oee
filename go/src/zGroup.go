package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func fetchSingleGroup(c *gin.Context){
	var group Group
	groupID := c.Param("id")
	db.First(&group, groupID)
	if group.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No group found!"})
		return
	}
	_group :=  jsonGroup{
		ID : group.ID,
		Name : group.Name,
		About : group.About,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _group})
}

func createGroup(c *gin.Context){
	var group jsonGroup
	err := c.BindJSON(&group)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case group.Name == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"name 参数缺失"})
		return
	case group.About == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"about 参数缺失"})
		return
	}
	var queryGroup Group
	db.Where("name = ?", group.Name).First(&queryGroup)
	if queryGroup.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Group name 已经存在，不允许重复"})
		return
	}
	var queryGroup Group
	db.Where("about = ?", group.About).First(&queryGroup)
	if queryGroup.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Group about 已经存在，不允许重复"})
		return
	}
	dbGroup := Group{
		Name : group.Name,
		About : group.About,
	}
	db.Save(&dbGroup)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbGroup.ID, 
	})
}

func updateGroup(c *gin.Context){
	var group Group
	groupID := c.Param("id")
	db.First(&group, groupID)
	if group.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No group found!"})
		return
	}
	var jgroup jsonGroup
	err := c.BindJSON(&jgroup)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jgroup.Name != "" {
		var queryGroup Group
		db.Where("Name = ?", jgroup.Name).First(&queryGroup)
		if queryGroup.ID != 0 && queryGroup.Name != group.Name {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Group Name 已经存在，不允许重复"})
			return
		}
		db.Model(&group).Update("Name",jgroup.Name)
	}
	if jgroup.About != "" {
		var queryGroup Group
		db.Where("About = ?", jgroup.About).First(&queryGroup)
		if queryGroup.ID != 0 && queryGroup.About != group.About {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Group About 已经存在，不允许重复"})
			return
		}
		db.Model(&group).Update("About",jgroup.About)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":group.ID, 
	})
}

func deleteGroup(c *gin.Context){
	var group Group
	groupID := c.Param("id")
	db.First(&group, groupID)
	if group.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No group found!"})
		return
	}
	db.Delete(&group)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageGroup struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonGroup
	}

type rpageGroup struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonGroup `json:"records"`
	}

func fetchPageGroup(c *gin.Context){
	var totalcount int
	var pagecount int
	var group pageGroup
	var _jGroups []jsonGroup
	Db := db
	itemGroups := make([]Group,0)
	err := c.BindJSON(&group)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case group.Pageindex == 0:
		group.Pageindex = 1
		fallthrough
	case group.Pagesize == 0:
		group.Pagesize = 20
	}
	if group.Sortfield != "" {
		var build strings.Builder
		build.WriteString(group.Sortfield)
		if group.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemGroups).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Group 表数据查询错误"})
		return
	}

	if group.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % group.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / group.Pagesize + yu
	}
	Db = Db.Limit(group.Pagesize).Offset((group.Pageindex - 1)* group.Pagesize)
	if err := Db.Find(&itemGroups).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Group 表数据查询错误"})
		return
	}

	for _, item := range itemGroups {
		_jGroups = append(_jGroups, jsonGroup{
			ID : item.ID,
			Name : item.Name,
			About : item.About,
		})
	}
	_page := rpageGroup{
		Pageindex : group.Pageindex,
		Pagesize : group.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jGroups,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

