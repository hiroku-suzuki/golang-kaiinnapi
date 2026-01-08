# 会員API モックサーバ（Kaiin API）README

このリポジトリは、会員情報を返す **モックAPI** を Gin で提供します。

---

## 提供API

### GET /api/Kaiin/:kaiinNo

- **目的**：会員情報を返す（モック）
- **対応会員番号**：`00010001` のみ

**リクエスト例**

```bash
curl -i http://192.168.91.132:8080/api/Kaiin/00010001
```

**存在しない会員番号（404）**

```bash
curl -i http://192.168.91.132:8080/api/Kaiin/99999999
```

---

## レスポンス

### 成功（200）

- `model` に会員情報を返します
- `errors` / `warnings` は通常省略（空の場合は出しません）

### Not Found（404）

- `model: null`
- `errors` に `kaiinNo not found` を返します

### ヘッダ

モックとして必要最小限の no-cache 系のみ付与します。

- `Cache-Control: no-cache`
- `Pragma: no-cache`
- `Expires: -1`

`Content-Type / Date / Content-Length` などは Gin が通常自動付与するため、固定しません。

---

## 実装方針（レイヤ構成）

- **Controller層**：HTTP入出力（パラメータ取得・ヘッダ付与・JSON返却）
- **Service層**：取得処理（将来の加工・バリデーションの置き場）
- **Repository層**：データ取得（今回は in-memory `map`）

---

## ディレクトリ構成（主要ファイル）

```
.
├─ main.go
├─ controllers/
│  └─ kaiin_controller.go
├─ services/
│  └─ kaiin_service.go
├─ repositories/
│  └─ kaiin_repository.go
├─ models/
│  └─ kaiin.go
└─ mockdata/
   └─ kaiin_sample.go   # サンプル（追跡対象）
```

---

## モックデータの扱い（mockdata/ と gitignore）

### やりたいこと

- **モックデータは `mockdata/` 配下に置く**
- **ただし実データ（ローカル用）は Git 管理しない**

この方針にするため、以下の運用にします。

- `mockdata/kaiin_sample.go` は **追跡（commit）** する（リポジトリだけでビルドできるようにする）
- ローカルで置くモックは `mockdata/kaiin.local.go` のように作り、**gitignore** する

### .gitignore 例

```gitignore
# mockdata のローカル差分はコミットしない
mockdata/*.local.go
mockdata/*.json
mockdata/*.csv
```

> ※ `mockdata/` 全体を ignore してしまうと、クリーン環境でビルドできなくなる可能性があるため、  
> 「サンプルだけ追跡＋ローカル差分だけ ignore」を推奨します。

### モックデータの差し替え

`repositories.NewInMemoryKaiinRepository(data)` に注入する `data` を差し替えるだけです。

- 例：`main.go` 側で `mockdata.KaiinData`（sample/local）を渡す

---

## 起動方法

### 1) 依存取得

```bash
go mod tidy
```

### 2) 起動

```bash
go run .
```

デフォルトで `:8080` で起動します。

---

## ルーティングとDI（概要）

`/api` グループ配下に `GET /Kaiin/:kaiinNo` を登録します。

```go
api := r.Group("/api")
{
    api.GET("/Kaiin/:kaiinNo", kaiinCtrl.GetKaiin)
}
```

Repository には **データを注入**しており、Repository 層は `mockdata` を直接 import しません。

---

## 実装メモ

- 対応会員番号は `00010001` のみ（mapキー一致で返却）
- API追加が増えてきたら `App構造体` や `RegisterRoutes()` に DI/ルーティングを集約するのがおすすめ
