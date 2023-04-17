Đã đến lúc xuất bản trang web của bạn để người dùng cuối truy cập. Cách nhanh và dễ nhất là đóng gói thư mục `dodata` chung với `DO CMS runtime` trong 1 Docker image. Sau đó triển khai image này lên một môi trường container.

`DO CMS runtime` đã được đóng gói sẵn ở dạng Docker image, đoạn script Dockerfile bên dưới sẽ tạo 1 image mới bao gồm thư mục `dodata` đóng gói chung với `DO CMS runtime`:

```dockerfile
FROM btnguyen2k/docmsruntime:latest as docmsruntime
LABEL maintainer="Your name"
COPY ./dodata /app/dodata
WORKDIR /app
EXPOSE 8000
CMD ["/app/main"]
```

Nếu nội dung trang web cần qua bước tiền xử lý trước khi đóng gói:

```dockerfile
FROM golang:1.17-alpine AS builder_docs
RUN mkdir /build
COPY . /build
RUN cd /build \
    && go install github.com/btnguyen2k/docms/docli \
    && docli build --purge --src dosrc --out dodata

FROM btnguyen2k/docmsruntime:latest as docmsruntime
LABEL maintainer="Your name"
COPY --from=builder_docs /build/dodata /app/dodata
WORKDIR /app
EXPOSE 8000
CMD ["/app/main"]
```

Các bước trên có thể được thực hiện bằng cách gõ lệnh tay, hoặc tự động thông qua qui trình CI/CD. Bạn có thể tham khảo ví dụ sử dụng GitHub Actions để tiền xử lý, đóng gói nội dung trang web và triển khai lên Azure Container App [ở đây](https://github.com/btnguyen2k/docms/blob/main/.github/workflows/dodocs.yml).

Xem thêm:
- [Công cụ DOCLI](../../components/cli/)
- [DO CMS runtime](../../components/runtime/)
