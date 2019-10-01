package main

import (
	"sync"
	"fmt"
)

func main() {
	testConfigMap := configMap{}

	for i:=0; i < 100; i++ {
		go func() {
			testConfigMap.Lock()
			defer testConfigMap.Unlock()

			for j:= 0 ; j < 100; j++ {
				key := fmt.Sprintf("%s-%s", i, j)
				testConfigMap.Config[key] = "TEST"
			}

		}()
	}

	testConfigMap2 := configMap{}

	go func() {
		testConfigMap.Config = testConfigMap2.Config
	}()
}

type configMap struct {
	Config  map[string]string
	*sync.Mutex
}

