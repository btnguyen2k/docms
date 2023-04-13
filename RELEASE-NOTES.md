# DOCMS Release Notes

## 2023-04-13 - v0.1.1

CLI `v0.1.1`:
- Command `new` - helper to create assets with default metadata.
  - `new site`: Create new site content metadata
  - `new topic`: Create new topic metadata
  - `new document`: Create new document metadata

Runtime `v0.1.1`:
- Bug fixes.

## 2023-04-12 - v0.1.0

CLI:
- Pre-process website content with command `docli build`.
- Build fulltext index while pre-processing website content.

Runtime:
- Render website content, supporting 3 themes: `bootstrap` (default), `coderdocs` and `prettydocs`.
- Fulltext search.
- Multi-languages support.
