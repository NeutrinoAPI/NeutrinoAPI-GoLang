# Getting started

The general-purpose API

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=NeutrinoAPI-GoLang&projectName=neutrinoapi_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=neutrinoapi_lib)

## How to Use

The following section explains how to use the NeutrinoAPI library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=neutrinoapi_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=neutrinoapi_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=neutrinoapi_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=neutrinoapi_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=neutrinoapi_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=NeutrinoAPI-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=neutrinoapi_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=neutrinoapi_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| userId | Your user ID |
| apiKey | Your API key |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [www_pkg](#www_pkg)
* [imaging_pkg](#imaging_pkg)
* [telephony_pkg](#telephony_pkg)
* [e-commerce_pkg](#e_commerce_pkg)
* [geolocation_pkg](#geolocation_pkg)
* [securityandnetworking_pkg](#securityandnetworking_pkg)
* [datatools_pkg](#datatools_pkg)

## <a name="www_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".www_pkg") www_pkg

### Get instance

Factory for the ``` WWW ``` interface can be accessed from the package www_pkg.

```go
wWW := www_pkg.NewWWW()
```

### <a name="browser_bot"></a>![Method: ](https://apidocs.io/img/method.png ".www_pkg.BrowserBot") BrowserBot

> Browser bot can extract content, interact with keyboard and mouse events, and execute JavaScript on a website. See: https://www.neutrinoapi.com/api/browser-bot/


```go
func (me *WWW_IMPL) BrowserBot(
            url string,
            timeout *int64,
            delay *int64,
            selector *string,
            exec []string,
            userAgent *string,
            ignoreCertificateErrors *bool)(*models_pkg.BrowserBotResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| url |  ``` Required ```  | The URL to load |
| timeout |  ``` Optional ```  ``` DefaultValue ```  | Timeout in seconds. Give up if still trying to load the page after this number of seconds |
| delay |  ``` Optional ```  ``` DefaultValue ```  | Delay in seconds to wait before executing any selectors or JavaScript |
| selector |  ``` Optional ```  | Extract content from the page DOM using this selector. Commonly known as a CSS selector, you can find a good reference <a href="https://www.w3schools.com/cssref/css_selectors.asp" target="_blank">here</a> |
| exec |  ``` Optional ```  ``` Collection ```  ``` DefaultValue ```  | Execute JavaScript on the page. Each array element should contain a valid JavaScript statement in string form. If a statement returns any kind of value it will be returned in the 'exec-results' response.<br/><br/>For your convenience you can also use the following special shortcut functions:<br/><div style='padding-left:32px; font-family:inherit; font-size:inherit;'>sleep(seconds); Just wait/sleep for the specified number of seconds.<br/>click('selector'); Click on the first element matching the given selector.<br/>focus('selector'); Focus on the first element matching the given selector.<br/>keys('characters'); Send the specified keyboard characters. Use click() or focus() first to send keys to a specific element.<br/>enter(); Send the Enter key.<br/>tab(); Send the Tab key.<br/></div><br/>Example:<br/><div style='padding-left:32px; font-family:inherit; font-size:inherit;'>[ "click('#button-id')", "sleep(1)", "click('.field-class')", "keys('1234')", "enter()" ]</div> |
| userAgent |  ``` Optional ```  | Override the browsers default user-agent string with this one |
| ignoreCertificateErrors |  ``` Optional ```  ``` DefaultValue ```  | Ignore any TLS/SSL certificate errors and load the page anyway |


#### Example Usage

```go
url := "url"
timeout,_ := strconv.ParseInt("30", 10, 8)
delay,_ := strconv.ParseInt("2", 10, 8)
selector := "selector"
execValue := []byte("[]")
var exec []string
json.Unmarshal(execValue,&exec)
userAgent := "user-agent"
ignoreCertificateErrors := false

var result *models_pkg.BrowserBotResponse
result,_ = wWW.BrowserBot(url, timeout, delay, selector, exec, userAgent, ignoreCertificateErrors)

```


### <a name="html_clean"></a>![Method: ](https://apidocs.io/img/method.png ".www_pkg.HTMLClean") HTMLClean

> Clean and sanitize untrusted HTML. See: https://www.neutrinoapi.com/api/html-clean/


```go
func (me *WWW_IMPL) HTMLClean(
            content string,
            outputType string)([]byte,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| content |  ``` Required ```  | The HTML content. This can be either a URL to load HTML from or an actual HTML content string |
| outputType |  ``` Required ```  | The level of sanitization, possible values are:<br/><b>plain-text</b>: reduce the content to plain text only (no HTML tags at all)<br/><br/><b>simple-text</b>: allow only very basic text formatting tags like b, em, i, strong, u<br/><br/><b>basic-html</b>: allow advanced text formatting and hyper links<br/><br/><b>basic-html-with-images</b>: same as basic html but also allows image tags<br/><br/><b>advanced-html</b>: same as basic html with images but also allows many more common HTML tags like table, ul, dl, pre<br/> |


#### Example Usage

```go
content := "content"
outputType := "output-type"

var result []byte
result,_ = wWW.HTMLClean(content, outputType)

```


### <a name="url_info"></a>![Method: ](https://apidocs.io/img/method.png ".www_pkg.URLInfo") URLInfo

> Parse, analyze and retrieve content from the supplied URL. See: https://www.neutrinoapi.com/api/url-info/


```go
func (me *WWW_IMPL) URLInfo(
            url string,
            fetchContent *bool)(*models_pkg.URLInfoResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| url |  ``` Required ```  | The URL to probe |
| fetchContent |  ``` Optional ```  ``` DefaultValue ```  | If this URL responds with html, text, json or xml then return the response. This option is useful if you want to perform further processing on the URL content (e.g. with the HTML Extract or HTML Clean APIs) |


#### Example Usage

```go
url := "url"
fetchContent := false

var result *models_pkg.URLInfoResponse
result,_ = wWW.URLInfo(url, fetchContent)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="imaging_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".imaging_pkg") imaging_pkg

### Get instance

Factory for the ``` IMAGING ``` interface can be accessed from the package imaging_pkg.

```go
imaging := imaging_pkg.NewIMAGING()
```

### <a name="image_watermark"></a>![Method: ](https://apidocs.io/img/method.png ".imaging_pkg.ImageWatermark") ImageWatermark

> Watermark one image with another image. See: https://www.neutrinoapi.com/api/image-watermark/


```go
func (me *IMAGING_IMPL) ImageWatermark(
            imageUrl string,
            watermarkUrl string,
            opacity *int64,
            format *string,
            position *string,
            width *int64,
            height *int64)([]byte,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| imageUrl |  ``` Required ```  | The URL to the source image |
| watermarkUrl |  ``` Required ```  | The URL to the watermark image |
| opacity |  ``` Optional ```  ``` DefaultValue ```  | The opacity of the watermark (0 to 100) |
| format |  ``` Optional ```  ``` DefaultValue ```  | The output image format, can be either png or jpg |
| position |  ``` Optional ```  ``` DefaultValue ```  | The position of the watermark image, possible values are:<br/>center, top-left, top-center, top-right, bottom-left, bottom-center, bottom-right |
| width |  ``` Optional ```  | If set resize the resulting image to this width (in px) while preserving aspect ratio |
| height |  ``` Optional ```  | If set resize the resulting image to this height (in px) while preserving aspect ratio |


#### Example Usage

```go
imageUrl := "image-url"
watermarkUrl := "watermark-url"
opacity,_ := strconv.ParseInt("50", 10, 8)
format := "png"
position := "center"
width,_ := strconv.ParseInt("147", 10, 8)
height,_ := strconv.ParseInt("147", 10, 8)

var result []byte
result,_ = imaging.ImageWatermark(imageUrl, watermarkUrl, opacity, format, position, width, height)

```


### <a name="qr_code"></a>![Method: ](https://apidocs.io/img/method.png ".imaging_pkg.QRCode") QRCode

> Generate a QR code as a PNG image. See: https://www.neutrinoapi.com/api/qr-code/


```go
func (me *IMAGING_IMPL) QRCode(
            content string,
            width *int64,
            height *int64,
            fgColor *string,
            bgColor *string)([]byte,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| content |  ``` Required ```  | The content to encode into the QR code (e.g. a URL or a phone number) |
| width |  ``` Optional ```  ``` DefaultValue ```  | The width of the QR code (in px) |
| height |  ``` Optional ```  ``` DefaultValue ```  | The height of the QR code (in px) |
| fgColor |  ``` Optional ```  ``` DefaultValue ```  | The QR code foreground color |
| bgColor |  ``` Optional ```  ``` DefaultValue ```  | The QR code background color |


#### Example Usage

```go
content := "content"
width,_ := strconv.ParseInt("256", 10, 8)
height,_ := strconv.ParseInt("256", 10, 8)
fgColor := "#000000"
bgColor := "#ffffff"

var result []byte
result,_ = imaging.QRCode(content, width, height, fgColor, bgColor)

```


### <a name="image_resize"></a>![Method: ](https://apidocs.io/img/method.png ".imaging_pkg.ImageResize") ImageResize

> Resize an image and output as either JPEG or PNG. See: https://www.neutrinoapi.com/api/image-resize/


```go
func (me *IMAGING_IMPL) ImageResize(
            imageUrl string,
            width int64,
            height int64,
            format *string)([]byte,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| imageUrl |  ``` Required ```  | The URL to the source image |
| width |  ``` Required ```  | The width to resize to (in px) while preserving aspect ratio |
| height |  ``` Required ```  | The height to resize to (in px) while preserving aspect ratio |
| format |  ``` Optional ```  ``` DefaultValue ```  | The output image format, can be either png or jpg |


#### Example Usage

```go
imageUrl := "image-url"
width,_ := strconv.ParseInt("147", 10, 8)
height,_ := strconv.ParseInt("147", 10, 8)
format := "png"

var result []byte
result,_ = imaging.ImageResize(imageUrl, width, height, format)

```


### <a name="html5_render"></a>![Method: ](https://apidocs.io/img/method.png ".imaging_pkg.HTML5Render") HTML5Render

> Render HTML content to PDF, JPG or PNG. See: https://www.neutrinoapi.com/api/html5-render/


```go
func (me *IMAGING_IMPL) HTML5Render(
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
            pageHeight *int64)([]byte,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| content |  ``` Required ```  | The HTML content. This can be either a URL to load HTML from or an actual HTML content string |
| format |  ``` Optional ```  ``` DefaultValue ```  | Which format to output, available options are: PDF, PNG, JPG |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Set the document page size, can be one of: A0 - A9, B0 - B10, Comm10E, DLE or Letter |
| title |  ``` Optional ```  | The document title |
| margin |  ``` Optional ```  ``` DefaultValue ```  | The document margin (in mm) |
| marginLeft |  ``` Optional ```  ``` DefaultValue ```  | The document left margin (in mm) |
| marginRight |  ``` Optional ```  ``` DefaultValue ```  | The document right margin (in mm) |
| marginTop |  ``` Optional ```  ``` DefaultValue ```  | The document top margin (in mm) |
| marginBottom |  ``` Optional ```  ``` DefaultValue ```  | The document bottom margin (in mm) |
| landscape |  ``` Optional ```  ``` DefaultValue ```  | Set the document to lanscape orientation |
| zoom |  ``` Optional ```  ``` DefaultValue ```  | Set the zoom factor when rendering the page (2.0 for double size, 0.5 for half size) |
| grayscale |  ``` Optional ```  ``` DefaultValue ```  | Render the final document in grayscale |
| mediaPrint |  ``` Optional ```  ``` DefaultValue ```  | Use @media print CSS styles to render the document |
| mediaQueries |  ``` Optional ```  ``` DefaultValue ```  | Activate all @media queries before rendering. This can be useful if you wan't to render the mobile version of a responsive website |
| forms |  ``` Optional ```  ``` DefaultValue ```  | Generate real (fillable) PDF forms from HTML forms |
| css |  ``` Optional ```  | Inject custom CSS into the HTML. e.g. 'body { background-color: red;}' |
| imageWidth |  ``` Optional ```  ``` DefaultValue ```  | If rendering to an image format (PNG or JPG) use this image width (in pixels) |
| imageHeight |  ``` Optional ```  | If rendering to an image format (PNG or JPG) use this image height (in pixels). The default is automatic which dynamically sets the image height based on the content |
| renderDelay |  ``` Optional ```  ``` DefaultValue ```  | Number of milliseconds to wait before rendering the page (can be useful for pages with animations etc) |
| headerTextLeft |  ``` Optional ```  | Text to print to the left-hand side header of each page. e.g. 'My header - Page {page_number} of {total_pages}' |
| headerTextCenter |  ``` Optional ```  | Text to print to the center header of each page |
| headerTextRight |  ``` Optional ```  | Text to print to the right-hand side header of each page |
| headerSize |  ``` Optional ```  ``` DefaultValue ```  | The height of your header (in mm) |
| headerFont |  ``` Optional ```  ``` DefaultValue ```  | Set the header font. Fonts available: Times, Courier, Helvetica, Arial |
| headerFontSize |  ``` Optional ```  ``` DefaultValue ```  | Set the header font size (in pt) |
| headerLine |  ``` Optional ```  ``` DefaultValue ```  | Draw a full page width horizontal line under your header |
| footerTextLeft |  ``` Optional ```  | Text to print to the left-hand side footer of each page. e.g. 'My footer - Page {page_number} of {total_pages}' |
| footerTextCenter |  ``` Optional ```  | Text to print to the center header of each page |
| footerTextRight |  ``` Optional ```  | Text to print to the right-hand side header of each page |
| footerSize |  ``` Optional ```  ``` DefaultValue ```  | The height of your footer (in mm) |
| footerFont |  ``` Optional ```  ``` DefaultValue ```  | Set the footer font. Fonts available: Times, Courier, Helvetica, Arial |
| footerFontSize |  ``` Optional ```  ``` DefaultValue ```  | Set the footer font size (in pt) |
| footerLine |  ``` Optional ```  ``` DefaultValue ```  | Draw a full page width horizontal line above your footer |
| pageWidth |  ``` Optional ```  | Set the PDF page width explicitly (in mm) |
| pageHeight |  ``` Optional ```  | Set the PDF page height explicitly (in mm) |


#### Example Usage

```go
content := "content"
format := "PDF"
pageSize := "A4"
title := "title"
margin,_ := strconv.ParseInt("0", 10, 8)
marginLeft,_ := strconv.ParseInt("0", 10, 8)
marginRight,_ := strconv.ParseInt("0", 10, 8)
marginTop,_ := strconv.ParseInt("0", 10, 8)
marginBottom,_ := strconv.ParseInt("0", 10, 8)
landscape := false
zoom := "1"
grayscale := false
mediaPrint := false
mediaQueries := false
forms := false
css := "css"
imageWidth,_ := strconv.ParseInt("1024", 10, 8)
imageHeight,_ := strconv.ParseInt("147", 10, 8)
renderDelay,_ := strconv.ParseInt("0", 10, 8)
headerTextLeft := "header-text-left"
headerTextCenter := "header-text-center"
headerTextRight := "header-text-right"
headerSize,_ := strconv.ParseInt("9", 10, 8)
headerFont := "Courier"
headerFontSize,_ := strconv.ParseInt("11", 10, 8)
headerLine := false
footerTextLeft := "footer-text-left"
footerTextCenter := "footer-text-center"
footerTextRight := "footer-text-right"
footerSize,_ := strconv.ParseInt("9", 10, 8)
footerFont := "Courier"
footerFontSize,_ := strconv.ParseInt("11", 10, 8)
footerLine := false
pageWidth,_ := strconv.ParseInt("147", 10, 8)
pageHeight,_ := strconv.ParseInt("147", 10, 8)

var result []byte
result,_ = imaging.HTML5Render(content, format, pageSize, title, margin, marginLeft, marginRight, marginTop, marginBottom, landscape, zoom, grayscale, mediaPrint, mediaQueries, forms, css, imageWidth, imageHeight, renderDelay, headerTextLeft, headerTextCenter, headerTextRight, headerSize, headerFont, headerFontSize, headerLine, footerTextLeft, footerTextCenter, footerTextRight, footerSize, footerFont, footerFontSize, footerLine, pageWidth, pageHeight)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="telephony_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".telephony_pkg") telephony_pkg

### Get instance

Factory for the ``` TELEPHONY ``` interface can be accessed from the package telephony_pkg.

```go
telephony := telephony_pkg.NewTELEPHONY()
```

### <a name="phone_verify"></a>![Method: ](https://apidocs.io/img/method.png ".telephony_pkg.PhoneVerify") PhoneVerify

> Make an automated call to any valid phone number and playback a unique security code. See: https://www.neutrinoapi.com/api/phone-verify/


```go
func (me *TELEPHONY_IMPL) PhoneVerify(
            number string,
            codeLength *int64,
            securityCode *int64,
            playbackDelay *int64,
            countryCode *string,
            languageCode *string)(*models_pkg.PhoneVerifyResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | The phone number to send the verification code to |
| codeLength |  ``` Optional ```  ``` DefaultValue ```  | The number of digits to use in the security code (between 4 and 12) |
| securityCode |  ``` Optional ```  | Pass in your own security code. This is useful if you have implemented TOTP or similar 2FA methods. If not set then we will generate a secure random code |
| playbackDelay |  ``` Optional ```  ``` DefaultValue ```  | The delay in milliseconds between the playback of each security code |
| countryCode |  ``` Optional ```  | ISO 2-letter country code, assume numbers are based in this country.<br/>If not set numbers are assumed to be in international format (with or without the leading + sign) |
| languageCode |  ``` Optional ```  ``` DefaultValue ```  | The language to playback the verification code in, available languages are:<ul><li>de - German</li><li>en - English</li><li>es - Spanish</li><li>fr - French</li><li>it - Italian</li><li>pt - Portuguese</li><li>ru - Russian</li></ul> |


#### Example Usage

```go
number := "number"
codeLength,_ := strconv.ParseInt("6", 10, 8)
securityCode,_ := strconv.ParseInt("147", 10, 8)
playbackDelay,_ := strconv.ParseInt("800", 10, 8)
countryCode := "country-code"
languageCode := "en"

var result *models_pkg.PhoneVerifyResponse
result,_ = telephony.PhoneVerify(number, codeLength, securityCode, playbackDelay, countryCode, languageCode)

```


### <a name="sms_message"></a>![Method: ](https://apidocs.io/img/method.png ".telephony_pkg.SMSMessage") SMSMessage

> Send a free-form message to any mobile device via SMS. See: https://www.neutrinoapi.com/api/sms-message/


```go
func (me *TELEPHONY_IMPL) SMSMessage(
            number string,
            message string,
            countryCode *string)(*models_pkg.SMSMessageResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | The phone number to send a message to |
| message |  ``` Required ```  | The SMS message to send. Messages are truncated to a maximum of 150 characters for ASCII content OR 70 characters for UTF content |
| countryCode |  ``` Optional ```  | ISO 2-letter country code, assume numbers are based in this country.<br/>If not set numbers are assumed to be in international format (with or without the leading + sign) |


#### Example Usage

```go
number := "number"
message := "message"
countryCode := "country-code"

var result *models_pkg.SMSMessageResponse
result,_ = telephony.SMSMessage(number, message, countryCode)

```


### <a name="sms_verify"></a>![Method: ](https://apidocs.io/img/method.png ".telephony_pkg.SMSVerify") SMSVerify

> Send a unique security code to any mobile device via SMS. See: https://www.neutrinoapi.com/api/sms-verify/


```go
func (me *TELEPHONY_IMPL) SMSVerify(
            number string,
            codeLength *int64,
            securityCode *int64,
            countryCode *string,
            languageCode *string)(*models_pkg.SMSVerifyResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | The phone number to send a verification code to |
| codeLength |  ``` Optional ```  ``` DefaultValue ```  | The number of digits to use in the security code (must be between 4 and 12) |
| securityCode |  ``` Optional ```  | Pass in your own security code. This is useful if you have implemented TOTP or similar 2FA methods. If not set then we will generate a secure random code |
| countryCode |  ``` Optional ```  | ISO 2-letter country code, assume numbers are based in this country.<br/>If not set numbers are assumed to be in international format (with or without the leading + sign) |
| languageCode |  ``` Optional ```  ``` DefaultValue ```  | The language to send the verification code in, available languages are:<ul><li>de - German</li><li>en - English</li><li>es - Spanish</li><li>fr - French</li><li>it - Italian</li><li>pt - Portuguese</li><li>ru - Russian</li></ul> |


#### Example Usage

```go
number := "number"
codeLength,_ := strconv.ParseInt("5", 10, 8)
securityCode,_ := strconv.ParseInt("147", 10, 8)
countryCode := "country-code"
languageCode := "en"

var result *models_pkg.SMSVerifyResponse
result,_ = telephony.SMSVerify(number, codeLength, securityCode, countryCode, languageCode)

```


### <a name="verify_security_code"></a>![Method: ](https://apidocs.io/img/method.png ".telephony_pkg.VerifySecurityCode") VerifySecurityCode

> Check if a security code from one of the verify APIs is valid. See: https://www.neutrinoapi.com/api/verify-security-code/


```go
func (me *TELEPHONY_IMPL) VerifySecurityCode(securityCode string)(*models_pkg.VerifySecurityCodeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| securityCode |  ``` Required ```  | The security code to verify |


#### Example Usage

```go
securityCode := "security-code"

var result *models_pkg.VerifySecurityCodeResponse
result,_ = telephony.VerifySecurityCode(securityCode)

```


### <a name="phone_playback"></a>![Method: ](https://apidocs.io/img/method.png ".telephony_pkg.PhonePlayback") PhonePlayback

> Make an automated call to any valid phone number and playback an audio message. See: https://www.neutrinoapi.com/api/phone-playback/


```go
func (me *TELEPHONY_IMPL) PhonePlayback(
            number string,
            audioUrl string)(*models_pkg.PhonePlaybackResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | The phone number to call. Must be in valid international format |
| audioUrl |  ``` Required ```  | A URL to a valid audio file. Accepted audio formats are:<ul><li>MP3</li><li>WAV</li><li>OGG</ul></ul>You can use the following MP3 URL for testing:<br/>https://www.neutrinoapi.com/test-files/test1.mp3 |


#### Example Usage

```go
number := "number"
audioUrl := "audio-url"

var result *models_pkg.PhonePlaybackResponse
result,_ = telephony.PhonePlayback(number, audioUrl)

```


### <a name="hlr_lookup"></a>![Method: ](https://apidocs.io/img/method.png ".telephony_pkg.HLRLookup") HLRLookup

> Connect to the global mobile cellular network and retrieve the status of a mobile device. See: https://www.neutrinoapi.com/api/hlr-lookup/


```go
func (me *TELEPHONY_IMPL) HLRLookup(
            number string,
            countryCode *string)(*models_pkg.HLRLookupResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | A phone number |
| countryCode |  ``` Optional ```  | ISO 2-letter country code, assume numbers are based in this country.<br/>If not set numbers are assumed to be in international format (with or without the leading + sign) |


#### Example Usage

```go
number := "number"
countryCode := "country-code"

var result *models_pkg.HLRLookupResponse
result,_ = telephony.HLRLookup(number, countryCode)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="e_commerce_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".e-commerce_pkg") e-commerce_pkg

### Get instance

Factory for the ``` E-COMMERCE ``` interface can be accessed from the package e-commerce_pkg.

```go
eCommerce := e-commerce_pkg.NewE-COMMERCE()
```

### <a name="bin_lookup"></a>![Method: ](https://apidocs.io/img/method.png ".e-commerce_pkg.BINLookup") BINLookup

> Perform a BIN (Bank Identification Number) or IIN (Issuer Identification Number) lookup. See: https://www.neutrinoapi.com/api/bin-lookup/


```go
func (me *ECOMMERCE_IMPL) BINLookup(
            binNumber string,
            customerIp *string)(*models_pkg.BINLookupResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| binNumber |  ``` Required ```  | The BIN or IIN number (the first 6 digits of a credit card number) |
| customerIp |  ``` Optional ```  | Pass in the customers IP address and we will return some extra information about them |


#### Example Usage

```go
binNumber := "bin-number"
customerIp := "customer-ip"

var result *models_pkg.BINLookupResponse
result,_ = eCommerce.BINLookup(binNumber, customerIp)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="geolocation_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".geolocation_pkg") geolocation_pkg

### Get instance

Factory for the ``` GEOLOCATION ``` interface can be accessed from the package geolocation_pkg.

```go
geolocation := geolocation_pkg.NewGEOLOCATION()
```

### <a name="geocode_address"></a>![Method: ](https://apidocs.io/img/method.png ".geolocation_pkg.GeocodeAddress") GeocodeAddress

> Geocode an address, partial address or just the name of a place. See: https://www.neutrinoapi.com/api/geocode-address/


```go
func (me *GEOLOCATION_IMPL) GeocodeAddress(
            address string,
            countryCode *string,
            languageCode *string,
            fuzzySearch *bool)(*models_pkg.GeocodeAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| address |  ``` Required ```  | The address, partial address or name of a place to try and locate |
| countryCode |  ``` Optional ```  | The ISO 2-letter country code to be biased towards (the default is no country bias) |
| languageCode |  ``` Optional ```  ``` DefaultValue ```  | The language to display results in, available languages are:<ul><li>de, en, es, fr, it, pt, ru</li></ul> |
| fuzzySearch |  ``` Optional ```  ``` DefaultValue ```  | If no matches are found for the given address, start performing a recursive fuzzy search until a geolocation is found. We use a combination of approximate string matching and data cleansing to find possible location matches |


#### Example Usage

```go
address := "address"
countryCode := "country-code"
languageCode := "en"
fuzzySearch := false

var result *models_pkg.GeocodeAddressResponse
result,_ = geolocation.GeocodeAddress(address, countryCode, languageCode, fuzzySearch)

```


### <a name="ip_info"></a>![Method: ](https://apidocs.io/img/method.png ".geolocation_pkg.IPInfo") IPInfo

> Get location information about an IP address and do reverse DNS (PTR) lookups. See: https://www.neutrinoapi.com/api/ip-info/


```go
func (me *GEOLOCATION_IMPL) IPInfo(
            ip string,
            reverseLookup *bool)(*models_pkg.IPInfoResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| ip |  ``` Required ```  | IPv4 or IPv6 address |
| reverseLookup |  ``` Optional ```  ``` DefaultValue ```  | Do a reverse DNS (PTR) lookup. This option can add extra delay to the request so only use it if you need it |


#### Example Usage

```go
ip := "ip"
reverseLookup := false

var result *models_pkg.IPInfoResponse
result,_ = geolocation.IPInfo(ip, reverseLookup)

```


### <a name="geocode_reverse"></a>![Method: ](https://apidocs.io/img/method.png ".geolocation_pkg.GeocodeReverse") GeocodeReverse

> Convert a geographic coordinate (latitude and longitude) into a real world address or location. See: https://www.neutrinoapi.com/api/geocode-reverse/


```go
func (me *GEOLOCATION_IMPL) GeocodeReverse(
            latitude string,
            longitude string,
            languageCode *string)(*models_pkg.GeocodeReverseResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| latitude |  ``` Required ```  | The location latitude in decimal degrees format |
| longitude |  ``` Required ```  | The location longitude in decimal degrees format |
| languageCode |  ``` Optional ```  ``` DefaultValue ```  | The language to display results in, available languages are:<ul><li>de, en, es, fr, it, pt, ru</li></ul> |


#### Example Usage

```go
latitude := "latitude"
longitude := "longitude"
languageCode := "en"

var result *models_pkg.GeocodeReverseResponse
result,_ = geolocation.GeocodeReverse(latitude, longitude, languageCode)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="securityandnetworking_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".securityandnetworking_pkg") securityandnetworking_pkg

### Get instance

Factory for the ``` SECURITYANDNETWORKING ``` interface can be accessed from the package securityandnetworking_pkg.

```go
securityAndNetworking := securityandnetworking_pkg.NewSECURITYANDNETWORKING()
```

### <a name="ip_blocklist"></a>![Method: ](https://apidocs.io/img/method.png ".securityandnetworking_pkg.IPBlocklist") IPBlocklist

> The IP Blocklist API will detect potentially malicious or dangerous IP addresses. See: https://www.neutrinoapi.com/api/ip-blocklist/


```go
func (me *SECURITYANDNETWORKING_IMPL) IPBlocklist(ip string)(*models_pkg.IPBlocklistResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| ip |  ``` Required ```  | An IPv4 or IPv6 address |


#### Example Usage

```go
ip := "ip"

var result *models_pkg.IPBlocklistResponse
result,_ = securityAndNetworking.IPBlocklist(ip)

```


### <a name="email_verify"></a>![Method: ](https://apidocs.io/img/method.png ".securityandnetworking_pkg.EmailVerify") EmailVerify

> SMTP based email address verification. See: https://www.neutrinoapi.com/api/email-verify/


```go
func (me *SECURITYANDNETWORKING_IMPL) EmailVerify(
            email string,
            fixTypos *bool)(*models_pkg.EmailVerifyResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | An email address |
| fixTypos |  ``` Optional ```  ``` DefaultValue ```  | Automatically attempt to fix typos in the address |


#### Example Usage

```go
email := "email"
fixTypos := false

var result *models_pkg.EmailVerifyResponse
result,_ = securityAndNetworking.EmailVerify(email, fixTypos)

```


### <a name="host_reputation"></a>![Method: ](https://apidocs.io/img/method.png ".securityandnetworking_pkg.HostReputation") HostReputation

> Check the reputation of an IP address, domain name, FQDN or URL against a comprehensive list of blacklists and blocklists. See: https://www.neutrinoapi.com/api/host-reputation/


```go
func (me *SECURITYANDNETWORKING_IMPL) HostReputation(
            host string,
            listRating *int64)(*models_pkg.HostReputationResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| host |  ``` Required ```  | An IP address, domain name, FQDN or URL.<br/>If you supply a domain/URL it will be checked against the URI DNSBL lists |
| listRating |  ``` Optional ```  ``` DefaultValue ```  | Only check lists with this rating or better |


#### Example Usage

```go
host := "host"
listRating,_ := strconv.ParseInt("3", 10, 8)

var result *models_pkg.HostReputationResponse
result,_ = securityAndNetworking.HostReputation(host, listRating)

```


### <a name="ip_probe"></a>![Method: ](https://apidocs.io/img/method.png ".securityandnetworking_pkg.IPProbe") IPProbe

> Analyze and extract provider information for an IP address. See: https://www.neutrinoapi.com/api/ip-probe/


```go
func (me *SECURITYANDNETWORKING_IMPL) IPProbe(ip string)(*models_pkg.IPProbeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| ip |  ``` Required ```  | IPv4 or IPv6 address |


#### Example Usage

```go
ip := "ip"

var result *models_pkg.IPProbeResponse
result,_ = securityAndNetworking.IPProbe(ip)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="datatools_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".datatools_pkg") datatools_pkg

### Get instance

Factory for the ``` DATATOOLS ``` interface can be accessed from the package datatools_pkg.

```go
dataTools := datatools_pkg.NewDATATOOLS()
```

### <a name="user_agent_info"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.UserAgentInfo") UserAgentInfo

> Parse, validate and get detailed user-agent information from a user agent string. See: https://www.neutrinoapi.com/api/user-agent-info/


```go
func (me *DATATOOLS_IMPL) UserAgentInfo(userAgent string)(*models_pkg.UserAgentInfoResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| userAgent |  ``` Required ```  | A user agent string |


#### Example Usage

```go
userAgent := "user-agent"

var result *models_pkg.UserAgentInfoResponse
result,_ = dataTools.UserAgentInfo(userAgent)

```


### <a name="phone_validate"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.PhoneValidate") PhoneValidate

> Parse, validate and get location information about a phone number. See: https://www.neutrinoapi.com/api/phone-validate/


```go
func (me *DATATOOLS_IMPL) PhoneValidate(
            number string,
            countryCode *string,
            ip *string)(*models_pkg.PhoneValidateResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | A phone number. This can be in international format (E.164) or local format. If passing local format you should use the 'country-code' or 'ip' options as well |
| countryCode |  ``` Optional ```  | ISO 2-letter country code, assume numbers are based in this country.<br/>If not set numbers are assumed to be in international format (with or without the leading + sign) |
| ip |  ``` Optional ```  | Pass in a users IP address and we will assume numbers are based in the country of the IP address |


#### Example Usage

```go
number := "number"
countryCode := "country-code"
ip := "ip"

var result *models_pkg.PhoneValidateResponse
result,_ = dataTools.PhoneValidate(number, countryCode, ip)

```


### <a name="convert"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.Convert") Convert

> A powerful unit conversion tool. See: https://www.neutrinoapi.com/api/convert/


```go
func (me *DATATOOLS_IMPL) Convert(
            fromValue string,
            fromType string,
            toType string)(*models_pkg.ConvertResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| fromValue |  ``` Required ```  | The value to convert from (e.g. 10.95) |
| fromType |  ``` Required ```  | The type of the value to convert from (e.g. USD) |
| toType |  ``` Required ```  | The type to convert to (e.g. EUR) |


#### Example Usage

```go
fromValue := "from-value"
fromType := "from-type"
toType := "to-type"

var result *models_pkg.ConvertResponse
result,_ = dataTools.Convert(fromValue, fromType, toType)

```


### <a name="bad_word_filter"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.BadWordFilter") BadWordFilter

> Detect bad words, swear words and profanity in a given text. See: https://www.neutrinoapi.com/api/bad-word-filter/


```go
func (me *DATATOOLS_IMPL) BadWordFilter(
            content string,
            censorCharacter *string)(*models_pkg.BadWordFilterResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| content |  ``` Required ```  | The content to scan. This can be either a URL to load content from or an actual content string |
| censorCharacter |  ``` Optional ```  | The character to use to censor out the bad words found |


#### Example Usage

```go
content := "content"
censorCharacter := "censor-character"

var result *models_pkg.BadWordFilterResponse
result,_ = dataTools.BadWordFilter(content, censorCharacter)

```


### <a name="email_validate"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.EmailValidate") EmailValidate

> Parse, validate and clean an email address. See: https://www.neutrinoapi.com/api/email-validate/


```go
func (me *DATATOOLS_IMPL) EmailValidate(
            email string,
            fixTypos *bool)(*models_pkg.EmailValidateResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | An email address |
| fixTypos |  ``` Optional ```  ``` DefaultValue ```  | Automatically attempt to fix typos in the address |


#### Example Usage

```go
email := "email"
fixTypos := false

var result *models_pkg.EmailValidateResponse
result,_ = dataTools.EmailValidate(email, fixTypos)

```


[Back to List of Controllers](#list_of_controllers)



