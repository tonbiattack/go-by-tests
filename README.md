# Go by Tests

[![Go examples](https://github.com/tonbiattack/go-by-tests/actions/workflows/ci.yml/badge.svg)](https://github.com/tonbiattack/go-by-tests/actions/workflows/ci.yml)
[![GitHub Pages](https://img.shields.io/badge/demo-GitHub%20Pages-181717?logo=github)](https://tonbiattack.github.io/go-by-tests/)

> **公開ページ**: [Go by Tests をブラウザで開く](https://tonbiattack.github.io/go-by-tests/)

**Go by Tests** は、Go 1.22 と標準 `testing` パッケージの **34 テーマ**を通じて、実務で見落としやすい挙動を確かめる実行可能ドキュメント教材です。説明を読む前に、短い Source と Test の組を並べて読み、期待値・例外・状態変化をコードで観測します。

> **設計原則**: ブラウザに表示するコードと `go test ./...` で実行するコードを同じファイルに保つ。説明だけ、テストだけ、表示だけを独立させず、観測可能な事実を一つのリポジトリに残します。

| 提供するもの | 内容 | 利用方法 |
|---|---|---|
| 実行可能な教材 | Goの言語仕様・標準ライブラリの挙動を、最小のSource/Testで記録する。 | [`examples/`](./examples/) を `go test` で実行する。 |
| 学習用サイト | SourceとTest、要点、観測結果を対比して読む。 | [公開ページ](https://tonbiattack.github.io/go-by-tests/) を開く。 |
| 品質ゲート | 整形、テスト、型検査、静的ビルド、リンク検査を自動化する。 | [GitHub Actions](https://github.com/tonbiattack/go-by-tests/actions) で確認する。 |

## 学習マップ

全 **34 テーマ**を7カテゴリに分けています。初めての場合は **Language → Collections → Error Handling → Concurrency** の順に進むと、値の表現、可変状態、失敗、並行処理を段階的に確認できます。気になる挙動から直接始めても構いません。

| カテゴリ | テーマ数 | まず読むテーマ | 身に付ける問い |
|---|---:|---|---|
| Language | 7 | [`string: len`](https://tonbiattack.github.io/go-by-tests/go/string/utf8-length/) | `string`、`rune`、`interface`、`range` の実行時の扱いをどう区別するか。 |
| Collections | 5 | [`slice: append`](https://tonbiattack.github.io/go-by-tests/go/slice/append-backing-array/) | sliceの共有、`nil`、mapのゼロ値、コピーをどう扱うか。 |
| Error Handling | 6 | [`error: wrapping`](https://tonbiattack.github.io/go-by-tests/go/errors/wrap-and-is/) | ラップした原因、`defer`、`recover`、closeエラーをどう観測するか。 |
| Concurrency | 10 | [`context: cancel`](https://tonbiattack.github.io/go-by-tests/go/context/cancel/) | `context`、channel、`sync.Once`、goroutineの境界をどう設計するか。 |
| Date / Time | 2 | [`time.Time: Equal`](https://tonbiattack.github.io/go-by-tests/go/time/equal-vs-operator/) | 同じ瞬間・期間をどの比較と単位で扱うか。 |
| Encoding | 3 | [`encoding/json: omitempty`](https://tonbiattack.github.io/go-by-tests/go/encoding/json-omitempty/) | JSONのゼロ値、非公開フィールド、未知フィールドをどう扱うか。 |
| HTTP | 1 | [`http.Response.Body`](https://tonbiattack.github.io/go-by-tests/go/http/response-body-close/) | HTTPレスポンスのリソースをいつ解放するか。 |

各テーマの Source と Test は [公開ページ](https://tonbiattack.github.io/go-by-tests/) で対比して読めます。学習の進め方は [LEARNING_PATH.md](./LEARNING_PATH.md) にまとめています。

## すぐに始める

### 必要条件

Go **1.22 以上**、Node.js **22 以上**、および pnpm を用意してください。Goの標準テストは `go test` で実行し、教材サイトは Astro で静的生成します。[1] [2] [3]

### インストールと閲覧

```bash
git clone https://github.com/tonbiattack/go-by-tests.git
cd go-by-tests
pnpm install --frozen-lockfile
pnpm dev
```

ローカルサーバーを起動したら、表示されたURLで教材サイトを開きます。公開済みのサイトだけを閲覧する場合は、[Go by Tests](https://tonbiattack.github.io/go-by-tests/) を直接開いてください。

### 実行可能なGo教材を検証する

```bash
go test ./...
```

### 品質ゲートを個別に実行する

| コマンド | 目的 |
|---|---|
| `go test ./...` | 全34テーマの期待値、panic、エラー、状態遷移を検証する。 |
| `gofmt -w $(find examples -type f -name '*.go' -print)` | `examples/` 配下のGoコードを整形する。 |
| `pnpm check` | Astroコンポーネントと教材メタデータの型整合性を確認する。 |
| `GITHUB_ACTIONS=true pnpm build` | GitHub Pages用の `/go-by-tests/` サブパスで静的ビルドする。 |
| `pnpm verify:links` | 静的出力に含まれる内部リンクを検査する。 |

GitHub Actionsでは、Goの整形とテスト、Astroの型検査と静的ビルド、GitHub Pages用リンク検査を順に実行します。各品質ゲートの結果は [Actions](https://github.com/tonbiattack/go-by-tests/actions) で確認できます。

## このリポジトリが保証すること

ブラウザに表示する Source/Test と、ローカルで `go test` が実行するコードは同じ `.go` ファイルです。Astro はビルド時に `examples/` を読み込み、テーマ、タグ、観測結果、Source/Testの対応だけを `src/data/lessons.ts` で管理します。これにより、説明用コードと検証対象の乖離を防ぎます。

```text
Browser
  ↓
GitHub Pages
  ↑
Astro static build ──→ src/pages + src/components
  ↑                         ↑
src/data/lessons.ts ─────── examples/**/*.go
                                  ↑
                         gofmt + go test
```

| 品質ゲート | 守る契約 | 自動化場所 |
|---|---|---|
| gofmt | すべての教材コードが標準フォーマットに従う。 | Go / GitHub Actions |
| go test | 期待値、panic、error、リソース解放、並行処理の状態を固定する。 | Go / GitHub Actions |
| Astro check | 表示コンポーネントと教材メタデータの整合性を確認する。 | Astro / GitHub Actions |
| Static build | GitHub Pagesで公開可能なHTMLを生成する。 | Astro / GitHub Actions |
| Link verification | サブパス下の内部リンクが実在することを確認する。 | Node.js / GitHub Actions |

## リポジトリ構成

```text
.
├── examples/                 # 実行されるGoソースと_test.go
├── src/
│   ├── data/lessons.ts       # テーマ、Source/Testの対応、観測結果
│   ├── pages/go/             # 静的教材ページ
│   └── components/           # ナビゲーションとコード比較UI
├── scripts/verify-pages-links.mjs
├── .github/workflows/        # CIとGitHub Pagesデプロイ
├── LEARNING_PATH.md          # 初めて読むときの順序と観測方法
└── CONTRIBUTING.md           # 教材を追加・更新するときの約束
```

## 姉妹教材

同じ「短い Source/Test の対で、見落としやすい挙動を証明する」方針で、Java版・TypeScript版も公開しています。各言語の固有の仕様を扱うため、内容をそのまま複製せず、同じ設計上の問いをそれぞれのランタイムで確認します。

| 教材 | リポジトリ | 公開ページ |
|---|---|---|
| Java by Tests | [GitHub](https://github.com/tonbiattack/java-by-tests) | [Live Demo](https://tonbiattack.github.io/java-by-tests/) |
| TypeScript by Tests | [GitHub](https://github.com/tonbiattack/typescript-by-tests) | [Live Demo](https://tonbiattack.github.io/typescript-by-tests/) |
| Go by Tests | [GitHub](https://github.com/tonbiattack/go-by-tests) | [Live Demo](https://tonbiattack.github.io/go-by-tests/) |

## 新しいテーマを提案・追加する

新しい教材は、テーマ数を増やすためではなく、**一つの見落としやすい契約を再現可能な証拠として残すため**に追加します。追加・更新時には、Source、Test、`lessons.ts` のメタデータ、観測結果を同じ変更で同期してください。具体的な手順と確認項目は [CONTRIBUTING.md](./CONTRIBUTING.md) を参照してください。提案や不具合は [Issues](https://github.com/tonbiattack/go-by-tests/issues) で受け付けます。

## ライセンス

このリポジトリには現時点でライセンスファイルを設定していません。再利用、配布、派生物の作成を予定する場合は、リポジトリ所有者へ確認してください。

## References

[1] [The Go Programming Language — Testing](https://go.dev/pkg/testing/)
[2] [The Go Programming Language — gofmt](https://pkg.go.dev/cmd/gofmt)
[3] [Astro Documentation](https://docs.astro.build/)
[4] [GitHub Pages — Custom GitHub Actions Workflows](https://docs.github.com/pages/getting-started-with-github-pages/using-custom-workflows-with-github-pages)
