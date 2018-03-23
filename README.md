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

* [imaging_pkg](#imaging_pkg)
* [telephony_pkg](#telephony_pkg)
* [datatools_pkg](#datatools_pkg)
* [securityandnetworking_pkg](#securityandnetworking_pkg)
* [geolocation_pkg](#geolocation_pkg)
* [e-commerce_pkg](#e_commerce_pkg)

## <a name="imaging_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".imaging_pkg") imaging_pkg

### Get instance

Factory for the ``` IMAGING ``` interface can be accessed from the package imaging_pkg.

```go
imaging := imaging_pkg.NewIMAGING()
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
| width |  ``` Required ```  | Width to resize to (in px) |
| height |  ``` Required ```  | Height to resize to (in px) |
| format |  ``` Optional ```  ``` DefaultValue ```  | The output image format, can be either png or jpg |


#### Example Usage

```go
imageUrl := "image-url"
width,_ := strconv.ParseInt("101", 10, 8)
height,_ := strconv.ParseInt("101", 10, 8)
format := "png"

var result []byte
result,_ = imaging.ImageResize(imageUrl, width, height, format)

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
| fgColor |  ``` Optional ```  ``` DefaultValue ```  | The QR code foreground color (you should always use a dark color for this) |
| bgColor |  ``` Optional ```  ``` DefaultValue ```  | The QR code background color (you should always use a light color for this) |


#### Example Usage

```go
content := "content"
width,_ := strconv.ParseInt("250", 10, 8)
height,_ := strconv.ParseInt("250", 10, 8)
fgColor := "#000000"
bgColor := "#ffffff"

var result []byte
result,_ = imaging.QRCode(content, width, height, fgColor, bgColor)

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
| position |  ``` Optional ```  ``` DefaultValue ```  | The position of the watermark image, possible values are: center, top-left, top-center, top-right, bottom-left, bottom-center, bottom-right |
| width |  ``` Optional ```  | If set resize the resulting image to this width (preserving aspect ratio) |
| height |  ``` Optional ```  | If set resize the resulting image to this height (preserving aspect ratio) |


#### Example Usage

```go
imageUrl := "image-url"
watermarkUrl := "watermark-url"
opacity,_ := strconv.ParseInt("50", 10, 8)
format := "png"
position := "center"
width,_ := strconv.ParseInt("101", 10, 8)
height,_ := strconv.ParseInt("101", 10, 8)

var result []byte
result,_ = imaging.ImageWatermark(imageUrl, watermarkUrl, opacity, format, position, width, height)

```


### <a name="html5_render"></a>![Method: ](https://apidocs.io/img/method.png ".imaging_pkg.HTML5Render") HTML5Render

> Render HTML and HTML5 content to PDF, JPG or PNG


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
            headerFontSize *string,
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
| renderDelay |  ``` Optional ```  | Number of milliseconds to wait before rendering the page (can be useful for pages with animations etc) |
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
zoom := "1.0"
grayscale := false
mediaPrint := false
mediaQueries := false
forms := false
css := "css"
imageWidth,_ := strconv.ParseInt("1024", 10, 8)
imageHeight,_ := strconv.ParseInt("101", 10, 8)
renderDelay,_ := strconv.ParseInt("101", 10, 8)
headerTextLeft := "header-text-left"
headerTextCenter := "header-text-center"
headerTextRight := "header-text-right"
headerSize,_ := strconv.ParseInt("9", 10, 8)
headerFont := "Courier"
headerFontSize := "11"
headerLine := false
footerTextLeft := "footer-text-left"
footerTextCenter := "footer-text-center"
footerTextRight := "footer-text-right"
footerSize,_ := strconv.ParseInt("9", 10, 8)
footerFont := "Courier"
footerFontSize,_ := strconv.ParseInt("11", 10, 8)
footerLine := false
pageWidth,_ := strconv.ParseInt("101", 10, 8)
pageHeight,_ := strconv.ParseInt("101", 10, 8)

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

### <a name="hlr_lookup"></a>![Method: ](https://apidocs.io/img/method.png ".telephony_pkg.HLRLookup") HLRLookup

> Connect to the global mobile cellular network and retrieve the status of a mobile device


```go
func (me *TELEPHONY_IMPL) HLRLookup(
            number string,
            countryCode *string)(*models_pkg.HLRLookupResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | A phone number |
| countryCode |  ``` Optional ```  | ISO 2-letter country code, assume numbers are based in this country. If not set numbers are assumed to be in international format (with or without the leading + sign) |


#### Example Usage

```go
number := "number"
countryCode := "country-code"

var result *models_pkg.HLRLookupResponse
result,_ = telephony.HLRLookup(number, countryCode)

```


### <a name="phone_playback"></a>![Method: ](https://apidocs.io/img/method.png ".telephony_pkg.PhonePlayback") PhonePlayback

> Make an automated call to any valid phone number and playback an audio message


```go
func (me *TELEPHONY_IMPL) PhonePlayback(
            number string,
            audioUrl string)(*models_pkg.PhonePlaybackResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | The phone number to call. Must be valid international format |
| audioUrl |  ``` Required ```  | A URL to a valid audio file. Accepted audio formats are: MP3, WAV, OGG |


#### Example Usage

```go
number := "number"
audioUrl := "audio-url"

var result *models_pkg.PhonePlaybackResponse
result,_ = telephony.PhonePlayback(number, audioUrl)

```


### <a name="verify_security_code"></a>![Method: ](https://apidocs.io/img/method.png ".telephony_pkg.VerifySecurityCode") VerifySecurityCode

> Check if a security code from one of the verify APIs is valid


```go
func (me *TELEPHONY_IMPL) VerifySecurityCode(securityCode int64)(*models_pkg.VerifySecurityCodeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| securityCode |  ``` Required ```  | The security code to verify |


#### Example Usage

```go
securityCode,_ := strconv.ParseInt("101", 10, 8)

var result *models_pkg.VerifySecurityCodeResponse
result,_ = telephony.VerifySecurityCode(securityCode)

```


### <a name="sms_verify"></a>![Method: ](https://apidocs.io/img/method.png ".telephony_pkg.SMSVerify") SMSVerify

> Send a unique security code to any mobile device via SMS


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
| securityCode |  ``` Optional ```  | ass in your own security code. This is useful if you have implemented TOTP or similar 2FA methods. If not set then we will generate a secure random code (only numerical security codes are currently supported) |
| countryCode |  ``` Optional ```  | ISO 2-letter country code, assume numbers are based in this country. If not set numbers are assumed to be in international format (with or without the leading + sign) |
| languageCode |  ``` Optional ```  ``` DefaultValue ```  | The language to send the verification code in, available languages are: de - German, en - English, es - Spanish, fr - Fench, it - Italian, pt - Portuguese, ru - Russian |


#### Example Usage

```go
number := "number"
codeLength,_ := strconv.ParseInt("5", 10, 8)
securityCode,_ := strconv.ParseInt("101", 10, 8)
countryCode := "country-code"
languageCode := "en"

var result *models_pkg.SMSVerifyResponse
result,_ = telephony.SMSVerify(number, codeLength, securityCode, countryCode, languageCode)

```


### <a name="phone_verify"></a>![Method: ](https://apidocs.io/img/method.png ".telephony_pkg.PhoneVerify") PhoneVerify

> Make an automated call to any valid phone number and playback a unique security code


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
| securityCode |  ``` Optional ```  | Pass in your own security code. This is useful if you have implemented TOTP or similar 2FA methods. If not set then we will generate a secure random code (only numerical security codes are currently supported) |
| playbackDelay |  ``` Optional ```  ``` DefaultValue ```  | The delay in milliseconds between the playback of each security code |
| countryCode |  ``` Optional ```  | ISO 2-letter country code, assume numbers are based in this country. If not set numbers are assumed to be in international format (with or without the leading + sign) |
| languageCode |  ``` Optional ```  ``` DefaultValue ```  | The language to playback the verification code in, available languages are: de - German, en - English, es - Spanish, fr - Fench, it - Italian, pt - Portuguese, ru - Russian |


#### Example Usage

```go
number := "number"
codeLength,_ := strconv.ParseInt("6", 10, 8)
securityCode,_ := strconv.ParseInt("101", 10, 8)
playbackDelay,_ := strconv.ParseInt("800", 10, 8)
countryCode := "country-code"
languageCode := "en"

var result *models_pkg.PhoneVerifyResponse
result,_ = telephony.PhoneVerify(number, codeLength, securityCode, playbackDelay, countryCode, languageCode)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="datatools_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".datatools_pkg") datatools_pkg

### Get instance

Factory for the ``` DATATOOLS ``` interface can be accessed from the package datatools_pkg.

```go
dataTools := datatools_pkg.NewDATATOOLS()
```

### <a name="email_validate"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.EmailValidate") EmailValidate

> Parse, validate and clean an email address


```go
func (me *DATATOOLS_IMPL) EmailValidate(
            email string,
            fixTypos *bool)(*models_pkg.EmailValidateResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email address |
| fixTypos |  ``` Optional ```  ``` DefaultValue ```  | Automatically attempt to fix typos in the address |


#### Example Usage

```go
email := "email"
fixTypos := false

var result *models_pkg.EmailValidateResponse
result,_ = dataTools.EmailValidate(email, fixTypos)

```


### <a name="bad_word_filter"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.BadWordFilter") BadWordFilter

> Detect bad words, swear words and profanity in a given text


```go
func (me *DATATOOLS_IMPL) BadWordFilter(
            content string,
            censorCharacter *string)(*models_pkg.BadWordFilterResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| content |  ``` Required ```  | The text content to check. This can be either a URL to load content from or an actual content string |
| censorCharacter |  ``` Optional ```  | The character to use to censor out the bad words found |


#### Example Usage

```go
content := "content"
censorCharacter := "censor-character"

var result *models_pkg.BadWordFilterResponse
result,_ = dataTools.BadWordFilter(content, censorCharacter)

```


### <a name="convert"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.Convert") Convert

> A powerful unit and currency conversion tool


```go
func (me *DATATOOLS_IMPL) Convert(
            fromValue string,
            fromType string,
            toType string)(*models_pkg.ConvertResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| fromValue |  ``` Required ```  | The value to convert from |
| fromType |  ``` Required ```  | The type of the value to convert from |
| toType |  ``` Required ```  | The type to convert to |


#### Example Usage

```go
fromValue := "from-value"
fromType := "from-type"
toType := "to-type"

var result *models_pkg.ConvertResponse
result,_ = dataTools.Convert(fromValue, fromType, toType)

```


### <a name="phone_validate"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.PhoneValidate") PhoneValidate

> Parse, validate and get location information about a phone number


```go
func (me *DATATOOLS_IMPL) PhoneValidate(
            number string,
            countryCode *string,
            ip *string)(*models_pkg.PhoneValidateResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | The phone number |
| countryCode |  ``` Optional ```  | ISO 2-letter country code, assume numbers are based in this country. If not set numbers are assumed to be in international format (with or without the leading + sign) |
| ip |  ``` Optional ```  | Pass in a users IP address and we will assume numbers are based in the country of the IP address |


#### Example Usage

```go
number := "number"
countryCode := "country-code"
ip := "ip"

var result *models_pkg.PhoneValidateResponse
result,_ = dataTools.PhoneValidate(number, countryCode, ip)

```


### <a name="user_agent_info"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.UserAgentInfo") UserAgentInfo

> Parse, validate and get detailed user-agent information from a user agent string


```go
func (me *DATATOOLS_IMPL) UserAgentInfo(userAgent string)(*models_pkg.UserAgentInfoResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| userAgent |  ``` Required ```  | A user-agent string |


#### Example Usage

```go
userAgent := "user-agent"

var result *models_pkg.UserAgentInfoResponse
result,_ = dataTools.UserAgentInfo(userAgent)

```


### <a name="html_clean"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.HTMLClean") HTMLClean

> Clean and sanitize untrusted HTML


```go
func (me *DATATOOLS_IMPL) HTMLClean(
            content string,
            outputType string)([]byte,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| content |  ``` Required ```  | The HTML content. This can be either a URL to load HTML from or an actual HTML content string |
| outputType |  ``` Required ```  | The level of sanitization, possible values are: plain-text, simple-text, basic-html, basic-html-with-images, advanced-html |


#### Example Usage

```go
content := "content"
outputType := "output-type"

var result []byte
result,_ = dataTools.HTMLClean(content, outputType)

```


### <a name="html_extract"></a>![Method: ](https://apidocs.io/img/method.png ".datatools_pkg.HTMLExtract") HTMLExtract

> Extract specific HTML tag contents or attributes from complex HTML or XHTML content


```go
func (me *DATATOOLS_IMPL) HTMLExtract(
            content string,
            tag string,
            attribute *string,
            baseUrl *string)(*models_pkg.HTMLExtractResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| content |  ``` Required ```  | The HTML content. This can be either a URL to load HTML from or an actual HTML content string |
| tag |  ``` Required ```  | The HTML tag(s) to extract data from. This can just be a simple tag name like 'img' OR a CSS/jQuery style selector |
| attribute |  ``` Optional ```  | If set, then extract data from the specified tag attribute. If not set, then data will be extracted from the tags inner content |
| baseUrl |  ``` Optional ```  | The base URL to replace into realive links |


#### Example Usage

```go
content := "content"
tag := "tag"
attribute := "attribute"
baseUrl := "base-url"

var result *models_pkg.HTMLExtractResponse
result,_ = dataTools.HTMLExtract(content, tag, attribute, baseUrl)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="securityandnetworking_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".securityandnetworking_pkg") securityandnetworking_pkg

### Get instance

Factory for the ``` SECURITYANDNETWORKING ``` interface can be accessed from the package securityandnetworking_pkg.

```go
securityAndNetworking := securityandnetworking_pkg.NewSECURITYANDNETWORKING()
```

### <a name="host_reputation"></a>![Method: ](https://apidocs.io/img/method.png ".securityandnetworking_pkg.HostReputation") HostReputation

> Check the reputation of an IP address or domain against a comprehensive list of blacklists and blocklists (DNSBLs)


```go
func (me *SECURITYANDNETWORKING_IMPL) HostReputation(host string)(*models_pkg.HostReputationResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| host |  ``` Required ```  | An IPv4 address or a domain name. If you supply a domain name it will be checked against the URI DNSBL list |


#### Example Usage

```go
host := "host"

var result *models_pkg.HostReputationResponse
result,_ = securityAndNetworking.HostReputation(host)

```


### <a name="url_info"></a>![Method: ](https://apidocs.io/img/method.png ".securityandnetworking_pkg.URLInfo") URLInfo

> Parse, analyze and retrieve content from the supplied URL


```go
func (me *SECURITYANDNETWORKING_IMPL) URLInfo(
            url string,
            fetchContent bool)(*models_pkg.URLInfoResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| url |  ``` Required ```  | The URL to process |
| fetchContent |  ``` Required ```  | If this URL responds with html, text, json or xml then return the response. This option is useful if you want to perform further processing on the URL content |


#### Example Usage

```go
url := "url"
fetchContent := false

var result *models_pkg.URLInfoResponse
result,_ = securityAndNetworking.URLInfo(url, fetchContent)

```


### <a name="ip_probe"></a>![Method: ](https://apidocs.io/img/method.png ".securityandnetworking_pkg.IPProbe") IPProbe

> Analyze and extract provider information for an IP address


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


### <a name="ip_blocklist"></a>![Method: ](https://apidocs.io/img/method.png ".securityandnetworking_pkg.IPBlocklist") IPBlocklist

> The IP Blocklist API will detect potentially malicious or dangerous IP addresses


```go
func (me *SECURITYANDNETWORKING_IMPL) IPBlocklist(ip string)(*models_pkg.IPBlocklistResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| ip |  ``` Required ```  | An IPv4 address |


#### Example Usage

```go
ip := "ip"

var result *models_pkg.IPBlocklistResponse
result,_ = securityAndNetworking.IPBlocklist(ip)

```


### <a name="email_verify"></a>![Method: ](https://apidocs.io/img/method.png ".securityandnetworking_pkg.EmailVerify") EmailVerify

> SMTP based email address verification


```go
func (me *SECURITYANDNETWORKING_IMPL) EmailVerify(
            email string,
            fixTypos *bool)(*models_pkg.EmailVerifyResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | An email address |
| fixTypos |  ``` Optional ```  | Automatically attempt to fix typos in the address |


#### Example Usage

```go
email := "email"
fixTypos := false

var result *models_pkg.EmailVerifyResponse
result,_ = securityAndNetworking.EmailVerify(email, fixTypos)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="geolocation_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".geolocation_pkg") geolocation_pkg

### Get instance

Factory for the ``` GEOLOCATION ``` interface can be accessed from the package geolocation_pkg.

```go
geolocation := geolocation_pkg.NewGEOLOCATION()
```

### <a name="geocode_reverse"></a>![Method: ](https://apidocs.io/img/method.png ".geolocation_pkg.GeocodeReverse") GeocodeReverse

> Convert a geographic coordinate (latitude and longitude) into a real world address or location.


```go
func (me *GEOLOCATION_IMPL) GeocodeReverse(
            latitude float64,
            longitude float64,
            languageCode *string)(*models_pkg.GeocodeReverseResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| latitude |  ``` Required ```  | The location latitude |
| longitude |  ``` Required ```  | The location longitude |
| languageCode |  ``` Optional ```  ``` DefaultValue ```  | The language to display results in, available languages are: de, en, es, fr, it, pt, ru |


#### Example Usage

```go
latitude := 101.592891887572
longitude := 101.592891887572
languageCode := "en"

var result *models_pkg.GeocodeReverseResponse
result,_ = geolocation.GeocodeReverse(latitude, longitude, languageCode)

```


### <a name="ip_info"></a>![Method: ](https://apidocs.io/img/method.png ".geolocation_pkg.IPInfo") IPInfo

> Get location information about an IP address and do reverse DNS (PTR) lookups.


```go
func (me *GEOLOCATION_IMPL) IPInfo(
            ip string,
            reverseLookup *bool)(*models_pkg.IPInfoResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| ip |  ``` Required ```  | The IP address |
| reverseLookup |  ``` Optional ```  ``` DefaultValue ```  | Do a reverse DNS (PTR) lookup. This option can add extra delay to the request so only use it if you need it |


#### Example Usage

```go
ip := "ip"
reverseLookup := false

var result *models_pkg.IPInfoResponse
result,_ = geolocation.IPInfo(ip, reverseLookup)

```


### <a name="geocode_address"></a>![Method: ](https://apidocs.io/img/method.png ".geolocation_pkg.GeocodeAddress") GeocodeAddress

> Geocode an address, partial address or the name of a location


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
| address |  ``` Required ```  | The address or partial address to try and locate |
| countryCode |  ``` Optional ```  | The ISO 2-letter country code to be biased towards (default is no country bias) |
| languageCode |  ``` Optional ```  ``` DefaultValue ```  | The language to display results in, available languages are: de, en, es, fr, it, pt, ru |
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
| customerIp |  ``` Optional ```  | Pass in a customers remote IP address. The API will then determine the country of the IP address and match it against the BIN country. This feature is designed for fraud prevention and detection checks. |


#### Example Usage

```go
binNumber := "bin-number"
customerIp := "customer-ip"

var result *models_pkg.BINLookupResponse
result,_ = eCommerce.BINLookup(binNumber, customerIp)

```


[Back to List of Controllers](#list_of_controllers)



