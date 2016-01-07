/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 01/07/2016
 */
package imaging

import(
    "github.com/apimatic/unirest-go"
    "neutrinoapi"
    "neutrinoapi/apihelper"

)
/*
 * Client structure as interface implementation
 */
type IMAGING_IMPL struct { }

/**
 * Generate a QR code as a PNG image. See: https://www.neutrinoapi.com/api/qr-code/
 * @param    string         content      parameter: Required
 * @param    *string        bgColor      parameter: Optional
 * @param    *string        fgColor      parameter: Optional
 * @param    *int           height       parameter: Optional
 * @param    *int           width        parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *IMAGING_IMPL) QRCode (
            content string,
            bgColor *string,
            fgColor *string,
            height *int,
            width *int) ([]byte, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/qr-code"

    //variable to hold errors
    var err error = nil
    
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithQueryParameters(queryBuilder, map[string]interface{} {
        "width" : apihelper.ToString(*width, "250"),
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
        "bg-color" : apihelper.ToString(*bgColor, "#ffffff"),
        "fg-color" : apihelper.ToString(*fgColor, "#000000"),
        "height" : apihelper.ToString(*height, "250"),

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
 * Convert HTML content into PDF documents. See: https://www.neutrinoapi.com/api/html-to-pdf/
 * @param    string         content        parameter: Required
 * @param    *int           htmlWidth      parameter: Optional
 * @param    *int           margin         parameter: Optional
 * @param    *string        title          parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *IMAGING_IMPL) HTMLToPDF (
            content string,
            htmlWidth *int,
            margin *int,
            title *string) ([]byte, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/html-to-pdf"

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
        "html-width" : apihelper.ToString(*htmlWidth, "1024"),
        "margin" : apihelper.ToString(*margin, "10"),
        "title" : title,

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
 * Resize an image and output as either JPEG or PNG. See: https://www.neutrinoapi.com/api/image-resize/
 * @param    int            height        parameter: Required
 * @param    string         imageUrl      parameter: Required
 * @param    int            width         parameter: Required
 * @param    *string        format        parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *IMAGING_IMPL) ImageResize (
            height int,
            imageUrl string,
            width int,
            format *string) ([]byte, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/image-resize"

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

        "height" : height,
        "image-url" : imageUrl,
        "width" : width,
        "format" : apihelper.ToString(*format, "png"),

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
 * Watermark one image with another image. See: https://www.neutrinoapi.com/api/image-watermark/
 * @param    string         imageUrl          parameter: Required
 * @param    string         watermarkUrl      parameter: Required
 * @param    *string        format            parameter: Optional
 * @param    *int           height            parameter: Optional
 * @param    *int           opacity           parameter: Optional
 * @param    *string        position          parameter: Optional
 * @param    *int           width             parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *IMAGING_IMPL) ImageWatermark (
            imageUrl string,
            watermarkUrl string,
            format *string,
            height *int,
            opacity *int,
            position *string,
            width *int) ([]byte, error) {
    //the base uri for api requests
    queryBuilder := neutrinoapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/image-watermark"

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

        "image-url" : imageUrl,
        "watermark-url" : watermarkUrl,
        "format" : apihelper.ToString(*format, "png"),
        "height" : height,
        "opacity" : apihelper.ToString(*opacity, "50"),
        "position" : apihelper.ToString(*position, "center"),
        "width" : width,

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

