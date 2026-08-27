# Go by Tests

Goのゼロ値、`nil`、`error`、`slice`、`channel`、`context`など、実務で混乱しやすい挙動を**実行可能なGoソースと標準`testing`パッケージのテスト**で確認する静的リファレンスです。

> **設計原則**: 説明より先に、SourceとTestを並べて読む。サイトに表示するコードは、同じリポジトリで`go test ./...`により実行されます。

## 内容

現在は**34テーマ**を収録しています。UTF-8のバイト数と`rune`、`range`のバイト位置・開始時の反復長、`strings.Builder`のコピー、stringと`[]byte`、`append`とバッキング配列、`slices.Clone`、`nil slice`と空slice、mapのゼロ値とnil map、typed nil interface、比較不能なslice、error wrapping・`errors.Join`・`errors.As`、`defer`とcloseエラー・`recover`の直接呼び出し、context cancellation・deadline・値キー・キャンセル原因、closed/nil channel・close後の送信、`sync.Once`、`sync.Map`、`time.Time.Equal`と`time.Duration`、HTTP Response Body、JSONの`omitempty`・非公開フィールド・未知フィールド、Go 1.22のgoroutine range変数を扱います。

## ローカル実行

```bash
git clone https://github.com/tonbiattack/go-by-tests.git
cd go-by-tests
go test ./...
pnpm install --frozen-lockfile
pnpm dev
```

Goコードは`gofmt`で整形します。静的サイトはAstroで生成し、GitHub ActionsではGoフォーマット、Goテスト、Astro型検査・ビルド、GitHub Pagesリンク検査を行います。

## 構成

```text
examples/                 実行されるGoソースと_test.go
src/data/lessons.ts       表示する教材のメタデータと実ファイル参照
src/pages/go/             静的教材ページ
scripts/verify-pages-links.mjs
.github/workflows/        CIとGitHub Pagesデプロイ
```

## ライセンス

このリポジトリのライセンスは今後追加予定です。教材を利用・再配布する前に、ライセンスファイルの追加を確認してください。
