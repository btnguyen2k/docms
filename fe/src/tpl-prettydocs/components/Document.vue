<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else-if="!document.id" class="alert alert-danger m-4" role="alert">{{ $t('error_document_not_found', {topic: $route.params.tid, document: $route.params.did}) }}</div>
  <div v-else :class="_styleClassForTopic(topic)">
    <div class="page-wrapper">
      <lego-page-header :topic="topic" :document="document" />

      <div class="doc-wrapper">
        <div class="container">
          <div id="doc-header" class="doc-header text-center">
            <h1 class="doc-title"><i v-if="document.icon" :class="document.icon"></i> {{ $localedText(document.title) }}</h1>
            <!--<div class="meta"><i class="far fa-clock"></i> Last updated: June 13th, 2022</div>-->
          </div>
          <div class="doc-body row">
            <div class="doc-content col-md-9 col-12 order-1">
              <div class="content-inner">
                <section class="doc-section">
                  <div class="section-block img-fit img-center" v-html="documentContentRendered"></div>
                </section>
              </div>
            </div>

            <lego-sidebar :topic-id="topic.id" :document-list="documentList" :document-id="document.id" />
          </div>
        </div>
      </div>
    </div>

    <lego-page-footer />
  </div>
</template>

<script>
import {markdownRender} from '@/_shared/utils/docms_utils'
import {useRoute} from 'vue-router'
import {watch} from 'vue'
import '@/_shared/assets/markdown-gfm.css'
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import legoSidebar from './_sidebar.vue'

export default {
  name: 'Document',
  inject: ['$global', '$siteTopics'],
  components: {legoPageHeader, legoPageFooter, legoSidebar},
  mounted() {
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
      return markdownRender(this.$localedText(this.document.content), true)
    },
  },
  methods: {
    _styleClassForTopic(topic) {
      const styleList = ["body-blue", "body-green", "body-red", "body-pink", "body-purple", "body-orange"]
      return this.$styleByHash(topic.id, styleList)
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
