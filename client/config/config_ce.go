// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !ent

package config

import (
	"github.com/hashicorp/go-hclog"
	structsc "github.com/hashicorp/nomad/nomad/structs/config"
)

// GetVaultConfigs returns the set of Vault configurations available for this
// client. In Nomad CE we only use the default Vault.
func (c *Config) GetVaultConfigs(logger hclog.Logger) map[string]*structsc.VaultConfig {
	if c.VaultConfigs["default"] == nil || !c.VaultConfigs["default"].IsEnabled() {
		return nil
	}

	if len(c.VaultConfigs) > 1 {
		logger.Warn("multiple Vault configurations are only supported in Nomad Enterprise")
	}
	return c.VaultConfigs
}

// GetConsulConfigs returns the set of Consul configurations the fingerprint needs
// to check. In Nomad CE we only check the default Consul.
func (c *Config) GetConsulConfigs(logger hclog.Logger) map[string]*structsc.ConsulConfig {
	if c.ConsulConfigs["default"] == nil {
		return nil
	}

	if len(c.ConsulConfigs) > 1 {
		logger.Warn("multiple Consul configurations are only supported in Nomad Enterprise")
	}

	return c.ConsulConfigs
}
