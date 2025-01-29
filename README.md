
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


## Installation

Prerequisites:
   - Install Go
   - Install Docker


  ### Clone the repo

- Clone the repo & intall packages

  ```bash
    git clone https://github.com/ayushdiitr/url-shortener.git
    cd url-shortener
  ```
- Install Dependencies
  
  ```bash
  go mod tidy
  ```
- Create .env
  
  ```bash
    DOMAIN=http://localhost:3000  # Your API domain
    API_QUOTA=10                  # Max requests per user
    DB_URL= "db:6379"             # Redis connection
    DB_PASSWORD= ""
    PORT= ":3000"
  ```
- Docker Compose
  
  ```bash
  docker-compose up -d
  ```

  ### API Endpoints

#### Shorten the URL

```http
  POST /api/items
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `url` | `string` | **Required**. URL that you want to shorten |


#### Resolve a shortened URL

```bash
  curl -X GET http://localhost:3000/{shortened_url}
```
- Redirects to the original URL.





    
