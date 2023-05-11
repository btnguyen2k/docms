import {h, resolveComponent} from 'vue'
import {createRouter, createWebHistory} from 'vue-router'

export function initRouter(components) {
    const router = createRouter({
        meta: {base: '/cms'}, // public metadata, can be accessed from other components
        history: createWebHistory('/cms'),
        // linkActiveClass: 'active',
        scrollBehavior: () => {
            // always scroll to top
            return {top: 0}
        },
        strict: true,
        routes: configRoutes(components),
    })

    return router
}

function configRoutes(components) {
    return [
        {
            path: '/_editor', name: 'Editor', meta: {label: 'Editor'},
            component: components['_editor'],
        },
        {
            path: '/_search', name: 'Search', meta: {labelId: 'search'},
            component: components['_search'],
        },
        {
            path: '/_tagsearch', name: 'TagSearch', meta: {labelId: 'search'},
            component: components['_tagsearch'],
        },
        // {
        //     path: '/:tid/:did/:img', name: 'CaptureImgLinks',
        //     redirect: to => {
        //         if (to.params.img.endsWith('.jpg') || to.params.img.endsWith('.jpeg') || to.params.img.endsWith('.gif')
        //             || to.params.img.endsWith('.png') || to.params.img.endsWith('.svg')) {
        //             window.location = APP_CONFIG.api_client.be_api_base_url + '/img/' + to.params.tid + '/' + to.params.did + '/' + to.params.img
        //         } else {
        //             return {name: 'Home'}
        //         }
        //     },
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
                    path: '', name: 'Home', meta: {labelId: 'home'},
                    component: components['home'],
                },
                {
                    path: '/:tid', name: '_Topic', redirect: to => {
                        return {name: 'Topic', params: {tid: to.params.tid}}
                    },
                },
                {
                    path: '/:tid/', name: 'Topic', meta: {labelId: 'topic'},
                    component: components['topic'],
                },
                {
                    path: '/:tid/:did', name: '_Document', redirect: to => {
                        return {name: 'Document', params: {tid: to.params.tid, did: to.params.did}}
                    },
                },
                {
                    path: '/:tid/:did/', name: 'Document', meta: {labelId: 'document'},
                    component: components['document'],
                }
            ]
        },
    ]
}
