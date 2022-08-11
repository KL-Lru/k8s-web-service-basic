## Sample Web Service

Kubernetes実演用のサンプルウェブアプリケーションです.

## ディレクトリ構成

```
.
├── cmd
│   ├── batch_worker
│   ├── heartbeat_check
│   ├── stream_worker
│   └── web_server
├── pkg
├── docker
├── k8s
│   ├── connected
│   ├── stable
│   └── unstable
├── Makefile
└── terraform
```

### cmd
golangのbinaryエントリーポイントが入っています.
当ディレクトリ内に含まれるのは次の処理群です.

<details>
<summary>web_server</summary>
HTTPリクエストを受け付けるWebサーバです.
簡素化のため, Routingは下記のみ含まれます.

```bash
GET  /        Hello
GET  /healthz ヘルスチェック
GET  /users   SQLデータ参照
POST /tasks   MQへのメッセージ発行
```

#### Requirement
`GET /users`へのアクセスに対して正常応答するには, 次のusersテーブルを有する, DBへの正確な環境変数値が必要です.
```sql
CREATE TABLE users (id VARCHAR(50), name VARCHAR(255));
```
```.env
POSTGRES_USER=<アクセスユーザ>
POSTGRES_PASSWORD=<アクセスユーザのパスワード>
POSTGRES_DB_HOST=<SQLの稼働しているhost名 or IPアドレス>
POSTGRES_DATABASE=<データベース名>
```

`POST /tasks`へのアクセスに対して正常応答するには, 次の環境変数及びGCPサービスへのアクセス用のcredentialが必要です.
```.env
PROJECT_ID=<利用するGCP Project ID>
GOOGLE_APPLICATION_CREDENTIALS=<credentialへのパス>
```
</details>

<details>
<summary>batch_worker</summary>
実行するとSQLアクセスを実行し, 処理結果を表示して即時終了するタスク記述です.

#### Requirement
タスクの実行を正常終了させるには, 次のusersテーブルを有する, DBへの正確な環境変数値が必要です.
```sql
CREATE TABLE users (id VARCHAR(50), name VARCHAR(255));
```
```.env
POSTGRES_USER=<アクセスユーザ>
POSTGRES_PASSWORD=<アクセスユーザのパスワード>
POSTGRES_DB_HOST=<SQLの稼働しているhost名 or IPアドレス>
POSTGRES_DATABASE=<データベース名>
```
</details>

<details>
<summary>stream_worker</summary>
裏稼働のWorkerサーバです.
実行するとMQへのStreaming接続を実行, メッセージを常時Subscribeし, 受け取ったメッセージを表示します.

#### Requirement
正常稼働させるためには, 次の環境変数及びGCPサービスへのアクセス用のcredentialが必要です.

```.env
PROJECT_ID=<利用するGCP Project ID>
SUBSCRIBE_ID=<設定したPub/SubのSubscriber ID>
GOOGLE_APPLICATION_CREDENTIALS=<credentialへのパス>
```
</details>

<details>
<summary>heartbeat_check</summary>
特定パス上にあるファイル内部の情報が, 特定時刻以降になっているかを判定します.
Worker系のHealth Check処理のサンプルとして設置されています.
かなり作りが粗末なので, 少し変更を加える可能性があります.
</details>

### pkg

cmdに含まれるサービス群で利用するUtility

### docker

heartbeatを除く, 各cmdの起動に必要となる最小構成を取るためのDockerfileです.
batch-serverのみalpine(bourne shell利用のため), 他はscratchベースになります.

### k8s

Kubernetesのリソース定義が含まれています.
- unstable: 稼働はするものの定義が不十分な状態, 一部コンテナは確定エラーになる
- connected: 接続性定義まで完了している状態, 計算資源やヘルスチェック考慮なし
- stable: 接続性定義, 計算資源量, ヘルスチェックまで定義完了した状態

### terraform

Terraformの定義ファイルが含まれています.
十分な権限を持ったユーザにて, CLI上でログインが完了している場合において,
`terraform apply`で必要になるリソース群定義をDB Migration以外完了します.
この際一部リソースは利用料金が発生します.
ご注意ください.

### Makefile

コンテナイメージのbuild等の省略エイリアスが含まれています.
ローカル向けに当リポジトリのコンテナイメージをbuildする際は, `make`のみで, 次のコマンド相当の処理が実行されます.
```bash
docker build . -f docker/web_server.dockerfile -t test/samples/web_server:v1.0
docker build . -f docker/stream_worker.dockerfile -t test/samples/stream_worker:v1.0
docker build . -f docker/batch_worker.dockerfile -t test/samples/batch_worker:v1.0
```

他, クラウド上でコンテナイメージを参照出来るようにする場合は, 
次のコマンドで, 指定のGCP Project内の東京Region内のArtifact Registry内にコンテナイメージ群を入れ込むことが出来ます. (事前にDocker形式のリポジトリを作成しておく必要があります.)
```
make push REPO=asia-northeast1-docker.pkg.dev/<利用するProject名>
```
