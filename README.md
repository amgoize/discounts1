# Discounts Service

**Discounts Service** is a RESTful API written in **Go** using **Gorilla Mux** for routing and **GORM** for database management. This service allows users to manage clients and their associated discount coupons efficiently.

## Features

- **Client Management**: Create, update, delete, and retrieve client data.
- **Coupon Management**: Create, update, delete, and retrieve coupons.
- **Database Integration**: Uses GORM to manage interactions with a relational database.
- **Error Handling**: Provides meaningful error messages and status codes.

## Project Structure
**models/:** Contains the data models (Client, Coupon) that represent the structure of the data in the database.
**repositories/:** Handles database interactions, abstracting the data access logic from the rest of the application.
**services/:** Implements the business logic, interacting with the repositories and preparing data for the handlers.
**handlers/:** Manages HTTP requests and responses, communicating with services to fulfill client requests.
