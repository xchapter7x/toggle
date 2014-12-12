package goutil

import "reflect"

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

func findErrorValue(responseInterfaceArray []interface{}) (err error) {
	for _, res := range responseInterfaceArray {
		if e, ok := res.(error); ok {
			err = e
		}
	}
	return
}

func CallChain(preverr error, functor interface{}, iargs ...interface{}) (responseInterfaceArray []interface{}, err error) {
	if err = preverr; err == nil {
		args := createReflectValueArgsArray(iargs)
		responseValuesArray := reflect.ValueOf(functor).Call(args)
		responseInterfaceArray = createInterfaceArrayFromValuesArray(responseValuesArray)
		err = findErrorValue(responseInterfaceArray)
	}
	return
}

func CallChainP(preverr error, responseInterfaceArray []interface{}, functor interface{}, iargs ...interface{}) (err error) {
	res, err := CallChain(preverr, functor, iargs...)
	_, err = CallChain(err, UnpackArray, res, responseInterfaceArray)
	_, err = CallChain(err, findErrorValue, responseInterfaceArray)
	return
}
