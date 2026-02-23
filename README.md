# Country Info Service

A simple REST API service that combines data from REST Countries and Currency APIs.

## What it does

This service has three endpoints:
1. Check if external APIs are working
2. Get information about a country
3. Get currency exchange rates for a country and its neighbors

## How to run
```bash
go run main.go
```

The service starts on `http://localhost:8080`

## Endpoints

### 1. Status Check
```bash
curl http://localhost:8080/countryinfo/v1/status/
```

Returns:
```json
{
  "restcountriesapi": 200,
  "currenciesapi": 200,
  "version": "v1",
  "uptime": 120
}
```

### 2. Country Info
```bash
curl http://localhost:8080/countryinfo/v1/info/no
```

Returns country name, population, area, languages, borders, flag, and capital.

### 3. Exchange Rates
```bash
curl http://localhost:8080/countryinfo/v1/exchange/no
```

Returns currency exchange rates between Norway and its neighbors (Sweden, Finland, Russia).

## Examples
```bash
# Norwegian info
curl http://localhost:8080/countryinfo/v1/info/no

# Swedish info
curl http://localhost:8080/countryinfo/v1/info/se

# German exchange rates
curl http://localhost:8080/countryinfo/v1/exchange/de
```

## Project Structure
```
assignment-1/
├── main.go           # Main server
├── handlers/         # Handle HTTP requests
├── services/         # Call external APIs
└── models/           # Data structures
```

## External APIs Used

- REST Countries: `http://129.241.150.113:8080/v3.1/`
- Currency API: `http://129.241.150.113:9090/currency/`

## Notes

- Only uses Go standard library (no external packages)
- 10 second timeout on external API calls
- Returns JSON for all responses
```

```
## Use of AI

This project was primarily implemented manually. AI tools (such as ChatGPT and Claude) were used as assistive tools for:

- Understanding Go concepts and REST API design
- Debugging issues and interpreting error messages
- Suggesting code structure and helping formulate function comments more clearly
- Reviewing documentation and assisting with formatting parts of the README where formatting was difficult

All code was reviewed, understood, and adapted by myself before being included in the project.