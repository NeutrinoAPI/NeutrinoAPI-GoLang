/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 01/07/2016
 */
package geolocation

import(
    "encoding/json"
    "github.com/apimatic/unirest-go"
    "neutrinoapi"
    "neutrinoapi/apihelper"
    "neutrinoapi/models"
)
/*
 * Client structure as interface implementation
 */
type GEOLOCATION_IMPL struct { }

/**
 * Get location information about an IP address and do reverse DNS (PTR) lookups. See: https://www.neutrinoapi.com/api/ip-info/
 * @param    string        ip                 parameter: Required
 * @param    *bool         reverseLookup      parameter: Optional
 * @return	Returns the *models.IPInfoResponse response from the API call
 */
func (me *GEOLOCATION_IMPL) IPInfo (
            ip string,
            reverseLookup *bool) (*models.IPInfoResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/ip-info"

    //variable to hold errors
    var err error = nil
    
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithQueryParameters(queryBuilder, map[string]interface{} {
        "user-id" : neutrinoapi.Config.UserId,
        "api-key" : neutrinoapi.Config.ApiKey,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
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

        "ip" : ip,
        "output-case" : "camel",
        "reverse-lookup" : apihelper.ToString(*reverseLookup, "false"),

    }


    //prepare API request
    request := unirest.Post(queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal *models.IPInfoResponse = &models.IPInfoResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Geocode an address or partial address. See: https://www.neutrinoapi.com/api/geocode-address/
 * @param    string         address           parameter: Required
 * @param    *string        countryCode       parameter: Optional
 * @param    *string        languageCode      parameter: Optional
 * @return	Returns the *models.GeocodeAddressResponse response from the API call
 */
func (me *GEOLOCATION_IMPL) GeocodeAddress (
            address string,
            countryCode *string,
            languageCode *string) (*models.GeocodeAddressResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/geocode-address"

    //variable to hold errors
    var err error = nil
    
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithQueryParameters(queryBuilder, map[string]interface{} {
        "user-id" : neutrinoapi.Config.UserId,
        "api-key" : neutrinoapi.Config.ApiKey,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
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

        "address" : address,
        "output-case" : "camel",
        "country-code" : countryCode,
        "language-code" : apihelper.ToString(*languageCode, "en"),

    }


    //prepare API request
    request := unirest.Post(queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal *models.GeocodeAddressResponse = &models.GeocodeAddressResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Reverse geocoding is the process of taking a coordinate (latitude and longitude) and mapping this to a real world address or location. See: https://www.neutrinoapi.com/api/geocode-reverse/
 * @param    float32        latitude          parameter: Required
 * @param    float32        longitude         parameter: Required
 * @param    *string        languageCode      parameter: Optional
 * @return	Returns the *models.GeocodeReverseResponse response from the API call
 */
func (me *GEOLOCATION_IMPL) GeocodeReverse (
            latitude float32,
            longitude float32,
            languageCode *string) (*models.GeocodeReverseResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/geocode-reverse"

    //variable to hold errors
    var err error = nil
    
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithQueryParameters(queryBuilder, map[string]interface{} {
        "user-id" : neutrinoapi.Config.UserId,
        "api-key" : neutrinoapi.Config.ApiKey,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
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

        "latitude" : latitude,
        "longitude" : longitude,
        "output-case" : "camel",
        "language-code" : apihelper.ToString(*languageCode, "en"),

    }


    //prepare API request
    request := unirest.Post(queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal *models.GeocodeReverseResponse = &models.GeocodeReverseResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

