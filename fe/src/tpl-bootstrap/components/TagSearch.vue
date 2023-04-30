<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else>
    <legoPageHeader active="Home" />

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
            <form class="form-inline" @submit.prevent="$doTagSearch" method="get">
              <div class="input-group pb-4">
                <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control" v-model="$global.searchQuery">
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
                <router-link :to="{name: 'Document', params: {tid: document.topic.id, did: document.id}, query: {q: searchTerm}}" class="text-decoration-none">
                  <i v-if="document.icon" :class="document.icon"></i> {{ $localedText(document.title) }}
                </router-link>
              </h2>
              <p class="card-text">{{ $localedText(document.summary) }}</p>
              <p v-if="document.tags && $localedText(document.tags).length>0" style="font-size: small">
                <router-link v-for="tag in $localedText(document.tags)" v-bind:key="tag"
                             :to="{name: 'TagSearch', query:{q: tag, l: $i18n.locale}}"
                             class="badge bg-secondary text-decoration-none link-light me-1" style="font-size: 0.65rem !important;">
                  {{ tag }}
                </router-link>
              </p>
              <router-link :to="{name: 'Document', params: {tid: document.topic.id, did: document.id}, query: {q: searchTerm}}" class="text-decoration-none">
                {{ $t('read') }} <i class="bi bi-arrow-right-circle"></i>
              </router-link>
            </div>
          </div>
          <div class="mb-4">
            <small>{{searchTotalResults}} results in {{searchDuration}} seconds</small>
          </div>
        </div>

        <lego-sidebar no-search />
      </div>
    </div>

    <lego-page-footer />
  </div>
</template>

<script>
import {switchLanguage} from "@/_shared/i18n"
import {watch} from 'vue'
import {useRoute} from "vue-router"
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import legoSidebar from './_sidebar.vue'

export default {
  name: 'Search',
  inject: ['$global'],
  components: {legoPageHeader, legoPageFooter, legoSidebar},
  mounted() {
    const vue = this
    const route = useRoute()
    watch(
        () => route.query.q,
        async () => {
          vue.$global.searchQuery = this.searchTerm
          vue._search(vue)
        }
    )
    watch(
        () => route.query.l,
        async () => vue._search(vue),
    )
    switchLanguage(this.searchLocale)
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
    _fetchSiteMeta(vue) {
      vue.$fetchSiteMeta(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
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
              vue._search(vue)
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
      vue.$tagSearch(vue.searchTerm,
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.searchHits = apiResp.data.docs
              vue.searchDuration = apiResp.data.duration
              vue.searchTotalResults = apiResp.data.total
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          }
      )
    },
  },
  data() {
    return {
      searchHits: [],
      status: -1,
      errorMsg: '',
      searchDuration: 0,
      searchTotalResults: 0,
    }
  },
}
</script>
