<template>
  <div class="hero overlay inner-page bg-primary py-5">
    <div class="container">
      <div class="row align-items-center justify-content-center text-center pt-5">
        <div class="col-lg-6">
          <h1 class="heading text-white mb-3" data-aos="zoom-in">{{ $localedText(document.title) }}</h1>
        </div>
      </div>
    </div>
  </div>

  <div class="section">
    <div class="container">
      <div class="row">
        <div class="offset-lg-3 offset-md-2 offset-sm-1 col-auto mb-5 mb-lg-0" data-aos="fade-up" data-aos-delay="100">
          <div class="contact-info">
            <div v-if="documentData.location" class="mt-4">
              <i class="fas fa-location-dot"></i>
              <h4 class="mb-2">{{ documentData.location.label }}</h4>
              <p v-html="markdownRender(documentData.location.content, {inline: true})"></p>
            </div>
            <div v-if="documentData.website" class="mt-4">
              <i class="fas fa-globe"></i>
              <h4 class="mb-2">{{ documentData.website.label }}</h4>
              <p v-html="markdownRender(documentData.website.content, {inline: true})"></p>
            </div>
            <div v-if="documentData.email" class="mt-4">
              <i class="fas fa-envelope"></i>
              <h4 class="mb-2">{{ documentData.email.label }}</h4>
              <p v-html="markdownRender(documentData.email.content, {inline: true})"></p>
            </div>
            <div v-if="documentData.discord" class="mt-4">
              <i class="fab fa-discord"></i>
              <h4 class="mb-2">{{ documentData.discord.label }}</h4>
              <p v-html="markdownRender(documentData.discord.content, {inline: true})"></p>
            </div>
            <div v-if="documentData.facebook" class="mt-4">
              <i class="fab fa-facebook"></i>
              <h4 class="mb-2">{{ documentData.facebook.label }}</h4>
              <p v-html="markdownRender(documentData.facebook.content, {inline: true})"></p>
            </div>
            <div v-if="documentData.github" class="mt-4">
              <i class="fab fa-github"></i>
              <h4 class="mb-2">{{ documentData.github.label }}</h4>
              <p v-html="markdownRender(documentData.github.content, {inline: true})"></p>
            </div>
            <div v-if="documentData.linkedin" class="mt-4">
              <i class="fab fa-linkedin"></i>
              <h4 class="mb-2">{{ documentData.linkedin.label }}</h4>
              <p v-html="markdownRender(documentData.linkedin.content, {inline: true})"></p>
            </div>
            <div v-if="documentData.slack" class="mt-4">
              <i class="fab fa-slack"></i>
              <h4 class="mb-2">{{ documentData.slack.label }}</h4>
              <p v-html="markdownRender(documentData.slack.content, {inline: true})"></p>
            </div>
            <div v-if="documentData.twitter" class="mt-4">
              <i class="fab fa-twitter"></i>
              <h4 class="mb-2">{{ documentData.twitter.label }}</h4>
              <p v-html="markdownRender(documentData.twitter.content, {inline: true})"></p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {markdownRender} from '@/_shared/utils/docms_utils'

export default {
  name: 'DocumentStyleContact',
  inject: ['$siteMeta'],
  props: ['topic', 'document', 'documentList'],
  computed: {
    documentData() {
      const content = this.$localedText(this.document.content)
      return JSON.parse(content)
    },
  },
  methods: {
    markdownRender
  },
}
</script>
