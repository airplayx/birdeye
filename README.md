# Birdeye Go SDK

A Go SDK for accessing the Birdeye API. This SDK provides a simple interface for interacting with Birdeye's blockchain data services.

## Features

- Generic client implementation for Birdeye API
- Type-safe API calls
- Built-in HTTP client configuration
- Support for the following API endpoints:
  - Trade Data
  - Token Overview
  - Token List

## Installation

```bash
go get github.com/airplayx/birdeye
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/airplayx/birdeye"
)

func main() {
    // Create a new client instance
    client := birdeye.NewClient("your-api-key")
    
    // Use the client to make API calls
    // ...
}
```

## Project Structure

- `birdeye.go` - Core client implementation and generic interfaces
- `trade_data.go` - Trade data API implementation
- `token_overview.go` - Token overview API implementation
- `token_list.go` - Token list API implementation

## Requirements

- Go 1.23.3 or higher

## License

[MIT License](LICENSE) 