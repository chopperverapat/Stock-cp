# Stock Product Management System

Welcome to the Stock Product Management System project repository! This project aims to provide a web-based system for managing stock products, transactions, and authentication using the Go programming language : the Gin web framework and database : gorm and sqlite .

## Table of Contents

- [Stock Product Management System](#stock-product-management-system)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Installation](#installation)
  - [Usage](#usage)
  - [API Endpoints](#api-endpoints)
  - [Contributing](#contributing)

## Introduction

This project is developed using Go and Gin framework to create a stock product management system. It includes features like authentication, product management, and transaction tracking.

## Features

- User authentication (login and registration)
- Product management (create, read, update, delete)
- Transaction tracking (create, read)
- API endpoints for interaction with the system

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/chopperverapat/stock_product.git

2. Navigate to the project directory:
   ```bash
   cd stock_product
   ```

3. Install dependencies:
  ```go
  go mod download
  ```

4. Set up your database configuration in db/config.go.


5. Run the application:
   ```go
   go run server.go
   ```


## Usage

- Access the web application by visiting [http://localhost:8080](http://localhost:8080) in your web browser.

## API Endpoints

The API endpoints are designed to interact with the system programmatically. Here are some of the available endpoints:

- `/api/v2/login`: Authenticate user and generate a JWT token.
- `/api/v2/register`: Register a new user.
- `/api/v2/product`: Get a list of products.
- `/api/v2/product/:id`: Get product details by ID.
- `/api/v2/transaction`: Get a list of transactions.
- `/api/v2/transaction`: Create a new transaction (requires JWT token).
Refer to the source code for a complete list of API endpoints and their functionality.

## Contributing

Contributions to this project are welcome! If you find any issues or would like to add new features, feel free to open an issue or submit a pull request.

