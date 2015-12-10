/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 12/10/2015
 */

package ecommerce


import "neutrinoapi/models"

/*
 * Interface for the ECOMMERCE_IMPL
 */
type ECOMMERCE interface {
    CreateBINLookup (string, *string) (*models.BINLookupResponse, error)
}

/*
 * Factory for the ECOMMERCE interaface returning ECOMMERCE_IMPL
 */
func NewECOMMERCE() ECOMMERCE {
    return &ECOMMERCE_IMPL{}
}