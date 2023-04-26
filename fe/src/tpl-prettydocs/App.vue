<template>
  <router-view></router-view>
</template>

<style lang="css">
/* Google Fonts */
@import "//fonts.googleapis.com/css?family=Open+Sans:300italic,400italic,600italic,700italic,800italic,400,300,600,700,800";

/* FontAwesome */
@import "@/_shared/assets/fontawesome-free-6.4.0-web/css/all.min.css";

/* Bootstrap 5.x */
@import "bootstrap/dist/css/bootstrap.min.css";
@import "bootstrap";
@import "bootstrap-icons";

/* Plugin: elegant font */
/*@import "@/assets/plugins/elegant_font/css/style.css";*/

/* Application's main style */
@import "./assets/css/styles.css";
</style>

<script>
import {triggerPopstate, triggerResize} from '@/_shared/utils/docms_utils'

window.addEventListener('popstate', () => triggerPopstate())
window.addEventListener('resize', () => triggerResize())

export default {
  name: 'App',
  inject: ['$global'],
  provide() {
    return {
      $calcTagCloudCSS: this.calcTagCloudCSS,
    }
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
    }
  }
}
</script>
