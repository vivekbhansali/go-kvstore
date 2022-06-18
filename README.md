# go-kvstore

Simple Key-Value Store written in Go. This is a learning project with the aim to understand Go Programming language and also different Database Storage techniques. 

The idea and the inspiration of this project comes from similar project by [gerlacdt](https://github.com/gerlacdt) called [db-key-value-store](https://github.com/gerlacdt/db-key-value-store).

## Build

To build the binary
```
make build
```

To build and run the app
```
make run
```

To remove the compiled binary of the app
```
make clean
```

## Usage

Open a terminal, and run the app.

Open another terminal and start interacting with the KV store. 

Here are some examples on how to use the KV store using [httpie](https://github.com/httpie/httpie) HTTP client.

### Set a new `key-value` in the store

```bash
http POST "http://localhost:8080/db/key" value
```

### Get an existing `value` for a given `key` from the store

```bash
http GET "http://localhost:8080/db/key"
```

### Update an existing `value` for a given `key` in the store

```bash
http PUT "http://localhost:8080/db/key" value
```

### Delete an existing `key-value` set from the store

```bash
http DELETE "http://localhost:8080/db/key"
```