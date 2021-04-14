# First api in go

## objective:
   - build an api in go to be used as azure function(s)

## steps:
    - add standart web sever in go that will be used to local test
    - add unitest to cover net/http code
    - write openapi specs for functions
    - generate golang code from the openapi specs
    - implement ann add unit tests
    - deploy function in azure and connect it to the iot hub, validate by running the iot simulation

## Go tips:
    - run the web server (local test with Postman for instance):
        go run .
    - run tests with coverage: 
        go test -cover