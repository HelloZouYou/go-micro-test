# Inventory Service

This is the Inventory service

Generated with

```
micro new payment-web --namespace=mu.micro.book --alias=inventory --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.web.inventory
- Type: web
- Alias: inventory

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
./inventory-web
```

Build a docker image
```
make docker
```