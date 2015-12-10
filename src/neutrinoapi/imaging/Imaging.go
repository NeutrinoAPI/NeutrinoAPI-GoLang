/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 12/10/2015
 */

package imaging


/*
 * Interface for the IMAGING_IMPL
 */
type IMAGING interface {
    CreateQRCode (string, *string, *string, *int, *int) ([]byte, error)

    CreateHTMLToPDF (string, *int, *int, *string) ([]byte, error)

    CreateImageResize (int, string, int, *string) ([]byte, error)

    CreateImageWatermark (string, string, *string, *int, *int, *string, *int) ([]byte, error)
}

/*
 * Factory for the IMAGING interaface returning IMAGING_IMPL
 */
func NewIMAGING() IMAGING {
    return &IMAGING_IMPL{}
}