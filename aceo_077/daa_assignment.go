package main

import (
	"fmt"
)

func main() {
	newToy := NewToy(SpiderMan{}, HalfFly{}, FullShoot{}, Mortal{})
	newToy.PerformDialogue()
	newToy.PerformFly()
	fmt.Println(" ")
	newToy.SetSuperHero(SuperMan{})
	newToy.SetHealthType(Immortal{})
	newToy.PerformDialogue()
	newToy.HealthType()
	fmt.Println(" ")
	newToy.SetSuperHero(BatMan{})
	newToy.PerformDialogue()
	newToy.SetFlyBeh(NoFly{})
	newToy.SetShootBeh(HalfShoot{})
	newToy.SetHealthType(Mortal{})
	newToy.PerformFly()
	newToy.MakeShoot()
	newToy.HealthType()
}

type DialogueReciter interface {
	Recite()
}

type Fly interface {
	PerformFly()
}

type Shoot interface {
	MakeShoot()
}

type Health interface {
	HealthType()
}
type Mortal struct {
}

func (m Mortal) HealthType() {
	fmt.Println("I am mortal :(")
}

type Immortal struct {
}

func (im Immortal) HealthType() {
	fmt.Println("I am immortal hahaha")
}

type FullShoot struct {
}

func (fs FullShoot) MakeShoot() {
	fmt.Println("Shoot with super abilities")
}

type HalfShoot struct {
}

func (hs HalfShoot) MakeShoot() {
	fmt.Println("Shoot with guns")
}

type NoShoot struct {
}

func (ns NoShoot) MakeShoot() {
	fmt.Println("Cannot shoot")
}

type FullFly struct {
}

func (ff FullFly) PerformFly() {
	fmt.Println("Fly by himself")
}

type HalfFly struct {
}

func (hf HalfFly) PerformFly() {
	fmt.Println("Fly with assist")
}

type NoFly struct {
}

func (nf NoFly) PerformFly() {
	fmt.Println("Cannot fly")
}

type SpiderMan struct{}

func (spm SpiderMan) Recite() {
	fmt.Println("No Man Can Win Every Battle, " +
		"But No Man Should Fall Without A Struggle")
}

type SuperMan struct{}

func (sum SuperMan) Recite() {
	fmt.Println("There is a superhero in all of us, " +
		"we just need the courage to put on the cape")
}

type BatMan struct{}

func (sum BatMan) Recite() {
	fmt.Println("It's not who I am underneath, " +
		"but what I do that defines me!")
}

type toy struct {
	DialogueReciter DialogueReciter
	FlyOption       Fly
	ShootOption     Shoot
	HealthOption    Health
}

func NewToy(dr DialogueReciter, fly Fly, sh Shoot, h Health) *toy {
	return &toy{
		DialogueReciter: dr,
		FlyOption:       fly,
		ShootOption:     sh,
		HealthOption:    h,
	}
}

func (t *toy) PerformDialogue() {
	t.DialogueReciter.Recite()
}

func (t *toy) SetSuperHero(dr DialogueReciter) {
	t.DialogueReciter = dr
}

func (t *toy) PerformFly() {
	t.FlyOption.PerformFly()
}

func (t *toy) SetFlyBeh(fly Fly) {
	t.FlyOption = fly
}

func (t *toy) MakeShoot() {
	t.ShootOption.MakeShoot()
}

func (t *toy) SetShootBeh(sh Shoot) {
	t.ShootOption = sh
}

func (t *toy) HealthType() {
	t.HealthOption.HealthType()
}

func (t *toy) SetHealthType(sho Health) {
	t.HealthOption = sho
}
