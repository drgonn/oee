package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func fetchSingleRole(c *gin.Context){
	var role Role
	roleID := c.Param("id")
	db.First(&role, roleID)
	if role.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No role found!"})
		return
	}
	_role :=  jsonRole{
		ID : role.ID,
		Name : role.Name,
		Permissions : role.Permissions,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _role})
}

func createRole(c *gin.Context){
	var role jsonRole
	err := c.BindJSON(&role)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case role.Name == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"name 参数缺失"})
		return
	case role.Permissions == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"permissions 参数缺失"})
		return
	}
	var queryRole Role
	db.Where("name = ?", role.Name).First(&queryRole)
	if queryRole.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Role name 已经存在，不允许重复"})
		return
	}
	dbRole := Role{
		Name : role.Name,
		Permissions : role.Permissions,
	}
	db.Save(&dbRole)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbRole.ID, 
	})
}

func updateRole(c *gin.Context){
	var role Role
	roleID := c.Param("id")
	db.First(&role, roleID)
	if role.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No role found!"})
		return
	}
	var jrole jsonRole
	err := c.BindJSON(&jrole)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jrole.Name != "" {
		var queryRole Role
		db.Where("Name = ?", jrole.Name).First(&queryRole)
		if queryRole.ID != 0 && queryRole.Name != role.Name {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Role Name 已经存在，不允许重复"})
			return
		}
		db.Model(&role).Update("Name",jrole.Name)
	}
	if jrole.Permissions != 0 {
		db.Model(&role).Update("Permissions",jrole.Permissions)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":role.ID, 
	})
}

func deleteRole(c *gin.Context){
	var role Role
	roleID := c.Param("id")
	db.First(&role, roleID)
	if role.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No role found!"})
		return
	}
	var user User
	db.Where("roleID",role.ID).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "role还拥有user，不能删除"})
		return
	}
	db.Delete(&role)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageRole struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonRole
	}

type rpageRole struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonRole `json:"records"`
	}

func fetchPageRole(c *gin.Context){
	var totalcount int
	var pagecount int
	var role pageRole
	var _jRoles []jsonRole
	Db := db
	itemRoles := make([]Role,0)
	err := c.BindJSON(&role)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case role.Pageindex == 0:
		role.Pageindex = 1
		fallthrough
	case role.Pagesize == 0:
		role.Pagesize = 20
	}
	if role.Sortfield != "" {
		var build strings.Builder
		build.WriteString(role.Sortfield)
		if role.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemRoles).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Role 表数据查询错误"})
		return
	}

	if role.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % role.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / role.Pagesize + yu
	}
	Db = Db.Limit(role.Pagesize).Offset((role.Pageindex - 1)* role.Pagesize)
	if err := Db.Find(&itemRoles).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Role 表数据查询错误"})
		return
	}

	for _, item := range itemRoles {
		_jRoles = append(_jRoles, jsonRole{
			ID : item.ID,
			Name : item.Name,
			Permissions : item.Permissions,
		})
	}
	_page := rpageRole{
		Pageindex : role.Pageindex,
		Pagesize : role.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jRoles,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

