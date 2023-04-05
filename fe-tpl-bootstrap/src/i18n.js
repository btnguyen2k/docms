//#DO CMS frontend, template PrettyDocs
import { createI18n } from 'vue-i18n'
import { localStorageGet, localStorageSet } from '@/utils/app_utils'

const messages = {
    en: {
        _name: 'English',
        _flag: 'cif-gb',

        language: 'Language',
        wait: 'Please wait...',
        search: 'Search',
        search_result: 'Search result',
        search_no_result: 'Search returns no result, please try another query',
        search_prompt: 'Enter search query',
        home: 'Home',
        topics: 'Topics',
        topic: 'Topic',
        documents: 'Documents',
        document: 'Document',
        read: 'Read',
        explore: 'Explore',
        contact_info: 'Contact us at',

        error: 'Error',
        error_topic_not_found: 'Topic "{topic}" not found',
        error_document_not_found: 'Document "{topic}/{document}" not found',

        empty_topic: 'This topic has no document page',
    },
    vi: {
        _name: 'Tiếng Việt',
        _flag: 'cif-vn',

        language: 'Ngôn ngữ',
        wait: 'Vui lòng giờ giây lát...',
        search: 'Tìm kiếm',
        search_result: 'Kết quả tìm kiếm',
        search_no_result: 'Tìm kiếm không trả về kết quả, vui lòng thử với câu truy vấn khác',
        search_prompt: 'Nhập câu truy vấn tìm kiếm',
        home: 'Trang nhà',
        topics: 'Chủ đề',
        topic: 'Chủ đề',
        documents: 'Bài viết',
        document: 'Bài viết',
        read: 'Xem',
        explore: 'Khám phá',
        contact_info: 'Thông tin liên hệ',

        error: 'Lỗi',
        error_topic_not_found: 'Không tìm thấy chủ đề "{topic}"',
        error_document_not_found: 'Không tìm thấy bài viết "{topic}/{document}"',

        empty_topic: 'Chủ đề này hiện chưa có bài viết nào',
    }
}

let savedLocale = localStorageGet('_l')
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
        localStorageSet('_l', i18n.global.locale)
        oldLocale = i18n.global.locale
    }
})

export default i18n

export function swichLanguage(locale, refreshPage) {
    if (locale=='') {
        locale = localStorageGet('_l')
    }
    if (locale !== oldLocale) {
        i18n.global.locale = messages[locale] ? locale : 'en'
        if (refreshPage) {
            window.location.reload()
        }
    }
}
