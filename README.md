# Redis Cache Client

Redis Cache is the [Redis](https://redis.io/) cache client for [Casbin](https://github.com/casbin/casbin). With this library, Casbin can cache results in Redis.

## Installation

    go get github.com/casbin/redis-adapter/v2

## Simple Example

```go
package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/cache-redis/v2"
	"github.com/go-redis/redis"
)

func main() {
	// Initialize Redis client
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	// Initialize a Redis cache and use it in a Casbin enforcer:
	с := cacheredis.NewRedisCache(client)
	e, _ := casbin.NewCachedEnforcer()

	e.EnableCache(true)
	e.SetCache(с)
}
```

## Need improvements
* Add logger
* Add error handling for all operations
* Add tests

## Getting Help

- [Casbin](https://github.com/casbin/casbin)

## License

This project is under Apache 2.0 License. See the [LICENSE](LICENSE) file for the full license text.
