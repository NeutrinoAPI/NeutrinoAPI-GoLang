/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package configuration_pkg


/** The base Uri for API calls */
const BASEURI string = "https://neutrinoapi.com"



type CONFIGURATION_IMPL struct {
    /** Your user ID */
    /** Replace the value of UserId with an appropriate value */
    user-id string
    /** Your API key */
    /** Replace the value of ApiKey with an appropriate value */
    api-key string
}

/*
 * Getter function returning user-id
 */
func (me *CONFIGURATION_IMPL) UserId() string{
    return me.user-id
}

/*
 * Setter function setting up user-id
 */
func (me *CONFIGURATION_IMPL) SetUserId(userId string) {
    me.user-id = userId
}

/*
 * Getter function returning api-key
 */
func (me *CONFIGURATION_IMPL) ApiKey() string{
    return me.api-key
}

/*
 * Setter function setting up api-key
 */
func (me *CONFIGURATION_IMPL) SetApiKey(apiKey string) {
    me.api-key = apiKey
}

