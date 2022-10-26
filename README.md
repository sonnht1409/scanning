# scanning

## description

Build a simple code scanning application that detects sensitive keywords in public git repos. The application must fulfil the following requirements:

    - A user can CRUD repositories. A repository contains a name and a link to the repo
    - A user can trigger a scan against a repository
    - A user can view the Security Scan Result ("Result") List

How to do a scan:

    - Just keep it simple by iterating the words on the codebase to detect secrets findings
    - A secret starts with the prefix public_key or private_key

# solution

<img src="https://i.postimg.cc/C5JdnQs2/Screen-Shot-2022-10-26-at-17-35-02.png">

# setup

Prerequisite:
    - MySQL 8.0.0+
    - Redis latest
    - Go 1.19+
    - Docker

## Create Database

[Run this script](https://github.com/sonnht1409/scanning/blob/main/service/migration/scannings.sql)

## configuration

[Update configuartion file](https://github.com/sonnht1409/scanning/blob/main/service/config/config.yaml)

```yaml
db:
  address: {your-mysql-host}
  port: {your-mysql-port}
  name: scanning
  user: {your-mysql-username}
  password: {your-mysql-password}
redis:
  url: {your-redis-host}:{your-redis-port}
  db: 0
```

### notes

If using docker, pls set host of local redis address & db host into `host.docker.internal`

## To run rest api server
Using Docker
Run these commands in order
```sh
1. cd service
2. docker build -t scanning-server -f dockerfile.server .
3. docker run -p 8080:8080 scanning-server
```

Build binary image
Run these commands in order
```sh
1. cd service
2. make build
3. make server
```

## To run worker
Using Docker
Run these commands in order
```sh
1. cd service
2. docker build -t scanning-worker -f dockerfile.worker .
3. docker run scanning-worker
```

Build binary image
Run these commands in order
```sh
1. cd service
2. make build
3. make worker
```

# api collection
[API collection](https://www.getpostman.com/collections/07350ddbc1c5630472dc)

# features
- Signup a repository for scanning
    Request
    ```sh
        curl --location --request POST 'localhost:8080/api/v1/scanning' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "repo_name":"gin",
            "repo_url":"https://github.com/gin-gonic/gin"
        }'
    ```
    Response
    ```json
        {
            "message": "ok",
            "scan_unique_id": "77769549-899a-4347-903a-03c4e37dd0cb"
        }
    ```
- View scanning result for a repository
    Request
    ```sh
        curl --location --request GET 'localhost:8080/api/v1/scanning/results?repo_name=gin'
    ```
    Response
    ```json
    {
        "data": [
            {
                "id": 19,
                "repo_name": "gin",
                "repo_url": "https://github.com/gin-gonic/gin",
                "scan_unique_id": "77769549-899a-4347-903a-03c4e37dd0cb",
                "scanning_status": "SUCCESS",
                "created_at": "2022-10-26T17:23:39+07:00",
                "updated_at": "2022-10-26T17:24:49+07:00",
                "queued_at": "2022-10-26T17:23:39+07:00",
                "finished_at": "2022-10-26T17:24:49+07:00",
                "scanning_at": "2022-10-26T17:23:39+07:00"
            },
            {
                "id": 18,
                "repo_name": "gin",
                "repo_url": "https://github.com/gin-gonic/gin",
                "scan_unique_id": "e22af4c9-3c16-4e33-a235-03feec9f312a",
                "scanning_status": "SUCCESS",
                "created_at": "2022-10-26T17:18:27+07:00",
                "updated_at": "2022-10-26T17:18:46+07:00",
                "queued_at": "2022-10-26T17:18:27+07:00",
                "finished_at": "2022-10-26T17:18:46+07:00",
                "scanning_at": "2022-10-26T17:18:27+07:00"
            },
            {
                "id": 17,
                "repo_name": "gin",
                "repo_url": "https://github.com/gin-gonic/gin",
                "scan_unique_id": "a14e8c04-035c-4fa0-8a8b-cf93ac9b5ae3",
                "scanning_status": "SUCCESS",
                "created_at": "2022-10-26T17:16:34+07:00",
                "updated_at": "2022-10-26T17:16:40+07:00",
                "queued_at": "2022-10-26T17:16:34+07:00",
                "finished_at": "2022-10-26T17:16:40+07:00",
                "scanning_at": "2022-10-26T17:16:34+07:00"
            }
            ],
        "message": "ok"
    }
    ```
- View one scanning process
    Request
    ```sh
        curl --location --request GET 'localhost:8080/api/v1/scanning/result?scan_unique_id=6ae2b12a-9366-4a1e-9c97-212cd9d24e06'
    ```
    Response
    ```json
    {
        "data": {
            "id": 1,
            "repo_name": "go-crud",
            "repo_url": "https://github.com/sonnht1409/go-crud",
            "scan_unique_id": "6ae2b12a-9366-4a1e-9c97-212cd9d24e06",
            "scanning_status": "FAILURE",
            "created_at": "2022-10-25T16:54:25+07:00",
            "updated_at": "2022-10-25T18:14:19+07:00",
            "queued_at": "2022-10-25T16:54:25+07:00",
            "finished_at": "2022-10-25T18:14:19+07:00",
            "scanning_at": "2022-10-25T18:14:13+07:00",
            "findings": [
                {
                    "rule_name": "PublicKeyCheck",
                    "location": {
                        "path": "abc.go",
                        "line": 3
                    }
                },
                {
                    "rule_name": "PrivateKeyCheck",
                    "location": {
                        "path": "xyz.go",
                        "line": 3
                    }
                }
            ]
        },
        "message": "ok"
    }
    ```
- Retry one scanning process
     Request
    ```sh
        curl --location --request POST 'localhost:8080/api/v1/scanning/retry' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "scan_unique_id":"77769549-899a-4347-903a-03c4e37dd0cb"
        }'
    ```
    Response
    ```json
    {
        "message": "ok"
    }
    ```

# error

- Redis connection:
    ```sh
    docker run -d -p 6379:6379 redis:latest
    ```