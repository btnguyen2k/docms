import {initRouter} from '@/_shared/router'

const router = initRouter({
    '_editor': () => import('@/_shared/components/Editor.vue'),
    '_search': () => import('@/tpl-prettydocs/components/Search.vue'),
    '_tagsearch': () => import('@/tpl-prettydocs/components/TagSearch.vue'),
    'home': () => import('@/tpl-prettydocs/components/Home.vue'),
    'topic': () => import('@/tpl-prettydocs/components/Topic.vue'),
    'document': () => import('@/tpl-prettydocs/components/Document.vue'),
})

export default router
