package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func fetchSingleValvetype(c *gin.Context){
	var valvetype Valvetype
	valvetypeID := c.Param("id")
	db.First(&valvetype, valvetypeID)
	if valvetype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No valvetype found!"})
		return
	}
	_valvetype :=  jsonValvetype{
		ID : valvetype.ID,
		Name : valvetype.Name,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _valvetype})
}

func createValvetype(c *gin.Context){
	var valvetype jsonValvetype
	err := c.BindJSON(&valvetype)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case valvetype.Name == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"name 参数缺失"})
		return
	}
	var queryValvetype Valvetype
	db.Where("name = ?", valvetype.Name).First(&queryValvetype)
	if queryValvetype.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valvetype name 已经存在，不允许重复"})
		return
	}
	dbValvetype := Valvetype{
		Name : valvetype.Name,
	}
	db.Save(&dbValvetype)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbValvetype.ID, 
	})
}

func updateValvetype(c *gin.Context){
	var valvetype Valvetype
	valvetypeID := c.Param("id")
	db.First(&valvetype, valvetypeID)
	if valvetype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No valvetype found!"})
		return
	}
	var jvalvetype jsonValvetype
	err := c.BindJSON(&jvalvetype)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jvalvetype.Name != "" {
		var queryValvetype Valvetype
		db.Where("Name = ?", jvalvetype.Name).First(&queryValvetype)
		if queryValvetype.ID != 0 && queryValvetype.Name != valvetype.Name {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valvetype Name 已经存在，不允许重复"})
			return
		}
		db.Model(&valvetype).Update("Name",jvalvetype.Name)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":valvetype.ID, 
	})
}

func deleteValvetype(c *gin.Context){
	var valvetype Valvetype
	valvetypeID := c.Param("id")
	db.First(&valvetype, valvetypeID)
	if valvetype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No valvetype found!"})
		return
	}
	var valve Valve
	db.Where("valvetypeID",valvetype.ID).First(&valve)
	if valve.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "valvetype还拥有valve，不能删除"})
		return
	}
	db.Delete(&valvetype)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageValvetype struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonValvetype
	}

type rpageValvetype struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonValvetype `json:"records"`
	}

func fetchPageValvetype(c *gin.Context){
	var totalcount int
	var pagecount int
	var valvetype pageValvetype
	var _jValvetypes []jsonValvetype
	Db := db
	itemValvetypes := make([]Valvetype,0)
	err := c.BindJSON(&valvetype)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case valvetype.Pageindex == 0:
		valvetype.Pageindex = 1
		fallthrough
	case valvetype.Pagesize == 0:
		valvetype.Pagesize = 20
	}
	if valvetype.Sortfield != "" {
		var build strings.Builder
		build.WriteString(valvetype.Sortfield)
		if valvetype.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemValvetypes).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valvetype 表数据查询错误"})
		return
	}

	if valvetype.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % valvetype.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / valvetype.Pagesize + yu
	}
	Db = Db.Limit(valvetype.Pagesize).Offset((valvetype.Pageindex - 1)* valvetype.Pagesize)
	if err := Db.Find(&itemValvetypes).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Valvetype 表数据查询错误"})
		return
	}

	for _, item := range itemValvetypes {
		_jValvetypes = append(_jValvetypes, jsonValvetype{
			ID : item.ID,
			Name : item.Name,
		})
	}
	_page := rpageValvetype{
		Pageindex : valvetype.Pageindex,
		Pagesize : valvetype.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jValvetypes,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

