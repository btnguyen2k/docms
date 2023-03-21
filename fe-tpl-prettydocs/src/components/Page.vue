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
<!--            <button @click="coderDocsSidebarToggler" ref="docs-sidebar-toggler" id="docs-sidebar-toggler" class="docs-sidebar-toggler docs-sidebar-visible me-2 d-xl-none" type="button">-->
<!--              <span></span>-->
<!--              <span></span>-->
<!--              <span></span>-->
<!--            </button>-->
<!--            <div class="site-logo">-->
<!--              <a class="navbar-brand" style="cursor: pointer" @click="goHome">-->
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
<!--            &lt;!&ndash;            <a href="https://themes.3rdwavemedia.com/bootstrap-templates/startup/coderdocs-free-bootstrap-5-documentation-template-for-software-projects/" class="btn btn-primary d-none d-lg-flex">Download</a>&ndash;&gt;-->
<!--          </div>-->
<!--        </div>-->
<!--      </div>-->
<!--    </header>-->

<!--    <div class="docs-wrapper">-->
<!--      <div ref="docs-sidebar" id="docs-sidebar" class="docs-sidebar">-->
<!--        <div class="top-search-box d-lg-none p-3">-->
<!--          <form class="search-form" @submit.prevent="popup('not implemented yet')">-->
<!--            <input type="text" placeholder="Search the docs..." name="search" class="form-control search-input">-->
<!--            <button type="submit" class="btn search-btn" value="Search"><ficon :icon="['fas', 'search']"/></button>-->
<!--          </form>-->
<!--        </div>-->
<!--        <nav id="docs-nav" class="docs-nav navbar">-->
<!--          <ul class="section-items list-unstyled nav flex-column pb-3">-->
<!--            <template v-for="itopic in topicList">-->
<!--              <li class="nav-item section-title">-->
<!--                <a :class="itopic.id==topic.id?'nav-link scrollto active':'nav-link scrollto'" style="cursor: pointer" @click="doViewTopic(itopic.id)"><span class="theme-icon-holder me-2"><ficon fixedWidth :icon="_iconize(itopic.icon)"/></span>{{itopic.title}}</a>-->
<!--              </li>-->
<!--              <template v-if="itopic.id==topic.id">-->
<!--                <li v-for="ipage in pageList" :key="ipage.id" class="nav-item">-->
<!--                  <a :class="ipage.id==page.id?'nav-link scrollto active':'nav-link scrollto'" style="cursor: pointer" @click="doViewPage(ipage.id)">&lt;!&ndash;<span class="theme-icon-holder me-2"><ficon fixedWidth :icon="_iconize(ipage.icon)"/></span>&ndash;&gt;{{ipage.title}}</a>-->
<!--                </li>-->
<!--              </template>-->
<!--            </template>-->
<!--          </ul>-->
<!--        </nav>-->
<!--      </div>-->

<!--      <div class="docs-content">-->
<!--        <div class="container">-->
<!--          <article class="docs-article" :id="page.id">-->
<!--            <h1 class="docs-heading"><span class="me-2"><ficon :icon="_iconize(page.icon)"/></span>{{page.title}}-->
<!--              <span class="docs-time">Last updated: {{pageUpdatedTimestamp}}</span></h1>-->
<!--            <section class="docs-intro">-->
<!--              {{page.summary}}-->
<!--            </section>-->
<!--            <section class="docs-section" :id="page.id+'-content'" v-html="pageContentRendered">-->
<!--            </section>-->
<!--          </article>-->

<!--          <footer class="footer">-->
<!--            <div class="footer-bottom text-center py-5">-->
<!--              &lt;!&ndash;/* This template is free as long as you keep the footer attribution link. If you'd like to use the template without the attribution link, you can buy the commercial license via our website: themes.3rdwavemedia.com Thank you for your support. :) */&ndash;&gt;-->
<!--              <small class="copyright">Designed with <ficon style="color: #fb866a;" :icon="['fas', 'heart']"/> by <a class="theme-link" href="http://themes.3rdwavemedia.com" target="_blank">Xiaoying Riley</a> for developers</small>-->
<!--              <br/>-->
<!--              <small class="copyright">Site built with <a class="theme-link" href="https://github.com/btnguyen2k/libro" target="_blank">Libro</a>.</small>-->
<!--            </div>-->
<!--          </footer>-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->
<!--  </div>-->
</template>

<!--<script>-->
<!--import clientUtils from "@/utils/api_client"-->
<!--import {registerPopstate, unregisterPopstate} from "@/components/utils"-->
<!--import {registerResize, unregisterResize} from "@/components/utils"-->
<!--import {iconize, markdownRender} from "@/components/utils"-->

<!--const regTrailingSlash = /\/+$/-->

