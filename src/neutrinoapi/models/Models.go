/*
 * neutrinoapi
 *
 * This file was automatically generated for NeutrinoAPI.com by APIMATIC BETA v2.0 on 12/10/2015
 */

package models


/*
 * Structure for the custom type VerifySecurityCodeResponse
 */
type VerifySecurityCodeResponse struct {
    Verified        bool            `json:"verified"` //True if the code is valid
}

/*
 * Structure for the custom type IPBlocklistResponse
 */
type IPBlocklistResponse struct {
    IsBot           bool            `json:"isBot"` //IP is hosting a malicious bot or is part of a botnet
    IsExploitBot    bool            `json:"isExploitBot"` //IP is hosting an exploit finding bot or exploit scanning software
    IsMalware       bool            `json:"isMalware"` //IP is involved in distributing malware
    IsSpider        bool            `json:"isSpider"` //IP is a hostile spider or crawler
    IsDshield       bool            `json:"isDshield"` //IP has been flagged on DShield (dshield.org)
    ListCount       int             `json:"listCount"` //The number of blocklists the IP is listed on
    IsProxy         bool            `json:"isProxy"` //IP has been detected as an anonymous web proxy or anonymous HTTP proxy
    IsHijacked      bool            `json:"isHijacked"` //hijacked netblocks or netblocks controlled by criminal organizations
    IsTor           bool            `json:"isTor"` //IP is coming from a Tor node
    IsSpyware       bool            `json:"isSpyware"` //IP is being used by spyware, malware, botnets or for other malicious activities
    IsSpamBot       bool            `json:"isSpamBot"` //IP address is hosting a spam bot, comment spamming or other spamming software
    IsListed        bool            `json:"isListed"` //Is this IP on a blocklist
    IsVpn           bool            `json:"isVpn"` //IP has been detected as coming from a VPN hosting provider
}

/*
 * Structure for the custom type HLRLookupResponse
 */
type HLRLookupResponse struct {
    NumberValid              bool            `json:"numberValid"` //Is this a valid phone number (mobile or otherwise)
    InternationalCallingCode string          `json:"internationalCallingCode"` //Numbers international calling code
    Mnc                      string          `json:"mnc"` //The mobile MNC number (only set if HLR lookup valid)
    NumberType               string          `json:"numberType"` //The number type, possible values are: mobile, fixed-line, premium-rate, toll-free, voip, unknown
    HlrValid                 bool            `json:"hlrValid"` //Was the HLR lookup successful. If true then this is a working and registered cell-phone or mobile device (SMS and phone calls will be delivered)
    HlrStatus                string          `json:"hlrStatus"` //The HLR lookup status. See API docs for specific status details
    PortedNetwork            string          `json:"portedNetwork"` //If the number has been ported, the ported to mobile carrier name (only set if HLR lookup valid)
    Imsi                     string          `json:"imsi"` //The mobile IMSI number (only set if HLR lookup valid)
    Mcc                      string          `json:"mcc"` //The mobile MCC number (only set if HLR lookup valid)
    InternationalNumber      string          `json:"internationalNumber"` //Number represented in international format
    LocalNumber              string          `json:"localNumber"` //Number represented in local format
    CountryCode              string          `json:"countryCode"` //Number location ISO 2-letter country code
    IsPorted                 bool            `json:"isPorted"` //Has this number been ported to another network
    Msin                     string          `json:"msin"` //The mobile MSIN number (only set if HLR lookup valid)
    Location                 string          `json:"location"` //Number location (could be a city, region or country)
    OriginNetwork            string          `json:"originNetwork"` //The origin mobile carrier name (only set if HLR lookup valid)
    IsMobile                 bool            `json:"isMobile"` //Is this a mobile number
}

/*
 * Structure for the custom type BINLookupResponse
 */
type BINLookupResponse struct {
    Country         string          `json:"country"` //Full country name of the issuer
    IpCity          string          `json:"ipCity"` //The city name (if detectable) from the customer IP
    IpMatchesBin    bool            `json:"ipMatchesBin"` //True if the customer IP address country matches the BIN country
    CardType        string          `json:"cardType"` //The card type, will always be one of: DEBIT, CREDIT, CHARGE CARD
    CardCategory    string          `json:"cardCategory"` //The card category (if known)
    IpCountryCode   string          `json:"ipCountryCode"` //The ISO 2-letter country code detected from the customer IP
    IpCountry       string          `json:"ipCountry"` //The country detected from the customer IP
    Issuer          string          `json:"issuer"` //The card issuer (if known)
    IpBlocklisted   bool            `json:"ipBlocklisted"` //True if the customer IP is listed on one of our blocklists, see the IP Blocklist API for more details
    Valid           bool            `json:"valid"` //Is this a valid BIN or IIN number
    IpBlocklists    []string        `json:"ipBlocklists"` //An array of strings indicating which blocklists this IP is listed on
    IssuerWebsite   string          `json:"issuerWebsite"` //The card issuer website (if known)
    CountryCode     string          `json:"countryCode"` //ISO 2-letter country code of the issuer
    IpRegion        string          `json:"ipRegion"` //The region name (if detectable) from the customer IP
    CardBrand       string          `json:"cardBrand"` //The card brand (e.g. Visa or Mastercard)
    IssuerPhone     string          `json:"issuerPhone"` //The card issuer phone number (if known)
}

