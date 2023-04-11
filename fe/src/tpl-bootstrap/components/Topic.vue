<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else-if="!topic.id" class="alert alert-danger m-4" role="alert">{{ $t('error_topic_not_found', {topic: $route.params.tid}) }}</div>
  <div v-else>
    <legoPageHeader />

    <header class="bg-light border-bottom mb-4">
      <div class="container px-5">
        <nav aria-label="breadcrumb">
          <ol class="breadcrumb">
            <li class="breadcrumb-item"><router-link :to="{ name: 'Home' }">{{ $t('home') }}</router-link></li>
            <li class="breadcrumb-item active" aria-current="page">{{ $localedText(topic.title) }}</li>
          </ol>
        </nav>

        <div class="text-center my-4">
          <h1 class="fw-bolder"><i v-if="topic.icon" :class="topic.icon"></i> {{ $localedText(topic.title) }}</h1>
          <p class="lead mb-0">{{ $localedText(topic.description) }}</p>
        </div>

<!--        <div class="row justify-content-center d-sm-flex d-md-flex d-lg-none">-->
<!--          <div class="col-lg-6">-->
<!--            <form class="form-inline" @submit.prevent="$doSearch" method="get">-->
<!--              <div class="input-group pb-4">-->
<!--                <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control" v-model="$global.searchQuery">-->
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
                  <i v-if="document.icon" :class="document.icon"></i> {{ $localedText(document.title) }}
                </router-link>
              </h2>
              <p class="card-text">{{ $localedText(document.summary) }}</p>
              <router-link :to="{name: 'Document', params: {tid: topic.id, did: document.id}}" class="text-decoration-none">
                {{ $t('read') }} <i class="bi bi-arrow-right-circle"></i>
              </router-link>
            </div>
          </div>
        </div>

        <lego-sidebar :topic-id="topic.id" />
      </div>
    </div>

    <lego-page-footer />
  </div>
</template>

<script>
import {useRoute} from 'vue-router'
import {watch} from 'vue'
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import legoSidebar from './_sidebar.vue'

export default {
  name: 'Topic',
  inject: ['$siteTopics'],
  components: {legoPageHeader, legoPageFooter, legoSidebar},
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
  methods: {
    _fetchSiteMeta(vue) {
      vue.$fetchSiteMeta(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue._fetchTopics(vue, vue.$route.params.tid)
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _fetchTopics(vue, topicId) {
      vue.$fetchSiteTopics(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.$siteTopics.forEach(t => {
                if (t.id == topicId) {
                  vue.topic = t
                  vue._fetchDocuments(vue, topicId)
                }
              })
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _fetchDocuments(vue, topicId) {
      vue.$fetchDocumentsForTopic(topicId,
          () => vue.status=0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.documentList = apiResp.data
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => vue.errorMsg = err
      )
    },
  },
  data() {
    return {
      topic: {},
      documentList: [],
      status: -1,
      errorMsg: '',
    }
  },
}
</script>
