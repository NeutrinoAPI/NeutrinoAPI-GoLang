/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 12/10/2015
 */

package datatools


import "neutrinoapi/models"

/*
 * Interface for the DATATOOLS_IMPL
 */
type DATATOOLS interface {
    CreatePhoneValidate (string, *string) (*models.PhoneValidateResponse, error)

    CreateUserAgentInfo (string) (*models.UserAgentInfoResponse, error)

    CreateCodeHighlight (string, string, *bool) ([]byte, error)

    CreateBadWordFilter (string, *string) (*models.BadWordFilterResponse, error)

    CreateConvert (string, string, string) (*models.ConvertResponse, error)

    CreateEmailValidate (string, *bool) (*models.EmailValidateResponse, error)

    CreateHTMLClean (string, string) ([]byte, error)

    CreateHTMLExtract (string, string, *string, *string) (*models.HTMLExtractResponse, error)
}

/*
 * Factory for the DATATOOLS interaface returning DATATOOLS_IMPL
 */
func NewDATATOOLS() DATATOOLS {
    return &DATATOOLS_IMPL{}
}