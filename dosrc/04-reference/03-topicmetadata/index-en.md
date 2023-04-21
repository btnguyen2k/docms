Each `topic directory` has its own metadata file. This topic metadata file contains the following fields:

**icon** - `string`, topic's icon, supporting [FontAwesome icons](https://fontawesome.com/search?m=free), example:
```yaml
icon: fas fa-book
```

```bs-alert info flex
<i class="fas fa-circle-info fa-2xl me-3"></i>
`DO CMS runtime` supports [FontAwesome icons](https://fontawesome.com/search?m=free) and [Bootstrap icons](https://icons.getbootstrap.com/).
```

```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-3"></i>
FontAwesome and Bootstrap icons are supported by the frontend templates included in the offical `DO CMS runtime`. If you use custom frontend templates from 3rd parties, please check their documentations.
```

**title** - `string` or `map[language-code:text]`, topic's title, example:
```yaml
title: "Introduction"
```

or
```yaml
title:
  en: Introduction
  vi: Giới thiệu
```

When the `title` field is a string, it is equivalent to `map[default-language-code:text]`. That means, the following setting
```yaml
title: "Introduction"
```
is equivalent to:
```yaml
title:
  en: "Introduction"
```
as `en` is the default language specified in the [site metadata file](../sitemetadata/).

**description** - `string` or `map[language-code:text]`, topic's short description, example:
```yaml
description: "An introduction of DO CMS: what it is and how it work."
```

or
```yaml
description:
  en: "An introduction of DO CMS: what it is and how it work."
  vi: "Giới thiệu về DO CMS: tổng quan và cách thức hoạt động."
```

When the `description` field is a string, it is equivalent to `map[default-language-code:text]`. That means, the following setting
```yaml
description: "An introduction of DO CMS: what it is and how it work."
```
is equivalent to:
```yaml
description:
  en: "An introduction of DO CMS: what it is and how it work."
```
as `en` is the default language specified in the [site metadata file](../sitemetadata/).

Example of a full topic metadata file:
```yaml
title:
  en: Introduction
  vi: Giới thiệu
description:
  en: "An introduction of DO CMS: what it is and how it work."
  vi: "Giới thiệu về DO CMS: tổng quan và cách thức hoạt động."
icon: fas fa-lightbulb
```

See also:
- [Site metadata file](../sitemetadata/)
- [Docment metadata file](../documentmetadata/)
