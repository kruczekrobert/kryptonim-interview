package tests

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"kryptonim-interview/app"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sort"
	"strings"
	"testing"
	"time"
)

func NewFeaturesContext(t *testing.T) FeaturesContext {
	return FeaturesContext{
		t:            t,
		container:    app.NewContainer(),
		mockRegistry: make(map[string]*MockConfig),
	}
}

type FeaturesContext struct {
	t              *testing.T
	apiRouter      *gin.Engine
	mockServer     *gin.Engine
	mockHTTPServer *http.Server
	mockStopCh     chan struct{}
	apiResponse    *httptest.ResponseRecorder
	container      *app.Container
	mockRegistry   map[string]*MockConfig
	err            string
}

func TestFeature(t *testing.T) {
	f := NewFeaturesContext(t)
	f.mockServerStart()
	defer f.mockServerStop()

	//FEATURE: no db logs and no logs in tests
	//f.container.DB().Config.Logger = logger.Default.LogMode(logger.Silent)
	//log.SetOutput(ioutil.Discard)

	gin.SetMode(gin.TestMode)

	godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: f.InitializeTestSuite,
		ScenarioInitializer:  f.InitializeScenario,
		Options: &godog.Options{
			Randomize:     1,
			StopOnFailure: true,
			Format:        "progress",
			Strict:        true,
			Paths:         []string{"features"},
			Tags:          os.Getenv("GODOG_TAGS"),
			TestingT:      t,
		},
	}.Run()

}

func (f *FeaturesContext) clear(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	f.apiRouter = nil
	f.apiResponse = nil
	f.err = ""
	f.mockRegistry = make(map[string]*MockConfig)
	f.clearDatabase()

	return ctx, nil
}

func (f *FeaturesContext) setup() {
	f.clearDatabase()
}

func (f *FeaturesContext) InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		f.setup()
	})
}

func (f *FeaturesContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(f.clear)

	// api steps
	ctx.Step(`^router "([^"]*)"$`, f.initRouter)
	ctx.Step(`^have request$`, f.makeJsonRequest)
	ctx.Step(`^have json response with status "([^"]*)"$`, f.assertJsonResponse)

	// gorm steps
	ctx.Step(`^models "([^"]*)"$`, f.models)

	// mock steps
	ctx.Step(`^json mock "([^"]*)"$`, f.jsonMock)
	ctx.Step(`^json mock "([^"]*)" was called`, f.jsonMockWasCalled)
	ctx.Step(`^json mock "([^"]*)" was not called`, f.jsonMockWasNotCalled)

}

func (f *FeaturesContext) mockServerStart() {
	f.mockServer = gin.Default()
	f.mockStopCh = make(chan struct{})

	f.mockServer.Any("/mocks/*path", func(c *gin.Context) {
		method := c.Request.Method
		rawPath := strings.TrimPrefix(c.Param("path"), "/")
		fullURL := f.buildNormalizedURL(rawPath, c.Request.URL.Query())

		for _, mock := range f.mockRegistry {
			if method == mock.Method && fullURL == mock.Path {
				c.JSON(mock.Status, mock.Response)
				mock.Called = true
				return
			}
		}
		c.JSON(404, gin.H{"error": "not found"})
	})

	f.mockHTTPServer = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: f.mockServer,
	}

	go func() {
		if err := f.mockHTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[204848] mock server error: %v", err)
		}
	}()

	go func() {
		<-f.mockStopCh
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = f.mockHTTPServer.Shutdown(ctx)
	}()
}

func (f *FeaturesContext) buildNormalizedURL(path string, query url.Values) string {
	queryParts := make([]string, 0, len(query))
	for key, values := range query {
		for _, value := range values {
			queryParts = append(queryParts, fmt.Sprintf("%s=%s", key, value))
		}
	}
	sort.Strings(queryParts)
	if len(queryParts) == 0 {
		return path
	}
	return fmt.Sprintf("%s?%s", path, strings.Join(queryParts, "&"))
}

func (f *FeaturesContext) mockServerStop() {
	if f.mockStopCh != nil {
		close(f.mockStopCh)
	}
}
