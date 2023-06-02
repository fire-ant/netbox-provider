package ipamrole

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_ipam_role", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier

	})
}
