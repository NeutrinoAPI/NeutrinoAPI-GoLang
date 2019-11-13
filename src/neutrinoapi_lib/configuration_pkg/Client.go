/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package configuration_pkg

import(
	"neutrinoapi_lib/apihelper_pkg"
)
/* Setting up enums for Environments and Servers 
*/
type Environments int

type Servers int

// Environment Enums
const (
     MULTICLOUD Environments = 1 + iota
     //AWS endpoint
     AWS
     //GCP endpoint
     GCP
     //MS Azure endpoint
     MSA
)

// Servers Enums
const (
 	 ENUM_DEFAULT Servers = 1 + iota
)

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

// Setting up Default Environment
var Environment = MULTICLOUD

// A map of environments and their corresponding servers/baseurls
var EnvironmentsMap = map[Environments](map[Servers]string){

    MULTICLOUD : map[Servers]string{
        ENUM_DEFAULT:"https://neutrinoapi.net/",
    },

    AWS : map[Servers]string{
        ENUM_DEFAULT:"https://aws.neutrinoapi.net/",
    },

    GCP : map[Servers]string{
        ENUM_DEFAULT:"https://gcp.neutrinoapi.net/",
    },

    MSA : map[Servers]string{
        ENUM_DEFAULT:"https://msa.neutrinoapi.net/",
    },
}
 
// Make and return the map of parameters
func GetBaseURIParameters(config CONFIGURATION) map[string]interface{} {
     kvpMap := map[string]interface{}{
    }
    return kvpMap;
}

// Gets the URL for a particular alias in the current environment and appends it with template parameters
// return the baseurl
func GetBaseURI(server Servers, config CONFIGURATION) string {
    url := EnvironmentsMap[Environment][server]
    appendedURL, _ := apihelper_pkg.AppendUrlWithTemplateParameters(url, GetBaseURIParameters(config))
    return appendedURL

}
