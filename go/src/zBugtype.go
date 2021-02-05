package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func fetchSingleBugtype(c *gin.Context){
	var bugtype Bugtype
	bugtypeID := c.Param("id")
	db.First(&bugtype, bugtypeID)
	if bugtype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No bugtype found!"})
		return
	}
	_bugtype :=  jsonBugtype{
		ID : bugtype.ID,
		Name : bugtype.Name,
		Sn : bugtype.Sn,
		About : bugtype.About,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _bugtype})
}

func createBugtype(c *gin.Context){
	var bugtype jsonBugtype
	err := c.BindJSON(&bugtype)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case bugtype.Name == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"name 参数缺失"})
		return
	case bugtype.Sn == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"sn 参数缺失"})
		return
	case bugtype.About == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"about 参数缺失"})
		return
	}
	var queryBugtype Bugtype
	db.Where("name = ?", bugtype.Name).First(&queryBugtype)
	if queryBugtype.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bugtype name 已经存在，不允许重复"})
		return
	}
	var queryBugtype Bugtype
	db.Where("sn = ?", bugtype.Sn).First(&queryBugtype)
	if queryBugtype.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bugtype sn 已经存在，不允许重复"})
		return
	}
	var queryBugtype Bugtype
	db.Where("about = ?", bugtype.About).First(&queryBugtype)
	if queryBugtype.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bugtype about 已经存在，不允许重复"})
		return
	}
	dbBugtype := Bugtype{
		Name : bugtype.Name,
		Sn : bugtype.Sn,
		About : bugtype.About,
	}
	db.Save(&dbBugtype)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbBugtype.ID, 
	})
}

func updateBugtype(c *gin.Context){
	var bugtype Bugtype
	bugtypeID := c.Param("id")
	db.First(&bugtype, bugtypeID)
	if bugtype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No bugtype found!"})
		return
	}
	var jbugtype jsonBugtype
	err := c.BindJSON(&jbugtype)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jbugtype.Name != "" {
		var queryBugtype Bugtype
		db.Where("Name = ?", jbugtype.Name).First(&queryBugtype)
		if queryBugtype.ID != 0 && queryBugtype.Name != bugtype.Name {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bugtype Name 已经存在，不允许重复"})
			return
		}
		db.Model(&bugtype).Update("Name",jbugtype.Name)
	}
	if jbugtype.Sn != "" {
		var queryBugtype Bugtype
		db.Where("Sn = ?", jbugtype.Sn).First(&queryBugtype)
		if queryBugtype.ID != 0 && queryBugtype.Sn != bugtype.Sn {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bugtype Sn 已经存在，不允许重复"})
			return
		}
		db.Model(&bugtype).Update("Sn",jbugtype.Sn)
	}
	if jbugtype.About != "" {
		var queryBugtype Bugtype
		db.Where("About = ?", jbugtype.About).First(&queryBugtype)
		if queryBugtype.ID != 0 && queryBugtype.About != bugtype.About {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bugtype About 已经存在，不允许重复"})
			return
		}
		db.Model(&bugtype).Update("About",jbugtype.About)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":bugtype.ID, 
	})
}

func deleteBugtype(c *gin.Context){
	var bugtype Bugtype
	bugtypeID := c.Param("id")
	db.First(&bugtype, bugtypeID)
	if bugtype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No bugtype found!"})
		return
	}
	var bug Bug
	db.Where("bugtypeID",bugtype.ID).First(&bug)
	if bug.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "bugtype还拥有bug，不能删除"})
		return
	}
	db.Delete(&bugtype)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageBugtype struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonBugtype
	}

type rpageBugtype struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonBugtype `json:"records"`
	}

func fetchPageBugtype(c *gin.Context){
	var totalcount int
	var pagecount int
	var bugtype pageBugtype
	var _jBugtypes []jsonBugtype
	Db := db
	itemBugtypes := make([]Bugtype,0)
	err := c.BindJSON(&bugtype)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case bugtype.Pageindex == 0:
		bugtype.Pageindex = 1
		fallthrough
	case bugtype.Pagesize == 0:
		bugtype.Pagesize = 20
	}
	if bugtype.Sortfield != "" {
		var build strings.Builder
		build.WriteString(bugtype.Sortfield)
		if bugtype.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemBugtypes).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bugtype 表数据查询错误"})
		return
	}

	if bugtype.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % bugtype.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / bugtype.Pagesize + yu
	}
	Db = Db.Limit(bugtype.Pagesize).Offset((bugtype.Pageindex - 1)* bugtype.Pagesize)
	if err := Db.Find(&itemBugtypes).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Bugtype 表数据查询错误"})
		return
	}

	for _, item := range itemBugtypes {
		_jBugtypes = append(_jBugtypes, jsonBugtype{
			ID : item.ID,
			Name : item.Name,
			Sn : item.Sn,
			About : item.About,
		})
	}
	_page := rpageBugtype{
		Pageindex : bugtype.Pageindex,
		Pagesize : bugtype.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jBugtypes,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

