package tests

import (
	"fmt"
	"github.com/cucumber/godog"
	"kryptonim-interview/app/errs"
	"kryptonim-interview/app/exchanges"
	"reflect"
)

var modelTypes = modelTypesRegistry{
	"ExchangeRates": exchanges.ExchangeRates{},
}

type modelTypesRegistry map[string]interface{}

func (r modelTypesRegistry) NewModel(modelType string) interface{} {
	mType, ok := r[modelType]
	if !ok {
		panic(fmt.Sprintf("no such model `%s`", modelType))
	}
	return reflect.New(reflect.TypeOf(mType)).Interface()
}

func (f *FeaturesContext) saveModels(models ...interface{}) {
	for _, model := range models {
		f.saveModel(model)
	}
}

func (f *FeaturesContext) saveModel(model interface{}) {
	errs.FatalOnError(f.container.DB().Model(model).Save(model).Error, "[1410591059] save model")
}

func (f *FeaturesContext) models(model string, data *godog.Table) error {

	objectsFields := f.tableToMaps(data)

	var objects []interface{}

	for _, values := range objectsFields {
		object := modelTypes.NewModel(model)
		f.mapToStruct(object, values)
		objects = append(objects, object)
	}

	f.saveModels(objects...)

	return nil
}
