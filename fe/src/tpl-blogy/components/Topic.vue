<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">
    {{ errorMsg }}
    <hr/>
    <p class="btn btn-outline-primary mb-0" @click="$reload()"><i class="bi bi-arrow-clockwise"></i> {{ $t('reload') }}</p>
  </div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert"><i class="bi bi-hourglass"></i> {{ $t('wait') }}</div>
  <div v-else-if="!topic.id" class="alert alert-danger m-4" role="alert">
    {{ $t('error_topic_not_found', {topic: $route.params.tid}) }}
    <hr/>
    <span v-html="$t('transfer_to_home', {url: $router.resolve({name: 'Home'}).href})"></span>
  </div>
  <div v-else>
    <lego-page-header active="topic" :topic="topic" />

    <div class="section search-result-wrap">
      <div class="container">
        <div class="row">
          <div class="col-12">
            <div class="heading">
              <span class=" fw-bold"><i v-if="topic.icon!=''" :class="topic.icon"/> {{ $localedText(topic.title) }}</span>
              <br/><span>{{ $localedText(topic.description) }}</span>
            </div>
          </div>
        </div>
        <div class="row posts-entry">
          <div class="col-lg-8">
            <div class="blog-entry d-flex blog-entry-search-item" v-for="doc in documentList" v-bind:key="doc.id" data-aos="fade-up">
              <router-link :to="{name: 'Document', params: {tid: topic.id, did: doc.id}}" class="img-link me-4">
                <img data-aos="zoom-in" :src="$calcDocumentEntryImgUrl(doc, topic.id, '//placehold.co/440x440/214252/90A1A9?text='+$localedText(doc.id).replaceAll(' ','%20'), 's')" class="img-fluid">
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

          <sidebar />
        </div>
      </div>
    </div>

    <lego-page-footer :document-list="latestDocuments" />
  </div>
</template>

<script>
import {useRoute} from 'vue-router'
import {watch} from 'vue'
import legoPageHeader from './_pageHeader.vue'
import legoPageFooter from './_pageFooter.vue'
import sidebar from './_sidebar.vue'
import {switchLanguage} from '@/_shared/i18n'

export default {
  name: 'Topic',
  inject: ['$global', '$siteTopics', '$latestDocuments'],
  components: {legoPageHeader, legoPageFooter, sidebar},
  mounted() {
    const vue = this
    if (vue.$route.query.l) {
      switchLanguage(vue.$route.query.l, false)
    }
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
  computed: {
    latestDocuments() {
      const result = []
      for (let i = 0; i < this.$latestDocuments.length; i++) {
        let found = false
        for (let j = 0; j < this.documentList.length; j++) {
          if (this.$latestDocuments[i].id == this.documentList[j].id) {
            found = true
            break
          }
        }
        if (!found) {
          result.push(this.$latestDocuments[i])
        }
      }
      return result
    },
  },
  methods: {
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
                  vue.$updatePageTitle({topic: t})
                  vue.topic = t
                  vue._fetchDocuments(vue, topicId)
                }
              })
            } else {
              // vue.errorMsg = vue.status+": "+apiResp.message
            }
            if (!vue.topic.id) {
              vue.$transferToHome(3)
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
