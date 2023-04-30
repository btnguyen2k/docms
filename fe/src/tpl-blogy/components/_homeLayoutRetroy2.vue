<template>
  <!-- Start retroy layout blog posts/style 2 -->
  <div class="section bg-light" v-if="$props['documentList'] && $props['documentList'].length>=4">
    <div class="container">
      <!--      <div class="row mb-4">-->
      <!--        <div class="col-sm-6">-->
      <!--          <h2 class="posts-entry-title">Travel</h2>-->
      <!--        </div>-->
      <!--        <div class="col-sm-6 text-sm-end"><a href="category.html" class="read-more">View All</a></div>-->
      <!--      </div>-->
      <div class="row align-items-stretch retro-layout-alt">
        <div class="col-md-5 order-md-2">
          <router-link :to="{name: 'Document', params: {tid: doc3.topic.id, did: doc3.id}}" class="hentry img-1 h-100 gradient">
            <div class="featured-img" :style="'background-image: url('+($calcDocumentEntryImgUrl(doc3, doc3.topic.id, '//placehold.co/700x700/214252/90A1A9?text='+$localedText(doc3.topic.title).replaceAll(' ','%20')))+');'"></div>
            <div class="text">
              <span>{{ $unixTimestampToReadable(doc3.tu) }}</span>
              <h2>{{ $localedText(doc3.title) }}</h2>
            </div>
          </router-link>
        </div>
        <div class="col-md-7">
          <router-link :to="{name: 'Document', params: {tid: doc0.topic.id, did: doc0.id}}" class="hentry img-2 v-height mb30 gradient">
            <div class="featured-img" :style="'background-image: url('+($calcDocumentEntryImgUrl(doc0, doc0.topic.id, '//placehold.co/700x440/0d6efd/90A1A9?text='+$localedText(doc0.topic.title).replaceAll(' ','%20')))+');'"></div>
            <div class="text text-sm">
              <span>{{ $unixTimestampToReadable(doc0.tu) }}</span>
              <h2>{{ $localedText(doc0.title) }}</h2>
            </div>
          </router-link>
          <div class="two-col d-block d-md-flex justify-content-between">
            <router-link :to="{name: 'Document', params: {tid: doc1.topic.id, did: doc1.id}}" class="hentry v-height img-2 gradient">
              <div class="featured-img" :style="'background-image: url('+($calcDocumentEntryImgUrl(doc1, doc1.topic.id, '//placehold.co/700x700/6610f2/90A1A9?text='+$localedText(doc1.topic.title).replaceAll(' ','%20')))+');'"></div>
              <div class="text text-sm">
                <span>{{ $unixTimestampToReadable(doc1.tu) }}</span>
                <h2>{{ $localedText(doc1.title) }}</h2>
              </div>
            </router-link>
            <router-link :to="{name: 'Document', params: {tid: doc2.topic.id, did: doc2.id}}" class="hentry v-height img-2 ms-auto float-end gradient">
              <div class="featured-img" :style="'background-image: url('+($calcDocumentEntryImgUrl(doc2, doc2.topic.id, '//placehold.co/700x700/6f42c1/90A1A9?text='+$localedText(doc2.topic.title).replaceAll(' ','%20')))+');'"></div>
              <div class="text text-sm">
                <span>{{ $unixTimestampToReadable(doc2.tu) }}</span>
                <h2>{{ $localedText(doc2.title) }}</h2>
              </div>
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- End retroy layout blog posts -->
</template>

<script>
export default {
  name: 'lego-home-layout-retroy2',
  inject: ['$siteMeta'],
  props: ['document-list'],
  computed: {
    doc0() {
      return this.$props['documentList'][0]
    },
    doc1() {
      return this.$props['documentList'][1]
    },
    doc2() {
      return this.$props['documentList'][2]
    },
    doc3() {
      return this.$props['documentList'][3]
    }
  }
}
</script>
