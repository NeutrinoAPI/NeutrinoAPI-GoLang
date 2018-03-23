/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io )
 */

package models_pkg



/*
 * Structure for the custom type BINLookupResponse
 */
type BINLookupResponse struct {
    Country         string          `json:"country" form:"country"` //Full country name of the issuer
    IpCity          string          `json:"ipCity" form:"ipCity"` //The city name (if detectable) from the customer IP
    IpMatchesBin    bool            `json:"ipMatchesBin" form:"ipMatchesBin"` //True if the customer IP address country matches the BIN country
    CardType        string          `json:"cardType" form:"cardType"` //The card type, will always be one of: DEBIT, CREDIT, CHARGE CARD
    CardCategory    string          `json:"cardCategory" form:"cardCategory"` //The card category (if known)
    IpCountryCode   string          `json:"ipCountryCode" form:"ipCountryCode"` //The ISO 2-letter country code detected from the customer IP
    IpCountry       string          `json:"ipCountry" form:"ipCountry"` //The country detected from the customer IP
    Issuer          string          `json:"issuer" form:"issuer"` //The card issuer (if known)
    IpBlocklisted   bool            `json:"ipBlocklisted" form:"ipBlocklisted"` //True if the customer IP is listed on one of our blocklists, see the IP Blocklist API for more details
    Valid           bool            `json:"valid" form:"valid"` //Is this a valid BIN or IIN number
    IpBlocklists    []string        `json:"ipBlocklists" form:"ipBlocklists"` //An array of strings indicating which blocklists this IP is listed on
    IssuerWebsite   string          `json:"issuerWebsite" form:"issuerWebsite"` //The card issuer website (if known)
    CountryCode     string          `json:"countryCode" form:"countryCode"` //ISO 2-letter country code of the issuer
    IpRegion        string          `json:"ipRegion" form:"ipRegion"` //The region name (if detectable) from the customer IP
    CardBrand       string          `json:"cardBrand" form:"cardBrand"` //The card brand (e.g. Visa or Mastercard)
    IssuerPhone     string          `json:"issuerPhone" form:"issuerPhone"` //The card issuer phone number (if known)
}

/*
 * Structure for the custom type VerifySecurityCodeResponse
 */
type VerifySecurityCodeResponse struct {
    Verified        bool            `json:"verified" form:"verified"` //True if the code is valid
}

/*
 * Structure for the custom type PhoneVerifyResponse
 */
type PhoneVerifyResponse struct {
    NumberValid     bool            `json:"numberValid" form:"numberValid"` //Is this a valid phone number
    Calling         bool            `json:"calling" form:"calling"` //True if the call is being made now
    SecurityCode    string          `json:"securityCode" form:"securityCode"` //The security code generated, you can save this code to perform your own verification or you can use the Verify Security Code API
}

/*
 * Structure for the custom type HostReputationResponse
 */
type HostReputationResponse struct {
    IsListed        bool            `json:"isListed" form:"isListed"` //Is this host blacklisted
    Lists           []*Blacklist    `json:"lists" form:"lists"` //An array of objects for each DNSBL checked
    ListCount       int64           `json:"listCount" form:"listCount"` //The number of DNSBL's the host is listed on
}

/*
 * Structure for the custom type SMSVerifyResponse
 */
type SMSVerifyResponse struct {
    NumberValid     bool            `json:"numberValid" form:"numberValid"` //Is this a valid phone number
    SecurityCode    string          `json:"securityCode" form:"securityCode"` //The security code generated, you can save this code to perform your own verification or you can use the Verify Security Code API
    Sent            bool            `json:"sent" form:"sent"` //True if the SMS has been sent
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
    ResponseTime    int64           `json:"responseTime" form:"responseTime"` //the DNSBL server response time in milliseconds
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
    AddressComponents map[string]string `json:"addressComponents" form:"addressComponents"` //The components which make up the address such as road, city, state etc. May also include additional metadata about the address
}

/*
 * Structure for the custom type PhonePlaybackResponse
 */
type PhonePlaybackResponse struct {
    Calling         bool            `json:"calling" form:"calling"` //True if the call is being made now
    NumberValid     bool            `json:"number-valid" form:"number-valid"` //Is this a valid phone number
}

/*
 * Structure for the custom type BadWordFilterResponse
 */
