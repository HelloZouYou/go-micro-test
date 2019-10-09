# Orders Service

This is the Orders service

Generated with

```
micro new orders-web --namespace=mu.micro.book --alias=orders --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.web.orders
- Type: web
- Alias: orders

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./orders-web
```

Build a docker image
```
make docker
```