/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package telephony_pkg

import "neutrinoapi_lib/models_pkg"
import "neutrinoapi_lib/configuration_pkg"

/*
 * Interface for the TELEPHONY_IMPL
 */
type TELEPHONY interface {
    HLRLookup (string, *string) (*models_pkg.HLRLookupResponse, error)

    PhonePlayback (string, string) (*models_pkg.PhonePlaybackResponse, error)

    VerifySecurityCode (string) (*models_pkg.VerifySecurityCodeResponse, error)

    SMSVerify (string, *int64, *int64, *string, *string) (*models_pkg.SMSVerifyResponse, error)

    SMSMessage (string, string, *string) (*models_pkg.SMSMessageResponse, error)

    PhoneVerify (string, *int64, *int64, *int64, *string, *string) (*models_pkg.PhoneVerifyResponse, error)
}

/*
 * Factory for the TELEPHONY interaface returning TELEPHONY_IMPL
 */
func NewTELEPHONY(config configuration_pkg.CONFIGURATION) *TELEPHONY_IMPL {
    client := new(TELEPHONY_IMPL)
    client.config = config
    return client
}