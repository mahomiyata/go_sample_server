# Go Sample Server for Line Bot

Line Botで送信したメッセージをデータベースに保存します。

## 準備

Docker Composeを利用してPostgresとサーバーを準備します。

```bash
$ docker-compose build
$ docker-compose up
```

## APIの説明

### GET /notes

全ユーザーのノートを古い方から5件取得します。

### GET /notes/:id/:start

指定のIDを持つユーザーのノートを5件ずつ表示します。

```bash
(EXAMPLE)

GET /notes/123123/1
→ user_idが123123のユーザーのノートを新しい方から5件表示

GET /notes/123123/2
→ user_idが123123のユーザーのノートで、新しい方から6件〜10件目を表示
```

### POST /notes

Line Botから送信されたメッセージを新規のノートとして保存します。