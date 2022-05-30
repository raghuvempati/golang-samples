package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
	apiv1 "k8s.io/api/core/v1"
)

func main() {
	file, err := os.Open("configmap.yaml")
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	configMap := &apiv1.ConfigMap{}

	err = yaml.Unmarshal(bytes, configMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(configMap)
}
