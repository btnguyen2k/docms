import {initRouter} from '@/_shared/router'

const router = initRouter({
    '_editor': () => import('@/_shared/components/Editor.vue'),
    '_search': () => import('@/tpl-coderdocs/components/Search.vue'),
    '_tagsearch': () => import('@/tpl-coderdocs/components/TagSearch.vue'),
    'home': () => import('@/tpl-coderdocs/components/Home.vue'),
    'topic': () => import('@/tpl-coderdocs/components/Topic.vue'),
    'document': () => import('@/tpl-coderdocs/components/Document.vue'),
})

export default router
