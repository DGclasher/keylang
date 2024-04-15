# KeyLang
Generate license keys in the GO

### Description
KeyLang is a tool to manage software licenses. It helps you create unique license keys for your products and validate them easily. With KeyLang, you can keep track of who's using your software and ensure they have valid licenses, making license management a breeze.

### Deploy locally
```
wget https://github.com/DGclasher/keylang/raw/main/docker-compose.yml
```
```
wget https://github.com/DGclasher/keylang/raw/main/.env.example -O .env
```
```
mkdir storage
```
Start the API
```
docker compose up -d
```
Example usage
```
curl -X POST localhost:8080/api/license/create -d '{"user_email":"example@email.com","product_id":"12345ABC"}' -H 'content-type: application/json'
```

#### [API Docs](./docs/api.md)
