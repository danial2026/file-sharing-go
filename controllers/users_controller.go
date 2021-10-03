package controllers

import (
	"github.com/danial2026/file-sharing-go/domain"
	"github.com/danial2026/file-sharing-go/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
    "io/ioutil"
    "time"
    "os"
)

type FileType int

const (
	PROFILE_PIC FileType = iota
    POST_PIC
    Unknown
)

func FileTypFeromString(str string) FileType {
    switch str {
    case "PROFILE_PIC":
        return PROFILE_PIC
    case "POST_PIC":
        return POST_PIC
    }
    return Unknown
}

func (f FileType) Path() string {
	switch f {
    case PROFILE_PIC:
		return "downloadFile/users/"
    case POST_PIC:
		return "downloadFile/posts/"
    }
    return "Unknown"
}

var (
	UsersController = usersController{}
)

type usersController struct{}

type response struct {
    Timestamp string `json:"timestamp"`
    Status   int      `json:"status"`
    Message string `json:"message"`
    Path string `json:"path"`
}

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

    value, ok := c.Request.URL.Query()["value"]

    if !ok || len(value[0]) < 1 {
        UsersController.BadRequestResponse(c)
        return
    }

    fileName := string(value[0])
    //stringRes := response{
    //    Status: 200,
    //    Message: []string{fileType[0]}}

    //res,_ := json.Marshal(stringRes)

    /*
    jData, err := json.Marshal(fileType[0])
    if err != nil {
        // handle error
    }
    */

    c.Header("Content-Type", "application/json")


    c.JSON(http.StatusOK, gin.H{
            "code" : http.StatusOK,
            "message": fileName,
    })

    //c.JSON(200, fileType[0])

    //c.JSON(200, res)
    //fmt.Println(fileType[0])

    //c.JSON(200, fileType[0])
}

func (controller usersController) DownloadFile(c *gin.Context) {

    value, ok := c.Request.URL.Query()["value"]
    if !ok || len(value[0]) < 1 {
        UsersController.BadRequestResponse(c)
        return
    }
    fileName := string(value[0])

    filetype, ok := c.Request.URL.Query()["fileType"]
    if !ok || len(filetype[0]) < 1 {
        UsersController.BadRequestResponse(c)
        return
    }
    fileType := string(filetype[0])

    dir, err := os.Getwd()
    if err != nil {
        UsersController.NotFoundResponse(c)
        return
    }

    f := FileTypFeromString(fileType)
    filePath := dir + "/" + f.Path() + fileName + ".png"

    c.Header("filePath", filePath)

    data, err:= ioutil.ReadFile(filePath)
    if err != nil {
        UsersController.NotFoundResponse(c)
        return
    }

    c.Writer.WriteHeader(http.StatusOK)
    c.Header("Content-Disposition", "attachment; filename=" + fileName + ".png")
    //c.Header("Content-Type", "application/octet-stream")
    c.Header("Content-Type", "image/jpeg")
    //c.Header("Content-Length", string(len(data)))
    c.Writer.Write(data)
}

func (controller usersController) BadRequestResponse(c *gin.Context) {

    UsersController.JsonResponse(c, http.StatusBadRequest, "Bad Request")
}

func (controller usersController) NotFoundResponse(c *gin.Context) {

    UsersController.JsonResponse(c, http.StatusNotFound, "Not Found")
}

func (controller usersController) JsonResponse(c *gin.Context, http_status int, messege string) {

    c.Header("Content-Type", "application/json")

     var response_json response

    response_json.Status=http_status
    response_json.Message=messege

    time_stamp , _ := time.Now().UTC().MarshalText()
    response_json.Timestamp=string(time_stamp)

    response_json.Path=c.Request.URL.Path
    c.JSON(response_json.Status, response_json)
}
