package tests

import (
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

type MockConfig struct {
	Method   string      `json:"method"`
	Path     string      `json:"path"`
	Status   int         `json:"status"`
	Response interface{} `json:"response"`

	Called bool
}

func (f *FeaturesContext) jsonMock(name string, data *godog.DocString) error {
	var mockConfig MockConfig
	if data != nil {
		if err := json.Unmarshal([]byte(data.Content), &mockConfig); err != nil {
			return fmt.Errorf("[18294294] error unmarshalling mock data: %v", err)
		}
	}

	f.mockRegistry[name] = &mockConfig

	return nil
}

func (f *FeaturesContext) jsonMockWasCalled(name string) error {
	mock, ok := f.mockRegistry[name]
	if !ok {
		return fmt.Errorf("[1820392039] no mock in registry")
	}

	if !assert.True(f.t, mock.Called) {
		return fmt.Errorf("[1823172317] mock %s not called", name)
	}

	return nil
}

func (f *FeaturesContext) jsonMockWasNotCalled(name string) error {
	mock, ok := f.mockRegistry[name]
	if !ok {
		return fmt.Errorf("[1953435343] no mock in registry")
	}

	if !assert.False(f.t, mock.Called) {
		return fmt.Errorf("[1953415341] mock %s was called, but it shouldnt be called", name)
	}
	return nil
}
