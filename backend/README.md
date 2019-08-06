# Getting Started

## 1. Database 立ち上げ

`../database` に構成などが定義されているので、読んでみてください。

```console
❯ make -f integration.mk database-init 
make -C ../database init
which goose || go get -u github.com/pressly/goose/cmd/goose
/Users/j-chikamori/go/bin/goose
docker-compose up -d
Starting treasure-app-db ... done
mysql -u root -h localhost --protocol tcp -e "create database \`treasure_app\`" -p
Enter password: passwordとタイプする
goose -dir migrations mysql "root:password@tcp(127.0.0.1:3306)/treasure_app" up
2019/08/06 17:02:55 OK    1_init.sql
2019/08/06 17:02:55 goose: no migrations to run. current version: 1
```

**portが被っている時**  
`lsof -i :3306` で portを既に使ってるプロセス探して、 `kill` してしまおう

## 2. Firebase プロジェクト作成

1. [Firebase](https://firebase.google.com/) からプロジェクトを作成する
    1. Google アナリティクスはなしでおｋ
    
## 3. 設定ファイル準備(.env)

```console
❯ make init            
cp .env.example .env
```

**Firebase API Key**

1. [Firebase](https://firebase.google.com/) のプロジェクトページを開く
1. `Project Overview` の右にある `⚙(歯車)` ボタンを押して、プロジェクトの設定へ遷移
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

```console
❯ make -f integration.mk create-token UID=demo
go run ./cmd/customtoken/main.go demo
{
  "kind": "identitytoolkit#VerifyCustomTokenResponse",
  "idToken": "idtoken",
  "refreshToken": "refreshToken",
  "expiresIn": "3600",
  "isNewUser": true
}

token file created
```

## 5. Hello World

```console
❯ make run          
go run cmd/api/main.go
Listening on port 1991
```

サーバーを立ち上げた状態で、別シェルから以下を叩く。

```console
❯ make -f integration.mk req-private
curl -H "Authorization: Bearer tokenhogehoge" localhost:1991/private
{"message":" Hello  from private endpoint! You have 0 comments"}

❯ make -f integration.mk req-public 
curl localhost:1991/public
{"message":"Hello from a public endpoint! You don't need to be authenticated to see this."}
```

うまくいかない時は、サーバーを立ち上げたシェルに標準出力が何かが出ているはずなので、トラブルシューティングしてみよう！  
例えば、データベースが動いていない時は、以下のようなメッセージが出ている。

```console
2019/08/06 16:40:05 Error 1045: Access denied for user 'root'@'172.18.0.1' (using password: YES)
```
