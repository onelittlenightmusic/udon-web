# Simple web CRUD API "udon-web"

This repo shows an example implementation of web service. The original instruction is [Yuri's article "Golang. CRUD in REST API in a generic way."](https://fenyuk.medium.com/golang-crud-in-rest-api-in-a-generic-way-9c395a60309e).

## Usage

The golang based web service can be run as follows,

```sh
git clone https://github.com/onelittlenightmusic/udon-web 
cd udon-web
go run .
```

You can access the web server at `http://localhost:6000`.

## Spec

The web service has just a simple API set.

```yaml
  /api/udons:
    get: # Get all udon instances
    post: # Create an udon instance
  /api/udons/{udon_id}:
    get: # Get an udon instance
    delete: # Delete an udon instance
    put: # Update an udon instance
```

## Configure

The port can be changed in `config.json`.

```json
{
  "port": 6000
}
```