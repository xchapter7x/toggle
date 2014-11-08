package toggle

import (
	"fmt"
	"reflect"
)

func createReflectValueArgsArray(iargs []interface{}) (args []reflect.Value) {
	for _, arg := range iargs {
		args = append(args, reflect.ValueOf(arg))
	}
	return
}

func createInterfaceArrayFromValuesArray(responseValuesArray []reflect.Value) (responseInterfaceArray []interface{}) {
	for _, ri := range responseValuesArray {
		responseInterfaceArray = append(responseInterfaceArray, ri.Interface())
	}
	return
}

func getActivePointer(flg string, defaultFeature, newFeature interface{}) (activePointer interface{}) {
	if IsActive(flg) {
		activePointer = newFeature

	} else {
		activePointer = defaultFeature
	}
	return
}

func Flip(flg string, defaultFeature, newFeature interface{}, iargs ...interface{}) (responseInterfaceArray []interface{}) {
	args := createReflectValueArgsArray(iargs)
	ptr := getActivePointer(flg, defaultFeature, newFeature)
	responseValuesArray := reflect.ValueOf(ptr).Call(args)
	responseInterfaceArray = createInterfaceArrayFromValuesArray(responseValuesArray)
	return
}

func SetFeatureStatus(featureName, featureStatus string) (err error) {
	if _, exists := featureList[featureName]; exists {
		featureList[featureName].status = featureStatus

	} else {
		err = fmt.Errorf("Feature toggle doesnt exist")
	}
	return
}

func IsActive(flg string) (active bool) {
	if feature, exists := featureList[flg]; !exists || feature.status == FEATURE_OFF {
		active = false

	} else {
		active = true
	}
	return
}

type feature struct {
	name     string
	status   string
	filter   func(...interface{}) bool
	settings map[string]interface{}
}

const (
	FEATURE_ON     = "true"
	FEATURE_OFF    = "false"
	FEATURE_FILTER = "filter:x:x"
)

var featureList map[string]*feature
var namespace string
var toggleEngine storageEngine

func Init(ns string, engine storageEngine) {
	featureList = make(map[string]*feature)
	namespace = ns

	if engine != nil {
		toggleEngine = engine

	} else {
		toggleEngine = NewDefaultEngine()
	}
}

func ShowFeatures() map[string]*feature {
	return featureList
}

func getFeatureStatusValue(featureSignature, statusValue string) (status string) {
	var err error

	if status, err = toggleEngine.GetFeatureStatusValue(featureSignature); err != nil {
		status = statusValue
	}
	return
}

func RegisterFeature(featureSignature string) (err error) {
	err = RegisterFeatureWithStatus(featureSignature, FEATURE_OFF)
	return
}

func RegisterFeatureWithStatus(featureSignature, statusValue string) (err error) {
	if _, exists := featureList[featureSignature]; !exists {
		featureList[featureSignature] = &feature{
			name:   featureSignature,
			status: getFeatureStatusValue(featureSignature, statusValue),
		}

	} else {
		err = fmt.Errorf("feature already registered")
	}
	return
}
