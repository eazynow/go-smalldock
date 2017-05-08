# go-smalldock
Just a test of making a small API service with minimal docker resources

## Requirements
You'll want to have Docker installed on your machine, as well as having cloned this repo.

## Usage

### Building

```
$ make build-linux
```

### Running

To run in command line:

```
$ docker run -p 8080:8080 -d smalldock-linux
39f9ae32f1c84c7588498cd109654a553ff901c15c71f67f4d96334a181c3a5c
```

### Stopping

First check the container ID

``` 
$ docker ps
CONTAINER ID 	IMAGE
39f9ae....

$ docker stop 39
39
```

## Testing

When running, just run the following from another terminal

```
$ curl http://localhost:8080/health
OK at 2017-05-08 19:01:20.657432957 +0000 UTC
```