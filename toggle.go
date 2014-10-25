package toggle

import "reflect"

func Flip(flg string, dflt interface{}, newFtr interface{}, iargs ...interface{}) (r []interface{}) {
	var cr []reflect.Value
	var args []reflect.Value

	for _, arg := range iargs {
		args = append(args, reflect.ValueOf(arg))
	}
	var ptr interface{}

	if IsActive(flg) {
		ptr = newFtr

	} else {
		ptr = dflt
	}
	cr = reflect.ValueOf(ptr).Call(args)

	for _, ri := range cr {
		r = append(r, ri.Interface())
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

var featureList map[string]feature
var namespace string

func Init(ns string) {
	featureList = make(map[string]feature)
	namespace = ns
}

func ShowFeatures() map[string]feature {
	return featureList
}

func RegisterFeature(featureSignature string) {

	if _, exists := featureList[featureSignature]; !exists {
		featureList[featureSignature] = feature{
			name:   featureSignature,
			status: FEATURE_OFF,
		}
	}
}
