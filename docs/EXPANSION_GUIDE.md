# Go by Tests 教材拡張指示書

## 目的

Go by Tests を「見落としやすい挙動だけを集めた教材」から、次の方向へ拡張する。

> **Go の言語仕様・標準ライブラリ・実行時の重要な挙動を、実行可能な標準 testing テストから理解する教材**

ただし、Go 入門サイトや文法リファレンスにはしない。

追加判断の最重要基準は、**そのテストを見ることで Go の理解が一段深くなるか**である。

## 教材に追加してよいテーマ

### Pitfall

Go で特に誤解やバグにつながりやすい挙動。

例:

- nil interface と typed nil
- slice の backing array 共有
- `append` による再確保
- map のゼロ値
- pointer receiver / value receiver
- loop variable と closure
- `defer` の引数評価タイミング
- error wrapping

### Behavior

意外ではなくても、Go を実務で扱ううえで理解価値が高く、テストで観測しやすいもの。

例:

- zero value
- array と slice のコピー差
- range の挙動
- map lookup の comma-ok
- `errors.Is` / `errors.As`
- channel の close
- buffered / unbuffered channel
- context cancellation
- JSON の zero value / `omitempty`

### Concept

Go の設計思想や実行モデルを、テスト結果から理解できるもの。

例:

- interface の動的型と動的値
- method set
- embedding
- goroutine
- channel synchronization
- select
- context propagation
- generics / constraints
- comparable
- memory sharing

## 追加しないテーマ

- `1 + 2 == 3` のような単純な構文確認
- `if`、`for`、変数宣言などの基本文法だけを説明する内容
- testing を使う意味がほぼない内容
- 外部サービスや実ネットワークが必須の内容
- ランダム性や実時刻に依存し、安定しないテスト
- Gin、GORM など特定フレームワークだけの使い方
- テーマ数を増やすためだけの細分化

## Go 版で優先する領域

### 優先度 A

既存テーマとの重複を確認し、未整備なら優先する。

- nil interface / typed nil
- slice の共有
- append と capacity / 再確保
- copy と backing array
- array と slice
- map の zero value / lookup / delete
- pointer receiver / value receiver
- method set
- defer の実行順序と評価タイミング
- panic / recover
- errors.Is / errors.As / wrapping
- goroutine lifecycle
- buffered / unbuffered channel
- channel close
- select
- context cancellation / deadline

### 優先度 B

- range の値コピー
- closure
- zero value
- struct embedding
- interface satisfaction
- type assertion / type switch
- generics
- comparable
- `sync.Once`
- `sync.Mutex` / `RWMutex`
- `WaitGroup`
- atomic
- `time.Time` comparison
- JSON `omitempty`
- JSON unknown fields
- `http.Response.Body` close

### 優先度 C

最小かつ決定的なテストで示せる場合に追加する。

- memory model
- race を起こさない synchronization の契約
- escape analysis に関連する観測可能な挙動
- scheduler に依存しない goroutine 設計
- generics の高度な constraints

## Go らしさを優先する

Java / TypeScript 版とのテーマ数合わせはしない。

Go では特に次を重視する。

- zero value
- value semantics
- pointer semantics
- slice / map
- interface
- explicit error handling
- defer
- goroutine / channel
- context
- standard library
- 小さな abstraction

他言語に対応テーマがあっても、Go 固有の学習価値が弱ければ無理に追加しない。

## 並行処理テーマの重要ルール

並行処理教材は flaky test にしない。

禁止例:

- `time.Sleep` の長さだけで goroutine の完了順を保証する
- scheduler の実行順を期待値にする
- race condition の偶然の再現を成功条件にする

代わりに channel、WaitGroup、context などを使い、同期点を明示する。

## 各教材の作り方

既存の Source + Test 形式を維持する。

1. 一つの挙動に焦点を絞る
2. Source と `*_test.go` を最小にする
3. 値、error、panic、順序、状態変化を観測する
4. `src/data/lessons.ts` に観測結果を記録する
5. なぜその挙動になるか短く説明する
6. 実務での注意点があれば記載する
7. `gofmt` を適用する
8. テストを決定的にする

## 分類

新規テーマは可能なら以下の観点を明示する。

- `Pitfall`: 誤解しやすさが中心
- `Behavior`: 実行時契約が中心
- `Concept`: Go の概念理解が中心

UI に分類機能を追加すること自体は必須ではない。分類管理のためにサイトを複雑化しない。

## テーマ追加時の判断質問

- Go 経験者でも結果を説明しにくいか
- 実務のバグ、レビュー、設計判断に役立つか
- `go test` で観測する意味があるか
- 既存テーマと重複していないか
- Go 固有の設計思想や実行モデルを理解できるか
- flaky にならず再現可能か

3 個以上が弱い場合は追加しない。

## 実施手順

1. README、LEARNING_PATH、`src/data/lessons.ts`、既存 `examples/` を確認する
2. 既存テーマと重複しない候補を洗い出す
3. Go 固有性と実務価値で優先順位を付ける
4. 一度に大量追加せず、意味のまとまり単位で実装する
5. Source / Test / Metadata / 説明を同期する
6. `gofmt` と `go test ./...` を実行する
7. `pnpm check`、静的ビルド、リンク検査を通す
8. テーマ数が変わる場合は README / LEARNING_PATH も更新する

## 完了条件

- 標準 `testing` で挙動が固定されている
- テストが決定的で flaky ではない
- Source と Test が最小
- 表示内容と実行コードが一致
- 既存テーマと重複しない
- Go 固有の学習価値がある
- 単なる文法チュートリアルではない
- `go test ./...` とサイト側品質ゲートが成功する

テーマ数ではなく、**Go の挙動を再現可能なテストとして残せているか**を品質基準にする。
