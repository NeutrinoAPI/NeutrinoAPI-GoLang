/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package imaging_pkg

import "neutrinoapi_lib/configuration_pkg"

/*
 * Interface for the IMAGING_IMPL
 */
type IMAGING interface {
    ImageWatermark (string, string, *int64, *string, *string, *int64, *int64) ([]byte, error)

    QRCode (string, *int64, *int64, *string, *string) ([]byte, error)

    ImageResize (string, int64, int64, *string) ([]byte, error)

    HTML5Render (string, *string, *string, *string, *int64, *int64, *int64, *int64, *int64, *bool, *float64, *bool, *bool, *bool, *bool, *string, *int64, *int64, *int64, *string, *string, *string, *int64, *string, *int64, *bool, *string, *string, *string, *int64, *string, *int64, *bool, *int64, *int64) ([]byte, error)
}

/*
 * Factory for the IMAGING interaface returning IMAGING_IMPL
 */
func NewIMAGING(config configuration_pkg.CONFIGURATION) *IMAGING_IMPL {
    client := new(IMAGING_IMPL)
    client.config = config
    return client
}