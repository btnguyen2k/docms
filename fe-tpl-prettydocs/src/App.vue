<template>
  <router-view></router-view>
</template>

<style lang="css">
/* Google Fonts */
@import "//fonts.googleapis.com/css?family=Open+Sans:300italic,400italic,600italic,700italic,800italic,400,300,600,700,800";

/* FontAwesome */
@import "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.3.0/css/all.min.css";

/* Bootstrap 5 */
@import "bootstrap/dist/css/bootstrap.min.css";
@import "bootstrap";

/* Plugin: elegant font */
/*@import "@/assets/plugins/elegant_font/css/style.css";*/

/*@import "@coreui/coreui/dist/css/coreui.min.css";*/
/*@import "@coreui/icons/css/all.min.css";*/

/* Application's main style */
@import "@/assets/css/styles.css";

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
import {triggerPopstate, triggerResize} from '@/components/utils'
window.addEventListener('popstate', () => triggerPopstate())
window.addEventListener('resize', () => triggerResize())

export default {
  name: 'App'
}
</script>