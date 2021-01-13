package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"redteam/model"
)

func ProjectCreate(c *gin.Context) {
	// 계정번호
	num := c.Keys["number"].(int)

	// DB
	db, _ := c.Get("db")
	conn := db.(sql.DB)

	// 프로젝트 생성
	p := model.Project{}
	err := c.ShouldBindJSON(&p)
	err = p.ProjectCreate(&conn, num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOk": 0,
			"error": err,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"isOk": 1,
		})
		return
	}
}

func GetProject(c *gin.Context) {
	// 계정번호
	num := c.Keys["number"].(int)

	// DB
	db, _ := c.Get("db")
	conn := db.(sql.DB)

	// 프로젝트 조회
	projects, err := model.ReadProject(&conn, num)
	if err != nil {
		log.Println("GetProject error occured, account :", c.Keys["email"])
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"isOk": 0,
			"error": err,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"project_list" : projects,
		})
	}
}

func EndProjectList(c *gin.Context) {
	// 사용자 계정정보
	num := c.Keys["number"].(int)

	db, _ := c.Get("db") // httpheader.go 의 DBMiddleware 에 셋팅되어있음.
	conn := db.(sql.DB)

	p := model.Project{}
	c.ShouldBindJSON(&p)

	err := p.EndProject(&conn, num)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : http.StatusBadRequest,
			"isOk": 0,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOk": 1,
	})
}

func StartProjectList(c *gin.Context) {
	// 사용자 계정정보
	num := c.Keys["number"].(int)

	db, _ := c.Get("db") // httpheader.go 의 DBMiddleware 에 셋팅되어있음.
	conn := db.(sql.DB)

	p := model.ProjectStart{}
	c.ShouldBindJSON(&p)

	// 프로젝트 상태변경
	err := p.StartProject(&conn, num)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : http.StatusBadRequest,
			"isOk": 0,
		})
		return
	}

	//kafka producer & consumer
	err = p.Kafka(&conn, num)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : http.StatusBadRequest,
			"isOk": 0,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"isOk": 1,
		})
	}
}

func GetTag(c *gin.Context) {
	num := c.Keys["number"].(int)

	c.JSON(http.StatusOK, gin.H{
		"isOk":   1,
		"status": http.StatusOK,
		"tags":   model.GetTag(num), // 태그들
	})
}