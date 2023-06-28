//#DO CMS frontend, template PrettyDocs
//import 'core-js/stable'
import {createApp} from 'vue'
import App from './App'

const app = createApp(App)

import i18n from '@/_shared/i18n'
app.use(i18n)

import router from './router'
app.use(router)

import globalPlugin from '@/_shared/plugin_app.js'
app.use(globalPlugin, {
    router: router
})

/* PrettyDocs only */
app.config.globalProperties.$styleByHash = (input, styleList) => {
    let hash = 0
    for (let i = 0; i < input.length; i++) {
        hash = ((hash << 5) - hash) + input.charCodeAt(i)
        hash = hash & hash
    }
    const mod = ((hash % styleList.length) + styleList.length) % styleList.length
    return styleList[mod]
}

import AOS from 'aos'
import 'aos/dist/aos.css'
app.use(AOS.init())

app.mount('#app')
