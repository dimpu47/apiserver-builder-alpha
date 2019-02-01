/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"log"

	examplecontroller "github.com/kubernetes-incubator/apiserver-builder-alpha/example/pkg/controller"
	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/controller"
)

var kubeconfig = flag.String("kubeconfig", "", "path to kubeconfig")

func main() {
	flag.Parse()
	config, err := controller.GetConfig(*kubeconfig)
	if err != nil {
		log.Fatalf("Could not create Config for talking to the apiserver: %v\n", err)
	}

	controllers, _ := examplecontroller.GetAllControllers(config)
	controller.StartControllerManager(controllers...)

	// Blockforever
	select {}
}
