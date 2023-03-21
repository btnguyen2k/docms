<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else>
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
            <form class="form-inline search-form justify-content-center" action="" method="get">
              <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control search-input">
              <button type="submit" class="btn search-btn" :value="$t('search')"><fa-icon icon="fas fa-search" /></button>
            </form>
          </div>

          <div class="social-container">
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
            <div v-for="topic in topicList" v-bind:key="topic.id" :class="_topicBlockStyleClass(topic)+' item col-12 col-lg-4'">
              <div class="item-inner">
                <div v-if="topic.icon!=''" class="icon-holder icon"><fa-icon :icon="topic.icon"/></div>
                <h3 class="title">{{ _localedText(topic.title) }}</h3>
                <p class="intro">{{ _localedText(topic.description) }}</p>
                <a class="link" href="start.html"><span></span></a>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
    <footer class="footer text-center">
      <div class="container">
        <!--/* This template is free as long as you keep the footer attribution link. If you'd like to use the template without the attribution link, you can buy the commercial license via our website: themes.3rdwavemedia.com Thank you for your support. :) */-->
        <small class="copyright">Designed with <span class="sr-only">love</span><fa-icon icon="fas fa-heart"/> by <a href="https://themes.3rdwavemedia.com/" target="_blank">Xiaoying Riley</a> for developers.</small>
      </div>
    </footer>
  </div>
</template>

<script>
import clientUtils from "@/utils/api_client"
import i18n from "@/i18n"

export default {
  name: 'Home',
  unmounted() {
    document.body.classList.remove('landing-page')
  },
  mounted() {
    document.body.classList.add('landing-page')

    this.status = 0
    const vue = this
    clientUtils.apiDoGet(clientUtils.apiSite,
        (apiResp) => {
          let status = apiResp.status
          if (status == 200) {
            vue.siteMeta = apiResp.data
            clientUtils.apiDoGet(clientUtils.apiTopics,
                (apiResp) => {
                  vue.status = apiResp.status
                  if (vue.status == 200) {
                    vue.topicList = apiResp.data
                  } else {
                    vue.errorMsg = apiResp.message
                  }
                },
                (err) => {
                  vue.errorMsg = err
                })
          } else {
            vue.errorMsg = apiResp.message
          }
        },
        (err) => {
          vue.errorMsg = err
        })
  },
  computed: {
    _siteNameFirst() {
      return this.siteMeta.name ? this.siteMeta.name.split(' ')[0] : ""
    },
    _siteNameLast() {
      return this.siteMeta.name ? this.siteMeta.name.slice(this._siteNameFirst.length).trim() : ""
    },
  },
  methods: {
    _topicBlockStyleClass(topic) {
      const styleList = ["item-blue", "item-green", "item-red", "item_pink", "item-purple", "item-orange"]
      let hash = 0
      for (let i = 0; i < topic.id.length; i++) {
        hash = ((hash << 5) - hash) + topic.id.charCodeAt(i)
        hash = hash & hash
      }
      return styleList[hash % styleList.length]
    },
    _localedText(inputMap) {
      return inputMap[i18n.global.locale]?inputMap[i18n.global.locale]:inputMap
    },
    doViewTopic(tid) {
      this.$router.push({name: "Topic", params: {tid: tid}})
    },
    goHome() {
      this.$router.push({name: "Home"})
    },
    popup(msg) {
      alert(msg)
    }
  },
  data() {
    return {
      siteMeta: {},
      topicList: [],
      status: -1,
      errorMsg: '',
    }
  },
}
</script>
