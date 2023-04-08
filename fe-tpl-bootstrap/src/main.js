//#DO CMS frontend, template Bootstrap
//import 'core-js/stable'
import {createApp} from 'vue'
import App from './App'

const app = createApp(App)

import i18n from './i18n'
app.use(i18n)

import router from './router'
app.use(router)

import globalPlugin from '@/components/plugin_app.js'
app.use(globalPlugin)

import {library} from '@fortawesome/fontawesome-svg-core'
import {fab} from '@fortawesome/free-brands-svg-icons'
import {far} from '@fortawesome/free-regular-svg-icons'
import {fas} from '@fortawesome/free-solid-svg-icons'
library.add(fab, far, fas)

import {Alert, Button, Carousel, Collapse, Dropdown, Modal, Offcanvas, Popover, ScrollSpy, Tab, Toast, Tooltip} from "bootstrap"
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

app.mount('#app')
