/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io )
 */

package configuration_pkg



type CONFIGURATION interface {
        UserId()  string
        ApiKey()  string
        SetUserId(userId   string)
        SetApiKey(apiKey   string)
}   
/*
 * Factory for the CONFIGURATION interaface returning CONFIGURATION_IMPL
 */
func NewCONFIGURATION() CONFIGURATION{
    configuration := new(CONFIGURATION_IMPL)
    return configuration
}