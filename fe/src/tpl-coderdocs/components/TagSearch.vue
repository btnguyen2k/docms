<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">
    {{ errorMsg }}
    <hr/>
    <p class="btn btn-outline-primary mb-0" @click="$reload()"><i class="bi bi-arrow-clockwise"></i> {{ $t('reload') }}</p>
  </div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert"><i class="bi bi-hourglass"></i> {{ $t('wait') }}</div>
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
              <router-link :to="{name: 'Document', params: {tid: document.topic.id, did: document.id}, query: {q: searchTerm}}" class="text-decoration-none">
                <h2 class="section-heading"><i v-if="document.icon!=''" :class="document.icon"></i> {{ $localedText(document.title) }}</h2>
              </router-link>
              <p>{{ $localedText(document.summary) }}</p>
              <p v-if="document.tags && $localedText(document.tags).length>0" style="font-size: small">
                <router-link v-for="tag in $localedText(document.tags)" v-bind:key="tag"
                             :to="{name: 'TagSearch', query:{q: tag, l: $i18n.locale}}"
                             :class="$calcTagCloudCSS(tag)+' me-1'" style="font-size: 0.65rem !important;">
                  {{ tag }}
                </router-link>
              </p>
            </section>
          </article>

          <small>{{searchTotalResults}} results in {{searchDuration}} seconds</small>

          <lego-page-footer />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {registerPopstate, unregisterPopstate} from "@/_shared/utils/docms_utils"
import {switchLanguage} from "@/_shared/i18n"
import {watch} from 'vue'
import {useRoute} from "vue-router"
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import legoSidebar from "./_sidebar.vue"

const regTrailingSlash = /\/+$/

export default {
  name: 'Search',
  inject: ['$global', '$coderDocsResponsiveSidebar', '$calcTagCloudCSS'],
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
