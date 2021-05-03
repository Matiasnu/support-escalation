# :memo: Go Support Escalation App


**Server: Golang  
Client: React, semantic-ui-react  
Database: Local MongoDB**


# Highlights

1. DB connection string, name and collection name moved to `.env` file as environment variable. Using `github.com/joho/godotenv` to read the environment variables.

## Application Requirement

### golang server requirement

1. golang https://golang.org/dl/
2. gorilla/mux package for router `go get -u github.com/gorilla/mux`
3. mongo-driver package to connect with mongoDB `go get go.mongodb.org/mongo-driver`
4. github.com/joho/godotenv to access the environment variable.

## :computer: Start the application

1. Make sure your mongoDB is started
2. Create a `.env` file inside the `go-server` and copy the keys from `.env.example` and update the DB connection string.
3. From go-server directory, open a terminal and run
   `go run main.go`
4. From client directory,  
   a. install all the dependencies using `npm install`  
   b. start client `npm start`

## :panda_face: Walk through the application

Open application at http://localhost:3000

## Author

Matias Nuñez

# References

https://godoc.org/go.mongodb.org/mongo-driver/mongo  
https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial  
https://vkt.sh/go-mongodb-driver-cookbook/

# License

MIT License

Copyright (c) 2019 Shubham Chadokar
