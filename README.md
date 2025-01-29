
# URL Shortener with Rate Limiting

This is a simple and efficient URL shortener built using Go and Redis. It allows users to shorten URLs with optional custom aliases while enforcing a rate limiter to prevent abuse.


## Features

- Shorten URLs with randomly generated or custom short codes.
- Redirect shortened URLs to their original destinations.
- Rate limiting: Users can make up to 10 requests per 30 minutes.
- Expiry time for shortened URLs (default: 24 hours).
- Uses Redis for fast key-value storage.



## Tech Stack

- **Go** (Fiber Framework)
- **Redis** 
- **Docker** (Optional for containerization)


Installation

Prerequisites

Install Go

Install Redis

Install Docker (optional)

Clone the Repository

git clone https://github.com/yourusername/url-shortener.git
cd url-shortener

Install Dependencies

go mod tidy

Set Up Environment Variables

Create a .env file in the root directory and configure the following:

DOMAIN=http://localhost:3000  # Your API domain
API_QUOTA=10                  # Max requests per user
REDIS_HOST=localhost:6379      # Redis connection

Start Redis Server

redis-server

Run the Application

go run main.go

API Endpoints

Shorten a URL

POST /api/shorten

Request Body:

{
  "url": "https://example.com",
  "custom_short": "mycustom",  // Optional
  "expiry": 24                   // Expiry in hours (default: 24)
}

Response:

{
  "url": "https://example.com",
  "custom_short": "http://localhost:3000/mycustom",
  "expiry": 24,
  "rate_limit": 9,
  "rate_limit_reset": 30
}

Resolve a Shortened URL

GET /{short_url}

Example:

curl -X GET http://localhost:3000/mycustom

Redirects to the original URL.

Rate Limiting

Each user can make up to 10 requests per 30 minutes.

The rate limit resets automatically after 30 minutes.

If the limit is exceeded, the API responds with:

{
  "error": "Rate limit exceeded",
  "rate_limit_reset": 30
}

Docker Setup (Optional)

To run the application inside a Docker container:

docker-compose up --build

Future Improvements

Persistent database storage (e.g., MongoDB/PostgreSQL).

User authentication.

Analytics dashboard for URL usage tracking.

License

This project is open-source under the MIT License.


    