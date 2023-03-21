<template><div>//TODO</div>
<!--  <div v-if="foundStatus<0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>-->
<!--  <div v-else-if="foundStatus==0" class="alert alert-danger m-4" role="alert">-->
<!--    {{ $t('error_product_not_found', {domain: currentHost}) }}-->
<!--  </div>-->
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
<!--            <div class="top-search-box d-none d-lg-flex">-->
<!--              <form class="search-form" @submit.prevent="popup('not implemented yet')">-->
<!--                <input type="text" placeholder="Search the docs..." name="search" class="form-control search-input">-->
<!--                <button type="submit" class="btn search-btn" value="Search"><ficon :icon="['fas', 'search']"/></button>-->
<!--              </form>-->
<!--            </div>-->
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
<!--            &lt;!&ndash;            <a href="https://themes.3rdwavemedia.com/bootstrap-templates/startup/coderdocs-free-bootstrap-5-documentation-template-for-software-projects/"&ndash;&gt;-->
<!--            &lt;!&ndash;               class="btn btn-primary d-none d-lg-flex">Download</a>&ndash;&gt;-->
<!--          </div>-->
<!--        </div>-->
<!--      </div>-->
<!--    </header>-->

<!--    <div class="page-header theme-bg-dark py-5 text-center position-relative">-->
<!--      <div class="theme-bg-shapes-right"></div>-->
<!--      <div class="theme-bg-shapes-left"></div>-->
<!--      <div class="container">-->
<!--        <h1 class="page-heading single-col-max mx-auto">{{ topic.title }}</h1>-->
<!--        <div class="page-intro single-col-max mx-auto">-->
<!--          {{ topic.summary }}-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->

<!--    <div class="page-content">-->
<!--      <div class="container">-->
<!--        <div class="docs-overview py-5">-->
<!--          <div v-if="pageList.length==0" class="row justify-content-center">-->
<!--            <div class="col-12 py-3">-->
<!--              <div class="card border-warning">-->
<!--                <div class="card-body">-->
<!--                  <h5 class="card-title mb-3">-->
<!--                    <span class="theme-icon-holder card-icon-holder me-2"><ficon :icon="['fas', 'exclamation']"/></span>-->
<!--                    <span class="card-title-text">{{ $t('empty_topic') }}</span>-->
<!--                  </h5>-->
<!--                </div>-->
<!--              </div>-->
<!--            </div>-->
<!--          </div><div v-else class="row justify-content-center">-->
<!--            <div v-for="page in pageList" class="col-12 col-lg-4 py-3">-->
<!--              <div class="card shadow-sm">-->
<!--                <div class="card-body">-->
<!--                  <h5 class="card-title mb-3">-->
<!--                    <span class="theme-icon-holder card-icon-holder me-2"><ficon :icon="_iconize(page.icon)"/></span>-->
<!--                    <span class="card-title-text">{{ page.title }}</span>-->
<!--                  </h5>-->
<!--                  <div class="card-text">{{ page.summary }}</div>-->
<!--                  <a class="card-link-mask" href="javascript:;" @click="doViewPage(page.id)"></a>-->
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
<!--        <div class="section-intro text-white mx-auto col-lg-6 col-sm-12">-->
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
<!--        <div class="section-intro text-white mx-auto col-lg-6 col-sm-12">-->
<!--          <strong>{{ $t('topics') }}</strong>-->
<!--          <ul class="list-unstyled">-->
<!--            <li v-for="topic in topicList">-->
<!--              <a class="text-white" @click="doViewTopic(topic.id)" style="cursor: pointer">{{ topic.title }}</a>-->
<!--            </li>-->
<!--          </ul>-->
<!--        </div>-->
<!--      </div>-->
<!--    </section>-->

<!--    <footer class="footer">-->
<!--      <div class="footer-bottom text-center py-5">-->
<!--        &lt;!&ndash;/* This template is free as long as you keep the footer attribution link. If you'd like to use the template without the attribution link, you can buy the commercial license via our website: themes.3rdwavemedia.com Thank you for your support. :) */&ndash;&gt;-->
<!--        <small class="copyright">Designed with <ficon style="color: #fb866a;" :icon="['fas', 'heart']"/> by <a class="theme-link" href="http://themes.3rdwavemedia.com" target="_blank">Xiaoying Riley</a> for developers.</small>-->
<!--        <br/>-->
<!--        <small class="copyright">Site built with <a class="theme-link" href="https://github.com/btnguyen2k/libro" target="_blank">Libro</a>.</small>-->
<!--      </div>-->
<!--    </footer>-->
<!--  </div>-->
</template>

