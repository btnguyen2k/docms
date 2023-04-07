import i18n from "@/i18n"
import {apiDocument, apiDocuments, apiDoGet, apiSite, apiTopics} from "@/utils/api_client"
import router from '@/router'

export default {
    install: (app) => {
        app.config.unwrapInjectedRef = true

        // use $localedText(inputMap) to pick up the correct i18n message
        app.config.globalProperties.$localedText = (inputMap) => {
            return inputMap && inputMap[i18n.global.locale] ? inputMap[i18n.global.locale] : inputMap
        }

        // use $search(query) to perform search
        app.config.globalProperties.$search = (query) => {
            return router.push({name: 'Search', query: {q: query, l: i18n.global.locale}})
        }

        // use $fetchSiteMeta to fetch site's metadata from server
        app.config.globalProperties.$fetchSiteMeta = (callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            apiDoGet(apiSite,
                (apiResp) => callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined'),
                (err) => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }

        // use $fetchSiteTopics to fetch site's topics metadata from server
        app.config.globalProperties.$fetchSiteTopics = (callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            apiDoGet(apiTopics,
                (apiResp) => callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined'),
                (err) => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }

        // use $fetchDocumentsForTopic to fetch topic's documents metadata from server
        app.config.globalProperties.$fetchDocumentsForTopic = (topicId, callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            apiDoGet(apiDocuments.replaceAll(':topic-id', topicId),
                (apiResp) => callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined'),
                (err) => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }

        // use $fetchDocument to fetch document's metadata from server
        app.config.globalProperties.$fetchDocument = (topicId, documentId, callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            apiDoGet(apiDocument.replaceAll(':topic-id', topicId).replaceAll(':document-id', documentId),
                (apiResp) => callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined'),
                (err) => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }
    }
}
