/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package imaging_pkg


import(
	"errors"
	"fmt"
	"github.com/apimatic/unirest-go"
	"neutrinoapi_lib/apihelper_pkg"
	"neutrinoapi_lib/configuration_pkg"
)
/*
 * Client structure as interface implementation
 */
type IMAGING_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Watermark one image with another image. See: https://www.neutrinoapi.com/api/image-watermark/
 * @param    string         imageUrl          parameter: Required
 * @param    string         watermarkUrl      parameter: Required
 * @param    *int64         opacity           parameter: Optional
 * @param    *string        format            parameter: Optional
 * @param    *string        position          parameter: Optional
 * @param    *int64         width             parameter: Optional
 * @param    *int64         height            parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *IMAGING_IMPL) ImageWatermark (
            imageUrl string,
            watermarkUrl string,
            opacity *int64,
            format *string,
            position *string,
            width *int64,
            height *int64) ([]byte, error) {
    //the endpoint path uri
    _pathUrl := "/image-watermark"

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

        "image-url" : imageUrl,
        "watermark-url" : watermarkUrl,
        "opacity" : apihelper_pkg.ToString(*opacity, "50"),
        "format" : apihelper_pkg.ToString(*format, "png"),
        "position" : apihelper_pkg.ToString(*position, "center"),
        "width" : width,
        "height" : height,

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
 * Generate a QR code as a PNG image. See: https://www.neutrinoapi.com/api/qr-code/
 * @param    string         content      parameter: Required
 * @param    *int64         width        parameter: Optional
 * @param    *int64         height       parameter: Optional
 * @param    *string        fgColor      parameter: Optional
 * @param    *string        bgColor      parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *IMAGING_IMPL) QRCode (
            content string,
            width *int64,
            height *int64,
            fgColor *string,
            bgColor *string) ([]byte, error) {
    //the endpoint path uri
    _pathUrl := "/qr-code"

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
        "width" : apihelper_pkg.ToString(*width, "256"),
        "height" : apihelper_pkg.ToString(*height, "256"),
        "fg-color" : apihelper_pkg.ToString(*fgColor, "#000000"),
        "bg-color" : apihelper_pkg.ToString(*bgColor, "#ffffff"),

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
 * Resize an image and output as either JPEG or PNG. See: https://www.neutrinoapi.com/api/image-resize/
 * @param    string         imageUrl      parameter: Required
 * @param    int64          width         parameter: Required
 * @param    int64          height        parameter: Required
 * @param    *string        format        parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *IMAGING_IMPL) ImageResize (
            imageUrl string,
            width int64,
            height int64,
            format *string) ([]byte, error) {
    //the endpoint path uri
    _pathUrl := "/image-resize"

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

        "image-url" : imageUrl,
        "width" : width,
        "height" : height,
        "format" : apihelper_pkg.ToString(*format, "png"),

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
 * Render HTML content to PDF, JPG or PNG. See: https://www.neutrinoapi.com/api/html5-render/
 * @param    string          content                parameter: Required
 * @param    *string         format                 parameter: Optional
 * @param    *string         pageSize               parameter: Optional
 * @param    *string         title                  parameter: Optional
 * @param    *int64          margin                 parameter: Optional
 * @param    *int64          marginLeft             parameter: Optional
 * @param    *int64          marginRight            parameter: Optional
 * @param    *int64          marginTop              parameter: Optional
 * @param    *int64          marginBottom           parameter: Optional
 * @param    *bool           landscape              parameter: Optional
 * @param    *float64        zoom                   parameter: Optional
 * @param    *bool           grayscale              parameter: Optional
 * @param    *bool           mediaPrint             parameter: Optional
 * @param    *bool           mediaQueries           parameter: Optional
 * @param    *bool           forms                  parameter: Optional
 * @param    *string         css                    parameter: Optional
 * @param    *int64          imageWidth             parameter: Optional
 * @param    *int64          imageHeight            parameter: Optional
 * @param    *int64          renderDelay            parameter: Optional
 * @param    *string         headerTextLeft         parameter: Optional
 * @param    *string         headerTextCenter       parameter: Optional
 * @param    *string         headerTextRight        parameter: Optional
 * @param    *int64          headerSize             parameter: Optional
 * @param    *string         headerFont             parameter: Optional
 * @param    *int64          headerFontSize         parameter: Optional
 * @param    *bool           headerLine             parameter: Optional
 * @param    *string         footerTextLeft         parameter: Optional
 * @param    *string         footerTextCenter       parameter: Optional
 * @param    *string         footerTextRight        parameter: Optional
 * @param    *int64          footerSize             parameter: Optional
 * @param    *string         footerFont             parameter: Optional
 * @param    *int64          footerFontSize         parameter: Optional
 * @param    *bool           footerLine             parameter: Optional
 * @param    *int64          pageWidth              parameter: Optional
 * @param    *int64          pageHeight             parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *IMAGING_IMPL) HTML5Render (
            content string,
            format *string,
            pageSize *string,
            title *string,
            margin *int64,
            marginLeft *int64,
            marginRight *int64,
            marginTop *int64,
            marginBottom *int64,
            landscape *bool,
            zoom *float64,
            grayscale *bool,
            mediaPrint *bool,
            mediaQueries *bool,
            forms *bool,
            css *string,
            imageWidth *int64,
            imageHeight *int64,
            renderDelay *int64,
            headerTextLeft *string,
            headerTextCenter *string,
            headerTextRight *string,
            headerSize *int64,
            headerFont *string,
            headerFontSize *int64,
            headerLine *bool,
            footerTextLeft *string,
            footerTextCenter *string,
            footerTextRight *string,
            footerSize *int64,
            footerFont *string,
            footerFontSize *int64,
            footerLine *bool,
            pageWidth *int64,
            pageHeight *int64) ([]byte, error) {
    //the endpoint path uri
    _pathUrl := "/html5-render"

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

        "output-case" : "camel",
        "content" : content,
        "format" : apihelper_pkg.ToString(*format, "PDF"),
        "page-size" : apihelper_pkg.ToString(*pageSize, "A4"),
        "title" : title,
        "margin" : apihelper_pkg.ToString(*margin, "0"),
        "margin-left" : apihelper_pkg.ToString(*marginLeft, "0"),
        "margin-right" : apihelper_pkg.ToString(*marginRight, "0"),
        "margin-top" : apihelper_pkg.ToString(*marginTop, "0"),
        "margin-bottom" : apihelper_pkg.ToString(*marginBottom, "0"),
        "landscape" : apihelper_pkg.ToString(*landscape, false),
        "zoom" : apihelper_pkg.ToString(*zoom, "1"),
        "grayscale" : apihelper_pkg.ToString(*grayscale, false),
        "media-print" : apihelper_pkg.ToString(*mediaPrint, false),
        "media-queries" : apihelper_pkg.ToString(*mediaQueries, false),
        "forms" : apihelper_pkg.ToString(*forms, false),
        "css" : css,
        "image-width" : apihelper_pkg.ToString(*imageWidth, "1024"),
        "image-height" : imageHeight,
        "render-delay" : apihelper_pkg.ToString(*renderDelay, "0"),
        "header-text-left" : headerTextLeft,
        "header-text-center" : headerTextCenter,
        "header-text-right" : headerTextRight,
        "header-size" : apihelper_pkg.ToString(*headerSize, "9"),
        "header-font" : apihelper_pkg.ToString(*headerFont, "Courier"),
        "header-font-size" : apihelper_pkg.ToString(*headerFontSize, "11"),
        "header-line" : apihelper_pkg.ToString(*headerLine, false),
        "footer-text-left" : footerTextLeft,
        "footer-text-center" : footerTextCenter,
        "footer-text-right" : footerTextRight,
        "footer-size" : apihelper_pkg.ToString(*footerSize, "9"),
        "footer-font" : apihelper_pkg.ToString(*footerFont, "Courier"),
        "footer-font-size" : apihelper_pkg.ToString(*footerFontSize, "11"),
        "footer-line" : apihelper_pkg.ToString(*footerLine, false),
        "page-width" : pageWidth,
        "page-height" : pageHeight,

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

