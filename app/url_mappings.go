package app

import (
	"github.com/danial2026/file-sharing-go/controllers"
)

func mapUrls() {
    prefix:="/api/v1/files"

	router.GET(prefix+"/users/:id", controllers.UsersController.Get)
	router.POST(prefix+"/users", controllers.UsersController.Save)

    router.GET(prefix+"/download", controllers.UsersController.Log)
    router.GET(prefix+"/downloadFile", controllers.UsersController.DownloadFile)
}
