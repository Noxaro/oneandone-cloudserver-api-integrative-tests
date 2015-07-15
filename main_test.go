/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"fmt"
	oaocs "github.com/Noxaro/oneandone-cloudserver-api"
	"github.com/docker/machine/log"
	"os"
)

var api *oaocs.API

func init() {

	token, err := getEnvironmentVar("token")
	if err != nil {
		log.Error("The 1&1 cloud server api endpoint must be set in the environment variable 'endpoint'")
		os.Exit(1)
	}

	apiEndpoint, err := getEnvironmentVar("endpoint")
	if err != nil {
		log.Error("The 1&1 cloud server api endpoint must be set in the environment variable 'endpoint'")
		os.Exit(1)
	}

	api = oaocs.New(token, apiEndpoint)
}

func getEnvironmentVar(name string) (string, error) {
	osVar := os.Getenv(name)
	if osVar == "" {
		return "", fmt.Errorf("The environment variable is not set")
	}
	os.Unsetenv(name)
	return osVar, nil
}

func GetAPI() *oaocs.API {
	return api
}
