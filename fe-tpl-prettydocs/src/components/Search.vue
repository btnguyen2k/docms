<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else class="body-orange">
    <div class="page-wrapper">
      <header id="header" class="header">
        <div class="container">
          <div class="branding">
            <h1 class="logo">
              <router-link :to="{ name: 'Home' }">
                <fa-icon v-if="siteMeta.icon!=''" aria-hidden="true" class="icon" style="padding-right: 8px" :icon="siteMeta.icon" />
                <span class="text-highlight">{{ _siteNameFirst }}</span><span class="text-bold">{{ _siteNameLast }}</span>
              </router-link>
            </h1>
          </div>

          <ol class="breadcrumb">
            <li class="breadcrumb-item"><router-link :to="{ name: 'Home' }">{{ $t('home') }}</router-link></li>
            <li class="breadcrumb-item active">{{ $t('search') }}</li>
          </ol>

          <div class="top-search-box">
            <form class="form-inline search-form justify-content-center" @submit.prevent="doSearch" method="get">
              <input type="text" :placeholder="$t('search')" name="q" class="form-control search-input" v-model="searchQuery">
              <button type="submit" class="btn search-btn" :value="$t('search')"><fa-icon icon="fas fa-search"></fa-icon></button>
            </form>
          </div>
        </div>
      </header>

      <div class="doc-wrapper">
        <div class="container">
          <div id="doc-header" class="doc-header text-center">
            <h1 class="doc-title">{{$t('search_result')}}</h1>
          </div>
          <div class="doc-body row">
            <div class="doc-content col-md-9 col-12 order-1">
              <div class="content-inner">
                <section class="doc-section">
                  <div v-if="searchHits.length==0" class="alert alert-secondary" role="alert">{{ $t('search_no_result') }}</div>
                  <div v-else v-for="doc in searchHits" v-bind:key="doc.id" class="section-block">
                    <router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}, query: {q: searchTerm}}" class="nav-link">
                      <h3 class="block-title"><fa-icon v-if="doc.icon!=''" :icon="doc.icon" class="pe-1"/>{{ _localedText(doc.title) }}</h3>
                    </router-link>
                    <p style="cursor: pointer" @click="doViewDocument(doc.topic.id, doc.id)">{{ _localedText(doc.summary) }}</p>
                  </div>
                  <div class="section-block">
                    <small>{{searchTotalResults}} results in {{searchDuration}} seconds</small>
                  </div>
                </section>
              </div>
            </div>
<!--            <div class="doc-sidebar col-md-3 col-12 order-0 d-none d-md-flex">-->
<!--              <div id="doc-nav" class="doc-nav">-->
<!--                <nav id="doc-menu" class="nav doc-menu flex-column sticky">-->
<!--                  <li v-for="topic in topicList" v-bind:key="topic.id" :class="'nav-item'+($route.params.tid==topic.id?' active':'')">-->
<!--                    <router-link class="nav-link scrollto" :to="{name: 'Topic', params: {tid: topic.id}}">-->
<!--                      <fa-icon v-if="topic.icon!=''" :icon="topic.icon" class="pe-1"/>{{ _localedText(topic.title) }}-->
<!--                    </router-link>-->
<!--                  </li>-->
<!--                </nav>-->
<!--              </div>-->
<!--            </div>-->
          </div>
        </div>
      </div>
    </div>
    <footer class="footer text-center">
      <div class="container">
        <!--/* This template is free as long as you keep the footer attribution link. If you'd like to use the template without the attribution link, you can buy the commercial license via our website: themes.3rdwavemedia.com Thank you for your support. :) */-->
        <small class="copyright d-none d-lg-inline float-none">
          Powered by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
          Theme <span style="font-family: monospace, monospace;">PrettyDocs</span> designed with <span class="sr-only">love</span><fa-icon icon="fas fa-heart"/> by <a href="https://themes.3rdwavemedia.com/" target="_blank">Xiaoying Riley</a> for developers.
        </small>
        <small class="copyright d-lg-none float-none">
          <fa-icon icon="fas fa-bolt-lightning"></fa-icon> by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
          <fa-icon icon="fas fa-pen-ruler"></fa-icon> <span style="font-family: monospace, monospace;">PrettyDocs</span> by <a href="https://themes.3rdwavemedia.com/" target="_blank">Xiaoying Riley</a>.
        </small>
        <small v-if="siteMeta.tags.build" class="d-none d-lg-inline float-end">Build: {{ siteMeta.tags.build }}</small>
        <ul class="nav nav-pills float-start align-middle" style="margin-top: -6px !important;"><!--style="font-size: 0.85em !important; "-->
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" data-bs-toggle="dropdown" href="#" role="button" aria-expanded="false"><fa-icon icon="fas fa-gears"></fa-icon></a>
            <ul class="dropdown-menu" style="font-size: small">
              <li v-for="lang in _siteLanguages" v-bind:key="lang">
                <a class="dropdown-item" href="#" @click="swichLanguage(lang, false)">{{siteMeta.languages[lang]}}</a>
              </li>
              <li class="d-lg-none"><hr class="dropdown-divider"></li>
              <li v-if="siteMeta.tags.build" class="dropdown-item d-lg-none">Build: {{ siteMeta.tags.build }}</li>
            </ul>
          </li>
        </ul>
      </div>
    </footer>
  </div>
</template>

<script>
import {extractLeadingFromName, extractTrailingFromName} from "@/components/utils"
import i18n, {swichLanguage} from "@/i18n"
import {apiDoPost, apiDoGet, apiSite, apiSearch} from "@/utils/api_client";
import { watch } from 'vue'
import {dropdown} from "bootstrap"
import {useRoute} from "vue-router";

export default {
  name: 'Search',
  components: [dropdown],
  mounted() {
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
    this.searchQuery = this.searchTerm
    this._fetchSiteMeta(this)
  },
  computed: {
    _siteNameFirst() {
      return extractLeadingFromName(this.siteMeta.name)
    },
    _siteNameLast() {
      return extractTrailingFromName(this.siteMeta.name)
    },
    _siteLanguages() {
      return Object.keys(this.siteMeta.languages).filter(value => {return value != 'default'})
    },
    searchTerm() {
      return this.$route.query.q ? this.$route.query.q.trim() : ""
    },
    searchLocale() {
      return this.$route.query.l ? this.$route.query.l.trim() : ""
    },
  },
  methods: {
    swichLanguage,
    _localedText(inputMap) {
      return inputMap[i18n.global.locale]?inputMap[i18n.global.locale]:inputMap
    },
    _fetchSiteMeta(vue) {
      vue.status = 0
      apiDoGet(apiSite,
          (apiResp) => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.siteMeta = apiResp.data
              vue._search(vue)
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          (err) => {
            vue.errorMsg = err
          })
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
          },
          (err) => {
            vue.errorMsg = err
          })
    },
    doSearch() {
      this.$router.push({name: 'Search', query: {q: this.searchQuery, l: i18n.global.locale}})
    },
    doViewDocument(tid, did) {
      this.$router.push({name: 'Document', params: {tid: tid, did: did}, query: {q: this.searchQuery}})
    },
  },
  data() {
    return {
      siteMeta: {},
      // topicList: [],
      // topic: {},
      // documentList: [],
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
