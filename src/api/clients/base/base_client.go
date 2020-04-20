package base

import (
	"encoding/json"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

type BaseClient struct {
	Domain string
}

func (baseClient *BaseClient) Get(path string, query map[string]string, returnObject interface{}) error {
	if &path == nil {
		path = ""
	}
	client := &http.Client{}
	req, requestErr := http.NewRequest("GET", baseClient.Domain+path+"?"+ path, nil)
	if requestErr != nil {
		logger.Errorf("error when creating request: %v", requestErr,"[method:get]",
			"[struct:base_client]","[event:exception]",fmt.Sprintf("request:%v",req.URL))
		return requestErr
	}

	response, responseErr := client.Do(req)

	if responseErr != nil {
		logger.Errorf("error when sending GET request: %v", requestErr,"[method:get]","[struct:base_client]","[event:exception]",
			fmt.Sprintf("request:%s",req.URL))
		return responseErr
	}
	defer response.Body.Close()

	decodeErr := json.NewDecoder(response.Body).Decode(returnObject)

	if decodeErr != nil {
		logger.Errorf("[method:get][struct:base_client][event:exception] error when decoding request: %v", decodeErr)
		return decodeErr
	}

	logger.Infof("Success Get request [method:get][struct:base_client][event:get_request][request:%s]", req.URL)
	return nil
}
