# statsd-dumper
Listens to incoming messages and dumps them to `STDOUT`. Useful to debug/monitor statsd metrics on a local environment.

## Usage
### Running in a container

```docker run -it --rm -p 8125:8125/udp jeanreis/statsd-dumper```

### Build and run from source

```go run .```

## Config
The port to listen on can be configured with the `STATSD_PORT` env var.