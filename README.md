toggle
======

[![wercker status](https://app.wercker.com/status/9c11e691895a9782a234fcc9bb313819/m "wercker status")](https://app.wercker.com/project/bykey/9c11e691895a9782a234fcc9bb313819)


func **Flip**(flg string, defaultFeature, newFeature interface{}, iargs ...interface{}) (responseInterfaceArray []interface{})

func **SetFeatureStatus**(featureName, featureStatus string) (err error)

func **IsActive**(flg string) (active bool)

func **Init**(ns string, engine storageEngine)

func **ShowFeatures**() map[string]*feature

func **RegisterFeature**(featureSignature string) (err error)

func **RegisterFeatureWithStatus**(featureSignature, statusValue string) (err error)

func **NewDefaultEngine**() (engine storageEngine)



type **DefaultEngine** struct

func (s \*DefaultEngine) **GetFeatureStatusValue**(featureSignature string) (status string, err error)


Sample usage:
(./sample/main.go)
```
package main

import (
	"fmt"

	"github.com/xchapter7x/goutil/unpack"
	"github.com/xchapter7x/toggle"
)

func TestA(s string) (r string) {
	r = fmt.Sprintln("testa", s)
	fmt.Println(r)
	return
}

func TestB(s string) (r string) {
	r = fmt.Sprintln("testb", s)
	fmt.Println(r)
	return
}

func main() {
	toggle.Init("MAINTEST", nil)
	toggle.RegisterFeature("test")
	f := toggle.Flip("test", TestA, TestB, "argstring")
	var output string
	unpack.Unpack(f, &output)
	fmt.Println(output)

}
```


```
$ test=true go run sample/main.go
testb argstring

testb argstring



$ go run sample/main.go
testa argstring

testa argstring
```
