package app

import (
	"./../controllers"
)

func mapUrls() {
    prefix:=""/api/v1/files""

	router.GET(prefix+"/users/:id", controllers.UsersController.Get)
	router.POST(prefix+"/users", controllers.UsersController.Save)
}
