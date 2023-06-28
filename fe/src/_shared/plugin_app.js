//#DO CMS frontend
import i18n from './i18n'
import {
    apiInfo,
    apiDocument,
    apiDocuments,
    apiDocumentsForTopic,
    apiDoGet,
    apiDoPost,
    apiSearch,
    apiSite, apiTags,
    apiTagSearch,
    apiTopics
} from './utils/api_client'
import {extractLeadingFromName, extractTrailingFromName} from './utils/docms_utils'
import {computed} from 'vue'
import MD5 from 'crypto-js/md5'

String.prototype.md5 = function () {
    return MD5(this).toString()
}

import {library} from '@fortawesome/fontawesome-svg-core'
import {fab} from '@fortawesome/free-brands-svg-icons'
import {far} from '@fortawesome/free-regular-svg-icons'
import {fas} from '@fortawesome/free-solid-svg-icons'

import {
    Alert,
    Button,
    Carousel,
    Collapse,
    Dropdown,
    Modal,
    Offcanvas,
    Popover,
    ScrollSpy,
    Tab,
    Toast,
    Tooltip
} from 'bootstrap'

import VueGtag from 'vue-gtag'

class Global {
    get router() {
        return this._router
    }

    set router(value) {
        this._router = value
    }

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

    get tagCloud() {
        return this._tagCloud ? this._tagCloud : {}
    }

    set tagCloud(value) {
        this._tagCloud = value
    }
}

function initGtag(app, global) {
    if (!global.router) {
        setTimeout(() => {
            initGtag(app, global)
        }, 100)
        return
    }

    const data = global.serverInfo
    app.use(VueGtag, {
        config: {id: data.tracking.gtag},
        appName: data.app.name,
        pageTrackerScreenviewEnabled: true
    }, global.router)
}

function initGlobal(app, global) {
    apiDoGet(apiInfo,
        apiResp => {
            if (apiResp.status == 200) {
                console.log(apiResp.data)
                console.log(apiResp.extras)
                global.serverInfo = apiResp.data
                initGtag(app, global)
            } else {
                console.error("Error while fetching server info, retry shortly: ", apiResp)
                setTimeout(() => {
                    initGlobal(app, global)
                }, 5000)
            }
        },
        (err) => {
            console.error("Error while fetching server info, retry shortly: ", err)
            setTimeout(() => {
                initGlobal(app, global)
            }, 5000)
        },
    )
}

