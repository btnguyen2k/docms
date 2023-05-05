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
  <div v-else :class="_styleClassForTopic(topic)">
    <div class="page-wrapper">
      <lego-page-header :topic="topic"/>

      <div class="doc-wrapper">
        <div class="container">
          <div id="doc-header" class="doc-header text-center">
            <h1 class="doc-title"><i v-if="topic.icon" :class="topic.icon"></i> {{ $localedText(topic.title) }}</h1>
            <!--<div class="meta"><i class="far fa-clock"></i> Last updated: June 13th, 2022</div>-->
          </div>
          <div class="doc-body row">
            <div class="doc-content col-md-9 col-12 order-1">
              <div class="content-inner">
                <section class="doc-section">
                  <div class="section-block">
                    <p>{{ $localedText(topic.description) }}</p>
                  </div>

                  <div v-if="documentList.length==0" class="alert alert-warning" role="alert">{{ $t('empty_topic') }}</div>
                  <div v-else v-for="document in documentList" v-bind:key="document.id" class="section-block">
                    <router-link :to="{name: 'Document', params: {tid: topic.id, did: document.id}}" class="nav-link">
                      <h3 class="block-title"><i v-if="document.icon!=''" :class="document.icon" class="pe-1"/>{{ $localedText(document.title) }}</h3>
                    </router-link>
                    <router-link :to="{name: 'Document', params: {tid: topic.id, did: document.id}}" class="text-decoration-none text-muted">
                      {{ $localedText(document.summary) }}
                    </router-link>
                    <p v-if="document.tags && $localedText(document.tags).length>0" style="font-size: small">
                      <router-link v-for="tag in $localedText(document.tags)" v-bind:key="tag"
                                   :to="{name: 'TagSearch', query:{q: tag, l: $i18n.locale}}"
                                   class="badge bg-secondary text-decoration-none link-light me-1" style="font-size: 0.65rem !important;">
                        {{ tag }}
                      </router-link>
                    </p>
                  </div>
                </section>
              </div>
            </div>

            <lego-sidebar :topic-id="topic.id" :document-list="documentList" />
          </div>
        </div>
      </div>
    </div>

    <lego-page-footer />
  </div>
</template>

<script>
import {useRoute} from 'vue-router'
import {watch} from 'vue'
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import legoSidebar from './_sidebar.vue'
import {APP_CONFIG} from '@/_shared/utils/app_config'

export default {
  name: 'Topic',
  inject: ['$siteTopics'],
  components: {legoPageHeader, legoPageFooter, legoSidebar},
  mounted() {
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
                  const appNameAndVersion = APP_CONFIG.app.name + ' v' + APP_CONFIG.app.version
                  document.title = vue.$localedText(vue.topic.title) + ' | ' + appNameAndVersion
                  vue._fetchDocuments(vue)
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
