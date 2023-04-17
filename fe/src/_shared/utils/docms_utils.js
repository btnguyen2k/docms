/*
DO CMS utilities.

@author Thanh Nguyen <btnguyen2k@gmail.com>
@DOCMS
*/

// import hljs from 'highlight.js'          // all languages
import hljs from 'highlight.js/lib/common'  // only common languages
// import hljs from 'highlight.js/lib/core' // no language
import 'highlight.js/styles/default.css'
/* additional languages if not include "all" */
hljs.registerLanguage('dockerfile', require('highlight.js/lib/languages/dockerfile'));

import katex from 'katex'
import 'katex/contrib/mhchem/mhchem'
import 'katex/dist/katex.min.css'
import {marked} from "marked"
import DOMPurify from "dompurify"
import {APP_CONFIG} from "@/_shared/utils/app_config"

function htmlDecode(html) {
    // TODO
    return html
    // const doc = new DOMParser().parseFromString(html, "text/html")
    // return doc.documentElement.textContent
}

//ref: https://github.com/markedjs/marked/issues/1538#issuecomment-575838181
let katexId = 0
const nextKatexId = () => `__special_katext_id_${katexId++}__`
const mathExpMap = {}

const reUrlWithProtocol = /^([a-z]+:)/i

function renderGithubGist(gist) {
    const srcScript = "https://gist.github.com/" + gist.trim() + ".js"
    const srcIframe = "data:text/html;charset=utf-8,&lt;head>&lt;base target='_blank' />&lt;/head>&lt;body>&lt;script src='" + srcScript + "'>&lt;/script>&lt;/body>"
    let result = "<div class=\"ratio ratio-21x9\"><iframe title=\"GitHub Gist\" src=\"" + srcIframe + "\"></iframe></div>"
    return result
}

function _parseParams(params, ignoreFirstN) {
    let result = {}
    const tokens = params.trim().split(/\s+/)
    if (ignoreFirstN > 0) {
        for (let i = 0; i < ignoreFirstN && i < tokens.length; i++) {
            const k = '$' + i
            result[k] = tokens[i].trim()
        }
    } else {
        ignoreFirstN = 0
    }
    for (let i = ignoreFirstN; i < tokens.length; i++) {
        const k = tokens[i].trim()
        result[k] = true
    }
    return result
}

function renderBootstrapAlert(paramsStr, text) {
    const params = _parseParams(paramsStr, 1)
    const style = ['secondary', 'success', 'danger', 'warning', 'info', 'light', 'dark'].indexOf(params['$0']) >= 0 ? params['$0'] : 'primary'
    const flex = params['flex'] ? true : false
    let result = '<div class="alert alert-' + style + (flex ? ' d-flex' : '') + ' align-items-center" role="alert">'
    const lines = text.split(/[\r\n]+/)
    const title = lines.length > 0 && lines[0].trim() != '' ? lines[0].trim() : ''
    let body = ''
    if (lines.length > 1) {
        body = lines.slice(1).join("\n")
    }
    if (title != '') {
        result += '<h4 class="alert-heading">' + title + '</h4>'
    }
    if (body != '') {
        result += '<div>' + markdownRender(body, true) + '</div>'
    }
    result += '</div>'
    return result
}

let bsTabsGroupId = 0
let bsTabsGroupArr = []

