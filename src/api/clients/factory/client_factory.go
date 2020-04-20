package factory

import (
	"../../constants"
	"../rest"
	"../base"
)

func ClientFactory(client string) base.BaseClientContract {
	//ClientFactory using a string in constants. More Clients can be added if needed as a switch case.
	switch client {
	case constants.EXAMPLECLIENT:
		return rest.NewExampleClient()
	default:
		return nil
	}
}