export default {
    install: (app, params) => {
        app.config.unwrapInjectedRef = true

        // use FontAwesome icons
        library.add(fab, far, fas)

        // enable Bootstrap components
        app.component(Alert)
        app.component(Button)
        app.component(Carousel)
        app.component(Collapse)
        app.component(Dropdown)
        app.component(Modal)
        app.component(Offcanvas)
        app.component(Popover)
        app.component(ScrollSpy)
        app.component(Tab)
        app.component(Toast)
        app.component(Tooltip)

        /*-- read/write global variable */
        let global = new Global()
        app.provide('$global', global)
        global.router = params.router ? params.router : undefined

        initGlobal(app, global)

        /*-- read-only global variable */
        app.provide('$searchQuery', computed(() => global.searchQuery))
        app.provide('$siteMeta', computed(() => global.siteMeta))
        app.provide('$siteLanguages', computed(() => global.siteLanguages))
        app.provide('$siteFirstName', computed(() => global.siteFirstName))
        app.provide('$siteLastName', computed(() => global.siteLastName))
        app.provide('$siteTopics', computed(() => global.siteTopics))
        app.provide('$tagCloud', computed(() => global.tagCloud))

        // use $reload() to reload the browser tab
        app.config.globalProperties.$reload = () => {
            window.location.reload()
        }

        // use $transferToHome(seconds) to redirect to Home page
        app.config.globalProperties.$transferToHome = (delayInSeconds) => {
            setTimeout(() => {
                return global.router.push({name: 'Home'})
            }, delayInSeconds * 1000)
        }

        // use $transferToTopic(topicId, seconds) to redirect to Home page
        app.config.globalProperties.$transferToTopic = (topicId, delayInSeconds) => {
            setTimeout(() => {
                return global.router.push({name: 'Topic', params: {tid: topicId}})
            }, delayInSeconds * 1000)
        }

        // use $updatePageTitle({...}) update browser's title
        app.config.globalProperties.$updatePageTitle = (opts) => {
            opts = opts ? opts : {}
            const calcTitle = () => {
                const siteName = global.siteMeta.name
                let title = ''
                if (opts.topic) {
                    title = app.config.globalProperties.$localedText(opts.topic.title)
                } else if (opts.document || opts.doc) {
                    const doc = opts.document ? opts.document : opts.doc
                    title = app.config.globalProperties.$localedText(doc.title)
                } else if (opts.search) {
                    title = i18n.global.t('search') + ': ' + opts.search
                }
                if (title) {
                    return title + (siteName ? (' | ' + siteName) : '')
                }
                return siteName + ' | ' + app.config.globalProperties.$localedText(global.siteMeta.description)
            }
            document.title = calcTitle()
        }

        // use $pickupFromHash(input, list) to pick up one item from the list based on hash of input
        app.config.globalProperties.$pickupFromHash = (input, list) => {
            let hash = 0
            for (let i = 0; i < input.length; i++) {
                hash = ((hash << 5) - hash) + input.charCodeAt(i)
                hash = hash & hash
            }
            const mod = ((hash % list.length) + list.length) % list.length
            return list[mod]
        }

        // use $trimText(input, maxLength) to trim a text down to a max-length
        app.config.globalProperties.$trimText = (input, maxLength) => {
            const result = input.replace(new RegExp('^(.{' + maxLength + '}[^\\s]*).*'), "$1")
            return result.length < input.length ? result + '...' : result
        }

        // use $localedText(inputMap) to pick up the correct i18n message
        app.config.globalProperties.$localedText = (inputMap) => {
            return inputMap && inputMap[i18n.global.locale] ? inputMap[i18n.global.locale] : inputMap
        }

        // use $doSearch to navigate to search view
        app.config.globalProperties.$doSearch = () => {
            return global.router.push({name: 'Search', query: {q: global.searchQuery, l: i18n.global.locale}})
        }
        // use $doTagSearch to navigate to tag-search view
        app.config.globalProperties.$doTagSearch = () => {
            return global.router.push({name: 'TagSearch', query: {q: global.searchQuery, l: i18n.global.locale}})
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
        // use $tagSearch(query) to perform search
        app.config.globalProperties.$tagSearch = (query, callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            apiDoPost(apiTagSearch, {query: query},
                apiResp => callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined'),
                err => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }

        // use $fetchTagCloud to fetch site's tag-cloud from server
        app.config.globalProperties.$fetchTagCloud = (callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            apiDoGet(apiTags,
                apiResp => {
                    if (apiResp.status == 200) {
                        global.tagCloud = apiResp.data
                        let tagSize = []
                        for (let t in global.tagCloud[i18n.global.locale]) {
                            tagSize.push(global.tagCloud[i18n.global.locale][t].length)
                        }
                        global.minTagSize = Math.min(...tagSize)
                        global.maxTagSize = Math.max(...tagSize)
                    }
                    callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined')
                },
                err => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }
        app.config.globalProperties.$calcTagSize = (tag, minSize, maxSize, steps) => {
            let sizeList = []
            for (let i = 0; i < steps + 1; i++) {
                const val = minSize + (maxSize - minSize) * i / steps
                sizeList.push(val)
            }
            const size = global.tagCloud[i18n.global.locale][tag] ? global.tagCloud[i18n.global.locale][tag].length : 0
            if (size < global.minTagSize) {
                return 0
            }
            if (global.maxTagSize - global.minTagSize < sizeList.length) {
                return sizeList[size - global.minTagSize]
            }
            for (let i = 0; i < steps + 1; i++) {
                const val = global.minTagSize + (global.maxTagSize - global.minTagSize) * i / steps
                if (val >= size) {
                    return sizeList[i]
                }
            }
            return sizeList[0]
        }

        const emptyFunc = () => {
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
                        app.config.globalProperties.$fetchTagCloud(null, emptyFunc, emptyFunc)
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
            apiDoGet(apiDocumentsForTopic.replaceAll(':topic-id', topicId),
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

        // use $fetchLatestDocuments to fetch latest documents' metadata from server
        app.config.globalProperties.$fetchLatestDocuments = (topicId, numDocs, callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            const uri = apiDocuments + '?p=latest&n=' + numDocs + (topicId ? '&t=' + topicId : '')
            apiDoGet(uri,
                apiResp => callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined'),
                err => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }
        // use $fetchSpecialDocuments to fetch special purpose documents' metadata from server
        app.config.globalProperties.$fetchSpecialDocuments = (callbackPrefetch, callbackSuccess, callbackError) => {
            if (callbackPrefetch) {
                callbackPrefetch()
            }
            const uri = apiDocuments + '?p=special'
            apiDoGet(uri,
                apiResp => callbackSuccess ? callbackSuccess(apiResp) : console.error('no success-callback function defined'),
                err => callbackError ? callbackError(err) : console.error('no error-callback function defined', err),
            )
        }
    }
}
