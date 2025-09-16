# Architecture Overview

Spanish version: [esp/architecture.md](esp/architecture.md)

## Code Structure

The codebase is organized into several directories and files to maintain a Clean Architecture. The main components are:
- `cmd`: Contains the entry points for the application.
- `models`: Defines the data models used in the application.
- `packages`: Includes different packages that encapsulate specific functionalities:
    - `cache`: Manages caching mechanisms.
    - `database`: Handles database operations.
    - `service`: Contains business logic services.
    - `server`: API server implementation and metrics.
- `storage`: Stores data files like JSON files.

Clean Architecture is used as it the most idiomatic way to organize Go code and promotes more clear separation of concerns and makes the codebase easier to test and maintain. The code should be organized as simple as possible (we are not Java devs :) )


## Project Structure

Basically there are 3 components now:
1. API server, that is responsible for handling HTTP requests and responses.
2. Redis cache, that is used to store frequently accessed data.
3. Database, that is used to store persistent data.

BUT, by the task the Database should be replaced with a JSON file storage. However, it still keeps the interface functionality to future changes.


