/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package models_pkg



/*
 * Structure for the custom type HLRLookupResponse
 */
type HLRLookupResponse struct {
    NumberValid              bool            `json:"numberValid" form:"numberValid"` //True if this a valid phone number
    InternationalCallingCode int64           `json:"internationalCallingCode" form:"internationalCallingCode"` //The numbers international calling code
    Mnc                      string          `json:"mnc" form:"mnc"` //The mobile MNC number (Mobile Network Code)
    NumberType               string          `json:"numberType" form:"numberType"` //The number type, possible values are: <ul> <li>mobile</li> <li>fixed-line</li> <li>premium-rate</li> <li>toll-free</li> <li>voip</li> <li>unknown</li> </ul>
    HlrValid                 bool            `json:"hlrValid" form:"hlrValid"` //Was the HLR lookup successful. If true then this is a working and registered cell-phone or mobile device (SMS and phone calls will be delivered)
    HlrStatus                string          `json:"hlrStatus" form:"hlrStatus"` //The HLR lookup status, possible values are: <ul> <li>ok - the HLR lookup was successful and the device is connected</li> <li>absent - the number was once registered but the device has been switched off or out of network range for some time</li> <li>unknown - the number is not known by the mobile network</li> <li>invalid - the number is not a valid mobile MSISDN number</li> <li>fixed-line - the number is a registered fixed-line not mobile</li> <li>voip - the number has been detected as a VOIP line</li> <li>failed - the HLR lookup has failed, we could not determine the real status of this number</li> </ul>
    PortedNetwork            string          `json:"portedNetwork" form:"portedNetwork"` //If the number has been ported, the ported to carrier name
    Imsi                     string          `json:"imsi" form:"imsi"` //The mobile IMSI number (International Mobile Subscriber Identity)
    Mcc                      string          `json:"mcc" form:"mcc"` //The mobile MCC number (Mobile Country Code)
    InternationalNumber      string          `json:"internationalNumber" form:"internationalNumber"` //The number represented in full international format
    LocalNumber              string          `json:"localNumber" form:"localNumber"` //The number represented in local dialing format
    CountryCode              string          `json:"countryCode" form:"countryCode"` //The number location as an ISO 2-letter country code
    IsPorted                 bool            `json:"isPorted" form:"isPorted"` //Has this number been ported to another network
    Msin                     string          `json:"msin" form:"msin"` //The mobile MSIN number (Mobile Subscription Identification Number)
    Location                 string          `json:"location" form:"location"` //The number location. Could be a city, region or country depending on the type of number
    OriginNetwork            string          `json:"originNetwork" form:"originNetwork"` //The origin mobile carrier name
    IsMobile                 bool            `json:"isMobile" form:"isMobile"` //True if this is a mobile number (only true with 100% certainty, if the number type is unknown this value will be false)
    IsRoaming                bool            `json:"isRoaming" form:"isRoaming"` //Is this number currently roaming from its origin country
    Country                  string          `json:"country" form:"country"` //The phone number country
    CountryCode3             string          `json:"countryCode3" form:"countryCode3"` //The number location as an ISO 3-letter country code
    CurrencyCode             string          `json:"currencyCode" form:"currencyCode"` //ISO 4217 currency code associated with the country
    RoamingCountryCode       string          `json:"roamingCountryCode" form:"roamingCountryCode"` //If the number is currently roaming, the ISO 2-letter country code of the roaming in country
    Msc                      string          `json:"msc" form:"msc"` //The mobile MSC number (Mobile Switching Center)
}

/*
 * Structure for the custom type URLInfoResponse
 */
