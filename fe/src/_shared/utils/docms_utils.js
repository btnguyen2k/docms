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
import {marked} from 'marked'
import mermaid from 'mermaid'

mermaid.initialize({startOnLoad: false})
import DOMPurify from 'dompurify'
import {APP_CONFIG} from '@/_shared/utils/app_config'

//ref: https://github.com/markedjs/marked/issues/1538#issuecomment-575838181
let katexId = 0
const nextKatexId = () => `__special_katext_id_${katexId++}__`
const mathExpMap = {}
const reKatexId = /(__special_katext_id_\d+__)/g

let bsTabsGroupId = 0
const bsTabsGroupArr = []

let mermaidId = 0

const reUrlWithProtocol = /^([a-z]+:)/i

function _parseParams(params, ignoreFirstN) {
    let result = {}
    const paramsTokens = params.trim().split(/\s+/)
    if (ignoreFirstN > 0) {
        for (let i = 0; i < ignoreFirstN && i < paramsTokens.length; i++) {
            const k = '$' + i
            result[k] = paramsTokens[i].trim()
        }
    } else {
        ignoreFirstN = 0
    }
    for (let i = ignoreFirstN; i < paramsTokens.length; i++) {
        const paramAndValue = paramsTokens[i].trim()
        const paramAndValueTokens = paramAndValue.split(/=/, 2)
        if (paramAndValueTokens.length == 1) {
            result[paramAndValue] = true
        } else {
            result[paramAndValueTokens[0]] = paramAndValueTokens[1]
        }
    }
    return result
}

class MyRenderer extends marked.Renderer {
    constructor(options) {
        super(options)
    }

    _inlineMathToIds(text) {
        if (reKatexId.test(text)) {
            return text
        }
        // inline Mathematics and Chemical formulas: _not_ allowing newlines between $...$
        text = text.replace(/\$(.+?)\$/g, (_match, expression) => {
            expression = unescapeHtml(expression)
            const id = nextKatexId()
            mathExpMap[id] = {type: 'inline', expression}
            return id
        })

        return text
    }

    _renderInlineDoTags(text, tags) {
        text = text.replace(/\[\[do-tag\s(.+?)\]\]/g, (_match, _exp) => {
            const tokens = _exp.trim().split('.')
            let value = tags
            for (let i = 0; i < tokens.length; i++) {
                value = typeof value == 'object' ? value[tokens[i]] : undefined
            }
            if (value == undefined) {
                return '<code title="Error: Tag not found/Invalid value!">' + _exp + '</code>'
            }
            return String(value)
        })
        return text
    }

    _renderGithubGist(gist) {
        const srcScript = "https://gist.github.com/" + gist.trim() + ".js"
        const srcIframe = "data:text/html;charset=utf-8,&lt;head>&lt;base target='_blank' />&lt;/head>&lt;body>&lt;script src='" + srcScript + "'>&lt;/script>&lt;/body>"
        const result = '<div class="{css-class}" data-aos="fade-up"><iframe title="GitHub Gist" src="{src-iframe}"></iframe></div>'
        return result
            .replaceAll('{css-class}', 'ratio ratio-21x9 mb-4')
            .replaceAll('{src-iframe}', srcIframe)
    }

    _renderBootstrapAlert(paramsStr, text) {
        const params = _parseParams(paramsStr, 1)
        const style = ['secondary', 'success', 'danger', 'warning', 'info', 'light', 'dark'].indexOf(params['$0']) >= 0 ? params['$0'] : 'primary'
        const flex = params['flex'] ? true : false
        let result = '<div data-aos="fade-up" class="alert alert-' + style + (flex ? ' d-flex' : '') + ' align-items-center" role="alert">'
        const lines = text.split(/\n/)
        const title = lines.length > 0 && lines[0].trim() != '' ? lines[0].trim() : ''
        let body = ''
        if (lines.length > 1) {
            body = lines.slice(1).join("\n")
        }
        if (title != '') {
            result += '<h4 class="alert-heading">' + title + '</h4>'
        }
        if (body != '') {
            result += '<div>' + markdownRender(body, this.options) + '</div>'
        }
        result += '</div>'
        return result
    }

