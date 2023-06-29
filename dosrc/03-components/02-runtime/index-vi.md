`DO CMS runtime` chịu trách nhiệm thể hiện nội dung trang web, và hỗ trợ một số tính năng khác như _i18n_ hoặc _tìm kiếm toàn văn_ (fulltext search).

Để tiện sử dụng, `DO CMS runtime` đã được đóng gói sẵn thành [Docker image](https://hub.docker.com/r/btnguyen2k/docmsruntime).

Câu lệnh sau sẽ chạy một bản `DO CMS runtime` từ Docker image:

```shell
docker run --name docms \
    -p 8000:8000 \
    -v /path/to/my/dodata:/app/dodata \
    btnguyen2k/docmsruntime
```

```bs-alert info flex
<i class="fas fa-circle-info fa-2xl me-3"></i>
Bản Docker image được đánh thẻ như sau:

`:stable` hoặc `:version` (ví dụ `:1.2.3`): bản phát hành ổn định<br/>
`:latest` hoặc `:yyyymmdd` hoặc `:yyyymmdd-hash`: bản phát hành có thể có 1 số lỗi nhỏ, nhưng luôn bao gồm các tính năng mới nhất (tính tới thời điểm build).
```

## Biến môi trường

### HTTP_LISTEN_ADDR/HTTP_LISTEN_PORT

Chỉ định địa chỉ & cổng lắng nghe. Giá trị mặc định của HTTP_LISTEN_ADDR là `0.0.0.0` và HTTP_LISTEN_PORT là `8000`. Câu lệnh sau chi định `DO CMS runtime` lắng nghe `127.0.0.1:3000`:
```shell
docker run --name docms \
    -p 8000:3000 \
    -v /path/to/my/dodata:/app/dodata \
    -e HTTP_LISTEN_ADDR=127.0.0.1 -e HTTP_LISTEN_PORT=3000 \
    btnguyen2k/docmsruntime
```

### GOOGLE_TAG

Chỉ định Google Tag ID để đẩy data lên Google Analytics, ví dụ:
```shell
docker run --name docms \
    -p 8000:8000 \
    -e GOOGLE_TAG=G-ABC123D4E5 \
    btnguyen2k/docmsruntime
```

```bs-alert info

`GOOGLE_TAG` được hỗ trợ từ bản runtime `v0.3.1` hoặc Docker image `:20230426`.
```

### DOCMS_DATA_DIR

Chỉ định thư mục gốc chứa nội dung trang web. Giá trị mặc định là `./dodata`. Câu lệnh sau chỉ định `DO CMS runtime` tải nội dung trang web từ thư mục `/mydata`:
```shell
docker run --name docms \
    -p 8000:8000 \
    -v /path/to/my/dodata:/mydata \
    -e DOCMS_DATA_DIR=/mydata \
    btnguyen2k/docmsruntime
```

### FRONTEND_TEMPLATE

Chỉ định giao diện đồ hoạ cho phần thể hiện nội dung trang web. Hiện tại có 3 giao diện có thể lựa chọn là `bootstrap` (mặc định), `coderdocs` and `prettydocs`. Câu lệnh sau sẽ chỉ định `coderdocs` làm giao diện thể hiện:
```shell
docker run --name docms \
    -p 8000:8000 \
    -e FRONTEND_TEMPLATE=coderdocs \
    btnguyen2k/docmsruntime
```

Giao diện **Bootstrap**:
```bs-cards cols-3 cols-sm-1 equals lightbox=Bootstrap
    [[bs-card
    -img:bootstrap1.png
    -title:Trang nhà
    ]]

    [[bs-card
    -img:bootstrap2.png
    -title:Trang chủ đề
    ]]

    [[bs-card
    -img:bootstrap3.png
    -title:Trang bài viết
    ]]
```

Giao diện **CoderDocs**:
```bs-cards cols-3 cols-sm-1 equals lightbox=CoderDocs
    [[bs-card
    -img:coderdocs1.png
    -title:Trang nhà
    ]]

    [[bs-card
    -img:coderdocs2.png
    -title:Trang chủ đề
    ]]

    [[bs-card
    -img:coderdocs3.png
    -title:Trang bài viết
    ]]
```

Giao diện **PrettyDocs**:
```bs-cards cols-3 cols-sm-1 equals lightbox=PrettyDocs
    [[bs-card
    -img:prettydocs1.png
    -title:Trang nhà
    ]]

    [[bs-card
    -img:prettydocs2.png
    -title:Trang chủ đề
    ]]

    [[bs-card
    -img:prettydocs3.png
    -title:Trang bài viết
    ]]
```

```bs-alert info

Giao diện đồ hoạ `blogy` được thêm vào từ phiên bản `v0.3.1` (hoặc Docker image `:20230511`). `blogy` thích hợp cho các trang web thien về blog (trang web [được tạo](../docli/#hỗ-trợ-tạo-metadata) với tham số `--mode blog`).
```

### MIME_TYPES

Chỉ định các loại định dạng tập tin được phép đính kèm theo tài liệu, giá trị phải là map ở định dạng 1 chuỗi JSON hợp lệ, ví dụ:
```shell
docker run --name docms \
    -p 8000:8000 \
    -e MIME_TYPES='{".jpg": "image/jpeg", ".gif": "image/gif"}' \
    btnguyen2k/docmsruntime
```

```bs-alert info

`MIME_TYPES` được hỗ trợ từ bản runtime `v0.4.0` hoặc Docker image `:20230628`.
```

Các loại định dạng tập tin được cấu hình mặc định:
```json
  ## Media MIME types
  media_mime {
    # photo files
    ".jpg":  "image/jpeg",
    ".jpeg": "image/jpeg",
    ".png":  "image/png",
    ".gif":  "image/gif",
    ".svg":  "image/svg+xml",
    # music files
    ".aac":  "audio/aac",
    ".mp3":  "audio/mp3",
    ".wav":  "audio/wav",
    ".wma":  "audio/wma",
    # video files
    ".mp4":  "audio/mp4",
    ".mov":  "audio/mov",
    ".avi":  "audio/avi",
    # documents
    ".xml":  "application/xml",
    ".json": "application/json",
    ".pdf":  "application/pdf",
    ".csv":  "application/csv",
    ".tsv":  "application/tsv",
    ".doc":  "application/doc",
    ".docx": "application/docx",
    ".xls":  "application/xls",
    ".xlsx": "application/xlsx",
    ".ppt":  "application/ppt",
    ".pptx": "application/pptx",
    ".odm":  "application/odm",
    ".odt":  "application/odt",
    ".odp":  "application/odp",
    ".ods":  "application/ods",
  }
```

### MIME_TYPES_ADD

Thêm định dạng tập tin vào danh sách có sẵn, giá trị phải là map ở định dạng 1 chuỗi JSON hợp lệ, ví dụ:
```shell
docker run --name docms \
    -p 8000:8000 \
    -e MIME_TYPES_ADD='{".mp4": "video/mp4"}' \
    btnguyen2k/docmsruntime
```

```bs-alert info

`MIME_TYPES_ADD` is supported since runtime `v0.4.0` or Docker image `:20230628`.
```
