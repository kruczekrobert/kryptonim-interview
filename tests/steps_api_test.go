package tests

import (
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"kryptonim-interview/app"
	"kryptonim-interview/app/errs"
	"kryptonim-interview/app/routers"
)

var apiRouters = map[string]func(container *app.Container) *gin.Engine{
	"Core": routers.SetupCore,
}

type requestData struct {
	Method  string
	Path    string
	Payload map[string]interface{}
}

type responseData struct {
	Response interface{} `json:"response"`
}

func (f *FeaturesContext) initRouter(name string) error {
	router, ok := apiRouters[name]
	if !ok {
		return fmt.Errorf("[14537537] router not exists, %s", name)
	}
	f.apiRouter = router(f.container)
	return nil
}

func (f *FeaturesContext) makeJsonRequest(data *godog.DocString) error {
	if f.apiRouter == nil {
		return fmt.Errorf("[14538538] empty api router")
	}

	var reqData *requestData

	if data != nil {
		err := json.Unmarshal([]byte(data.Content), &reqData)
		if err != nil {
			return fmt.Errorf("[147777] unmarshall error: %v", err)
		}
	}

	f.apiResponse = f.makeRequest(f.apiRouter, reqData.Method, reqData.Path, reqData.Payload)

	return nil
}

func (f *FeaturesContext) assertResponseStatus(status int) error {
	if f.apiResponse.Code != status {
		return fmt.Errorf("[1922432243] expected status: %d, actual status: %d", status, f.apiResponse.Code)
	}
	return nil
}

func (f *FeaturesContext) assertJsonResponse(status int, data *godog.DocString) error {
	err := f.assertResponseStatus(status)
	if err != nil {
		return errs.Wrap(err, "[19028028] assert response status")
	}

	var resData *responseData
	var jsonStr []byte

	if data != nil {
		err = json.Unmarshal([]byte(data.Content), &resData)
		if err != nil {
			return errs.Wrap(err, "[15225225] unmarshal data content")
		}

		jsonStr, err = json.Marshal(resData.Response)
		if err != nil {
			return errs.Wrap(err, "[15223223] marshal response")
		}
	}
	err = f.assertJsonEq(string(jsonStr), f.apiResponse.Body.String())
	if err != nil {
		return errs.Wrap(err, "[1525112511] assert json eq")
	}

	return nil
}

func (f *FeaturesContext) assertJsonEq(expected, actual string) error {
	if !assert.JSONEq(f.t, expected, actual) {
		return fmt.Errorf("[15521521] json eq")
	}
	return nil
}
