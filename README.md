# Marketplace Engine

REST APIs for Web3 Auth and Smart Contract Functionalities

[![.github/workflows/test.yml](https://github.com/TheLazarusNetwork/marketplace-engine/actions/workflows/test.yml/badge.svg)](https://github.com/TheLazarusNetwork/marketplace-engine/actions/workflows/test.yml)
[![Lint](https://github.com/TheLazarusNetwork/marketplace-engine/actions/workflows/lint.yml/badge.svg)](https://github.com/TheLazarusNetwork/marketplace-engine/actions/workflows/lint.yml)

# Getting Started

## Postgres for development

```bash
docker run --name="marketplace" --rm -d -p 5432:5432 \
-e POSTGRES_PASSWORD=lazarus \
-e POSTGRES_USER=lazarus \
-e POSTGRES_DB=lazarus \
postgres -c log_statement=all
```

## Steps to get started

- Run `go get ./...` to install dependencies
- Set up env variables or create `.env` file as per [`.env-sample`](https://github.com/TheLazarusNetwork/marketplace-engine/blob/main/.env-sample) file
- Run `go test ./...` to make sure setup is working
- Run `go run main.go` to start server

## Env Reference

| Key                         | Description                                                                                                                    |
| :-------------------------- | :----------------------------------------------------------------------------------------------------------------------------- |
| `JWT_PRIVATE_KEY`           | Private key used for signing the issuing JWT tokens                                                                            |
| `JWT_EXPIRATION_IN_HOURS`   | Hours after which issued JWT token should be invalid                                                                           |
| `AUTH_EULA`                 | EULA to be signed when requesting JWT token                                                                                    |
| `CREATOR_EULA`              | EULA to be signed when requesting Creator Role                                                                                 |
| `APP_NAME`                  | App name to be logged with logger                                                                                              |
| `GIN_MODE`                  | Gin mode used to specify type of logging for API requests, use `release` for production and `debug` for testing and deployment |
| `CREATIFY_CONTRACT_ADDRESS` | Contract address of creatify deployed on blockchain                                                                            |
| `POLYGON_RPC`               | RPC url of node which is connected to network where your smart contract is deployed                                            |
| `MNEMONIC`                  | Mnemonic of wallet which has operator role in smart contract                                                                   |
| `LOG_TO_FILE`               | Specify wheather to log to file unders logs folder                                                                             |
| `IPFS_NODE_URL`             | Node url of ipfs network                                                                                                       |
| `ALLOWED_ORIGIN`            | Origin which are allowed to access this APIs                                                                                   |

Note - Database env variables are self explainatory

## FAQs

### How to get polygon RPC url?

> You can get polygon RPC url from many providers, like alchemy.
> Also you can host your own node and use that.

### How to get ipfs node url?

> You can get ipfs node url from many providers like infura and pinata cloud
> Also you can host your own node and use that.

### What is MNEMONIC

> A MNEMONIC is set of words which is used to derive many private keys which can be used to access many wallets associated with private key. In this backend the first wallet is used as operator.

## API Reference

### Auth

For protected APIs use bearer token which can be obtained after calling authenticate API.

Use `Authorization` key in header in order to send token,
also since it is bearer token append `Bearer ` as suffix
for example, `Bearer mytokenabcd`

### APIs

#### Returns flow ID and EULA which should be signed and send to authenticate API in order to get the JWT which can be used for accessing all other APIs

```
  GET /flowid?walletAddress={{wallet address}}
```

#### Request JWT for the user who accepted the EULA by signing it with flow ID, i.e. flowid+EULA.

```
  POST /authenticate
```

| Parameter   | Type     | Description                                             |
| :---------- | :------- | :------------------------------------------------------ |
| `flowId`    | `string` | **Required**. flowId you got from flowId API            |
| `signature` | `string` | **Required**. signature obtained by signing flowId+EULA |

#### Get profile details of user present in the system.

Note - Some unset data is emitted.

```
  GET /profile
```

#### Updates profile data like name, profile picture URL, etc.

```
  PATCH /profile
```

| Parameter           | Type     |
| :------------------ | :------- |
| `name`              | `string` |
| `country`           | `string` |
| `profilePictureUrl` | `string` |

#### Returns flow ID and Eula which should be signed and passed to claim Role in order to successful verification and claim of role

```
  GET /roleId/{{roleId}}
```

#### Successfully complete role claim by sending signature which is obtained from signing eula+flowId which was returned from roleId API (aka Request role)

```
  POST /claimrole
```

| Parameter   | Type                   |
| :---------- | :--------------------- |
| `flowId`    | **Required**. `string` |
| `signature` | **Required**. `string` |

#### Delegate artifact creation to other creator

```
  POST /delegateArtifactCreation
```

| Parameter        | Type                   |
| :--------------- | :--------------------- |
| `creatorAddress` | **Required**. `string` |
| `metaDataHash`   | **Required**. `string` |

```
  POST /uploadtoipfs
```

| Parameter | Type                  |
| :-------- | :-------------------- |
| `file`    | **Required**. `Files` |
