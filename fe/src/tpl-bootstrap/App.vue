<template>
  <router-view></router-view>
</template>

<style lang="css">
/* FontAwesome */
@import "@/_shared/assets/fontawesome-free-6.4.0-web/css/all.min.css";

/* Bootstrap 5.x */
@import "bootstrap/dist/css/bootstrap.min.css";
@import "bootstrap";
@import "bootstrap-icons";
</style>

<script>
import {triggerPopstate, triggerResize} from '@/_shared/utils/docms_utils'

window.addEventListener('popstate', () => triggerPopstate())
window.addEventListener('resize', () => triggerResize())

export default {
  name: 'App',
  inject: ['$global'],
}
</script>
