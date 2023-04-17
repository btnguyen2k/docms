Mỗi một `thư mục chủ đề` có một tập tin metadata riêng. Tập tin này chứa các trường thông tin sau:

**icon** - `string`, icon của chủ đề, hỗ trợ [icon FontAwesome](https://fontawesome.com/search?m=free), ví dụ:
```yaml
icon: fas fa-book
```

```bs-alert info flex
<i class="fas fa-circle-info fa-2xl me-2"></i>
`DO CMS runtime` hỗ trợ bộ biểu tượng [FontAwesome](https://fontawesome.com/search?m=free) và [Bootstrap](https://icons.getbootstrap.com/).
```

```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-2"></i>
Bộ biểu tượng FontAwesome và Bootstrap hỗ trợ bởi các giao diện đồ hoạ đi kèm với phiên bản bản `DO CMS runtime` gốc. Nếu bạn sử dụng một bộ giao diện đồ hoạ của bên thứ 3, vui lòng kiểm tra tài liệu đi kèm.
```

**title** - `string` or `map[language-code:text]`, tiêu đề ngắn gọn, ví dụ:
```yaml
title: "Giới thiệu"
```

hoặc
```yaml
title:
  en: Introduction
  vi: Giới thiệu
```

Khi `title` là một string, nó sẽ tương đương với `map[ngôn-ngữ-mặc-định:text]`. Điều đó có nghĩa là dòng chỉ thị như sau
```yaml
title: "Giới thiệu"
```
sẽ tương đương với:
```yaml
title:
  vi: Giới thiệu
```
bởi vì `vi` đang được chỉ định là ngôn ngữ mặc định của trang web trong [tập tin metadata trang web](../sitemetadata/).

**description** - `string` or `map[language-code:text]`, mô tả ngắn dọn về chủ đề, ví dụ:
```yaml
description: "Giới thiệu về DO CMS: tổng quan và cách thức hoạt động."
```

hoặc
```yaml
description:
  en: "An introduction of DO CMS: what it is and how it work."
  vi: "Giới thiệu về DO CMS: tổng quan và cách thức hoạt động."
```

Khi `description` là một string, nó sẽ tương đương với `map[ngôn-ngữ-mặc-định:text]`. Điều đó có nghĩa là dòng chỉ thị như sau
```yaml
description: "Giới thiệu về DO CMS: tổng quan và cách thức hoạt động."
```
sẽ tương đương với:
```yaml
description:
  vi: "Giới thiệu về DO CMS: tổng quan và cách thức hoạt động."
```
bởi vì `vi` đang được chỉ định là ngôn ngữ mặc định của trang web trong [tập tin metadata trang web](../sitemetadata/).

Một ví dụ đầy đủ của tập tin metadata chủ đề:
```yaml
title:
  en: Introduction
  vi: Giới thiệu
description:
  en: "An introduction of DO CMS: what it is and how it work."
  vi: "Giới thiệu về DO CMS: tổng quan và cách thức hoạt động."
icon: fas fa-lightbulb
```

Xem thêm:
- [Tập tin metadata trang web](../sitemetadata/)
- [Tập tin metadata bài viết](../documentmetadata/)
