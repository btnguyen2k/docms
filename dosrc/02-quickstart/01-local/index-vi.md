Cách đơn giản nhất để bắt đầu với DO CMS ở môi trường local là sử dụng Docker và 1 trình soạn thảo văn bản (nếu hỗ trợ Markdown thì càng tốt).

- Đầu tiên, tạo một cấu trúc thư mục để lưu nội dung trang web. Chi tiết về cấu trúc thư mục sẽ được giới thiệu ở ngay phần sau.
- Kế tiếp, bắt đầu tạo nội dung cho trang web với trình soạn thảo văn bản ưa thích của bạn. DO CMS hỗ trợ <a href="https://github.github.com/gfm/" target="_blank">GitHub Flavored Markdown</a> và các công thức Toán học cũng như Hoá học.
- Cuối cùng, khởi chạy một bản `DO CMS runtime` ở dạng Docker container với thư mục nội dung được mount vào trong container.

## Cấu trúc thư mục nội dung

Một ví dụ của cây thư mục nội dung:

```plain
dodata/
├── 01-intro/
│   ├── 01-whatitis/
│   │   ├── index-en.md
│   │   ├── index-vi.md
│   │   └── meta.yaml
│   ├── 02-howitwork/
│   │   ├── docms-howitwork.svg
│   │   ├── index-en.md
│   │   ├── index-vi.md
│   │   └── meta.yaml
│   ├── 03-markdown/
│   │   ├── index-en.md
│   │   ├── index-vi.md
│   │   └── meta.yaml
│   └── meta.yaml
├── 02-quickstart/
│   ...
│   └── meta.yaml
├── 03-components/
│   ...
└── meta.yaml
```

- Cấp trên cùng là `thư mục gốc` chứa nội dung trang web.
- Cấp tiếp theo là các `thư mục chủ đề` chứa các bài viết nằm trong chủ đề.
  - `thư mục chủ đề` được đặt tên theo mẫu `(\d+)-(\w+)`. Phần đầu là một số tự nhiên qui định thứ tự hiển thị của chủ đề, phần cuối là id định danh. Hai phần được phân cách với nhau bằng ký tự gạch ngang (`-`).
  - Id định danh chủ đề phải là duy nhất, và nên được đặt hoàn toàn bằng chữ thường.
- Cấp cuối là `thư mục bài viết` lưu các tập tin của bài viết (ví dụ như các tập tin Markdown và hình ảnh).
  - `thư mục bài viết` được đặt tên theo mẫu `(\d+)-(\w+)`. Phần đầu là một số tự nhiên qui định thứ tự hiển thị của bài viết, phần cuối là id định danh. Hai phần được phân cách với nhau bằng ký tự gạch ngang (`-`).
  - Id định danh bài viết phải là duy nhất trong chủ đề, và nên được đặt hoàn toàn bằng chữ thường.
- `DO CMS runtime`, mặc định đọc nội dung trang web trong thư mục có tên `dodata`. Do vậy có thể đặt tên thư mục gốc là `dodata` cho tiện.
- Trong `thư mục gốc`, `thư mục chủ đề` and `thư mục tài liệu` có một tập tin `meta.yaml`:
  - `meta.yaml` trong `thư mục gốc` lưu *metadata* của trang web (ví dụ như tên, mô tả và thông tin liên hệ).
  - `meta.yaml` trong `thư mục chủ đề` lưu *metadata* của chủ đề (ví dụ như tiêu đề hoặc mô tả).
  - `meta.yaml` trong `thư mục tài liệu` lưu *metadata* của tài liệu (ví dụ như tiêu đề hoặc tóm tắt).

> Metadata cũng có thể được lưu trong tập tin `meta.json`. Nếu cả 2 tập tin `meta.yaml` và `meta.json` cùng tồn tại, tập tin định dạng YAML sẽ được ưu tiên.

## Thể hiện nội dung trang web với DO CMS runtime

Cách đơn giản nhất là khởi chạy phiên bản Docker của `DO CMS runtime`. Lưu ý: nhớ mount `thư mục gốc` vào container khi chạy.

Sau đây là câu lệnh mẫu để khởi chạy `DO CMS runtime` ở dạng Docker container:

```shell
docker run -p 8000:8000 -v ./dodata:/app/dodata btnguyen2k/docmsrutime:latest
```

Sau đó mở địa chỉ http://localhost:8000 trong một trình duyệt để truy cập vào trang web.

Xem thông tin chi tiết về cấu trúc cây thư mục và các tập tin metadata [ở đây](../../reference/metafile/).