type BadWordFilterResponse struct {
    BadWordsList    []string        `json:"badWordsList" form:"badWordsList"` //Array of the bad words found
    BadWordsTotal   int64           `json:"badWordsTotal" form:"badWordsTotal"` //Total number of bad words detected
    CensoredContent string          `json:"censoredContent" form:"censoredContent"` //The censored content (only set if censor-character has been set)
    IsBad           bool            `json:"isBad" form:"isBad"` //Does the text contain bad words
}

/*
 * Structure for the custom type GeocodeAddressResponse
 */
type GeocodeAddressResponse struct {
    Found           int64           `json:"found" form:"found"` //The number of possible matching locations found
    Locations       []*Location     `json:"locations" form:"locations"` //Array of matching location objects
}

/*
 * Structure for the custom type ConvertResponse
 */
type ConvertResponse struct {
    Valid           bool            `json:"valid" form:"valid"` //Was the coversion successful and produced a valid result
    Result          string          `json:"result" form:"result"` //The result of the conversion
    FromValue       string          `json:"fromValue" form:"fromValue"` //The value being converted from
    ToType          string          `json:"toType" form:"toType"` //The type being converted to
    FromType        string          `json:"fromType" form:"fromType"` //The type of the value being converted from
}

/*
 * Structure for the custom type Location
 */
type Location struct {
    Country           string          `json:"country" form:"country"` //The country of the location
    Address           string          `json:"address" form:"address"` //The fully formatted address
    City              string          `json:"city" form:"city"` //The city of the location
    CountryCode       string          `json:"countryCode" form:"countryCode"` //The ISO 2-letter country code of the location
    Latitude          float64         `json:"latitude" form:"latitude"` //The location latitude
    PostalCode        string          `json:"postalCode" form:"postalCode"` //The postal code for the location
    Longitude         float64         `json:"longitude" form:"longitude"` //The location longitude
    State             string          `json:"state" form:"state"` //The state of the location (if applicable)
    AddressComponents map[string]string `json:"addressComponents" form:"addressComponents"` //The components which make up the address such as road, city, state etc. May also include additional metadata about the address
}

/*
 * Structure for the custom type HLRLookupResponse
 */
type HLRLookupResponse struct {
    NumberValid              bool            `json:"numberValid" form:"numberValid"` //Is this a valid phone number (mobile or otherwise)
    InternationalCallingCode string          `json:"internationalCallingCode" form:"internationalCallingCode"` //Numbers international calling code
    Mnc                      string          `json:"mnc" form:"mnc"` //The mobile MNC number (only set if HLR lookup valid)
    NumberType               string          `json:"numberType" form:"numberType"` //The number type, possible values are: mobile, fixed-line, premium-rate, toll-free, voip, unknown
    HlrValid                 bool            `json:"hlrValid" form:"hlrValid"` //Was the HLR lookup successful. If true then this is a working and registered cell-phone or mobile device (SMS and phone calls will be delivered)
    HlrStatus                string          `json:"hlrStatus" form:"hlrStatus"` //The HLR lookup status. See API docs for specific status details
    PortedNetwork            string          `json:"portedNetwork" form:"portedNetwork"` //If the number has been ported, the ported to mobile carrier name (only set if HLR lookup valid)
    Imsi                     string          `json:"imsi" form:"imsi"` //The mobile IMSI number (only set if HLR lookup valid)
    Mcc                      string          `json:"mcc" form:"mcc"` //The mobile MCC number (only set if HLR lookup valid)
    InternationalNumber      string          `json:"internationalNumber" form:"internationalNumber"` //Number represented in international format
    LocalNumber              string          `json:"localNumber" form:"localNumber"` //Number represented in local format
    CountryCode              string          `json:"countryCode" form:"countryCode"` //Number location ISO 2-letter country code
    IsPorted                 bool            `json:"isPorted" form:"isPorted"` //Has this number been ported to another network
    Msin                     string          `json:"msin" form:"msin"` //The mobile MSIN number (only set if HLR lookup valid)
    Location                 string          `json:"location" form:"location"` //Number location (could be a city, region or country)
    OriginNetwork            string          `json:"originNetwork" form:"originNetwork"` //The origin mobile carrier name (only set if HLR lookup valid)
    IsMobile                 bool            `json:"isMobile" form:"isMobile"` //Is this a mobile number
    IsRoaming                bool            `json:"isRoaming" form:"isRoaming"` //Is this number currently roaming from its origin country
    Country                  string          `json:"country" form:"country"` //The phone number country
}

