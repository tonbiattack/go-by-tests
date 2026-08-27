# Go by Tests — デザイン検討

## アプローチ候補

### 1. パケット追跡の作業台
**Very Brief Intro**: Goの値、goroutine、channelが移動する経路を観測する、ネットワーク診断コンソールのような静かな作業面。挙動を「流れ」として確かめる感覚をつくる。

**Probability**: 0.06

### 2. Go Docの余白注釈
**Very Brief Intro**: 標準ライブラリの文書を読み解くための、明るく簡潔な技術ノート風の方向性。短い証拠と型シグネチャの視認性を最優先にする。

**Probability**: 0.03

### 3. 並行処理の信号地図
**Very Brief Intro**: channel、context、errorの状態遷移を地図記号のように配置する、構造的でやや図解的な学習環境。複数の実行経路を見失わないことを狙う。

**Probability**: 0.08

---

## 選定：パケット追跡の作業台

### Design Movement
**Swiss International Typographic Style** の情報階層と、GoのCLI・ネットワーク診断画面の密度を結び付ける。Java版と同じくSource/Testを同格に置きながら、値・rune・error・goroutineの流れを扱うGoらしい観測面にする。

### Core Principles
1. **Source と Test をひとつの観測面に置く。** すべてのテーマは、実装と`testing`の証拠を並べて始める。
2. **ゼロ値・nil・errorを曖昧にしない。** 既定値と成功・失敗の境界を、短い結果表示で固定する。
3. **データの流れを示す。** 値やsignalの向きを細いトレースで示し、並行処理の例でも視線を迷わせない。
4. **補足は観測ノートに留める。** 長い解説を避け、テストへ戻るための問いと結果だけを置く。

### Color Philosophy
背景は深い青墨の**runtime ink**にして、ターミナルとエディタの集中感を保つ。Goらしい鮮やかな**Trace Cyan**をデータ経路・型情報・現在地に、暖色の**Error Coral**を例外的な失敗やnil境界に限定する。Cyanは装飾ではなく「観測できた流れ」を示すために使う。

### Layout Paradigm
固定のパッケージレールと、可変の**実行トレース面**で構成する。デスクトップは左にカテゴリ・テーマ、右にメタ情報、50:50のSource/Test、観測ノートを順に流す。モバイルではレールを折り畳み、Source → Test → Evidenceの単一列へ再構成する。

### Signature Elements
1. **Trace Junction**: SourceとTestを結ぶCyanの細線と、値が通過したことを示す小さな接続ノード。
2. **型・値ラベル**: `[]byte`、`rune`、`nil`、`error`などを小型の等幅ラベルで示す。
3. **Evidence Ledger**: 結果を`input`、`observed`、`contract`として並べる短い観測台帳。

### Interaction Philosophy
操作はGoの標準ツールのように即時かつ明確に応答する。コピー、パッケージ移動、検索、前後の記事移動は、色とラベルの両方で状態を伝える。詳細ページでは←/→で前後の観測へ移動でき、入力中や検索中は無効化する。

### Animation
初期表示ではメタ情報、Source/Test、Evidence Ledgerを50ms間隔で短くフェード・スライドインさせる。Trace Junctionは遷移せず常に即時に見え、hoverでは境界線だけを160msで強める。すべての非必須アニメーションは`prefers-reduced-motion`で停止する。

### Typography System
見出しは**Space Grotesk**、日本語本文は**Noto Sans JP**、コードと観測ラベルは**IBM Plex Mono**を使う。見出しは太く低い字間、本文は余裕のある行間、コードは13px以上を保つ。

### Brand Essence
**Go by Testsは、Goのゼロ値・並行処理・標準ライブラリの挙動を、実行可能なテストで観測したい開発者のためのリファレンスである。**

性格: **観測的、簡潔、信頼できる**

### Brand Voice
観測できる結果を最初に示し、曖昧な一般論で終わらせない。見出し・CTA・マイクロコピーは、値や実行経路を確かめる行為を促す。

例:

> ゼロ値だけでは、存在を証明できない。

> この受信結果が、channelの状態を示す。

### Wordmark & Logo
ロゴマークは、二本の角括弧を思わせる縦線の間をCyanの一本線が斜めに通過する記号にする。入力から観測結果へ向かう流れを示し、ワードマークは`GO`を細く、`/TESTS`を太く組む。

### Signature Brand Color
**Trace Cyan — #00ADD8**。観測できたデータ経路とGo by Testsの固有の状態を示す鮮やかな青。
