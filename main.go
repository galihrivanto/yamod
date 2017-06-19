package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func modify(key, value string, kv map[interface{}]interface{}) {
	if strings.Contains(key, ".") {
		rootKey := key[:strings.Index(key, ".")]
		subV, ok := kv[rootKey]
		if !ok {
			return
		}

		subKv, ok := subV.(map[interface{}]interface{})
		if !ok {
			return
		}

		modify(strings.TrimPrefix(key, rootKey+"."), value, subKv)
	} else {
		kv[key] = value
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("incomplete parameter")
		os.Exit(-1)
	}

	target := os.Args[1]
	key := os.Args[2]
	value := os.Args[3]

	bi, err := ioutil.ReadFile(target)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	kv := make(map[interface{}]interface{})
	if err := yaml.Unmarshal(bi, &kv); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	modify(key, value, kv)

	bo, err := yaml.Marshal(kv)
	ioutil.WriteFile(target, bo, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
