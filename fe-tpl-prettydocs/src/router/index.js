import { h, resolveComponent } from 'vue'
// eslint-disable-next-line no-unused-vars
import { createRouter, createWebHashHistory, createMemoryHistory } from 'vue-router'

const router = createRouter({
    // history: createWebHashHistory(''),
    history: createMemoryHistory(''),
    linkActiveClass: 'active',
    scrollBehavior: () => {
        // always scroll to top
        return { top: 0 }
    },
    routes: configRoutes(),
})

// import appConfig from '@/utils/app_config'
// import utils from '@/utils/app_utils'
// import clientUtils from '@/utils/api_client'

// router.beforeEach((to, from, next) => {
//     if (!to.matched.some((record) => record.meta.allowGuest)) {
//         let session = utils.loadLoginSession()
//         if (session == null) {
//             //redirect to login page if not logged in
//             return next({
//                 name: 'Login',
//                 query: { returnUrl: router.resolve(to, from).href },
//             })
//         }
//         let lastUserTokenCheck = utils.localStorageGetAsInt(
//             utils.lskeyLoginSessionLastCheck,
//         )
//         if (lastUserTokenCheck + 60 < utils.getUnixTimestamp()) {
//             lastUserTokenCheck = utils.getUnixTimestamp()
//             let token = session.token
//             clientUtils.apiDoPost(
//                 clientUtils.apiVerifyLoginToken,
//                 { app: appConfig.APP_ID, token: token },
//                 (apiRes) => {
//                     if (apiRes.status != 200) {
//                         //redirect to login page if session verification failed
//                         console.error(
//                             'Session verification failed: ' + JSON.stringify(apiRes),
//                         )
//                         return next({
//                             name: 'Login',
//                             query: { returnUrl: router.resolve(to, from).href },
//                         })
//                         // return next({name: "Login", query: {returnUrl: to.fullPath}})
//                     } else {
//                         utils.localStorageSet(
//                             utils.lskeyLoginSessionLastCheck,
//                             lastUserTokenCheck,
//                         )
//                         next()
//                     }
//                 },
//                 (err) => {
//                     console.error('Session verification error: ' + err)
//                     //redirect to login page if cannot verify session
//                     return next({
//                         name: 'Login',
//                         query: { returnUrl: router.resolve(to, from).href },
//                     })
//                     // return next({name: "Login", query: {returnUrl: to.fullPath}})
//                 },
//             )
//         } else {
//             next()
//         }
//     } else {
//         next()
//     }
// })

export default router

import i18n from '@/i18n'

function configRoutes() {
    return [
        {
            path: '/_error', name: 'Error', meta: { label: i18n.global.t('message.error') },
            component: () => import('@/components/Error'), props: true,
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
                    path: '', name: 'Home', meta: { label: i18n.global.t('message.gome') },
                    component: () => import('@/components/Home')
                },
                {
                    path: '/:tid', name: 'Topic', meta: { label: i18n.global.t('message.topic') },
                    component: () => import('@/components/Topic')
                },
                {
                    path: '/:tid/:pid', name: 'Page', meta: { label: i18n.global.t('message.page') },
                    component: () => import('@/components/Page')
                }
            ]
        },
    ]
}
