package models

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Attack struct {
	MaxHurt           int32
	MinHurt           int32
	CoolingTime       time.Duration
	GlobalCoolingTime time.Duration
	Lock              int32
	AttackType        AttackType
}

func (a Attack) lock() bool {
	return atomic.CompareAndSwapInt32(&a.Lock, 0, 1)
}

func (a Attack) unlock() bool {
	return atomic.CompareAndSwapInt32(&a.Lock, 1, 0)
}

func (a Attack) Attack(target *Fighter, globalLock sync.Locker) error {
	if target.Life <= 0 {
		return ErrorTargetIsDead
	}
	if a.lock() {
		defer func() {
			go func() {
				time.Sleep(a.CoolingTime)
				a.unlock()
			}()
		}()
	} else {
		return ErrorAttackCoolingTime
	}
	globalLock.Lock()
	defer func() {
		go func() {
			time.Sleep(a.GlobalCoolingTime)
			globalLock.Unlock()
		}()
	}()
	target.Life -= (rand.Int31n(a.MaxHurt-a.MinHurt) + a.MinHurt)
	return nil
}

func (a Attack) CancleAttack() error {
	return nil
}
