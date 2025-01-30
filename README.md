# Automobile Rental Backend

A Go-based backend service for managing vehicle rentals, featuring user authentication, vehicle management, and booking systems.

## 🚀 Features

- User authentication with JWT
- Vehicle management system
- Booking and reservation handling
- Role-based access control
- PostgreSQL database integration
- Docker support
- CI/CD with GitHub Actions

## 📋 Prerequisites

- Go 1.21+
- PostgreSQL 15+
- Docker & Docker Compose (optional)
- Make (optional)

## 🛠️ Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go          # Application entry point
├── internal/
│   ├── config/             # Configuration setup
│   ├── handlers/           # HTTP request handlers
│   ├── middleware/         # Custom middleware
│   ├── models/             # Data models
│   ├── repositories/       # Database operations
│   ├── routes/             # Route definitions
│   ├── services/           # Business logic
│   └── utils/              # Helper functions
├── migrations/             # Database migrations
├── pkg/                    # Public packages
├── docker-compose.yaml     # Docker compose configuration
├── Dockerfile             # Docker build instructions
├── Makefile              # Build automation
├── go.mod                # Go dependencies
└── README.md             # Project documentation
```

## 🚀 Quick Start

### Local Development

1. Clone the repository
```bash
git clone <repository-url>
cd automobile-rental
```

2. Set up environment variables
```bash
cp .env.example .env
# Edit .env with your configurations
```

3. Install dependencies
```bash
go mod download
```

4. Run database migrations
```bash
make migrate
```

5. Start the application
```bash
make run
```

### Docker Setup

1. Build and start services
```bash
docker-compose up --build
```

2. Stop services
```bash
docker-compose down
```

## 🔑 Environment Variables

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=rental
DB_PORT=5432
JWT_SECRET=your-secret-key
```

## 🛣️ API Endpoints

### Authentication
- POST `/auth/signup` - Create new user account
- POST `/auth/login` - User login

### Vehicles
- GET `/vehicles` - List all vehicles
- POST `/vehicles` - Create new vehicle (Admin only)
- GET `/vehicles/:id` - Get vehicle details
- PUT `/vehicles/:id` - Update vehicle (Admin only)
- DELETE `/vehicles/:id` - Delete vehicle (Admin only)

### Bookings
- GET `/bookings` - List user's bookings
- POST `/bookings` - Create new booking
- GET `/bookings/:id` - Get booking details
- PUT `/bookings/:id` - Update booking status

## 🔒 Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your-token>
```

## 🧪 Testing

Run the test suite:

```bash
make test
```

## 🔨 Available Make Commands

```bash
make run          # Run the application
make test         # Run tests
make migrate      # Run database migrations
make docker-build # Build Docker images
make docker-up    # Start Docker services
make docker-down  # Stop Docker services
```

## 🐳 Docker Support

The application can be run using Docker Compose:

```bash
# Build and start services
docker-compose up --build

# Stop services
docker-compose down

# View logs
docker-compose logs -f
```

## 📝 Database Models

### User
```go
type User struct {
    gorm.Model
    FirstName string
    LastName  string
    Email     string
    Password  string
    Role      string
    Bookings  []Booking
}
```

### Vehicle
```go
type Vehicle struct {
    gorm.Model
    Make         string
    Model        string
    Year         int
    LicensePlate string
    DailyRate    float64
    Status       string
    Bookings     []Booking
}
```

### Booking
```go
type Booking struct {
    gorm.Model
    UserID     uint
    VehicleID  uint
    StartDate  time.Time
    EndDate    time.Time
    TotalCost  float64
    Status     string
    PaymentID  string
}
```

## 🔍 CI/CD

The project uses GitHub Actions for continuous integration:
- Runs tests on every push
- Includes PostgreSQL service container for integration tests
- Go 1.21 environment

## 📚 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 🪲 Common Issues & Solutions

1. Database Connection Issues
```bash
# Check PostgreSQL service
docker-compose ps

# View logs
docker-compose logs db
```

2. Migration Failures
```bash
# Reset migrations
make migrate-reset

# Run migrations again
make migrate
```