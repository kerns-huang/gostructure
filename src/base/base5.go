package main

import (
	"io/ioutil"
	"os"
)
func main() {
	f, err := os.Open("bash1.go")
	defer f.Close()
	if err != nil {
		return
	}
	b,err := ioutil.ReadAll(f)
	println(string(b))
}
