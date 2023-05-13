<template>
  <router-view></router-view>
</template>

<style lang="css">
/*!* Google Fonts *!*/
@import "//fonts.googleapis.com/css?family=Poppins:300,400,500,600,700&display=swap";

/* DevIcons */
@import "//cdn.jsdelivr.net/gh/devicons/devicon@latest/devicon.min.css";

/* FontAwesome */
@import "@/_shared/assets/fontawesome-free-6.4.0-web/css/all.min.css";

/* Bootstrap 5.x */
/*@import "bootstrap/dist/css/bootstrap.min.css";*/
@import "bootstrap";
@import "bootstrap-icons";

/*!* Application's main style *!*/
@import "./assets/css/theme.css";
</style>

<script>
import {
  registerResize, unregisterResize,
  triggerPopstate, triggerResize,
} from '@/_shared/utils/docms_utils'

window.addEventListener('popstate', () => triggerPopstate())
window.addEventListener('resize', () => triggerResize())

export default {
  name: 'App',
  inject: ['$global'],
  provide() {
    return {
      $coderDocsResponsiveSidebar: this.coderDocsResponsiveSidebar,
      $coderDocsSidebarToggler: this.coderDocsSidebarToggler,
      $calcTagCloudCSS: this.calcTagCloudCSS,
    }
  },
  unmounted() {
    unregisterResize(this.$coderDocsResponsiveSidebar) // CoderDocs
  },
  mounted() {
    registerResize(this.$coderDocsResponsiveSidebar) // CoderDocs: onresize
  },
  methods: {
    calcTagCloudCSS(tag) {
      const cssList = [
        'badge bg-primary text-decoration-none link-light',
        'badge bg-secondary text-decoration-none link-light',
        'badge bg-success text-decoration-none link-light',
        'badge bg-danger text-decoration-none link-light',
        'badge bg-warning text-dark text-decoration-none link-dark',
        'badge bg-info text-decoration-none link-light',
      ]
      return this.$pickupFromHash(tag, cssList)
    },
    coderDocsResponsiveSidebar() {
      const sidebar = this.$global.sidebar
      if (sidebar) {
        const w = window.innerWidth
        if (w >= 1024) {
          sidebar.classList.remove('sidebar-hidden')
          sidebar.classList.add('sidebar-visible')
        } else {
          sidebar.classList.remove('sidebar-visible')
          sidebar.classList.add('sidebar-hidden')
        }
      }
    },
    coderDocsSidebarToggler() {
      const sidebar = this.$global.sidebar
      if (sidebar) {
        if (sidebar.classList.contains('sidebar-visible')) {
          sidebar.classList.remove('sidebar-visible');
          sidebar.classList.add('sidebar-hidden');
        } else {
          sidebar.classList.remove('sidebar-hidden');
          sidebar.classList.add('sidebar-visible');
        }
      }
    }
  }
}
</script>
