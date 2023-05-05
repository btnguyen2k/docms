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
            <p v-if="document.tags && $localedText(document.tags).length>0" style="font-size: small">
              <router-link v-for="tag in $localedText(document.tags)" v-bind:key="tag"
                           :to="{name: 'TagSearch', query:{q: tag, l: $i18n.locale}}"
                           :class="$calcTagCloudCSS(tag)+' me-1'" style="font-size: 0.65rem !important;">
                {{ tag }}
              </router-link>
            </p>
            <section class="docs-section docms-content img-fit img-center" v-html="documentContentRendered"></section>
          </article>

          <lego-page-footer/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
/* Lightbox for Bootstrap 5 */
import Lightbox from 'bs5-lightbox'

import {markdownRender} from "@/_shared/utils/docms_utils"
import {useRoute } from 'vue-router'
import {watch } from 'vue'
import '@/_shared/assets/markdown-gfm.css'
import {registerPopstate, unregisterPopstate} from '@/_shared/utils/docms_utils'
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import legoSidebar from './_sidebar.vue'
import {APP_CONFIG} from '@/_shared/utils/app_config'

const regTrailingSlash = /\/+$/

export default {
  name: 'Document',
  inject: ['$global', '$siteMeta', '$siteTopics', '$coderDocsResponsiveSidebar', '$calcTagCloudCSS'],
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
      this._updateLightbox()
      return markdownRender(this.$localedText(this.document.content), {
        sanitize: true,
        tags: this.$siteMeta.tags,
      })
    },
  },
  methods: {
    _updateLightbox() {
      this.$nextTick(() => {
        document.querySelectorAll('[data-toggle="lightbox"]').forEach(el => el.addEventListener('click', Lightbox.initialize));
      })
    },
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
              const appNameAndVersion = APP_CONFIG.app.name + ' v' + APP_CONFIG.app.version
              document.title = vue.$localedText(vue.document.title) + ' | ' + appNameAndVersion
            } else {
              // vue.errorMsg = vue.status+": "+apiResp.message
            }
            if (!vue.document.id) {
              vue.$transferToTopic(vue.topic.id, 3)
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
