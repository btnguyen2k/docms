<template>
  <router-view></router-view>
</template>

<style lang="css">
/* FontAwesome */
@import "@/assets/fontawesome-free-6.4.0-web/css/all.min.css";

/* Bootstrap 5.x */
@import "bootstrap/dist/css/bootstrap.min.css";
@import "bootstrap";
@import "bootstrap-icons";

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
  name: 'App',
}
</script>
