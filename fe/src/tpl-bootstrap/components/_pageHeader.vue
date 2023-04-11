<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark sticky-top">
    <div class="container px-5">
      <router-link :to="{name: 'Home'}" class="navbar-brand">
        <i v-if="$siteMeta.icon!=''" :class="$siteMeta.icon" style="padding-right: 4px"/>
        <span class="text-primary">{{ $siteFirstName }}</span>{{ $siteLastName }}
      </router-link>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
              aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link :to="{name: 'Home'}" :class="'nav-link'+($props['active']=='Home'?' active':'')">{{ $t('home') }}</router-link>
          </li>
          <li class="nav-item dropdown">
            <a :class="'nav-link dropdown-toggle'+($props['active']!='Home'?' active':'')" id="navbarDropdownTopics" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">{{ $t('topics') }}</a>
            <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdownTopics">
              <li v-for="topic in $siteTopics" v-bind:key="topic.id">
                <router-link :to="{name: 'Topic', params: {tid: topic.id}}" class="dropdown-item">
                  <i v-if="topic.icon!=''" :class="topic.icon+' fa-fw'"/>
                  {{ $localedText(topic.title) }}
                </router-link>
              </li>
            </ul>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" id="navbarDropdownSettings" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false"><i class="bi bi-translate"></i></a>
            <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdownSettings">
              <li v-for="lang in $siteLanguages" v-bind:key="lang">
                <a class="dropdown-item" href="javascript:;" @click="swichLanguage(lang, false)">
                  <span class="badge text-dark">{{ lang }}</span> {{$siteMeta.languages[lang]}}
                </a>
              </li>
            </ul>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
import {swichLanguage} from "@/_shared/i18n"

export default {
  name: 'lego-page-header',
  inject: ['$siteMeta', '$siteFirstName', '$siteLastName', '$siteTopics', '$siteLanguages'],
  props: ['active'],
  methods: {
    swichLanguage,
  },
}
</script>
