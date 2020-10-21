package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
// 토큰안에 이름도 넣는다.

func Dashboard(c *gin.Context) {

	// email, name 을 출력할 수 있도록 만든다.
	c.JSON(http.StatusOK, gin.H{"email": c.Keys["email"], "name": c.Keys["name"]})

	return
}