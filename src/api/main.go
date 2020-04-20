package api

import (
	"../api/configuration"
	"../api/controller"
)

func main() {
	//you would probaably have a dockerfile that would fire the main func o a bash script
	//also if you were to use gingonic as a service engine you'd need a router file that would configure the paths
	configuration.Init()	//fires the configuration logic
	controller.ExamplController()

}