type URLInfoResponse struct {
    HttpStatusMessage int64           `json:"httpStatusMessage" form:"httpStatusMessage"` //The HTTP status message assoicated with the status code
    ServerRegion      string          `json:"serverRegion" form:"serverRegion"` //The servers IP geo-location: full region name (if detectable)
    Query             map[string]string `json:"query" form:"query"` //A key-value map of the URL query paramaters
    ServerName        string          `json:"serverName" form:"serverName"` //The name of the server software hosting this URL
    UrlPort           int64           `json:"urlPort" form:"urlPort"` //The URL port
    ServerCountry     string          `json:"serverCountry" form:"serverCountry"` //The servers IP geo-location: full country name
    Real              bool            `json:"real" form:"real"` //Is this URL actually serving real content
    ServerCity        string          `json:"serverCity" form:"serverCity"` //The servers IP geo-location: full city name (if detectable)
    UrlPath           string          `json:"urlPath" form:"urlPath"` //The URL path
    Url               string          `json:"url" form:"url"` //The fully qualified URL. This may be different to the URL requested if http-redirect is true
    Valid             bool            `json:"valid" form:"valid"` //Is this a valid well-formed URL
    ServerHostname    string          `json:"serverHostname" form:"serverHostname"` //The servers hostname (PTR record)
    LoadTime          int64           `json:"loadTime" form:"loadTime"` //The time taken to load the URL content in seconds
    HttpOk            bool            `json:"httpOk" form:"httpOk"` //True if this URL responded with an HTTP OK (200) status
    ContentSize       int64           `json:"contentSize" form:"contentSize"` //The size of the URL content in bytes
    HttpStatus        int64           `json:"httpStatus" form:"httpStatus"` //The HTTP status code this URL responded with. An HTTP status of 0 indicates a network level issue
    ServerCountryCode string          `json:"serverCountryCode" form:"serverCountryCode"` //The servers IP geo-location: ISO 2-letter country code
    ContentEncoding   string          `json:"contentEncoding" form:"contentEncoding"` //The encoding format the URL uses
    ServerIp          string          `json:"serverIp" form:"serverIp"` //The IP address of the server hosting this URL
    UrlProtocol       string          `json:"urlProtocol" form:"urlProtocol"` //The URL protocol, usually http or https
    ContentType       string          `json:"contentType" form:"contentType"` //The content-type this URL serves
    HttpRedirect      bool            `json:"httpRedirect" form:"httpRedirect"` //True if this URL responded with an HTTP redirect
    Content           string          `json:"content" form:"content"` //The actual content this URL responded with. Only set if the 'fetch-content' option was used
    IsTimeout         bool            `json:"isTimeout" form:"isTimeout"` //True if a timeout occurred while loading the URL. You can set the timeout with the request parameter 'timeout'
}

/*
 * Structure for the custom type IPProbeResponse
 */
type IPProbeResponse struct {
    Valid               bool            `json:"valid" form:"valid"` //Is this a valid IPv4 or IPv6 address
    Country             string          `json:"country" form:"country"` //Full country name
    ProviderType        string          `json:"providerType" form:"providerType"` //The detected provider type, possible values are: <ul> <li>isp - IP belongs to an internet service provider. This includes both mobile, home and business internet providers</li> <li>hosting - IP belongs to a hosting company. This includes website hosting, cloud computing platforms and colocation facilities</li> <li>vpn - IP belongs to a VPN provider</li> <li>proxy - IP belongs to a proxy service. This includes HTTP/SOCKS proxies and browser based proxies</li> <li>university - IP belongs to a university/college/campus</li> <li>government - IP belongs to a government department. This includes military facilities</li> <li>commercial - IP belongs to a commercial entity such as a corporate headquarters or company office</li> <li>unknown - could not identify the provider type</li> </ul>
    CountryCode         string          `json:"countryCode" form:"countryCode"` //ISO 2-letter country code
    Hostname            string          `json:"hostname" form:"hostname"` //The IPs full hostname (PTR)
    ProviderDomain      string          `json:"providerDomain" form:"providerDomain"` //The domain name of the provider
    City                string          `json:"city" form:"city"` //Full city name (if detectable)
    ProviderWebsite     string          `json:"providerWebsite" form:"providerWebsite"` //The website URL for the provider
    Ip                  string          `json:"ip" form:"ip"` //The IP address
    Region              string          `json:"region" form:"region"` //Full region name (if detectable)
    ProviderDescription string          `json:"providerDescription" form:"providerDescription"` //A description of the provider (usually extracted from the providers website)
    ContinentCode       string          `json:"continentCode" form:"continentCode"` //ISO 2-letter continent code
    IsHosting           bool            `json:"isHosting" form:"isHosting"` //True if this IP belongs to a hosting company. Note that this can still be true even if the provider type is VPN/proxy, this occurs in the case that the IP is detected as both types
    IsIsp               bool            `json:"isIsp" form:"isIsp"` //True if this IP belongs to an internet service provider. Note that this can still be true even if the provider type is VPN/proxy, this occurs in the case that the IP is detected as both types
    CountryCode3        string          `json:"countryCode3" form:"countryCode3"` //ISO 3-letter country code
    CurrencyCode        string          `json:"currencyCode" form:"currencyCode"` //ISO 4217 currency code associated with the country
    IsVpn               bool            `json:"isVpn" form:"isVpn"` //True if this IP ia a VPN
    IsProxy             bool            `json:"isProxy" form:"isProxy"` //True if this IP ia a proxy
    Asn                 string          `json:"asn" form:"asn"` //The autonomous system (AS) number
    AsCidr              string          `json:"asCidr" form:"asCidr"` //The autonomous system (AS) CIDR range
    AsCountryCode       string          `json:"asCountryCode" form:"asCountryCode"` //The autonomous system (AS) ISO 2-letter country code
    AsCountryCode3      string          `json:"asCountryCode3" form:"asCountryCode3"` //The autonomous system (AS) ISO 3-letter country code
    AsDomains           []string        `json:"asDomains" form:"asDomains"` //Array of all the domains associated with the autonomous system (AS)
    AsDescription       string          `json:"asDescription" form:"asDescription"` //The autonomous system (AS) description / company name
    AsAge               int64           `json:"asAge" form:"asAge"` //The age of the autonomous system (AS) in number of years since registration
    HostDomain          string          `json:"hostDomain" form:"hostDomain"` //The IPs host domain
    VpnDomain           string          `json:"vpnDomain" form:"vpnDomain"` //The domain of the VPN provider (may be empty if the VPN domain is not detectable)
}

