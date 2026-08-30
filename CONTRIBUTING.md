# Contributing to Go by Tests

Go by Tests は、Goの挙動を短い Source/Test の対で記録する教材です。変更の目的はテーマ数を増やすことではなく、**一つの見落としやすい契約を再現可能な証拠として残すこと**です。

## 変更の最小単位

新しい教材または既存教材の修正では、次の四点を同じ変更に含めてください。

| 対象 | 役割 | 配置 |
|---|---|---|
| Source | 確認したい挙動を最小限に表すGo実装 | `examples/` |
| Test | 正常系、境界値、panic、error、状態変化を固定する標準 `testing` テスト | `examples/` の `*_test.go` |
| Metadata | タイトル、カテゴリ、タグ、Source/Testの対応、観測結果 | `src/data/lessons.ts` |
| 説明 | 何を確認するか、なぜ見落としやすいか | `src/data/lessons.ts` の `oneLine`、`checks`、`note` |

Sourceだけ、Testだけ、表示データだけを追加する変更は受け入れません。ブラウザに表示するコードと `go test ./...` で検証するコードを同一に保つことが、このリポジトリの中心的な制約です。

## テーマの選び方

良いテーマは、短く、観測可能で、再現可能です。`nil`、ゼロ値、sliceの共有、map、`error`、`context`、channel、`time.Time`、JSONなどで、「見た目は似ているが実行時の契約が異なる」一点を選びます。外部サービス、実時間、ネットワーク、ランダム値には依存させず、テスト単体で同じ結果を再現できるようにします。

| 確認すること | 例 |
|---|---|
| 正常系 | 入力に対して期待する値・順序・状態を返すか。 |
| 境界値 | 空文字、空slice、ゼロ値、`nil`、閾値前後でどう振る舞うか。 |
| 失敗 | 期待するpanicやerrorをどのように検証するか。 |
| 不変条件 | sliceの共有、リソース解放、contextのキャンセル状態を守るか。 |

## 実装手順

1. Issueまたはプルリクエストの説明に、確認したい契約を一文で書きます。
2. `*_test.go` を先に追加し、失敗することを確認します。
3. テストを通す最小のSourceを追加します。
4. `lessons.ts` にメタデータと、期待値・実際値・panic・errorのいずれを観測するかを記録します。
5. Goコードを整形し、ローカルで品質ゲートを通します。

```bash
gofmt -w $(find examples -type f -name '*.go' -print)
go test ./...
pnpm install --frozen-lockfile
pnpm check
GITHUB_ACTIONS=true pnpm build
pnpm verify:links
```

## プルリクエストの確認項目

| 確認項目 | 合格条件 |
|---|---|
| 焦点 | 一つの言語仕様またはAPI契約に焦点を絞っている。 |
| テスト | 正常系に加え、境界値または失敗時の挙動を確認している。 |
| 表示との一致 | Source/Testのパス、ファイル名、観測結果が実在するコードと一致する。 |
| 実行性 | `go test ./...`、`pnpm check`、静的ビルド、リンク検査が成功する。 |
| 文書 | タイトル、説明、タグがGoの言語仕様またはAPIの挙動を正確に表す。 |
| 独立性 | 実ネットワーク、実時刻、外部資格情報へ依存しない。 |

## Issueの使い分け

新しい題材の提案、教材コード・表示・リンク・CIに関する不具合は、[Issues](https://github.com/tonbiattack/go-by-tests/issues) で報告してください。提案では、確認したい挙動、最小の再現例、期待する観測結果を記載してください。

## References

[1] [The Go Programming Language — Testing](https://go.dev/pkg/testing/)
[2] [The Go Programming Language — gofmt](https://pkg.go.dev/cmd/gofmt)
