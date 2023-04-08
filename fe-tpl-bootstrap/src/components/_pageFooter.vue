<template>
  <footer class="py-5 bg-dark">
    <div class="container px-5">
      <p class="m-0 text-center text-white">
        Powered by <a href="https://github.com/btnguyen2k/docms" target="_blank">DO CMS</a>.
        Copyright &copy; {{ $siteMeta.name }}.
        <small v-if="$siteMeta.tags.build">Build {{ $siteMeta.tags.build }}</small>
      </p>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'lego-page-footer',
  inject: ['$siteMeta'],
}
</script>
