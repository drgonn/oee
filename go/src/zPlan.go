package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
func fetchSinglePlan(c *gin.Context){
	var plan Plan
	planID := c.Param("id")
	db.First(&plan, planID)
	if plan.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No plan found!"})
		return
	}
	_plan :=  jsonPlan{
		ID : plan.ID,
		Name : plan.Name,
		Week : plan.Week,
		Current : plan.Current,
		Follow : plan.Follow,
		Time : plan.Time,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _plan})
}

func createPlan(c *gin.Context){
	var plan jsonPlan
	err := c.BindJSON(&plan)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case plan.Name == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"name 参数缺失"})
		return
	case plan.Week == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"week 参数缺失"})
		return
	case plan.Current == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"current 参数缺失"})
		return
	case plan.Follow == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"follow 参数缺失"})
		return
	case plan.ProjectID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"ProjectID 参数缺失"})
		return
	case plan.UserID == 0:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"UserID 参数缺失"})
		return
	}
	var queryPlan Plan
	db.Where("name = ?", plan.Name).First(&queryPlan)
	if queryPlan.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Plan name 已经存在，不允许重复"})
		return
	}
	var queryPlan Plan
	db.Where("week = ?", plan.Week).First(&queryPlan)
	if queryPlan.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Plan week 已经存在，不允许重复"})
		return
	}
	var queryPlan Plan
	db.Where("current = ?", plan.Current).First(&queryPlan)
	if queryPlan.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Plan current 已经存在，不允许重复"})
		return
	}
	var queryPlan Plan
	db.Where("follow = ?", plan.Follow).First(&queryPlan)
	if queryPlan.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Plan follow 已经存在，不允许重复"})
		return
	}
	var project Project
	projectID := plan.ProjectID
	db.First(&project, projectID)
	if project.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Project found!"})
		return
	}
	dbPlan := Plan{
		Name : plan.Name,
		Week : plan.Week,
		Current : plan.Current,
		Follow : plan.Follow,
		ProjectID : plan.ProjectID,
	}
	db.Save(&dbPlan)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbPlan.ID, 
	})
}

func updatePlan(c *gin.Context){
	var plan Plan
	planID := c.Param("id")
	db.First(&plan, planID)
	if plan.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No plan found!"})
		return
	}
	var jplan jsonPlan
	err := c.BindJSON(&jplan)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jplan.ProjectID != 0 {
		var project Project
		projectID := jplan.ProjectID
		db.First(&project, projectID)
		if project.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Project found!"})
			return
		}
	db.Model(&plan).Update("projectID",plan.ProjectID)
	}
	if jplan.Name != "" {
		var queryPlan Plan
		db.Where("Name = ?", jplan.Name).First(&queryPlan)
		if queryPlan.ID != 0 && queryPlan.Name != plan.Name {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Plan Name 已经存在，不允许重复"})
			return
		}
		db.Model(&plan).Update("Name",jplan.Name)
	}
	if jplan.Week != 0 {
		var queryPlan Plan
		db.Where("Week = ?", jplan.Week).First(&queryPlan)
		if queryPlan.ID != 0 && queryPlan.Week != plan.Week {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Plan Week 已经存在，不允许重复"})
			return
		}
		db.Model(&plan).Update("Week",jplan.Week)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":plan.ID, 
	})
}

func deletePlan(c *gin.Context){
	var plan Plan
	planID := c.Param("id")
	db.First(&plan, planID)
	if plan.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No plan found!"})
		return
	}
	db.Delete(&plan)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pagePlan struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonPlan
	}

type rpagePlan struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonPlan `json:"records"`
	}

func fetchPagePlan(c *gin.Context){
	var totalcount int
	var pagecount int
	var plan pagePlan
	var _jPlans []jsonPlan
	Db := db
	itemPlans := make([]Plan,0)
	err := c.BindJSON(&plan)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case plan.Pageindex == 0:
		plan.Pageindex = 1
		fallthrough
	case plan.Pagesize == 0:
		plan.Pagesize = 20
	}
	if plan.ProjectID != 0 { 
		var project Project
		projectID := plan.ProjectID
		db.First(&project, projectID)
		if project.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No Project found!"})
		return
		} else {
			Db = Db.Where("projectID = ?",plan.ProjectID)
		}
	}
	if plan.Sortfield != "" {
		var build strings.Builder
		build.WriteString(plan.Sortfield)
		if plan.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemPlans).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Plan 表数据查询错误"})
		return
	}

	if plan.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % plan.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / plan.Pagesize + yu
	}
	Db = Db.Limit(plan.Pagesize).Offset((plan.Pageindex - 1)* plan.Pagesize)
	if err := Db.Find(&itemPlans).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Plan 表数据查询错误"})
		return
	}

	for _, item := range itemPlans {
		_jPlans = append(_jPlans, jsonPlan{
			ID : item.ID,
			Name : item.Name,
			Week : item.Week,
			Current : item.Current,
			Follow : item.Follow,
			Time : item.Time,
			ProjectID : item.ProjectID,
			UserID : item.UserID,
		})
	}
	_page := rpagePlan{
		Pageindex : plan.Pageindex,
		Pagesize : plan.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jPlans,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

