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

## API

### V1
    - /api/v1/data/{"temperature": float,"humidity": float}
    - to validate openapi structure, there is (swagger)[https://apitools.dev/swagger-parser/online/] web site.

    - api generated automatically from docker image dowmloaded from: https://hub.docker.com/r/openapitools/openapi-generator-cli/
    - using docker command: 

```docker
$> sudo docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i local/openapi.yaml -g go -o /local/openapi
$> chown -R OCondat:OCondat .
$> mv openapi ../
$> cd ../openapi
$> git init
$> git remote add origin git@github.com:Olivier-Condat/openapi-go-first-api-test.git
$> git push -u origin master
```
In openapi workspace, in go.mod file, replace:

```golang
module github.com/GIT_USER_ID/GIT_REPO_ID => module github.com/Olivier-Condat/openapi-go-first-api-test
```

Commit and push then. 

In the project that will use the openapi dependency =>
```golang
 go install github.com/Olivier-Condat/openapi-go-first-api-test
``` 
Local go.mod will be updated => 

```golang
github.com/Olivier-Condat/openapi-go-first-api-test v0.0.0-20210414163125-5dc39a9a7eb3
```


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

