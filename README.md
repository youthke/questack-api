# questack-api

## 起動方法

### mysqlの起動

dbディレクトリにて

```cmd
docker-compose up -d
```

### serverの起動

プロジェクトのルートディレクトリにて

```cmd
go run cmd/main.go
```

## resources 

### Question

|name|type|description|
|:---|:---|:---|
|id|int|primary key|
|title|string||
|author|string|questioner's name|
|content|string|question body|

### Owner

|name|type|description|
|:---|:---|:---|
|id|int|primary key|
|name|string|user's name|
|mail|string|mail address|
|description|string|presentation|
|stacks|[]Stack|question boxes|

### Stack

|name|type|description|
|:---:|:---:|:---|
|id|string|primary key|
|owner|Owner||
|name|string||
|questions|[]Question||

### 注意点
もしこれをサーバーにあげる場合jwtのkeyを複雑で誰にもわからない内緒な文字列に変更してください