function renderBootstrapTabs(paramsStr, text) {
    const params = _parseParams(paramsStr)
    const savedIndex = bsTabsGroupId
    bsTabsGroupArr.push([])

    const lines = text.replaceAll(/^ {4}/gms, '').split(/[\r\n]+/)
    let intab = false
    let tabTitle = ''
    let tabBody = ''
    for (let i = 0; i < lines.length; i++) {
        if (lines[i] == '[[bs-tab' || lines[i].startsWith('[[bs-tab ')) {
            if (intab) {
                //error
                return '<pre>' + text + '</pre>'
            }
            intab = true
            tabTitle = lines[i].slice('[[bs-tab'.length).trim()
        } else if (lines[i] == ']]') {
            if (!intab) {
                //error
                return '<pre>' + text + '</pre>'
            }
            intab = false
            bsTabsGroupArr[bsTabsGroupId].push({title: tabTitle, body: tabBody})
            tabTitle = ''
            tabBody = ''
        } else {
            if (intab) {
                tabBody += lines[i] + '\n'
            } else if (lines[i].trim()!=''){
                //error
                return '<pre>' + text + '</pre>'
            }
        }
    }
    bsTabsGroupId++
    
    let tabHeader = '<ul class="nav ' + (params['vertical'] ? 'nav-pills flex-column me-3' : 'nav-tabs') + '" id="tabGroup-' + savedIndex + '" role="tablist"' + (params['vertical'] ? ' aria-orientation="vertical"' : '') + '>'
    let tabContent = '<div class="tab-content' + (params['vertical'] ? '' : ' border') + '">'
    for (let i = 0; i < bsTabsGroupArr[savedIndex].length; i++) {
        const tabId = 'tab-' + savedIndex + '-' + i
        const tabTarget = 'tabtarget-' + savedIndex + '-' + i

        const btn = '<button class="${btn-class}" id="${tab-id}" data-bs-toggle="tab" data-bs-target="#${tab-target}" type="button" role="tab" aria-controls="${tab-target}" aria-selected="${selected}">${tab-title}</button>'
            .replaceAll('${btn-class}', 'nav-link' + (i == 0 ? ' active' : ''))
            .replaceAll('${tab-id}', tabId)
            .replaceAll('${tab-target}', tabTarget)
            .replaceAll('${selected}', i == 0 ? 'true' : 'false')
            .replaceAll('${tab-title}', escapeHtml(bsTabsGroupArr[savedIndex][i].title))
        tabHeader += '<li class="nav-item" role="presentation">' + btn + '</li>'

        const pane = '<div class="${pane-class}" id="${tab-target}" role="tabpanel" aria-labelledby="${tab-id}">${tab-body}</div>'
            .replaceAll('${pane-class}', 'tab-pane fade' + (i == 0 ? ' show active' : '') + (params['vertical'] ? ' p-0' : ' p-2'))
            .replaceAll('${tab-id}', tabId)
            .replaceAll('${tab-target}', tabTarget)
            .replaceAll('${tab-body}', markdownRender(bsTabsGroupArr[savedIndex][i].body, true))
        tabContent += pane
    }
    tabHeader += '</ul>'
    tabContent += '</div>'
    let result = '<div class="container mb-3' + (params['vertical'] ? ' d-flex align-items-start' : '') + '">' + tabHeader + tabContent + '</div>'
    return result
}

function renderBootstrapTab(params, text) {
    const lines = text.split(/[\r\n]+/)
    const tabTitle = lines[0].trim()
    const tabBody = lines.length > 1 ? lines.slice(1).join('\n') : ''
    bsTabsGroupArr[bsTabsGroupId].push({title: tabTitle, body: tabBody})
    return ''
}

class MyRenderer extends marked.Renderer {
    constructor(options) {
        super(options)
    }

    code(code, infoString, escaped) {
        infoString = infoString == '' || infoString === undefined ? 'plaintext' : infoString.toLowerCase().trim()
        if (infoString == 'katex') {
            const id = nextKatexId()
            mathExpMap[id] = {type: 'block', expression: code/*, el: el*/}
            return id
        }
        if (infoString.startsWith('gh-gist ')) {
            return renderGithubGist(infoString.slice('gh-gist'.length))
        }
        if (infoString == 'bs-alert' || infoString.startsWith('bs-alert ')) {
            return renderBootstrapAlert(infoString.slice('bs-alert'.length), code)
        }
        if (infoString == 'bs-tabs' || infoString.startsWith('bs-tabs ')) {
            return renderBootstrapTabs(infoString.slice('bs-tabs'.length), code)
        }
        if (infoString == 'bs-tab' || infoString.startsWith('bs-tab ')) {
            return renderBootstrapTab(infoString.slice('bs-tab'.length), code)
        }
        return super.code(code, infoString, escaped)
    }

    link(href, title, text) {
        let result = super.link(href, title, text)
        if (reUrlWithProtocol.test(href)) {
            if (text.startsWith('<img')) {
                result = result.replaceAll('<a', '<a target="_blank" ')
            } else {
                result = result.replaceAll('<a', '<a class="external-link" target="_blank" ')
            }
        }
        return result
    }

    listitem(text) {
        return super.listitem(replaceMathWithIds(htmlDecode(text), 'listitem'))
    }

    paragraph(text) {
        return super.paragraph(replaceMathWithIds(htmlDecode(text), 'paragraph'))
    }

    // tablecell(content, flags) {
    //     return super.tablecell(replaceMathWithIds(htmlDecode(content), 'tablecell'), flags)
    // }

    text(text) {
        return super.text(replaceMathWithIds(htmlDecode(text), 'text'))
    }

