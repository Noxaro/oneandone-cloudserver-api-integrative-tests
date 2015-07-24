/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"github.com/docker/machine/log"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationServerAppliancesGet(t *testing.T) {
	sApps, err := GetAPI().GetServerAppliances()
	assert.Nil(t, err)

	for index, _ := range sApps {
		sApp, err := GetAPI().GetServerAppliance(sApps[index].Id)
		assert.Nil(t, err)
		assert.Equal(t, sApps[index].Id, sApp.Id)
		assert.Equal(t, sApps[index].Architecture, sApp.Architecture)
		assert.Equal(t, sApps[index].IsAutomaticInstall, sApp.IsAutomaticInstall)
		assert.Equal(t, sApps[index].Licenses, sApp.Licenses)
		assert.Equal(t, sApps[index].MinHddSize, sApp.MinHddSize)
		assert.Equal(t, sApps[index].Name, sApp.Name)
		assert.Equal(t, sApps[index].Os, sApp.Os)
		assert.Equal(t, sApps[index].OsFamily, sApp.OsFamily)
		assert.Equal(t, sApps[index].OsImageType, sApp.OsImageType)
		assert.Equal(t, sApps[index].OsVersion, sApp.OsVersion)

		time.Sleep(1 * time.Second)
	}
}

func TestIntegrationServerApplianceList(t *testing.T) {
	families, err := GetAPI().ServerApplianceListFamilies()
	assert.Nil(t, err)
	assert.True(t, len(families) >= 1)

	systems, err := GetAPI().ServerApplianceListOperationSystems(families[0])
	assert.Nil(t, err)
	assert.True(t, len(systems) >= 1)

	types, err := GetAPI().ServerApplianceListTypes(families[0], systems[0])
	assert.Nil(t, err)
	assert.True(t, len(types) >= 1)

	architectures, err := GetAPI().ServerApplianceListArchitectures(families[0], systems[0], types[0])
	assert.Nil(t, err)
	assert.True(t, len(architectures) >= 1)

	system, err := GetAPI().ServerApplianceFindNewest(families[0], systems[0], types[0], architectures[0], true)
	assert.Nil(t, err)
	log.Debug("_______________________________")
	log.Debug(system)
}
