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

	storageName := "ITTestStorage" + GetTimestamp()
	sharedStorageSettings := oaocs.SharedStorageSettings{
		Name:        storageName,
		Description: "Test",
		Size:        100,
	}

	sharedStorage, err := GetAPI().CreateSharedStorage(sharedStorageSettings)
	assert.Nil(t, err)

	assert.Equal(t, storageName, sharedStorage.Name)
	assert.Equal(t, "Test", sharedStorage.Description)
	assert.Equal(t, 100, sharedStorage.Size)

	sharedStorageId = sharedStorage.Id
}

func TestIntegrationSharedStorageUpdate(t *testing.T) {
	sharedStorage, err := GetAPI().GetSharedStorage(sharedStorageId)
	sharedStorage.WaitForState("ACTIVE")

	storageName := "ITTestStorageRename" + GetTimestamp()
	config := oaocs.SharedStorageSettings{
		Name:        storageName,
		Description: "Test2",
		Size:        200,
	}

	assert.Nil(t, err)

	sharedStorage, err = sharedStorage.UpdateConfig(config)
	assert.Nil(t, err)

	sharedStorage.WaitForState("ACTIVE")
	sharedStorage, err = GetAPI().GetSharedStorage(sharedStorageId)
	assert.Nil(t, err)

	assert.Equal(t, storageName, sharedStorage.Name)
	assert.Equal(t, "Test2", sharedStorage.Description)
	assert.Equal(t, 200, sharedStorage.Size)
}

func TestIntegrationSharedStoragesServer(t *testing.T) {
	serverName := "IT Test Server" + GetTimestamp()
	latestAppliance, err := GetAPI().ServerApplianceFindNewest("Linux", "Ubuntu", "Minimal", 64, true)
	assert.Nil(t, err)
	config := oaocs.ServerCreateData{
		Name:        serverName,
		ApplianceId: latestAppliance.Id,
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
}

func TestIntegrationSharedStorageDelete(t *testing.T) {
	sharedStorage, err := GetAPI().GetSharedStorage(sharedStorageId)
	assert.Nil(t, err)

	_, err = sharedStorage.Delete()
	assert.Nil(t, err)
}
