<template>
  <div class="site-mobile-menu site-navbar-target">
    <div class="site-mobile-menu-header">
      <div class="site-mobile-menu-close">
        <span class="icofont-close js-menu-toggle"></span>
      </div>
    </div>
    <div class="site-mobile-menu-body"></div>
  </div>

  <nav class="site-nav">
    <div class="container">
      <div class="menu-bg-wrap">
        <div class="site-navigation">
          <div class="row g-0 align-items-center">
            <div class="col-2">
              <router-link :to="{name: 'Home'}" class="logo m-0 float-start">{{ $siteFirstName }}<span class="text-primary">{{ $siteLastName }}</span></router-link>
            </div>
            <div class="col-8 text-center">
              <form class="search-form d-inline-block d-lg-none" @submit.prevent="$doSearch" method="get" style="width: 65% !important;">
                <input type="text" class="form-control" :placeholder="$t('search')" v-model="$global.searchQuery">
                <span class="bi-search"></span>
              </form>

              <ul class="js-clone-nav d-none d-lg-inline-block text-start site-menu mx-auto">
                <li :class="$props['active']=='home'?'active':''"><router-link :to="{name: 'Home'}">{{ $t('home') }}</router-link></li>
                <li v-for="topic in $siteTopics.slice(0,4)" v-bind:key="topic.id" :class="$props['active']=='topic'&&$props['topic'].id==topic.id?'active':''">
                  <router-link :to="{name: 'Topic', params: {tid: topic.id}}">{{ $localedText(topic.title) }}</router-link>
                </li>
                <li class="has-children" v-if="$siteTopics.length > 4">
                  <a>{{ $t('topics') }}</a>
                  <ul class="dropdown text-nowrap">
                    <li v-for="topic in $siteTopics" v-bind:key="topic.id">
                      <router-link :to="{name: 'Topic', params:{tid: topic.id}}">
                        <i v-if="topic.icon!=''" :class="topic.icon+' fa-fw'"></i> {{ $localedText(topic.title) }}
                      </router-link>
                    </li>
                  </ul>
                </li>
                <li class="has-children">
                  <a><i class="fas fa-language"></i></a>
                  <ul class="dropdown text-nowrap">
                    <li v-for="lang in $siteLanguages" v-bind:key="lang">
                      <a class="dropdown-item" href="javascript:;" @click="swichLanguage(lang, false)">
                        <span class="badge text-dark">{{ lang }}</span> {{$siteMeta.languages[lang]}}
                      </a>
                    </li>
                  </ul>
                </li>
              </ul>
            </div>
            <div class="col-2 text-end">
              <a href="#" class="burger ms-auto float-end site-menu-toggle js-menu-toggle d-inline-block d-lg-none light">
                <span></span>
              </a>
              <form class="search-form d-none d-lg-inline-block" @submit.prevent="$doSearch" method="get">
                <input type="text" class="form-control" :placeholder="$t('search_prompt')" v-model="$global.searchQuery">
                <span class="bi-search"></span>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
import {switchLanguage} from "@/_shared/i18n"

export default {
  name: 'lego-page-header',
  inject: ['$global', '$siteMeta', '$siteFirstName', '$siteLastName', '$siteTopics', '$siteLanguages'],
  props: ['active', 'topic'],
  methods: {
    swichLanguage: switchLanguage,
  },
}
</script>
