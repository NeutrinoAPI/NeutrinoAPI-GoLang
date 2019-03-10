/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package NeutrinoAPIClient

import(
	"neutrinoapi_lib/configuration_pkg"
	"neutrinoapi_lib/imaging_pkg"
	"neutrinoapi_lib/telephony_pkg"
	"neutrinoapi_lib/datatools_pkg"
	"neutrinoapi_lib/securityandnetworking_pkg"
	"neutrinoapi_lib/geolocation_pkg"
	"neutrinoapi_lib/ecommerce_pkg"
	"neutrinoapi_lib/www_pkg"
)


/*
 * Interface for the NEUTRINOAPI_IMPL
 */
type NEUTRINOAPI interface {
    Imaging()               imaging_pkg.IMAGING
    Telephony()             telephony_pkg.TELEPHONY
    DataTools()             datatools_pkg.DATATOOLS
    SecurityAndNetworking()       securityandnetworking_pkg.SECURITYANDNETWORKING
    Geolocation()           geolocation_pkg.GEOLOCATION
    ECommerce()             ecommerce_pkg.ECOMMERCE
    WWW()                   www_pkg.WWW
    Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the NEUTRINOAPI interface returning NEUTRINOAPI_IMPL
 */
func NewNEUTRINOAPI() NEUTRINOAPI {
    neutrinoAPIClient := new(NEUTRINOAPI_IMPL)
    neutrinoAPIClient.config = configuration_pkg.NewCONFIGURATION()

    return neutrinoAPIClient
}
