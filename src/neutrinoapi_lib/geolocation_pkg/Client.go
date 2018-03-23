/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io )
 */

package geolocation_pkg


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
type GEOLOCATION_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Convert a geographic coordinate (latitude and longitude) into a real world address or location.
 * @param    float64        latitude          parameter: Required
 * @param    float64        longitude         parameter: Required
 * @param    *string        languageCode      parameter: Optional
 * @return	Returns the *models_pkg.GeocodeReverseResponse response from the API call
 */
func (me *GEOLOCATION_IMPL) GeocodeReverse (
            latitude float64,
            longitude float64,
            languageCode *string) (*models_pkg.GeocodeReverseResponse, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/geocode-reverse"

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
        "latitude" : latitude,
        "longitude" : longitude,
        "language-code" : apihelper_pkg.ToString(*languageCode, "en"),

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
    var retVal *models_pkg.GeocodeReverseResponse = &models_pkg.GeocodeReverseResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Get location information about an IP address and do reverse DNS (PTR) lookups.
 * @param    string        ip                 parameter: Required
 * @param    *bool         reverseLookup      parameter: Optional
 * @return	Returns the *models_pkg.IPInfoResponse response from the API call
 */
func (me *GEOLOCATION_IMPL) IPInfo (
            ip string,
            reverseLookup *bool) (*models_pkg.IPInfoResponse, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/ip-info"

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
        "ip" : ip,
        "reverse-lookup" : apihelper_pkg.ToString(*reverseLookup, false),

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
    var retVal *models_pkg.IPInfoResponse = &models_pkg.IPInfoResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Geocode an address, partial address or the name of a location
 * @param    string         address           parameter: Required
 * @param    *string        countryCode       parameter: Optional
 * @param    *string        languageCode      parameter: Optional
 * @param    *bool          fuzzySearch       parameter: Optional
 * @return	Returns the *models_pkg.GeocodeAddressResponse response from the API call
 */
func (me *GEOLOCATION_IMPL) GeocodeAddress (
            address string,
            countryCode *string,
            languageCode *string,
            fuzzySearch *bool) (*models_pkg.GeocodeAddressResponse, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/geocode-address"

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
        "address" : address,
        "country-code" : countryCode,
        "language-code" : apihelper_pkg.ToString(*languageCode, "en"),
        "fuzzy-search" : apihelper_pkg.ToString(*fuzzySearch, false),

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
    var retVal *models_pkg.GeocodeAddressResponse = &models_pkg.GeocodeAddressResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

