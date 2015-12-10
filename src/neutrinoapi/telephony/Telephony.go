/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 12/10/2015
 */

package telephony


import "neutrinoapi/models"

/*
 * Interface for the TELEPHONY_IMPL
 */
type TELEPHONY interface {
    CreatePhonePlayback (string, string) (*models.PhonePlaybackResponse, error)

    CreateVerifySecurityCode (int) (*models.VerifySecurityCodeResponse, error)

    CreateHLRLookup (string, *string) (*models.HLRLookupResponse, error)

    CreatePhoneVerify (string, *int, *string, *string, *int, *int) (*models.PhoneVerifyResponse, error)

    CreateSMSVerify (string, *int, *string, *string, *int) (*models.SMSVerifyResponse, error)
}

/*
 * Factory for the TELEPHONY interaface returning TELEPHONY_IMPL
 */
func NewTELEPHONY() TELEPHONY {
    return &TELEPHONY_IMPL{}
}