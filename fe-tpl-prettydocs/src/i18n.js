//#DO CMS frontend, template PrettyDocs
import { createI18n } from 'vue-i18n'
import utils from '@/utils/app_utils'

const messages = {
    en: {
        wait: 'Please wait...',
        error: 'Error',
        search: 'Search',
        search_prompt: 'Enter search term',
        home: 'Home',
        topics: 'Topics',
        topic: 'Topic',
        documents: 'Documents',
        document: 'Document',

        error_topic_not_found: 'Topic "{topic}" not found.',

        // click_to_home: 'Click here to navigate to the home page.',
        // click_to_topic: 'Click here to navigate to the topic index.',
        // contacts: 'Contacts',
        // pages: 'Pages',
        //
        // page_updated_timestamp: 'Last updated',
        // page_created_timestamp: 'Creation time',
        //
        // error_unknown: "Error occurred, try again after a moment! If the problem persists, please contact website administrator.",
        // error_product_not_found: 'No product mapped to domain "{domain}".',

        // error_page_not_found: 'Page "{page}" not found.',
        //
        // empty_topic: 'This topic has no document page.',
    },
    vi: {
        wait: 'Vui lòng giờ giây lát...',
        error: 'Lỗi',
        search: 'Tìm kiếm',
        search_prompt: 'Nhập câu truy vấn tìm kiếm',
        home: 'Trang nhà',
        topics: 'Chủ đề',
        topic: 'Chủ đề',
        documents: 'Bài viết',
        document: 'Bài viết',

        error_topic_not_found: 'Không tìm thấy chủ đề "{topic}".',

        // click_to_home: 'Nhấn vào đây để đi đến trang chủ.',
        // click_to_topic: 'Nhấn vào đây để đi đến trang nhà của chủ đề.',
        // contacts: 'Liên hệ',
        // pages: 'Trang tài liệu',
        //
        // page_updated_timestamp: 'Cập nhật lần cuối',
        // page_created_timestamp: 'Thời điểm tạo',
        //
        // error_unknown: "Có lỗi xảy ra, vui lòng thử lại sau. Nếu sự cố vẫn còn tiếp diễn, hãy liên lạc với người quản trị website.",
        // error_product_not_found: 'Không có sản phẩm nào tương ứng với tên miền "{domain}".',
        // error_page_not_found: 'Không tìm thấy trang tài liệu "{page}".',
        //
        // empty_topic: 'Chủ đề này hiện chưa có bài viết nào.',
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
