# Go Admin
![Build](https://github.com/quangdangfit/go-admin/workflows/.github/workflows/go.yml/badge.svg)

RBAC scaffolding based on Gin + Gorm + Casbin + Dig


## How to run

### Required

- Postgres
- Redis

### Config
Simply run `make startup`, or run following commands step - by - step:
- Copy config file: `cp config/config.sample.yaml config/config.yaml`
  

- You should modify `config/config.yaml`

```yaml
database:
  host: localhost
  port: 5432
  name: github.com/quangdangfit/go-admin
  env: development
  user: postgres
  password: 1234
  sslmode: disable

redis:
  enable: true
  host: localhost
  port: 6397
  password:
  database: 0

cache:
  enable: true
  expiry_time: 3600
```

### Run
```shell script
$ go run main.go 
```

### Document
* API document at: `http://localhost:8888/swagger/index.html`

### Techstack
- RESTful API
- Gorm
- Swagger
- Logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- Cron Job
- Redis
- Dig (Dependency Injection)