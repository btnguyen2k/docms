/*
Client to make call to API server using Axios.

@author Thanh Nguyen <btnguyen2k@gmail.com>
@DOCMS
*/
import Axios from "axios"
import {APP_ID, APP_CONFIG} from "./app_config"
import {parseJwt, getLoginSession, saveLoginSession} from "./app_utils"
import i18n from "@/_shared/i18n"

const apiClient = Axios.create({
    baseURL: APP_CONFIG.api_client.be_api_base_url,
    timeout: 10000,
});

const headerAppId = APP_CONFIG.api_client.header_app_id
const headerAccessToken = APP_CONFIG.api_client.header_access_token
const headerLanguage = APP_CONFIG.api_client.header_language
const appId = APP_ID

const apiInfo = "/info"
const apiSite = "/api/site"
const apiTopics = "/api/topics"
const apiDocuments = "/api/documents"
const apiDocumentsForTopic = "/api/documents/:topic-id"
const apiDocument = "/api/document/:topic-id/:document-id"
const apiSearch = "/api/search"
const apiTagSearch = "/api/tag_search"
const apiTags = "/api/tags"

function _apiOnSuccess(method, resp, apiUri, callbackSuccessful) {
    // if (method == 'GET' && Object.prototype.hasOwnProperty.call(resp, "data") && resp.data.status == 403) {
    //     console.error("Error 403 from API [" + apiUri + "], redirecting to login page...")
    //     router.push({name: "Login", query: {app: APP_ID, returnUrl: router.currentRoute.fullPath}})
    //     return
    // }
    if (Object.prototype.hasOwnProperty.call(resp, "data") &&
        Object.prototype.hasOwnProperty.call(resp.data, "extras") &&
        Object.prototype.hasOwnProperty.call(resp.data.extras, "_access_token_")) {
        console.log("Update new access token from API [" + apiUri + "]")
        let jwt = parseJwt(resp.data.extras._access_token_)
        saveLoginSession({uid: jwt.payloadObj.uid, token: resp.data.extras._access_token_})
    }
    if (callbackSuccessful != null) {
        callbackSuccessful(resp.data)
    }
}

function _apiOnError(err, apiUri, callbackError) {
    console.error("Error calling api [" + apiUri + "]: " + err)
    if (callbackError != null) {
        callbackError(err)
    }
}

let cacheExpiryMs = 3600000
let cache = {}

function setCacheExpiry(_cacheExpiryMs) {
    cacheExpiryMs = _cacheExpiryMs
}

function buildHeaders() {
    const session = getLoginSession()
    const headers = {}
    headers[headerAppId] = appId
    headers[headerAccessToken] = session != null ? session.token : ""
    headers[headerLanguage] = i18n.global.locale
    return headers
}

function apiDoGet(apiUri, callbackSuccessful, callbackError) {
    if (cacheExpiryMs > 0) {
        const cacheEntry = cache[apiUri]
        if (cacheEntry && cacheEntry.expiry > new Date().valueOf()) {
            // console.log('Cache hit:', apiUri)
            _apiOnSuccess('GET', cacheEntry.data, apiUri, callbackSuccessful)
            return
        }
    }
    // console.log('Cache missed:', apiUri)
    const headers = buildHeaders()
    return apiClient.get(apiUri, {
        headers: headers, cache: false
    }).then(res => {
        const cacheEntry = {expiry: new Date().valueOf() + cacheExpiryMs, data: res,}
        cache[apiUri] = cacheEntry
        _apiOnSuccess('GET', res, apiUri, callbackSuccessful)
    }).catch(err => _apiOnError(err, apiUri, callbackError))
}

function apiDoPost(apiUri, data, callbackSuccessful, callbackError) {
    const headers = buildHeaders()
    return apiClient.post(apiUri, data, {
        headers: headers,
        cache: false,
    }).then(res => _apiOnSuccess('POST', res, apiUri, callbackSuccessful))
        .catch(err => _apiOnError(err, apiUri, callbackError))
}

export {
    setCacheExpiry,

    apiInfo,
    apiSite,
    apiTopics,
    apiDocuments,
    apiDocumentsForTopic,
    apiDocument,
    apiSearch, apiTagSearch, apiTags,

    apiDoGet,
    apiDoPost,
}
