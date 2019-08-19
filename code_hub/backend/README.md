# Slide

https://go-talks.appspot.com/github.com/voyagegroup/talks/2019/treasure-go-day2/intro.slide#1

# Getting Started

## 1. Database 立ち上げ

`../database` に構成などが定義されているので、読んでみてください。

以下のターゲットを叩くと、databaseの準備をします。

```console
❯ make -f integration.mk database-init
```

```console
which goose || GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose
/Users/j-chikamori/go/bin/goose
docker-compose up -d
ptreasure-app-db is up-to-date
mysql -u root -h localhost --protocol tcp -e "create database \`treasure_app\`" -p
Enter password:
goose -dir migrations mysql "root:password@tcp(127.0.0.1:3306)/treasure_app" up
2019/08/08 11:29:17 OK    1_init.sql
2019/08/08 11:29:18 OK    2_article.sql
2019/08/08 11:29:18 goose: no migrations to run. current version: 2
```

**エラー起きたら**  
- ポート被り  
`lsof -i :3306` で portを既に使ってるプロセス探して、 `kill` してしまおう
- `Lost connection to MySQL server`  
とかって出たら、まだmysql起動してないだけなので、もう一回やってみよう

## 2. Firebase プロジェクト作成

1. [Firebase](https://firebase.google.com/) からプロジェクトを作成する
    1. Google アナリティクスはなしでおｋ
    
## 3. 設定ファイル準備(.env)

**make init**

設定ファイルをコピーする。

```console
❯ make init            
cp .env.example .env
```

**Firebase API Key**

1. [Firebase](https://firebase.google.com/) のプロジェクトページを開く
1. `Project Overview` の右にある `⚙(歯車)` ボタンを押して、 `プロジェクトの設定` へ遷移
1. `ウェブ API キー` をコピーして、`.env` ファイルの `FIREBASE_WEB_API_KEY` に設定する

```
DATABASE_DATASOURCE=root:password@tcp(localhost:3306)/treasure_app?parseTime=true&charset=utf8mb4&interpolateParams=true
FIREBASE_WEB_API_KEY=コピーしたAPIキーをここに貼り付ける
#https://cloud.google.com/docs/authentication/production#finding_credentials_automatically
GOOGLE_APPLICATION_CREDENTIALS=
PORT=1991
```

**Google Application Credentials**

1. [Firebase](https://firebase.google.com/) のプロジェクトページを開く
1. `Project Overview` の右にある `⚙(歯車)` ボタンを押して、プロジェクトの設定へ遷移
1. タブの `サービスアカウント` へ遷移
1. `新しい秘密鍵の生成` ボタンを押して、鍵をダウンロードする
1. ダウンロードした鍵を、 `$HOME/.config/gcloud/application_default_credentials.json` へリネームして置く (もしくは、`.env` の `GOOGLE_APPLICATION_CREDENTIALS` に設定ファイルの場所を教える)

```console
❯ mv /Users/j-chikamori/Downloads/treasure-app-demo-pei-firebase-adminsdk-92eds-2cb2584124.json $HOME/.config/gcloud/application_default_credentials.json
```

## 4. Firebaseユーザーの作成

使っているAPI: [Exchange custom token for an ID and refresh token](https://firebase.google.com/docs/reference/rest/auth/#section-refresh-token)  
それぞれのキーについては、上記リンクを見れば大体分かる。

トークンの作成

```console
❯ make -f integration.mk create-token UID=demo
go run ./cmd/customtoken/main.go demo .idToken
```

```console
{
  "kind": "identitytoolkit#VerifyCustomTokenResponse",
  "idToken": "idtoken",
  "refreshToken": "refreshToken",
  "expiresIn": "3600",
  "isNewUser": true
}

.IdToken created
```

## 5. Hello World

*サーバー立ち上げ*

```console
❯ make run
go run cmd/api/main.go
```

```console
2019/08/08 11:32:47 server.go:51: Listening on port 1991
```

サーバーを立ち上げた状態で、別シェルから以下を叩く。

*認証ありエンドポイントへのリクエスト*

```console
❯ make -f integration.mk req-private
curl -v -H "Authorization: Bearer Hoge" localhost:1991/private
```

```console
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 1991 (#0)
> GET /private HTTP/1.1
> Host: localhost:1991
> User-Agent: curl/7.54.0
> Accept: */*
> Authorization: Bearer Hoge 
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Vary: Origin
< Date: Thu, 08 Aug 2019 02:33:22 GMT
< Content-Length: 70
<
* Connection #0 to host localhost left intact
{"message":"Hello  from private endpoint! Your firebase uuid is demo"}
```

*認証なしエンドポイントへのリクエスト*

```console
❯ make -f integration.mk req-public
curl -v localhost:1991/public
```

```console
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 1991 (#0)
> GET /public HTTP/1.1
> Host: localhost:1991
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Vary: Origin
< Date: Thu, 08 Aug 2019 02:33:32 GMT
< Content-Length: 91
<
* Connection #0 to host localhost left intact
{"message":"Hello from a public endpoint! You don't need to be authenticated to see this."}
```

うまくいかない時は、サーバーを立ち上げたシェルに標準出力が何かが出ているはずなので、トラブルシューティングしてみよう！  
例えば、データベースが動いていない時は、以下のようなメッセージが出ている。

```console
2019/08/06 16:40:05 Error 1045: Access denied for user 'root'@'172.18.0.1' (using password: YES)
```

## 6. いざ開発

コードを書き換える度に `go run` し直すのは面倒くさいので、書き換えるとrebuildしてくれるコマンドを用意しているので活用してください。

realizeをgo getする

```console
❯ make dev-deps                     
  GO111MODULE=off go get -u -v \
                  github.com/oxequa/realize
```

```
  github.com/oxequa/realize (download)
  github.com/oxequa/interact (download)
  github.com/fatih/color (download)
  github.com/fsnotify/fsnotify (download)
```

realizeを使って実行する

```console
❯ make refresh-run
realize start
```

``` console
[11:36:48][BACKEND] : Watching 21 file/s 14 folder/s
[11:36:48][BACKEND] : Install started
[11:36:51][BACKEND] : Install completed in 3.653 s
[11:36:51][BACKEND] : Running..
[11:36:51][BACKEND] : 2019/08/08 11:36:51 server.go:51: Listening on port 1991
```