/*
 * Structure for the custom type BrowserBotResponse
 */
type BrowserBotResponse struct {
    Url               string          `json:"url" form:"url"` //The page URL
    Content           string          `json:"content" form:"content"` //The complete raw, decompressed and decoded page content. Usually will be either HTML, JSON or XML
    MimeType          string          `json:"mimeType" form:"mimeType"` //The document MIME type
    Title             string          `json:"title" form:"title"` //The document title
    IsError           bool            `json:"isError" form:"isError"` //True if an error has occurred loading the page. Check the 'error-message' field for details
    IsTimeout         bool            `json:"isTimeout" form:"isTimeout"` //True if a timeout occurred while loading the page. You can set the timeout with the request parameter 'timeout'
    ErrorMessage      string          `json:"errorMessage" form:"errorMessage"` //Contains the error message if an error has occurred ('is-error' will be true)
    HttpStatusCode    int64           `json:"httpStatusCode" form:"httpStatusCode"` //The HTTP status code the URL returned
    HttpStatusMessage string          `json:"httpStatusMessage" form:"httpStatusMessage"` //The HTTP status message the URL returned
    IsHttpOk          bool            `json:"isHttpOk" form:"isHttpOk"` //True if the HTTP status is OK (200)
    IsHttpRedirect    bool            `json:"isHttpRedirect" form:"isHttpRedirect"` //True if the URL responded with an HTTP redirect
    HttpRedirectUrl   string          `json:"httpRedirectUrl" form:"httpRedirectUrl"` //The redirected URL if the URL responded with an HTTP redirect
    ServerIp          string          `json:"serverIp" form:"serverIp"` //The HTTP servers IP address
    LoadTime          int64           `json:"loadTime" form:"loadTime"` //The number of seconds taken to load the page (from initial request until DOM ready)
    ResponseHeaders   map[string]string `json:"responseHeaders" form:"responseHeaders"` //Map containing all the HTTP response headers the URL responded with
    IsSecure          bool            `json:"isSecure" form:"isSecure"` //True if the page is secured using TLS/SSL
    SecurityDetails   map[string]string `json:"securityDetails" form:"securityDetails"` //Map containing details of the TLS/SSL setup
    Elements          []string        `json:"elements" form:"elements"` //Array containing all the elements matching the supplied selector. Each element object will contain the text content, HTML content and all current element attributes
    ExecResults       []string        `json:"execResults" form:"execResults"` //If you executed any JavaScript this array holds the results as objects
}

/*
 * Structure for the custom type BINLookupResponse
 */
type BINLookupResponse struct {
    Country         string          `json:"country" form:"country"` //The full country name of the issuer
    IpCity          string          `json:"ipCity" form:"ipCity"` //The city of the customers IP (if detectable)
    IpMatchesBin    bool            `json:"ipMatchesBin" form:"ipMatchesBin"` //True if the customers IP country matches the BIN country
    CardType        string          `json:"cardType" form:"cardType"` //The card type, will always be one of: DEBIT, CREDIT, CHARGE CARD
    CardCategory    string          `json:"cardCategory" form:"cardCategory"` //The card category. There are many different card categories the most common card categories are: CLASSIC, BUSINESS, CORPORATE, PLATINUM, PREPAID
    IpCountryCode   string          `json:"ipCountryCode" form:"ipCountryCode"` //The ISO 2-letter country code of the customers IP
    IpCountry       string          `json:"ipCountry" form:"ipCountry"` //The country of the customers IP
    Issuer          string          `json:"issuer" form:"issuer"` //The card issuer
    IpBlocklisted   bool            `json:"ipBlocklisted" form:"ipBlocklisted"` //True if the customers IP is listed on one of our blocklists, see the <a href="http://www.neutrinoapi.com/api/ip-blocklist/">IP Blocklist API</a>
    Valid           bool            `json:"valid" form:"valid"` //Is this a valid BIN or IIN number
    IpBlocklists    []string        `json:"ipBlocklists" form:"ipBlocklists"` //An array of strings indicating which blocklists this IP is listed on
    IssuerWebsite   string          `json:"issuerWebsite" form:"issuerWebsite"` //The card issuers website
    CountryCode     string          `json:"countryCode" form:"countryCode"` //The ISO 2-letter country code of the issuer
    IpRegion        string          `json:"ipRegion" form:"ipRegion"` //The region of the customers IP (if detectable)
    CardBrand       string          `json:"cardBrand" form:"cardBrand"` //The card brand (e.g. Visa or Mastercard)
    IssuerPhone     string          `json:"issuerPhone" form:"issuerPhone"` //The card issuers phone number
    CountryCode3    string          `json:"countryCode3" form:"countryCode3"` //The ISO 3-letter country code of the issuer
    CurrencyCode    string          `json:"currencyCode" form:"currencyCode"` //ISO 4217 currency code associated with the country of the issuer
    IpCountryCode3  string          `json:"ipCountryCode3" form:"ipCountryCode3"` //The ISO 3-letter country code of the customers IP
}

