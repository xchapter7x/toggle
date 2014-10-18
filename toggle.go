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
		ptr = dflt

	} else {
		ptr = newFtr
	}
	cr = reflect.ValueOf(ptr).Call(args)

	for _, ri := range cr {
		r = append(r, ri.Interface())
	}
	return
}

func IsActive(flg) (active bool) {

	if _, exists := featureList[flg]; !exists {
		active = false

	} else {

	}
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

var featureList []feature

func RegisterFeature(featureSignature string) {

	if _, exists := featureList[flg]; !exists {
		featureList = append(featureList, feature{
			name:   featureSignature,
			status: FEATURE_OFF,
		})
	}
}
