<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else-if="!topic.id" class="alert alert-danger m-4" role="alert">{{ $t('error_topic_not_found', {topic: $route.params.tid}) }}</div>
  <div v-else class="docs-page">
    <lego-page-header />

    <div class="docs-wrapper">
      <lego-sidebar :topic-id="topic.id" :document-list="documentList" />

      <div class="docs-content">
        <div class="container">
          <div v-if="documentList.length==0" class="alert alert-warning mt-4" role="alert">{{ $t('empty_topic') }}</div>
          <article v-for="document in documentList" v-bind:key="document.id" class="docs-article pt-1 pb-2">
<!--            <header class="docs-header">-->
<!--              <router-link :to="{name: 'Document', params: {tid: topic.id, did: document.id}}" class="text-decoration-none">-->
<!--                <h1 class="docs-heading mb-2"><i v-if="document.icon!=''" :class="document.icon"></i> {{ $localedText(document.title) }}</h1>-->
<!--              </router-link>-->
<!--              <section class="docs-intro">-->
<!--                <p>{{ $localedText(document.summary) }}</p>-->
<!--              </section>-->
<!--            </header>-->
            <section class="docs-section">
              <router-link :to="{name: 'Document', params: {tid: topic.id, did: document.id}}" class="text-decoration-none">
                <h2 class="section-heading"><i v-if="document.icon!=''" :class="document.icon"></i> {{ $localedText(document.title) }}</h2>
              </router-link>
              <p>{{ $localedText(document.summary) }}</p>
            </section>
          </article>

          <lego-page-footer/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { useRoute } from 'vue-router'
import { watch } from 'vue'
import {registerPopstate, unregisterPopstate} from "@/components/utils"
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import legoSidebar from './_sidebar.vue'

const regTrailingSlash = /\/+$/

export default {
  name: 'Topic',
  inject: ['$global', '$siteMeta', '$siteLanguages', '$siteTopics', '$coderDocsResponsiveSidebar'],
  components: {legoPageHeader, legoPageFooter, legoSidebar},
  unmounted() {
    unregisterPopstate(this.handleBackFoward)
  },
  mounted() {
    registerPopstate(this.handleBackFoward)

    const vue = this
    const route = useRoute()
    watch(
        () => route.params.tid,
        async newTid => {
          if (newTid) {
            vue._fetchTopics(vue, newTid)
          }
        }
    )
    this._fetchSiteMeta(this)
  },
  methods: {
    handleBackFoward() {
      const pathBase = this.$router.options.meta.base.replace(regTrailingSlash, '')
      const vuePath = window.location.pathname.slice(pathBase.length) // remove the 'base' prefix
      const result = this.$router.resolve(vuePath)
      if (result && result.resolved && result.resolved.name=='Topic') {
        this._fetchTopics(this, this.$route.params.tid)
      }
    },
    _fetchSiteMeta(vue) {
      vue.$fetchSiteMeta(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.$global.siteMeta = apiResp.data
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
              vue.$global.siteTopics = apiResp.data
              vue.$global.siteTopics.forEach(t => {
                if (t.id == topicId) {
                  vue.topic = t
                  vue._fetchDocuments(vue)
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
    _fetchDocuments(vue) {
      vue.$fetchDocumentsForTopic(vue.topic.id,
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.documentList = apiResp.data
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
            vue.$nextTick(()=>{
              // CoderDocs: onload
              this.$coderDocsResponsiveSidebar()
            })
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
  },
  data() {
    return {
      topic: {},
      documentList: [],
      status: -1,
      errorMsg: '',
    }
  },
}
</script>