/*
 * Structure for the custom type GeocodeReverseResponse
 */
type GeocodeReverseResponse struct {
    Country           string          `json:"country" form:"country"` //The country of the location
    Found             bool            `json:"found" form:"found"` //True if these coordinates map to a real location
    Address           string          `json:"address" form:"address"` //The fully formatted address
    City              string          `json:"city" form:"city"` //The city of the location
    CountryCode       string          `json:"countryCode" form:"countryCode"` //The ISO 2-letter country code of the location
    PostalCode        string          `json:"postalCode" form:"postalCode"` //The postal code for the location
    State             string          `json:"state" form:"state"` //The state of the location
    AddressComponents map[string]string `json:"addressComponents" form:"addressComponents"` //The components which make up the address such as road, city, state, etc
    CountryCode3      string          `json:"countryCode3" form:"countryCode3"` //The ISO 3-letter country code of the location
    CurrencyCode      string          `json:"currencyCode" form:"currencyCode"` //ISO 4217 currency code associated with the country
    LocationType      string          `json:"locationType" form:"locationType"` //The detected location type ordered roughly from most to least precise, possible values are: <ul> <li>address - indicates a precise street address</li> <li>street - accurate to the street level but may not point to the exact location of the house/building number</li> <li>city - accurate to the city level, this includes villages, towns, suburbs, etc</li> <li>postal-code - indicates a postal code area (no house or street information present)</li> <li>railway - location is part of a rail network such as a station or railway track</li> <li>natural - indicates a natural feature, for example a mountain peak or a waterway</li> <li>island - location is an island or archipelago</li> <li>administrative - indicates an administrative boundary such as a country, state or province</li> </ul>
    LocationTags      []string        `json:"locationTags" form:"locationTags"` //Array of strings containing any location tags associated with the address. Tags are additional pieces of metadata about a specific location, there are thousands of different tags. Some examples of tags: shop, office, cafe, bank, pub
    Latitude          int64           `json:"latitude" form:"latitude"` //The location latitude
    Longitude         int64           `json:"longitude" form:"longitude"` //The location longitude
    Timezone          map[string]string `json:"timezone" form:"timezone"` //Map containing timezone details for the location: <ul> <li>id - the time zone ID as per the IANA time zone database (tzdata)</li> <li>name - the time zone name</li> <li>abbr - the time zone abbreviation</li> <li>date - the current date within the time zone (ISO format)</li> <li>time - the current time within the time zone (ISO format)</li> </ul>
}

/*
 * Structure for the custom type Location
 */
type Location struct {
    Country           string          `json:"country" form:"country"` //The country of the location
    Address           string          `json:"address" form:"address"` //The fully formatted address
    City              string          `json:"city" form:"city"` //The city of the location
    CountryCode       string          `json:"countryCode" form:"countryCode"` //The ISO 2-letter country code of the location
    CountryCode3      string          `json:"countryCode3" form:"countryCode3"` //The ISO 3-letter country code of the location
    Latitude          float64         `json:"latitude" form:"latitude"` //The location latitude
    PostalCode        string          `json:"postalCode" form:"postalCode"` //The postal code for the location
    Longitude         float64         `json:"longitude" form:"longitude"` //The location longitude
    State             string          `json:"state" form:"state"` //The state of the location (if applicable)
    AddressComponents map[string]string `json:"addressComponents" form:"addressComponents"` //The components which make up the address such as road, city, state etc. May also include additional metadata about the address
}

/*
 * Structure for the custom type PhoneValidateResponse
 */
