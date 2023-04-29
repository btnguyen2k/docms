<template>
  <!-- Start posts-entry/text right -->
  <section class="section posts-entry" v-if="$props['documentList'] && $props['documentList'].length>=3">
    <div class="container">
      <div class="row mb-4">
        <div class="col-sm-6">
          <h2 class="posts-entry-title">{{ $localedText($props['topic'].title) }}</h2>
        </div>
        <div class="col-sm-6 text-sm-end"><router-link :to="{name: 'Topic', params: {tid: topic.id}}" class="read-more">{{ $t('view_all') }}</router-link></div>
      </div>
      <div class="row g-3">
        <div class="col-md-9">
          <div class="row g-3">
            <div class="col-md-6">
              <div class="blog-entry">
                <router-link :to="{name: 'Document', params: {tid: doc0.topic.id, did: doc0.id}}" class="img-link">
                  <img :src="$calcDocumentEntryImgUrl(doc0, doc0.topic.id, '//placehold.co/700x700/6f42c1/90A1A9?text='+$localedText(doc0.id).replaceAll(' ','%20'))" class="img-fluid">
                </router-link>
                <span class="date">{{ $unixTimestampToReadable(doc0.tu) }}</span>
                <h2><router-link :to="{name: 'Document', params: {tid: doc0.topic.id, did: doc0.id}}">{{ $localedText(doc0.title) }}</router-link></h2>
                <p>{{ $localedText(doc0.summary) }}</p>
                <p><router-link :to="{name: 'Document', params: {tid: doc0.topic.id, did: doc0.id}}" class="btn btn-sm btn-outline-primary">{{ $t('read') }}</router-link></p>
              </div>
            </div>
            <div class="col-md-6">
              <div class="blog-entry">
                <router-link :to="{name: 'Document', params: {tid: doc1.topic.id, did: doc1.id}}" class="img-link">
                  <img :src="$calcDocumentEntryImgUrl(doc1, doc1.topic.id, '//placehold.co/700x700/6610f2/90A1A9?text='+$localedText(doc1.id).replaceAll(' ','%20'))" class="img-fluid">
                </router-link>
                <span class="date">{{ $unixTimestampToReadable(doc1.tu) }}</span>
                <h2><router-link :to="{name: 'Document', params: {tid: doc1.topic.id, did: doc1.id}}">{{ $localedText(doc1.title) }}</router-link></h2>
                <p>{{ $localedText(doc1.summary) }}</p>
                <p><router-link :to="{name: 'Document', params: {tid: doc1.topic.id, did: doc1.id}}" class="btn btn-sm btn-outline-primary">{{ $t('read') }}</router-link></p>
              </div>
            </div>
          </div>
        </div>
        <div class="col-md-3">
          <ul class="list-unstyled blog-entry-sm">
            <li v-for="doc in docOthers" v-bind:key="doc.id">
              <span class="date">{{ $unixTimestampToReadable(doc.tu) }}</span>
              <h3><router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}}">{{ $localedText(doc.title) }}</router-link></h3>
              <p>{{ $localedText(doc.summary) }}</p>
              <p><router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}}" class="read-more">{{ $t('read') }}</router-link></p>
            </li>
          </ul>
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
  name: 'lego-home-posts-entry-text-right',
  inject: ['$siteMeta'],
  props: ['topic', 'document-list'],
  computed: {
    doc0() {
      return this.$props['documentList'][0]
    },
    doc1() {
      return this.$props['documentList'][1]
    },
    docOthers() {
      return this.$props['documentList'].slice(2, 5)
    },
  }
}
</script>
