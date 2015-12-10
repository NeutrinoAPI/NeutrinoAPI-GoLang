/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 12/10/2015
 */

package securityandnetworking


import "neutrinoapi/models"

/*
 * Interface for the SECURITYANDNETWORKING_IMPL
 */
type SECURITYANDNETWORKING interface {
    CreateURLInfo (bool, string) (*models.URLInfoResponse, error)

    CreateHostReputation (string) (*models.HostReputationResponse, error)

    CreateIPBlocklist (string) (*models.IPBlocklistResponse, error)
}

/*
 * Factory for the SECURITYANDNETWORKING interaface returning SECURITYANDNETWORKING_IMPL
 */
func NewSECURITYANDNETWORKING() SECURITYANDNETWORKING {
    return &SECURITYANDNETWORKING_IMPL{}
}