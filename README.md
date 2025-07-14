# CLI using

## Getting weather in user location, default format (default is 4)
```bash
go run .
```
or just run binary file

## Getting weather in specified city with specified format
```bash
go run . --city=London --format=[1 <= n <= 5]
```

## Testing

```bash
go test -v ./geo/ && go test -v ./weather/
```
