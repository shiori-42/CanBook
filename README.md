# CanBook

![Green and Pink Bright Playful Simple Business Project Presentation](https://github.com/shiori-42/CanBook/assets/147027038/67a377c2-61a7-491a-87e5-d6ce08d3712c.png)
![27](https://github.com/shiori-42/CanBook/assets/147027038/c544fcfe-f9e8-47b4-89f4-81f67de40e07.png)

## 製品概要

大学生のための、キャンパス内で教科書を売買・レンタルできる特化型フリマアプリです。

[作品スライド](https://www.canva.com/design/DAGAO-n2Ewo/s37XOtN8EWCaMm2fmXKKUw/edit?utm_content=DAGAO-n2Ewo&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)

## 使用方法（現状、ローカルでのみ動作）

### ローカルにクローンする：

```bash
$ git clone https://github.com/shiori-42/CanBook.git
```

### .envファイルの作成(バックエンド)

`go/backend`ディレクトリに`.env`ファイルを作成し、以下の情報を設定してください。デプロイが完了していないため、ローカルでDockerを使い、PostgreSQLのインスタンスを立ち上げます。

```bash
PORT=8080
POSTGRES_USER=myuser
POSTGRES_PASSWORD=mypassword
POSTGRES_DB=mydatabase
POSTGRES_HOST=localhost
SECRET=uu5pveql
GO_ENV=dev
API_DOMAIN=localhost
FE_URL=http://localhost:3000
```

### .env.localファイルの作成(フロントエンド)

`frontend`ディレクトリに`.env.local`ファイルを作成し、次のように設定してください。

```bash
NEXT_PUBLIC_API_URL=http://localhost:8080
```

この変数は、フロントエンドがバックエンドAPIにアクセスするためのエンドポイントURLを指定しています。デプロイが完了していないため、ローカル環境を指定します。

### 画像を保存する`images`ディレクトリの作成(バックエンド)

`go`ディレクトリ配下に`images`ディレクトリを作成してください。これは出品時にアップロードする画像の保存場所として使用されます。

### Dockerを使用してPostgreSQLのインスタンスを立ち上げる

`go/backend`ディレクトリ内で、以下のコマンドを実行してPostgreSQLを起動します。

```bash
$ docker-compose up -d
```

### サーバーを起動する

**バックエンド：**

```bash
# go/backendディレクトリ内で
$ GO_ENV=dev go run main.go
```

**フロントエンド：**

```bash
# frontend/appディレクトリ内で
$ npm run dev
```

## 製品開発のきっかけ、課題

国内最大級の女性＆ノンバイナリーの方向けハッカソン「Dots to Code」で開発を開始しました。教科書や技術書の出費が負担となっている私たち開発者の経験から着想を得て、この課題を解決するためにプロダクトを開発しました。

## 製品説明

現在、ユーザーの新規登録・ログイン、商品の出品機能、商品検索機能が実装されています。新規ユーザー登録や商品出品が可能で、キーワードを指定して目的の商品を検索することができます。

## 解決できること

キャンパス内で直接受け渡しするため、送料がかかりません。出品者は梱包・発送の手間がなくなり手軽に出品でき、購入者は教科書代の負担を軽減できます。

## 今後の展望

- まずはプレゼン資料に沿った機能の完成
  - チャット機能の完成
  - 検索機能の強化
- 次にリファクタリング
  - CSRFトークンを組み合わせたセキュリティ強化
- さらに拡張
  - Google Maps APIを利用して、近隣のキャンパスの商品を探すことができるようにする

## 開発技術

- **フロントエンド**
  - TypeScript、Next.js
- **バックエンド**
  - Go、Echo、GorillaのWebSocketパッケージ
- **データベース**
  - PostgreSQL

## 対応デバイス

- パソコン（デプロイ後はスマホにも対応）
