<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">
    {{ errorMsg }}
    <hr/>
    <p class="btn btn-outline-primary mb-0" @click="$reload()"><i class="bi bi-arrow-clockwise"></i> {{ $t('reload') }}</p>
  </div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert"><i class="bi bi-hourglass"></i> {{ $t('wait') }}</div>
  <div v-else class="body-orange">
    <div class="page-wrapper">
      <lego-page-header search />

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
                  <div v-else v-for="document in searchHits" v-bind:key="document.id" class="section-block">
                    <router-link :to="{name: 'Document', params: {tid: document.topic.id, did: document.id}, query: {q: searchTerm}}" class="nav-link">
                      <h3 class="block-title"><i v-if="document.icon!=''" :class="document.icon" class="pe-1"/>{{ $localedText(document.title) }}</h3>
                    </router-link>
                    <router-link :to="{name: 'Document', params: {tid: document.topic.id, did: document.id}}" class="text-decoration-none text-muted">
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
                  <div class="section-block">
                    <small>{{searchTotalResults}} results in {{searchDuration}} seconds</small>
                  </div>
                </section>
              </div>
            </div>

            <lego-sidebar />
          </div>
        </div>
      </div>
    </div>

    <legoPageFooter />
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
  inject: ['$global', '$siteMeta'],
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
      vue.$search(vue.searchTerm,
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
      searchQuery: '',
      searchDuration: 0,
      searchTotalResults: 0,
    }
  },
}
</script>
