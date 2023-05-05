<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">
    {{ errorMsg }}
    <hr/>
    <p class="btn btn-outline-primary mb-0" @click="$reload()"><i class="bi bi-arrow-clockwise"></i> {{ $t('reload') }}</p>
  </div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert"><i class="bi bi-hourglass"></i> {{ $t('wait') }}</div>
  <div v-else>
    <legoPageHeader active="Home" />

    <header class="bg-dark py-5">
      <div class="container px-5">
        <div class="row gx-5 justify-content-center">
          <div class="col-lg-6">
            <div class="text-center my-5">
              <h1 class="display-6 fw-bolder text-white mb-2" style="font-size:1.5rem !important;">{{ $localedText($siteMeta.description) }}</h1>
              <!--<p class="lead text-white-50 mb-4">{{ $localedText($siteMeta.description) }}</p>-->
              <div class="d-grid gap-3 d-flex justify-content-center" v-if="Object.keys($siteMeta.contacts).length > 0" style="font-size: larger">
                <a v-if="$siteMeta.contacts.website" :title="$siteMeta.contacts.website" :href="$siteMeta.contacts.website" target="_blank"><i class="bi bi-globe" /></a>
                <a v-if="$siteMeta.contacts.email" :title="$siteMeta.contacts.email" :href="'mailto:'+$siteMeta.contacts.email" target="_blank"><i class="bi bi-envelope" /></a>
                <a v-if="$siteMeta.contacts.github" :title="$siteMeta.contacts.github" :href="$siteMeta.contacts.github" target="_blank"><i class="bi bi-github" /></a>
                <a v-if="$siteMeta.contacts.facebook" :title="$siteMeta.contacts.facebook" :href="$siteMeta.contacts.facebook" target="_blank"><i class="bi bi-facebook" /></a>
                <a v-if="$siteMeta.contacts.linkedin" :title="$siteMeta.contacts.linkedin" :href="$siteMeta.contacts.linkedin" target="_blank"><i class="bi bi-linkedin" /></a>
                <a v-if="$siteMeta.contacts.twitter" :title="$siteMeta.contacts.twitter" :href="$siteMeta.contacts.twitter" target="_blank"><i class="bi bi-twitter" /></a>
                <a v-if="$siteMeta.contacts.slack" :title="$siteMeta.contacts.slack" :href="$siteMeta.contacts.slack" target="_blank"><i class="bi bi-slack" /></a>
                <a v-if="$siteMeta.contacts.discord" :title="$siteMeta.contacts.discord" :href="$siteMeta.contacts.discord" target="_blank"><i class="bi bi-discord" /></a>
              </div>
            </div>
          </div>
        </div>
        <div class="row gx-5 justify-content-center">
          <div class="col-lg-6">
            <form class="form-inline" @submit.prevent="$doSearch" method="get">
              <div class="input-group mb-3">
                <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control form-control-lg" v-model="$global.searchQuery">
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
          <div class="col-lg-4 col-md-6 col-12 mb-5 mb-lg-0 pb-4" v-for="topic in $siteTopics" v-bind:key="topic.id">
            <div class="feature bg-primary bg-gradient text-white rounded-3 mb-3" style="cursor: pointer" @click="$router.push({name:'Topic',params:{tid:topic.id}})">
              <i v-if="topic.icon!=''" :class="topic.icon" /><i v-else class="bi bi-square" />
            </div>
            <h2 class="h4 fw-bolder">
              <router-link :to="{name: 'Topic', params: {tid: topic.id}}" class="text-decoration-none">
                {{ $localedText(topic.title) }}
              </router-link>
            </h2>
            <p style="margin-bottom: 0rem !important;">{{ $localedText(topic.description) }}</p>
            <router-link :to="{name: 'Topic', params: {tid: topic.id}}" class="text-decoration-none">
              {{ $t('explore') }} <i class="bi bi-arrow-right"></i>
            </router-link>
          </div>
        </div>
      </div>
    </section>

    <lego-page-footer />
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
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'

export default {
  name: 'Home',
  inject: ['$global', '$siteMeta', '$siteTopics'],
  components: {legoPageHeader, legoPageFooter},
  mounted() {
    this._fetchSiteMeta(this)
  },
  methods: {
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