/*
 * Structure for the custom type HTMLExtractResponse
 */
type HTMLExtractResponse struct {
    Total           int64           `json:"total" form:"total"` //The total number of values extracted
    Values          []string        `json:"values" form:"values"` //Array of extracted values
}

/*
 * Structure for the custom type URLInfoResponse
 */
type URLInfoResponse struct {
    HttpStatusMessage string          `json:"httpStatusMessage" form:"httpStatusMessage"` //The HTTP status message assoicated with the status code
    ServerRegion      string          `json:"serverRegion" form:"serverRegion"` //Server IP geo-location: full region name (if detectable)
    Query             map[string]interface{} `json:"query" form:"query"` //A key:value map of the URL query paramaters
    ServerName        string          `json:"serverName" form:"serverName"` //The name of the server software hosting this URL
    UrlPort           int64           `json:"urlPort" form:"urlPort"` //The URL port
    ServerCountry     string          `json:"serverCountry" form:"serverCountry"` //Server IP geo-location: full country name
    Real              bool            `json:"real" form:"real"` //Is this URL actually serving real content
    ServerCity        string          `json:"serverCity" form:"serverCity"` //Server IP geo-location: full city name (if detectable)
    UrlPath           string          `json:"urlPath" form:"urlPath"` //The URL path
    Url               string          `json:"url" form:"url"` //The fully qualified URL. This may be different to the URL requested if http-redirect is True
    Valid             bool            `json:"valid" form:"valid"` //Is this a valid well-formed URL
    ServerHostname    string          `json:"serverHostname" form:"serverHostname"` //The server hostname (PTR)
    LoadTime          float64         `json:"loadTime" form:"loadTime"` //The time taken to load the URL content (in seconds)
    HttpOk            bool            `json:"httpOk" form:"httpOk"` //True if this URL responded with an HTTP OK (200) status
    ContentSize       int64           `json:"contentSize" form:"contentSize"` //The size of the URL content in bytes
    HttpStatus        int64           `json:"httpStatus" form:"httpStatus"` //The HTTP status code this URL responded with
    ServerCountryCode string          `json:"serverCountryCode" form:"serverCountryCode"` //Server IP geo-location: ISO 2-letter country code
    ContentEncoding   string          `json:"contentEncoding" form:"contentEncoding"` //The encoding type the URL uses
    ServerIp          string          `json:"serverIp" form:"serverIp"` //The IP address of the server hosting this URL
    UrlProtocol       string          `json:"urlProtocol" form:"urlProtocol"` //The URL protocol (usually http or https)
    ContentType       string          `json:"contentType" form:"contentType"` //The content-type the URL points to
    HttpRedirect      bool            `json:"httpRedirect" form:"httpRedirect"` //True if this URL responded with a HTTP redirect
    Content           string          `json:"content" form:"content"` //The actual content this URL responded with. Only set if the 'fetch-content' option was used
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
    Email           string          `json:"email" form:"email"` //The full email address (this could be different to the supplied address if fix-typos is used)
    IsDisposable    bool            `json:"isDisposable" form:"isDisposable"` //True if this address is a disposable, temporary or darknet related email address
    TyposFixed      bool            `json:"typosFixed" form:"typosFixed"` //True if typos have been fixed
    IsPersonal      bool            `json:"isPersonal" form:"isPersonal"` //True if this address belongs to a person. False if this is a role based address, e.g. admin@, help@, office@, etc.
    Provider        string          `json:"provider" form:"provider"` //The email service provider domain
}

/*
 * Structure for the custom type PhoneValidateResponse
 */
type PhoneValidateResponse struct {
    Valid                    bool            `json:"valid" form:"valid"` //Is this a valid phone number
    InternationalCallingCode string          `json:"internationalCallingCode" form:"internationalCallingCode"` //Numbers international calling code
    CountryCode              string          `json:"countryCode" form:"countryCode"` //Number location ISO 2-letter country code
    Location                 string          `json:"location" form:"location"` //Number location (could be a city, region or country)
    IsMobile                 bool            `json:"isMobile" form:"isMobile"` //Is this a mobile number
    Type                     string          `json:"type" form:"type"` //The number type, possible values are: mobile, fixed-line, premium-rate, toll-free, voip, unknown
    InternationalNumber      string          `json:"internationalNumber" form:"internationalNumber"` //Number represented in international format
    LocalNumber              string          `json:"localNumber" form:"localNumber"` //Number represented in local format
    Country                  string          `json:"country" form:"country"` //The phone number country
}

