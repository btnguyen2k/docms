Bạn có thể nhận thấy rằng một vài tính năng của DO CMS không hoạt động (ví dụ tìm kiếm toàn văn - fulltext search) khi nội dung trang web được thể hiện từ `thư mục gốc`. Để mở khoá toàn bộ tính năng, nội dung trang web cần được tiền xử lý trước khi thể hiện bởi `DO CMS runtime`.

Nội dung trang web được tiền xử lý thông qua công cụ `DOCLI`. Cài đặt `DOCLI` với lệnh sau:

```shell
docli build --src ./dosrc --out ./dodata
```

Bạn hãy lưu nội dung trang web vào thư mục `dosrc` và chạy câu lệnh sau để tiền xử lý:

```shell
docli build --src ./dosrc --out ./dodata
```

`DOCLI` thực hiện các bước sau:
- Kiểm tra nội dung trang web ở thư mục nguồn (`dosrc`), báo lỗi nếu có.
- Sao chép nội dung trang web từ thư mục nguồn (`dosrc`) sang thư mục đích (`dodata`).
- Xây dựng chỉ mục toàn văn (fulltext index) và lưu vào thư mục đích (`dodata`).

Sau khi tiền xử lý, thể hiện nội dung trang web từ thư mục `dodata` sẽ mở khóa tất cả các tính năng của DO CMS.

```bs-alert warning

`DOCLI` được cài đặt vào thư mục `$GOPATH/bin`. Lưu ý đưa thư mục này vào đường dẫn hệ thống.
```

Xem thêm:
- [Công cụ DOCLI](../../components/cli/)
- [DO CMS runtime](../../components/runtime/)
