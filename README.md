# GOSON-Server

A fast and extensible JSON RESTful server written in Go, inspired by [json-server](https://github.com/typicode/json-server).  

> [!IMPORTANT]  
> **v0.2.0** This version is in active development and may have substantial changes in future releases.
---

## Features

- **Plug-and-play REST API**: Instantly expose CRUD endpoints for any JSON file.
- **Full CRUD support**: List, filter, create, update, and delete for any collection.
- **Advanced Filtering**: Query by any field, with operators like `field_lt`, `field_gte`, `field_cs`, etc.
- **Native Pagination & Ordering**: Safe paging (never out-of-bounds) and ascending/descending ordering.
- **JSON:API-style responses**: Clean, consistent, and ready for frontend consumption.
- **Robust data validation**: Handles empty bodies, malformed JSON, and preserves unique IDs.

## Quick Start

### Installation

```bash
# Download a prebuilt executable (see Releases for your platform)
# Or build from source:
git clone https://github.com/tdalexm/goson-server.git
cd goson-server
go build -o goson-server
```

### Basic Usage

1. Create a `db.json` file:
    ```json
    {
      "pokemon": [
        { "id": "1", "name": "Bulbasaur", "type": "Grass/Poison", "level": 5 }
      ]
    }
    ```

> [!NOTE]
> The .zip release contains an example `db.json` ready for testing.

2. Start the server:
    ```bash
    ./goson-server --db=db.json --port=8080
    ```

3. Your REST API is now available at `http://localhost:8080`

## API Reference

### Main Endpoints

| Method | Route                 | Description                            |
|--------|-----------------------|----------------------------------------|
| GET    | `/:collection`        | List & filter resources                |
| GET    | `/:collection/:id`    | Get resource by ID                     |
| POST   | `/:collection`        | Create a new resource                  |
| PUT    | `/:collection/:id`    | Overwrite a resource (idempotent)      |
| PATCH  | `/:collection/:id`    | Partial update to a resource           |
| DELETE | `/:collection/:id`    | Delete a resource                      |

#### Sample List Query

```http
GET /pokemon?page=2&limit=5&type_cs=grass&sort=desc
```

- `page`, `limit`: Safe paging—never returns out of bounds.
- Filtering: Use suffixes like `_lt`, `_gt`, `_lte`, `_gte`, `_ne`, `_cs` for less/greater/contains/not-equal, etc.
- `sort=desc`: Returns the collection in reverse order.

## Filtering

You can filter the results of any `GET /:collection` request by passing filter parameters as query strings.  
Filters are based on the field name and can include special suffixes to specify the matching operation:

| Suffix    | Meaning                    | Example Query                | Description                                |
|-----------|----------------------------|------------------------------|--------------------------------------------|
|   | Equals                     | `?type=Fire`                 | Exact match (`type == "Fire"`)             |
| `_lt`     | Less than                  | `?level_lt=10`               | Field less than value (`level < 10`)       |
| `_lte`    | Less than or equal         | `?level_lte=5`               | Field <= value                             |
| `_gt`     | Greater than               | `?level_gt=20`               | Field > value                              |
| `_gte`    | Greater than or equal      | `?level_gte=12`              | Field >= value                             |
| `_ne`     | Not equal                  | `?name_ne=Bulbasaur`         | Field not equal to value                   |
| `_cs`     | Contains substring         | `?name_cs=saur`              | Field contains value (case-insensitive)    |

**You can combine multiple filters in one request:**  
Example:  
```
GET /pokemon?type=Fire&level_gte=5&name_cs=char
```
This will return all pokemon of type "Fire" whose level is **at least 5** and whose name contains "char" (like "Charmander").

**Additional notes:**
- Filtering by `id` is not allowed. To query by `id`, use `GET /:collection/:id`.
- All filters are **case-insensitive** for string values.
- When combining filters, only items matching **all** criteria (AND logic) are returned.

## CLI Configuration

```bash
./goson-server --help

--db string       Path to the JSON database file (default "db.json")
--port string     Port to serve the API on (default "8080")
--help            Show help

Examples:
  ./goson-server --db=./data/db.json --port=8000
```

## What's New

- **Better Hexagonal architecture implementation**
- **Pagination and filtering:** Supports combinations of filters with safe range slicing.
- **JSON:API style responses**.
- **ID field protection:** IDs cannot be updated or deleted directly.

## Requirements

- Go 1.21+
- Gin Web Framework (`v1.9`)

## Roadmap

- Real persistence
- Concurrency

## Feedback

- [Report a bug](https://github.com/tdalexm/goson-server/issues)
- [Request a feature](https://github.com/tdalexm/goson-server/issues)

## Acknowledgements

- [typicode/json-server](https://github.com/typicode/json-server)
