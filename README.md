# acectl

A multi-interface command line tool for remote IBM ACE administration.

# Interfaces

## CLI

### Usage:

`acectl register --id EXAMPLEID --host 127.0.0.1 --port <optional, default 22> --username exampleusername --password examplepassword`

Registers an ACE broker to the persistent storage.

`acectl unregister EXAMPLEID`

Unregisters (deletes) the specified broker from persistent storage.

`acectl get EXAMPLEID`

Retrieves stored metadata for the specified broker.

`acectl list`

Lists all registered brokers.

`acectl status EXAMPLEID`

Retrieves status of the broker.

`acectl stop EXAMPLEID`

Stops specified broker.

`acectl start EXAMPLEID`

Starts specified broker.

## HTTP Server

### API

```json
POST /register [Registers a new broker]
{
    "id": "EXAMPLEID",
    "host": "127.0.0.1",
    "port": "",
    "username": "example",
    "password": "example"
}
```

```json
GET /broker [Lists all registered brokers]
```

```json
GET /broker/:brokerId [Fetches a specific broker's metadata]
```

```json
DELETE /broker/:brokerId [Unregisters a broker]
```

```json
GET /broker/:brokerId/status [Reports broker status]
```

```json
GET /broker/:brokerId/stop [Stops broker]
```

```json
GET /broker/:brokerId/start [Starts broker]
```