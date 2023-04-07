<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else class="docs-page">
    <lego-page-header />

    <div class="docs-wrapper">
      <lego-sidebar />

      <div class="docs-content">
        <div class="container">
          <div v-if="searchHits.length==0" class="alert alert-warning mt-4" role="alert">{{ $t('search_no_result') }}</div>
          <article v-for="document in searchHits" v-bind:key="document.id" class="docs-article pt-1 pb-2">
<!--            <header class="docs-header">-->
<!--              <router-link :to="{name: 'Document', params: {tid: document.topic.id, did: document.id}}" class="text-decoration-none">-->
<!--                <h1 class="docs-heading mb-2"><i v-if="document.icon!=''" :class="document.icon"></i> {{ $localedText(document.title) }}</h1>-->
<!--              </router-link>-->
<!--              <section class="docs-intro">-->
<!--                <p>{{ $localedText(document.summary) }}</p>-->
<!--              </section>-->
<!--            </header>-->
            <section class="docs-section">
              <router-link :to="{name: 'Document', params: {tid: document.topic.id, did: document.id}}" class="text-decoration-none">
                <h2 class="section-heading"><i v-if="document.icon!=''" :class="document.icon"></i> {{ $localedText(document.title) }}</h2>
              </router-link>
              <p>{{ $localedText(document.summary) }}</p>
            </section>
          </article>

          <small>{{searchTotalResults}} results in {{searchDuration}} seconds</small>

          <lego-page-footer />
        </div>
      </div>
    </div>
  </div>

<!--  <div v-else class="body-orange">-->
<!--    <div class="page-wrapper">-->

<!--      <div class="doc-wrapper">-->
<!--        <div class="container">-->
<!--          <div id="doc-header" class="doc-header text-center">-->
<!--            <h1 class="doc-title">{{$t('search_result')}}</h1>-->
<!--          </div>-->
<!--          <div class="doc-body row">-->
<!--            <div class="doc-content col-md-9 col-12 order-1">-->
<!--              <div class="content-inner">-->
<!--                <section class="doc-section">-->
<!--                  <div v-if="searchHits.length==0" class="alert alert-secondary" role="alert">{{ $t('search_no_result') }}</div>-->
<!--                  <div v-else v-for="doc in searchHits" v-bind:key="doc.id" class="section-block">-->
<!--                    <router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}, query: {q: searchTerm}}" class="nav-link">-->
<!--                      <h3 class="block-title"><fa-icon v-if="doc.icon!=''" :icon="doc.icon" class="pe-1"/>{{ _localedText(doc.title) }}</h3>-->
<!--                    </router-link>-->
<!--                    <p style="cursor: pointer" @click="doViewDocument(doc.topic.id, doc.id)">{{ _localedText(doc.summary) }}</p>-->
<!--                  </div>-->
<!--                  <div class="section-block">-->
<!--                    <small>{{searchTotalResults}} results in {{searchDuration}} seconds</small>-->
<!--                  </div>-->
<!--                </section>-->
<!--              </div>-->
<!--            </div>-->
<!--&lt;!&ndash;            <div class="doc-sidebar col-md-3 col-12 order-0 d-none d-md-flex">&ndash;&gt;-->
<!--&lt;!&ndash;              <div id="doc-nav" class="doc-nav">&ndash;&gt;-->
<!--&lt;!&ndash;                <nav id="doc-menu" class="nav doc-menu flex-column sticky">&ndash;&gt;-->
<!--&lt;!&ndash;                  <li v-for="topic in topicList" v-bind:key="topic.id" :class="'nav-item'+($route.params.tid==topic.id?' active':'')">&ndash;&gt;-->
<!--&lt;!&ndash;                    <router-link class="nav-link scrollto" :to="{name: 'Topic', params: {tid: topic.id}}">&ndash;&gt;-->
<!--&lt;!&ndash;                      <fa-icon v-if="topic.icon!=''" :icon="topic.icon" class="pe-1"/>{{ _localedText(topic.title) }}&ndash;&gt;-->
<!--&lt;!&ndash;                    </router-link>&ndash;&gt;-->
<!--&lt;!&ndash;                  </li>&ndash;&gt;-->
<!--&lt;!&ndash;                </nav>&ndash;&gt;-->
<!--&lt;!&ndash;              </div>&ndash;&gt;-->
<!--&lt;!&ndash;            </div>&ndash;&gt;-->
<!--          </div>-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->
<!--    <footer class="footer text-center">-->
<!--      <div class="container">-->
<!--        &lt;!&ndash;/* This template is free as long as you keep the footer attribution link. If you'd like to use the template without the attribution link, you can buy the commercial license via our website: themes.3rdwavemedia.com Thank you for your support. :) */&ndash;&gt;-->
<!--        <small class="copyright d-none d-lg-inline float-none">-->
<!--          Powered by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.-->
<!--          Theme <span style="font-family: monospace, monospace;">PrettyDocs</span> designed with <span class="sr-only">love</span><fa-icon icon="fas fa-heart"/> by <a href="https://themes.3rdwavemedia.com/" target="_blank">Xiaoying Riley</a> for developers.-->
<!--        </small>-->
<!--        <small class="copyright d-lg-none float-none">-->
<!--          <fa-icon icon="fas fa-bolt-lightning"></fa-icon> by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.-->
<!--          <fa-icon icon="fas fa-pen-ruler"></fa-icon> <span style="font-family: monospace, monospace;">PrettyDocs</span> by <a href="https://themes.3rdwavemedia.com/" target="_blank">Xiaoying Riley</a>.-->
<!--        </small>-->
<!--        <small v-if="siteMeta.tags.build" class="d-none d-lg-inline float-end">Build: {{ siteMeta.tags.build }}</small>-->
<!--        <ul class="nav nav-pills float-start align-middle" style="margin-top: -6px !important;">&lt;!&ndash;style="font-size: 0.85em !important; "&ndash;&gt;-->
<!--          <li class="nav-item dropdown">-->
<!--            <a class="nav-link dropdown-toggle" data-bs-toggle="dropdown" href="#" role="button" aria-expanded="false"><fa-icon icon="fas fa-gears"></fa-icon></a>-->
<!--            <ul class="dropdown-menu" style="font-size: small">-->
<!--              <li v-for="lang in _siteLanguages" v-bind:key="lang">-->
<!--                <a class="dropdown-item" href="#" @click="swichLanguage(lang, false)">{{siteMeta.languages[lang]}}</a>-->
<!--              </li>-->
<!--              <li class="d-lg-none"><hr class="dropdown-divider"></li>-->
<!--              <li v-if="siteMeta.tags.build" class="dropdown-item d-lg-none">Build: {{ siteMeta.tags.build }}</li>-->
<!--            </ul>-->
<!--          </li>-->
<!--        </ul>-->
<!--      </div>-->
<!--    </footer>-->
<!--  </div>-->
</template>

