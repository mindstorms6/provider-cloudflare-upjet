/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"fmt"
	"strings"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/mindstorms6/provider-cloudflare-upjet/config/branch"
	"github.com/mindstorms6/provider-cloudflare-upjet/config/repository"
)

const (
	resourcePrefix = "cloudflare-upjet"
	modulePath     = "github.com/mindstorms6/provider-cloudflare-upjet"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cloudflare.crossplane.io"),
		// ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),

		// have to use a custom branch currently
		ujconfig.WithSkipList([]string{"map"}),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
	)

	println(fmt.Sprintf("RG: %s", pc.RootGroup))
	for name, resource := range pc.Resources {
		group := pc.RootGroup
		if resource.ShortGroup != "" {
			group = strings.ToLower(resource.ShortGroup) + "." + pc.RootGroup
		}
		println(fmt.Sprintf("%s:%s", group, name))
	}

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		repository.Configure,
		branch.Configure,
	} {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc
}
