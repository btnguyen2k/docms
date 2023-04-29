<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">{{ errorMsg }}</div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else-if="!topic.id" class="alert alert-danger m-4" role="alert">{{ $t('error_topic_not_found', {topic: $route.params.tid}) }}</div>
  <div v-else>
    <lego-page-header active="topic" :topic="topic" />

    <div class="section search-result-wrap">
      <div class="container">
        <div class="row">
          <div class="col-12">
            <div class="heading">{{ $t('topic')}}: {{ $localedText(topic.title) }}</div>
          </div>
        </div>
        <div class="row posts-entry">
          <div class="col-lg-8">
            <div class="blog-entry d-flex blog-entry-search-item" v-for="doc in documentList" v-bind:key="doc.id">
              <router-link :to="{name: 'Document', params: {tid: topic.id, did: doc.id}}" class="img-link me-4">
                <img :src="$calcDocumentEntryImgUrl(doc, topic.id, '//placehold.co/440x440/214252/90A1A9?text='+$localedText(doc.id).replaceAll(' ','%20'))" class="img-fluid">
              </router-link>
              <div class="col-9">
                <span class="date">{{ $unixTimestampToReadable(doc.tu) }}</span>
                <h2><router-link :to="{name: 'Document', params: {tid: topic.id, did: doc.id}}">{{ $localedText(doc.title) }}</router-link></h2>
                <p>{{ $localedText(doc.summary) }}</p>
                <p><router-link :to="{name: 'Document', params: {tid: topic.id, did: doc.id}}" class="btn btn-sm btn-outline-primary">{{ $t('read') }}</router-link></p>
              </div>
            </div>
<!--            <div class="row text-start pt-5 border-top">-->
<!--              <div class="col-md-12">-->
<!--                <div class="custom-pagination">-->
<!--                  <span>1</span>-->
<!--                  <a href="#">2</a>-->
<!--                  <a href="#">3</a>-->
<!--                  <a href="#">4</a>-->
<!--                  <span>...</span>-->
<!--                  <a href="#">15</a>-->
<!--                </div>-->
<!--              </div>-->
<!--            </div>-->
          </div>

          <div class="col-lg-4 sidebar">
            <div class="sidebar-box search-form-wrap mb-4">
              <form class="sidebar-search-form" @submit.prevent="$doSearch" method="get">
                <span class="bi-search"></span>
                <input type="text" class="form-control" id="s" :placeholder="$t('search_prompt')" v-model="$global.searchQuery">
              </form>
            </div>

<!--            <div class="sidebar-box">-->
<!--              <h3 class="heading">Popular Posts</h3>-->
<!--              <div class="post-entry-sidebar">-->
<!--                <ul>-->
<!--                  <li>-->
<!--                    <a href="">-->
<!--                      <img src="images/img_1_sq.jpg" alt="Image placeholder" class="me-4 rounded">-->
<!--                      <div class="text">-->
<!--                        <h4>There’s a Cool New Way for Men to Wear Socks and Sandals</h4>-->
<!--                        <div class="post-meta">-->
<!--                          <span class="mr-2">March 15, 2018 </span>-->
<!--                        </div>-->
<!--                      </div>-->
<!--                    </a>-->
<!--                  </li>-->
<!--                  <li>-->
<!--                    <a href="">-->
<!--                      <img src="images/img_2_sq.jpg" alt="Image placeholder" class="me-4 rounded">-->
<!--                      <div class="text">-->
<!--                        <h4>There’s a Cool New Way for Men to Wear Socks and Sandals</h4>-->
<!--                        <div class="post-meta">-->
<!--                          <span class="mr-2">March 15, 2018 </span>-->
<!--                        </div>-->
<!--                      </div>-->
<!--                    </a>-->
<!--                  </li>-->
<!--                  <li>-->
<!--                    <a href="">-->
<!--                      <img src="images/img_3_sq.jpg" alt="Image placeholder" class="me-4 rounded">-->
<!--                      <div class="text">-->
<!--                        <h4>There’s a Cool New Way for Men to Wear Socks and Sandals</h4>-->
<!--                        <div class="post-meta">-->
<!--                          <span class="mr-2">March 15, 2018 </span>-->
<!--                        </div>-->
<!--                      </div>-->
<!--                    </a>-->
<!--                  </li>-->
<!--                </ul>-->
<!--              </div>-->
<!--            </div>-->

            <div class="sidebar-box">
              <h3 class="heading">{{ $t('topics') }}</h3>
              <ul class="categories">
                <li v-for="topic in $siteTopics" v-bind:key="topic.id">
                  <router-link :to="{name: 'Topic', params:{tid: topic.id}}">{{ $localedText(topic.title) }} <span>({{ topic.num_docs }})</span></router-link>
                </li>
              </ul>
            </div>

            <div class="sidebar-box">
              <h3 class="heading">{{ $t('tag_cloud') }}</h3>
              <ul class="tags" style="font-size: small">
                <li v-for="(_, tag) in $localedText($tagCloud)" v-bind:key="tag">
                  <router-link :to="{name: 'TagSearch', query:{q: tag, l: $i18n.locale}}" :class="calcTagCloudCSS(tag)">{{ tag }}</router-link>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>

    <lego-page-footer :document-list="$latestDocuments" />
  </div>
</template>

<script>
import {useRoute} from 'vue-router'
import {watch} from 'vue'
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import {APP_CONFIG} from '@/_shared/utils/app_config'

export default {
  name: 'Topic',
  inject: ['$global', '$siteTopics', '$tagCloud', '$latestDocuments'],
  components: {legoPageHeader, legoPageFooter},
  mounted() {
    const vue = this
    const route = useRoute()
    watch(
        () => route.params.tid,
        async newTid => {
          if (newTid) {
            vue._fetchTopics(vue, newTid)
          }
        }
    )
    this._fetchSiteMeta(this)
  },
  methods: {
    calcTagCloudCSS(tag) {
      const cssList = [
        'bg-primary link-light',
        'bg-secondary link-light',
        'bg-success link-light',
        'bg-danger link-light',
        'bg-warning text-dark link-dark',
        'bg-info link-light',
        'bg-light text-dark link-dark',
      ]
      return this.$pickupFromHash(tag, cssList)
    },
    _fetchSiteMeta(vue) {
      vue.$fetchSiteMeta(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue._fetchTopics(vue, vue.$route.params.tid)
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _fetchTopics(vue, topicId) {
      vue.$fetchSiteTopics(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.$siteTopics.forEach(t => {
                if (t.id == topicId) {
                  vue.topic = t
                  const appNameAndVersion = APP_CONFIG.app.name + ' v' + APP_CONFIG.app.version
                  document.title = vue.$localedText(vue.topic.title) + ' | ' + appNameAndVersion
                  vue._fetchDocuments(vue, topicId)
                }
              })
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _fetchDocuments(vue, topicId) {
      vue.$fetchDocumentsForTopic(topicId,
          () => vue.status=0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.documentList = apiResp.data
            } else {
              vue.errorMsg = vue.status+": "+apiResp.message
            }
          },
          err => vue.errorMsg = err
      )
    },
  },
  data() {
    return {
      topic: {},
      documentList: [],
      status: -1,
      errorMsg: '',
    }
  },
}
</script>
