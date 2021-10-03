package controllers

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
		"code":    http.StatusOK,
		"message": fileName,
	})
}

func (controller filesController) StoreUserPicBase64(c *gin.Context) {

	userIdZuul, ok := c.Request.URL.Query()["userIdZuul"]
	if !ok || len(userIdZuul[0]) < 1 {
		FilesController.BadRequestResponse(c)
		return
	}
	userId := string(userIdZuul[0])

	fileBase64, ok := c.Request.URL.Query()["fileBase64"]
	if !ok || len(fileBase64[0]) < 1 {
		FilesController.BadRequestResponse(c)
		return
	}
	file, err := base64.StdEncoding.DecodeString(fileBase64[0])
	if err != nil {
		FilesController.BadRequestResponse(c)
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		FilesController.NotFoundResponse(c)
		return
	}

	f := FileTypFeromString("PROFILE_PIC")

	filename := userId
	out, err := os.Create(dir + "/" + f.Path() + filename)
	if err != nil {
		FilesController.BadRequestResponse(c)
		return
	}

	defer out.Close()

	if _, err := out.Write(file); err != nil {
		FilesController.NotFoundResponse(c)
		return
	}
	if err := out.Sync(); err != nil {
		FilesController.NotFoundResponse(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"filepath": dir + "/" + f.Path() + filename})
}

func (controller filesController) StoreUserPic(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		//   c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		FilesController.NotFoundResponse(c)
		return
	}

	f := FileTypFeromString("PROFILE_PIC")

	filename := header.Filename
	out, err := os.Create(dir + "/" + f.Path() + filename)
	if err != nil {
		FilesController.BadRequestResponse(c)
		return
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		FilesController.BadRequestResponse(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"filepath": dir + "/" + f.Path() + filename})
}

func (controller filesController) DownloadFileHandler(c *gin.Context) {

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

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		FilesController.NotFoundResponse(c)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename="+fileName+".png")
	//c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Type", "image/jpeg")
	//c.Header("Content-Length", string(len(data)))
	c.Writer.Write(data)
}