<!--export default {-->
<!--  name: 'Page',-->
<!--  mounted() {-->
<!--    this.$once("hook:beforeDestroy", () => {-->
<!--      unregisterPopstate(this.handleBackFoward)-->
<!--      unregisterResize(this.coderDocsResponsiveSidebar) // CoderDocs-->
<!--    })-->
<!--    registerPopstate(this.handleBackFoward)-->
<!--    registerResize(this.coderDocsResponsiveSidebar) // CoderDocs: onresize-->
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
<!--    pageContentRendered() {-->
<!--      return markdownRender(this.page.content, true)-->
<!--    },-->
<!--    pageUpdatedTimestamp() {-->
<!--      const d = new Date(this.page.t_updated)-->
<!--      return d.toLocaleString()-->
<!--    },-->
<!--  },-->
<!--  methods: {-->
<!--    handleBackFoward() {-->
<!--      const pathBase = this.$router.options.base.replace(regTrailingSlash, '')-->
<!--      const vuePath = window.location.pathname.slice(pathBase.length) // remove the 'base' prefix-->
<!--      const result = this.$router.resolve(vuePath)-->
<!--      if (result && result.resolved && result.resolved.name=='Page') {-->
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

<!--            // now we verify if page exists-->
<!--            for (let i = vue.pageList.length-1; i>=0; i&#45;&#45;) {-->
<!--              const pid = vue.pageList[i].id-->
<!--              vue.pageMap[pid] = vue.pageList[i]-->
<!--            }-->
<!--            const pid = vue.$route.params.pid-->
<!--            if (pid=='') {-->
<!--              if (vue.pageList.length>0) {-->
<!--                this.$router.push({name: "Page", params: {tid: tid, pid: vue.pageList[0].id}}).finally(()=>this.initPage())-->
<!--              } else {-->
<!--                this.$router.push({name: "Topic", params: {tid: tid}})-->
<!--              }-->
<!--              return-->
<!--            }-->
<!--            vue.page = vue.pageMap[pid]-->
<!--            if (!vue.page) {-->
<!--              vue.$router.push({-->
<!--                name: "Error",-->
<!--                params: {errorMsg: vue.$i18n.t('error_page_not_found', {page: pid}), topic: vue.topic}-->
<!--              })-->
<!--              return-->
<!--            }-->

<!--            vue.foundStatus = 1 // found product-->
<!--            vue.$nextTick(()=>{-->
<!--              // CoderDocs: onload-->
<!--              vue.coderDocsResponsiveSidebar()-->
<!--            })-->
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
<!--      if (tid==this.$route.params.tid) {-->
<!--        this.doViewPage(this.pageList[0].id)-->
<!--      } else {-->
<!--        this.foundStatus = -1-->
<!--        this.$route.params.pid = ''-->
<!--        this.fetchTopicInfo(tid)-->
<!--      }-->
<!--    },-->
<!--    doViewPage(pid) {-->
<!--      if (pid!=this.$route.params.pid) {-->
<!--        const tid = this.$route.params.tid-->
<!--        this.$router.push({name: "Page", params: {tid: tid, pid: pid}}).finally(()=>this.initPage())-->
<!--      }-->
<!--    },-->
<!--    goHome() {-->
<!--      this.$router.push({name: "Home"})-->
<!--    },-->
<!--    popup(msg) {-->
<!--      alert(msg)-->
<!--    },-->
<!--    coderDocsResponsiveSidebar() { // CoderDocs-->
<!--      const w = window.innerWidth-->
<!--      const sidebar = this.$refs['docs-sidebar']-->
<!--      if (sidebar) {-->
<!--        if (w >= 1024) {-->
<!--          sidebar.classList.remove('sidebar-hidden')-->
<!--          sidebar.classList.add('sidebar-visible')-->
<!--        } else {-->
<!--          sidebar.classList.remove('sidebar-visible')-->
<!--          sidebar.classList.add('sidebar-hidden')-->
<!--        }-->
<!--      }-->
<!--    },-->
<!--    coderDocsSidebarToggler() { // CoderDocs-->
<!--      const sidebar = this.$refs['docs-sidebar']-->
<!--      if (sidebar) {-->
<!--        if (sidebar.classList.contains('sidebar-visible')) {-->
<!--          sidebar.classList.remove('sidebar-visible');-->
<!--          sidebar.classList.add('sidebar-hidden');-->
<!--        } else {-->
<!--          sidebar.classList.remove('sidebar-hidden');-->
<!--          sidebar.classList.add('sidebar-visible');-->
<!--        }-->
<!--      }-->
<!--    },-->
<!--  },-->
<!--  data() {-->
<!--    return {-->
<!--      product: {},-->
<!--      topicList: [],-->
<!--      topic: {},-->
<!--      pageList: [],-->
<!--      pageMap: {},-->
<!--      page: {},-->
<!--      foundStatus: -1,-->
<!--      errorMsg: '',-->
<!--    }-->
<!--  },-->
<!--}-->
<!--</script>-->
