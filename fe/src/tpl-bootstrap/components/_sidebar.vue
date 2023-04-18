<template>
  <div class="col-lg-4">
    <div id="sidebar" class="sticky-top">
      <form v-if="$props['noSearch']===undefined" class="form-inline _d-none _d-sm-none _d-md-none _d-lg-inline" @submit.prevent="$doSearch" method="get">
        <div class="input-group pb-4">
          <input type="text" :placeholder="$t('search_prompt')" name="q" class="form-control" v-model="$global.searchQuery">
          <button type="submit" class="btn btn-primary" :value="$t('search')"><i class="bi bi-search" /></button>
        </div>
      </form>

      <div class="card mb-4">
        <div class="card-header">{{ $t('topics') }}</div>
        <div class="card-body">
          <div class="row">
            <nav class="nav flex-column">
              <template v-for="topic in $siteTopics" v-bind:key="topic.id">
                <router-link :to="{name: 'Topic', params: {tid: topic.id}}"
                             :class="'nav-link'+($props['topicId']==topic.id?' active':'')" style="padding-top: 0px !important; padding-bottom: 0px !important;">
                  <i :class="'fa-fw '+(topic.icon!=''?topic.icon:'fas fa-square')" :style="topic.icon==''?'color: transparent !important;':''" />
                  {{ $localedText(topic.title) }}
                </router-link>
                <nav v-if="$props['topicId']==topic.id" class="nav flex-column ps-4">
                  <router-link v-for="document in $props['documentList']" v-bind:key="document.id"
                               :to="{name: 'Document', params: {tid: topic.id, did: document.id}}"
                               :class="'nav-link'+($props['documentId']?' active':'')" style="padding-top: 0px !important; padding-bottom: 0px !important;">
                    <i :class="'fa-fw '+(document.icon!=''?document.icon:'fas fa-square')" :style="document.icon==''?'color: transparent !important;':''" />
                    {{ $localedText(document.title) }}
                  </router-link>
                </nav>
              </template>
            </nav>
          </div>
        </div>
      </div>
      <!--<div class="card mb-4">-->
      <!--  <div class="card-header">Side Widget</div>-->
      <!--  <div class="card-body">You can put anything you want inside of these side widgets. They are easy to use, and feature the Bootstrap 5 card component!</div>-->
      <!--</div>-->
    </div>
  </div>
</template>

<style>
#sidebar.sticky-top {
  top: 60px !important;
}
</style>

<script>
export default {
  name: 'lego-page-header',
  inject: ['$global', '$siteTopics'],
  props: ['topic-id', 'document-list', 'document-id', 'no-search'],
}
</script>
