# Golang - A Product Subscription Service

This is a product (imaginary) subscription service, where users can subscribe based on three planes - Gold, Bronze &
Silver. This project has been built to showcase various concurrency instances.

## Install

- Install `Go`
- Install `docker` and `docker-compose`
- Install `Git`
- Clone the repository using `git clone <repository address>`
  - Run `docker compose up -d` or `docker-compose up -d`
  - Run `make start` to start the server
  - Server will be running on `http://localhost:8080`
  - To stop the server run `make stop`
  - To see mailhog emails, open `http://localhost:8025/` in the browser

## Use cases of concurrencies

  - Sending email
  - Graceful shutdown
  - Generate an invoice
  - Generate an user manual as a PDF
