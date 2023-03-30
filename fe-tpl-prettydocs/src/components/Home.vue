<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else class="landing-page">
    <div class="page-wrapper">
      <header class="header text-center">
        <div class="container">
          <div class="branding">
            <h1 class="logo">
              <fa-icon v-if="siteMeta.icon!=''" aria-hidden="true" class="icon" style="padding-right: 8px" :icon="siteMeta.icon" />
              <span class="text-highlight">{{ _siteNameFirst }}</span><span class="text-bold">{{ _siteNameLast }}</span>
            </h1>
          </div>

          <div class="tagline">
            {{ _localedText(siteMeta.description) }}
          </div>

          <div class="main-search-box pt-3 pb-4 d-inline-block">
            <form class="form-inline search-form justify-content-center" @submit.prevent="doSearch" method="get">
              <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control search-input" v-model="searchQuery">
              <button type="submit" class="btn search-btn" :value="$t('search')"><fa-icon icon="fas fa-search" /></button>
            </form>
          </div>

          <div class="social-container" v-if="Object.keys(siteMeta.contacts).length > 0">
            {{ $t('contact_info') }}
            <a class="link p-1 fa-lg" v-if="siteMeta.contacts.website" :title="siteMeta.contacts.website" :href="siteMeta.contacts.website" target="_blank"><fa-icon icon="fas fa-globe" /></a>
            <a class="link p-1 fa-lg" v-if="siteMeta.contacts.email" :title="siteMeta.contacts.email" :href="'mailto:'+siteMeta.contacts.email" target="_blank"><fa-icon icon="fas fa-envelope" /></a>
            <a class="link p-1 fa-lg" v-if="siteMeta.contacts.github" :title="siteMeta.contacts.github" :href="siteMeta.contacts.github" target="_blank"><fa-icon icon="fab fa-github" /></a>
            <a class="link p-1 fa-lg" v-if="siteMeta.contacts.facebook" :title="siteMeta.contacts.facebook" :href="siteMeta.contacts.facebook" target="_blank"><fa-icon icon="fab fa-facebook" /></a>
            <a class="link p-1 fa-lg" v-if="siteMeta.contacts.linkedin" :title="siteMeta.contacts.linkedin" :href="siteMeta.contacts.linkedin" target="_blank"><fa-icon icon="fab fa-linkedin" /></a>
            <a class="link p-1 fa-lg" v-if="siteMeta.contacts.slack" :title="siteMeta.contacts.slack" :href="siteMeta.contacts.slack" target="_blank"><fa-icon icon="fab fa-slack" /></a>
            <a class="link p-1 fa-lg" v-if="siteMeta.contacts.twitter" :title="siteMeta.contacts.twitter" :href="siteMeta.contacts.twitter" target="_blank"><fa-icon icon="fab fa-twitter" /></a>
          </div>
        </div>
      </header>

      <section class="cards-section text-center">
        <div class="container">
          <div id="cards-wrapper" class="cards-wrapper row">
            <div v-for="topic in topicList" v-bind:key="topic.id" :class="_styleClassForTopic(topic)+' item col-12 col-lg-4'">
              <div class="item-inner">
                <div v-if="topic.icon!=''" class="icon-holder icon"><fa-icon :icon="topic.icon"/></div>
                <h3 class="title">{{ _localedText(topic.title) }}</h3>
                <p class="intro">{{ _localedText(topic.description) }}</p>
                <router-link class="link" :to="{ name: 'Topic', params: { tid: topic.id } }"><span></span></router-link>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
    <footer class="footer text-center">
      <div class="container">
        <!--/* This template is free as long as you keep the footer attribution link. If you'd like to use the template without the attribution link, you can buy the commercial license via our website: themes.3rdwavemedia.com Thank you for your support. :) */-->
        <small class="copyright">
          Powered by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
          Theme <span style="font-family: monospace, monospace;">PrettyDocs</span> designed with <span class="sr-only">love</span><fa-icon icon="fas fa-heart"/> by <a href="https://themes.3rdwavemedia.com/" target="_blank">Xiaoying Riley</a> for developers.
        </small>
      </div>
    </footer>
  </div>
</template>

<script>
import { apiDoGet, apiSite, apiTopics } from "@/utils/api_client"
import i18n from "@/i18n"
import { styleByHash, extractLeadingFromName, extractTrailingFromName } from "./utils"

export default {
  name: 'Home',
  mounted() {
    this._fetchSiteMeta(this)
  },
  computed: {
    _siteNameFirst() {
      return extractLeadingFromName(this.siteMeta.name)
    },
    _siteNameLast() {
      return extractTrailingFromName(this.siteMeta.name)
    },
  },
  methods: {
    _styleClassForTopic(topic) {
      const styleList = ["item-blue", "item-green", "item-red", "item-pink", "item-purple", "item-orange"]
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
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          (err) => {
            vue.errorMsg = err
          })
    },
    doSearch() {
      this.$router.push({name: 'Search', query: {q: this.searchQuery}})
    },
  },
  data() {
    return {
      siteMeta: {},
      topicList: [],
      status: -1,
      errorMsg: '',
      searchQuery: '',
    }
  },
}
</script>
