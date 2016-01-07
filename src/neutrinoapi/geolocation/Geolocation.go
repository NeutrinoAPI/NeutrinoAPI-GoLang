/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 01/07/2016
 */

package geolocation


import "neutrinoapi/models"

/*
 * Interface for the GEOLOCATION_IMPL
 */
type GEOLOCATION interface {
    IPInfo (string, *bool) (*models.IPInfoResponse, error)

    GeocodeAddress (string, *string, *string) (*models.GeocodeAddressResponse, error)

    GeocodeReverse (float32, float32, *string) (*models.GeocodeReverseResponse, error)
}

/*
 * Factory for the GEOLOCATION interaface returning GEOLOCATION_IMPL
 */
func NewGEOLOCATION() GEOLOCATION {
    return &GEOLOCATION_IMPL{}
}