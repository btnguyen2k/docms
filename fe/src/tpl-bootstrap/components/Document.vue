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
    <legoPageHeader />

    <div class="container px-5">
      <nav aria-label="breadcrumb">
        <ol class="breadcrumb">
          <li class="breadcrumb-item"><router-link :to="{ name: 'Home' }">{{ $t('home') }}</router-link></li>
          <li class="breadcrumb-item"><router-link :to="{ name: 'Topic', params: {tid: topic.id} }">{{ $localedText(topic.title) }}</router-link></li>
          <li class="breadcrumb-item active" aria-current="page">{{ $localedText(document.title) }}</li>
        </ol>
      </nav>

      <div class="row">
        <div class="col-lg-8">
          <article>
            <header class="mb-4">
              <h1 class="fw-bolder mb-1"><i v-if="document.icon" :class="document.icon"></i> {{ $localedText(document.title) }}</h1>
              <!--<div class="text-muted fst-italic mb-2">Posted on January 1, 2023 by Start Bootstrap</div>-->
              <p v-if="document.tags && $localedText(document.tags).length>0" style="font-size: small">
                <router-link v-for="tag in $localedText(document.tags)" v-bind:key="tag"
                             :to="{name: 'TagSearch', query:{q: tag, l: $i18n.locale}}"
                             class="badge bg-secondary text-decoration-none link-light me-1" style="font-size: 0.65rem !important;">
                  {{ tag }}
                </router-link>
              </p>
            </header>
            <section class="mb-5">
              <div class="docms-content img-fit img-center" v-html="documentContentRendered"></div>
            </section>
          </article>
        </div>

        <lego-sidebar :topic-id="topic.id" :document-list="documentList" :document-id="document.id"/>
      </div>
    </div>

    <lego-page-footer />
  </div>
</template>

<script>
/* Lightbox for Bootstrap 5 */
import Lightbox from 'bs5-lightbox'

import {markdownRender} from '@/_shared/utils/docms_utils'
import {useRoute} from 'vue-router'
import {watch} from 'vue'
import '@/_shared/assets/markdown-gfm.css'
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import legoSidebar from './_sidebar.vue'
import {APP_CONFIG} from '@/_shared/utils/app_config'

export default {
  name: 'Document',
  inject: ['$global', '$siteMeta', '$siteTopics'],
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
              const appNameAndVersion = APP_CONFIG.app.name + ' v' + APP_CONFIG.app.version
              document.title = vue.$localedText(vue.document.title) + ' | ' + appNameAndVersion
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