<!--<script>-->
<!--import clientUtils from "@/utils/api_client"-->
<!--import {iconize, registerPopstate, unregisterPopstate} from "@/components/utils"-->

<!--const regTrailingSlash = /\/+$/-->

<!--export default {-->
<!--  name: 'Topic',-->
<!--  mounted() {-->
<!--    this.$once("hook:beforeDestroy", () => {-->
<!--      unregisterPopstate(this.handleBackFoward)-->
<!--    })-->
<!--    registerPopstate(this.handleBackFoward)-->
<!--    this.initPage()-->
<!--  },-->
<!--  computed: {-->
<!--    prodNameFirst() {-->
<!--      return this.product.name ? this.product.name.slice(0, 2) : ""-->
<!--    },-->
<!--    prodNameLast() {-->
<!--      return this.product.name ? this.product.name.slice(2) : ""-->
<!--    },-->
<!--    currentHost() {-->
<!--      return window.location.host-->
<!--    },-->
<!--  },-->
<!--  methods: {-->
<!--    handleBackFoward() {-->
<!--      const pathBase = this.$router.options.base.replace(regTrailingSlash, '')-->
<!--      const vuePath = window.location.pathname.slice(pathBase.length) // remove the 'base' prefix-->
<!--      const result = this.$router.resolve(vuePath)-->
<!--      if (result && result.resolved && result.resolved.name=='Topic') {-->
<!--        this.initPage()-->
<!--      }-->
<!--    },-->
<!--    initPage() {-->
<!--      this.foundStatus = -1-->
<!--      const vue = this-->
<!--      const apiUrl = clientUtils.apiProduct.replaceAll(':domain', vue.currentHost)-->
<!--      clientUtils.apiDoGet(apiUrl,-->
<!--          (apiRes) => {-->
<!--            let foundProduct = apiRes.status == 200 && apiRes.data.is_published-->
<!--            if (foundProduct) {-->
<!--              vue.product = apiRes.data-->
<!--              vue.topicList = vue.product.topics-->
<!--              this.fetchTopicInfo(vue.$route.params.tid)-->
<!--            } else {-->
<!--              vue.foundStatus = 0-->
<!--            }-->
<!--          },-->
<!--          (err) => {-->
<!--            vue.errorMsg = err-->
<!--          }-->
<!--      )-->
<!--    },-->
<!--    fetchTopicInfo(tid) {-->
<!--      const vue = this-->
<!--      // now we need to verify if topic exists-->
<!--      const apiUrl = clientUtils.apiTopic.replaceAll(':domain', vue.currentHost).replaceAll(':tid', tid)-->
<!--      clientUtils.apiDoGet(apiUrl,-->
<!--          (apiRes)=>{-->
<!--            const foundTopic = apiRes.status == 200-->
<!--            if (!foundTopic) {-->
<!--              vue.$router.push({-->
<!--                name: "Error",-->
<!--                params: {errorMsg: vue.$i18n.t('error_topic_not_found', {topic: tid})}-->
<!--              })-->
<!--              return-->
<!--            }-->

<!--            vue.topic = apiRes.data-->
<!--            vue.pageList = vue.topic.pages-->
<!--            vue.foundStatus = 1 // found product-->
<!--          },-->
<!--          (err)=>{-->
<!--            vue.errorMsg = err-->
<!--          }-->
<!--      )-->
<!--    },-->
<!--    _iconize(icon) {-->
<!--      return iconize(icon)-->
<!--    },-->
<!--    doViewTopic(tid) {-->
<!--      if (tid!=this.$route.params.tid) {-->
<!--        this.$router.push({name: "Topic", params: {tid: tid}}).finally(()=>this.initPage())-->
<!--      }-->
<!--    },-->
<!--    doViewPage(pid) {-->
<!--      const tid = this.$route.params.tid-->
<!--      this.$router.push({name: "Page", params: {tid: tid, pid: pid}})-->
<!--    },-->
<!--    goHome() {-->
<!--      this.$router.push({name: "Home"})-->
<!--    },-->
<!--    popup(msg) {-->
<!--      alert(msg)-->
<!--    },-->
<!--  },-->
<!--  data() {-->
<!--    return {-->
<!--      product: {},-->
<!--      topicList: [],-->
<!--      topic: {},-->
<!--      pageList: [],-->
<!--      foundStatus: -1,-->
<!--      errorMsg: '',-->
<!--    }-->
<!--  },-->
<!--}-->
<!--</script>-->
