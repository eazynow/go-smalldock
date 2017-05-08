# go-smalldock
Just a test of making a small API service with minimal docker resources

## Requirements
You'll want to have Docker installed on your machine, as well as having cloned this repo.

## Building

```
$ make build-linux
```

## Running

To run interactively in command line:

```
$ docker run -p 127.0.0.1:8081:8080 -it  smalldock-linux
```


## Testing

When running, just run the following from another terminal

```
$ curl http://localhost:8081/health
OK at 2017-05-08 19:01:20.657432957 +0000 UTC
```