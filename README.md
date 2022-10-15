Pataka is an application that manages feature flag.

# Overview

Pataka is a back-end services that you can run, and it provides an API that you can call via HTTP.

Pataka provides:

- CRUD operation for your feature manifest.
- CRUD operation for your feature audiences.

# Ideas

- CLI to manage feature and audience.
- Golang client library.
- Audience group to easily target the feature's audience.

# Usage

To run pataka, you need to have a postgresql database. After you have a database, please setup the `.env` and run the database migration.

```sh
make migrate
```

Generate an access token to your Pataka API. For the Pataka client, please use read only token.

```sh
go run main.go create token --scopes write:feature,write:audience,read:feature,read:audience
go run main.go create token --scopes read:feature,read:audience
```

Then, run your server

```sh
go run main.go serve
```

# License

Pataka is released under the Apache 2.0 license. See [LICENSE](./LICENSE)
