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

## API V1
    - /api/v1/{"temperature": float,"humidity": float}

## Go tips:
    - Visual Code packages:
       - Remote - WSL
       - Go (VS will ask you other Go dependencies to be installed at first run)
       - Remote - Containers (not used by now)

    - Remove Go autoformat on save (mainly by replacing spaces by tabs!):
       - in ~/.vscode-server/data/Machine/setting.json 
       ```     
       "[go]": {
        "editor.formatOnSave": false 
        }
        ```

    - run the web server (local test with Postman for instance):
        go run .
    - run tests with coverage: 
        go test -cover

## Useful links:

Ofcourse, nothing new in this repository, only code adaptation from various links.
- how to [unit test net/http](https://golang.org/src/net/http/httptest/example_test.go)
- how to unit test [http handlers](https://blog.questionable.services/article/testing-http-handlers-go/)
- how to unit test [gorilla/mux](https://stackoverflow.com/questions/34435185/unit-testing-for-functions-that-use-gorilla-mux-url-parameters)
- nice tutorial to build [REST api in go](https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj)

