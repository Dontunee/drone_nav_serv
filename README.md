# drone_nav_serv

The drone navigation service is an Api that enables drones to upload gathered data

## Prerequisites

Before you begin, ensure you have met the following requirements:

* [Go](https://golang.org/dl/)

## Installing

Clone this repo and run `make download` to install the dependencies.

## How to run locally:
run this command on the terminal `make run/api`

This will run on port 4000 by default 

**Endpoint Doc**

localhost:4000/v1/data
Request
```json
{
  "x": "123.12",
  "y": "456.56",
  "z": "789.89",
  "vel": "20.0"
}
```

Response
```json
{
  "loc": 1389.57
}
```

