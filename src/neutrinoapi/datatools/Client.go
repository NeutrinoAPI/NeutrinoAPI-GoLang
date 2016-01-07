/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 01/07/2016
 */
package datatools

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
type DATATOOLS_IMPL struct { }

/**
 * Parse, validate and get location information about a phone number. See: https://www.neutrinoapi.com/api/phone-validate/
 * @param    string         number           parameter: Required
 * @param    *string        countryCode      parameter: Optional
 * @return	Returns the *models.PhoneValidateResponse response from the API call
 */
func (me *DATATOOLS_IMPL) PhoneValidate (
            number string,
            countryCode *string) (*models.PhoneValidateResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/phone-validate"

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

        "number" : number,
        "output-case" : "camel",
        "country-code" : countryCode,

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
    var retVal *models.PhoneValidateResponse = &models.PhoneValidateResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Parse, validate and get detailed user-agent information from a user-agent string. See: https://www.neutrinoapi.com/api/user-agent-info/
 * @param    string        userAgent       parameter: Required
 * @return	Returns the *models.UserAgentInfoResponse response from the API call
 */
func (me *DATATOOLS_IMPL) UserAgentInfo (
            userAgent string) (*models.UserAgentInfoResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/user-agent-info"

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

        "output-case" : "camel",
        "user-agent" : userAgent,

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
    var retVal *models.UserAgentInfoResponse = &models.UserAgentInfoResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Code highlight will take raw source code and convert into nicely formatted HTML with syntax and keyword highlighting. See: https://www.neutrinoapi.com/api/code-highlight/
 * @param    string        content               parameter: Required
 * @param    string        mtype                 parameter: Required
 * @param    *bool         addKeywordLinks       parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *DATATOOLS_IMPL) CodeHighlight (
            content string,
            mtype string,
            addKeywordLinks *bool) ([]byte, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/code-highlight"

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
    }

    //form parameters
    parameters := map[string]interface{} {

        "content" : content,
        "type" : mtype,
        "add-keyword-links" : apihelper.ToString(*addKeywordLinks, "false"),

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
    return response.RawBody, nil
}

/**
 * Detect bad words, swear words and profanity in a given text. See: https://www.neutrinoapi.com/api/bad-word-filter/
 * @param    string         content              parameter: Required
 * @param    *string        censorCharacter      parameter: Optional
 * @return	Returns the *models.BadWordFilterResponse response from the API call
 */
func (me *DATATOOLS_IMPL) BadWordFilter (
            content string,
            censorCharacter *string) (*models.BadWordFilterResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/bad-word-filter"

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

        "content" : content,
        "output-case" : "camel",
        "censor-character" : censorCharacter,

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
    var retVal *models.BadWordFilterResponse = &models.BadWordFilterResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * A powerful unit and currency conversion tool. See: https://www.neutrinoapi.com/api/convert/
 * @param    string        fromType        parameter: Required
 * @param    string        fromValue       parameter: Required
 * @param    string        toType          parameter: Required
 * @return	Returns the *models.ConvertResponse response from the API call
 */
func (me *DATATOOLS_IMPL) Convert (
            fromType string,
            fromValue string,
            toType string) (*models.ConvertResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/convert"

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

        "from-type" : fromType,
        "from-value" : fromValue,
        "output-case" : "camel",
        "to-type" : toType,

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
    var retVal *models.ConvertResponse = &models.ConvertResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Parse, validate and clean an email address. See: https://www.neutrinoapi.com/api/email-validate/
 * @param    string        email           parameter: Required
 * @param    *bool         fixTypos        parameter: Optional
 * @return	Returns the *models.EmailValidateResponse response from the API call
 */
func (me *DATATOOLS_IMPL) EmailValidate (
            email string,
            fixTypos *bool) (*models.EmailValidateResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/email-validate"

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

        "email" : email,
        "output-case" : "camel",
        "fix-typos" : apihelper.ToString(*fixTypos, "false"),

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
    var retVal *models.EmailValidateResponse = &models.EmailValidateResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

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
func (me *DATATOOLS_IMPL) HTMLClean (
            content string,
            outputType string) ([]byte, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/html-clean"

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
    }

    //form parameters
    parameters := map[string]interface{} {

        "content" : content,
        "output-type" : outputType,

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
    return response.RawBody, nil
}

/**
 * Extract HTML tag contents or attributes from complex HTML or XHTML content. See: https://www.neutrinoapi.com/api/html-extract-tags/
 * @param    string         content         parameter: Required
 * @param    string         tag             parameter: Required
 * @param    *string        attribute       parameter: Optional
 * @param    *string        baseUrl         parameter: Optional
 * @return	Returns the *models.HTMLExtractResponse response from the API call
 */
func (me *DATATOOLS_IMPL) HTMLExtract (
            content string,
            tag string,
            attribute *string,
            baseUrl *string) (*models.HTMLExtractResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/html-extract-tags"

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

        "content" : content,
        "output-case" : "camel",
        "tag" : tag,
        "attribute" : attribute,
        "base-url" : baseUrl,

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
    var retVal *models.HTMLExtractResponse = &models.HTMLExtractResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

