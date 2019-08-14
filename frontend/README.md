# Getting Started

## 1. 環境構築

**make setup**

設定ファイルのコピーとパッケージのセットアップを行う

```console
❯ make setup
```

**Firebase API Key**

API講義で作ったFirebaseプロジェクトを使う

1. [Firebase](https://firebase.google.com/) のプロジェクトページを開く
1. `Project Overview` の右にある `⚙(歯車)` ボタンを押して、 `プロジェクトの設定` へ遷移
1. `Firebase SDK snippet` をコピーして、`.env` ファイルの `FIREBASE_...` に設定する

```
FIREBASE_APIKEY="REPLACE_ME"
FIREBASE_AUTHDOMAIN="REPLACE_ME"
FIREBASE_DATABASEURL="REPLACE_ME"
FIREBASE_PROJECTID="REPLACE_ME"
FIREBASE_MESSAGINGSENDERID="REPLACE_ME"
FIREBASE_APPID="REPLACE_ME"
```

## 2. Hello World

**フロント立ち上げ**

```console
❯ make dev
```
