package rest

import (
	"../base"
	"../../configuration"
)

type applicationClient struct {
	base.BaseClient

}

func NewExampleClient() base.BaseClientContract {
	var applicationClient applicationClient

	applicationClient.Domain = configuration.GetConfigurationInstance().ConfigurationInstance.UserClient
	return &applicationClient
}
