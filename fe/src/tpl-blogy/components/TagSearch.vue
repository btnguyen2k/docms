<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">
    {{ errorMsg }}
    <hr/>
    <p class="btn btn-outline-primary mb-0" @click="$reload()"><i class="bi bi-arrow-clockwise"></i> {{ $t('reload') }}</p>
  </div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert"><i class="bi bi-hourglass"></i> {{ $t('wait') }}</div>
  <div v-else>
    <lego-page-header active="home" />

    <div class="section search-result-wrap">
      <div class="container">
        <div class="row">
          <div class="col-12">
            <div class="heading">{{ $t('tag')}}: "{{ searchTerm }}"</div>
          </div>
        </div>
        <div class="row posts-entry">
          <div class="col-lg-8">
            <div v-if="searchHits.length==0" class="alert alert-warning" role="alert">{{ $t('search_no_result') }}</div>
            <div class="blog-entry d-flex blog-entry-search-item" v-for="doc in searchHits" v-bind:key="doc.id" data-aos="fade-up">
              <router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}, query: {q: searchTerm}}" class="img-link me-4">
                <img data-aos="zoom-in" :src="$calcDocumentEntryImgUrl(doc, doc.topic.id, '//placehold.co/440x440/214252/90A1A9?text='+$localedText(doc.id).replaceAll(' ','%20'), 's')" class="img-fluid">
              </router-link>
              <div class="col-9">
                <span class="date">{{ $unixTimestampToReadable(doc.tu) }} &bullet; <router-link :to="{name: 'Topic', params: {tid: doc.topic.id}}">{{ $localedText(doc.topic.title) }}</router-link></span>
                <h2><router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}, query: {q: searchTerm}}">{{ $localedText(doc.title) }}</router-link></h2>
                <p>{{ $localedText(doc.summary) }}</p>
                <p><router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}, query: {q: searchTerm}}" class="btn btn-sm btn-outline-primary">{{ $t('read') }}</router-link></p>
              </div>
            </div>
            <!--            <div class="row text-start pt-5 border-top">-->
            <!--              <div class="col-md-12">-->
            <!--                <div class="custom-pagination">-->
            <!--                  <span>1</span>-->
            <!--                  <a href="#">2</a>-->
            <!--                  <a href="#">3</a>-->
            <!--                  <a href="#">4</a>-->
            <!--                  <span>...</span>-->
            <!--                  <a href="#">15</a>-->
            <!--                </div>-->
            <!--              </div>-->
            <!--            </div>-->
          </div>

          <sidebar />
        </div>
      </div>
    </div>

    <lego-page-footer :document-list="$latestDocuments" />
  </div>
</template>

<script>
import {switchLanguage} from '@/_shared/i18n'
import {watch} from 'vue'
import {useRoute} from "vue-router"
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import sidebar from './_sidebar.vue'

export default {
  name: 'Search',
  inject: ['$global', '$latestDocuments'],
  components: {legoPageHeader, legoPageFooter, sidebar},
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
              vue.$updatePageTitle({search: vue.searchTerm})
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