    _renderBootstrapCards(paramsStr, text) {
        const params = _parseParams(paramsStr)
        const bsCardArr = []
        const errorResult = '<pre title="Cannot parse cards content">' + text + '</pre>'

        const lines = text.replaceAll(/^ {4}/gms, '').split(/\n/)
        let incard = false
        let cardHeader = ''
        let cardImg = ''
        let cardTitle = ''
        let cardSubtitle = ''
        let cardText = ''
        let cardFooter = ''
        let cardParams = {}
        for (let i = 0; i < lines.length; i++) {
            if (lines[i] == '[[bs-card' || lines[i].startsWith('[[bs-card ')) {
                if (incard) {
                    //error
                    return errorResult
                }
                incard = true
                cardParams = _parseParams(lines[i].slice('[[bs-card'.length).trim())
            } else if (lines[i] == ']]') {
                if (!incard) {
                    //error
                    return errorResult
                }
                incard = false
                bsCardArr.push({
                    header: cardHeader,
                    img: cardImg,
                    title: cardTitle,
                    subtitle: cardSubtitle,
                    text: cardText,
                    footer: cardFooter,
                    params: cardParams,
                })
                cardHeader = cardImg = cardTitle = cardSubtitle = cardText = cardFooter = ''
            } else {
                if (incard) {
                    if (lines[i].toLowerCase().startsWith("-header:")) {
                        cardHeader = lines[i].slice('-header:'.length).trim()
                    } else if (lines[i].toLowerCase().startsWith("-img:")) {
                        cardImg = lines[i].slice('-img:'.length).trim()
                    } else if (lines[i].toLowerCase().startsWith("-title:")) {
                        cardTitle = lines[i].slice('-title:'.length).trim()
                    } else if (lines[i].toLowerCase().startsWith("-subtitle:")) {
                        cardSubtitle = lines[i].slice('-subtitle:'.length).trim()
                    } else if (lines[i].toLowerCase().startsWith("-footer:")) {
                        cardFooter = lines[i].slice('-footer:'.length).trim()
                    } else {
                        cardText += lines[i] + '\n'
                    }
                } else if (lines[i].trim() != '') {
                    //error
                    return errorResult
                }
            }
        }

        const cardsEqual = params['equals'] ? true : (params['equal'] ? true : false)
        const lightbox = params['lightbox'] ? params['lightbox'] : false
        let cardsContent = ''
        for (let i = 0; i < bsCardArr.length; i++) {
            const cardData = bsCardArr[i]
            const noMute = cardData.params['no-muted'] ? true : (cardData.params['no-mute'] ? true : false)
            const cssClassText = cardData.params['text'] ? ' text-' + cardData.params['text'] : ''
            const cssClassBorder = cardData.params['border'] ? ' border-' + cardData.params['border'] : ''
            const cssClassBg = cardData.params['bg'] ? ' bg-' + cardData.params['bg'] : (cardData.params['background'] ? ' bg-' + cardData.params['background'] : '')
            const card = '<div class="col" data-aos="fade-up"><div class="card ${cardCssClass}">\n${cardImg}\n${cardHeader}\n<div class="card-body">\n${cardTitle}\n${cardSubTitle}\n${cardText}\n</div>\n${cardFooter}\n</div></div>'
                .replaceAll('${cardHeader}', cardData.header != '' ? '<div class="card-header">' + cardData.header + '</div>' : '')
                .replaceAll('${cardImg}', cardData.img != '' ? (lightbox ? '<a data-type="image" data-gallery="' + ('lightbox-gallery-' + lightbox) + '" data-toggle="lightbox" href="' + cardData.img + '">' : '') + '<img data-aos="zoom-in" src="' + cardData.img + '" class="card-img-top docms-reset">' + (lightbox ? '</a>' : '') : '')
                .replaceAll('${cardTitle}', cardData.title != '' ? '<h5 className="card-title">' + cardData.title + '</h5>' : '')
                .replaceAll('${cardSubTitle}', cardData.subtitle != '' ? '<h6 class="card-subtitle mb-2' + (noMute ? '' : ' text-muted') + '">' + cardData.subtitle + '</h6>' : '')
                .replaceAll('${cardText}', '<div class="card-text">' + markdownRender(cardData.text, this.options) + '</div>')
                .replaceAll('${cardFooter}', cardData.footer != '' ? ('<div class="card-footer' + (noMute ? '' : ' text-muted') + '">' + cardData.footer + '</div>') : '')
                .replaceAll('${cardCssClass}', (cardsEqual ? ' h-100' : '') + cssClassBorder + cssClassText + cssClassBg)
            cardsContent += card
        }
        const colsCssClass = params['cols'] ? ' row-cols-' + params['cols'] : (params['col'] ? ' row-cols-' + params['col'] : '')
        const colsSmCssClass = params['cols-sm'] ? ' row-cols-sm-' + params['cols-sm'] : (params['col-sm'] ? ' row-cols-sm-' + params['col-sm'] : '')
        const colsMdCssClass = params['cols-md'] ? ' row-cols-md-' + params['cols-md'] : (params['col-md'] ? ' row-cols-md-' + params['col-md'] : '')
        const colsLgCssClass = params['cols-lg'] ? ' row-cols-lg-' + params['cols-lg'] : (params['col-lg'] ? ' row-cols-lg-' + params['col-lg'] : '')
        let result = '<div class="row mb-3 g-3' + (colsCssClass + colsSmCssClass + colsMdCssClass + colsLgCssClass) + '">' + cardsContent + '</div>'
        return result
    }

