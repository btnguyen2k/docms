<template>
  <footer class="footer">
    <div class="container text-center py-5">
      <!--/* This template is free as long as you keep the footer attribution link. If you'd like to use the template without the attribution link, you can buy the commercial license via our website: themes.3rdwavemedia.com Thank you for your support. :) */-->
      <small class="copyright d-none d-sm-none d-md-none d-lg-inline float-none">
        Powered by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
        Theme <span style="font-family: monospace, monospace;">CoderDocs</span> designed with <span class="sr-only">love</span><i class="fas fa-heart" style="color: #fb866a;"></i> by <a class="theme-link" href="http://themes.3rdwavemedia.com" target="_blank">Xiaoying Riley</a> for developers.
      </small>
      <small class="copyright d-none d-sm-none d-md-inline d-lg-none float-none">
        <fa-icon icon="fas fa-bolt-lightning"></fa-icon> by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
        <fa-icon icon="fas fa-pen-ruler"></fa-icon> <span style="font-family: monospace, monospace;">CoderDocs</span> by <a href="https://themes.3rdwavemedia.com/" target="_blank">Xiaoying Riley</a>.
      </small>
      <small class="copyright d-md-none d-lg-none float-none">
        <fa-icon icon="fas fa-bolt-lightning"></fa-icon> by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
        <fa-icon icon="fas fa-pen-ruler"></fa-icon> by <a href="https://themes.3rdwavemedia.com/" target="_blank">Xiaoying Riley</a>.
      </small>
      <small v-if="$siteMeta.tags.build" class="d-none d-lg-inline float-end">Build: {{ $siteMeta.tags.build }}</small>
      <ul class="nav nav-pills float-start align-middle" style="margin-top: -6px !important;"><!--style="font-size: 0.85em !important; "-->
        <li class="nav-item dropdown">
          <a class="nav-link dropdown-toggle" data-bs-toggle="dropdown" href="#" role="button" aria-expanded="false"><fa-icon icon="fas fa-language"></fa-icon></a>
          <ul class="dropdown-menu" style="font-size: small">
            <li v-for="lang in $siteLanguages" v-bind:key="lang">
              <a class="dropdown-item" href="#" @click="swichLanguage(lang, false)"><span class="badge text-dark">{{ lang }}</span> {{$siteMeta.languages[lang]}}</a>
            </li>
            <li class="d-lg-none"><hr class="dropdown-divider"></li>
            <li v-if="$siteMeta.tags.build" class="dropdown-item d-lg-none">Build: {{ $siteMeta.tags.build }}</li>
          </ul>
        </li>
      </ul>
      <lego-social-list class="social-list list-unstyled pt-4 mb-0" />
    </div>
  </footer>
</template>

<script>
import legoSocialList from './_socialList.vue'
import {swichLanguage} from "@/i18n"

export default {
  name: 'lego-page-footer',
  inject: ['$siteMeta', '$siteLanguages'],
  components: {legoSocialList},
  methods: {
    swichLanguage,
  }
}
</script>