    image(href, title, text) {
        // console.log(href, title, text)
        const re = /^(https:)|(http:)|(\/)/i
        let beBase = APP_CONFIG.api_client.be_api_base_url
        if (beBase && !re.test(href)) {
            const imgUrl = new URL(href, document.baseURI)
            if (beBase.endsWith("/")) {
                beBase = beBase.slice(0, beBase.length - 1)
            }
            href = beBase + imgUrl.href.slice(imgUrl.origin.length)
        }
        return super.image(href, title, text)
    }

    table(header, body) {
        let output = super.table(header, body)
        output = output.replaceAll('<table>', '<table class="table table-bordered table-striped">')
        return output
    }

    blockquote(src) {
        let output = super.blockquote(src)
        output = output.replaceAll('<blockquote>', '<blockquote class="blockquote">')
        return output
    }

    checkbox(checked) {
        let output = super.checkbox(checked)
        output = output.replaceAll('<input ', '<input class="form-check-input" ')
        return output
    }
}

const myRenderer = new MyRenderer()

function escapeHtml(html) {
    return html.replace(/</g, '&lt;')
        .replace(/>/g, '&gt;')
        .replace(/"/g, '&quot;')
        .replace(/'/g, '&#39;')
        .replace(/&/g, '&amp;')
}

function unescapeHtml(html) {
    return html.replace(/&lt;/g, '<')
        .replace(/&gt;/g, '>')
        .replace(/&quot;/g, '"')
        .replace(/&#39;/g, "'")
        .replace(/&amp;/g, '&')
}

const reKatexId = /(__special_katext_id_\d+__)/g

function replaceMathWithIds(text, el) {
    if (reKatexId.test(text)) {
        return text
    }
    // block Mathematics and Chemical formulas: allowing newlines between $$...$$
    text = text.replace(/\$\$\s([\s\S]+?)\s\$\$/g, (_match, expression) => {
        expression = unescapeHtml(expression).replace(/\\(\s)/g, (_match, capture) => {
            return '\\\\' + capture
        })
        const id = nextKatexId()
        mathExpMap[id] = {type: 'block', expression, el: el}
        return id
    })

    // inline Mathematics and Chemical formulas: _not_ allowing newlines between $...$
    text = text.replace(/\$([^$\n]+?)\$/g, (_match, expression) => {
        expression = unescapeHtml(expression)
        const id = nextKatexId()
        mathExpMap[id] = {type: 'inline', expression, el: el}
        return id
    })

    return text
}

const markedOpts = {
    gfm: true,
    renderer: myRenderer,
    langPrefix: 'hljs language-', // highlight.js css expects a top-level 'hljs' class
    highlight: function (code, lang) {
        const language = hljs.getLanguage(lang) ? lang : 'plaintext'
        return hljs.highlight(code, {language}).value
    },
}

function markdownRender(markdownInput, sanitize) {
    const html = marked.parse(markdownInput, markedOpts)
    const latexHtml = html.replace(reKatexId, (_match, capture) => {
        const token = mathExpMap[capture]
        return katex.renderToString(token.expression, {
            displayMode: token.type == 'block',
            output: 'html',
            throwOnError: false
        })
    })
    return sanitize ? DOMPurify.sanitize(latexHtml, {
        ADD_TAGS: ['iframe'], ADD_DATA_URI_TAGS: ['iframe'], // needed for embed GitHug Gist
        ADD_ATTR: ['target'],
    }) : latexHtml
}

/*----------------------------------------------------------------------*/

let popstateHandler = null

function registerPopstate(func) {
    popstateHandler = func
}

function unregisterPopstate(func) {
    if (func == popstateHandler) {
        popstateHandler = null
        return true
    }
    return false
}

function triggerPopstate() {
    if (popstateHandler) {
        popstateHandler()
    }
}

let resizeHandler = null

function registerResize(func) {
    resizeHandler = func
}

function unregisterResize(func) {
    if (func == resizeHandler) {
        resizeHandler = null
        return true
    }
    return false
}

function triggerResize() {
    if (resizeHandler) {
        resizeHandler()
    }
}

function extractLeadingFromName(input) {
    input = input.trim()
    let tokens = input.split(' ')
    return tokens.length > 1 ? tokens[0].trim() : input.slice(0, 2)
}

function extractTrailingFromName(input) {
    input = input.trim()
    let tokens = input.split(' ')
    return tokens.length > 1 ? input.slice(tokens[0].length).trim() : input.slice(2)
}

export {
    registerPopstate, unregisterPopstate, triggerPopstate,
    registerResize, unregisterResize, triggerResize,
    markdownRender,
    extractLeadingFromName, extractTrailingFromName,
}
