# Voucher Management API

This is a Go-based REST API for managing vouchers, brands, and redemptions.

## Prerequisites

- Go 1.18 or later
- MySQL database

## Setup

1. Clone this repository:

   ```bash
   git@github.com:zain120500/voucher_management.git
   ```

2. Install dependencies:
    
    ```bash
    go mod tidy
    ```
3. Set up your MySQL database:

   Make sure you have a MySQL database named voucher_db or modify the DSN in main.go to match your database configuration.

4. Run database migrations::

    ```bash
    migrate -database "mysql://root:root@tcp(localhost:3306)/voucher_db" -path db/migrations up
    ```

5. Run the application:

    ```bash
    go run main.go
    ```
   The application will run on http://localhost:8080.

6. Run tests:

   Execute the test suite to ensure the application is working correctly:

    ```bash
    go test -v ./pkg/handler_test      
    ```

# Voucher Management API

## POST /brand

Create a new brand.

   
    {
    "name": "Brand Name",
    "vouchers": [
            {
            "title": "Voucher Title",
            "cost_in_point": 100
            }
        ]
    }
  


## POST /voucher

Create a new voucher.

    {
        "title": "Voucher Title",
        "cost_in_point": 100,
        "brand_id": 1
    }

## GET /voucher?id={id}

Get a voucher by ID.

## GET /voucher/brand?id={brand_id}

Get all vouchers for a specific brand.

## POST /transaction/redemption

Create a new redemption transaction.


    {
    "customer_name": "John Doe",
    "vouchers": [
            {
            "id": 1,
            "title": "Voucher A",
            "cost_in_point": 100,
            "brand_id": 1
            }
        ]
    }

## GET /transaction/redemption?transactionId={id}

Get redemption transaction details by ID.




