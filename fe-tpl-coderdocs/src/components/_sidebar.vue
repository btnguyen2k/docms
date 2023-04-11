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
              <span class="theme-icon-holder me-2"><i v-if="topic.icon!=''" :class="topic.icon"></i></span>{{ $localedText(topic.title) }}
            </router-link>
          </li>
          <template v-if="topic.id==$props['topicId']">
            <li v-for="document in $props['documentList']" v-bind:key="document.id" class="nav-item">
              <router-link :class="'nav-link scrollto'+(document.id==$props['documentId']?' active':'')" :to="{name: 'Document', params: {tid: topic.id, did: document.id}}">
                <i v-if="document.icon!=''" :class="document.icon"></i> {{ $localedText(document.title) }}
              </router-link>
            </li>
          </template>
        </template>
      </ul>
    </nav>
  </div>
</template>

<script>
export default {
  name: 'lego-sidebar',
  inject: ['$global', '$siteTopics'],
  props: ['topic-id', 'document-list', 'document-id'],
  mounted() {
    this.$global.sidebar = this.$refs['docs-sidebar']
  }
}
</script>
