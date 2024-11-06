Each `document directory` has its own metadata file. This document metadata file contains the following fields:

## icon

`string`, document's icon, supporting [FontAwesome icons](https://fontawesome.com/search?m=free), example:
```yaml
icon: fas fa-file
```

```bs-alert info flex
<i class="fas fa-circle-info fa-2xl me-3"></i>
`DO CMS runtime` supports [FontAwesome icons](https://fontawesome.com/search?m=free) and [Bootstrap icons](https://icons.getbootstrap.com/).
```

```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-3"></i>
FontAwesome and Bootstrap icons are supported by the frontend templates included in the offical `DO CMS runtime`. If you use custom frontend templates from 3rd parties, please check their documentations.
```

## title

`string` or `map[language-code:text]`, document's title, example:
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

## summary

`string` or `map[language-code:text]`, document's summary, example:
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

## file

`string` or `map[language-code:file-name]`, document's content file, example:
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

## tags

`Array("tag1","tag2",...)` or `map[language-code:Array("tag1","tag2",...)]`, arbitrary info attached to the document, example:
```yaml
tags: [cloud, virtual machine]
```

or
```yaml
tags:
  en: [cloud, virtual machine]
  vi: [đám mây, máy ảo]
```

When the `tags` field is an array, it is equivalent to `map[default-language-code:Array(...)]`. That means, the following setting
```yaml
tags: [cloud, virtual machine]
```
is equivalent to:
```yaml
tags:
  en: [cloud, virtual machine]
```
as `en` is the default language specified in the [site metadata file](../sitemetadata/).

## img

`string` - document's entry image, either a single url to the image, example:
```yaml
img: entry-img.jpg
```

or a `map[id:img-url]` specifying multiple images, encoded as a JSON string, example:
```yaml
img: '{"square":"entry-img-s.jpg","vertical":"entry-img-v.jpg","hortizontal":"entry-img-h.jpg"}
```

## tu and tc

`UNIX timestamp` - timestamp when the document was created/last updated, example:
```yaml
tc: 1683215751
tu: 1683215751
```

## page

`string` - indicate that this document is a special purpose page, used in `blog` [mode](../sitemetadata/#mode), example:
```yaml
page: about
```

## style

`string` - indicate that the GUI should present this document in a specific style, example:
```yaml
style: fullscreen
```

## Example

Example of a full document metadata file:
```yaml
title:
  en: What it is
  vi: DO CMS là gì
summary:
  en: "DO CMS is a Content Management System that helps authors publish website content through a CI/CD flow. Unlike other CMS, there is no UI to create, update and publish content in DO CMS. Instead, website content is built and published via CI/CD pipelines."
  vi: "DO CMS là Hệ thống quản lý nội dung giúp tác giả xuất bản nội dung trang web thông qua luồng CI/CD. Sẽ không có giao diện để người dùng tạo, cập nhật và xuất bản nội dung lên trang web. Thay vào đó, nội dung của trang web sẽ được xây dựng và xuất bản thông qua qui trình CI/CD."
icon:
img: '{"s":"entry-img-s.jpg","v":"entry-img-v.jpg","h":"entry-img-h.jpg"}'
file:
  en: index-en.md
  vi: index-vi.md
tags:
  en: [cms, content, devops]
  vi: [cms, nội dung, devops]
tc: 1683215751
tu: 1683215751
```

See also:
- [Site metadata file](../sitemetadata/)
- [Topic metadata file](../topicmetadata/)
