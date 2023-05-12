<template>
  <footer class="site-footer">
    <div class="container">
      <div class="row">
        <div class="col-lg-4">
          <div class="widget">
            <!--<h3 class="mb-4">{{ $siteMeta.name }}</h3>-->
            <p>{{ $localedText($siteMeta.description) }}</p>
          </div>
          <div class="widget">
            <h3>{{ $t('contact_info') }}</h3>
            <ul class="list-unstyled social">
              <li class="me-1" v-if="$siteMeta.contacts.website">
                <a :title="$siteMeta.contacts.website" :href="$siteMeta.contacts.website" target="_blank"><span class="icon-globe"></span></a>
              </li>
              <li class="me-1" v-if="$siteMeta.contacts.email">
                <a :title="$siteMeta.contacts.email" :href="'mailto:'+$siteMeta.contacts.email" target="_blank"><span class="icon-envelope"></span></a>
              </li>
              <li class="me-1" v-if="$siteMeta.contacts.github">
                <a :title="$siteMeta.contacts.github" :href="$siteMeta.contacts.github" target="_blank"><span class="icon-github"></span></a>
              </li>
              <li class="me-1" v-if="$siteMeta.contacts.facebook">
                <a :title="$siteMeta.contacts.facebook" :href="$siteMeta.contacts.facebook" target="_blank"><span class="icon-facebook"></span></a>
              </li>
              <li class="me-1" v-if="$siteMeta.contacts.linkedin">
                <a :title="$siteMeta.contacts.linkedin" :href="$siteMeta.contacts.linkedin" target="_blank"><span class="icon-linkedin"></span></a>
              </li>
              <li class="me-1" v-if="$siteMeta.contacts.twitter">
                <a :title="$siteMeta.contacts.twitter" :href="$siteMeta.contacts.twitter" target="_blank"><span class="icon-twitter"></span></a>
              </li>
              <li class="me-1" v-if="$siteMeta.contacts.slack">
                <a :title="$siteMeta.contacts.slack" :href="$siteMeta.contacts.slack" target="_blank"><span class="icon-slack"></span></a>
              </li>
              <li class="me-1" v-if="$siteMeta.contacts.discord">
                <a :title="$siteMeta.contacts.discord" :href="$siteMeta.contacts.discord" target="_blank"><span class="bi bi-discord"></span></a>
              </li>
            </ul>
          </div>
        </div>
        <div class="col-lg-4 ps-lg-5">
          <div class="widget">
            <h3 class="mb-4">{{ $siteMeta.name }}</h3>
            <ul class="list-unstyled float-start links">
              <li v-for="pageId in specialPages" v-bind:key="pageId">
                <router-link :to="{name: 'Document', params: {tid: $specialDocuments[pageId].topic.id, did: $specialDocuments[pageId].id}}">{{ $localedText($specialDocuments[pageId].title) }}</router-link>
              </li>
            </ul>
<!--            <ul class="list-unstyled float-start links">-->
<!--              <li><a href="#">Partners</a></li>-->
<!--              <li><a href="#">Business</a></li>-->
<!--              <li><a href="#">Careers</a></li>-->
<!--              <li><a href="#">Blog</a></li>-->
<!--              <li><a href="#">FAQ</a></li>-->
<!--              <li><a href="#">Creative</a></li>-->
<!--            </ul>-->
          </div>
        </div>
        <div class="col-lg-4">
          <div class="widget" v-if="$props['documentList']">
            <h3 class="mb-4">{{ $t('recent_posts') }}</h3>
            <div class="post-entry-footer">
              <ul>
                <li v-for="doc in $props['documentList'].slice(0,3)" v-bind:key="doc.id">
                  <router-link :to="{name: 'Document', params: {tid: doc.topic.id, did: doc.id}}">
                    <img :src="$calcDocumentEntryImgUrl(doc, doc.topic.id, '//placehold.co/700x700/?text='+doc.id, 's')" class="me-4 rounded">
                    <div class="text">
                      <h4>{{ $localedText(doc.title) }}</h4>
                      <div class="post-meta">
                        <span class="mr-2">{{ $unixTimestampToReadable(doc.tu) }}</span>
                      </div>
                    </div>
                  </router-link>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
      <div class="row mt-5">
        <div class="col-12 text-center">
          <!--
              **==========
              NOTE:
              Please don't remove this copyright link unless you buy the license here https://untree.co/license/
              **==========
            -->
          <!-- License information: https://untree.co/license/ -->
          <p>Copyright &copy; {{ $siteMeta.name}} - Template <code>Blogy</code>, designed with love by <a href="https://untree.co" target="_blank">Untree.co</a>
          </p>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'lego-page-footer',
  inject: ['$siteMeta', '$latestDocuments', '$specialDocuments'],
  props: ['document-list'],
  mounted() {
  },
  computed: {
    now() {
      return new Date()
    },
    specialPages() {
      return ['about', 'contact']
    },
  },
  methods: {
  },
  data() {
    return {
    }
  },
}
</script>
