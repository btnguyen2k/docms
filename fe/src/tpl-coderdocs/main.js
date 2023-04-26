//#DO CMS frontend, template CoderDocs
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

app.mount('#app')
