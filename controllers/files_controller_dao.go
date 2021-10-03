package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
    "io/ioutil"
    "os"
)

var (
	FilesController = filesController{}
)

type filesController struct{}

func (controller filesController) Log(c *gin.Context) {

    value, ok := c.Request.URL.Query()["value"]
    if !ok || len(value[0]) < 1 {
        FilesController.BadRequestResponse(c)
        return
    }
    fileName := string(value[0])

    c.Header("Content-Type", "application/json")

    c.JSON(http.StatusOK, gin.H{
            "code" : http.StatusOK,
            "message": fileName,
    })
}

func (controller filesController) DownloadFile(c *gin.Context) {
	
    value, ok := c.Request.URL.Query()["value"]
    if !ok || len(value[0]) < 1 {
        FilesController.BadRequestResponse(c)
        return
    }
    fileName := string(value[0])

    filetype, ok := c.Request.URL.Query()["fileType"]
    if !ok || len(filetype[0]) < 1 {
        FilesController.BadRequestResponse(c)
        return
    }
    fileType := string(filetype[0])

    dir, err := os.Getwd()
    if err != nil {
        FilesController.NotFoundResponse(c)
        return
    }

    f := FileTypFeromString(fileType)
    filePath := dir + "/" + f.Path() + fileName + ".png"

    data, err:= ioutil.ReadFile(filePath)
    if err != nil {
        FilesController.NotFoundResponse(c)
        return
    }

    c.Writer.WriteHeader(http.StatusOK)
    c.Header("Content-Disposition", "attachment; filename=" + fileName + ".png")
    //c.Header("Content-Type", "application/octet-stream")
    c.Header("Content-Type", "image/jpeg")
    //c.Header("Content-Length", string(len(data)))
    c.Writer.Write(data)
}