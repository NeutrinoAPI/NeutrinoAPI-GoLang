/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package www_pkg

import "neutrinoapi_lib/configuration_pkg"
import "neutrinoapi_lib/models_pkg"

/*
 * Interface for the WWW_IMPL
 */
type WWW interface {
    URLInfo (string, *bool, *bool, *int64) (*models_pkg.URLInfoResponse, error)

    HTMLClean (string, string) ([]byte, error)

    BrowserBot (string, *int64, *int64, *string, []string, *string, *bool) (*models_pkg.BrowserBotResponse, error)
}

/*
 * Factory for the WWW interaface returning WWW_IMPL
 */
func NewWWW(config configuration_pkg.CONFIGURATION) *WWW_IMPL {
    client := new(WWW_IMPL)
    client.config = config
    return client
}