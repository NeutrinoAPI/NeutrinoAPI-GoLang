/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 01/07/2016
 */

package datatools


import "neutrinoapi/models"

/*
 * Interface for the DATATOOLS_IMPL
 */
type DATATOOLS interface {
    PhoneValidate (string, *string) (*models.PhoneValidateResponse, error)

    UserAgentInfo (string) (*models.UserAgentInfoResponse, error)

    CodeHighlight (string, string, *bool) ([]byte, error)

    BadWordFilter (string, *string) (*models.BadWordFilterResponse, error)

    Convert (string, string, string) (*models.ConvertResponse, error)

    EmailValidate (string, *bool) (*models.EmailValidateResponse, error)

    HTMLClean (string, string) ([]byte, error)

    HTMLExtract (string, string, *string, *string) (*models.HTMLExtractResponse, error)
}

/*
 * Factory for the DATATOOLS interaface returning DATATOOLS_IMPL
 */
func NewDATATOOLS() DATATOOLS {
    return &DATATOOLS_IMPL{}
}