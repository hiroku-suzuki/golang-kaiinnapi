# golang-mock
## auth-mock
### 実フォルダ構成

```text
auth-mock/ # 認証APIモックサーバ
  ├── main.go
  ├── apperrors/
  ├── controllers/
  ├── services/
  ├── repositories/
  └── models/

kaiin-mock/ # 会員APIモックサーバ
  ├── main.go
  ├── controllers/
  ├── services/
  ├── repositories/
  └── models/

```

### 各ディレクトリの役割

#### main.go

- 依存関係を組み立てるだけの層。
- `repositories` → `services` → `controllers` を `new` して HTTP サーバを起動する。

#### models/

- `User`, `Client`, `AuthenticateRequest`, `AuthenticateResult` など **データの形** を表す struct を置く。
- 基本的に「ビジネスルールは持たない入れ物」。

#### repositories/

- `UserRepository` / `ClientRepository` の **インターフェースと実装** を置く。
- まずはインメモリ版（`map`）だけを実装する。
- 将来 DB にする場合も、`repositories` の実装を差し替えるだけで `services` 層を変更せずに済む。

#### services/

- 認証のビジネスロジックを持つ層。  
- 主な責務：
  - RP 認証（`Authorization` ヘッダの Basic 認証情報を解析）
  - リクエストの業務的バリデーション
  - ユーザ認証（ID/PASS）
  - トークン生成（`access_token` / `refresh_token` / `id_token`）
- HTTP の都合（ステータスコードや JSON キー名）は知らない。

#### controllers/

- HTTP ハンドラ層。  
- 主な責務：
  - `http.Request` から JSON をデコードする
  - `Authorization` ヘッダを取り出す
  - `services.AuthService` を呼び出す
  - 戻ってきた `result` または `error` を、HTTP ステータス + JSON にマッピングしてレスポンスする

#### apperrors/

- アプリケーション共通のエラーコード・エラー型を定義する層。
- `services` 層はここで定義したエラーを返し、`controllers` 層がそれを見て HTTP ステータス（400 / 401 / 500 など）を決定する。
# golang-kaiinnapi
# golang-kaiinnapi
