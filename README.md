## Backend folder structure for future reference

```
project/
├── cmd/
│   └── app/
│       └── main.go      # Main application logic
├── internal/
│   ├── user/            # Feature: User
│   │   ├── handler/      # User-specific HTTP Handlers
│   │   ├── service/      # User-specific Business Logic
│   │   ├── repository/   # User-specific Data Access
│   │   └── user.go       # User Model
│   ├── product/         # Feature: Product
│   │   ├── handler/      # Product-specific HTTP Handlers
│   │   ├── service/      # Product-specific Business Logic
│   │   ├── repository/   # Product-specific Data Access
│   │   └── product.go    # Product Model
│   ├── order/           # Feature: Order
│   │   ├── handler/      # Order-specific HTTP Handlers
│   │   ├── service/      # Order-specific Business Logic
│   │   ├── repository/   # Order-specific Data Access
│   │   └── order.go      # Order Model
├── pkg/                 # Shared utilities or helpers
│   └── logger.go        # Logging utilities
├── configs/             # Configuration files
├── go.mod               # Go module definition
└── go.sum               # Go module checksum file
```
