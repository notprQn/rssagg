# API Documentation

This document provides an overview of the RESTful API implemented in the provided Go code. The API is designed to manage users, feeds, and posts, with additional functionality for handling feed follows.

## Table of Contents

- [API Link](#api-link)
- [Endpoints](#endpoints)
    - [Users](#users)
    - [Feeds](#feeds)
    - [Feed Follows](#feed-follows)
    - [Posts](#posts)
    - [Health Check](#health-check)
- [Authentication](#authentication)
- [Scraping](#scraping)

## API Link

[rssagg-production.up.railway.app](https://rssagg-production.up.railway.app)

## Endpoints

### Users

- **Create User**
  - `POST /v1/users`
  - Creates a new user.
  - Parametes:
    `name (string): User name.`

- **Get Users**
  - `GET /v1/users`
  - Retrieves information about users. Requires authentication.

### Feeds

- **Create Feed**
  - `POST /v1/feeds`
  - Creates a new feed. Requires authentication.
  - Parametes:
    `name (string): User name.`
    `url (string): website url.`

- **Get Feeds**
  - `GET /v1/feeds`
  - Retrieves information about feeds.

### Feed Follows

- **Get Feed Follows**
  - `GET /v1/feed_follows`
  - Retrieves information about feed follows. Requires authentication.
  - Parametes:
    `feed_id (uuid): feed id.`

- **Create Feed Follow**
  - `POST /v1/feed_follows`
  - Creates a new feed follow. Requires authentication.

- **Delete Feed Follow**
  - `DELETE /v1/feed_follows/{feedFollowID}` 
  - Deletes a feed follow by ID. Requires authentication.

### Posts

- **Get Posts**
  - `GET /v1/posts`
  - Retrieves information about posts. Requires authentication.

### Health Check

- **Readiness Check**
  - `GET /v1/healthz`
  - Checks the readiness of the API.

## Authentication

The API uses token-based authentication. To access endpoints that require authentication, include the token in the request header.

```plaintext
Authorization: ApiKey + YOUR_TOKEN
```

### Scraping

The application includes a scraping routine that runs concurrently to collect data from external sources. It starts automatically with the server. The scraping configuration can be adjusted in the startScraping function.

Note: Ensure that your PostgreSQL database is properly configured and accessible before running the application.

Feel free to customize the code according to your requirements. For more details on Go-Chi and other libraries used, refer to their respective documentation.

Enjoy using the API! If you encounter any issues or have suggestions, please create an issue on the GitHub repository.
