# DOCMS Release Notes

## 2023-04-xx

Runtime `v0.2.0`:
- Support [KaTeX](https://katex.org/) block with syntax _```katex_
- Links to external websites will open new windows/tabs.
- Support [Bootstrap Alerts](https://getbootstrap.com/docs/5.3/components/alerts/) with syntax ```bs-alert

## 2023-04-13

CLI `v0.1.1`:
- Command `new` - helper to create assets with default metadata.
  - `new site`: Create new site content metadata
  - `new topic`: Create new topic metadata
  - `new document`: Create new document metadata

Runtime `v0.1.1`:
- Bug fixes.

## 2023-04-12

CLI `v0.1.0`:
- Pre-process website content with command `docli build`.
- Build fulltext index while pre-processing website content.

Runtime `v0.1.0`:
- Render website content, supporting 3 themes: `bootstrap` (default), `coderdocs` and `prettydocs`.
- Fulltext search.
- Multi-languages support.
