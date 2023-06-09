The site metadata file resides at the `root directory` and contains the following fields:

## name

`string`, website's short name, example:
```yaml
name: DO CMS
```

## icon

`string`, website's icon, example:
```yaml
icon: fas fa-globe
```

```bs-alert info flex
<i class="fas fa-circle-info fa-2xl me-3"></i>
`DO CMS runtime` supports [FontAwesome icons](https://fontawesome.com/search?m=free) and [Bootstrap icons](https://icons.getbootstrap.com/).
```

```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-3"></i>
FontAwesome and Bootstrap icons are supported by the frontend templates included in the offical `DO CMS runtime`. If you use custom frontend templates from 3rd parties, please check their documentations.
```

## languages

`map[language-code:display-label]`, website's support languages, example:
```yaml
languages:
  en: English
  vi: Tiếng Việt
  default: en # default language
```

> Language codes should follow [ISO 639-1 codes](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
>
> The special entry `default: en` sets `en` as the default language.

## description

`string` or `map[language-code:text]`, website's short description, example:
```yaml
description: "Content Management System where its content is built through CI/CD pipeline"
```

or
```yaml
description:
  en: "Content Management System where its content is built through CI/CD pipeline"
  vi: "Hệ thống Quản trị nội dung với dữ liệu được xây dựng thông qua qui trình CI/CD"
```

When the `description` field is a string, it is equivalent to `map[default-language-code:text]`. That means, the following setting
```yaml
description: "Content Management System where its content is built through CI/CD pipeline"
```
is equivalent to:
```yaml
description:
  en: "Content Management System where its content is built through CI/CD pipeline"
```
as `en` is the default language specified in the `languages` field above.

## contacts

`map[string:string]`, website's contact information, example:
```yaml
contacts:
  website: "https://github.com/btnguyen2k/docms"
  email: "btnguyen2k (at) gmail (dot) com"
  github: "https://github.com/btnguyen2k/"
  facebook: ""
  linkedin: "https://www.linkedin.com/in/btnguyen2k/"
  slack: ""
  twitter: ""
  discord: ""
```

## tagalias

`map[language-code:map[string:Array(string)]]`, map similar tags to one, example:
```yaml
tagalias:
  en:
    cms: [content management, content management system, docms, do cms]
    ci/cd: [cicd, ci cd, ci-cd]
    localhost: [local]
    docli: [cli]
  vi:
    cms: [quản trị nội dung, hệ thống quản trị nội dung, quản lý nội dung, hệ thống quản lý nội dung, docms, do cms]
    ci/cd: [cicd, ci cd, ci-cd]
    localhost: [local]
    docli: [cli]
```

## tags

`map[string:object]`, extra key-object data attached to site metadata, example:
```yaml
tags:
  build: ${build_datetime}
  tags1:
    key1: value 1
    key2: value 2
  tags2: [1, 2, 3]
```

```bs-alert primary

Special placeholders will be replaced with value by `DO CLI` when pre-processing website content:
- `${build_datetime}` - pre-processing timestamp, format `YYYYMMDDHHmmss`
- `${build_date}` - pre-processing timestamp, format `YYYYMMDD`
- `${build_time}` - pre-processing timestamp, format `HHmmss`
```

```bs-alert info flex
<i class="fas fa-circle-info me-3 fa-xl"></i>
Tags can be embedded to the document using syntax `[[do-tag`. Refer to [Supported Markdown syntax](../markdown/) for details.
```

## mode

`string`, website's mode, current allowed values are `document` and `blog`, example:
```yaml
mode: document
```

> The default mode is `document`.

When `blog` mode is on, the following features are available:
- API `/api/documents` supports 2 more request types:
  - `?p=latest`: return latest N blog posts.
  - `?p=special`: return list of "special"-purpose documents (non blog posts).
- Documents (blog posts) are sorted by creation time, newest first.
- `/feeds` returns latest blog posts as RSS newsfeed.

## author

`map[string:string]`, default author's info for blog posts, used in `blog` mode, example:
```yaml
author:
  name: Thanh Nguyen
  avatar: //www.gravatar.com/avatar/9d6ee977a36db465f103ea5c0e4b859c
```

## Example

Example of a full site metadata file:
```yaml
name: DO CMS
description:
  en: Content Management System where its content is built through CI/CD pipeline
  vi: Hệ thống Quản trị nội dung với dữ liệu được xây dựng thông qua qui trình CI/CD
icon: fas fa-code
languages:
  en: English
  vi: Tiếng Việt
  default: en
contacts:
  website: "https://github.com/btnguyen2k/docms"
  email: "btnguyen2k (at) gmail (dot) com"
  github: "https://github.com/btnguyen2k/"
  facebook: ""
  linkedin: "https://www.linkedin.com/in/btnguyen2k/"
  slack: ""
  twitter: ""
  discord: ""
tags:
  build: ${build_datetime}
tagalias:
  en:
    cms: [content management, content management system, docms, do cms]
    ci/cd: [cicd, ci cd, ci-cd]
    localhost: [local]
    docli: [cli]
  vi:
    cms: [quản trị nội dung, hệ thống quản trị nội dung, quản lý nội dung, hệ thống quản lý nội dung, docms, do cms]
    ci/cd: [cicd, ci cd, ci-cd]
    localhost: [local]
    docli: [cli]
mode: document
```

See also:
- [Topic metadata file](../topicmetadata/)
- [Docment metadata file](../documentmetadata/)
