# go-kvstore

Key-Value Store written in Go

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

### Set a new key-value in the store

```bash
http POST "http://localhost:8080/key" key=value
```

### Get an existing key-value set from the store

```bash
http GET "http://localhost:8080/key"
```

### Update an existing value for a given in the store

```bash
http PUT "http://localhost:8080/key" value
```

### Delete an existing key-value set from the store

```bash
http DELETE "http://localhost:8080/key"
```