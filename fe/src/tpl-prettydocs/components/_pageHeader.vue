<template>
  <header id="header" class="header">
    <div class="container">
      <div class="branding">
        <h1 class="logo">
          <router-link :to="{ name: 'Home' }">
            <fa-icon v-if="$siteMeta.icon!=''" aria-hidden="true" class="icon" style="padding-right: 8px" :icon="$siteMeta.icon" />
            <span class="text-highlight">{{ $siteFirstName }}</span><span class="text-bold">{{ $siteLastName }}</span>
          </router-link>
        </h1>
      </div>

      <ol class="breadcrumb">
        <li class="breadcrumb-item"><router-link :to="{name: 'Home'}">{{ $t('home') }}</router-link></li>

        <li class="breadcrumb-item active" v-if="$props['search']!==undefined">{{ $t('search') }}</li>

        <li class="breadcrumb-item active" v-if="$props['search']===undefined && $props['document']===undefined">{{ $localedText($props['topic'].title) }}</li>

        <li class="breadcrumb-item" v-if="$props['search']===undefined && $props['document']!==undefined"><router-link :to="{name: 'Topic', params: {tid: topic.id}}">{{ $localedText($props['topic'].title) }}</router-link></li>
        <li class="breadcrumb-item active" v-if="$props['search']===undefined && $props['document']!==undefined">{{ $localedText($props['document'].title) }}</li>
      </ol>

      <div class="top-search-box">
        <form class="form-inline search-form justify-content-center" @submit.prevent="$doSearch" method="get">
          <input type="text" :placeholder="$t('search')" name="q" class="form-control search-input" v-model="$global.searchQuery">
          <button type="submit" class="btn search-btn" :value="$t('search')"><fa-icon icon="fas fa-search"></fa-icon></button>
        </form>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'lego-page-header',
  inject: ['$global', '$siteMeta', '$siteFirstName', '$siteLastName'],
  props: ['search', 'topic', 'document'],
}
</script>
