import { h, resolveComponent } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    history: createWebHistory('/cms'),
    // linkActiveClass: 'active',
    scrollBehavior: () => {
        // always scroll to top
        return { top: 0 }
    },
    routes: configRoutes(),
})

export default router

import i18n from '@/i18n'

function configRoutes() {
    return [
        // {
        //     path: '/_error', name: 'Error', meta: { label: i18n.global.t('error') },
        //     component: () => import('@/components/Error'), props: true,
        // },
        {
            path: '/',
            component: {
                render() {
                    return h(resolveComponent('router-view'))
                }
            },
            children: [
                {
                    path: '', name: 'Home', meta: { label: i18n.global.t('home') },
                    component: () => import('@/components/Home')
                },
                {
                    path: '/:tid', name: 'Topic', meta: { label: i18n.global.t('topic') },
                    component: () => import('@/components/Topic')
                },
                {
                    path: '/:tid/:did', name: 'Document', meta: { label: i18n.global.t('document') },
                    component: () => import('@/components/Document')
                }
            ]
        },
    ]
}