type PhoneValidateResponse struct {
    Valid                    bool            `json:"valid" form:"valid"` //Is this a valid phone number
    InternationalCallingCode int64           `json:"internationalCallingCode" form:"internationalCallingCode"` //The international calling code
    CountryCode              string          `json:"countryCode" form:"countryCode"` //The phone number country as an ISO 2-letter country code
    Location                 string          `json:"location" form:"location"` //The phone number location. Could be a city, region or country depending on the type of number
    IsMobile                 bool            `json:"isMobile" form:"isMobile"` //True if this is a mobile number. If the number type is unknown this value will be false, use HLR lookup instead
    Type                     string          `json:"type" form:"type"` //The predicted number type. Note: type detection is not possible in some countries which have no predictable prefix pattern (you can use the HLR Lookup API in these cases) Possible values are: <ul> <li>mobile</li> <li>fixed-line</li> <li>premium-rate</li> <li>toll-free</li> <li>voip</li> <li>unknown (use HLR lookup)</li> </ul>
    InternationalNumber      string          `json:"internationalNumber" form:"internationalNumber"` //The number represented in full international format (E.164)
    LocalNumber              string          `json:"localNumber" form:"localNumber"` //The number represented in local dialing format
    Country                  string          `json:"country" form:"country"` //The phone number country
    CountryCode3             string          `json:"countryCode3" form:"countryCode3"` //The phone number country as an ISO 3-letter country code
    CurrencyCode             string          `json:"currencyCode" form:"currencyCode"` //ISO 4217 currency code associated with the country
}

/*
 * Structure for the custom type IPInfoResponse
 */
type IPInfoResponse struct {
    Valid           bool            `json:"valid" form:"valid"` //Is this a valid IPv4 or IPv6 address
    Country         string          `json:"country" form:"country"` //Full country name
    Hostname        string          `json:"hostname" form:"hostname"` //The IPs full hostname (only set if reverse-lookup has been used)
    City            string          `json:"city" form:"city"` //Name of the city (if detectable)
    CountryCode     string          `json:"countryCode" form:"countryCode"` //ISO 2-letter country code
    Latitude        int64           `json:"latitude" form:"latitude"` //Location latitude
    Region          string          `json:"region" form:"region"` //Name of the region (if detectable)
    Longitude       int64           `json:"longitude" form:"longitude"` //Location longitude
    ContinentCode   string          `json:"continentCode" form:"continentCode"` //ISO 2-letter continent code
    Ip              string          `json:"ip" form:"ip"` //The IP address
    CountryCode3    string          `json:"countryCode3" form:"countryCode3"` //ISO 3-letter country code
    CurrencyCode    string          `json:"currencyCode" form:"currencyCode"` //ISO 4217 currency code associated with the country
    HostDomain      string          `json:"hostDomain" form:"hostDomain"` //The IPs host domain (only set if reverse-lookup has been used)
    Timezone        map[string]string `json:"timezone" form:"timezone"` //Map containing timezone details for the location: <ul> <li>id - the time zone ID as per the IANA time zone database (tzdata)</li> <li>name - the time zone name</li> <li>abbr - the time zone abbreviation</li> <li>date - the current date within the time zone (ISO format)</li> <li>time - the current time within the time zone (ISO format)</li> </ul>
}

/*
 * Structure for the custom type HostReputationResponse
 */
type HostReputationResponse struct {
    IsListed        bool            `json:"isListed" form:"isListed"` //Is this host blacklisted
    Lists           []*Blacklist    `json:"lists" form:"lists"` //An array of objects for each DNSBL checked, with the following keys: <ul> <li>is-listed - true if listed, false if not</li> <li>list-name - the name of the DNSBL</li> <li>list-host - the domain/hostname of the DNSBL</li> <li>list-rating - the list rating [1-3] with 1 being the best rating and 3 the lowest rating</li> <li>txt-record - the TXT record returned for this listing (if listed)</li> <li>return-code - the specific return code for this listing (if listed)</li> <li>response-time - the DNSBL server response time in milliseconds</li> </ul>
    ListCount       int64           `json:"listCount" form:"listCount"` //The number of DNSBLs the host is listed on
    Host            string          `json:"host" form:"host"` //The IP address or host name
}

/*
 * Structure for the custom type IPBlocklistResponse
 */
