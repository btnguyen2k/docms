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
<!--              <span class="text-highlight">{{ siteMeta.name }}</span>-->
            </h1>
          </div>

          <div class="tagline">
            {{ _localedText(siteMeta.description) }}
          </div>

          <div class="main-search-box pt-3 pb-4 d-inline-block">
            <form class="form-inline search-form justify-content-center" action="" method="get">
              <input type="text" placeholder="{{ $t('search_msg') }}" name="q" class="form-control search-input">
              <button type="submit" class="btn search-btn" value="{{ $t('search') }}"><fa-icon icon="fas fa-search" /></button>
            </form>
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
<!--  <div v-else>-->
<!--    <header class="header fixed-top">-->
<!--      <div class="branding docs-branding">-->
<!--        <div class="container-fluid position-relative py-2">-->
<!--          <div class="docs-logo-wrapper">-->
<!--            <div class="site-logo">-->
<!--              <a class="navbar-brand" @click="goHome" style="cursor: pointer">-->
<!--                <img class="logo-icon me-2" :src="$router.options.base.replace(/\/+$/, '')+'/images/coderdocs-logo.svg'" alt="logo">-->
<!--                <span class="logo-text">{{ prodNameFirst }}<span class="text-alt">{{ prodNameLast }}</span></span>-->
<!--              </a>-->
<!--            </div>-->
<!--          </div>-->
<!--          <div class="docs-top-utilities d-flex justify-content-end align-items-center">-->
<!--            <ul class="social-list list-inline mx-md-3 mx-lg-5 mb-0 d-none d-lg-flex">-->
<!--              <li v-if="product.contacts.website" class="list-inline-item">-->
<!--                <a :href="product.contacts.website" title="Website"><ficon fixedWidth :icon="['fas', 'globe']"/></a>-->
<!--              </li>-->
<!--              <li v-if="product.contacts.email" class="list-inline-item">-->
<!--                <a :href="'mailto:'+product.contacts.email" title="Email"><ficon fixedWidth :icon="['fas', 'envelope']"/></a>-->
<!--              </li>-->
<!--              <li v-if="product.contacts.github" class="list-inline-item">-->
<!--                <a :href="product.contacts.github" title="GitHub"><ficon fixedWidth :icon="['fab', 'github']"/></a>-->
<!--              </li>-->
<!--              <li v-if="product.contacts.facebook" class="list-inline-item">-->
<!--                <a :href="product.contacts.facebook" title="Facebook"><ficon fixedWidth :icon="['fab', 'facebook']"/></a>-->
<!--              </li>-->
<!--              <li v-if="product.contacts.linkedin" class="list-inline-item">-->
<!--                <a :href="product.contacts.linkedin" title="LinkedIn"><ficon fixedWidth :icon="['fab', 'linkedin']"/></a>-->
<!--              </li>-->
<!--              <li v-if="product.contacts.slack" class="list-inline-item">-->
<!--                <a :href="product.contacts.slack" title="Slack"><ficon fixedWidth :icon="['fab', 'slack']"/></a>-->
<!--              </li>-->
<!--              <li v-if="product.contacts.twitter" class="list-inline-item">-->
<!--                <a :href="product.contacts.twitter" title="Twitter"><ficon fixedWidth :icon="['fab', 'twitter']"/></a>-->
<!--              </li>-->
<!--            </ul>-->
<!--&lt;!&ndash;            <a href="https://themes.3rdwavemedia.com/bootstrap-templates/startup/coderdocs-free-bootstrap-5-documentation-template-for-software-projects/"&ndash;&gt;-->
<!--&lt;!&ndash;               class="btn btn-primary d-none d-lg-flex">Download</a>&ndash;&gt;-->
<!--          </div>-->
<!--        </div>-->
<!--      </div>-->
<!--    </header>-->

<!--    <div class="page-header theme-bg-dark py-5 text-center position-relative">-->
<!--      <div class="theme-bg-shapes-right"></div>-->
<!--      <div class="theme-bg-shapes-left"></div>-->
<!--      <div class="container">-->
<!--&lt;!&ndash;        <h1 class="page-heading single-col-max mx-auto">Documentation</h1>&ndash;&gt;-->
<!--        <div class="page-intro single-col-max mx-auto">-->
<!--          {{ product.desc }}-->
<!--        </div>-->
<!--        <div class="main-search-box pt-3 d-block mx-auto">-->
<!--          <form class="search-form w-100" @submit.prevent="popup('not implemented yet')">-->
<!--            <input type="text" placeholder="Search the docs..." name="search" class="form-control search-input">-->
<!--            <button type="submit" class="btn search-btn" value="Search"><ficon :icon="['fas', 'search']"/></button>-->
<!--          </form>-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->

