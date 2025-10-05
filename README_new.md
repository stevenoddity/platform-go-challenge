# GWI challenge

A dashboard for assets. 


<i>This project is the [technical assignment for GWI](https://github.com/GlobalWebIndex/platform-go-challenge).</i>


## Table of Contents

- Technical Implementation
- Software versions
- Installation & Deployment
- Features
- Data population
- Authentication
- Potential enhancements
- Bugs
- License
- Contact

## Technical Implementation
- The backend part has been implemented by using Go.

The main goals during development were:
- Separation of concerns: each entity has a route, a service and a model.
- Security:
    - JWT
    - Sanitize input (validator)
    - Unified response
- Error handling
- Testing
- Clear naming and documentation

## Software versions
- Python: 1.25.1
- Docker: 28.4.0

## Installation & Deployment

```bash
# Clone the repository
git clone git@github.com:stevenoddity/platform-go-challenge.git

# Navigate to the project directory
cd platform-go-challenge

# Build docker image
docker build -t gwi-challenge .

# Run docker image 
docker run -p 5000:5000 gwi-challenge

```

```bash
Application can be accessed at http://127.0.0.1:8080/
```
Command for running tests:
```bash
TODO
```

# Features
- List Favorites:
```bash
curl -X GET -H "Authorization: Bearer $JWT" "http://127.0.0.1:8080/favorites?user_id=1"
```
- Add a new Favorite
```bash
curl -X GET -H "Authorization: Bearer $JWT" "http://127.0.0.1:8080/favorites?user_id=1"
```
- Delete a Favorite
```bash
curl -X GET -H "Authorization: Bearer $JWT" "http://127.0.0.1:8080/favorites?user_id=1"
```
- Edit Asset's description
```bash
curl -X GET -H "Authorization: Bearer $JWT" "http://127.0.0.1:8080/favorites?user_id=1"
```

# Data population




# Authentication
## JWT token for user_id = 1:
```bash
 eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.fKCJWNXwhs7ukzI7vpAN2v1z5PBFmiqLlAEhoxbuDB4
```
## Generate JWT token ad hoc:
Use secret-key at https://www.jwt.io/
```bash
secret_key="gwi-jwt-secret"
```

# Potential Enhancements
- Enhance testing
- Use database
- Logging
- Pagination
- Authorization on authenticated users
- Expiration in JWT
- Blacklist old JWT tokens
- Rate limiter on endpoints

# Bugs

Pending 

# Licences

# Contact

Don't hesitate to ask me for clarifications at steveofsam@gmail.com
