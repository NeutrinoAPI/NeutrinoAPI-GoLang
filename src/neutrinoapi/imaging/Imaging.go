/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 01/07/2016
 */

package imaging


/*
 * Interface for the IMAGING_IMPL
 */
type IMAGING interface {
    QRCode (string, *string, *string, *int, *int) ([]byte, error)

    HTMLToPDF (string, *int, *int, *string) ([]byte, error)

    ImageResize (int, string, int, *string) ([]byte, error)

    ImageWatermark (string, string, *string, *int, *int, *string, *int) ([]byte, error)
}

/*
 * Factory for the IMAGING interaface returning IMAGING_IMPL
 */
func NewIMAGING() IMAGING {
    return &IMAGING_IMPL{}
}