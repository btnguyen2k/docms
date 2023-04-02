<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else-if="!document.id" class="alert alert-danger m-4" role="alert">{{ $t('error_document_not_found', {topic: $route.params.tid, document: $route.params.did}) }}</div>
  <div v-else :class="_styleClassForTopic(topic)">
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
            <li class="breadcrumb-item"><router-link :to="{name: 'Home'}">{{ $t('home') }}</router-link></li>
            <li class="breadcrumb-item"><router-link :to="{name: 'Topic', params: {tid: topic.id}}">{{ _localedText(topic.title) }}</router-link></li>
            <li class="breadcrumb-item active">{{ _localedText(document.title) }}</li>
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
            <h1 class="doc-title"><fa-icon v-if="document.icon" :icon="document.icon"></fa-icon> {{ _localedText(document.title) }}</h1>
            <!--            <div class="meta"><i class="far fa-clock"></i> Last updated: June 13th, 2022</div>-->
          </div>
          <div class="doc-body row">
            <div class="doc-content col-md-9 col-12 order-1">
              <div class="content-inner">
                <section class="doc-section">
                  <div class="section-block img-fit img-center" v-html="documentContentRendered"></div>
                </section>
              </div>
            </div>
            <div class="doc-sidebar col-md-3 col-12 order-0 d-none d-md-flex">
              <div id="doc-nav" class="doc-nav">
                <nav id="doc-menu" class="nav doc-menu flex-column sticky">
                  <template v-for="topic in topicList" v-bind:key="topic.id">
                    <li :class="'nav-item'+($route.params.tid==topic.id?' active':'')">
                      <router-link class="nav-link scrollto" :to="{name: 'Topic', params: {tid: topic.id}}">
                        <fa-icon v-if="topic.icon!=''" :icon="topic.icon" class="pe-1"/>{{ _localedText(topic.title) }}
                      </router-link>
                    </li>
                    <nav v-if="topic.id==$route.params.tid" class="nav doc-sub-menu nav flex-column">
                      <li v-for="doc in documentList" v-bind:key="doc.id" :class="'nav-item'+($route.params.did==doc.id?' active':'')">
                        <router-link class="nav-link scrollto" :to="{name: 'Document', params: {tid: topic.id, did: doc.id}}">
                          <fa-icon v-if="doc.icon!=''" :icon="doc.icon" class="pe-1"/>{{ _localedText(doc.title) }}
                        </router-link>
                      </li>
                    </nav>
                  </template>
                </nav>
              </div>
            </div>
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
import { apiDoGet, apiSite, apiTopics, apiDocuments, apiDocument } from "@/utils/api_client"
import i18n, {swichLanguage} from "@/i18n"
import { styleByHash, extractLeadingFromName, extractTrailingFromName, markdownRender } from "./utils"
import { useRoute } from 'vue-router'
import { watch } from 'vue'
import './markdown-gfm.css'
import {dropdown} from "bootstrap"

export default {
  name: 'Document',
  components: [dropdown],
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
    this.searchQuery = this.$route.query.q ? this.$route.query.q : ''
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
    documentContentRendered() {
      return markdownRender(this._localedText(this.document.content), true)
    },
  },
  methods: {
    swichLanguage,
    _styleClassForTopic(topic) {
      const styleList = ["body-blue", "body-green", "body-red", "body-pink", "body-purple", "body-orange"]
      return styleByHash(topic.id, styleList)
    },
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
              vue._fetchTopics(vue, vue.$route.params.tid)
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          (err) => {
            vue.errorMsg = err
          })
    },
    _fetchTopics(vue, topicId) {
      vue.status = 0
      apiDoGet(apiTopics,
          (apiResp) => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.topicList = apiResp.data
              vue.topicList.forEach(t => {
                if (t.id == topicId) {
                  vue.topic = t
                  vue._fetchDocuments(vue, vue.$route.params.did)
                }
              })
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          (err) => {
            vue.errorMsg = err
          })
    },
    _fetchDocuments(vue, docId) {
      vue.status = 0
      apiDoGet(apiDocuments.replaceAll(':topic-id', vue.topic.id),
          (apiResp) => {
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
          (err) => {
            vue.errorMsg = err
          })
    },
    _fetchDocument(vue, docId) {
      vue.status = 0
      apiDoGet(apiDocument.replaceAll(':topic-id', vue.topic.id).replaceAll(':document-id', docId),
          (apiResp) => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.document = apiResp.data
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          (err) => {
            vue.errorMsg = err
          })
    },
    doViewDocument(tid, did) {
      this.$router.push({name: 'Document', params: {tid: tid, did: did}})
    },
    doSearch() {
      this.$router.push({name: 'Search', query: {q: this.searchQuery, l: i18n.global.locale}})
    },
  },
  data() {
    return {
      siteMeta: {},
      topicList: [],
      topic: {},
      documentList: [],
      document: {},
      status: -1,
      errorMsg: '',
      searchQuery: '',
    }
  },
}
</script>