/*
 * Structure for the custom type IPInfoResponse
 */
type IPInfoResponse struct {
    Valid           bool            `json:"valid" form:"valid"` //Is this a valid IP address
    Country         string          `json:"country" form:"country"` //Full country name
    Hostname        string          `json:"hostname" form:"hostname"` //The IPs hostname (only set if reverse-lookup has been used)
    City            string          `json:"city" form:"city"` //Full city name (if detectable)
    CountryCode     string          `json:"countryCode" form:"countryCode"` //ISO 2-letter country code
    Latitude        float64         `json:"latitude" form:"latitude"` //Location latitude
    Region          string          `json:"region" form:"region"` //Full region name (if detectable)
    Longitude       float64         `json:"longitude" form:"longitude"` //Location longitude
    ContinentCode   string          `json:"continentCode" form:"continentCode"` //ISO 2-letter continent code
}

/*
 * Structure for the custom type IPBlocklistResponse
 */
type IPBlocklistResponse struct {
    Ip              string          `json:"ip" form:"ip"` //The IP address
    IsBot           bool            `json:"isBot" form:"isBot"` //IP is hosting a malicious bot or is part of a botnet
    IsExploitBot    bool            `json:"isExploitBot" form:"isExploitBot"` //IP is hosting an exploit finding bot or exploit scanning software
    IsMalware       bool            `json:"isMalware" form:"isMalware"` //IP is involved in distributing malware
    IsSpider        bool            `json:"isSpider" form:"isSpider"` //IP is a hostile spider or crawler
    IsDshield       bool            `json:"isDshield" form:"isDshield"` //IP has been flagged on DShield (dshield.org)
    ListCount       int64           `json:"listCount" form:"listCount"` //The number of blocklists the IP is listed on
    IsProxy         bool            `json:"isProxy" form:"isProxy"` //IP has been detected as an anonymous web proxy or anonymous HTTP proxy
    IsHijacked      bool            `json:"isHijacked" form:"isHijacked"` //hijacked netblocks or netblocks controlled by criminal organizations
    IsTor           bool            `json:"isTor" form:"isTor"` //IP is coming from a Tor node
    IsSpyware       bool            `json:"isSpyware" form:"isSpyware"` //IP is being used by spyware, malware, botnets or for other malicious activities
    IsSpamBot       bool            `json:"isSpamBot" form:"isSpamBot"` //IP address is hosting a spam bot, comment spamming or other spamming software
    IsListed        bool            `json:"isListed" form:"isListed"` //Is this IP on a blocklist
    IsVpn           bool            `json:"isVpn" form:"isVpn"` //IP has been detected as coming from a VPN hosting provider
    LastSeen        int64           `json:"lastSeen" form:"lastSeen"` //The last time this IP was seen on a blocklist (in Unix time or 0 if not listed recently)
    Blocklists      []string        `json:"blocklists" form:"blocklists"` //An array of strings indicating which blocklists this IP is listed on (empty if not listed)
}

/*
 * Structure for the custom type EmailVerifyResponse
 */
type EmailVerifyResponse struct {
    Valid           bool            `json:"valid" form:"valid"` //Is this a valid email address (syntax and domain is valid)
    Verified        bool            `json:"verified" form:"verified"` //True if this address has passed SMTP verification. Check the smtp-status and smtp-response fields for specific verification details
    Email           string          `json:"email" form:"email"` //The full email address (this could be different to the supplied address if typos-fixed is true)
    TyposFixed      bool            `json:"typosFixed" form:"typosFixed"` //True if typos have been fixed
    SyntaxError     bool            `json:"syntaxError" form:"syntaxError"` //True if this address has a syntax error
    DomainError     bool            `json:"domainError" form:"domainError"` //True if this address has a domain error (e.g. no valid mail server records)
    Domain          string          `json:"domain" form:"domain"` //The email domain
    Provider        string          `json:"provider" form:"provider"` //The email service provider domain
    IsFreemail      bool            `json:"isFreemail" form:"isFreemail"` //True if this address is a free-mail address
    IsDisposable    bool            `json:"isDisposable" form:"isDisposable"` //True if this address is a disposable, temporary or darknet related email address
    IsPersonal      bool            `json:"isPersonal" form:"isPersonal"` //True if this address is for a person. False if this is a role based address, e.g. admin@, help@, office@, etc.
    SmtpStatus      string          `json:"smtpStatus" form:"smtpStatus"` //The SMTP verification status for the address (see online docs for full details)
    SmtpResponse    string          `json:"smtpResponse" form:"smtpResponse"` //The raw SMTP response message received during verification
    IsCatchAll      bool            `json:"isCatchAll" form:"isCatchAll"` //True if this email domain has a catch-all policy (it will accept mail for any username)
    IsDeferred      bool            `json:"isDeferred" form:"isDeferred"` //True if the mail server responded with a temporary failure (either a 4xx response code or unresponsive server). You can retry this address later, we recommend waiting at least 15 minutes before retrying
}

