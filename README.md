# Notifier Service

A notifier service to send out notifications on different channels(SMS, email, push). This service can also be used to run cron jobs if required.

- Prerequisites
  - go
  - make
- Building
  - `make build` - will build the project
- Running
  - env
    - REDIS_PORT=7002
    - REDIS_HOST=localhost
    - REDIS_USERNAME=default
  - `make run` - will run the project