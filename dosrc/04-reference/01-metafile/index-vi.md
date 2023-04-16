Tóm tắt nhanh về cấu trúc thư mục lưu trữ nội dung trang web:

- Cấp trên cùng là `thư mục gốc` chứa nội dung trang web.
- Cấp tiếp theo là các `thư mục chủ đề` chứa các bài viết nằm trong chủ đề.
- Cấp cuối là `thư mục bài viết` lưu các tập tin của bài viết (ví dụ như các tập tin Markdown và hình ảnh).

Mỗi thư mục có một _tập tin metadata file_ ngay bên dưới:
- Tập tin metadata của `thư mục gốc` lưu các thông tin như tên, mô tả và địa chỉ liên lạc của trang web; tập tin metadata của `thư mục chủ đề` lưu các thông tin như tiêu đề hay mô tả của chủ đề; và tập tin metadata của `thư mục bài viết` lưu các thông tin về bài viết như tiêu đề hay nội dung tóm tắt.
- Tập tin metadata được lưu ở định dạng YAML (tên `meta.yaml` hoặc `meta.yml`) hoặc JSON (tên `meta.json`). Nếu đồng thời có nhiều tập tin, `meta.yaml` sẽ được ưu tiên trước nhất, kế đến là `meta.yml` và `meta.json` là cuối cùng.

Một ví dụ của cây thư mục nội dung:

```
dodata/                      <-- thư mục gốc
├── 01-intro/                <---- thư mục chủ đề, id của chủ đề này là "intro"
│   ├── 01-whatitis/         <------ thư mục bài viết, bài viết này nằm trong chủ đề "intro", id của bài viết này là "whatitis"
│   │   ├── index-en.md      <-------- nội dung bài viết (ngôn ngữ "en")
│   │   ├── index-vi.md      <-------- nội dung bài viết (ngôn ngữ "vi")
│   │   └── meta.yaml        <-------- tập tin metadata của bài viết "whatitis"
│   ├── 02-howitwork/        <------ thư mục bài viết
│   │   ...
│   │   └── meta.yaml        <-------- tập tin metadata của bài viết "howitwork"
│   ├── 03-markdown/         <------ thư mục bài viết
│   │   ...
│   │   └── meta.yaml        <-------- tập tin metadata của bài viết "markdown"
│   └── meta.yaml            <------ tập tin metadata của chủ đề "intro"
├── 02-quickstart/           <---- thư mục chủ đề
│   ...
│   └── meta.yaml            <------ tập tin metadata của chủ đề "quickstart"
├── 03-components/           <---- thư mục chủ đề
│   ...
└── meta.yaml                <---- tập tin metadata của trang web, nằm ngay dưới thư mục gốc
```

Xem thêm:
- [Tập tin metadata trang web](../sitemetadata/)
- [Tập tin metadata chủ đề](../topicmetadata/)
- [Tập tin metadata bài viết](../documentmetadata/)
