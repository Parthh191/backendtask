# User Management API with Age Calculation

A RESTful API built with Go and PostgreSQL to manage users with their date of birth and dynamically calculated age.

## 🚀 Features

- **User Management**: Create, read, update, and delete users
- **Age Calculation**: Automatically calculates user age from date of birth
- **Search**: Search users by name
- **RESTful API**: Standard REST endpoints
- **Database Migrations**: SQL migration files included
- **Logging**: Request logging and error handling
- **CORS Support**: Cross-origin requests enabled
- **Error Handling**: Comprehensive error handling and validation
- **PostgreSQL**: Uses PostgreSQL database (compatible with Neon)

## 🛠️ Tech Stack

- **Language**: Go 1.22+
- **Web Framework**: Gin
- **Database**: PostgreSQL (Neon-compatible)
- **Environment Management**: godotenv

## 📋 Prerequisites

- Go 1.22 or higher
- Neon PostgreSQL database account
- Git

## 🔧 Installation & Setup

### 1. Setup Repository
```bash
cd /home/coderwizard/backendtask
```

### 2. Get Your Neon Database URL

1. Go to [console.neon.tech](https://console.neon.tech)
2. Create or select your project
3. Click the "Connection String" button
4. Copy the PostgreSQL connection string

### 3. Configure Environment Variables

Create a `.env` file in the root directory:

```env
DATABASE_URL=postgresql://username:password@host.neon.tech:5432/dbname?sslmode=require
PORT=8080
ENV=development
```

**Example with Neon:**
```env
DATABASE_URL=postgresql://neondb_owner:abc123xyz@ep-cool-wave-12345.us-east-1.aws.neon.tech/neondb?sslmode=require
PORT=8080
ENV=development
```

### 4. Install Dependencies
```bash
go mod download
go mod tidy
```

### 5. Run Database Migrations

Execute the migration to create the users table in Neon:

```bash
psql $DATABASE_URL -f db/migrations/001_create_users_table.up.sql
```

### 6. Start the Server

```bash
go run cmd/server/main.go
```

The API will start on `http://localhost:8080`

## 📚 API Endpoints

### Health Check
- **GET** `/health` - API health status

### Users API (v1)

#### Create User
- **POST** `/api/v1/users`
- Request Body:
```json
{
  "name": "John Doe",
  "dob": "1990-05-15"
}
```
- Response: `201 Created`

#### Get All Users
- **GET** `/api/v1/users`
- Response: `200 OK` - Returns array of users with calculated ages

#### Get User by ID
- **GET** `/api/v1/users/:id`
- Response: `200 OK`

#### Update User
- **PUT** `/api/v1/users/:id`
- Request Body:
```json
{
  "name": "Jane Doe",
  "dob": "1992-03-20"
}
```
- Response: `200 OK`

#### Delete User
- **DELETE** `/api/v1/users/:id`
- Response: `204 No Content`

#### Search Users
- **GET** `/api/v1/users/search?name=John`
- Response: `200 OK` - Returns matching users

## 📊 Example Response

```json
{
  "id": 1,
  "name": "John Doe",
  "dob": "1990-05-15",
  "age": 34,
  "created_at": "2025-12-19T10:30:00Z",
  "updated_at": "2025-12-19T10:30:00Z"
}
```

## 🗂️ Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go                    # Entry point
├── config/
│   └── config.go                      # Configuration management
├── db/
│   └── migrations/
│       ├── 001_create_users_table.up.sql
│       └── 001_create_users_table.down.sql
├── internal/
│   ├── handler/
│   │   └── user_handler.go            # HTTP handlers
│   ├── logger/
│   │   └── logger.go                  # Logging utility
│   ├── middleware/
│   │   └── middleware.go              # HTTP middleware
│   ├── models/
│   │   └── user.go                    # Data models
│   ├── repository/
│   │   └── user_repository.go         # Database access layer
│   ├── routes/
│   │   └── routes.go                  # Route definitions
│   └── service/
│       └── user_service.go            # Business logic
├── .env                               # Environment variables (local)
├── .env.example                       # Example environment variables
├── go.mod                             # Go module file
├── go.sum                             # Go dependencies lock file
└── README.md                          # This file
```

## 🔐 Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | `postgresql://user:pass@host:5432/db` |
| `PORT` | Server port | `8080` |
| `ENV` | Environment mode | `development` or `production` |

## 🧪 Testing Endpoints

### Using cURL

**Create a user:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "dob": "1990-05-15"
  }'
```

**Get all users:**
```bash
curl http://localhost:8080/api/v1/users
```

**Get user by ID:**
```bash
curl http://localhost:8080/api/v1/users/1
```

**Update user:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "dob": "1992-03-20"
  }'
```

**Delete user:**
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

**Search users:**
```bash
curl "http://localhost:8080/api/v1/users/search?name=John"
```

## 🚀 Running in Production

1. Update `.env` for production:
   ```env
   ENV=production
   DATABASE_URL=your_neon_connection_string
   PORT=8080
   ```

2. Build the binary:
   ```bash
   go build -o backendtask cmd/server/main.go
   ```

3. Run the binary:
   ```bash
   ./backendtask
   ```

## 📝 Database Schema

### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    dob DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🐛 Error Handling

The API returns appropriate HTTP status codes and error messages:

- `200 OK` - Successful GET/PUT request
- `201 Created` - User successfully created
- `204 No Content` - Successful DELETE request
- `400 Bad Request` - Invalid input or validation error
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

## 📄 License

This project is open source and available under the MIT License.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📧 Support

For issues and questions, please create an issue in the repository.
# backendtask