<!--    <div class="page-content">-->
<!--      <div class="container">-->
<!--        <div class="docs-overview py-5">-->
<!--          <div class="row justify-content-center">-->
<!--            <div v-for="topic in topicList" class="col-12 col-lg-4 py-3">-->
<!--              <div class="card shadow-sm border-primary">-->
<!--                <div class="card-header bg-primary">-->
<!--                  <h5 class="card-title">-->
<!--                    <span class="theme-icon-holder card-icon-holder me-2"><ficon :icon="_iconize(topic.icon)"/></span>-->
<!--                    <span class="card-title-text text-white">{{ topic.title }}</span>-->
<!--                  </h5>-->
<!--                </div>-->
<!--                <div class="card-body">-->
<!--                  <div class="card-text">{{ topic.summary }}</div>-->
<!--                  <a class="card-link-mask" style="cursor: pointer" @click="doViewTopic(topic.id)"></a>-->
<!--                </div>-->
<!--              </div>-->
<!--            </div>-->
<!--          </div>-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->

<!--    <section class="cta-section py-5 theme-bg-dark position-relative">-->
<!--      <div class="theme-bg-shapes-right"></div>-->
<!--      <div class="theme-bg-shapes-left"></div>-->
<!--      <div class="container row mx-5">-->
<!--        <div class="section-intro text-white mx-auto col-12">-->
<!--          <strong>{{ $t('contacts') }}</strong>-->
<!--          <ul class="list-unstyled">-->
<!--            <li v-if="product.contacts.website">-->
<!--              <a :href="product.contacts.website" title="Website"><ficon class="text-white" fixedWidth :icon="['fas', 'globe']"/></a>-->
<!--              <a :href="product.contacts.website" title="Website" class="text-white mx-1">{{ product.contacts.website }}</a>-->
<!--            </li>-->
<!--            <li v-if="product.contacts.email">-->
<!--              <a :href="'mailto:'+product.contacts.email" title="Email"><ficon class="text-white" fixedWidth :icon="['fas', 'envelope']"/></a>-->
<!--              <a :href="'mailto:'+product.contacts.email" title="Email" class="text-white mx-1">{{ product.contacts.email }}</a>-->
<!--            </li>-->
<!--            <li v-if="product.contacts.github">-->
<!--              <a :href="product.contacts.github" title="GitHub"><ficon class="text-white" fixedWidth :icon="['fab', 'github']"/></a>-->
<!--              <a :href="product.contacts.github" title="GitHub" class="text-white mx-1">{{ product.contacts.github }}</a>-->
<!--            </li>-->
<!--            <li v-if="product.contacts.facebook">-->
<!--              <a :href="product.contacts.facebook" title="Facebook"><ficon class="text-white" fixedWidth :icon="['fab', 'facebook']"/></a>-->
<!--              <a :href="product.contacts.facebook" title="Facebook" class="text-white mx-1">{{ product.contacts.facebook }}</a>-->
<!--            </li>-->
<!--            <li v-if="product.contacts.linkedin">-->
<!--              <a :href="product.contacts.linkedin" title="LinkedIn"><ficon class="text-white" fixedWidth :icon="['fab', 'linkedin']"/></a>-->
<!--              <a :href="product.contacts.linkedin" title="LinkedIn" class="text-white mx-1">{{ product.contacts.linkedin }}</a>-->
<!--            </li>-->
<!--            <li v-if="product.contacts.slack">-->
<!--              <a :href="product.contacts.slack" title="Slack"><ficon class="text-white" fixedWidth :icon="['fab', 'slack']"/></a>-->
<!--              <a :href="product.contacts.slack" title="Slack" class="text-white mx-1">{{ product.contacts.slack }}</a>-->
<!--            </li>-->
<!--            <li v-if="product.contacts.twitter">-->
<!--              <a :href="product.contacts.twitter" title="Twitter"><ficon class="text-white" fixedWidth :icon="['fab', 'twitter']"/></a>-->
<!--              <a :href="product.contacts.twitter" title="Twitter" class="text-white mx-1">{{ product.contacts.twitter }}</a>-->
<!--            </li>-->
<!--          </ul>-->
<!--        </div>-->
<!--      </div>-->
<!--    </section>-->

