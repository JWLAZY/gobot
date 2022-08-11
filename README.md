# gobot
use golang to drive raspberry

# demo

```
    CGO_ENABLED=0 GOOS=linux GOARCH=arm go build cmd/simple/simple.go
    ./simple pin1 pin2 pwm speed(50+)
```