# Birdeye Go SDK

A Go SDK for accessing the Birdeye API. This SDK provides a simple interface for interacting with Birdeye's blockchain data services.

## Features

- Generic client implementation for Birdeye API
- Type-safe API calls
- Built-in HTTP client configuration with 30-second timeout
- Comprehensive API endpoint support
- Automatic request/response handling
- JSON response parsing

## Supported API Endpoints

The SDK supports the following API endpoints:

- Trade Data
  - Single trade data
  - Multiple trade data
- Token Information
  - Token Overview
  - Token List (V1, V3, V3Scroll)
  - Token Trending
  - Token Security
  - Token Creation Info
- Market Data
  - Single market data
  - Multiple market data
- Metadata
  - Single metadata
  - Multiple metadata
- Additional Features
  - Holder Information
  - New Listings
  - Top Traders
  - Markets
  - Mint/Burn Transactions
  - Trades

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
    
    // Example: Get token overview
    // tokenOverview, err := client.GetTokenOverview("token-address")
    
    // Example: Get trade data
    // tradeData, err := client.GetTradeData("token-address")
    
    // Example: Get market data
    // marketData, err := client.GetMarketData("token-address")
}
```

## API Base URL

The SDK uses the following base URL for all API requests:
```
https://public-api.birdeye.so/defi
```

## Project Structure

- `birdeye.go` - Core client implementation and generic interfaces
- `trade_data.go` - Trade data API implementation
- `token_overview.go` - Token overview API implementation
- `token_list.go` - Token list API implementation

## Requirements

- Go 1.23.3 or higher

## Error Handling

The SDK provides detailed error messages for various scenarios:
- Request creation errors
- HTTP request errors
- Unexpected status codes
- Response parsing errors

## License

[MIT License](LICENSE) 