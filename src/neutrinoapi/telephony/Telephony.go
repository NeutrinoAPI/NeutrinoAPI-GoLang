/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 01/07/2016
 */

package telephony


import "neutrinoapi/models"

/*
 * Interface for the TELEPHONY_IMPL
 */
type TELEPHONY interface {
    PhonePlayback (string, string) (*models.PhonePlaybackResponse, error)

    VerifySecurityCode (int) (*models.VerifySecurityCodeResponse, error)

    HLRLookup (string, *string) (*models.HLRLookupResponse, error)

    PhoneVerify (string, *int, *string, *string, *int, *int) (*models.PhoneVerifyResponse, error)

    SMSVerify (string, *int, *string, *string, *int) (*models.SMSVerifyResponse, error)
}

/*
 * Factory for the TELEPHONY interaface returning TELEPHONY_IMPL
 */
func NewTELEPHONY() TELEPHONY {
    return &TELEPHONY_IMPL{}
}