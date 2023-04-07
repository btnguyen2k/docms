<template>
  <router-view></router-view>
</template>

<style lang="css">
/*!* Google Fonts *!*/
@import "//fonts.googleapis.com/css?family=Poppins:300,400,500,600,700&display=swap";

/* FontAwesome */
@import "@/assets/fontawesome-free-6.4.0-web/css/all.min.css";

/* Bootstrap 5.x */
/*@import "bootstrap/dist/css/bootstrap.min.css";*/
@import "bootstrap";
@import "bootstrap-icons";

/*!* Application's main style *!*/
@import "@/assets/css/theme.css";

.img-fit img {
  max-width: 95%;
}

.img-center img {
  margin-left: auto;
  margin-right: auto;
  display: block;
}
</style>

<script>
import {
  extractLeadingFromName, extractTrailingFromName,
  registerResize, unregisterResize,
  triggerPopstate, triggerResize,
} from '@/components/utils'

window.addEventListener('popstate', () => triggerPopstate())
window.addEventListener('resize', () => triggerResize())

class Global {
  get siteMeta() {
    return this._siteMeta ? this._siteMeta : {}
  }

  set siteMeta(value) {
    this._siteMeta = value
    this._siteLanguages = Object.keys(this._siteMeta.languages).filter(value => {
      return value != 'default'
    })
  }

  get siteTopics() {
    return this._siteTopics ? this._siteTopics : []
  }

  set siteTopics(value) {
    this._siteTopics = value
  }

  get siteLanguages() {
    return this._siteLanguages ? this._siteLanguages : []
  }

  get siteFirstName() {
    return this._siteMeta.name ? extractLeadingFromName(this._siteMeta.name) : ''
  }

  get siteLastName() {
    return this._siteMeta.name ? extractTrailingFromName(this._siteMeta.name) : ''
  }

  get searchQuery() {
    return this._searchQuery !== undefined ? this._searchQuery : ''
  }

  set searchQuery(value) {
    this._searchQuery = value
  }

  get sidebar() {
    return this._sidebar
  }
  set sidebar(value) {
    this._sidebar = value
  }
}

import {computed} from 'vue'

export default {
  name: 'App',
  unmounted() {
    unregisterResize(this.coderDocsResponsiveSidebar) // CoderDocs
  },
  mounted() {
    registerResize(this.coderDocsResponsiveSidebar) // CoderDocs: onresize
  },
  data() {
    return {
      global: new Global(),
    }
  },
  methods: {
    coderDocsResponsiveSidebar() { // CoderDocs
      const w = window.innerWidth
      const sidebar = this.global.sidebar
      if (sidebar) {
        if (w >= 1024) {
          sidebar.classList.remove('sidebar-hidden')
          sidebar.classList.add('sidebar-visible')
        } else {
          sidebar.classList.remove('sidebar-visible')
          sidebar.classList.add('sidebar-hidden')
        }
      }
    },
  },
  provide() {
    return {
      /*-- read/write global variable */
      $global: this.global,
      $searchQuery: computed(() => this.global.searchQuery),

      /*-- read-only global variable */
      $siteMeta: computed(() => this.global.siteMeta),
      $siteLanguages: computed(() => this.global.siteLanguages),
      $siteFirstName: computed(() => this.global.siteFirstName),
      $siteLastName: computed(() => this.global.siteLastName),
      $siteTopics: computed(() => this.global.siteTopics),

      /*-- global functions */
      $coderDocsResponsiveSidebar: this.coderDocsResponsiveSidebar,
    }
  }
}
</script>
