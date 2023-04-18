<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else class="landing-page">
    <div class="page-wrapper">
      <header class="header text-center">
        <div class="container">
          <div class="branding">
            <h1 class="logo">
              <i v-if="$siteMeta.icon!=''" aria-hidden="true" class="icon" style="padding-right: 8px" :class="$siteMeta.icon" />
              <span class="text-highlight">{{ $siteFirstName }}</span><span class="text-bold">{{ $siteLastName }}</span>
            </h1>
          </div>

          <div class="tagline">
            {{ $localedText($siteMeta.description) }}
          </div>

          <div class="main-search-box pt-3 pb-4 d-inline-block">
            <form class="form-inline search-form justify-content-center" @submit.prevent="$doSearch" method="get">
              <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control search-input" v-model="$global.searchQuery">
              <button type="submit" class="btn search-btn" :value="$t('search')"><i class="fas fa-search" /></button>
            </form>
          </div>

          <div class="social-container" v-if="Object.keys($siteMeta.contacts).length > 0">
            {{ $t('contact_info') }}
            <a class="link p-1 fa-lg" v-if="$siteMeta.contacts.website" :title="$siteMeta.contacts.website" :href="$siteMeta.contacts.website" target="_blank"><i class="fas fa-globe" /></a>
            <a class="link p-1 fa-lg" v-if="$siteMeta.contacts.email" :title="$siteMeta.contacts.email" :href="'mailto:'+$siteMeta.contacts.email" target="_blank"><i class="fas fa-envelope" /></a>
            <a class="link p-1 fa-lg" v-if="$siteMeta.contacts.github" :title="$siteMeta.contacts.github" :href="$siteMeta.contacts.github" target="_blank"><i class="fab fa-github" /></a>
            <a class="link p-1 fa-lg" v-if="$siteMeta.contacts.facebook" :title="$siteMeta.contacts.facebook" :href="$siteMeta.contacts.facebook" target="_blank"><i class="fab fa-facebook" /></a>
            <a class="link p-1 fa-lg" v-if="$siteMeta.contacts.linkedin" :title="$siteMeta.contacts.linkedin" :href="$siteMeta.contacts.linkedin" target="_blank"><i class="fab fa-linkedin" /></a>
            <a class="link p-1 fa-lg" v-if="$siteMeta.contacts.slack" :title="$siteMeta.contacts.slack" :href="$siteMeta.contacts.slack" target="_blank"><i class="fab fa-slack" /></a>
            <a class="link p-1 fa-lg" v-if="$siteMeta.contacts.twitter" :title="$siteMeta.contacts.twitter" :href="$siteMeta.contacts.twitter" target="_blank"><i class="fab fa-twitter" /></a>
            <a class="link p-1 fa-lg" v-if="$siteMeta.contacts.discord" :title="$siteMeta.contacts.discord" :href="$siteMeta.contacts.discord" target="_blank"><i class="fab fa-discord" /></a>
          </div>
        </div>
      </header>

      <section class="cards-section text-center">
        <div class="container">
          <div id="cards-wrapper" class="cards-wrapper row">
            <div v-for="topic in $siteTopics" v-bind:key="topic.id" :class="_styleClassForTopic(topic)+' item col-12 col-lg-4'">
              <div class="item-inner">
                <div v-if="topic.icon!=''" class="icon-holder icon"><i :class="topic.icon"/></div>
                <h3 class="title">{{ $localedText(topic.title) }}</h3>
                <p class="intro">{{ $localedText(topic.description) }}</p>
                <router-link class="link" :to="{ name: 'Topic', params: { tid: topic.id } }"><span></span></router-link>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>

    <lego-page-footer />
  </div>
</template>

<script>
import legoPageFooter from './_pageFooter.vue'

export default {
  name: 'Home',
  inject: ['$global', '$siteMeta', '$siteTopics', '$siteFirstName', '$siteLastName'],
  components: {legoPageFooter},
  mounted() {
    this._fetchSiteMeta(this)
  },
  methods: {
    _styleClassForTopic(topic) {
      const styleList = ["item-blue", "item-green", "item-red", "item-pink", "item-purple", "item-orange"]
      return this.$styleByHash(topic.id, styleList)
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
            if (vue.status != 200) {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
  },
  data() {
    return {
      status: -1,
      errorMsg: '',
    }
  },
}
</script>
