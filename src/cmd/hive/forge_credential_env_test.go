package main

import (
	"slices"
	"testing"

	cruxconfig "github.com/hivecommons/hive/crux/config"
	"github.com/hivecommons/hive/pkg/config"
)

func TestSyncForgeCredentialEnvRegistersConfiguredForgejoName(t *testing.T) {
	cfg := &config.Config{}
	cfg.Forgejo.TokenEnv = "CUSTOM_FORGEJO_PAT"
	syncForgeCredentialEnv(cfg)
	if got := cruxconfig.ExtraCredentialEnvVars(); !slices.Contains(got, "CUSTOM_FORGEJO_PAT") {
		t.Fatalf("registry = %v, want it to include CUSTOM_FORGEJO_PAT", got)
	}

	cfg.Forgejo.TokenEnv = ""
	syncForgeCredentialEnv(cfg)
	got := cruxconfig.ExtraCredentialEnvVars()
	if !slices.Contains(got, "CUSTOM_FORGEJO_PAT") {
		t.Fatalf("registry = %v, want the previous custom name kept after reload", got)
	}
	if !slices.Contains(got, config.DefaultForgejoTokenEnv) {
		t.Fatalf("registry = %v, want it to include the default %s", got, config.DefaultForgejoTokenEnv)
	}
}
