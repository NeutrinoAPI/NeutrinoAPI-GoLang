/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 12/10/2015
 */
package telephony

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
type TELEPHONY_IMPL struct { }

/**
 * Make an automated call to any valid phone number and playback an audio message. See: https://www.neutrinoapi.com/api/phone-playback/
 * @param    string        audioUrl        parameter: Required
 * @param    string        number          parameter: Required
 * @return	Returns the *models.PhonePlaybackResponse response from the API call
 */
func (me *TELEPHONY_IMPL) CreatePhonePlayback (
            audioUrl string,
            number string) (*models.PhonePlaybackResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/phone-playback"

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

        "audio-url" : audioUrl,
        "number" : number,
        "output-case" : "camel",

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
    var retVal *models.PhonePlaybackResponse = &models.PhonePlaybackResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Check if a security code from one of the verify APIs is valid. See: https://www.neutrinoapi.com/api/verify-security-code/
 * @param    int           securityCode      parameter: Required
 * @return	Returns the *models.VerifySecurityCodeResponse response from the API call
 */
func (me *TELEPHONY_IMPL) CreateVerifySecurityCode (
            securityCode int) (*models.VerifySecurityCodeResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/verify-security-code"

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
        "security-code" : securityCode,

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
    var retVal *models.VerifySecurityCodeResponse = &models.VerifySecurityCodeResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Mobile network HLR lookup service. See: https://www.neutrinoapi.com/api/hlr-lookup/
 * @param    string         number           parameter: Required
 * @param    *string        countryCode      parameter: Optional
 * @return	Returns the *models.HLRLookupResponse response from the API call
 */
func (me *TELEPHONY_IMPL) CreateHLRLookup (
            number string,
            countryCode *string) (*models.HLRLookupResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/hlr-lookup"

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
    var retVal *models.HLRLookupResponse = &models.HLRLookupResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Make an automated call to any valid phone number and playback a unique security code. See: https://www.neutrinoapi.com/api/phone-verify/
 * @param    string         number             parameter: Required
 * @param    *int           codeLength         parameter: Optional
 * @param    *string        countryCode        parameter: Optional
 * @param    *string        languageCode       parameter: Optional
 * @param    *int           playbackDelay      parameter: Optional
 * @param    *int           securityCode       parameter: Optional
 * @return	Returns the *models.PhoneVerifyResponse response from the API call
 */
func (me *TELEPHONY_IMPL) CreatePhoneVerify (
            number string,
            codeLength *int,
            countryCode *string,
            languageCode *string,
            playbackDelay *int,
            securityCode *int) (*models.PhoneVerifyResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/phone-verify"

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
        "code-length" : apihelper.ToString(*codeLength, "6"),
        "country-code" : countryCode,
        "language-code" : apihelper.ToString(*languageCode, "en"),
        "playback-delay" : apihelper.ToString(*playbackDelay, "800"),
        "security-code" : securityCode,

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
    var retVal *models.PhoneVerifyResponse = &models.PhoneVerifyResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Send a unique security code to any mobile device via SMS. See: https://www.neutrinoapi.com/api/sms-verify/
 * @param    string         number            parameter: Required
 * @param    *int           codeLength        parameter: Optional
 * @param    *string        countryCode       parameter: Optional
 * @param    *string        languageCode      parameter: Optional
 * @param    *int           securityCode      parameter: Optional
 * @return	Returns the *models.SMSVerifyResponse response from the API call
 */
func (me *TELEPHONY_IMPL) CreateSMSVerify (
            number string,
            codeLength *int,
            countryCode *string,
            languageCode *string,
            securityCode *int) (*models.SMSVerifyResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/sms-verify"

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
        "code-length" : apihelper.ToString(*codeLength, "5"),
        "country-code" : countryCode,
        "language-code" : apihelper.ToString(*languageCode, "en"),
        "security-code" : securityCode,

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
    var retVal *models.SMSVerifyResponse = &models.SMSVerifyResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

