<template>
  <ul :class="$props.class">
    <li class="list-inline-item" v-if="$siteMeta.contacts.website">
      <a :title="$siteMeta.contacts.website" :href="$siteMeta.contacts.website" target="_blank"><i class="bi bi-globe" /></a>
    </li>
    <li class="list-inline-item" v-if="$siteMeta.contacts.email">
      <a :title="$siteMeta.contacts.email" :href="'mailto:'+$siteMeta.contacts.email" target="_blank"><i class="bi bi-envelope-fill" /></a>
    </li>
    <li class="list-inline-item" v-if="$siteMeta.contacts.github">
      <a :title="$siteMeta.contacts.github" :href="$siteMeta.contacts.github" target="_blank"><i class="bi bi-github" /></a>
    </li>
    <li class="list-inline-item" v-if="$siteMeta.contacts.facebook">
      <a :title="$siteMeta.contacts.facebook" :href="$siteMeta.contacts.facebook" target="_blank"><i class="bi bi-facebook" /></a>
    </li>
    <li class="list-inline-item" v-if="$siteMeta.contacts.linkedin">
      <a :title="$siteMeta.contacts.linkedin" :href="$siteMeta.contacts.linkedin" target="_blank"><i class="bi bi-linkedin" /></a>
    </li>
    <li class="list-inline-item" v-if="$siteMeta.contacts.slack">
      <a :title="$siteMeta.contacts.slack" :href="$siteMeta.contacts.slack" target="_blank"><i class="bi bi-slack" /></a>
    </li>
    <li class="list-inline-item" v-if="$siteMeta.contacts.twitter">
      <a :title="$siteMeta.contacts.twitter" :href="$siteMeta.contacts.twitter" target="_blank"><i class="bi bi-twitter" /></a>
    </li>
    <li class="list-inline-item" v-if="$siteMeta.contacts.discord">
      <a :title="$siteMeta.contacts.discord" :href="$siteMeta.contacts.discord" target="_blank"><i class="bi bi-discord" /></a>
    </li>
  </ul>
</template>

<script>
export default {
  name: 'lego-social-list',
  props: ['class'],
  inject: ['$siteMeta'],
}
</script>
