/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package www_pkg

import "neutrinoapi_lib/models_pkg"
import "neutrinoapi_lib/configuration_pkg"

/*
 * Interface for the WWW_IMPL
 */
type WWW interface {
    URLInfo (string, *bool) (*models_pkg.URLInfoResponse, error)

    BrowserBot (string, *int64, *int64, *string, []string, *string, *bool) (*models_pkg.BrowserBotResponse, error)

    HTMLClean (string, string) ([]byte, error)
}

/*
 * Factory for the WWW interaface returning WWW_IMPL
 */
func NewWWW(config configuration_pkg.CONFIGURATION) *WWW_IMPL {
    client := new(WWW_IMPL)
    client.config = config
    return client
}