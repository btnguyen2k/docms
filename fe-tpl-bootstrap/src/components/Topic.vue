<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else-if="!topic.id" class="alert alert-danger m-4" role="alert">{{ $t('error_topic_not_found', {topic: $route.params.tid}) }}</div>
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
            <li class="breadcrumb-item active" aria-current="page">{{ _localedText(topic.title) }}</li>
          </ol>
        </nav>

        <div class="text-center my-4">
          <h1 class="fw-bolder"><i v-if="topic.icon" :class="topic.icon"></i> {{ _localedText(topic.title) }}</h1>
          <p class="lead mb-0">{{ _localedText(topic.description) }}</p>
        </div>

<!--        <div class="row justify-content-center">-->
<!--          <div class="col-lg-6">-->
<!--            <form class="form-inline" @submit.prevent="doSearch" method="get">-->
<!--              <div class="input-group pb-4">-->
<!--                <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control" v-model="searchQuery">-->
<!--                <button type="submit" class="btn btn-primary" :value="$t('search')"><i class="bi bi-search" /></button>-->
<!--              </div>-->
<!--            </form>-->
<!--          </div>-->
<!--        </div>-->
      </div>
    </header>

    <div class="container px-5">
      <div class="row">
        <div class="col-lg-8">
          <div v-if="documentList.length==0" class="alert alert-warning" role="alert">{{ $t('empty_topic') }}</div>
          <div v-else class="card mb-4" v-for="document in documentList.slice(0)" v-bind:key="document.id">
            <div class="card-body">
              <!--<div class="small text-muted">January 1, 2023</div>-->
              <h2 class="card-title">
                <router-link :to="{name: 'Document', params: {tid: topic.id, did: document.id}}" class="text-decoration-none">
                  <i v-if="document.icon" :class="document.icon"></i> {{ _localedText(document.title) }}
                </router-link>
              </h2>
              <p class="card-text">{{ _localedText(document.summary) }}</p>
              <router-link :to="{name: 'Document', params: {tid: topic.id, did: document.id}}" class="text-decoration-none">
                {{ $t('read') }} <i class="bi bi-arrow-right-circle"></i>
              </router-link>
            </div>
          </div>
<!--          <div class="row" v-if="false">-->
<!--            <div v-for="document in documentList.slice(1)" v-bind:key="document.id" class="col-lg-6">-->
<!--              <div class="card mb-4">-->
<!--                <div class="card-body">-->
<!--                  &lt;!&ndash;<div class="small text-muted">January 1, 2023</div>&ndash;&gt;-->
<!--                  <h2 class="card-title h4">-->
<!--                    <router-link :to="{name: 'Document', params: {tid: topic.id, did: document.id}}" class="text-decoration-none">-->
<!--                      <i v-if="document.icon" :class="document.icon"></i> {{ _localedText(document.title) }}-->
<!--                    </router-link>-->
<!--                  </h2>-->
<!--                  <p class="card-text">{{ _localedText(document.summary) }}</p>-->
<!--                  <router-link :to="{name: 'Document', params: {tid: topic.id, did: document.id}}" class="text-decoration-none">-->
<!--                    {{ $t('read') }} <i class="bi bi-arrow-right-circle"></i>-->
<!--                  </router-link>-->
<!--                </div>-->
<!--              </div>-->
<!--            </div>-->
<!--          </div>-->
        </div>

        <div class="col-lg-4">
          <form class="form-inline" @submit.prevent="doSearch" method="get">
            <div class="input-group pb-4">
              <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control" v-model="searchQuery">
              <button type="submit" class="btn btn-primary" :value="$t('search')"><i class="bi bi-search" /></button>
            </div>
          </form>

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
import { apiDoGet, apiSite, apiTopics, apiDocuments } from "@/utils/api_client"
import i18n, {swichLanguage} from "@/i18n"
import { styleByHash, extractLeadingFromName, extractTrailingFromName } from "./utils"
import { useRoute } from 'vue-router'
import { watch } from 'vue'

export default {
  name: 'Topic',
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
                  vue._fetchDocuments(vue)
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
    _fetchDocuments(vue) {
      vue.status = 0
      apiDoGet(apiDocuments.replaceAll(':topic-id', vue.topic.id),
          (apiResp) => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.documentList = apiResp.data
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
      status: -1,
      errorMsg: '',
      searchQuery: '',
    }
  },
}
</script>
