import i18n from "@/i18n"
import {apiDocument, apiDocuments, apiDoGet, apiDoPost, apiSearch, apiSite, apiTopics} from "@/utils/api_client"
import router from '@/router'
import {extractLeadingFromName, extractTrailingFromName} from "@/components/utils"

class Global {
    get siteMeta() {
        return this._siteMeta ? this._siteMeta : {}
    }

    set siteMeta(value) {
        this._siteMeta = value
        this._siteLanguages = Object.keys(this._siteMeta.languages).filter(value => {
            return value != 'default'
        })
    }

    get siteTopics() {
        return this._siteTopics ? this._siteTopics : []
    }

    set siteTopics(value) {
        this._siteTopics = value
    }

    get siteLanguages() {
        return this._siteLanguages ? this._siteLanguages : []
    }

    get siteFirstName() {
        return this._siteMeta.name ? extractLeadingFromName(this._siteMeta.name) : ''
    }

    get siteLastName() {
        return this._siteMeta.name ? extractTrailingFromName(this._siteMeta.name) : ''
    }

    get searchQuery() {
        return this._searchQuery !== undefined ? this._searchQuery : ''
    }

    set searchQuery(value) {
        this._searchQuery = value
    }

    get sidebar() {
        return this._sidebar
    }

    set sidebar(value) {
        this._sidebar = value
    }
}

import {computed} from 'vue'

export default {
    install: (app) => {
        app.config.unwrapInjectedRef = true

        /*-- read/write global variable */
        let global = new Global()
        app.provide('$global', global)

        /*-- read-only global variable */
        app.provide('$searchQuery', computed(() => global.searchQuery))
        app.provide('$siteMeta', computed(() => global.siteMeta))
        app.provide('$siteLanguages', computed(() => global.siteLanguages))
        app.provide('$siteFirstName', computed(() => global.siteFirstName))
        app.provide('$siteLastName', computed(() => global.siteLastName))
        app.provide('$siteTopics', computed(() => global.siteTopics))

        // use $localedText(inputMap) to pick up the correct i18n message
        app.config.globalProperties.$localedText = (inputMap) => {
            return inputMap && inputMap[i18n.global.locale] ? inputMap[i18n.global.locale] : inputMap
        }

        // use $doSearch to navigate to search view
        app.config.globalProperties.$doSearch = () => {
            return router.push({name: 'Search', query: {q: global.searchQuery, l: i18n.global.locale}})
        }

        // use $search(query) to perform search
        app.config.globalProperties.$search = (query, callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            apiDoPost(apiSearch, {query: query},
                apiResp => callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined'),
                err => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }

        // use $fetchSiteMeta to fetch site's metadata from server
        app.config.globalProperties.$fetchSiteMeta = (callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            apiDoGet(apiSite,
                apiResp => {
                    if (apiResp.status == 200) {
                        global.siteMeta = apiResp.data
                    }
                    callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined')
                },
                err => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }

        // use $fetchSiteTopics to fetch site's topics metadata from server
        app.config.globalProperties.$fetchSiteTopics = (callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            apiDoGet(apiTopics,
                apiResp => {
                    if (apiResp.status == 200) {
                        global.siteTopics = apiResp.data
                    }
                    callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined')
                },
                err => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }

        // use $fetchDocumentsForTopic to fetch topic's documents metadata from server
        app.config.globalProperties.$fetchDocumentsForTopic = (topicId, callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            apiDoGet(apiDocuments.replaceAll(':topic-id', topicId),
                apiResp => callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined'),
                err => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }

        // use $fetchDocument to fetch document's metadata from server
        app.config.globalProperties.$fetchDocument = (topicId, documentId, callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            apiDoGet(apiDocument.replaceAll(':topic-id', topicId).replaceAll(':document-id', documentId),
                apiResp => callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined'),
                err => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }

        /* CoderDocs only */
        app.config.globalProperties.$coderDocsResponsiveSidebar = () => {
            const w = window.innerWidth
            const sidebar = global.sidebar
            if (sidebar) {
                if (w >= 1024) {
                    sidebar.classList.remove('sidebar-hidden')
                    sidebar.classList.add('sidebar-visible')
                } else {
                    sidebar.classList.remove('sidebar-visible')
                    sidebar.classList.add('sidebar-hidden')
                }
            }
        }

        /* CoderDocs only */
        app.config.globalProperties.$coderDocsSidebarToggler = () => {
            const sidebar = global.sidebar
            if (sidebar) {
                if (sidebar.classList.contains('sidebar-visible')) {
                    sidebar.classList.remove('sidebar-visible');
                    sidebar.classList.add('sidebar-hidden');
                } else {
                    sidebar.classList.remove('sidebar-hidden');
                    sidebar.classList.add('sidebar-visible');
                }
            }
        }
    }
}
