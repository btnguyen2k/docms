DOCLI là một công cụ dòng lệnh để tiền xử lý nội dung trang web. Các tính năng của nó bao gồm:
- Kiểm tra lỗi cho nội dung trang web.
- Xây dựng chỉ mục toàn văn.

## Cài đặt

Đầu tiên, <a href="https://go.dev/doc/install" target="_blank">tải</a> và cài đặt Go. Sau đó cài đặt `DOCLI` bằng câu lệnh sau:
```shell
go install github.com/btnguyen2k/docms/docli@latest
```

> `DOCLI` được cài đặt vào thư mục `$GOPATH/bin`. Lưu ý đưa thư mục này vào đường dẫn hệ thống.

## Sử dụng

**Hướng dẫn sử dụng**

Lệnh `docli help` hiển thị danh sách các lệnh được hỗ trợ, và `docli help lệnh` hiển thị hướng dẫn sử dụng cho 1 lệnh cụ thể. Ví dụ:

```shell
$ docli help
NAME:
   docli - Pre-process and build content for DO CMS

USAGE:
   docli [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   build, b  Build DO CMS data
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

**Tiền xử lý nội dung trang web**

Nội dung trang web được tiền xử lý với lệnh `docli build`. Ví dụ:

```shell
docli build --src dosrc --out dodata
```

Các tham số của lệnh:

|Tham số|Bắt buộc|Giá trị mặc định|Mô tả|
|---|---|---|---|
|`--src source_dir`|không|`dosrc`|thư mục nguồn nơi lưu nội dung trang web|
|`--out output_dir`|không|`dodata`|thư mục đích lưu dữ liệu sau khi tiền xử lý|
|`--purge`|không||xoá sạch thư mục đích trước khi bắt đầu xử lý|
