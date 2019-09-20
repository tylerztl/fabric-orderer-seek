### Prerequisites
- Go 1.10+ installation or later
- GOPATH environment variable is set correctly
- Protoc Plugin
- Protocol Buffers

### Getting started
#### 1. Build images
```
docker build -t hyperledger/fabric-orderer-seek .
```
#### 2. Start mysql
```
cd yaml
docker-compose -f docker-compose-mysql.yaml up -d
```
#### 3. Start explorer
```
docker-compose -f docker-compose-explorer.yaml up -d
```
