/*
 * neutrinoapi_lib
 *
 * This file was automatically generated for NeutrinoAPI by APIMATIC v2.0 ( https://apimatic.io )
 */

package NeutrinoAPIClient

import(
	"neutrinoapi_lib/configuration_pkg"
	"neutrinoapi_lib/imaging_pkg"
	"neutrinoapi_lib/telephony_pkg"
	"neutrinoapi_lib/data tools_pkg"
	"neutrinoapi_lib/security and networking_pkg"
	"neutrinoapi_lib/geolocation_pkg"
	"neutrinoapi_lib/e-commerce_pkg"
)
/*
 * Client structure as interface implementation
 */
type NEUTRINOAPI_IMPL struct {
     imaging imaging_pkg.IMAGING
     telephony telephony_pkg.TELEPHONY
     data tools data tools_pkg.DATA TOOLS
     security and networking security and networking_pkg.SECURITY AND NETWORKING
     geolocation geolocation_pkg.GEOLOCATION
     e-commerce e-commerce_pkg.E-COMMERCE
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
func (me * NEUTRINOAPI_IMPL) Imaging() imaging_pkg.IMAGING {
    if(me.imaging) == nil {
        me.imaging = imaging_pkg.NewIMAGING(me.config)
    }
    return me.imaging
}
/**
     * Access to Telephony controller
     * @return Returns the Telephony() instance
*/
func (me * NEUTRINOAPI_IMPL) Telephony() telephony_pkg.TELEPHONY {
    if(me.telephony) == nil {
        me.telephony = telephony_pkg.NewTELEPHONY(me.config)
    }
    return me.telephony
}
/**
     * Access to DataTools controller
     * @return Returns the DataTools() instance
*/
func (me * NEUTRINOAPI_IMPL) DataTools() data tools_pkg.DATA TOOLS {
    if(me.data tools) == nil {
        me.data tools = data tools_pkg.NewDATA TOOLS(me.config)
    }
    return me.data tools
}
/**
     * Access to SecurityAndNetworking controller
     * @return Returns the SecurityAndNetworking() instance
*/
func (me * NEUTRINOAPI_IMPL) SecurityAndNetworking() security and networking_pkg.SECURITY AND NETWORKING {
    if(me.security and networking) == nil {
        me.security and networking = security and networking_pkg.NewSECURITY AND NETWORKING(me.config)
    }
    return me.security and networking
}
/**
     * Access to Geolocation controller
     * @return Returns the Geolocation() instance
*/
func (me * NEUTRINOAPI_IMPL) Geolocation() geolocation_pkg.GEOLOCATION {
    if(me.geolocation) == nil {
        me.geolocation = geolocation_pkg.NewGEOLOCATION(me.config)
    }
    return me.geolocation
}
/**
     * Access to ECommerce controller
     * @return Returns the ECommerce() instance
*/
func (me * NEUTRINOAPI_IMPL) ECommerce() e-commerce_pkg.E-COMMERCE {
    if(me.e-commerce) == nil {
        me.e-commerce = e-commerce_pkg.NewE-COMMERCE(me.config)
    }
    return me.e-commerce
}
