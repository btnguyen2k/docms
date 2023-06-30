Tập tin metadata trang web nằm ở `thư mục gốc` và chứa các trường thông tin sau:

## name

`string`, tên (ngắn gọn) của trang web, ví dụ:
```yaml
name: DO CMS
```

## icon

`string`, biểu tượng của trang web, hỗ trợ [icon FontAwesome](https://fontawesome.com/search?m=free), ví dụ:
```yaml
icon: fas fa-globe
```

```bs-alert info flex
<i class="fas fa-circle-info fa-2xl me-3"></i>
`DO CMS runtime` hỗ trợ bộ biểu tượng [FontAwesome](https://fontawesome.com/search?m=free) và [Bootstrap](https://icons.getbootstrap.com/).
```

```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-3"></i>
Bộ biểu tượng FontAwesome và Bootstrap hỗ trợ bởi các giao diện đồ hoạ đi kèm với phiên bản bản `DO CMS runtime` gốc. Nếu bạn sử dụng một bộ giao diện đồ hoạ của bên thứ 3, vui lòng kiểm tra tài liệu đi kèm.
```

## languages

`map[language-code:display-label]`, các ngôn ngữ được hỗ trợ trên trang web, ví dụ:
```yaml
languages:
  en: English
  vi: Tiếng Việt
  default: vi # ngôn ngữ mặc định của trang web
```

> Mã ngôn ngữ nên tuân theo chuẩn [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
>
> Dòng `default: vi` chỉ thị rằng `vi` là ngôn ngữ mặc định của trang web.

## description

`string` hoặc `map[language-code:text]`, mô tả ngắn dọn về trang web, ví dụ:
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

## contacts

`map[string:string]`, thông tin liên lạc, ví dụ:
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

`map[language-code:map[string:Array(string)]]`, gom nhóm các thẻ tương tự nhau vào làm một, ví dụ:
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

`map[string:object]`, các thông tin khác ở dạng key-object, ví dụ:
```yaml
tags:
  build: ${build_datetime}
  tags1:
    key1: value 1
    key2: value 2
  tags2: [1, 2, 3]
```

> Các chuỗi đặc biệt sẽ được thay bằng giá trị thực trong quá trình tiền xử lý với công cụ `DO CLI`:
> - `${build_datetime}` - thay thế bằng thời gian chạy tiền xửu lý ở định dạng`YYYYMMDDHHmmss`
> - `${build_date}` - pre-processing timestamp, format `YYYYMMDD`
> - `${build_time}` - pre-processing timestamp, format `HHmmss`

```bs-alert info flex
<i class="fas fa-circle-info me-3 fa-xl"></i>
Các tag có thể được nhúng vào trong tài liệu sử dụng cú pháp `[[do-tag`. Tham khảo chi tiết ở tài liệu [Cú pháp Markdown được hỗ trợ](../markdown/).
```

## mode

`string`, chế độ hoạt động của trang web, các giá trị hợp lệ là `document` và `blog`, ví dụ:
```yaml
mode: document
```

> Chế độ mặc định là `document`.

Ở chế độ `blog`, các chức năng sau sẽ được bật:
- API `/api/documents` hỗ trợ thêm 2 loại request:
  - `?p=latest`: trả về danh sách N bài viết mới nhất.
  - `?p=special`: trả về danh sách các tài liệu cho mục đích chuyên biệt (không phải là bài blog).
- Các bài viết được sắp xếp theo thời gian tạo, mới nhất lên trước.
- `/feeds` trả về danh sách các bài viết mới nhất ở dạng RSS newsfeed.

## author

`map[string:string]`, thông tin tác giả mặc định của các bài viết, dùng trong chế độ `blog`, ví dụ:
```yaml
author:
  name: Thanh Nguyen
  avatar: //www.gravatar.com/avatar/9d6ee977a36db465f103ea5c0e4b859c
```

## Ví dụ

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

Xem thêm:
- [Tập tin metadata chủ đề](../topicmetadata/)
- [Tập tin metadata bài viết](../documentmetadata/)
