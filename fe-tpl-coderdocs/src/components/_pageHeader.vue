<template>
  <header class="header fixed-top">
    <div class="branding docs-branding">
      <div class="container-fluid position-relative py-2">
        <div class="docs-logo-wrapper">
          <button @click="$coderDocsSidebarToggler" ref="docs-sidebar-toggler" id="docs-sidebar-toggler" class="docs-sidebar-toggler docs-sidebar-visible me-2 d-xl-none" type="button">
            <span></span>
            <span></span>
            <span></span>
          </button>
          <div class="site-logo">
            <router-link class="navbar-brand" :to="{name: 'Home'}">
              <span class="logo-text"><i v-if="$siteMeta.icon!=''" :class="$siteMeta.icon"></i> {{ $siteFirstName }}<span class="text-alt">{{ $siteLastName }}</span></span>
            </router-link>
          </div>
        </div>

        <div class="docs-top-utilities d-flex justify-content-end align-items-center">
          <div class="top-search-box d-none d-lg-flex">
            <form class="search-form" @submit.prevent="$doSearch">
              <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control search-input" v-model="$global.searchQuery">
              <button type="submit" class="btn search-btn" :value="$t('search')"><i class="fas fa-search"></i></button>
            </form>
          </div>

          <lego-social-list class="social-list list-inline mx-md-3 mx-lg-5 mb-0 d-none d-lg-flex" />
        </div>
      </div>
    </div>
  </header>
</template>

<script>
import legoSocialList from './_socialList.vue'

export default {
  name: 'lego-page-header',
  inject: ['$global', '$siteMeta', '$siteFirstName', '$siteLastName'],
  components: {legoSocialList},
}
</script>
