package services

import(
	"../clients/factory"
	"../constants"
)
func ExampleService() {

	factory.ClientFactory(constants.EXAMPLECLIENT)
	//Posibly some business logic goes here

	return
}
