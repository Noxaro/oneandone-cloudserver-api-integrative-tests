/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"testing"
	oaocs "github.com/Noxaro/oneandone-cloudserver-api"
	"github.com/stretchr/testify/assert"
)

var fwpId = ""

func TestIntegrationCreateFirewallPolicy(t *testing.T) {
	firewall, err := GetAPI().CreateFirewallPolicy(oaocs.FirewallPolicyCreateData{
		Name:        "[Docker Machine] ",
		Description: "Firewall policy for docker machine",
		Rules: []oaocs.FirewallPolicyRulesCreateData{
			oaocs.FirewallPolicyRulesCreateData{
				Protocol: "TCP",
				PortFrom: oaocs.Int2Pointer(1),
				PortTo:   oaocs.Int2Pointer(65535),
				SourceIp: "0.0.0.0",
			},
			oaocs.FirewallPolicyRulesCreateData{
				Protocol: "UDP",
				PortFrom: oaocs.Int2Pointer(1),
				PortTo: oaocs.Int2Pointer(65535),
				SourceIp: "0.0.0.0",
			},
			oaocs.FirewallPolicyRulesCreateData{
				Protocol: "ICMP",
				PortFrom: nil,
				PortTo: nil,
				SourceIp: "0.0.0.0",
			},
		},
	})
	assert.Nil(t, err)
	firewall.WaitForState("ACTIVE")
	fwpId = firewall.Id
}

func TestIntegrationGetFirewallPolicy(t *testing.T) {
	firewallPolicies, err := GetAPI().GetFirewallPolicies()
	assert.Nil(t, err)
	assert.True(t, len(firewallPolicies) >= 1)

	for index, _ := range firewallPolicies {
		fwp, err := GetAPI().GetFirewallPolicy(firewallPolicies[index].Id)
		assert.Nil(t, err)
		assert.Equal(t, firewallPolicies[index].Id, fwp.Id)
		assert.Equal(t, firewallPolicies[index].Name, fwp.Name)
		assert.Equal(t, firewallPolicies[index].Description, fwp.Description)
		assert.Equal(t, firewallPolicies[index].Status, fwp.Status)

		for rIndex, _ := range firewallPolicies[index].Rules {
			assert.Equal(t, firewallPolicies[index].Rules[rIndex].Id, fwp.Rules[rIndex].Id)
			assert.Equal(t, firewallPolicies[index].Rules[rIndex].Protocol, fwp.Rules[rIndex].Protocol)
			assert.Equal(t, firewallPolicies[index].Rules[rIndex].PortFrom, fwp.Rules[rIndex].PortFrom)
			assert.Equal(t, firewallPolicies[index].Rules[rIndex].PortTo, fwp.Rules[rIndex].PortTo)
			assert.Equal(t, firewallPolicies[index].Rules[rIndex].SourceIp, fwp.Rules[rIndex].SourceIp)
		}

	}
}

func TestIntegrationDeleteFirewallPolicy(t *testing.T) {
	firewall, err := GetAPI().GetFirewallPolicy(fwpId)
	assert.Nil(t, err)
	firewall.Delete()
	err = firewall.WaitUntilDeleted()
	assert.Nil(t, err)
}