    _renderBootstrapTabs(paramsStr, text) {
        const params = _parseParams(paramsStr)
        const savedIndex = bsTabsGroupId
        const errorResult = '<pre title="Cannot parse tabs content">' + text + '</pre>'
        bsTabsGroupArr.push([])

        const lines = text.replaceAll(/^ {4}/gms, '').split(/\n/)
        let intab = false
        let tabTitle = ''
        let tabBody = ''
        for (let i = 0; i < lines.length; i++) {
            if (lines[i] == '[[bs-tab' || lines[i].startsWith('[[bs-tab ')) {
                if (intab) {
                    //error
                    return errorResult
                }
                intab = true
                tabTitle = lines[i].slice('[[bs-tab'.length).trim()
            } else if (lines[i] == ']]') {
                if (!intab) {
                    //error
                    return errorResult
                }
                intab = false
                bsTabsGroupArr[bsTabsGroupId].push({title: tabTitle, body: tabBody})
                tabTitle = tabBody = ''
            } else {
                if (intab) {
                    tabBody += lines[i] + '\n'
                } else if (lines[i].trim() != '') {
                    //error
                    return errorResult
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
                .replaceAll('${tab-body}', markdownRender(bsTabsGroupArr[savedIndex][i].body, this.options))
            tabContent += pane
        }
        tabHeader += '</ul>'
        tabContent += '</div>'
        let result = '<div data-aos="fade-up" class="container mb-3' + (params['vertical'] ? ' d-flex align-items-start' : '') + '">' + tabHeader + tabContent + '</div>'
        return result
    }

    processInlineElements(text) {
        let result = this._inlineMathToIds(text)
        return result
    }

    code(code, infoString, escaped) {
        infoString = infoString == '' || infoString === undefined ? 'plaintext' : infoString.toLowerCase().trim()
        if (infoString == 'katex' || infoString.startsWith('katex ')) {
            const id = nextKatexId()
            mathExpMap[id] = {type: 'block', expression: code}
            return id
        }
        if (infoString == 'mermaid' || infoString.startsWith('mermaid ')) {
            const mid = 'm' + code.md5() + (mermaidId++)
            mermaid.render(mid, code)
                .then(value => {
                    const elList = document.querySelectorAll('.' + mid)
                    elList.forEach(el => {
                        el.innerHTML = '<div data-aos="fade-up">' + value.svg + '</div>'
                    })
                })
                .finally(() => {
                    //hack do get rid of mermaid's error box
                    const el = document.querySelector("#d" + mid)
                    if (el) {
                        el.remove()
                    }
                })
            return '<pre class="docms-mermaid mermaid ' + mid + '">' + code + '</pre>'
        }
        if (infoString.startsWith('gh-gist ')) {
            return this._renderGithubGist(infoString.slice('gh-gist'.length))
        }
        if (infoString == 'bs-alert' || infoString.startsWith('bs-alert ')) {
            return this._renderBootstrapAlert(infoString.slice('bs-alert'.length), code)
        }
        if (infoString == 'bs-tabs' || infoString.startsWith('bs-tabs ')) {
            return this._renderBootstrapTabs(infoString.slice('bs-tabs'.length), code)
        }
        if (infoString == 'bs-cards' || infoString.startsWith('bs-cards ')) {
            return this._renderBootstrapCards(infoString.slice('bs-cards'.length), code)
        }
        const htmlCode = super.code(code, infoString, escaped)
        return '<div data-aos="fade-up">' + htmlCode + '</div>'
    }

    toc = []

    heading(text, level, raw, slugger) {
        let output = '<h${level} ${cssClass} id="${id}">${text}</h${level}>\n'
        if (this.options.headerIds) {
            const id = this.options.headerPrefix + slugger.slug(raw)
            this.toc.push({id: id, level: level, text: text})
            output = output.replaceAll('${id}', id)
        } else {
            output = '<h${level} ${cssClass}>${text}</h${level}>\n'
        }
        return output.replaceAll('${level}', level)
            .replaceAll('${cssClass}', 'data-aos="fade-up"')
            .replaceAll('${text}', text)
        // const output = super.heading(text, level, raw, slugger)
        // console.log(output, slugger)
        // return output.replaceAll(/^<(h\d+) /gi, '<$1 data-aos="fade-up" ')
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
        const output = super.listitem(this.processInlineElements(text))
        return output.replaceAll(/^<li>/gi, '<li data-aos="fade-up">')
    }

    paragraph(text) {
        const output = super.paragraph(text)
        if (!this.options['inline']) {
            return output.replaceAll(/^<p>/gi, '<p data-aos="fade-up">')
        }
        return output
    }

    text(text) {
        return super.text(this.processInlineElements(text))
    }

    image(href, title, text) {
        const re = /^(https:)|(http:)|(\/)/i
        let beBase = APP_CONFIG.api_client.be_api_base_url
        if (beBase && !re.test(href)) {
            const imgUrl = new URL(href, document.baseURI)
            if (beBase.endsWith("/")) {
                beBase = beBase.slice(0, beBase.length - 1)
            }
            href = beBase + imgUrl.href.slice(imgUrl.origin.length)
        }
        const imgHtml = super.image(href, title, text)
        return imgHtml.replaceAll(/^<img /gi, '<img data-aos="zoom-in" ')
    }

    table(header, body) {
        let output = super.table(header, body)
        output = output.replaceAll('<table>', '<table data-aos="fade-up" class="table table-bordered table-striped">')
        return output
    }

    tablecell(content, flags) {
        return super.tablecell(this.processInlineElements(content), flags)
    }

    blockquote(src) {
        let output = super.blockquote(src)
        output = output.replaceAll('<blockquote>', '<blockquote data-aos="fade-up" class="blockquote">')
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

const defaultMarkedOpts = {
    gfm: true,
    renderer: myRenderer,
    langPrefix: 'hljs language-', // highlight.js css expects a top-level 'hljs' class
    highlight: function (code, lang) {
        const language = hljs.getLanguage(lang) ? lang : 'plaintext'
        const highlightedCode = hljs.highlight(code, {language}).value
        return highlightedCode
    },
}

function markdownRender(markdownInput, opts, tocContainer) {
    //parse: marked
    let markedOpts = {...defaultMarkedOpts}
    opts = typeof opts == 'object' ? opts : {}
    markedOpts = {...markedOpts, ...opts}
    delete markedOpts['sanitize']
    const myRenderer = new MyRenderer(markedOpts)
    markedOpts.renderer = myRenderer

    // process all instances of [[do-tag...]] first
    const tags = typeof opts['tags'] == 'object' ? opts['tags'] : {}
    markdownInput = myRenderer._renderInlineDoTags(markdownInput, tags)

    let html = marked.parse(markdownInput, markedOpts)
    if (opts['inline']) {
        html = html.replaceAll(/^<\s*p\s*>/gi, '').replaceAll(/<\s*\/\s*p\s*>$/gi, '')
    }

    //render: katex
    const latexHtml = html.replace(reKatexId, (_match, capture) => {
        const token = mathExpMap[capture]
        const renderedKatex = katex.renderToString(token.expression, {
            displayMode: token.type == 'block',
            output: 'html',
            throwOnError: false
        })
        return token.type == 'block' ? ('<div data-aos="fade-up">' + renderedKatex + '</div>') : renderedKatex
    })

    if (typeof tocContainer == 'object') {
        tocContainer.value = myRenderer.toc
    }

    const sanitize = typeof opts == 'boolean' ? opts : (typeof opts == 'object' ? opts['sanitize'] : false)
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
