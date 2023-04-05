<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <div class="container px-5">
        <router-link :to="{name: 'Home'}" class="navbar-brand">
          <i v-if="siteMeta.icon!=''" :class="siteMeta.icon" style="padding-right: 4px"/>
          <span class="text-primary">{{ _siteNameFirst }}</span>{{ _siteNameLast }}
        </router-link>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
                aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
            <li class="nav-item">
              <router-link :to="{name: 'Home'}" class="nav-link active">{{ $t('home') }}</router-link>
            </li>
            <li class="nav-item dropdown">
              <a class="nav-link dropdown-toggle" id="navbarDropdownTopics" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">{{ $t('topics') }}</a>
              <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdownTopics">
                <li v-for="topic in topicList" v-bind:key="topic.id">
                  <router-link :to="{name: 'Topic', params: {tid: topic.id}}" class="dropdown-item">
                    <i v-if="topic.icon!=''" :class="topic.icon"/>
                    {{ _localedText(topic.title) }}
                  </router-link>
                </li>
              </ul>
            </li>
            <li class="nav-item dropdown">
              <a class="nav-link dropdown-toggle" id="navbarDropdownSettings" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false"><i class="bi bi-translate"></i></a>
              <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdownSettings">
                <li v-for="lang in _siteLanguages" v-bind:key="lang">
                  <a class="dropdown-item" href="#/!" @click="swichLanguage(lang, false)">
                    <span class="badge text-dark">{{ lang }}</span> {{siteMeta.languages[lang]}}
                  </a>
                </li>
              </ul>
            </li>
          </ul>
        </div>
      </div>
    </nav>

    <header class="bg-dark py-5">
      <div class="container px-5">
        <div class="row gx-5 justify-content-center">
          <div class="col-lg-6">
            <div class="text-center my-5">
              <h1 class="display-6 fw-bolder text-white mb-2" style="font-size:1.5rem !important;">{{ _localedText(siteMeta.description) }}</h1>
              <!--<p class="lead text-white-50 mb-4">{{ _localedText(siteMeta.description) }}</p>-->
              <div class="d-grid gap-3 d-flex justify-content-center" v-if="Object.keys(siteMeta.contacts).length > 0" style="font-size: larger">
                <a v-if="siteMeta.contacts.website" :title="siteMeta.contacts.website" :href="siteMeta.contacts.website" target="_blank"><i class="bi bi-globe" /></a>
                <a v-if="siteMeta.contacts.email" :title="siteMeta.contacts.email" :href="'mailto:'+siteMeta.contacts.email" target="_blank"><i class="bi bi-envelope" /></a>
                <a v-if="siteMeta.contacts.github" :title="siteMeta.contacts.github" :href="siteMeta.contacts.github" target="_blank"><i class="bi bi-github" /></a>
                <a v-if="siteMeta.contacts.facebook" :title="siteMeta.contacts.facebook" :href="siteMeta.contacts.facebook" target="_blank"><i class="bi bi-facebook" /></a>
                <a v-if="siteMeta.contacts.linkedin" :title="siteMeta.contacts.linkedin" :href="siteMeta.contacts.linkedin" target="_blank"><i class="bi bi-linkedin" /></a>
                <a v-if="siteMeta.contacts.twitter" :title="siteMeta.contacts.twitter" :href="siteMeta.contacts.twitter" target="_blank"><i class="bi bi-twitter" /></a>
                <a v-if="siteMeta.contacts.slack" :title="siteMeta.contacts.slack" :href="siteMeta.contacts.slack" target="_blank"><i class="bi bi-slack" /></a>
                <a v-if="siteMeta.contacts.discord" :title="siteMeta.contacts.discord" :href="siteMeta.contacts.discord" target="_blank"><i class="bi bi-discord" /></a>
              </div>
            </div>
          </div>
        </div>
        <div class="row gx-5 justify-content-center">
          <div class="col-lg-6">
            <form class="form-inline" @submit.prevent="doSearch" method="get">
              <div class="input-group mb-3">
                <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control form-control-lg" v-model="searchQuery">
                <button type="submit" class="btn btn-light" :value="$t('search')"><i class="bi bi-search" /></button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </header>

    <section class="py-5 border-bottom">
      <div class="container px-5 my-5">
        <div class="row gx-5">
          <div class="col-lg-4 col-md-6 col-12 mb-5 mb-lg-0" v-for="topic in topicList" v-bind:key="topic.id">
            <div class="feature bg-primary bg-gradient text-white rounded-3 mb-3" style="cursor: pointer" @click="$router.push({name:'Topic',params:{tid:topic.id}})">
              <i v-if="topic.icon!=''" :class="topic.icon" /><i v-else class="bi bi-square" />
            </div>
            <h2 class="h4 fw-bolder">
              <router-link :to="{name: 'Topic', params: {tid: topic.id}}" class="text-decoration-none">
                {{ _localedText(topic.title) }}
              </router-link>
            </h2>
            <p>{{ _localedText(topic.description) }}</p>
            <router-link :to="{name: 'Topic', params: {tid: topic.id}}" class="text-decoration-none">
              {{ $t('explore') }} <i class="bi bi-arrow-right"></i>
            </router-link>
          </div>
        </div>
      </div>
    </section>

    <footer class="py-5 bg-dark">
      <div class="container px-5">
        <p class="m-0 text-center text-white">
          Powered by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
          Copyright &copy; {{ siteMeta.name }}.
          <small>Build {{ siteMeta.tags.build }}</small>
        </p>
      </div>
    </footer>
  </div>
</template>

<style lang="css">
.feature {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 4rem;
  width: 4rem;
  font-size: 2rem;
}
</style>

<script>
import {apiDoGet, apiSite, apiTopics} from "@/utils/api_client"
import i18n, {swichLanguage} from "@/i18n"
import {styleByHash, extractLeadingFromName, extractTrailingFromName} from "./utils"

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
    _siteLanguages() {
      return Object.keys(this.siteMeta.languages).filter(value => {
        return value != 'default'
      })
    },
  },
  methods: {
    swichLanguage,
    _styleClassForTopic(topic) {
      const styleList = ["item-blue", "item-green", "item-red", "item-pink", "item-purple", "item-orange"]
      return styleByHash(topic.id, styleList)
    },
    _localedText(inputMap) {
      return inputMap[i18n.global.locale] ? inputMap[i18n.global.locale] : inputMap
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
              vue.errorMsg = vue.status + ": " + apiResp.message
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
              vue.errorMsg = vue.status + ": " + apiResp.message
            }
          },
          (err) => {
            vue.errorMsg = err
          })
    },
    doSearch() {
      this.$router.push({name: 'Search', query: {q: this.searchQuery, l: i18n.global.locale}})
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