/*
 * Structure for the custom type IPProbeResponse
 */
type IPProbeResponse struct {
    Valid               bool            `json:"valid" form:"valid"` //Is this a valid IPv4 or IPv6 address
    Country             string          `json:"country" form:"country"` //Full country name
    ProviderType        string          `json:"providerType" form:"providerType"` //The detected provider type. See online API docs for specific provider type details
    CountryCode         string          `json:"countryCode" form:"countryCode"` //ISO 2-letter country code
    Hostname            string          `json:"hostname" form:"hostname"` //The IPs hostname (PTR)
    ProviderDomain      string          `json:"providerDomain" form:"providerDomain"` //The domain name of the provider
    City                string          `json:"city" form:"city"` //Full city name (if detectable)
    ProviderWebsite     string          `json:"providerWebsite" form:"providerWebsite"` //The website URL for the provider
    Ip                  string          `json:"ip" form:"ip"` //The IP address
    Region              string          `json:"region" form:"region"` //Full region name (if detectable)
    ProviderDescription string          `json:"providerDescription" form:"providerDescription"` //A description of the provider, usually extracted from the providers website or WHOIS record
    ContinentCode       string          `json:"continentCode" form:"continentCode"` //ISO 2-letter continent code
    IsHosting           bool            `json:"isHosting" form:"isHosting"` //True if this IP belongs to a hosting company. Note that this can still be true even if the provider type is VPN/proxy, this occurs in the case that the IP is detected as both types
    IsIsp               bool            `json:"isIsp" form:"isIsp"` //True if this IP belongs to an ISP. Note that this can still be true even if the provider type is VPN/proxy, this occurs in the case that the IP is detected as both types
}

/*
 * Structure for the custom type UserAgentInfoResponse
 */
type UserAgentInfoResponse struct {
    MobileScreenWidth      int64           `json:"mobileScreenWidth" form:"mobileScreenWidth"` //Mobile device screen width (in px)
    MobileBrand            string          `json:"mobileBrand" form:"mobileBrand"` //Mobile device brand
    MobileModel            string          `json:"mobileModel" form:"mobileModel"` //Mobile device model
    Producer               string          `json:"producer" form:"producer"` //Producer or manufacturer
    BrowserName            string          `json:"browserName" form:"browserName"` //Browser software name
    MobileScreenHeight     int64           `json:"mobileScreenHeight" form:"mobileScreenHeight"` //Mobile device screen height (in px)
    IsMobile               bool            `json:"isMobile" form:"isMobile"` //True if this is a mobile user-agent
    Type                   string          `json:"type" form:"type"` //The user-agent type, possible values are: desktop-browser, email-client, feed-reader, software-library, media-player, mobile-browser, robot, unknown
    Version                string          `json:"version" form:"version"` //Software version
    OperatingSystem        string          `json:"operatingSystem" form:"operatingSystem"` //Operating system
    MobileBrowser          string          `json:"mobileBrowser" form:"mobileBrowser"` //Mobile device browser
    IsAndroid              bool            `json:"isAndroid" form:"isAndroid"` //True if this is an Android based mobile user agent
    IsIos                  bool            `json:"isIos" form:"isIos"` //True if this is an iOS based mobile user agent
    OperatingSystemFamily  string          `json:"operatingSystemFamily" form:"operatingSystemFamily"` //The operating system family name, this is the OS name without any version information
    OperatingSystemVersion string          `json:"operatingSystemVersion" form:"operatingSystemVersion"` //The operating system version number (if detectable)
    Engine                 string          `json:"engine" form:"engine"` //The browser engine name
    EngineVersion          string          `json:"engineVersion" form:"engineVersion"` //The browser engine version (if detectable)
}