type IPBlocklistResponse struct {
    Ip              string          `json:"ip" form:"ip"` //The IP address
    IsBot           bool            `json:"isBot" form:"isBot"` //IP is hosting a malicious bot or is part of a botnet. Includes brute-force crackers
    IsExploitBot    bool            `json:"isExploitBot" form:"isExploitBot"` //IP is hosting an exploit finding bot or is running exploit scanning software
    IsMalware       bool            `json:"isMalware" form:"isMalware"` //IP is involved in distributing or is running malware
    IsSpider        bool            `json:"isSpider" form:"isSpider"` //IP is running a hostile web spider / web crawler
    IsDshield       bool            `json:"isDshield" form:"isDshield"` //IP has been flagged as an attack source on DShield (dshield.org)
    ListCount       int64           `json:"listCount" form:"listCount"` //The number of blocklists the IP is listed on
    IsProxy         bool            `json:"isProxy" form:"isProxy"` //IP has been detected as an anonymous web proxy or anonymous HTTP proxy
    IsHijacked      bool            `json:"isHijacked" form:"isHijacked"` //IP is part of a hijacked netblock or a netblock controlled by a criminal organization
    IsTor           bool            `json:"isTor" form:"isTor"` //IP is a Tor node or running a Tor related service
    IsSpyware       bool            `json:"isSpyware" form:"isSpyware"` //IP is involved in distributing or is running spyware
    IsSpamBot       bool            `json:"isSpamBot" form:"isSpamBot"` //IP address is hosting a spam bot, comment spamming or any other spamming type software
    IsListed        bool            `json:"isListed" form:"isListed"` //Is this IP on a blocklist
    IsVpn           bool            `json:"isVpn" form:"isVpn"` //IP belongs to a VPN provider. This field is only kept for backward compatibility, for VPN detection use the <a href="https://www.neutrinoapi.com/api/ip-probe/">IP Probe</a> API
    LastSeen        int64           `json:"lastSeen" form:"lastSeen"` //The last time this IP was seen on a blocklist (in Unix time or 0 if not listed recently)
    Blocklists      []string        `json:"blocklists" form:"blocklists"` //An array of strings indicating which blocklists this IP is listed on (empty if not listed)
    Sensors         []string        `json:"sensors" form:"sensors"` //An array of objects containing details on which sensors were used to detect this IP
}

/*
 * Structure for the custom type Blacklist
 */
type Blacklist struct {
    IsListed        bool            `json:"isListed" form:"isListed"` //true if listed, false if not
    ListHost        string          `json:"listHost" form:"listHost"` //the domain/hostname of the DNSBL
    ListRating      int64           `json:"listRating" form:"listRating"` //the list rating [1-3] with 1 being the best rating and 3 the lowest rating
    ListName        string          `json:"listName" form:"listName"` //the name of the DNSBL
    TxtRecord       string          `json:"txtRecord" form:"txtRecord"` //the TXT record returned for this listing (if listed)
    ReturnCode      string          `json:"returnCode" form:"returnCode"` //the specific return code for this listing (if listed)
    ResponseTime    int64           `json:"responseTime" form:"responseTime"` //the DNSBL server response time in milliseconds
}

/*
 * Structure for the custom type ConvertResponse
 */
type ConvertResponse struct {
    Valid           bool            `json:"valid" form:"valid"` //True if the coversion was successful and produced a valid result
    Result          string          `json:"result" form:"result"` //The result of the conversion in string format
    FromValue       string          `json:"fromValue" form:"fromValue"` //The value being converted from
    ToType          string          `json:"toType" form:"toType"` //The type being converted to
    FromType        string          `json:"fromType" form:"fromType"` //The type of the value being converted from
    ResultFloat     int64           `json:"resultFloat" form:"resultFloat"` //The result of the conversion as a floating-point number
}

/*
 * Structure for the custom type EmailVerifyResponse
 */
type EmailVerifyResponse struct {
    Valid           bool            `json:"valid" form:"valid"` //Is this a valid email address (syntax and domain is valid)
    Verified        bool            `json:"verified" form:"verified"` //True if this address has passed SMTP verification. Check the smtp-status and smtp-response fields for specific verification details
    Email           string          `json:"email" form:"email"` //The email address. If you have used the fix-typos option then this will be the fixed address
    TyposFixed      bool            `json:"typosFixed" form:"typosFixed"` //True if typos have been fixed
    SyntaxError     bool            `json:"syntaxError" form:"syntaxError"` //True if this address has a syntax error
    DomainError     bool            `json:"domainError" form:"domainError"` //True if this address has a domain error (e.g. no valid mail server records)
    Domain          string          `json:"domain" form:"domain"` //The email domain
    Provider        string          `json:"provider" form:"provider"` //The email service provider domain
    IsFreemail      bool            `json:"isFreemail" form:"isFreemail"` //True if this address is a free-mail address
    IsDisposable    bool            `json:"isDisposable" form:"isDisposable"` //True if this address is a disposable, temporary or darknet related email address
    IsPersonal      bool            `json:"isPersonal" form:"isPersonal"` //True if this address is for a person. False if this is a role based address, e.g. admin@, help@, office@, etc.
    SmtpStatus      string          `json:"smtpStatus" form:"smtpStatus"` //The SMTP verification status for the address: <ul> <li>ok - SMTP verification was successful, this is a real address that can receive mail</li> <li>invalid - this is not a valid email address (has either a domain or syntax error)</li> <li>absent - this address is not registered with the email service provider</li> <li>unresponsive - the mail server(s) for this address timed-out or refused to open an SMTP connection</li> <li>unknown - sorry, we could not reliably determine the real status of this address (this address may or may not exist)</li> </ul>
    SmtpResponse    string          `json:"smtpResponse" form:"smtpResponse"` //The raw SMTP response message received during verification
    IsCatchAll      bool            `json:"isCatchAll" form:"isCatchAll"` //True if this email domain has a catch-all policy (it will accept mail for any username)
    IsDeferred      bool            `json:"isDeferred" form:"isDeferred"` //True if the mail server responded with a temporary failure (either a 4xx response code or unresponsive server). You can retry this address later, we recommend waiting at least 15 minutes before retrying
}

