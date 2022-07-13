package main

import (
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

/**
 * 驱动一个电机
 */
func main() {

}

type Motor struct {
	CorrectPin rpio.Pin
	AntiPin    rpio.Pin
	PwmPin     rpio.Pin
}

func (motor *Motor) Init() {
	motor.CorrectPin.Mode(rpio.Output)
	motor.AntiPin.Mode(rpio.Output)
	motor.PwmPin.Mode(rpio.Output)
	motor.CorrectPin.High()
	motor.AntiPin.Low()
	rpio.SetFreq(motor.PwmPin, 1000)
	motor.PwmPin.DutyCycle(0, 20)
}

func (motor *Motor) Go() {
	motor.PwmPin.DutyCycle(10, 20)
}

func runOneMotor() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	motor := Motor{CorrectPin: 10, AntiPin: 11, PwmPin: 12}
	motor.Init()

	motor.Go()
	time.Sleep(5 * time.Second)
	os.Exit(0)
}
