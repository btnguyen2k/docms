import {h, resolveComponent} from 'vue'
import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    meta: {base: '/cms'}, // public metadata, can be accessed from other components

    history: createWebHistory('/cms'),
    // linkActiveClass: 'active',
    scrollBehavior: () => {
        // always scroll to top
        return {top: 0}
    },
    strict: true,
    routes: configRoutes(),
})

export default router

import i18n from '@/i18n'
import {APP_CONFIG} from '@/utils/app_config'

function configRoutes() {
    return [
        {
            path: '/_search', name: 'Search', meta: {label: i18n.global.t('search')},
            component: () => import('@/components/Search'), props: true,
        },
        {
            path: '/_editor', name: 'Editor', meta: {label: 'Editor'},
            component: () => import('@/components/Editor'),
        },
        {
            path: '/:tid/:did/:img', name: 'CaptureImgLinks',
            redirect: to => {
                if (to.params.img.endsWith('.jpg') || to.params.img.endsWith('.jpeg') || to.params.img.endsWith('.gif')
                    || to.params.img.endsWith('.png') || to.params.img.endsWith('.svg')) {
                    window.location = APP_CONFIG.api_client.be_api_base_url + '/img/' + to.params.tid + '/' + to.params.did + '/' + to.params.img
                } else {
                    return {name: 'Home'}
                }
            },
        },
        {
            path: '/',
            component: {
                render() {
                    return h(resolveComponent('router-view'))
                }
            },
            children: [
                {
                    path: '', name: 'Home', meta: {label: i18n.global.t('home')},
                    component: () => import('@/components/Home')
                },
                {
                    path: '/:tid', name: '_Topic', redirect: to => {
                        return {name: 'Topic', params: {tid: to.params.tid}}
                    },
                },
                {
                    path: '/:tid/', name: 'Topic', meta: {label: i18n.global.t('topic')},
                    component: () => import('@/components/Topic')
                },
                {
                    path: '/:tid/:did', name: '_Document', redirect: to => {
                        return {name: 'Document', params: {tid: to.params.tid, did: to.params.did}}
                    },
                },
                {
                    path: '/:tid/:did/', name: 'Document', meta: {label: i18n.global.t('document')},
                    component: () => import('@/components/Document')
                }
            ]
        },
    ]
}
