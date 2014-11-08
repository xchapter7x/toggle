toggle
======


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
