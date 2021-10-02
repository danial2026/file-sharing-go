package controllers

import (
    "fmt"
	"github.com/danial2026/file-sharing-go/domain"
	"github.com/danial2026/file-sharing-go/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
    "io/ioutil"
)

var (
	UsersController = usersController{}
)

type usersController struct{}

func (contoller usersController) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user id"})
		return
	}

	user, err := services.UsersService.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller usersController) Save(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid json body"})
		return
	}

	if err := services.UsersService.Save(&user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}


func (controller usersController) Log(c *gin.Context) {
	fmt.Println(c.Request.URL.Query())

    fileType, ok := c.Request.URL.Query()["fileType"]

    if !ok || len(fileType[0]) < 1 {
        return
    }
    res := []string{"foo", "bar"}
    c.JSON(200, res)

    fmt.Println(fileType[0])

    c.JSON(200, fileType[0])
}

func (controller usersController) DownloadFile(c *gin.Context) {

    data, err:= ioutil.ReadFile("cat.png")

    if err != nil {
        fmt.Println(err)
    }

    c.Writer.WriteHeader(http.StatusOK)
    //c.Header("Content-Disposition", "attachment; filename=a.tar")
    c.Header("Content-Type", "application/octet-stream")
    //c.Header("Content-Length", len(data))
    c.Writer.Write(data) //the memory take up 1.2~1.7G
}
