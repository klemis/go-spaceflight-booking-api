Here’s how you can structure the documentation for the `internal/api` package, covering the three endpoints you mentioned:

---

# API Documentation

## Overview

The `internal/api` package provides HTTP endpoints for managing bookings. The API is organized under the `/api/v1` namespace and allows users to create, retrieve, and delete bookings. The following endpoints are available:

- **POST /api/v1/bookings**: Create a new booking.
- **GET /api/v1/bookings**: Retrieve all bookings.
- **DELETE /api/v1/bookings/:id**: Delete a booking by its ID.

---

## Endpoints

#### 1. Create a Booking

- **Endpoint**: `/bookings`
- **Method**: `POST`
- **Description**: Creates a new booking for a space launch based on the provided details.

**Request Body**:
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "gender": "Male",
  "birthday": "1985-05-15T00:00:00Z",
  "destination_id": 1,
  "launch_date": "2024-12-01T00:00:00Z"
}
```

**Request Body Fields**:
- `first_name` (string, required): The first name of the person booking the launch. Must be between 2 and 50 characters.
- `last_name` (string, required): The last name of the person booking the launch. Must be between 2 and 50 characters.
- `gender` (string, required): The gender of the person booking the launch. Must be between 2 and 50 characters.
- `birthday` (ISO 8601 date, required): The date of birth of the person booking the launch.
- `destination_id` (integer, required): The ID of the destination for the launch. Must be a value between 1 and 7, representing predefined destinations.
  - Destination mapping:
    - `1` — **Mars**
    - `2` — **Moon**
    - `3` — **Pluto**
    - `4` — **Asteroid Belt**
    - `5` — **Europa**
    - `6` — **Titan**
    - `7` — **Ganymede**
- `launch_date` (ISO 8601 date, required): The date of the launch.

**Response**:
- `201 Created`: Returns the created booking details.
- `400 Bad Request`: If validation fails or the request body is invalid.
- `500 Internal Server Error`: If an internal error occurs.

---

#### 2. Get All Bookings

- **Endpoint**: `/bookings`
- **Method**: `GET`
- **Description**: Retrieves a list of all bookings.

**Response**:
```json
[
  {
    "id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "gender": "Male",
    "birthday": "1985-05-15T00:00:00Z",
    "launchpad_id": "5e9e4501f5090910d4566f83",
    "destination_id": 1,
    "launch_date": "2024-12-01T00:00:00Z"
  },
  {
    "id": 2,
    "first_name": "Jane",
    "last_name": "Doe",
    "gender": "Female",
    "birthday": "1989-05-15T00:00:00Z",
    "launchpad_id": "5e9e4501f5090910d4566f83",
    "destination_id": 3,
    "launch_date": "2024-10-01T00:00:00Z"
  }
]
```

**Response Fields**:
- `id` (integer): The unique identifier for the booking.
- `first_name` (string): The first name of the person who made the booking.
- `last_name` (string): The last name of the person who made the booking.
- `gender` (string): The gender of the person who made the booking.
- `birthday` (ISO 8601 date): The birth date of the person who made the booking.
- `launchpad_id` (string): The ID of the launchpad assigned for the booking.
- `destination_id` (integer): The ID of the destination.
- `launch_date` (ISO 8601 date): The date of the launch.

**Response Codes**:
- `200 OK`: Returns the list of bookings.
- `404 Not Found`: If no bookings are found.
- `500 Internal Server Error`: If an internal error occurs.

---

#### 3. Delete a Booking

- **Endpoint**: `/bookings/:id`
- **Method**: `DELETE`
- **Description**: Deletes a booking based on its ID.

**URL Parameters**:
- `id` (integer, required): The unique identifier of the booking to delete.

**Response**:
- `200 OK`: If the booking is successfully deleted.
- `404 Not Found`: If no booking with the given ID is found.
- `500 Internal Server Error`: If an internal error occurs.

---

## Error Handling

All endpoints return appropriate HTTP status codes. In case of an error, the response will include a JSON object with an `error` field describing the issue.