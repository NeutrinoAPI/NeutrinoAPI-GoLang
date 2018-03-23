/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io )
 */

package securityandnetworking_pkg

import "neutrinoapi_lib/models_pkg"
import "neutrinoapi_lib/configuration_pkg"

/*
 * Interface for the SECURITYANDNETWORKING_IMPL
 */
type SECURITYANDNETWORKING interface {
    HostReputation (string) (*models_pkg.HostReputationResponse, error)

    URLInfo (string, bool) (*models_pkg.URLInfoResponse, error)

    IPProbe (string) (*models_pkg.IPProbeResponse, error)

    IPBlocklist (string) (*models_pkg.IPBlocklistResponse, error)

    EmailVerify (string, *bool) (*models_pkg.EmailVerifyResponse, error)
}

/*
 * Factory for the SECURITYANDNETWORKING interaface returning SECURITYANDNETWORKING_IMPL
 */
func NewSECURITYANDNETWORKING(config configuration_pkg.CONFIGURATION) *SECURITYANDNETWORKING_IMPL {
    client := new(SECURITYANDNETWORKING_IMPL)
    client.config = config
    return client
}