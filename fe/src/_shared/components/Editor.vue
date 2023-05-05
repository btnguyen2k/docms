<template>
  <div class="container">
      <nav class="pb-2 g-0">
        <div class="nav nav-tabs" id="nav-tab" role="tablist">
          <button class="nav-link active" id="nav-hstyle-tab" data-bs-toggle="tab" data-bs-target="#nav-hstyle"
                  type="button" role="tab" aria-controls="nav-hstyle" aria-selected="true"><strong>Horizontal
            Editor</strong></button>
          <button class="nav-link" id="nav-vstyle-tab" data-bs-toggle="tab" data-bs-target="#nav-vstyle" type="button"
                  role="tab" aria-controls="nav-vstyle" aria-selected="false"><strong>Vertical Editor</strong></button>
        </div>
      </nav>
      <div class="tab-content" id="nav-tabContent">
        <div class="tab-pane fade show active" id="nav-hstyle" role="tabpanel" aria-labelledby="nav-hstyle-tab">
          <div class="row">
            <div class="col-6">
              <form>
                <label for="hmarkdown" class="form-label"><strong>Markdown</strong> (<input type="checkbox"
                                                                                            v-model="textWrap"/> text
                  wrap)</label>
                <textarea id="hmarkdown" :class="'form-control '+(textWrap?'text-wrap':'text-nowrap')"
                          style="height: 75vh !important;" v-model="markdownContent"></textarea>
              </form>
            </div>
            <div class="col-6">
              <label for="hhtml" class="form-label"><strong>Preview</strong></label>
              <div id="hhtml" class="form-control docms-content img-fit img-center" style="height: 75vh !important; overflow: auto;"
                   v-html="markdownRendered"></div>
            </div>
          </div>
        </div>
        <div class="tab-pane fade" id="nav-vstyle" role="tabpanel" aria-labelledby="nav-vstyle-tab">
          <div class="row">
            <div class="col-12">
              <form>
                <label for="vmarkdown" class="form-label"><strong>Markdown</strong> (<input type="checkbox"
                                                                                            v-model="textWrap"/> text
                  wrap)</label>
                <textarea id="vmarkdown" :class="'form-control '+(textWrap?'text-wrap':'text-nowrap')"
                          style="height: 45vh !important;" v-model="markdownContent"></textarea>
              </form>
            </div>
            <div class="col-12">
              <label for="vhtml" class="form-label"><strong>Preview</strong></label>
              <div id="vhtml" class="form-control img-fit img-center" style="height: 45vh !important; overflow: auto;"
                   v-html="markdownRendered"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  <div id="mermaid-temp" class="d-none"></div>
</template>

<script>
/* Lightbox for Bootstrap 5 */
import Lightbox from 'bs5-lightbox'

import {localStorageGet, localStorageSet} from "@/_shared/utils/app_utils"
import {markdownRender} from "@/_shared/utils/docms_utils"
import "@/_shared/assets/markdown-gfm.css"

export default {
  name: 'Editor',
  mounted() {
    //ref: https://css-tricks.com/snippets/javascript/support-tabs-in-textareas/
    HTMLTextAreaElement.prototype.getCaretPosition = function () { //return the caret position of the textarea
      return this.selectionStart
    }
    HTMLTextAreaElement.prototype.setCaretPosition = function (position) { //change the caret position of the textarea
      this.selectionStart = position
      this.selectionEnd = position
      this.focus()
    }
    HTMLTextAreaElement.prototype.hasSelection = function () { //if the textarea has selection then return true
      return this.selectionStart != this.selectionEnd
    }
    HTMLTextAreaElement.prototype.getSelectedText = function () { //return the selection text
      return this.value.substring(this.selectionStart, this.selectionEnd)
    }
    HTMLTextAreaElement.prototype.setSelection = function (start, end) { //change the selection area of the textarea
      this.selectionStart = start
      this.selectionEnd = end
      this.focus()
    }
    const tabSize = "    "
    const fnOnKeyDown = event => { // support tab on textarea
      if (event.keyCode == 9) { // tab
        const el = event.srcElement
        const newCaretPosition = el.getCaretPosition() + tabSize.length
        el.value = el.value.substring(0, el.getCaretPosition()) + tabSize + el.value.substring(el.getCaretPosition(), el.value.length)
        el.setCaretPosition(newCaretPosition)
        return false
      }
      if (event.keyCode == 8) { // backspace
        const el = event.srcElement
        if (el.value.substring(el.getCaretPosition() - tabSize.length, el.getCaretPosition()) == tabSize) { //it's a tab space
          const newCaretPosition = el.getCaretPosition() - tabSize.length + 1
          el.value = el.value.substring(0, el.getCaretPosition() - tabSize.length + 1) + el.value.substring(el.getCaretPosition(), el.value.length)
          el.setCaretPosition(newCaretPosition)
        }
      }
      // if (event.keyCode == 37) { //left arrow
      //   let newCaretPosition;
      //   if (textarea.value.substring(textarea.getCaretPosition() - 4, textarea.getCaretPosition()) == "    ") { //it's a tab space
      //     newCaretPosition = textarea.getCaretPosition() - 3;
      //     textarea.setCaretPosition(newCaretPosition);
      //   }
      // }
      // if (event.keyCode == 39) { //right arrow
      //   let newCaretPosition;
      //   if (textarea.value.substring(textarea.getCaretPosition() + 4, textarea.getCaretPosition()) == "    ") { //it's a tab space
      //     newCaretPosition = textarea.getCaretPosition() + 3;
      //     textarea.setCaretPosition(newCaretPosition);
      //   }
      // }
    };
    document.getElementById('hmarkdown').onkeydown = fnOnKeyDown
    document.getElementById('vmarkdown').onkeydown = fnOnKeyDown

    const savedContent = localStorageGet("_editor")
    if (savedContent) {
      this.markdownContent = savedContent
    }
    this.stopTimer()
    this.startTimer()
  },
  unmounted() {
    this.stopTimer()
  },
  computed: {
    markdownRendered() {
      this._updateLightbox()
      return markdownRender(this.markdownContent, {sanitize: true, tags: {
          build: new Date(),
          demo: {
            tag1: {
              key1: 'value 1',
              key2: 'value 2'
            },
            tag2: [1, "2", true]
          }
        }
      })
    },
  },
  methods: {
    _updateLightbox() {
      this.$nextTick(() => {
        document.querySelectorAll('[data-toggle="lightbox"]').forEach(el => el.addEventListener('click', Lightbox.initialize));
      })
    },
    startTimer() {
      this.timerInterval = setInterval(() => (localStorageSet("_editor", this.markdownContent)), 5000);
    },
    stopTimer() {
      if (this.timerInterval) {
        clearInterval(this.timerInterval)
      }
    }
  },
  data() {
    return {
      markdownContent: '',
      timerInterval: null,
      textWrap: false,
    }
  },
}
</script>
