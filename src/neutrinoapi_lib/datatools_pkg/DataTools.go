/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package datatools_pkg

import "neutrinoapi_lib/configuration_pkg"
import "neutrinoapi_lib/models_pkg"

/*
 * Interface for the DATATOOLS_IMPL
 */
type DATATOOLS interface {
    EmailValidate (string, *bool) (*models_pkg.EmailValidateResponse, error)

    UserAgentInfo (string) (*models_pkg.UserAgentInfoResponse, error)

    BadWordFilter (string, *string) (*models_pkg.BadWordFilterResponse, error)

    Convert (string, string, string) (*models_pkg.ConvertResponse, error)

    PhoneValidate (string, *string, *string) (*models_pkg.PhoneValidateResponse, error)
}

/*
 * Factory for the DATATOOLS interaface returning DATATOOLS_IMPL
 */
func NewDATATOOLS(config configuration_pkg.CONFIGURATION) *DATATOOLS_IMPL {
    client := new(DATATOOLS_IMPL)
    client.config = config
    return client
}