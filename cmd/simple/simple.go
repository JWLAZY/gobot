package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

/**
 * 驱动一个电机
 */
func main() {
	runOneMotor()
}

type Motor struct {
	CorrectPin rpio.Pin
	AntiPin    rpio.Pin
	PwmPin     rpio.Pin
}

func (motor *Motor) Inits() {
	motor.CorrectPin.Mode(rpio.Output)
	motor.AntiPin.Mode(rpio.Output)
	motor.PwmPin.Mode(rpio.Pwm)
	motor.CorrectPin.Low()
	motor.AntiPin.Low()
	motor.PwmPin.Pwm()
	// rpio.StartPwm()
	rpio.SetFreq(motor.PwmPin, 1)
	motor.PwmPin.DutyCycleWithPwmMode(0, 10, true)
}

func (motor *Motor) Go() {
	// motor.PwmPin.Freq(10)
	motor.CorrectPin.High()
	motor.PwmPin.DutyCycleWithPwmMode(10, 10, true)
}

func (motor *Motor) Stop() {
	// motor.PwmPin.Freq(0)
	motor.CorrectPin.Low()
	motor.PwmPin.DutyCycleWithPwmMode(0, 10, true)
}

func runOneMotor() {
	err := rpio.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	motor := Motor{CorrectPin: rpio.Pin(12), AntiPin: rpio.Pin(6), PwmPin: rpio.Pin(5)}
	motor.Inits()

	for i := 0; i < 5; i++ {
		fmt.Printf("wait go, %d\n", i)
		time.Sleep(1 * time.Second)
	}
	go motor.Go()
	for i := 0; i < 5; i++ {
		fmt.Printf("wait stop, %d\n", i)
		time.Sleep(1 * time.Second)
	}
	go motor.Stop()
	time.Sleep(4 * time.Second)
	rpio.Close()
	os.Exit(0)
}
