package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func fetchSinglePermission(c *gin.Context){
	var permission Permission
	permissionID := c.Param("id")
	db.First(&permission, permissionID)
	if permission.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No permission found!"})
		return
	}
	_permission :=  jsonPermission{
		ID : permission.ID,
		Name : permission.Name,
		About : permission.About,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _permission})
}

func createPermission(c *gin.Context){
	var permission jsonPermission
	err := c.BindJSON(&permission)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case permission.Name == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"name 参数缺失"})
		return
	case permission.About == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"about 参数缺失"})
		return
	}
	var queryPermission Permission
	db.Where("name = ?", permission.Name).First(&queryPermission)
	if queryPermission.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Permission name 已经存在，不允许重复"})
		return
	}
	var queryPermission Permission
	db.Where("about = ?", permission.About).First(&queryPermission)
	if queryPermission.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Permission about 已经存在，不允许重复"})
		return
	}
	dbPermission := Permission{
		Name : permission.Name,
		About : permission.About,
	}
	db.Save(&dbPermission)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbPermission.ID, 
	})
}

func updatePermission(c *gin.Context){
	var permission Permission
	permissionID := c.Param("id")
	db.First(&permission, permissionID)
	if permission.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No permission found!"})
		return
	}
	var jpermission jsonPermission
	err := c.BindJSON(&jpermission)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jpermission.Name != "" {
		var queryPermission Permission
		db.Where("Name = ?", jpermission.Name).First(&queryPermission)
		if queryPermission.ID != 0 && queryPermission.Name != permission.Name {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Permission Name 已经存在，不允许重复"})
			return
		}
		db.Model(&permission).Update("Name",jpermission.Name)
	}
	if jpermission.About != "" {
		var queryPermission Permission
		db.Where("About = ?", jpermission.About).First(&queryPermission)
		if queryPermission.ID != 0 && queryPermission.About != permission.About {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Permission About 已经存在，不允许重复"})
			return
		}
		db.Model(&permission).Update("About",jpermission.About)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":permission.ID, 
	})
}

func deletePermission(c *gin.Context){
	var permission Permission
	permissionID := c.Param("id")
	db.First(&permission, permissionID)
	if permission.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No permission found!"})
		return
	}
	db.Delete(&permission)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pagePermission struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonPermission
	}

type rpagePermission struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonPermission `json:"records"`
	}

func fetchPagePermission(c *gin.Context){
	var totalcount int
	var pagecount int
	var permission pagePermission
	var _jPermissions []jsonPermission
	Db := db
	itemPermissions := make([]Permission,0)
	err := c.BindJSON(&permission)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case permission.Pageindex == 0:
		permission.Pageindex = 1
		fallthrough
	case permission.Pagesize == 0:
		permission.Pagesize = 20
	}
	if permission.Sortfield != "" {
		var build strings.Builder
		build.WriteString(permission.Sortfield)
		if permission.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemPermissions).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Permission 表数据查询错误"})
		return
	}

	if permission.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % permission.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / permission.Pagesize + yu
	}
	Db = Db.Limit(permission.Pagesize).Offset((permission.Pageindex - 1)* permission.Pagesize)
	if err := Db.Find(&itemPermissions).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Permission 表数据查询错误"})
		return
	}

	for _, item := range itemPermissions {
		_jPermissions = append(_jPermissions, jsonPermission{
			ID : item.ID,
			Name : item.Name,
			About : item.About,
		})
	}
	_page := rpagePermission{
		Pageindex : permission.Pageindex,
		Pagesize : permission.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jPermissions,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

