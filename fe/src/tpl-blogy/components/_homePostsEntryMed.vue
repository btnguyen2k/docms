<template>
  <!-- Start posts-entry/med:many -->
  <section class="section">
    <div class="container">
      <div class="row mb-4" v-if="$props['topic']">
        <div class="col-sm-6">
          <h2 class="posts-entry-title">{{ $localedText($props['topic'].title) }}</h2>
        </div>
        <div class="col-sm-6 text-sm-end"><router-link :to="{name: 'Topic', params: {tid: topic.id}}" class="read-more">{{ $t('view_all') }}</router-link></div>
      </div>
      <div class="row">
        <div class="col-lg-4 mb-4" v-for="doc in docs" v-bind:key="doc.id" data-aos="fade-up">
          <div class="post-entry-alt">
            <router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}}" class="img-link" data-aos="zoom-in">
              <img :src="$calcDocumentEntryImgUrl(doc, doc.topic.id, '//placehold.co/700x440/214252/90A1A9?text='+$localedText(doc.topic.id).replaceAll(' ','%20'), 'h')" class="img-fluid">
            </router-link>
            <div class="excerpt">
              <h2><router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}}">{{ $localedText(doc.title) }}</router-link></h2>
              <div class="post-meta align-items-center text-left clearfix">
                <figure class="author-figure mb-0 me-3 float-start"><img :src="$calcAuthorAvatarUrl(doc.author)" class="img-fluid"></figure>
                <span class="d-inline-block mt-1">{{ doc.author.name }}</span>
                <span> - {{ $unixTimestampToReadable(doc.tu) }}</span>
              </div>
              <p>{{ $trimText($localedText(doc.summary), 70) }}</p>
              <p><router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}}" class="read-more">{{ $t('read') }}</router-link></p>
            </div>
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
  name: 'lego-home-posts-entry-med',
  inject: ['$siteMeta'],
  props: ['topic', 'document-list'],
  computed: {
    docs() {
      return this.$props['documentList'].slice(0,9)
    },
  }
}
</script>
