/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io )
 */

package ecommerce_pkg


import(
	"encoding/json"
	"neutrinoapi_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"neutrinoapi_lib/apihelper_pkg"
	"neutrinoapi_lib/configuration_pkg"
)
/*
 * Client structure as interface implementation
 */
type ECOMMERCE_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Perform a BIN (Bank Identification Number) or IIN (Issuer Identification Number) lookup. See: https://www.neutrinoapi.com/api/bin-lookup/
 * @param    string         binNumber       parameter: Required
 * @param    *string        customerIp      parameter: Optional
 * @return	Returns the *models_pkg.BINLookupResponse response from the API call
 */
func (me *ECOMMERCE_IMPL) BINLookup (
            binNumber string,
            customerIp *string) (*models_pkg.BINLookupResponse, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/bin-lookup"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "user-id" : neutrinoapi_lib.config.UserId,
        "api-key" : neutrinoapi_lib.config.ApiKey,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //form parameters
    parameters := map[string]interface{} {

        "output-case" : "camel",
        "bin-number" : binNumber,
        "customer-ip" : customerIp,

    }


    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Your API request has been rejected. Check the error code for details", _response.Code, _response.RawBody)
    } else if (_response.Code == 403) {
        err = apihelper_pkg.NewAPIError("You have failed to authenticate or are using an invalid API path", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("We messed up, sorry! Your request has caused a fatal exception", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.BINLookupResponse = &models_pkg.BINLookupResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

