Mỗi một `thư mục bài viết` có một tập tin metadata riêng. Tập tin này chứa các trường thông tin sau:

**icon** - `string`, icon của bài viết, hỗ trợ [icon FontAwesome](https://fontawesome.com/search?m=free), ví dụ:
```yaml
icon: fas fa-file
```

**title** - `string` or `map[language-code:text]`, tiêu đề bài viết, ví dụ:
```yaml
title: "DO CMS là gì"
```

hoặc
```yaml
title:
  en: What it is
  vi: DO CMS là gì
```

Khi `title` là một string, nó sẽ tương đương với `map[ngôn-ngữ-mặc-định:text]`. Điều đó có nghĩa là dòng chỉ thị như sau
```yaml
title: "DO CMS là gì"
```
sẽ tương đương với:
```yaml
title:
  vi: DO CMS là gì
```
bởi vì `vi` đang được chỉ định là ngôn ngữ mặc định của trang web trong [tập tin metadata trang web](../sitemetadata/).

**summary** - `string` or `map[language-code:text]`, tóm tắt ngắn gọn nội dung bài viết, ví dụ:
```yaml
summary: "DO CMS là Hệ thống quản lý nội dung giúp tác giả xuất bản nội dung trang web thông qua luồng CI/CD. Sẽ không có giao diện để người dùng tạo, cập nhật và xuất bản nội dung lên trang web. Thay vào đó, nội dung của trang web sẽ được xây dựng và xuất bản thông qua qui trình CI/CD."
```

hoặc
```yaml
summary:
  en: "DO CMS is a Content Management System that helps authors publish website content through a CI/CD flow. Unlike other CMS, there is no UI to create, update and publish content in DO CMS. Instead, website content is built and published via CI/CD pipelines."
  vi: "DO CMS là Hệ thống quản lý nội dung giúp tác giả xuất bản nội dung trang web thông qua luồng CI/CD. Sẽ không có giao diện để người dùng tạo, cập nhật và xuất bản nội dung lên trang web. Thay vào đó, nội dung của trang web sẽ được xây dựng và xuất bản thông qua qui trình CI/CD."
```

Khi `summary` là một string, nó sẽ tương đương với `map[ngôn-ngữ-mặc-định:text]`. Điều đó có nghĩa là dòng chỉ thị như sau
```yaml
summary: "DO CMS là Hệ thống quản lý nội dung giúp tác giả xuất bản nội dung trang web thông qua luồng CI/CD. Sẽ không có giao diện để người dùng tạo, cập nhật và xuất bản nội dung lên trang web. Thay vào đó, nội dung của trang web sẽ được xây dựng và xuất bản thông qua qui trình CI/CD."
```
sẽ tương đương với:
```yaml
summary:
  vi: "DO CMS là Hệ thống quản lý nội dung giúp tác giả xuất bản nội dung trang web thông qua luồng CI/CD. Sẽ không có giao diện để người dùng tạo, cập nhật và xuất bản nội dung lên trang web. Thay vào đó, nội dung của trang web sẽ được xây dựng và xuất bản thông qua qui trình CI/CD."
```
bởi vì `vi` đang được chỉ định là ngôn ngữ mặc định của trang web trong [tập tin metadata trang web](../sitemetadata/).

**file** - `string` or `map[language-code:tên-tập-tin]`, tập tin chứa nội dung bài viết, ví dụ:
```yaml
file: index.md
```

hoặc
```yaml
file:
  en: index-en.md
  vi: index-vi.md
```

Khi `file` là một string, nó sẽ tương đương với `map[ngôn-ngữ-mặc-định:tên-tập-tin]`. Điều đó có nghĩa là dòng chỉ thị như sau
```yaml
file: index.md
```
sẽ tương đương với:
```yaml
file:
  vi: index-vi.md
```
bởi vì `vi` đang được chỉ định là ngôn ngữ mặc định của trang web trong [tập tin metadata trang web](../sitemetadata/).

Một ví dụ đầy đủ của tập tin metadata bài viết:
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

Xem thêm:
- [Tập tin metadata trang web](../sitemetadata/)
- [Tập tin metadata chủ đề](../topicmetadata/)
