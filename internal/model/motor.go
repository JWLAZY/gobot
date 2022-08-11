package model

import (
	"github.com/stianeikeland/go-rpio/v4"
)

type Motor struct {
	CorrectPin rpio.Pin
	AntiPin    rpio.Pin
	PwmPin     rpio.Pin
}

func NewMotor(corrent int, anti int, pwm int) *Motor {
	motor := &Motor{CorrectPin: rpio.Pin(corrent), AntiPin: rpio.Pin(anti), PwmPin: rpio.Pin(pwm)}
	motor.Inits()
	return motor
}

func (motor *Motor) Inits() {
	motor.CorrectPin.Mode(rpio.Output)
	motor.AntiPin.Mode(rpio.Output)
	motor.PwmPin.Mode(rpio.Pwm)
	motor.CorrectPin.Low()
	motor.AntiPin.Low()
	motor.PwmPin.Freq(64000)
	motor.PwmPin.DutyCycleWithPwmMode(0, 100, true)
}

func (motor *Motor) Go(speed int) {
	motor.CorrectPin.High()
	motor.AntiPin.Low()
	motor.PwmPin.DutyCycleWithPwmMode(uint32(speed), 100, true)
}

func (motor *Motor) Back(speed int) {
	motor.AntiPin.High()
	motor.CorrectPin.Low()
	motor.PwmPin.DutyCycleWithPwmMode(uint32(speed), 100, true)
}

func (motor *Motor) Stop() {
	motor.CorrectPin.Low()
	motor.AntiPin.Low()
	motor.PwmPin.DutyCycleWithPwmMode(0, 100, true)
}
