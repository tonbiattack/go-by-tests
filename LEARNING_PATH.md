# Go by Tests 学習順ガイド

[Go by Tests の公開ページ](https://tonbiattack.github.io/go-by-tests/)には、Go 1.22の挙動を確認する34テーマがあります。このガイドでは、暗記ではなく **Source → Test → 観測結果** の順で、実際に挙動を確かめながら進める順序を示します。

> **基本の進め方**: テスト名を先に読み、結果を予想し、Sourceを一行だけ変えて失敗を観測してから、最小の変更で元のテストを通します。

## 最初の4テーマ

最初は、Goの文字列・slice・error・contextを1テーマずつ読むと、値の表現、共有状態、失敗、キャンセルという実務上の境界を短時間で確認できます。

| 順序 | テーマ | 確認する契約 |
|---:|---|---|
| 1 | [`string: len`](https://tonbiattack.github.io/go-by-tests/go/string/utf8-length/) | `len` は文字数ではなくUTF-8のバイト数を返す。 |
| 2 | [`slice: append`](https://tonbiattack.github.io/go-by-tests/go/slice/append-backing-array/) | 容量に余裕がある `append` は元のバッキング配列を共有し得る。 |
| 3 | [`error: wrapping`](https://tonbiattack.github.io/go-by-tests/go/errors/wrap-and-is/) | `%w` でラップした原因は `errors.Is` で照合する。 |
| 4 | [`context: cancel`](https://tonbiattack.github.io/go-by-tests/go/context/cancel/) | キャンセル済みのcontextは `context.Canceled` として観測する。 |

## 関心別の入口

学ぶ目的に合わせて開始地点を選ぶ場合は、次のカテゴリを入口にしてください。

| 関心 | 推奨カテゴリ | 代表的な論点 |
|---|---|---|
| 基本の値と比較 | [Language](https://tonbiattack.github.io/go-by-tests/) | UTF-8、`rune`、typed nil、比較不能な値、`range`。 |
| 可変状態とコピー | [Collections](https://tonbiattack.github.io/go-by-tests/) | sliceの共有、`nil`と空slice、mapのゼロ値、`slices.Clone`。 |
| エラーを失わない設計 | [Error Handling](https://tonbiattack.github.io/go-by-tests/) | wrapping、`errors.Join`、`errors.As`、`defer`、`recover`。 |
| goroutineとキャンセル | [Concurrency](https://tonbiattack.github.io/go-by-tests/) | context、channel、`sync.Once`、`sync.Map`、Go 1.22のrange変数。 |
| 時刻と外部境界 | [Date / Time](https://tonbiattack.github.io/go-by-tests/)・[HTTP](https://tonbiattack.github.io/go-by-tests/) | `time.Time.Equal`、`time.Duration`、`Response.Body.Close`。 |
| JSON入出力 | [Encoding](https://tonbiattack.github.io/go-by-tests/) | `omitempty`、非公開フィールド、未知フィールド。 |

## 各テーマで試すこと

1. テスト名を、自分の言葉で「何を保証しているか」に読み替えます。
2. `t.Errorf`、`t.Fatalf`、panicの検証、エラー照合のどれが契約を固定しているかを確認します。
3. Sourceの一行を変更し、期待値・実際値・エラー・panicのいずれが変化するかを観測します。
4. テストを通す最小の変更を加えます。
5. 境界値、失敗時の状態、リソース解放、並行処理のいずれかを追加で一つ確認します。

ローカルでは、次のコマンドで全教材を検証できます。

```bash
go test ./...
```

新しいテーマを追加する場合は、Source、Test、メタデータ、観測結果を同じ変更に含めてください。詳細は [CONTRIBUTING.md](./CONTRIBUTING.md) を参照してください。

## References

[1] [The Go Programming Language — Testing](https://go.dev/pkg/testing/)
[2] [Go by Tests — Live Demo](https://tonbiattack.github.io/go-by-tests/)
