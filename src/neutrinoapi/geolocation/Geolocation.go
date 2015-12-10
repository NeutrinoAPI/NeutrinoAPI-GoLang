/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 12/10/2015
 */

package geolocation


import "neutrinoapi/models"

/*
 * Interface for the GEOLOCATION_IMPL
 */
type GEOLOCATION interface {
    CreateIPInfo (string, *bool) (*models.IPInfoResponse, error)

    CreateGeocodeAddress (string, *string, *string) (*models.GeocodeAddressResponse, error)

    CreateGeocodeReverse (float32, float32, *string) (*models.GeocodeReverseResponse, error)
}

/*
 * Factory for the GEOLOCATION interaface returning GEOLOCATION_IMPL
 */
func NewGEOLOCATION() GEOLOCATION {
    return &GEOLOCATION_IMPL{}
}