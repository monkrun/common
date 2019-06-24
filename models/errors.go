package models

import "errors"

var (
	ErrorAttackCoolingTime = errors.New("attack failed.skill is cooling")
	ErrorTargetIsDead      = errors.New("target is already dead")
	ErrorCurrentIsDead     = errors.New("current fighter is already dead")
)
