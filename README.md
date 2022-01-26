# spbill-api

## Language, Framework, etc ...

- 使用言語：Go
- フレームワーク：Echo
- アーキテクチャ：Clean Architecture

## Database
- mysql

# Getting started

```bash
$ cd ~/spbill-api
$ cp .env.sample .env
$ docker-compose up -d
```

# ルーティング

```bash
// sign up
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"name": "sample", "email": "sample@test.com", password: "password"}' localhost:8000/signup

// sign in
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"email": "sample@test.com", password: "password"}' localhost:8000/login

// 一覧
$ curl -i -H 'Content-Type:application/json' localhost:8000/users

// ID指定
$ curl -i -H 'Content-Type:application/json' localhost:8000/users/3

// 登録
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"column1": "XX", "column2": "XX", ...}' localhost:8000/users

// 更新
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d '{"ID": 6,"column1": "YY", "column2": "YY", ...}' localhost:8000/users/6

// 削除
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X DELETE localhost:8000/users/6
```
