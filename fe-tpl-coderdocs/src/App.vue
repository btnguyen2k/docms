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
  registerResize, unregisterResize,
  triggerPopstate, triggerResize,
} from '@/components/utils'

window.addEventListener('popstate', () => triggerPopstate())
window.addEventListener('resize', () => triggerResize())

export default {
  name: 'App',
  unmounted() {
    unregisterResize(this.$coderDocsResponsiveSidebar) // CoderDocs
  },
  mounted() {
    registerResize(this.$coderDocsResponsiveSidebar) // CoderDocs: onresize
  },
}
</script>
