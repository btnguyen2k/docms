Each `document directory` has its own metadata file. This document metadata file contains the following fields:

**icon** - `string`, document's icon, supporting [FontAwesome icons](https://fontawesome.com/search?m=free), example:
```yaml
icon: fas fa-file
```

```bs-alert info flex
<i class="fas fa-circle-info fa-2xl me-2"></i>
`DO CMS runtime` hỗ trợ bộ biểu tượng [FontAwesome](https://fontawesome.com/search?m=free) và [Bootstrap](https://icons.getbootstrap.com/).
```

```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-2"></i>
Bộ biểu tượng FontAwesome và Bootstrap hỗ trợ bởi các giao diện đồ hoạ đi kèm với phiên bản bản `DO CMS runtime` gốc. Nếu bạn sử dụng một bộ giao diện đồ hoạ của bên thứ 3, vui lòng kiểm tra tài liệu đi kèm.
```

**title** - `string` or `map[language-code:text]`, document's title, example:
```yaml
title: "What it is"
```

or
```yaml
title:
  en: What it is
  vi: DO CMS là gì
```

When the `title` field is a string, it is equivalent to `map[default-language-code:text]`. That means, the following setting
```yaml
title: "What it is"
```
is equivalent to:
```yaml
title:
  en: What it is
```
as `en` is the default language specified in the [site metadata file](../sitemetadata/).

**summary** - `string` or `map[language-code:text]`, document's summary, example:
```yaml
summary: "DO CMS is a Content Management System that helps authors publish website content through a CI/CD flow. Unlike other CMS, there is no UI to create, update and publish content in DO CMS. Instead, website content is built and published via CI/CD pipelines."
```

or
```yaml
summary:
  en: "DO CMS is a Content Management System that helps authors publish website content through a CI/CD flow. Unlike other CMS, there is no UI to create, update and publish content in DO CMS. Instead, website content is built and published via CI/CD pipelines."
  vi: "DO CMS là Hệ thống quản lý nội dung giúp tác giả xuất bản nội dung trang web thông qua luồng CI/CD. Sẽ không có giao diện để người dùng tạo, cập nhật và xuất bản nội dung lên trang web. Thay vào đó, nội dung của trang web sẽ được xây dựng và xuất bản thông qua qui trình CI/CD."
```

When the `summary` field is a string, it is equivalent to `map[default-language-code:text]`. That means, the following setting
```yaml
summary: "DO CMS is a Content Management System that helps authors publish website content through a CI/CD flow. Unlike other CMS, there is no UI to create, update and publish content in DO CMS. Instead, website content is built and published via CI/CD pipelines."
```
is equivalent to:
```yaml
summary:
  en: "DO CMS is a Content Management System that helps authors publish website content through a CI/CD flow. Unlike other CMS, there is no UI to create, update and publish content in DO CMS. Instead, website content is built and published via CI/CD pipelines."
```
as `en` is the default language specified in the [site metadata file](../sitemetadata/).

**file** - `string` or `map[language-code:file-name]`, document's content file, example:
```yaml
file: index.md
```

or
```yaml
file:
  en: index-en.md
  vi: index-vi.md
```

When the `file` field is a string, it is equivalent to `map[default-language-code:file-name]`. That means, the following setting
```yaml
file: index.md
```
is equivalent to:
```yaml
file:
  en: index.md
```
as `en` is the default language specified in the [site metadata file](../sitemetadata/).

Example of a full document metadata file:
```yaml
title:
  en: What it is
  vi: DO CMS là gì
summary:
  en: "DO CMS is a Content Management System that helps authors publish website content through a CI/CD flow. Unlike other CMS, there is no UI to create, update and publish content in DO CMS. Instead, website content is built and published via CI/CD pipelines."
  vi: "DO CMS là Hệ thống quản lý nội dung giúp tác giả xuất bản nội dung trang web thông qua luồng CI/CD. Sẽ không có giao diện để người dùng tạo, cập nhật và xuất bản nội dung lên trang web. Thay vào đó, nội dung của trang web sẽ được xây dựng và xuất bản thông qua qui trình CI/CD."
icon:
file:
  en: index-en.md
  vi: index-vi.md
```

See also:
- [Site metadata file](../sitemetadata/)
- [Topic metadata file](../topicmetadata/)
