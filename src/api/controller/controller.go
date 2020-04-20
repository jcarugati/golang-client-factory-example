package controller

import "../services"

func ExamplController() {
	//This func would most likely receive a gin context when applied in a microservices infrastructure
	//orchestrates services regarding a certain functionality

	services.ExampleService()

	return

}
