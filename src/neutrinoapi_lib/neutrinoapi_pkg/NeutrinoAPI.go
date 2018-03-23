/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io )
 */

package NeutrinoAPIClient

import(
	"neutrinoapi_lib/configuration_pkg"
	"neutrinoapi_lib/imaging_pkg"
	"neutrinoapi_lib/telephony_pkg"
	"neutrinoapi_lib/data tools_pkg"
	"neutrinoapi_lib/security and networking_pkg"
	"neutrinoapi_lib/geolocation_pkg"
	"neutrinoapi_lib/e-commerce_pkg"
)


/*
 * Interface for the NEUTRINOAPI_IMPL
 */
type NEUTRINOAPI interface {
     Imaging()               imaging_pkg.IMAGING
     Telephony()             telephony_pkg.TELEPHONY
     DataTools()            data tools_pkg.DATA TOOLS
     SecurityAndNetworking()       security and networking_pkg.SECURITY AND NETWORKING
     Geolocation()           geolocation_pkg.GEOLOCATION
     ECommerce()            e-commerce_pkg.E-COMMERCE
     Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the NEUTRINOAPI interaface returning NEUTRINOAPI_IMPL
 */
func NewNEUTRINOAPI() NEUTRINOAPI {
    neutrinoAPIClient := new(NEUTRINOAPI_IMPL)
    neutrinoAPIClient.config = configuration_pkg.NewCONFIGURATION()
    return neutrinoAPIClient
}
