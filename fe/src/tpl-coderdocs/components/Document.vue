<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else-if="!document.id" class="alert alert-danger m-4" role="alert">{{ $t('error_document_not_found', {topic: $route.params.tid, document: $route.params.did}) }}</div>
  <div v-else class="docs-page">
    <lego-page-header />

    <div class="docs-wrapper">
      <lego-sidebar :topic-id="topic.id" :document-list="documentList" :document-id="document.id"/>

      <div class="docs-content">
        <div class="container">
          <div v-if="documentList.length==0" class="alert alert-warning mt-4" role="alert">{{ $t('empty_topic') }}</div>
          <article class="docs-article">
            <header class="docs-header">
              <h1 class="docs-heading mb-2"><i v-if="document.icon!=''" :class="document.icon"></i> {{ $localedText(document.title) }}</h1>
            </header>
            <section class="docs-section img-fit img-center" v-html="documentContentRendered"></section>
          </article>

          <lego-page-footer/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {markdownRender} from "@/_shared/utils/docms_utils"
import {useRoute } from 'vue-router'
import {watch } from 'vue'
import '@/_shared/assets/markdown-gfm.css'
import {registerPopstate, unregisterPopstate} from '@/_shared/utils/docms_utils'
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import legoSidebar from './_sidebar.vue'

const regTrailingSlash = /\/+$/

export default {
  name: 'Document',
  inject: ['$global', '$siteMeta', '$siteTopics', '$coderDocsResponsiveSidebar'],
  components: {legoPageHeader, legoPageFooter, legoSidebar},
  unmounted() {
    unregisterPopstate(this.handleBackFoward)
  },
  mounted() {
    registerPopstate(this.handleBackFoward)

    const vue = this
    const route = useRoute()
    watch(
        () => route.params.did,
        async newDid => {
          if (newDid) {
            vue._fetchDocument(vue, newDid)
          }
        }
    )
    this.$global.searchQuery = this.$route.query.q ? this.$route.query.q : ''
    this._fetchSiteMeta(this)
  },
  computed: {
    documentContentRendered() {
      return markdownRender(this.$localedText(this.document.content), {
        sanitize: true,
        tags: this.$siteMeta.tags,
      })
    },
  },
  methods: {
    handleBackFoward() {
      const pathBase = this.$router.options.meta.base.replace(regTrailingSlash, '')
      const vuePath = window.location.pathname.slice(pathBase.length) // remove the 'base' prefix
      const result = this.$router.resolve(vuePath)
      if (result && result.resolved && result.resolved.name=='Document') {
        this._fetchDocument(this, this.$route.params.did)
      }
    },
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
                  vue._fetchDocuments(vue, vue.$route.params.did)
                }
              })
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _fetchDocuments(vue, docId) {
      vue.$fetchDocumentsForTopic(vue.topic.id,
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.documentList = apiResp.data
              vue.documentList.forEach(d => {
                if (d.id == docId) {
                  vue._fetchDocument(vue, docId)
                }
              })
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _fetchDocument(vue, docId) {
      vue.$fetchDocument(vue.topic.id, docId,
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.document = apiResp.data
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
            vue.$nextTick(()=>{
              // CoderDocs: onload
              vue.$coderDocsResponsiveSidebar()
            })
          },
          err => {
            vue.errorMsg = err
          }
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
