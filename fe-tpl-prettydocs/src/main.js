//#DO CMS frontend, template PrettyDocs
//import 'core-js/stable'
import { createApp } from 'vue'
import App from './App'

const app = createApp(App)

import i18n from './i18n'
app.use(i18n)

import router from './router'
app.use(router)

import { library } from '@fortawesome/fontawesome-svg-core'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
library.add(fab, far, fas)

// app.provide('icons', icons)
app.component('fa-icon', FontAwesomeIcon)

app.mount('#app')
