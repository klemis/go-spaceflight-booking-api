# Database Schema

## Tables

### Bookings
The `bookings` table defines the booking of the flight.
- **id**: Primary key.
- **first_name**: First name of the customer.
- **last_name**: Last name of the customer.
- **gender**: Gender of the customer.
- **birthday**: Date of birth.
- **launchpad_id**: ID of the launchpad used for the booking.
- **destination_id**: ID of the destination.
- **launch_date**: Date of the flight.

### Schedules
The `schedules` table defines the flight schedules for each launchpad.
Each launchpad has a unique destination for each day of the week.

- **id**: Primary key.
- **launchpad_id**: ID of the launchpad. This refers to the launchpad's unique identifier.
- **destination_id**: ID of the destination. This corresponds to the specific destination that the flight from the launchpad will go to on a given day.
- **day_of_week**: Day of the week when the flight is scheduled. Stored as an integer (0 for Sunday, 1 for Monday, etc.).
- **created_at**: Timestamp of when the schedule was created.
- **updated_at**: Timestamp of the last update to the schedule.

## Migrations
- All migrations are located in `internal/database/migrations/`.
- Migrations are executed automatically by the `migrate` binary.