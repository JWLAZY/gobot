package model

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

type Car struct {
	leftMotor1 Motor
}

func NewCar(c, a, d int) *Car {
	err := rpio.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	car := &Car{leftMotor1: *NewMotor(c, a, d)}
	return car
}

func (car *Car) Go(speed int) {
	car.leftMotor1.Go(speed)
}

func (car *Car) Stop() {
	car.leftMotor1.Stop()
}

func (car *Car) Back(speed int) {
	car.leftMotor1.Back(speed)
}
