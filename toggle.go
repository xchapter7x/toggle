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

func SetFeatureStatus(featureName string, featureStatus int) (err error) {
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
	status   int
	filter   func(...interface{}) bool
	settings map[string]interface{}
}

const (
	FEATURE_ON = iota
	FEATURE_OFF
	FEATURE_FILTER
)

var featureList map[string]*feature
var namespace string

func Init(ns string) {
	featureList = make(map[string]*feature)
	namespace = ns
}

func ShowFeatures() map[string]*feature {
	return featureList
}

func RegisterFeature(featureSignature string) {

	if _, exists := featureList[featureSignature]; !exists {
		featureList[featureSignature] = &feature{
			name:   featureSignature,
			status: FEATURE_OFF,
		}
	}
}
