/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package ecommerce_pkg

import "neutrinoapi_lib/configuration_pkg"
import "neutrinoapi_lib/models_pkg"

/*
 * Interface for the ECOMMERCE_IMPL
 */
type ECOMMERCE interface {
    BINLookup (string, *string) (*models_pkg.BINLookupResponse, error)
}

/*
 * Factory for the ECOMMERCE interaface returning ECOMMERCE_IMPL
 */
func NewECOMMERCE(config configuration_pkg.CONFIGURATION) *ECOMMERCE_IMPL {
    client := new(ECOMMERCE_IMPL)
    client.config = config
    return client
}