/*
 * Structure for the custom type HostReputationResponse
 */
type HostReputationResponse struct {
    IsListed        bool            `json:"isListed"` //Is this host blacklisted
    Lists           []*Blacklist    `json:"lists"` //An array of objects for each DNSBL checked
    ListCount       int             `json:"listCount"` //The number of DNSBL's the host is listed on
}

/*
 * Structure for the custom type GeocodeAddressResponse
 */
type GeocodeAddressResponse struct {
    Found           int             `json:"found"` //The number of possible matching locations found
    Locations       []*Location     `json:"locations"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type BadWordFilterResponse
 */
type BadWordFilterResponse struct {
    BadWordsList    []string        `json:"badWordsList"` //Array of the bad words found
    BadWordsTotal   int             `json:"badWordsTotal"` //Total number of bad words detected
    CensoredContent string          `json:"censoredContent"` //The censored content (only set if censor-character has been set)
    IsBad           bool            `json:"isBad"` //Does the text contain bad words
}

/*
 * Structure for the custom type PhonePlaybackResponse
 */
type PhonePlaybackResponse struct {
    Calling         bool            `json:"calling"` //True if the call is being made now
    NumberValid     bool            `json:"number-valid"` //Is this a valid phone number
}

/*
 * Structure for the custom type HTMLExtractResponse
 */
type HTMLExtractResponse struct {
    Total           int             `json:"total"` //The total number of values extracted
    Values          []string        `json:"values"` //Array of extracted values
}

/*
 * Structure for the custom type PhoneVerifyResponse
 */
type PhoneVerifyResponse struct {
    NumberValid     bool            `json:"numberValid"` //Is this a valid phone number
    Calling         bool            `json:"calling"` //True if the call is being made now
    SecurityCode    string          `json:"securityCode"` //The security code generated, you can save this code to perform your own verification or you can use the Verify Security Code API
}

/*
 * Structure for the custom type SMSVerifyResponse
 */
type SMSVerifyResponse struct {
    NumberValid     bool            `json:"numberValid"` //Is this a valid phone number
    SecurityCode    string          `json:"securityCode"` //The security code generated, you can save this code to perform your own verification or you can use the Verify Security Code API
    Sent            bool            `json:"sent"` //True if the SMS has been sent
}

/*
 * Structure for the custom type Blacklist
 */
type Blacklist struct {
    IsListed        bool            `json:"isListed"` //true if listed, false if not
    ListHost        string          `json:"listHost"` //the domain/hostname of the DNSBL
    ListRating      int             `json:"listRating"` //the list rating [1-3] with 1 being the best rating and 3 the lowest rating
    ListName        string          `json:"listName"` //the name of the DNSBL
    TxtRecord       string          `json:"txtRecord"` //the TXT record returned for this listing (if listed)
}

/*
 * Structure for the custom type IPInfoResponse
 */
type IPInfoResponse struct {
    Valid           bool            `json:"valid"` //Is this a valid IP address
    Country         string          `json:"country"` //Full country name
    Hostname        string          `json:"hostname"` //IP hostname (if reverse-lookup has been used)
    City            string          `json:"city"` //Full city name (if detectable)
    CountryCode     string          `json:"countryCode"` //ISO 2-letter country code
    Latitude        float32         `json:"latitude"` //Location latitude
    Region          string          `json:"region"` //Full region name (if detectable)
    Longitude       float32         `json:"longitude"` //Location longitude
}

/*
 * Structure for the custom type URLInfoResponse
 */
type URLInfoResponse struct {
    HttpStatusMessage string          `json:"httpStatusMessage"` //The HTTP status message assoicated with the status code
    ServerRegion      string          `json:"serverRegion"` //Server IP geo-location: full region name (if detectable)
    Query             map[string]interface{} `json:"query"` //A key:value map of the URL query paramaters
    ServerName        string          `json:"serverName"` //The name of the server software hosting this URL
    UrlPort           int             `json:"urlPort"` //The URL port
    ServerCountry     string          `json:"serverCountry"` //Server IP geo-location: full country name
    Real              bool            `json:"real"` //Is this URL actually serving real content
    ServerCity        string          `json:"serverCity"` //Server IP geo-location: full city name (if detectable)
    UrlPath           string          `json:"urlPath"` //The URL path
    Url               string          `json:"url"` //The fully qualified URL. This may be different to the URL requested if http-redirect is True
    Valid             bool            `json:"valid"` //Is this a valid well-formed URL
    ServerHostname    string          `json:"serverHostname"` //The server hostname (PTR)
    LoadTime          float32         `json:"loadTime"` //The time taken to load the URL content (in seconds)
    HttpOk            bool            `json:"httpOk"` //True if this URL responded with an HTTP OK (200) status
    ContentSize       int             `json:"contentSize"` //The size of the URL content in bytes
    HttpStatus        int             `json:"httpStatus"` //The HTTP status code this URL responded with
    ServerCountryCode string          `json:"serverCountryCode"` //Server IP geo-location: ISO 2-letter country code
    ContentEncoding   string          `json:"contentEncoding"` //The encoding type the URL uses
    ServerIp          string          `json:"serverIp"` //The IP address of the server hosting this URL
    UrlProtocol       string          `json:"urlProtocol"` //The URL protocol (usually http or https)
    ContentType       string          `json:"contentType"` //The content-type the URL points to
    HttpRedirect      bool            `json:"httpRedirect"` //True if this URL responded with a HTTP redirect
}

/*
 * Structure for the custom type GeocodeReverseResponse
 */
type GeocodeReverseResponse struct {
    Country         string          `json:"country"` //The country of the location
    Found           bool            `json:"found"` //True if these coordinates map to a real location
    Address         string          `json:"address"` //The fully formatted address
    City            string          `json:"city"` //The city of the location
    CountryCode     string          `json:"countryCode"` //The ISO 2-letter country code of the location
    PostalCode      string          `json:"postalCode"` //The postal code for the location
}

/*
 * Structure for the custom type UserAgentInfoResponse
 */
type UserAgentInfoResponse struct {
    MobileScreenWidth  int             `json:"mobileScreenWidth"` //Mobile device screen width (in px)
    MobileBrand        string          `json:"mobileBrand"` //Mobile device brand
    MobileModel        string          `json:"mobileModel"` //Mobile device model
    Producer           string          `json:"producer"` //Producer or manufacturer
    BrowserName        string          `json:"browserName"` //Browser software name
    MobileScreenHeight int             `json:"mobileScreenHeight"` //Mobile device screen height (in px)
    IsMobile           bool            `json:"isMobile"` //True if this is a mobile user-agent
    Type               string          `json:"type"` //The user-agent type, possible values are: desktop-browser, email-client, feed-reader, software-library, media-player, mobile-browser, robot, unknown
    Version            string          `json:"version"` //Software version
    OperatingSystem    string          `json:"operatingSystem"` //Operating system
    MobileBrowser      string          `json:"mobileBrowser"` //Mobile device browser
}

/*
 * Structure for the custom type EmailValidateResponse
 */
type EmailValidateResponse struct {
    Valid           bool            `json:"valid"` //TODO: Write general description for this field
    SyntaxError     bool            `json:"syntaxError"` //TODO: Write general description for this field
    Domain          string          `json:"domain"` //TODO: Write general description for this field
    DomainError     bool            `json:"domainError"` //TODO: Write general description for this field
    IsFreemail      bool            `json:"isFreemail"` //TODO: Write general description for this field
    Email           string          `json:"email"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type PhoneValidateResponse
 */
type PhoneValidateResponse struct {
    Valid                    bool            `json:"valid"` //Is this a valid phone number
    InternationalCallingCode string          `json:"internationalCallingCode"` //Numbers international calling code
    CountryCode              string          `json:"countryCode"` //Number location ISO 2-letter country code
    Location                 string          `json:"location"` //Number location (could be a city, region or country)
    IsMobile                 bool            `json:"isMobile"` //Is this a mobile number
    Type                     string          `json:"type"` //The number type, possible values are: mobile, fixed-line, premium-rate, toll-free, voip, unknown
    InternationalNumber      string          `json:"internationalNumber"` //Number represented in international format
    LocalNumber              string          `json:"localNumber"` //Number represented in local format
}

/*
 * Structure for the custom type Location
 */
type Location struct {
    Country         string          `json:"country"` //TODO: Write general description for this field
    Address         string          `json:"address"` //TODO: Write general description for this field
    City            string          `json:"city"` //TODO: Write general description for this field
    CountryCode     string          `json:"countryCode"` //TODO: Write general description for this field
    Latitude        float32         `json:"latitude"` //TODO: Write general description for this field
    PostalCode      string          `json:"postalCode"` //TODO: Write general description for this field
    Longitude       float32         `json:"longitude"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ConvertResponse
 */
type ConvertResponse struct {
    Valid           bool            `json:"valid"` //Was the coversion successful and produced a valid result
    Result          string          `json:"result"` //The result of the conversion
    FromValue       string          `json:"fromValue"` //The value being converted from
    ToType          string          `json:"toType"` //The type being converted to
    FromType        string          `json:"fromType"` //The type of the value being converted from
}
