package models

import "errors"

type AttackType int

const (
	_                  AttackType = iota
	AttackTypePhysical            //物理攻击方式
	AttackTypeMagic               //魔法攻击方式
)

var (
	ErrorCurrentFighterIsDead = errors.New("current fighter is dead")
	ErrorTargetFighterIsDead  = errors.New("target fighter is dead")
)

//斗士基本属性
type Fighter struct {
	Life         int32  //生命值
	Mana         int32  //魔法值
	NormalAttack Attack //普通攻击
}
