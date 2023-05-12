//#DO CMS frontend, template Bootstrap
//import 'core-js/stable'
import {createApp} from 'vue'
import App from './App.vue'

const app = createApp(App)

import i18n from '@/_shared/i18n'
app.use(i18n)

import router from './router'
app.use(router)

import globalPlugin from '@/_shared/plugin_app.js'
app.use(globalPlugin, {
    router: router
})

import AOS from 'aos'
import 'aos/dist/aos.css'
app.use(AOS.init())

app.mount('#app')
