### Prerequisites
- Go 1.10+ installation or later
- GOPATH environment variable is set correctly
- Protoc Plugin
- Protocol Buffers

### Getting started
#### 1. Start mysql
```
cd yaml
docker-compose -f docker-compose-mysql.yaml up -d
```
#### 2. Start explorer
```
docker-compose -f docker-compose-explorer.yaml up -d
```
