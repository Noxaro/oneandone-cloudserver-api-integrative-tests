/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	oaocs "github.com/Noxaro/oneandone-cloudserver-api"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sharedStorageId = ""

func TestIntegrationSharedStorageCreate(t *testing.T) {

	sharedStorageSettings := oaocs.SharedStorageSettings{
		Name:        "ITTestStorage",
		Description: "Test",
		Size:        100,
	}

	sharedStorage, err := GetAPI().CreateSharedStorage(sharedStorageSettings)
	assert.Nil(t, err)

	assert.Equal(t, "ITTestStorage", sharedStorage.Name)
	assert.Equal(t, "Test", sharedStorage.Description)
	assert.Equal(t, 100, sharedStorage.Size)

	sharedStorageId = sharedStorage.Id
}

func TestIntegrationSharedStorageUpdate(t *testing.T) {
	sharedStorage, err := GetAPI().GetSharedStorage(sharedStorageId)
	sharedStorage.WaitForState("ACTIVE")

	config := oaocs.SharedStorageSettings{
		Name:        "ITTestStorage2",
		Description: "Test2",
		Size:        200,
	}

	assert.Nil(t, err)

	sharedStorage, err = sharedStorage.UpdateConfig(config)
	assert.Nil(t, err)

	sharedStorage.WaitForState("ACTIVE")
	sharedStorage, err = GetAPI().GetSharedStorage(sharedStorageId)
	assert.Nil(t, err)

	assert.Equal(t, "ITTestStorage2", sharedStorage.Name)
	assert.Equal(t, "Test2", sharedStorage.Description)
	assert.Equal(t, 200, sharedStorage.Size)
}

func TestIntegrationSharedStoragesServer(t *testing.T) {
	config := oaocs.ServerCreateData{
		Name:        "IT Test Server",
		ApplianceId: "C14988A9ABC34EA64CD5AAC0D33ABCAF",
		Hardware: oaocs.Hardware{
			Vcores:            1,
			CoresPerProcessor: 1,
			Ram:               1,
			Hdds: []oaocs.Hdd{
				oaocs.Hdd{
					Size:   40,
					IsMain: true,
				},
			},
		},
		PowerOn: true,
	}

	server, err := GetAPI().CreateServer(config)
	assert.Nil(t, err)

	server.WaitForState("POWERED_ON")

	sharedStorage, err := GetAPI().GetSharedStorage(sharedStorageId)
	assert.Nil(t, err)

	serverStoragePermissions := oaocs.SharedStorageServerPermissions{
		[]oaocs.SharedStorageServer{
			oaocs.SharedStorageServer{
				Id:     server.Id,
				Rights: "RW",
			},
		},
	}

	sharedStorage.UpdateServerPermissions(serverStoragePermissions)
	sharedStorage.WaitForState("ACTIVE")

	accessPermission, err := sharedStorage.GetServerPermission(server.Id)
	assert.Nil(t, err)
	assert.Equal(t, server.Id, accessPermission.Id)
	assert.Equal(t, "RW", accessPermission.Rights)

	accessPermission.DeleteServerPermission()

	server.Delete()
	sharedStorage.Delete()
}
