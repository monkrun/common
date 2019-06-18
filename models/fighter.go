package models

type AttackType int

const (
	_                  AttackType = iota
	AttackTypePhysical            //物理攻击方式
	AttackTypeMagic               //魔法攻击方式
)

//斗士基本属性
type Fighter struct {
	Life           int32      //生命值
	Mana           int32      //魔法值
	AttackInterval int64      //物理攻击时间间隔，单位毫秒
	Attack         int32      //攻击基准值
	AttackType     AttackType //攻击方式
}
