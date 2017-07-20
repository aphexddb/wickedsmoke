Build for ARM / Pi

```bash
SET GOARM=6
SET GOARCH=arm
SET GOOS=linux
go build -o bin/wickedsmoke-arm  analogpi.go api.go client.go cook.go hardware.go hub.go ip.go main.go
```