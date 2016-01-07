/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 01/07/2016
 */

package securityandnetworking


import "neutrinoapi/models"

/*
 * Interface for the SECURITYANDNETWORKING_IMPL
 */
type SECURITYANDNETWORKING interface {
    URLInfo (bool, string) (*models.URLInfoResponse, error)

    HostReputation (string) (*models.HostReputationResponse, error)

    IPBlocklist (string) (*models.IPBlocklistResponse, error)
}

/*
 * Factory for the SECURITYANDNETWORKING interaface returning SECURITYANDNETWORKING_IMPL
 */
func NewSECURITYANDNETWORKING() SECURITYANDNETWORKING {
    return &SECURITYANDNETWORKING_IMPL{}
}