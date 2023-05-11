<template>
  <router-view></router-view>
</template>

<style lang="css">
/* Google fonts */
@import "//fonts.googleapis.com/css2?family=Work+Sans:wght@400;600;700&display=swap";

/* FontAwesome */
@import "@/_shared/assets/fontawesome-free-6.4.0-web/css/all.min.css";

/* Bootstrap 5.x */
@import "bootstrap/dist/css/bootstrap.min.css";
@import "bootstrap";
@import "bootstrap-icons";

@import "@/tpl-blogy/assets/fonts/icomoon/style.css";
@import "@/tpl-blogy/assets/fonts/flaticon/font/flaticon.css";
@import "@/tpl-blogy/assets/css/style.css";
</style>

<script>
import {computed, getCurrentInstance} from 'vue'
import {triggerPopstate, triggerResize} from '@/_shared/utils/docms_utils'
import router from '@/tpl-blogy/router'

window.addEventListener('popstate', () => triggerPopstate())
window.addEventListener('resize', () => triggerResize())

export default {
  name: 'App',
  inject: ['$global'],
  provide() {
    return {
      $latestDocuments: computed(() => this.latestDocuments),
    }
  },
  mounted() {
    const app = getCurrentInstance()
    app.appContext.config.globalProperties.$unixTimestampToReadable = this.unixTimestampToReadable
    app.appContext.config.globalProperties.$calcDocumentEntryImgUrl = this.calcDocumentEntryImgUrl
    app.appContext.config.globalProperties.$calcAuthorAvatarUrl = this.calcAuthorAvatarUrl

    const vue = this
    document.addEventListener('click', function (event) {
      const specifiedElement = document.querySelector(".site-mobile-menu")
      const btnSidebarOpen = document.querySelector("#sidebarOpen")
      if (!specifiedElement || !btnSidebarOpen) {
        return
      }
      const isClickInside = specifiedElement.contains(event.target)
      const mt = btnSidebarOpen.contains(event.target)
      if (!isClickInside && !mt) {
        document.body.classList.remove('offcanvas-menu')
        document.querySelector("#sidebarOpen").classList.remove('active')
        const el = document.querySelector("#collapseTogglerLanguages")
        if (el && !el.classList.contains('collapsed')) {
          el.click()
        }
      }
    });

    vue._fetchLatestDocuments(vue)
  },
  methods: {
    _fetchLatestDocuments(vue) {
      vue.$fetchLatestDocuments(undefined, 10,
          () => {
          },
          apiResp => {
            if (apiResp.status == 200) {
              vue.latestDocuments = apiResp.data
            }
          },
          () => {
          },
      )
    },
    unixTimestampToReadable(timeInSeconds) {
      const d = new Date(timeInSeconds * 1000)
      const yyyy = d.getFullYear()
      const mm = ('0' + (d.getMonth() + 1)).slice(-2)  // Months are zero based. Add leading 0.
      const dd = ('0' + d.getDate()).slice(-2)         // Add leading 0.
      const hh = ('0' + d.getHours()).slice(-2)         // Add leading 0.
      const min = ('0' + d.getMinutes()).slice(-2)     // Add leading 0.
      return yyyy + '-' + mm + '-' + dd + ', ' + hh + ':' + min
    },
    calcDocumentEntryImgUrl(doc, topicId, defaultUrl, selection) {
      let img = doc.img
      if (img && selection) {
        const imgMap = JSON.parse(doc.img)
        img = imgMap && imgMap[selection] ? imgMap[selection] : ''
      }
      if (img != '') {
        const reAbsUrl = /^([a-z]+:)?\//i
        if (reAbsUrl.test(img)) {
          return img
        }
        const base = router.resolve({name: 'Document', params: {tid: topicId, did: doc.id}}).href
        return base + img
      }
      return defaultUrl
    },
    calcAuthorAvatarUrl(auth, defaultUrl) {
      if (auth && auth.avatar) {
        return auth.avatar
      }
      if (auth && auth.email) {
        return '//www.gravatar.com/avatar/' + auth.email.toLowerCase().md5()
      }
      return defaultUrl
    }
  },
  data() {
    return {
      latestDocuments: [],
    }
  },
}
</script>
