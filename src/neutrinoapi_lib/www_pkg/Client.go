/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package www_pkg


import(
	"errors"
	"fmt"
	"encoding/json"
	"neutrinoapi_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"neutrinoapi_lib/apihelper_pkg"
	"neutrinoapi_lib/configuration_pkg"
)
/*
 * Client structure as interface implementation
 */
type WWW_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Browser bot can extract content, interact with keyboard and mouse events, and execute JavaScript on a website. See: https://www.neutrinoapi.com/api/browser-bot/
 * @param    string          url                           parameter: Required
 * @param    *int64          timeout                       parameter: Optional
 * @param    *int64          delay                         parameter: Optional
 * @param    *string         selector                      parameter: Optional
 * @param    []string        exec                          parameter: Optional
 * @param    *string         userAgent                     parameter: Optional
 * @param    *bool           ignoreCertificateErrors       parameter: Optional
 * @return	Returns the *models_pkg.BrowserBotResponse response from the API call
 */
func (me *WWW_IMPL) BrowserBot (
            url string,
            timeout *int64,
            delay *int64,
            selector *string,
            exec []string,
            userAgent *string,
            ignoreCertificateErrors *bool) (*models_pkg.BrowserBotResponse, error) {
    //the endpoint path uri
    _pathUrl := "/browser-bot"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

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
        "url" : url,
        "timeout" : apihelper_pkg.ToString(*timeout, "30"),
        "delay" : apihelper_pkg.ToString(*delay, "2"),
        "selector" : selector,
        "exec" : apihelper_pkg.ToString(*exec, "[]"),
        "user-agent" : userAgent,
        "ignore-certificate-errors" : apihelper_pkg.ToString(*ignoreCertificateErrors, false),

    }


    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
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
    var retVal *models_pkg.BrowserBotResponse = &models_pkg.BrowserBotResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Clean and sanitize untrusted HTML. See: https://www.neutrinoapi.com/api/html-clean/
 * @param    string        content         parameter: Required
 * @param    string        outputType      parameter: Required
 * @return	Returns the []byte response from the API call
 */
func (me *WWW_IMPL) HTMLClean (
            content string,
            outputType string) ([]byte, error) {
    //the endpoint path uri
    _pathUrl := "/html-clean"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

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
    }

    //form parameters
    parameters := map[string]interface{} {

        "content" : content,
        "output-type" : outputType,

    }


    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
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
    return _response.RawBody, nil

}

/**
 * Parse, analyze and retrieve content from the supplied URL. See: https://www.neutrinoapi.com/api/url-info/
 * @param    string        url               parameter: Required
 * @param    *bool         fetchContent      parameter: Optional
 * @return	Returns the *models_pkg.URLInfoResponse response from the API call
 */
func (me *WWW_IMPL) URLInfo (
            url string,
            fetchContent *bool) (*models_pkg.URLInfoResponse, error) {
    //the endpoint path uri
    _pathUrl := "/url-info"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

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
        "url" : url,
        "fetch-content" : apihelper_pkg.ToString(*fetchContent, false),

    }


    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
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
    var retVal *models_pkg.URLInfoResponse = &models_pkg.URLInfoResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

