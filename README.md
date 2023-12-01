# iotcore-pubsub

## 手順

### トピック作成
```bash
$ gcloud pubsub topics create sensors
```

### デプロイ
```bash
make deploy-pubsub
```

デプロイコマンドにより以下の処理が実行される（と思われる）
1. Cloud Build でビルド
2. Artifact Registry にイメージをプッシュ
3. Cloud Run にデプロイ
- Pub/Sub トリガーを有効にする

### メッセージ送信
```bash
gcloud pubsub topics publish sensors --message="{\"temperature\": 25.11, \"humidity\": 80.29}"
```