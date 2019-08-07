/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package geolocation_pkg

import "neutrinoapi_lib/configuration_pkg"
import "neutrinoapi_lib/models_pkg"

/*
 * Interface for the GEOLOCATION_IMPL
 */
type GEOLOCATION interface {
    GeocodeReverse (string, string, *string, *string) (*models_pkg.GeocodeReverseResponse, error)

    IPInfo (string, *bool) (*models_pkg.IPInfoResponse, error)

    GeocodeAddress (string, *string, *string, *bool) (*models_pkg.GeocodeAddressResponse, error)
}

/*
 * Factory for the GEOLOCATION interaface returning GEOLOCATION_IMPL
 */
func NewGEOLOCATION(config configuration_pkg.CONFIGURATION) *GEOLOCATION_IMPL {
    client := new(GEOLOCATION_IMPL)
    client.config = config
    return client
}