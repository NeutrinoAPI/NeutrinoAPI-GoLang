/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 01/07/2016
 */

package ecommerce


import "neutrinoapi/models"

/*
 * Interface for the ECOMMERCE_IMPL
 */
type ECOMMERCE interface {
    BINLookup (string, *string) (*models.BINLookupResponse, error)
}

/*
 * Factory for the ECOMMERCE interaface returning ECOMMERCE_IMPL
 */
func NewECOMMERCE() ECOMMERCE {
    return &ECOMMERCE_IMPL{}
}