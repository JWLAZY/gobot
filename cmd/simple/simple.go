package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jwlazy/gobot/internal/model"
	"github.com/stianeikeland/go-rpio/v4"
)

/**
 * 驱动一个电机
 */
func main() {
	//4, 17, 18
	fmt.Println(os.Args)
	c, _ := strconv.Atoi(os.Args[1])
	a, _ := strconv.Atoi(os.Args[2])
	d, _ := strconv.Atoi(os.Args[3])
	s, _ := strconv.Atoi(os.Args[4])
	car := model.NewCar(c, a, d)

	defer rpio.Close()

	fmt.Printf("go, 5 second\n")
	go car.Go(s)
	time.Sleep(5 * time.Second)

	fmt.Printf("stop, 2 second\n")
	car.Stop()
	time.Sleep(2 * time.Second)

	fmt.Printf("back, 5 second\n")
	go car.Back(s)
	time.Sleep(5 * time.Second)

	car.Stop()

	rpio.Close()
	os.Exit(0)
}
