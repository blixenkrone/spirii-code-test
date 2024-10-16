## Spirii code challenge

## How to run
`$ go run ./cmd/main.go`

or


`$ docker build --pull --rm -f "Dockerfile" -t blixenkrone/spirii:latest .`

`$ docker run --rm -d -p 8080:8080 blixenkrone/spirii:latest`

## How to hit endpoints
Use either httpie or curl and target localhost ie.:

`http get localhost:8080/ping`

`curl -V GET localhost:8080/ping`

Valid routes are:
`/ping`

`/v1/chargers/{id}` // {id} is a number from 1-3

`/v1/top-consumers`



## How to test 
$ go test ./... -count 1 -short

## How to stop
Just sig kill it (once for soft, twice for force)

## Notes
Focus was put on concurrent data streams, HTTP API and tests.
Error handling and logging is pretty standard.
I didn't find the API design quite RESTful, so I only implemented the first route.

Done in ~1 hour since I don't principally allow myself to allocate more time for unpaid coding tests.
I started doing a PostgresQL/Docker implementation, but I didn't have time to finish it, so ended up with a cache/in-memory implementation.
No AI used - not that I won't use it in my day-to-day, but for the purpose of showing my thinking.

Some "panic"s in there, but do note I never use them unless something truly is unrecoverable.
Git commit logs are not reflective of how I usually work.
