/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io ).
 */

package NeutrinoAPIClient

import(
	"neutrinoapi_lib/configuration_pkg"
	"neutrinoapi_lib/imaging_pkg"
	"neutrinoapi_lib/telephony_pkg"
	"neutrinoapi_lib/datatools_pkg"
	"neutrinoapi_lib/securityandnetworking_pkg"
	"neutrinoapi_lib/geolocation_pkg"
	"neutrinoapi_lib/ecommerce_pkg"
	"neutrinoapi_lib/www_pkg"
)
/*
 * Client structure as interface implementation
 */
type NEUTRINOAPI_IMPL struct {
     imaging imaging_pkg.IMAGING
     telephony telephony_pkg.TELEPHONY
     datatools datatools_pkg.DATATOOLS
     securityandnetworking securityandnetworking_pkg.SECURITYANDNETWORKING
     geolocation geolocation_pkg.GEOLOCATION
     ecommerce ecommerce_pkg.ECOMMERCE
     www www_pkg.WWW
     config  configuration_pkg.CONFIGURATION
}

/**
     * Access to Configuration
     * @return Returns the Configuration instance
*/
func (me *NEUTRINOAPI_IMPL) Configuration() configuration_pkg.CONFIGURATION {
    return me.config
}
/**
     * Access to Imaging controller
     * @return Returns the Imaging() instance
*/
func (me *NEUTRINOAPI_IMPL) Imaging() imaging_pkg.IMAGING {
    if(me.imaging) == nil {
        me.imaging = imaging_pkg.NewIMAGING(me.config)
    }
    return me.imaging
}
/**
     * Access to Telephony controller
     * @return Returns the Telephony() instance
*/
func (me *NEUTRINOAPI_IMPL) Telephony() telephony_pkg.TELEPHONY {
    if(me.telephony) == nil {
        me.telephony = telephony_pkg.NewTELEPHONY(me.config)
    }
    return me.telephony
}
/**
     * Access to DataTools controller
     * @return Returns the DataTools() instance
*/
func (me *NEUTRINOAPI_IMPL) DataTools() datatools_pkg.DATATOOLS {
    if(me.datatools) == nil {
        me.datatools = datatools_pkg.NewDATATOOLS(me.config)
    }
    return me.datatools
}
/**
     * Access to SecurityAndNetworking controller
     * @return Returns the SecurityAndNetworking() instance
*/
func (me *NEUTRINOAPI_IMPL) SecurityAndNetworking() securityandnetworking_pkg.SECURITYANDNETWORKING {
    if(me.securityandnetworking) == nil {
        me.securityandnetworking = securityandnetworking_pkg.NewSECURITYANDNETWORKING(me.config)
    }
    return me.securityandnetworking
}
/**
     * Access to Geolocation controller
     * @return Returns the Geolocation() instance
*/
func (me *NEUTRINOAPI_IMPL) Geolocation() geolocation_pkg.GEOLOCATION {
    if(me.geolocation) == nil {
        me.geolocation = geolocation_pkg.NewGEOLOCATION(me.config)
    }
    return me.geolocation
}
/**
     * Access to ECommerce controller
     * @return Returns the ECommerce() instance
*/
func (me *NEUTRINOAPI_IMPL) ECommerce() ecommerce_pkg.ECOMMERCE {
    if(me.ecommerce) == nil {
        me.ecommerce = ecommerce_pkg.NewECOMMERCE(me.config)
    }
    return me.ecommerce
}
/**
     * Access to WWW controller
     * @return Returns the WWW() instance
*/
func (me *NEUTRINOAPI_IMPL) WWW() www_pkg.WWW {
    if(me.www) == nil {
        me.www = www_pkg.NewWWW(me.config)
    }
    return me.www
}

