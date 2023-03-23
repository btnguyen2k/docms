//#DO CMS frontend, template PrettyDocs
import { createI18n } from 'vue-i18n'
import utils from '@/utils/app_utils'

const messages = {
    en: {
        wait: 'Please wait...',
        search: 'Search',
        search_prompt: 'Enter search term',
        home: 'Home',
        topics: 'Topics',
        topic: 'Topic',
        documents: 'Documents',
        document: 'Document',
        contact_info: 'Contact us at',

        error: 'Error',
        error_topic_not_found: 'Topic "{topic}" not found',
        error_document_not_found: 'Document "{topic}/{document}" not found',

        empty_topic: 'This topic has no document page',
    },
    vi: {
        wait: 'Vui lòng giờ giây lát...',
        search: 'Tìm kiếm',
        search_prompt: 'Nhập câu truy vấn tìm kiếm',
        home: 'Trang nhà',
        topics: 'Chủ đề',
        topic: 'Chủ đề',
        documents: 'Bài viết',
        document: 'Bài viết',
        contact_info: 'Thông tin liên hệ',

        error: 'Lỗi',
        error_topic_not_found: 'Không tìm thấy chủ đề "{topic}"',
        error_document_not_found: 'Không tìm thấy bài viết "{topic}/{document}"',

        empty_topic: 'Chủ đề này hiện chưa có bài viết nào',
    }
}

let savedLocale = utils.localStorageGet('_l')
savedLocale = savedLocale ? (messages[savedLocale] ? savedLocale : 'en') : 'en'
const _i18n = createI18n({
    locale: savedLocale,
    fallbackLocale: 'en',
    messages: messages,
})

/* i18n.global is readonly, we need to clone a new instance and make it reactive */
const i18n = { ..._i18n }
import { reactive, watchEffect } from 'vue'
i18n.global = reactive(i18n.global)
let oldLocale = i18n.global.locale
watchEffect(() => {
    if (i18n.global.locale !== oldLocale) {
        utils.localStorageSet('_l', i18n.global.locale)
        oldLocale = i18n.global.locale
    }
})

export default i18n

export function swichLanguage(locale, refreshPage) {
    if (locale !== oldLocale) {
        i18n.global.locale = locale
        if (refreshPage) {
            window.location.reload()
        }
    }
}
