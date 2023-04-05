<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <div class="container px-5">
        <router-link :to="{name: 'Home'}" class="navbar-brand">
          <i v-if="siteMeta.icon!=''" :class="siteMeta.icon" style="padding-right: 4px"/>
          <span class="text-primary">{{ _siteNameFirst }}</span>{{ _siteNameLast }}
        </router-link>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation"><span class="navbar-toggler-icon"></span></button>
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
            <li class="nav-item">
              <router-link :to="{name: 'Home'}" class="nav-link">{{ $t('home') }}</router-link>
            </li>
            <li class="nav-item dropdown">
              <a class="nav-link dropdown-toggle active" id="navbarDropdownTopics" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">{{ $t('topics') }}</a>
              <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdownTopics">
                <li v-for="topic in topicList" v-bind:key="topic.id">
                  <router-link :to="{name: 'Topic', params: {tid: topic.id}}" class="dropdown-item">
                    <i v-if="topic.icon!=''" :class="topic.icon"/>
                    {{ _localedText(topic.title) }}
                  </router-link>
                </li>
              </ul>
            </li>
            <li class="nav-item dropdown">
              <a class="nav-link dropdown-toggle" id="navbarDropdownSettings" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false"><i class="bi bi-translate"></i></a>
              <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdownSettings">
                <li v-for="lang in _siteLanguages" v-bind:key="lang">
                  <a class="dropdown-item" href="#/!" @click="swichLanguage(lang, false)">
                    <span class="badge text-dark">{{ lang }}</span> {{siteMeta.languages[lang]}}
                  </a>
                </li>
              </ul>
            </li>
          </ul>
        </div>
      </div>
    </nav>

    <header class="bg-light border-bottom mb-4">
      <div class="container px-5">
        <nav aria-label="breadcrumb">
          <ol class="breadcrumb">
            <li class="breadcrumb-item"><router-link :to="{ name: 'Home' }">{{ $t('home') }}</router-link></li>
            <li class="breadcrumb-item active" aria-current="page">{{ $t('search') }}</li>
          </ol>
        </nav>

        <div class="text-center my-4">
          <h1 class="fw-bolder">{{ $t('search') }}</h1>
        </div>

        <div class="row justify-content-center">
          <div class="col-lg-6">
            <form class="form-inline" @submit.prevent="doSearch" method="get">
              <div class="input-group pb-4">
                <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control" v-model="searchQuery">
                <button type="submit" class="btn btn-primary" :value="$t('search')"><i class="bi bi-search" /></button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </header>

    <div class="container px-5">
      <div class="row">
        <div class="col-lg-8">
          <div v-if="searchHits.length==0" class="alert alert-warning" role="alert">{{ $t('search_no_result') }}</div>
          <div v-else class="card mb-4" v-for="document in searchHits" v-bind:key="document.id">
            <div class="card-body">
              <!--<div class="small text-muted">January 1, 2023</div>-->
              <h2 class="card-title">
                <router-link :to="{name: 'Document', params: {tid: document.topic.id, did: document.id}, query: {q: searchQuery}}" class="text-decoration-none">
                  <i v-if="document.icon" :class="document.icon"></i> {{ _localedText(document.title) }}
                </router-link>
              </h2>
              <p class="card-text">{{ _localedText(document.summary) }}</p>
              <router-link :to="{name: 'Document', params: {tid: document.topic.id, did: document.id}, query: {q: searchQuery}}" class="text-decoration-none">
                {{ $t('read') }} <i class="bi bi-arrow-right-circle"></i>
              </router-link>
            </div>
          </div>
        </div>

        <div class="col-lg-4">
          <div class="mb-4">
            <small>{{searchTotalResults}} results in {{searchDuration}} seconds</small>
          </div>

          <div class="card mb-4">
            <div class="card-header">{{ $t('topics') }}</div>
            <div class="card-body">
              <div class="row">
                <nav class="nav flex-column">
                  <template v-for="topic in topicList" v-bind:key="topic.id">
                    <router-link :to="{name: 'Topic', params: {tid: topic.id}}"
                                 :class="'nav-link'+($route.params.tid==topic.id?' active':'')" style="padding-top: 0px !important; padding-bottom: 0px !important;">
                      <i v-if="topic.icon!=''" :class="topic.icon"/> {{ _localedText(topic.title) }}
                    </router-link>
                  </template>
                </nav>
              </div>
            </div>
          </div>
          <!--          <div class="card mb-4">-->
          <!--            <div class="card-header">Side Widget</div>-->
          <!--            <div class="card-body">You can put anything you want inside of these side widgets. They are easy to use, and feature the Bootstrap 5 card component!</div>-->
          <!--          </div>-->
        </div>
      </div>
    </div>

    <footer class="py-5 bg-dark">
      <div class="container px-5">
        <p class="m-0 text-center text-white">
          Powered by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
          Copyright &copy; {{ siteMeta.name }}.
          <small>Build {{ siteMeta.tags.build }}</small>
        </p>
      </div>
    </footer>
  </div>
</template>

<script>
import {extractLeadingFromName, extractTrailingFromName} from "@/components/utils"
import i18n, {swichLanguage} from "@/i18n"
import {apiDoPost, apiDoGet, apiSite, apiSearch, apiTopics} from "@/utils/api_client";
import { watch } from 'vue'
import {useRoute} from "vue-router";

export default {
  name: 'Search',
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
              vue._fetchTopics(vue)
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          (err) => {
            vue.errorMsg = err
          })
    },
    _fetchTopics(vue) {
      vue.status = 0
      apiDoGet(apiTopics,
          (apiResp) => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.topicList = apiResp.data
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
      topicList: [],
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