<!--    <footer class="footer">-->
<!--      <div class="footer-bottom text-center py-5">-->
<!--&lt;!&ndash;        <ul class="social-list list-unstyled pb-4 mb-0">&ndash;&gt;-->
<!--&lt;!&ndash;          <li v-if="product.contacts.website" class="list-inline-item">&ndash;&gt;-->
<!--&lt;!&ndash;            <a :href="product.contacts.website" title="Website"><ficon fixedWidth :icon="['fas', 'globe']"/></a>&ndash;&gt;-->
<!--&lt;!&ndash;          </li>&ndash;&gt;-->
<!--&lt;!&ndash;          <li v-if="product.contacts.github" class="list-inline-item">&ndash;&gt;-->
<!--&lt;!&ndash;            <a :href="product.contacts.github" title="GitHub"><ficon fixedWidth :icon="['fab', 'github']"/></a>&ndash;&gt;-->
<!--&lt;!&ndash;          </li>&ndash;&gt;-->
<!--&lt;!&ndash;          <li v-if="product.contacts.facebook" class="list-inline-item">&ndash;&gt;-->
<!--&lt;!&ndash;            <a :href="product.contacts.facebook" title="Facebook"><ficon fixedWidth :icon="['fab', 'facebook']"/></a>&ndash;&gt;-->
<!--&lt;!&ndash;          </li>&ndash;&gt;-->
<!--&lt;!&ndash;          <li v-if="product.contacts.linkedin" class="list-inline-item">&ndash;&gt;-->
<!--&lt;!&ndash;            <a :href="product.contacts.linkedin" title="LinkedIn"><ficon fixedWidth :icon="['fab', 'linkedin']"/></a>&ndash;&gt;-->
<!--&lt;!&ndash;          </li>&ndash;&gt;-->
<!--&lt;!&ndash;          <li v-if="product.contacts.slack" class="list-inline-item">&ndash;&gt;-->
<!--&lt;!&ndash;            <a :href="product.contacts.slack" title="Slack"><ficon fixedWidth :icon="['fab', 'slack']"/></a>&ndash;&gt;-->
<!--&lt;!&ndash;          </li>&ndash;&gt;-->
<!--&lt;!&ndash;          <li v-if="product.contacts.twitter" class="list-inline-item">&ndash;&gt;-->
<!--&lt;!&ndash;            <a :href="product.contacts.twitter" title="Twitter"><ficon fixedWidth :icon="['fab', 'twitter']"/></a>&ndash;&gt;-->
<!--&lt;!&ndash;          </li>&ndash;&gt;-->
<!--&lt;!&ndash;        </ul>&ndash;&gt;-->
<!--        &lt;!&ndash;/* This template is free as long as you keep the footer attribution link. If you'd like to use the template without the attribution link, you can buy the commercial license via our website: themes.3rdwavemedia.com Thank you for your support. :) */&ndash;&gt;-->
<!--        <small class="copyright">Designed with <ficon style="color: #fb866a;" :icon="['fas', 'heart']"/> by <a class="theme-link" href="http://themes.3rdwavemedia.com" target="_blank">Xiaoying Riley</a> for developers.</small>-->
<!--        <br/>-->
<!--        <small class="copyright">Site built with <a class="theme-link" href="https://github.com/btnguyen2k/libro" target="_blank">Libro</a>.</small>-->
<!--      </div>-->
<!--    </footer>-->
<!--  </div>-->
</template>

<script>
import clientUtils from "@/utils/api_client"
import {iconize} from "./utils"
// import {FaIcon} from "@fortawesome/vue-fontawesome";
import i18n from "@/i18n"

export default {
  name: 'Home',
  // components: {FaIcon},
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
            console.log(apiResp.data)
            clientUtils.apiDoGet(clientUtils.apiTopics,
                (apiResp) => {
                  vue.status = apiResp.status
                  if (vue.status == 200) {
                    vue.topicList = apiResp.data
                    console.log(apiResp.data)
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
    locale() {
      return i18n.global.locale
    }
  },
  methods: {
    _iconize(icon) {
      return iconize(icon)
    },
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
