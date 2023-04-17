<template>
  <div class="doc-sidebar col-md-3 col-12 order-0 d-none d-md-flex">
    <div id="doc-nav" class="doc-nav">
      <nav id="doc-menu" class="nav doc-menu flex-column sticky">
        <template v-for="topic in $siteTopics" v-bind:key="topic.id">
          <li :class="'nav-item'+($props['topicId']==topic.id?' active':'')">
            <router-link class="nav-link scrollto" :to="{name: 'Topic', params: {tid: topic.id}}">
              <i :class="'fa-fw '+(topic.icon!=''?topic.icon:'fas fa-square')" :style="topic.icon==''?'color: transparent !important;':''"/>
              {{ $localedText(topic.title) }}
            </router-link>
          </li>
          <nav v-if="topic.id==$props['topicId']" class="nav doc-sub-menu nav flex-column">
            <li v-for="doc in $props['documentList']" v-bind:key="doc.id" :class="'nav-item'+($props['documentId']==doc.id?' active':'')">
              <router-link class="nav-link scrollto" :to="{name: 'Document', params: {tid: topic.id, did: doc.id}}">
                <i :class="'fa-fw '+(doc.icon!=''?topic.icon:'fas fa-square')" :style="doc.icon==''?'color: transparent !important;':''"/>
                {{ $localedText(doc.title) }}
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
