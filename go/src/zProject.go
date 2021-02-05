package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func fetchSingleProject(c *gin.Context){
	var project Project
	projectID := c.Param("id")
	db.First(&project, projectID)
	if project.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No project found!"})
		return
	}
	_project :=  jsonProject{
		ID : project.ID,
		Name : project.Name,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error_code":0, "data": _project})
}

func createProject(c *gin.Context){
	var project jsonProject
	err := c.BindJSON(&project)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case project.Name == "":
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"name 参数缺失"})
		return
	}
	var queryProject Project
	db.Where("name = ?", project.Name).First(&queryProject)
	if queryProject.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Project name 已经存在，不允许重复"})
		return
	}
	dbProject := Project{
		Name : project.Name,
	}
	db.Save(&dbProject)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":dbProject.ID, 
	})
}

func updateProject(c *gin.Context){
	var project Project
	projectID := c.Param("id")
	db.First(&project, projectID)
	if project.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No project found!"})
		return
	}
	var jproject jsonProject
	err := c.BindJSON(&jproject)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	}
	if jproject.Name != "" {
		var queryProject Project
		db.Where("Name = ?", jproject.Name).First(&queryProject)
		if queryProject.ID != 0 && queryProject.Name != project.Name {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Project Name 已经存在，不允许重复"})
			return
		}
		db.Model(&project).Update("Name",jproject.Name)
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"id":project.ID, 
	})
}

func deleteProject(c *gin.Context){
	var project Project
	projectID := c.Param("id")
	db.First(&project, projectID)
	if project.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "No project found!"})
		return
	}
	var plan Plan
	db.Where("projectID",project.ID).First(&plan)
	if plan.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "project还拥有plan，不能删除"})
		return
	}
	db.Delete(&project)
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
	})
}

type pageProject struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Order int `json:"order"`
	Sortfield string `json:"sortfield"`
	jsonProject
	}

type rpageProject struct {
	Pageindex int `json:"current"`
	Pagesize int `json:"pageSize"`
	Pagecount int `json:"pagecount"`
	Totalcount int `json:"totalcount"`
	Records []jsonProject `json:"records"`
	}

func fetchPageProject(c *gin.Context){
	var totalcount int
	var pagecount int
	var project pageProject
	var _jProjects []jsonProject
	Db := db
	itemProjects := make([]Project,0)
	err := c.BindJSON(&project)
	switch {
	case err != nil:
		c.JSON(200,gin.H{"success":false,"error_code": -1,"errmsg":"Post data err"})
		return
	case project.Pageindex == 0:
		project.Pageindex = 1
		fallthrough
	case project.Pagesize == 0:
		project.Pagesize = 20
	}
	if project.Sortfield != "" {
		var build strings.Builder
		build.WriteString(project.Sortfield)
		if project.Order == 0 {
			build.WriteString(" desc")
		} else {
			build.WriteString(" asc")
		}
	orderstr := build.String()
	Db = Db.Order(orderstr)
	}
	if err := Db.Find(&itemProjects).Count(&totalcount).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Project 表数据查询错误"})
		return
	}

	if project.Pagesize > totalcount { 
		pagecount = 1 
	} else {
		var yu int
		if  totalcount  % project.Pagesize == 0{ 
			yu = 0
		} else{ 
			yu = 1 
		} 
		pagecount = totalcount / project.Pagesize + yu
	}
	Db = Db.Limit(project.Pagesize).Offset((project.Pageindex - 1)* project.Pagesize)
	if err := Db.Find(&itemProjects).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error_code":-4, "errmsg": "Project 表数据查询错误"})
		return
	}

	for _, item := range itemProjects {
		_jProjects = append(_jProjects, jsonProject{
			ID : item.ID,
			Name : item.Name,
		})
	}
	_page := rpageProject{
		Pageindex : project.Pageindex,
		Pagesize : project.Pagesize,
		Pagecount : pagecount,
		Totalcount : totalcount,
		Records : _jProjects,
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"error_code": 0,
		"data": _page,
	})
}

