<template>
  <div class="site-cover site-cover-sm same-height overlay single-page"
       :style="'background-image: url('+($calcDocumentEntryImgUrl(document, topic.id, '//placehold.co/700x440/6c757d/6c757d', 'h'))">
    <div class="container">
      <div class="row same-height justify-content-center">
        <div class="col-md-6">
          <div class="post-entry text-center" data-aos="zoom-in">
            <h1 class="mb-4">{{ $localedText(document.title) }}</h1>
            <div class="post-meta align-items-center text-center">
              <figure class="author-figure mb-0 me-3 d-inline-block">
                <img :src="$calcAuthorAvatarUrl(document.author)" class="img-fluid">
              </figure>
              <span class="d-inline-block mt-1">{{ document.author.name }}</span>
              <span>&nbsp;-&nbsp;{{ $unixTimestampToReadable(document.tu) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <section class="section">
    <div class="container">
      <div class="row blog-entries element-animate">
        <div class="col-md-12 col-lg-8 main-content">
          <div class="d-inline-block mb-2" v-if="document.tags && $localedText(document.tags).length>0">
            <ul class="tags" style="font-size: small">
              <li v-for="tag in $localedText(document.tags)" v-bind:key="tag"><router-link :to="{name: 'TagSearch', query:{q: tag, l: $i18n.locale}}">{{ tag }}</router-link></li>
            </ul>
          </div>

          <div class="post-content-body">
            <p class="text-end" style="font-size: small">{{ $t('document_length', {words: documentContentNumWords, time: documentContentTimeRead}) }}</p>
            <div class="docms-content img-fit img-center" v-html="documentContentRendered"></div>
          </div>

          <!--            <div class="pt-5" v-if="document.tags && $localedText(document.tags).length>0">-->
          <!--              <ul class="tags" style="font-size: small">-->
          <!--                <li v-for="tag in $localedText(document.tags)" v-bind:key="tag"><router-link :to="{name: 'TagSearch', query:{q: tag, l: $i18n.locale}}">{{ tag }}</router-link></li>-->
          <!--              </ul>-->
          <!--            </div>-->

          <!--            <div class="pt-5 comment-wrap">-->
          <!--              <h3 class="mb-5 heading">6 Comments</h3>-->
          <!--              <ul class="comment-list">-->
          <!--                <li class="comment">-->
          <!--                  <div class="vcard">-->
          <!--                    <img src="images/person_1.jpg" alt="Image placeholder">-->
          <!--                  </div>-->
          <!--                  <div class="comment-body">-->
          <!--                    <h3>Jean Doe</h3>-->
          <!--                    <div class="meta">January 9, 2018 at 2:21pm</div>-->
          <!--                    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Pariatur quidem laborum necessitatibus, ipsam impedit vitae autem, eum officia, fugiat saepe enim sapiente iste iure! Quam voluptas earum impedit necessitatibus, nihil?</p>-->
          <!--                    <p><a href="#" class="reply rounded">Reply</a></p>-->
          <!--                  </div>-->
          <!--                </li>-->

          <!--                <li class="comment">-->
          <!--                  <div class="vcard">-->
          <!--                    <img src="images/person_2.jpg" alt="Image placeholder">-->
          <!--                  </div>-->
          <!--                  <div class="comment-body">-->
          <!--                    <h3>Jean Doe</h3>-->
          <!--                    <div class="meta">January 9, 2018 at 2:21pm</div>-->
          <!--                    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Pariatur quidem laborum necessitatibus, ipsam impedit vitae autem, eum officia, fugiat saepe enim sapiente iste iure! Quam voluptas earum impedit necessitatibus, nihil?</p>-->
          <!--                    <p><a href="#" class="reply rounded">Reply</a></p>-->
          <!--                  </div>-->

          <!--                  <ul class="children">-->
          <!--                    <li class="comment">-->
          <!--                      <div class="vcard">-->
          <!--                        <img src="images/person_3.jpg" alt="Image placeholder">-->
          <!--                      </div>-->
          <!--                      <div class="comment-body">-->
          <!--                        <h3>Jean Doe</h3>-->
          <!--                        <div class="meta">January 9, 2018 at 2:21pm</div>-->
          <!--                        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Pariatur quidem laborum necessitatibus, ipsam impedit vitae autem, eum officia, fugiat saepe enim sapiente iste iure! Quam voluptas earum impedit necessitatibus, nihil?</p>-->
          <!--                        <p><a href="#" class="reply rounded">Reply</a></p>-->
          <!--                      </div>-->


          <!--                      <ul class="children">-->
          <!--                        <li class="comment">-->
          <!--                          <div class="vcard">-->
          <!--                            <img src="images/person_4.jpg" alt="Image placeholder">-->
          <!--                          </div>-->
          <!--                          <div class="comment-body">-->
          <!--                            <h3>Jean Doe</h3>-->
          <!--                            <div class="meta">January 9, 2018 at 2:21pm</div>-->
          <!--                            <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Pariatur quidem laborum necessitatibus, ipsam impedit vitae autem, eum officia, fugiat saepe enim sapiente iste iure! Quam voluptas earum impedit necessitatibus, nihil?</p>-->
          <!--                            <p><a href="#" class="reply rounded">Reply</a></p>-->
          <!--                          </div>-->

          <!--                          <ul class="children">-->
          <!--                            <li class="comment">-->
          <!--                              <div class="vcard">-->
          <!--                                <img src="images/person_5.jpg" alt="Image placeholder">-->
          <!--                              </div>-->
          <!--                              <div class="comment-body">-->
          <!--                                <h3>Jean Doe</h3>-->
          <!--                                <div class="meta">January 9, 2018 at 2:21pm</div>-->
          <!--                                <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Pariatur quidem laborum necessitatibus, ipsam impedit vitae autem, eum officia, fugiat saepe enim sapiente iste iure! Quam voluptas earum impedit necessitatibus, nihil?</p>-->
          <!--                                <p><a href="#" class="reply rounded">Reply</a></p>-->
          <!--                              </div>-->
          <!--                            </li>-->
          <!--                          </ul>-->
          <!--                        </li>-->
          <!--                      </ul>-->
          <!--                    </li>-->
          <!--                  </ul>-->
          <!--                </li>-->

          <!--                <li class="comment">-->
          <!--                  <div class="vcard">-->
          <!--                    <img src="images/person_1.jpg" alt="Image placeholder">-->
          <!--                  </div>-->
          <!--                  <div class="comment-body">-->
          <!--                    <h3>Jean Doe</h3>-->
          <!--                    <div class="meta">January 9, 2018 at 2:21pm</div>-->
          <!--                    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Pariatur quidem laborum necessitatibus, ipsam impedit vitae autem, eum officia, fugiat saepe enim sapiente iste iure! Quam voluptas earum impedit necessitatibus, nihil?</p>-->
          <!--                    <p><a href="#" class="reply rounded">Reply</a></p>-->
          <!--                  </div>-->
          <!--                </li>-->
          <!--              </ul>-->
          <!--              &lt;!&ndash; END comment-list &ndash;&gt;-->

          <!--              <div class="comment-form-wrap pt-5">-->
          <!--                <h3 class="mb-5">Leave a comment</h3>-->
          <!--                <form action="#" class="p-5 bg-light">-->
          <!--                  <div class="form-group">-->
          <!--                    <label for="name">Name *</label>-->
          <!--                    <input type="text" class="form-control" id="name">-->
          <!--                  </div>-->
          <!--                  <div class="form-group">-->
          <!--                    <label for="email">Email *</label>-->
          <!--                    <input type="email" class="form-control" id="email">-->
          <!--                  </div>-->
          <!--                  <div class="form-group">-->
          <!--                    <label for="website">Website</label>-->
          <!--                    <input type="url" class="form-control" id="website">-->
          <!--                  </div>-->

          <!--                  <div class="form-group">-->
          <!--                    <label for="message">Message</label>-->
          <!--                    <textarea name="" id="message" cols="30" rows="10" class="form-control"></textarea>-->
          <!--                  </div>-->
          <!--                  <div class="form-group">-->
          <!--                    <input type="submit" value="Post Comment" class="btn btn-primary">-->
          <!--                  </div>-->

          <!--                </form>-->
          <!--              </div>-->
          <!--            </div>-->
        </div>

        <sidebar :topic="topic" :document-list="sameTopicDocuments" :document-list-title="$t('same_topic')" />
      </div>
    </div>
  </section>
</template>

<script>
/* Lightbox for Bootstrap 5 */
import Lightbox from 'bs5-lightbox'

import {markdownRender} from '@/_shared/utils/docms_utils'
import sidebar from './_sidebar.vue'

export default {
  name: 'DocumentStyleNormal',
  inject: ['$global', '$siteMeta', '$siteTopics', '$latestDocuments'],
  props: ['topic', 'document', 'documentList'],
  components: {sidebar},
  computed: {
    documentContentNumWords() {
      return this.$localedText(this.document.content).split(/[^\p{L}\d]+/u).length
    },
    documentContentTimeRead() {
      const numWords = this.documentContentNumWords
      const numMins = Math.round(numWords/130)
      return numMins
    },
    documentContentRendered() {
      this._updateLightbox()
      this._scrollToHash(document.location.hash)
      return markdownRender(this.$localedText(this.document.content), {
        sanitize: true,
        tags: this.$siteMeta.tags,
      })
    },
    sameTopicDocuments() {
      const result = []
      for (let i = 0; i < this.documentList.length; i++) {
        if (this.document.id != this.documentList[i].id) {
          result.push(this.documentList[i])
        }
      }
      return result
    },
  },
  methods: {
    _scrollToHash(hash) {
      if (hash) {
        this.$nextTick(() => {
          const el = document.getElementById(hash.slice(1))
          if (el) {
            el.scrollIntoView({block: 'center', behavior: 'smooth'})
          }
        })
      }
    },
    _updateLightbox() {
      this.$nextTick(() => {
        document.querySelectorAll('[data-toggle="lightbox"]').forEach(el => el.addEventListener('click', Lightbox.initialize));
      })
    },
  },
}
</script>
