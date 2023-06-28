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

  <section class="section">
    <div class="container">
      <div class="row blog-entries element-animate">
        <div class="col-12 main-content">
          <div class="post-content-body">
            <div class="docms-content img-fit img-center" v-html="documentContentRendered"></div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <div class="section">
    <div class="container">
      <div class="row">
        <div class="offset-lg-3 offset-md-2 offset-sm-1 col-auto mb-5 mb-lg-0" data-aos="fade-up">
        </div>
      </div>
    </div>
  </div>

</template>

<script>
import {markdownRender} from '@/_shared/utils/docms_utils'

export default {
  name: 'DocumentStyleNoSidebar',
  inject: ['$siteMeta'],
  props: ['document'],
  computed: {
    documentContentRendered() {
      return markdownRender(this.$localedText(this.document.content), {
        sanitize: true,
        tags: this.$siteMeta.tags,
      })
    },
  },
}
</script>