<script>
import {registerPopstate, unregisterPopstate} from "@/components/utils"
import {swichLanguage} from "@/i18n"
import {apiDoPost, apiSearch} from "@/utils/api_client"
import { watch } from 'vue'
import {useRoute} from "vue-router"
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from "@/components/_pageFooter.vue";
import legoSidebar from "@/components/_sidebar.vue";

const regTrailingSlash = /\/+$/

export default {
  name: 'Search',
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
        () => route.query.q,
        async () => vue._search(vue),
    )
    watch(
        () => route.query.l,
        async () => vue._search(vue),
    )
    swichLanguage(this.searchLocale)
    this.$global.searchQuery = this.searchTerm
    this._fetchSiteMeta(this)
  },
  computed: {
    searchTerm() {
      return this.$route.query.q ? this.$route.query.q.trim() : ""
    },
    searchLocale() {
      return this.$route.query.l ? this.$route.query.l.trim() : ""
    },
  },
  methods: {
    handleBackFoward() {
      const pathBase = this.$router.options.meta.base.replace(regTrailingSlash, '')
      const vuePath = window.location.pathname.slice(pathBase.length) // remove the 'base' prefix
      const result = this.$router.resolve(vuePath)
      if (result && result.resolved && result.resolved.name=='Search') {
        this._search(this)
      }
    },
    _fetchSiteMeta(vue) {
      vue.$fetchSiteMeta(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.$global.siteMeta = apiResp.data
              vue._fetchTopics(vue)
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _fetchTopics(vue) {
      vue.$fetchSiteTopics(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.$global.siteTopics = apiResp.data
              vue._search(this)
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _search(vue) {
      apiDoPost(apiSearch, {query: this.searchTerm},
          (apiResp) => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.searchHits = apiResp.data.docs
              vue.searchDuration = apiResp.data.duration
              vue.searchTotalResults = apiResp.data.total
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
            vue.$nextTick(()=>{
              // CoderDocs: onload
              vue.$coderDocsResponsiveSidebar()
            })
          },
          (err) => {
            vue.errorMsg = err
          })
    },
  },
  data() {
    return {
      searchHits: [],
      status: -1,
      errorMsg: '',
      searchQuery: '',
      searchDuration: 0,
      searchTotalResults: 0,
    }
  },
}
</script>
