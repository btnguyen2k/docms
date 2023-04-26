import {initRouter} from '@/_shared/router'

const router = initRouter({
    '_editor': () => import('@/_shared/components/Editor.vue'),
    '_search': () => import('@/tpl-bootstrap/components/Search.vue'),
    '_tagsearch': () => import('@/tpl-bootstrap/components/TagSearch.vue'),
    'home': () => import('@/tpl-bootstrap/components/Home.vue'),
    'topic': () => import('@/tpl-bootstrap/components/Topic.vue'),
    'document': () => import('@/tpl-bootstrap/components/Document.vue'),
})

export default router
