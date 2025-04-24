package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"kryptonim-interview/app/store"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
)

func (f *FeaturesContext) makeRequest(r *gin.Engine, method string, path string, payload map[string]interface{}) *httptest.ResponseRecorder {
	buf := bytes.NewBuffer(nil)

	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			panic(fmt.Sprintf("[14623623] marshal error: %v", err))
		}
		buf.Write(b)
	}

	req, _ := http.NewRequest(method, path, buf)

	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	return response
}

func (f *FeaturesContext) clearDatabase() {
	for _, v := range store.AppModels {
		err := f.container.DB().Migrator().DropTable(&v)
		if err != nil {
			panic(fmt.Sprintf("[14641641] cannot drop table: %v", err))
		}
	}
	err := store.Migrate(f.container.DB())
	if err != nil {
		panic(fmt.Sprintf("[14642642] cannot migrate db: %v", err))
	}
}

func (f *FeaturesContext) stringAsKind(value string, k reflect.Kind, t reflect.Type) interface{} {
	switch k {
	case reflect.String:
		return reflect.ValueOf(value).Convert(t).Interface()
	case reflect.Int:
		val, err := strconv.Atoi(value)
		require.NoErrorf(f.t, err, "[14157157] `%s` is not an integer: %s", value, err)
		return reflect.ValueOf(val).Convert(t).Interface()
	case reflect.Uint:
		val, err := strconv.ParseUint(value, 10, 64)
		require.NoErrorf(f.t, err, "[14155155] `%s` is not an uint: %s", value, err)
		return reflect.ValueOf(val).Convert(t).Interface()
	case reflect.Float64:
		val, err := strconv.ParseFloat(value, 64)
		require.NoErrorf(f.t, err, "[14158158] '%s' is not a float64: %s", value, err)
		return reflect.ValueOf(val).Convert(t).Interface()
	default:
		panic(fmt.Sprintf("[14159159] no such kind: `%s`", k))
	}
}

func (f *FeaturesContext) tableToMaps(table *godog.Table) []map[string]string {
	var results []map[string]string

	keys := table.Rows[0].Cells

	for _, values := range table.Rows[1:] {
		result := map[string]string{}
		for idx, value := range values.Cells {
			result[keys[idx].Value] = value.Value
		}
		results = append(results, result)
	}
	return results
}

func (f *FeaturesContext) mapToStruct(object interface{}, values map[string]string) {
	reflectedValue := reflect.ValueOf(object).Elem()
	for name, value := range values {
		field := reflectedValue.FieldByName(name)
		require.Truef(f.t, field.IsValid(), "[1415161516] no field `%s` in model `%s`", name, reflectedValue.Type().Name())
		field.Set(reflect.ValueOf(f.stringAsKind(value, field.Kind(), field.Type())))
	}
}
