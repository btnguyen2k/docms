<template>
  <div ref="docs-sidebar" id="docs-sidebar" class="docs-sidebar">
    <div class="top-search-box d-lg-none p-3">
      <form class="search-form" @submit.prevent="$doSearch">
        <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control search-input" v-model="$global.searchQuery">
        <button type="submit" class="btn search-btn" :value="$t('search')"><i class="fas fa-search"></i></button>
      </form>
    </div>

    <nav id="docs-nav" class="docs-nav navbar">
      <ul class="section-items list-unstyled nav flex-column pb-3">
        <template v-for="topic in $siteTopics" v-bind:key="topic.id">
          <li class="nav-item section-title">
            <router-link :class="'nav-link scrollto'+(topic.id==$props['topicId']?' active':'')" :to="{name: 'Topic', params: {tid: topic.id}}">
              <span class="theme-icon-holder me-2"><i :class="topic.icon!=''?topic.icon:'fas fa-square'" :style="topic.icon==''?'color: transparent !important;':''"></i></span>{{ $localedText(topic.title) }}
            </router-link>
          </li>
          <template v-if="topic.id==$props['topicId']">
            <li v-for="document in $props['documentList']" v-bind:key="document.id" class="nav-item">
              <router-link :class="'nav-link scrollto'+(document.id==$props['documentId']?' active':'')" :to="{name: 'Document', params: {tid: topic.id, did: document.id}}">
                <i v-if="document.icon!=''" :class="'fa-fw '+document.icon"></i> {{ $localedText(document.title) }}
              </router-link>
            </li>
          </template>
        </template>
      </ul>
    </nav>

    <div class="m-2">
      <ul class="list-unstyled mb-0">
        <li v-for="(_, tag) in $localedText($tagCloud)" v-bind:key="tag" class="d-inline-block me-1 mb-1">
          <router-link :to="{name: 'TagSearch', query:{q: tag, l: $i18n.locale}}" :class="$calcTagCloudCSS(tag)" :style="calcTagCloudSize(tag)">{{ tag }}</router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'lego-page-footer',
  inject: ['$global', '$siteTopics', '$tagCloud', '$calcTagCloudCSS'],
  props: ['topic-id', 'document-list', 'document-id'],
  mounted() {
    this.$global.sidebar = this.$refs['docs-sidebar']
  },
  methods: {
    calcTagCloudSize(tag) {
      const size = this.$calcTagSize(tag, 0.55, 1.15, 5)
      return 'font-size: ' + size + 'rem !important;'
    }
  }
}
</script>
