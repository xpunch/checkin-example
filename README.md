# Check In Example
```
Use different redis data structure to implement check in service.
Check in by user id, return continuous checkin days.
```
## Prepare

```
protoc --go_out=proto --micro_out=proto proto/check_in.proto
```

## Server

### Bitmap

```
cd server/bitmap
go run .
```

### Set

```
cd server/set
go run .
```

### ZSet

```
cd server/zset
go run .
```

## Client

```
cd client
go run . --uid 10000
```
