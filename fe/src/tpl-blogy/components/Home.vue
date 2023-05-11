<template>
  <div v-if="errorMsg!=''" class="alert alert-danger m-4" role="alert">
    {{ errorMsg }}
    <hr/>
    <p class="btn btn-outline-primary mb-0" @click="$reload()"><i class="bi bi-arrow-clockwise"></i> {{ $t('reload') }}</p>
  </div>
  <div v-else-if="status<=0" class="alert alert-info m-4" role="alert"><i class="bi bi-hourglass"></i> {{ $t('wait') }}</div>
  <div v-else>
    <lego-page-header active="home"/>

    <lego-home-layout-retroy2 :document-list="$latestDocuments.slice(0,4)"/>

    <template v-for="(topic, index) in largeTopics()" v-bind:key="topic.id">
      <lego-home-posts-entry-text-right :topic="topic" :document-list="latestDocsPerTopic[topic.id]" v-if="index%3==0"/>
      <lego-home-posts-entry-sm :topic="topic" :document-list="latestDocsPerTopic[topic.id]" v-if="index%3==1"/>
      <lego-home-posts-entry-text-left :topic="topic" :document-list="latestDocsPerTopic[topic.id]" v-if="index%3==2"/>
    </template>

    <lego-home-posts-entry-med :document-list="latestDocsSmallTopics" />

    <lego-page-footer :document-list="$latestDocuments.slice(4)" />
  </div>
</template>

<script>
import LegoPageFooter from "./_pageFooter.vue"
import LegoHomeLayoutRetroy2 from "./_homeLayoutRetroy2.vue"
import LegoHomePostsEntryMed from "./_homePostsEntryMed.vue"
import LegoHomePostsEntryTextLeft from "./_homePostsEntryTextLeft.vue"
import LegoHomePostsEntrySm from "./_homePostsEntrySm.vue"
import LegoHomePostsEntryTextRight from "./_homePostsEntryTextRight.vue"
import LegoHomeLayoutRetroy1 from "./_homeLayoutRetroy1.vue"
import LegoPageHeader from "./_pageHeader.vue"

const largeTopicThreshold = 4

export default {
  name: 'Home',
  // eslint-disable-next-line vue/no-unused-components
  components: {LegoPageFooter, LegoHomeLayoutRetroy2, LegoHomePostsEntryMed, LegoHomePostsEntryTextLeft, LegoHomePostsEntrySm, LegoHomePostsEntryTextRight, LegoHomeLayoutRetroy1, LegoPageHeader},
  inject: ['$global', '$siteTopics', '$latestDocuments'],
  mounted() {
    this._fetchSiteMeta(this)
  },
  methods: {
    largeTopics() {
      let result = []
      for (let i = 0; i < this.$siteTopics.length; i++) {
        if (this.$siteTopics[i].num_docs >= largeTopicThreshold) {
          result.push(this.$siteTopics[i])
        }
      }
      return result
    },
    _fetchSiteMeta(vue) {
      vue.$fetchSiteMeta(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              vue.$updatePageTitle()
              vue._fetchTopics(vue)
            } else {
              vue.errorMsg = vue.status + ": " + apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _fetchTopics(vue) {
      vue.$fetchSiteTopics(
          () => vue.status = 0,
          apiResp => {
            vue.status = apiResp.status
            if (vue.status == 200) {
              let smallTopics = ''
              for (const topic of vue.$siteTopics) {
                if (topic.num_docs >= largeTopicThreshold) {
                  vue._fetchLatestDocumentsForTopic(vue, topic)
                } else {
                  smallTopics += topic.id+","
                }
              }
              vue._fetchLatestDocumentsForSmallTopics(vue, smallTopics)
            } else {
              vue.errorMsg = vue.status + ": " + apiResp.message
            }
          },
          err => {
            vue.errorMsg = err
          },
      )
    },
    _fetchLatestDocumentsForSmallTopics(vue, topicId) {
      vue.$fetchLatestDocuments(topicId, 10,
          () => {
          },
          (apiResp) => {
            if (apiResp.status == 200) {
              vue.latestDocsSmallTopics = apiResp.data
            }
          },
          () => {
          },
      )
    },
    _fetchLatestDocumentsForTopic(vue, topic) {
      vue.$fetchLatestDocuments(topic.id, 10,
          () => {
          },
          (apiResp) => {
            if (apiResp.status == 200) {
              vue.latestDocsPerTopic[topic.id] = apiResp.data
            }
          },
          () => {
          },
      )
    },
  },
  data() {
    return {
      status: -1,
      errorMsg: '',
      latestDocsPerTopic: {},
      latestDocsSmallTopics: [],
    }
  },
}
</script>
