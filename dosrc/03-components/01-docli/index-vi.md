`DOCLI` là một công cụ dòng lệnh để tiền xử lý nội dung trang web. Các tính năng chính bao gồm:
- Kiểm tra lỗi trong metadata của nội dung trang web.
- Xây dựng chỉ mục toàn văn (fulltext index).

## Cài đặt

Đầu tiên, [tải](https://go.dev/doc/install) và cài đặt Go. Sau đó cài đặt `DOCLI` bằng câu lệnh sau:
```shell
go install github.com/btnguyen2k/docms/docli@latest
```

Cài đặt chỉ định một phiên bản cụ thể (*) qua câu lệnh:
```shell
go install github.com/btnguyen2k/docms/docli@cli-v0.3.1.1
```

```bs-alert warning

`DOCLI` được cài đặt vào thư mục `$GOPATH/bin`. Lưu ý đưa thư mục này vào đường dẫn hệ thống.

(*) Danh sách các phiên bản có thể tìm thấy [ở đây](https://github.com/btnguyen2k/docms/tags).
```

## Cách sử dụng

### Hướng dẫn

Lệnh `docli help` hiển thị danh sách các lệnh được hỗ trợ, và `docli help lệnh` hiển thị hướng dẫn sử dụng cho 1 lệnh cụ thể. Ví dụ:

```shell
$ docli help
NAME:
   docli - DO CMS website content preprocessing tool

USAGE:
   docli [global options] command [command options] [arguments...]

VERSION:
   0.3.1.1

AUTHOR:
   Thanh Nguyen <btnguyen2k (at) gmail (dot) com>

COMMANDS:
   build, b           Preprocess website content, ready for DO CMS runtime
   new, n, create, c  Helper to create assets with default metadata
   touch, t           Touch document metadata file to update timestamp
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   Copyright (c) 2023 - DO CMS
```

### Tiền xử lý nội dung trang web

Nội dung trang web được tiền xử lý với lệnh `docli build`. Ví dụ:

```shell
docli build --src dosrc --out dodata
```

`docli help build` hiển thị phần trợ giúp của lệnh `build`:

```shell
$ docli help build
NAME:
   docli build - Preprocess website content, ready for DO CMS runtime

USAGE:
   docli build [command options] [arguments...]

OPTIONS:
   --src value  source directory
   --out value  output directory
   --purge      purge output directory if not empty (default: false)
   --help, -h   show help
```

Các tham số của lệnh:

|Tham số|Bắt buộc|Giá trị mặc định|Mô tả|
|---|---|---|---|
|`--src source_dir`|không|`dosrc`|thư mục nguồn nơi lưu nội dung trang web|
|`--out output_dir`|không|`dodata`|thư mục đích lưu dữ liệu sau khi tiền xử lý|
|`--purge`|không||xoá sạch thư mục đích trước khi bắt đầu xử lý|

### Hỗ trợ tạo metadata

Từ phiên bản `cli-v0.1.1`, `DO CLI` cung cấp một số tiện ích giúp tạo metadata cho nội dung trang web qua lệnh `docli new`:

- `docli new [global options] site     [command options] [arguments...]`: tạo metadata cho nội dung trang web
- `docli new [global options] topic    [command options] [arguments...]`: tạo metada cho chủ đề
- `docli new [global options] document [command options] [arguments...]`: tạo metadata cho bài viết

**Tham số chung**

Lệnh `new` có các tham số chung sau, áp dụng cho toàn bộ các lệnh con:

|Tham số chung|Bắt buộc|Giá trị mặc định|Mô tả|
|---|---|---|---|
|`--dir value`|không|`doscr`|thư mục lưu nội các tập tin metadata|
|`--override`|không||ghi đè nội dung mới nếu tập tin metadata đã tồn tại|

```bs-alert warning

Các tham số chung phải được đặt _trước_ lệnh con (tức là ngay sau `new`).
```

**Lệnh `new site`**

Hỗ trợ tạo metadata tổng cho nội dung trang web:

```shell
docli new --override --dir doscr site --name "My Blog"
```

Các tham số của lệnh:

|Tham số|Bắt buộc|Giá trị mặc định|Mô tả|
|---|---|---|---|
|`--name value`|true||tên (ngắn gọn) của trang web|
|`--icon value`|false|`fas fa-globe`|icon của trang web (hỗ trợ icon [FontAwesome](https://fontawesome.com/search?m=free))|
|<span class="text-nowrap">`--languages value`</span>|false|`en:english,default:en`|các ngôn ngữ mà trang web hỗ trợ, định dạng `<mã1:tên1>[,<mã2:tên2>...],<default:mã>`|
|`--mode value`|false|`document`|chế độ hoạt động của trang web, các giá trị hợp lệ `document` and `blog`|
|`--author.name`|false|tên tác giả mặc định của trang web|
|`--author.email`|false|email tác giả mặc định của trang web|
|`--author.avatar`|false|đường dẫn tới ảnh đại diện tác giả mặc định của trang web|

```bs-alert warning

Mã ngôn ngữ nên tuân theo chuẩn [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
```

**Lệnh `new topic`**

Hỗ trợ tạo metadata cho chủ đề:

```shell
docli new --override --dir doscr topic --id topic1
```

Các tham số của lệnh:

|Tham số|Bắt buộc|Giá trị mặc định|Mô tả|
|---|---|---|---|
|`--id value`|true||định danh của chủ đề|
|<span class="text-nowrap">`--icon value`</span>|false|`fas fa-book`|icon của chủ đề (hỗ trợ icon [FontAwesome](https://fontawesome.com/search?m=free))|
|`--hidden`|false| |nếu thiết lập 'hidden', chủ đề không hiển thị trong danh sách trên frontend|
|`--img value`|false| |đường dẫn tới ảnh bìa của chủ đề|

```bs-alert warning

Định danh của chủ đề phải là duy nhất (không trùng lắp), và chỉ nên chứa chữ cái thường (`a-z`) và chữ số (`0-9`).

Tên thư mục lưu nôi dung chủ đề sẽ được tự động thêm tiền tố là 1 chuỗi số và ký tự `-` (định dạng `\d+-`). Chuỗi số tiền tố được dùng để xác định thứ tự hiển thị của chủ đề trên trang web.
```

**Lệnh `new document`**

Hỗ trợ tạo meetadata cho bài viết:

```shell
docli new --override --dir doscr document --id doc1 --topic topic1
```

Các tham số của lệnh:

|Tham số|Bắt buộc|Giá trị mặc định|Mô tả|
|---|---|---|---|
|`--topic value`|true||định danh của chủ đề chứa bài viết|
|`--id value`|true||định danh của bài viết|
|`--icon value`|false|`fas fa-file`|icon của bài viết (hỗ trợ icon [FontAwesome](https://fontawesome.com/search?m=free))|
|<span class="text-nowrap">`--use-timestamp`</span>|false||sử dụng thời gian hiện tại (định dạng `yyyyMMddHHmm`) gắn vào trước tên thư mục bài viết, xem thêm bên dưới (*)|
|`--img value`|false| |đường dẫn tới ảnh bìa của bài viết|
|`--page value`|false| |thiết lập bài viết thành 1 trang đặc biệt|
|`--style value`|false| |thiết lập phong cách hiển thị riêng cho bài viết|
|`--author.name`|false|tên tác giả của bài viết|
|`--author.email`|false|email tác giả của tài liệu|
|`--author.avatar`|false|đường dẫn tới ảnh đại diện tác giả của tài liệu|

```bs-alert warning

Định danh của bài viết phải là duy nhất trong chủ đề (trong cùng 1 chủ đề không có bài viết trùng lắp định danh), và chỉ nên chứa chữ cái thường (`a-z`) và chữ số (`0-9`).

(*) Tên thư mục lưu nôi dung bài viết sẽ được tự động thêm tiền tố là 1 chuỗi số và ký tự `-` (định dạng `\d+-`). Chuỗi số tiền tố được dùng để xác định thứ tự hiển thị của bài viết trên trang web. Nếu tham số `--use-timestamp` được sử dụng, thời gian hiện tại (định dạng `yyyyMMddHHmm`) sẽ được sử dụng làm chuỗi số tiền tố.
```

### Các tính năng khác

- `docli touch --topic <topic-id> --id <doc-id>`: cập nhật thời gian `last-updated` của tài liệu.
