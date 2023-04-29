import {initRouter} from '@/_shared/router'

const router = initRouter({
    '_editor': () => import('@/_shared/components/Editor.vue'),
    '_search': () => import('@/tpl-blogy/components/Search.vue'),
    '_tagsearch': () => import('@/tpl-blogy/components/TagSearch.vue'),
    'home': () => import('@/tpl-blogy/components/Home.vue'),
    'topic': () => import('@/tpl-blogy/components/Topic.vue'),
    'document': () => import('@/tpl-blogy/components/Document.vue'),
})

export default router
