// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package hcp_packer_image

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion   *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Bucket              *string           `mapstructure:"bucket_name" required:"true" cty:"bucket_name" hcl:"bucket_name"`
	IterationID         *string           `mapstructure:"iteration_id" required:"true" cty:"iteration_id" hcl:"iteration_id"`
	CloudProvider       *string           `mapstructure:"cloud_provider" required:"true" cty:"cloud_provider" hcl:"cloud_provider"`
	Region              *string           `mapstructure:"region" required:"true" cty:"region" hcl:"region"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"bucket_name":                &hcldec.AttrSpec{Name: "bucket_name", Type: cty.String, Required: false},
		"iteration_id":               &hcldec.AttrSpec{Name: "iteration_id", Type: cty.String, Required: false},
		"cloud_provider":             &hcldec.AttrSpec{Name: "cloud_provider", Type: cty.String, Required: false},
		"region":                     &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
	}
	return s
}

// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDatasourceOutput struct {
	CloudProvider *string           `mapstructure:"cloud_provider" cty:"cloud_provider" hcl:"cloud_provider"`
	ComponentType *string           `mapstructure:"component_type" cty:"component_type" hcl:"component_type"`
	CreatedAt     *string           `mapstructure:"created_at" cty:"created_at" hcl:"created_at"`
	BuildID       *string           `mapstructure:"build_id" cty:"build_id" hcl:"build_id"`
	IterationID   *string           `mapstructure:"iteration_id" cty:"iteration_id" hcl:"iteration_id"`
	PackerRunUUID *string           `mapstructure:"packer_run_uuid" cty:"packer_run_uuid" hcl:"packer_run_uuid"`
	ID            *string           `mapstructure:"id" cty:"id" hcl:"id"`
	Region        *string           `mapstructure:"region" cty:"region" hcl:"region"`
	Labels        map[string]string `mapstructure:"labels" cty:"labels" hcl:"labels"`
}

// FlatMapstructure returns a new FlatDatasourceOutput.
// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DatasourceOutput) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDatasourceOutput)
}

// HCL2Spec returns the hcl spec of a DatasourceOutput.
// This spec is used by HCL to read the fields of DatasourceOutput.
// The decoded values from this spec will then be applied to a FlatDatasourceOutput.
func (*FlatDatasourceOutput) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"cloud_provider":  &hcldec.AttrSpec{Name: "cloud_provider", Type: cty.String, Required: false},
		"component_type":  &hcldec.AttrSpec{Name: "component_type", Type: cty.String, Required: false},
		"created_at":      &hcldec.AttrSpec{Name: "created_at", Type: cty.String, Required: false},
		"build_id":        &hcldec.AttrSpec{Name: "build_id", Type: cty.String, Required: false},
		"iteration_id":    &hcldec.AttrSpec{Name: "iteration_id", Type: cty.String, Required: false},
		"packer_run_uuid": &hcldec.AttrSpec{Name: "packer_run_uuid", Type: cty.String, Required: false},
		"id":              &hcldec.AttrSpec{Name: "id", Type: cty.String, Required: false},
		"region":          &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"labels":          &hcldec.AttrSpec{Name: "labels", Type: cty.Map(cty.String), Required: false},
	}
	return s
}