package app

import (
	"github.com/danial2026/file-sharing-go/controllers"
)

func mapUrls() {
    prefix:="/api/v1/files"

    router.GET(prefix+"/download", controllers.FilesController.Log)
    router.GET(prefix+"/downloadFile", controllers.FilesController.DownloadFile)
}
