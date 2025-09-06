# RHGHospitality
RHG Hospitality Solution

# Reservation Backend
- Users -> Book rooms for an Hotel
- Admins -> Manage Hotels, Rooms, Bookings
- Authentication -> JWT
- Authorization -> Role-based access control
- Hotels -> CRUD operations, JSON 
- Rooms -> CRUD operations, JSON
- Scripts -> Database initialization, seeding
- Database -> MongoDB

# Database MongoDB Golang Driver

'go get go.mongodb.org/mongo-driver/v2/mongo'


# Go GIN Modules

'go get -u github.com/gin-gonic/gin'

# Database setup 

'docker-composer.yml 
version: '3.8'

services:
  mongo:
    image: mongo:latest
    container_name: mongo
    environment:
      - MONGO_INITDB_ROOT_USERNAME=xxxxxx
      - MONGO_INITDB_ROOT_PASSWORD=xxxxxx
    ports:
      - "27017:27017"
    volumes:
      - mongo-data:/data/db
    restart: always

volumes:
  mongo-data:

# .env file
MONGO_URI=mongodb://xxxxxx:xxxxxx@localhost:27017
DB_NAME=rhghospitality
JWT_SECRET=xxxxx
PORT=xxxx
ADMIN_USERNAME=xxxxx
ADMIN_PASSWORD=xxxxx
