<template>
  <div class="col-md-12 col-lg-4 sidebar">
    <div class="sidebar-box search-form-wrap">
      <form class="sidebar-search-form" @submit.prevent="$doSearch" method="get" style="position: relative;">
        <input type="text" class="form-control" id="s" :placeholder="$t('search_prompt')" v-model="$global.searchQuery">
        <span class="bi-search"></span>
      </form>
    </div>

    <!--            <div class="sidebar-box">-->
    <!--              <div class="bio text-center">-->
    <!--                <img src="images/person_2.jpg" alt="Image Placeholder" class="img-fluid mb-3">-->
    <!--                <div class="bio-body">-->
    <!--                  <h2>Hannah Anderson</h2>-->
    <!--                  <p class="mb-4">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Exercitationem facilis sunt repellendus excepturi beatae porro debitis voluptate nulla quo veniam fuga sit molestias minus.</p>-->
    <!--                  <p><a href="#" class="btn btn-primary btn-sm rounded px-2 py-2">Read my bio</a></p>-->
    <!--                  <p class="social">-->
    <!--                    <a href="#" class="p-2"><span class="fa fa-facebook"></span></a>-->
    <!--                    <a href="#" class="p-2"><span class="fa fa-twitter"></span></a>-->
    <!--                    <a href="#" class="p-2"><span class="fa fa-instagram"></span></a>-->
    <!--                    <a href="#" class="p-2"><span class="fa fa-youtube-play"></span></a>-->
    <!--                  </p>-->
    <!--                </div>-->
    <!--              </div>-->
    <!--            </div>-->

    <div class="sidebar-box" v-if="$props['documentList'] && $props['documentList'].length > 0">
      <h3 class="heading" v-if="$props['documentListTitle']">{{ $props['documentListTitle'] }}</h3>
      <div class="post-entry-sidebar">
        <ul>
          <li v-for="doc in $props['documentList']" v-bind:key="doc.id" data-aos="fade-up">
            <router-link :to="{name: 'Document', params: {tid: doc.topic?doc.topic.id:$props['topic'].id, did: doc.id}}">
              <img data-aos="zoom-in" :src="$calcDocumentEntryImgUrl(doc, doc.topic?doc.topic.id:$props['topic'].id, '//placehold.co/440x440/214252/90A1A9?text='+$localedText(doc.id).replaceAll(' ','%20'), 's')" class="me-4 rounded">
              <div class="text">
                <h4>{{ $localedText(doc.title) }}</h4>
                <div class="post-meta">
                  <span class="mr-2">{{ $unixTimestampToReadable(doc.tu) }}</span>
                </div>
              </div>
            </router-link>
          </li>
        </ul>
      </div>
    </div>

    <div class="sidebar-box">
      <h3 class="heading">{{ $t('topics') }}</h3>
      <ul class="categories">
        <li v-for="topic in visibleTopics" v-bind:key="topic.id">
          <router-link :to="{name: 'Topic', params:{tid: topic.id}}">
            <i v-if="topic.icon!=''" :class="topic.icon+' fa-fw'"></i> {{ $localedText(topic.title) }} <span>({{ topic.num_docs }})</span>
          </router-link>
        </li>
      </ul>
    </div>

    <div class="sidebar-box">
      <h3 class="heading">{{ $t('tag_cloud') }}</h3>
      <ul class="tags" style="font-size: small">
        <li v-for="(_, tag) in $localedText($tagCloud)" v-bind:key="tag" :data-aos="randomAOS()" :data-aos-delay="randomNumber(100,1000)" :data-aos-duration="randomNumber(100,1000)">
          <router-link :to="{name: 'TagSearch', query:{q: tag, l: $i18n.locale}}" :class="calcTagCloudCSS(tag)">{{ tag }}</router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'lego-page-header',
  inject: ['$global', '$tagCloud', '$siteTopics'],
  props: ['document-list', 'document-list-title', 'topic'],
  computed: {
    visibleTopics() {
      const result = []
      for (let i = 0; i < this.$siteTopics.length; i++) {
        if (!this.$siteTopics[i].hidden) {
          result.push(this.$siteTopics[i])
        }
      }
      return result
    },
  },
  methods: {
    randomNumber(min, max) {
      return min + Math.floor(Math.random() * (max-min))
    },
    randomAOS() {
      const arr = ['fade-up', 'fade-down', 'fade-right', 'fade-left', 'fade-up-right', 'fade-up-left', 'fade-down-right', 'fade-down-left']
      const i = this.randomNumber(0, arr.length)
      return arr[i]
    },
    calcTagCloudCSS(tag) {
      const cssList = [
        'bg-primary link-light',
        'bg-secondary link-light',
        'bg-success link-light',
        'bg-danger link-light',
        'bg-warning text-dark link-dark',
        'bg-info link-light',
        // 'bg-light text-dark link-dark',
      ]
      return this.$pickupFromHash(tag, cssList)
    },
  },
}
</script>
