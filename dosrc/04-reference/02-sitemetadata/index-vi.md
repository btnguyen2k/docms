Tập tin metadata trang web nằm ở `thư mục gốc` và chứa các trường thông tin sau:

**name** - `string`, tên (ngắn gọn) của trang web, ví dụ:
```yaml
name: DO CMS
```

**icon** - `string`, biểu tượng của trang web, hỗ trợ [icon FontAwesome](https://fontawesome.com/search?m=free), ví dụ:
```yaml
icon: fas fa-globe
```

```bs-alert info flex
<i class="fas fa-circle-info fa-2xl me-2"></i>
`DO CMS runtime` hỗ trợ bộ biểu tượng [FontAwesome](https://fontawesome.com/search?m=free) và [Bootstrap](https://icons.getbootstrap.com/).
```

```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-2"></i>
Bộ biểu tượng FontAwesome và Bootstrap hỗ trợ bởi các giao diện đồ hoạ đi kèm với phiên bản bản `DO CMS runtime` gốc. Nếu bạn sử dụng một bộ giao diện đồ hoạ của bên thứ 3, vui lòng kiểm tra tài liệu đi kèm.
```

**languages** - `map[language-code:display-label]`, các ngôn ngữ được hỗ trợ trên trang web, ví dụ:
```yaml
languages:
  en: English
  vi: Tiếng Việt
  default: vi # ngôn ngữ mặc định của trang web
```

> Mã ngôn ngữ nên tuân theo chuẩn [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
>
> Dòng `default: vi` chỉ thị rằng `vi` là ngôn ngữ mặc định của trang web.

**description** - `string` or `map[language-code:text]`, mô tả ngắn dọn về trang web, ví dụ:
```yaml
description: "Hệ thống Quản trị nội dung với dữ liệu được xây dựng thông qua qui trình CI/CD"
```

hoặc
```yaml
description:
  en: "Content Management System where its content is built through CI/CD pipeline"
  vi: "Hệ thống Quản trị nội dung với dữ liệu được xây dựng thông qua qui trình CI/CD"
```

Khi `description` là một string, nó sẽ tương đương với `map[ngôn-ngữ-mặc-định:text]`. Điều đó có nghĩa là dòng chỉ thị như sau
```yaml
description: "Hệ thống Quản trị nội dung với dữ liệu được xây dựng thông qua qui trình CI/CD"
```
sẽ tương đương với:
```yaml
description:
  vi: "Hệ thống Quản trị nội dung với dữ liệu được xây dựng thông qua qui trình CI/CD"
```
bởi vì `vi` đang được chỉ định là ngôn ngữ mặc định của trang web trong trường `languages` ở trên.

**contacts** - `map[string:string]`, thông tin liên lạc, ví dụ:
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

**tags** - `map[string:string]`, các thông tin khác ở dạng key-value, ví dụ:
```yaml
tags:
  build: ${build_datetime}
  mytag: My value
```

> Các chuỗi đặc biệt sẽ được thay bằng giá trị thực trong quá trình tiền xử lý với công cụ `DO CLI`:
> - `${build_datetime}` - thay thế bằng thời gian chạy tiền xửu lý ở định dạng`YYYYMMDDHHmmss`
> - `${build_date}` - pre-processing timestamp, format `YYYYMMDD`
> - `${build_time}` - pre-processing timestamp, format `HHmmss`

Một ví dụ đầy đủ của tập tin metadata trang web:
```yaml
name: DO CMS
description:
  en: Content Management System where its content is built through CI/CD pipeline
  vi: Hệ thống Quản trị nội dung với dữ liệu được xây dựng thông qua qui trình CI/CD
icon: fas fa-code
languages:
  en: English
  vi: Tiếng Việt
  default: vi
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
```

Xem thêm:
- [Tập tin metadata chủ đề](../topicmetadata/)
- [Tập tin metadata bài viết](../documentmetadata/)
