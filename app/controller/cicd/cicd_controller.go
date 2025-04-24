package cicd

import (
	"net/http"

	"github.com/Riyoukou/odyssey/app/model"
	"github.com/Riyoukou/odyssey/app/repository"
	"github.com/Riyoukou/odyssey/pkg/response"
	"github.com/gin-gonic/gin"
)

func HandleCICDFetchRepo(c *gin.Context) {
	switch c.Param("type") {
	}
}

func HandleCICDGetRepo(c *gin.Context) {
	switch c.Param("type") {
	}
}

func HandleCICDCreateRepo(c *gin.Context) {
	var err error
	switch c.Param("type") {
	case "cluster":
		var req model.ClusterTable
		if err := c.ShouldBind(&req); err != nil {
			break
		}
		err = repository.CreateCluster(req)
	}
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
	}
}

func HandleCICDDeleteRepo(c *gin.Context) {
	switch c.Param("type") {
	}
}
