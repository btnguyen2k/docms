<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else>
    <header class="header fixed-top">
      <div class="branding docs-branding">
        <div class="container-fluid position-relative py-2">
          <div class="docs-logo-wrapper">
            <div class="site-logo">
              <router-link class="navbar-brand" :to="{name: 'Home'}">
                <span class="logo-text"><i v-if="$siteMeta.icon!=''" :class="$siteMeta.icon"></i> {{ $siteFirstName }}<span class="text-alt">{{ $siteLastName }}</span></span>
                <!--<img class="logo-icon me-2" :src="$router.options.meta.base.replace(/\/+$/, '')+'/images/coderdocs-logo.svg'">-->
                <!--<span class="logo-text">{{ $siteFirstName }}<span class="text-alt">{{ $siteLastName }}</span></span>-->
              </router-link>
            </div>
          </div>

          <div class="docs-top-utilities d-flex justify-content-end align-items-center">
            <lego-social-list class="social-list list-inline mx-md-3 mx-lg-5 mb-0 d-none d-lg-flex"/>
          </div>
        </div>
      </div>
    </header>

    <div class="page-header theme-bg-dark py-5 text-center position-relative">
      <div class="theme-bg-shapes-right"></div>
      <div class="theme-bg-shapes-left"></div>
      <div class="container">
        <h1 class="page-heading single-col-max mx-auto d-none d-md-inline">{{ $siteMeta.name }}</h1>
        <div class="page-intro single-col-max mx-auto">{{ $localedText($siteMeta.description) }}</div>
        <div class="main-search-box pt-3 d-block mx-auto">
          <form class="search-form w-100" @submit.prevent="$doSearch">
            <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control search-input" v-model="$global.searchQuery">
            <button type="submit" class="btn search-btn" :value="$t('search')"><i class="fas fa-search"></i></button>
          </form>
        </div>
      </div>
    </div>

    <div class="page-content">
      <div class="container">
        <div class="docs-overview py-5">
          <div class="row justify-content-center">
            <div class="col-12 col-lg-4 py-3" v-for="topic in $siteTopics" v-bind:key="topic.id">
              <div class="card shadow-sm">
                <div class="card-body">
                  <h5 class="card-title mb-3">
                    <span class="theme-icon-holder card-icon-holder me-2"><i v-if="topic.icon!=''" :class="topic.icon"/></span>
                    <span class="card-title-text">{{ $localedText(topic.title) }}</span>
                  </h5>
                  <div class="card-text">
                    {{ $localedText(topic.description) }}
                  </div>
                  <router-link class="card-link-mask" :to="{name: 'Topic', params: {tid: topic.id}}"></router-link>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <section class="cta-section text-center py-5 theme-bg-dark position-relative">
      <div class="theme-bg-shapes-right"></div>
      <div class="theme-bg-shapes-left"></div>
<!--      <div class="container">-->
<!--        <h3 class="mb-2 text-white mb-3">Launch Your Software Project Like A Pro</h3>-->
<!--        <div class="section-intro text-white mb-3 single-col-max mx-auto">Want to launch your software project and start getting traction from your target users? Check out our premium <a class="text-white" href="https://themes.3rdwavemedia.com/bootstrap-templates/startup/coderpro-bootstrap-5-startup-template-for-software-projects/">Bootstrap 5 startup template CoderPro</a>! It has everything you need to promote your product.</div>-->
<!--        <div class="pt-3 text-center">-->
<!--          <a class="btn btn-light" href="https://themes.3rdwavemedia.com/bootstrap-templates/startup/coderpro-bootstrap-5-startup-template-for-software-projects/">Get CoderPro<i class="fas fa-arrow-alt-circle-right ml-2"></i></a>-->
<!--        </div>-->
<!--      </div>-->
    </section>

    <footer class="footer">
      <div class="footer-bottom text-center py-5 px-1">
        <lego-social-list class="social-list list-unstyled pb-4 mb-0" />

        <!--/* This template is free as long as you keep the footer attribution link. If you'd like to use the template without the attribution link, you can buy the commercial license via our website: themes.3rdwavemedia.com Thank you for your support. :) */-->
        <small class="copyright d-none d-sm-none d-md-none d-lg-inline float-none">
          Powered by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
          Theme <span style="font-family: monospace, monospace;">CoderDocs</span> designed with <span class="sr-only">love</span><i class="fas fa-heart" style="color: #fb866a;"></i> by <a class="theme-link" href="http://themes.3rdwavemedia.com" target="_blank">Xiaoying Riley</a> for developers.
        </small>
        <small class="copyright d-none d-sm-none d-md-inline d-lg-none float-none">
          <i class="fas fa-bolt-lightning"></i> by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
          <i class="fas fa-pen-ruler"></i> <span style="font-family: monospace, monospace;">CoderDocs</span> by <a href="https://themes.3rdwavemedia.com/" target="_blank">Xiaoying Riley</a>.
        </small>
        <small class="copyright d-md-none d-lg-none float-none">
          <i class="fas fa-bolt-lightning"></i> by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
          <i class="fas fa-pen-ruler"></i> by <a href="https://themes.3rdwavemedia.com/" target="_blank">Xiaoying Riley</a>.
        </small>
        <small v-if="$siteMeta.tags.build" class="d-none d-lg-inline float-end">Build: {{ $siteMeta.tags.build }}</small>
        <ul class="nav nav-pills float-start align-middle" style="margin-top: -6px !important;">
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" data-bs-toggle="dropdown" href="#" role="button" aria-expanded="false"><i class="fas fa-language"></i></a>
            <ul class="dropdown-menu" style="font-size: small">
              <li v-for="lang in $siteLanguages" v-bind:key="lang">
                <a class="dropdown-item" href="#" @click="swichLanguage(lang, false)"><span class="badge text-dark">{{ lang }}</span> {{$siteMeta.languages[lang]}}</a>
              </li>
              <li class="d-lg-none"><hr class="dropdown-divider"></li>
              <li v-if="$siteMeta.tags.build" class="dropdown-item d-lg-none">Build: {{ $siteMeta.tags.build }}</li>
            </ul>
          </li>
        </ul>
      </div>
    </footer>
  </div>
</template>

<script>
import {switchLanguage} from '@/_shared/i18n'
import legoSocialList from './_socialList.vue'

export default {
  name: 'Home',
  inject: ['$global', '$siteMeta', '$siteLanguages', '$siteTopics', '$siteFirstName', '$siteLastName'],
  components: {legoSocialList},
  mounted() {
    this._fetchSiteMeta(this)
  },
  methods: {
    swichLanguage: switchLanguage,
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
