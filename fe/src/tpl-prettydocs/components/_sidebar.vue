<template>
  <div class="doc-sidebar col-md-3 col-12 order-0 d-none d-md-flex">
    <div id="doc-nav" class="doc-nav">
      <nav id="doc-menu" class="nav doc-menu flex-column sticky">
        <template v-for="topic in $siteTopics" v-bind:key="topic.id">
          <li :class="'nav-item'+($props['topicId']==topic.id?' active':'')">
            <router-link class="nav-link scrollto" :to="{name: 'Topic', params: {tid: topic.id}}">
              <fa-icon v-if="topic.icon!=''" :icon="topic.icon" class="fa-fw"/>{{ $localedText(topic.title) }}
            </router-link>
          </li>
          <nav v-if="topic.id==$props['topicId']" class="nav doc-sub-menu nav flex-column">
            <li v-for="doc in $props['documentList']" v-bind:key="doc.id" :class="'nav-item'+($props['documentId']==doc.id?' active':'')">
              <router-link class="nav-link scrollto" :to="{name: 'Document', params: {tid: topic.id, did: doc.id}}">
                <fa-icon v-if="doc.icon!=''" :icon="doc.icon" class="fa-fw"/>{{ $localedText(doc.title) }}
              </router-link>
            </li>
          </nav>
        </template>
      </nav>
    </div>
  </div>
</template>

<script>
export default {
  name: 'lego-sidebar',
  inject: ['$siteTopics'],
  props: ['topic-id', 'document-list', 'document-id'],
}
</script>
