# Booking Service API

## Overview
This project is a service that manages spaceflight bookings using various launchpads and destinations.
It includes functionalities for creating, retrieving, and deleting bookings.

## Project Structure
- **cmd/**: Contains the command-line tools (API server, migrations, schedule generator).
    - **api/**: Starts the API server.
    - **migrate/**: Runs database migrations.
    - **schedule/**: Generates launchpads schedules.
- **internal/**: Core business logic and service implementations.
  - **api/**: API routing and handlers.
  - **database/**: Database handling, migrations, and interface.
  - **external/**: External API client for SpaceX API.
  - **service/**: Booking service implementation.
  - **utils/**: Utility functions and helpers.
- **models/**: Defines data structures (e.g., `Booking`, `Schedule`).
- **postman/**: Contains a postman collection with an example requests to the API.

### Prerequisites
- Docker & Docker Compose

### Setup Instructions
1. Clone the repository.
   ```bash
   git clone https://github.com/klemis/go-spaceflight-booking-api.git
   cd go-spaceflight-booking-api
   ```
2. Run the database, migration, schedule generator and API server using Docker Compose.
    ```bash
    docker-compose up --build
    ```