/*
 * Structure for the custom type EmailValidateResponse
 */
type EmailValidateResponse struct {
    Valid           bool            `json:"valid" form:"valid"` //Is this a valid email
    SyntaxError     bool            `json:"syntaxError" form:"syntaxError"` //True if this address has a syntax error
    Domain          string          `json:"domain" form:"domain"` //The email domain
    DomainError     bool            `json:"domainError" form:"domainError"` //True if this address has a domain error (e.g. no valid mail server records)
    IsFreemail      bool            `json:"isFreemail" form:"isFreemail"` //True if this address is a free-mail address
    Email           string          `json:"email" form:"email"` //The email address. If you have used the fix-typos option then this will be the fixed address
    IsDisposable    bool            `json:"isDisposable" form:"isDisposable"` //True if this address is a disposable, temporary or darknet related email address
    TyposFixed      bool            `json:"typosFixed" form:"typosFixed"` //True if typos have been fixed
    IsPersonal      bool            `json:"isPersonal" form:"isPersonal"` //True if this address belongs to a person. False if this is a role based address, e.g. admin@, help@, office@, etc.
    Provider        string          `json:"provider" form:"provider"` //The email service provider domain
}

/*
 * Structure for the custom type UserAgentInfoResponse
 */
type UserAgentInfoResponse struct {
    MobileScreenWidth      int64           `json:"mobileScreenWidth" form:"mobileScreenWidth"` //The estimated mobile device screen width in CSS 'px'
    MobileBrand            string          `json:"mobileBrand" form:"mobileBrand"` //The mobile device brand
    MobileModel            string          `json:"mobileModel" form:"mobileModel"` //The mobile device model
    Producer               string          `json:"producer" form:"producer"` //The producer or manufacturer of the user agent
    BrowserName            string          `json:"browserName" form:"browserName"` //The browser software name
    MobileScreenHeight     int64           `json:"mobileScreenHeight" form:"mobileScreenHeight"` //The estimated mobile device screen height in CSS 'px'
    IsMobile               bool            `json:"isMobile" form:"isMobile"` //True if this is a mobile user agent
    Type                   string          `json:"type" form:"type"` //The user agent type, possible values are: <ul> <li>desktop-browser</li> <li>mobile-browser</li> <li>email-client</li> <li>feed-reader</li> <li>software-library</li> <li>media-player (includes smart TVs)</li> <li>robot</li> <li>unknown</li> </ul>
    Version                string          `json:"version" form:"version"` //The browser software version
    OperatingSystem        string          `json:"operatingSystem" form:"operatingSystem"` //The full operating system name which may include the major version number or code name
    MobileBrowser          string          `json:"mobileBrowser" form:"mobileBrowser"` //The mobile device browser name (this is usually the same as the browser name)
    IsAndroid              bool            `json:"isAndroid" form:"isAndroid"` //True if this is an Android based mobile user agent
    IsIos                  bool            `json:"isIos" form:"isIos"` //True if this is an iOS based mobile user agent
    OperatingSystemFamily  string          `json:"operatingSystemFamily" form:"operatingSystemFamily"` //The operating system family name, this is the OS name without any version information
    OperatingSystemVersion string          `json:"operatingSystemVersion" form:"operatingSystemVersion"` //The operating system version number (if detectable)
    Engine                 string          `json:"engine" form:"engine"` //The browser engine name
    EngineVersion          string          `json:"engineVersion" form:"engineVersion"` //The browser engine version (if detectable)
}

/*
 * Structure for the custom type GeocodeAddressResponse
 */
