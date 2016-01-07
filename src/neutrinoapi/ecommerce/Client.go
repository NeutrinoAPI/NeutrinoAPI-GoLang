/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 01/07/2016
 */
package ecommerce

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
type ECOMMERCE_IMPL struct { }

/**
 * Perform a BIN (Bank Identification Number) or IIN (Issuer Identification Number) lookup. See: https://www.neutrinoapi.com/api/bin-lookup/
 * @param    string         binNumber       parameter: Required
 * @param    *string        customerIp      parameter: Optional
 * @return	Returns the *models.BINLookupResponse response from the API call
 */
func (me *ECOMMERCE_IMPL) BINLookup (
            binNumber string,
            customerIp *string) (*models.BINLookupResponse, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/bin-lookup"

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

        "bin-number" : binNumber,
        "output-case" : "camel",
        "customer-ip" : customerIp,

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
    var retVal *models.BINLookupResponse = &models.BINLookupResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

