/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestIntegrationGetDvdIsos(t *testing.T) {
	dvdIsos, err := GetAPI().GetDvdIsos()
	assert.Nil(t, err)

	if len(dvdIsos) == 0 {
		t.Error("There are no dvd isos in the result from the api")
	}

	for index, _ := range dvdIsos {
		time.Sleep(1 * time.Second)
		dvdIso, err := GetAPI().GetDvdIso(dvdIsos[index].Id)
		assert.Nil(t, err)
		assert.Equal(t, dvdIsos[index].Id, dvdIso.Id)
		assert.Equal(t, dvdIsos[index].Name, dvdIso.Name)
		assert.Equal(t, dvdIsos[index].Os, dvdIso.Os)
		assert.Equal(t, dvdIsos[index].OsVersion, dvdIso.OsVersion)
		assert.Equal(t, dvdIsos[index].OsFamily, dvdIso.OsFamily)
		assert.Equal(t, dvdIsos[index].Architecture, dvdIso.Architecture)
		assert.Equal(t, dvdIsos[index].Type, dvdIso.Type)
	}

}