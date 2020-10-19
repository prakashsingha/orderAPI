# Go API Server

This is the Order API. It provides RESTful operation for customer's orders.

## Overview

- API version: 1.0.0-oas3
- Build date: 2020-10-18T03:40:16.495Z[GMT]

## API Endpoints

Please follow the following URL to get the full structure of the API.

[OAS Documentation](https://app.swaggerhub.com/apis/prakashsingha/order-api/1.0.0-oas3)

## Data Structure

- Database: oms
  - Collections:
    - orders
    - payments

### Assumptions

- The details of hotel, room and customer in orders collection maintain referential integrity with the respective collections.
- Models are validated before adding/updating the collection.
- Only a room can be booked in an order.
- Only a customer detail can be added in an order.
- The database contains lookup collections like hotels and customers.
- The config data (/config/config.go) has database configurations. However, only 'DB_URL' is used to connect local db server.

### Running the server

To run the server, follow these simple steps:

```
- Set the environment variable, DB_URL, to match your mongodb URI.
- go run main.go
```
