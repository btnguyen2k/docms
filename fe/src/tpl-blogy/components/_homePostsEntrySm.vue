<template>
  <!-- Start posts-entry/small:4 -->
  <section class="section posts-entry posts-entry-sm bg-light" v-if="$props['documentList'] && $props['documentList'].length>=4">
    <div class="container">
      <div class="row mb-4" v-if="$props['topic']">
        <div class="col-sm-6">
          <h2 class="posts-entry-title"><i v-if="$props['topic'].icon!=''" :class="$props['topic'].icon"/> {{ $localedText($props['topic'].title) }}</h2>
        </div>
        <div class="col-sm-6 text-sm-end"><router-link :to="{name: 'Topic', params: {tid: topic.id}}" class="read-more">{{ $t('view_all') }}</router-link></div>
      </div>
      <div class="row">
        <div class="col-md-6 col-lg-3" v-for="doc in docs" v-bind:key="doc.id" data-aos="fade-up">
          <div class="blog-entry">
            <router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}}" class="img-link" data-aos="zoom-in">
              <img :src="$calcDocumentEntryImgUrl(doc, doc.topic.id, '//placehold.co/700x440/214252/90A1A9?text='+$localedText(doc.id).replaceAll(' ','%20'), 'h')" class="img-fluid">
            </router-link>
            <span class="date">{{ $unixTimestampToReadable(doc.tu) }}</span>
            <h2><router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}}">{{ $localedText(doc.title) }}</router-link></h2>
            <p>{{ $localedText(doc.summary) }}</p>
            <p><router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}}" class="read-more">{{ $t('read') }}</router-link></p>
          </div>
        </div>
      </div>
    </div>
  </section>
  <!-- End posts-entry -->
</template>

<style lang="css">
</style>

<script>
export default {
  name: 'lego-home-posts-entry-sm',
  inject: ['$siteMeta'],
  props: ['topic', 'document-list'],
  computed: {
    docs() {
      return this.$props['documentList'].slice(0, 4)
    },
  }
}
</script>
