/*
Application configurations.

@author Thanh Nguyen <btnguyen2k@gmail.com>
@DOCMS
*/
const APP_CONFIG = require('@/_shared/config.json')
if (!APP_CONFIG.api_client.be_api_base_url) {
    APP_CONFIG.api_client.be_api_base_url = process.env.VUE_APP_BE_API_BASE_URL
}
const APP_ID = APP_CONFIG.api_client.app_id
export {
    APP_ID,
    APP_CONFIG
}
