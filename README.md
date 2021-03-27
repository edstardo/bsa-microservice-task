# bsa-microservice-task

Go Installation
  https://golang.org/doc/install

Run proxy-microservice
  ```console
    cd proxy-microservice
    go run main.go
  ```

Run user-microservice
  ```console
    cd user-microservice
    go run main.go
  ```

Run auth-microservice
  ```console
    cd auth-microservice
    go run main.go
  ```

Hardcoded valid usersnames are ["user1", "user2"]

Test
* All three microservices should be running
* Unauthorized requests
  ```console
    curl localhost:3000/user/profile
    curl -H "Username: unauthorized-user" localhost:3000/user/profile
  ```
* Authorized requests
  ```console
    curl -H "Username: user1" localhost:3000/user/profile
    curl -H "Username: user2" localhost:3000/user/profile
  ```
* No auth required
  ```console
    curl localhost:3000/microservice/name
  ```

