<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">
    {{ errorMsg }}
    <hr/>
    <p class="btn btn-outline-primary mb-0" @click="$reload()"><i class="bi bi-arrow-clockwise"></i> {{ $t('reload') }}</p>
  </div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert"><i class="bi bi-hourglass"></i> {{ $t('wait') }}</div>
  <div v-else-if="!topic.id" class="alert alert-danger m-4" role="alert">
    {{ $t('error_topic_not_found', {topic: $route.params.tid}) }}
    <hr/>
    <span v-html="$t('transfer_to_home', {url: $router.resolve({name: 'Home'}).href})"></span>
  </div>
  <div v-else-if="!document.id" class="alert alert-danger m-4" role="alert">
    {{ $t('error_document_not_found', {topic: $route.params.tid, document: $route.params.did}) }}
    <hr/>
    <span v-html="$t('transfer_to_topic', {url: $router.resolve({name: 'Topic', params: {tid: topic.id}}).href})"></span>
  </div>
  <div v-else>
    <lego-page-header active="topic" :topic="topic" />

    <DocumentStyleContact v-if="document.style=='contact'" :topic="topic" :document="document" :document-list="documentList" />
    <DocumentStyleNormal v-else :topic="topic" :document="document" :document-list="documentList" />

    <lego-page-footer :document-list="latestDocuments" />
  </div>
</template>

<script>
import {useRoute} from 'vue-router'
import {watch} from 'vue'
import '@/_shared/assets/markdown-gfm.css'
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import {switchLanguage} from '@/_shared/i18n'
import DocumentStyleNormal from './_docStyleNormal.vue'
import DocumentStyleContact from './_docStyleContact.vue'

export default {
  name: 'Document',
  inject: ['$global', '$siteMeta', '$siteTopics', '$latestDocuments'],
  components: {legoPageHeader, legoPageFooter, DocumentStyleNormal, DocumentStyleContact},
  mounted() {
    const vue = this
    if (vue.$route.query.l) {
      switchLanguage(vue.$route.query.l, false)
    }
    const route = useRoute()
    watch(
        () => route.params.tid,
        async newTid => {
          vue._fetchTopics(vue, newTid)
        }
    )
    watch(
        () => route.params.did,
        async () => {
          vue._fetchTopics(vue, route.params.tid)
        }
    )
    this.$global.searchQuery = this.$route.query.q ? this.$route.query.q : ''
    this._fetchSiteMeta(this)
  },
  computed: {
    latestDocuments() {
      const result = []
      for (let i = 0; i < this.$latestDocuments.length; i++) {
        if (this.$latestDocuments[i].id != this.document.id) {
          result.push(this.$latestDocuments[i])
        }
      }
      return result
    },
  },
  methods: {
    _fetchSiteMeta(vue) {
      vue.$fetchSiteMeta(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue._fetchTopics(vue, vue.$route.params.tid)
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _fetchTopics(vue, topicId) {
      vue.$fetchSiteTopics(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.$siteTopics.forEach(t => {
                if (t.id == topicId) {
                  vue.topic = t
                  vue._fetchDocuments(vue, topicId, vue.$route.params.did)
                }
              })
            } else {
              // vue.errorMsg = vue.status+": "+apiResp.message
            }
            if (!vue.topic.id) {
              vue.$transferToHome(3)
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _fetchDocuments(vue, topicId, docId) {
      vue.$fetchDocumentsForTopic(topicId,
          () => vue.status=0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.documentList = apiResp.data
              vue.documentList.forEach(d => {
                if (d.id == docId) {
                  vue.$updatePageTitle({document: d})
                  vue.document = d
                  vue._fetchDocument(vue, docId)
                }
              })
            } else {
              // vue.errorMsg = vue.status+": "+apiResp.message
            }
            if (!vue.document.id) {
              vue.$transferToTopic(vue.topic.id, 3)
            }
          },
          err => vue.errorMsg = err
      )
    },
    _fetchDocument(vue, docId) {
      vue.$fetchDocument(vue.topic.id, docId,
          () => vue.status=0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.document = apiResp.data
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => vue.errorMsg = err
      )
    },
  },
  data() {
    return {
      topic: {},
      documentList: [],
      document: {},
      status: -1,
      errorMsg: '',
    }
  },
}
</script>
