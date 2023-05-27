# Simple E-Commerce: Synapsis.id Backend Intern Challenge

## Entity Relationship Diagram

![Entity Relationship Diagram](erd.png)

## Features

- Customer can view product list by product category
- Customer can add product to shopping cart
- Customers can see a list of products that have been added to the shopping cart
- Customer can delete product list in shopping cart
- Customers can checkout and make payment transactions
- Login and register customers

## How to Run

### Prerequisites

Make sure you have the following prerequisites installed:

- Docker (version >= 23.0.1)
- Docker Compose (version >= v2.16.0)

### Configuration

Create a .env or replace .env.example / .env.example.replicated file in the project directory:

```
PORT=4001

DATABASE_USER=aryahmph
DATABASE_PASSWORD=aryahmph
DATABASE_HOST=postgres
DATABASE_PORT=5432
DATABASE_NAME=synapsis_challange
DATABASE_MIGRATIONS_PATH=pkg/sql/schema
DATABASE_URL=postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable

JWT_ACCESS_TOKEN=secret
XENDIT_SECRET_KEY=xnd_development_ye78i7MafcV7CF1UPhdFS6m4NjfzXfqtYjlQ9KDiJ3fNMZXG3AFdYfStgtbpZ
```

### Installation

1. Clone the repository:

```bash
git clone https://github.com/aryahmph/synapsis-challange
```

2. Change to the project directory:

```bash
cd synapsis-challange
```

3. Start the Docker containers using Docker Compose:

```bash
docker-compose up -d
```

4. Connect to `http://localhost:4001/api/v1`

## API Documentation

I use OpenAPI to document the API that has been created. To view the documentation, you can use tools such as Jetbrains
[OpenAPI Plugin](https://www.jetbrains.com/help/idea/openapi.html), [Swagger UI](https://swagger.io/tools/swagger-ui/),
or cURL in the [apidoc.yaml](apidoc.yaml) file.

## Simulate Payment

I use Xendit as payment gateway to perform transactions on the system. To simulate the payment process, run the
following request:

```bash
curl -X POST \
  https://api.xendit.co/pool_virtual_accounts/simulate_payment \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Basic eG5kX2RldmVsb3BtZW50X25oRXhrOGY5WXVvazJXajZ1dkZNeUxSSExRVWhQVVZMQWFOMm5ER0hNd3BLejBtY0piQ2JhUFlOYWwwd3AwMTo=' \
  -d '{
    "transfer_amount": 0,
    "bank_account_number":"va-number-example",
    "bank_code": "BRI"
  }'

```

## Deployment

I use AWS EC2 Free tier to deploy the project that has been created, which can be accessed at the following
link: `http://54.255.190.44/api/v1`

### Application Docker Image

There are also images that have been uploaded to the docker hub at the
following [link](https://hub.docker.com/repository/docker/aryahmph/synapsis-challange).

## Notes

### Problems when trying to migrate a database

When you want to migrate a database, you must first change the enviroment of the database host to `localhost` instead of
`postgres`.

---
**Thank you!**