type GeocodeAddressResponse struct {
    Found           int64           `json:"found" form:"found"` //The number of possible matching locations found
    Locations       []*Location     `json:"locations" form:"locations"` //Array of matching location objects: <table> <tbody> <tr> <th>Key</th> <th>Type</th> <th>Description</th> </tr> <tr> <td>latitude</td> <td>float</td> <td>The location latitude</td> </tr> <tr> <td>longitude</td> <td>float</td> <td>The location longitude</td> </tr> <tr> <td>address</td> <td>string</td> <td>The fully formatted address</td> </tr> <tr> <td>address-components</td> <td>map</td> <td>The components which make up the address such as road, city, state, etc</td> </tr> <tr> <td>city</td> <td>string</td> <td>The city of the location</td> </tr> <tr> <td>state</td> <td>string</td> <td>The state of the location</td> </tr> <tr> <td>country</td> <td>string</td> <td>The country of the location</td> </tr> <tr> <td>country-code</td> <td>string</td> <td>The ISO 2-letter country code of the location</td> </tr> <tr> <td>country-code3</td> <td>string</td> <td>The ISO 3-letter country code of the location</td> </tr> <tr> <td>currency-code</td> <td>string</td> <td>ISO 4217 currency code associated with the country</td> </tr> <tr> <td>postal-code</td> <td>string</td> <td>The postal code for the location</td> </tr> <tr> <td>location-type</td> <td>string</td> <td>The detected location type ordered roughly from most to least precise, possible values are: <ul> <li>address - indicates a precise street address</li> <li>street - accurate to the street level but may not point to the exact location of the house/building number</li> <li>city - accurate to the city level, this includes villages, towns, suburbs, etc</li> <li>postal-code - indicates a postal code area (no house or street information present)</li> <li>railway - location is part of a rail network such as a station or railway track</li> <li>natural - indicates a natural feature, for example a mountain peak or a waterway</li> <li>island - location is an island or archipelago</li> <li>administrative - indicates an administrative boundary such as a country, state or province</li> </ul></td> </tr> <tr> <td>location-tags</td> <td>array</td> <td>Array of strings containing any location tags associated with the address. Tags are additional pieces of metadata about a specific location, there are thousands of different tags. Some examples of tags: shop, office, cafe, bank, pub</td> </tr> <tr> <td>timezone</td> <td>map</td> <td>Map containing timezone details for the location: <ul> <li>id - the time zone ID as per the IANA time zone database (tzdata)</li> <li>name - the time zone name</li> <li>abbr - the time zone abbreviation</li> <li>date - the current date within the time zone (ISO format)</li> <li>time - the current time within the time zone (ISO format)</li> </ul></td> </tr> </tbody> </table>
}

/*
 * Structure for the custom type PhoneVerifyResponse
 */
type PhoneVerifyResponse struct {
    NumberValid     bool            `json:"numberValid" form:"numberValid"` //True if this a valid phone number
    Calling         bool            `json:"calling" form:"calling"` //True if the call is being made now
    SecurityCode    string          `json:"securityCode" form:"securityCode"` //The security code generated, you can save this code to perform your own verification or you can use the <a href="https://www.neutrinoapi.com/api/verify-security-code/">Verify Security Code API</a>
}

/*
 * Structure for the custom type SMSVerifyResponse
 */
type SMSVerifyResponse struct {
    NumberValid     bool            `json:"numberValid" form:"numberValid"` //True if this a valid phone number
    SecurityCode    string          `json:"securityCode" form:"securityCode"` //The security code generated, you can save this code to perform your own verification or you can use the <a href="https://www.neutrinoapi.com/api/verify-security-code/">Verify Security Code API</a>
    Sent            bool            `json:"sent" form:"sent"` //True if the SMS has been sent
}

/*
 * Structure for the custom type BadWordFilterResponse
 */
type BadWordFilterResponse struct {
    BadWordsList    []string        `json:"badWordsList" form:"badWordsList"` //An array of the bad words found
    BadWordsTotal   int64           `json:"badWordsTotal" form:"badWordsTotal"` //Total number of bad words detected
    CensoredContent string          `json:"censoredContent" form:"censoredContent"` //The censored content (only set if censor-character has been set)
    IsBad           bool            `json:"isBad" form:"isBad"` //Does the text contain bad words
}

/*
 * Structure for the custom type VerifySecurityCodeResponse
 */
type VerifySecurityCodeResponse struct {
    Verified        bool            `json:"verified" form:"verified"` //True if the code is valid
}

/*
 * Structure for the custom type SMSMessageResponse
 */
type SMSMessageResponse struct {
    NumberValid     bool            `json:"numberValid" form:"numberValid"` //True if this a valid phone number
    Sent            bool            `json:"sent" form:"sent"` //True if the SMS has been sent
}

/*
 * Structure for the custom type PhonePlaybackResponse
 */
type PhonePlaybackResponse struct {
    Calling         bool            `json:"calling" form:"calling"` //True if the call is being made now
    NumberValid     bool            `json:"numberValid" form:"numberValid"` //True if this a valid phone number
}
