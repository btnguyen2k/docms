`DO CMS runtime` chịu trách nhiệm thể hiện nội dung trang web, và hỗ trợ một số tính năng khác như i18n hoặc tìm kiếm toàn văn (fulltext search).

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

**HTTP_LISTEN_ADDR** và **HTTP_LISTEN_PORT** - chỉ định địa chỉ & cổng lắng nghe. Giá trị mặc định của HTTP_LISTEN_ADDR là `0.0.0.0` và HTTP_LISTEN_PORT là `8000`. Câu lệnh sau chi định `DO CMS runtime` lắng nghe `127.0.0.1:3000`:
```shell
docker run --name docms \
    -p 8000:3000 \
    -v /path/to/my/dodata:/app/dodata \
    -e HTTP_LISTEN_ADDR=127.0.0.1 -e HTTP_LISTEN_PORT=3000 \
    btnguyen2k/docmsruntime
```

**GOOGLE_TAG** - chỉ định Google Tag ID để đẩy data lên Google Analytics, ví dụ:
```shell
docker run --name docms \
    -p 8000:8000 \
    -e GOOGLE_TAG=G-ABC123D4E5 \
    btnguyen2k/docmsruntime
```

```bs-alert info

`GOOGLE_TAG` được hỗ trợ từ bản runtime `v0.3.1` hoặc Docker image `:20230426`.
```

**DOCMS_DATA_DIR** - chỉ định thư mục gốc chứa nội dung trang web. Giá trị mặc định là `./dodata`. Câu lệnh sau chỉ định `DO CMS runtime` tải nội dung trang web từ thư mục `/mydata`:
```shell
docker run --name docms \
    -p 8000:8000 \
    -v /path/to/my/dodata:/mydata \
    -e DOCMS_DATA_DIR=/mydata \
    btnguyen2k/docmsruntime
```

**FRONTEND_TEMPLATE** - chỉ định giao diện đồ hoạ cho phần thể hiện nội dung trang web. Hiện tại có 3 giao diện có thể lựa chọn là `bootstrap` (mặc định), `coderdocs` and `prettydocs`. Câu lệnh sau sẽ chỉ định `coderdocs` làm giao diện thể hiện:
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
