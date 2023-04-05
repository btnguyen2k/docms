<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else-if="!document.id" class="alert alert-danger m-4" role="alert">{{ $t('error_document_not_found', {topic: $route.params.tid, document: $route.params.did}) }}</div>
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

    <div class="container px-5">
      <nav aria-label="breadcrumb">
        <ol class="breadcrumb">
          <li class="breadcrumb-item"><router-link :to="{ name: 'Home' }">{{ $t('home') }}</router-link></li>
          <li class="breadcrumb-item"><router-link :to="{ name: 'Topic', params: {tid: topic.id} }">{{ _localedText(topic.title) }}</router-link></li>
          <li class="breadcrumb-item active" aria-current="page">{{ _localedText(document.title) }}</li>
        </ol>
      </nav>

      <div class="row">
        <div class="col-lg-8">
          <article>
            <header class="mb-4">
              <h1 class="fw-bolder mb-1"><i v-if="document.icon" :class="document.icon"></i> {{ _localedText(document.title) }}</h1>
              <!--<div class="text-muted fst-italic mb-2">Posted on January 1, 2023 by Start Bootstrap</div>-->
<!--              <a class="badge bg-secondary text-decoration-none link-light" href="#!">Web Design</a>-->
<!--              <a class="badge bg-secondary text-decoration-none link-light" href="#!">Freebies</a>-->
            </header>
            <section class="mb-5">
              <div class="img-fit img-center" v-html="documentContentRendered"></div>
            </section>
          </article>
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
                    <nav v-if="topic.id==$route.params.tid" class="nav flex-column ps-4">
                      <router-link v-for="document in documentList" v-bind:key="document.id"
                                   :to="{name: 'Document', params: {tid: topic.id, did: document.id}}"
                                   :class="'nav-link'+($route.params.did==document.id?' active':'')" style="padding-top: 0px !important; padding-bottom: 0px !important;">
                        <i v-if="document.icon!=''" :class="document.icon"/> {{ _localedText(document.title) }}
                      </router-link>
                    </nav>
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
import { apiDoGet, apiSite, apiTopics, apiDocuments, apiDocument } from "@/utils/api_client"
import i18n, {swichLanguage} from "@/i18n"
import { styleByHash, extractLeadingFromName, extractTrailingFromName, markdownRender } from "./utils"
import { useRoute } from 'vue-router'
import { watch } from 'vue'
import './markdown-gfm.css'

export default {
  name: 'Document',
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
