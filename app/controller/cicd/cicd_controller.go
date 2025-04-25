package cicd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Riyoukou/odyssey/app/model"
	"github.com/Riyoukou/odyssey/app/repository"
	"github.com/Riyoukou/odyssey/pkg/response"
	"github.com/gin-gonic/gin"
)

func HandleCICDFetchRepo(c *gin.Context) {
	var (
		err    error
		result interface{}
	)
	switch c.Param("type") {
	case "cluster":
		result, err = repository.FetchClusters()
	case "project":
		result, err = repository.FetchProjects()
	case "service":
		result, err = repository.FetchServicesByProjectAndEnv(c.Query("project"), c.Query("env"))
	case "code_library":
		result, err = repository.FetchCodeLibraries()
	case "code_source":
		result, err = repository.FetchCodeSources()
	}
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}
	response.Success(c, result, fmt.Sprintf("%s fetched successfully", c.Param("type")))
}

func HandleCICDGetRepo(c *gin.Context) {
	var (
		err    error
		result interface{}
	)
	switch c.Param("type") {
	case "cluster":
		result, err = repository.GetClusterByName(c.Query("name"))
	case "project":
		result, err = repository.GetProjectByName(c.Query("name"))
	case "service":
		result, err = repository.GetServiceByNameAndProjectByEnv(c.Query("name"), c.Query("project"), c.Query("env"))
	case "code_library":
		result, err = repository.GetCodeLibraryByNameAndProject(c.Query("name"), c.Query("project"))
	case "code_source":
		result, err = repository.GetCodeSourceByName(c.Query("name"))
	}
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}
	response.Success(c, result, fmt.Sprintf("%s fetched successfully", c.Param("type")))
}

func HandleCICDCreateRepo(c *gin.Context) {
	var err error
	switch c.Param("type") {
	case "cluster":
		var req model.ClusterTable
		if err = c.ShouldBind(&req); err != nil {
			break
		}
		err = repository.CreateCluster(req)
	case "project":
		var req model.ProjectTable
		if err = c.ShouldBind(&req); err != nil {
			break
		}
		err = repository.CreateProject(req)
	case "service":
		var req model.ServiceTable
		if err = c.ShouldBind(&req); err != nil {
			break
		}
		err = repository.CreateService(req)
	case "code_library":
		var req model.CodeLibraryTable
		if err = c.ShouldBind(&req); err != nil {
			break
		}
		err = repository.CreateCodeLibrary(req)
	case "code_source":
		var req model.CodeSourceTable
		if err = c.ShouldBind(&req); err != nil {
			break
		}
		err = repository.CreateCodeSource(req)
	}
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}
	response.Success(c, nil, fmt.Sprintf("%s created successfully", c.Param("type")))
}

func HandleCICDUpdateRepo(c *gin.Context) {
	var err error
	switch c.Param("type") {
	case "cluster":
		var req model.ClusterTable
		if err = c.ShouldBind(&req); err != nil {
			break
		}
		err = repository.UpdateCluster(req)
	case "project":
		var req model.ProjectTable
		if err = c.ShouldBind(&req); err != nil {
			break
		}
		err = repository.UpdateProject(req)
	case "service":
		var req model.ServiceTable
		if err = c.ShouldBind(&req); err != nil {
			break
		}
		err = repository.UpdateServiceByNameAndProjectByEnv(req)
	case "code_library":
		var req model.CodeLibraryTable
		if err = c.ShouldBind(&req); err != nil {
			break
		}
		err = repository.UpdateCodeLibraryByNameAndProject(req)
	case "code_source":
		var req model.CodeSourceTable
		if err = c.ShouldBind(&req); err != nil {
			break
		}
		err = repository.UpdateCodeSourceByName(req)
	}
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}
	response.Success(c, nil, fmt.Sprintf("%s updated successfully", c.Param("type")))
}

func HandleCICDDeleteRepo(c *gin.Context) {
	var (
		err   error
		intID int64
	)
	intID, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}
	switch c.Param("type") {
	case "cluster":
		err = repository.DeleteCluster(intID)
	case "project":
		err = repository.DeleteProject(intID)
	case "service":
		err = repository.DeleteService(intID)
	case "code_library":
		err = repository.DeleteCodeLibrary(intID)
	}
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}
	response.Success(c, nil, fmt.Sprintf("%s deleted successfully", c.Param("type")))
}
