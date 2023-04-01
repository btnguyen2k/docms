/*
DO CMS utilities.

@author Thanh Nguyen <btnguyen2k@gmail.com>
@DOCMS
*/

// import hljs from 'highlight.js'          // all languages
import hljs from 'highlight.js/lib/common'  // only common languages
// import hljs from 'highlight.js/lib/core' // no language
import 'highlight.js/styles/default.css'

import katex from 'katex'
import 'katex/contrib/mhchem/mhchem'
import 'katex/dist/katex.min.css'
import {marked} from "marked"
import DOMPurify from "dompurify"
import {APP_CONFIG} from "@/utils/app_config";

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

class MyRenderer extends marked.Renderer {
    constructor(options) {
        super(options)
    }

    // listitem(text) {
    //     console.log("ListItem: ", text)
    //     return super.listitem(replaceMathWithIds(htmlDecode(text), 'listitem'))
    // }

    paragraph(text) {
        // console.log("Paragraph: ", text)
        return super.paragraph(replaceMathWithIds(htmlDecode(text), 'paragraph'))
    }

    // tablecell(content, flags) {
    //     console.log("TableCell: ", content, flags)
    //     return super.tablecell(replaceMathWithIds(htmlDecode(content), 'tablecell'), flags)
    // }

    text(text) {
        // console.log("Text: ", text)
        return super.text(replaceMathWithIds(htmlDecode(text), 'text'))
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
    // console.log("Test: ", text)
    // block Mathematics and Chemical formulas: allowing newlines between $$...$$
    text = text.replace(/\$\$\s([\s\S]+?)\s\$\$/g, (_match, expression) => {
        expression = unescapeHtml(expression).replace(/\\(\s)/g, (_match, capture) => {
            return '\\\\' + capture
        })
        // console.log("\tBlock match: ", _match)
        // console.log("\texpression: ", expression)
        const id = nextKatexId()
        mathExpMap[id] = {type: 'block', expression, el: el}
        return id
    })

    // inline Mathematics and Chemical formulas: _not_ allowing newlines between $...$
    text = text.replace(/[^$]\$([^$\n]+?)\$/g, (_match, expression) => {
        expression = unescapeHtml(expression)
        // console.log("\tInline match: ", _match)
        // console.log("\texpression: ", expression)
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
    return sanitize ? latexHtml : DOMPurify.sanitize(latexHtml, {ADD_ATTR: ['target']})
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

function styleByHash(input, styleList) {
    let hash = 0
    for (let i = 0; i < input.length; i++) {
        hash = ((hash << 5) - hash) + input.charCodeAt(i)
        hash = hash & hash
    }
    const mod = ((hash % styleList.length) + styleList.length) % styleList.length
    return styleList[mod]
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
    styleByHash, extractLeadingFromName, extractTrailingFromName,
}
