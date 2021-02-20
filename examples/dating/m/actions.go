package m

import (
	"fmt"

	b3 "github.com/huilong-cn/behavior3go"

	//. "github.com/huilong-cn/behavior3go/actions"
	//. "github.com/huilong-cn/behavior3go/composites"
	. "github.com/huilong-cn/behavior3go/config"
	. "github.com/huilong-cn/behavior3go/core"
	//. "github.com/huilong-cn/behavior3go/decorators"
)

//Buy 自定义action节点
type Buy struct {
	Action
	something string
}

func (this *Buy) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.something = setting.GetPropertyAsString("something")
}

func (this *Buy) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest: buy", tick.GetLastSubTree(), this.something)
	tick.Blackboard.Set("myself_"+this.something, true, "", "")
	tick.Blackboard.Remove("myself_money")
	return b3.SUCCESS
}

//Go 自定义action节点
type Go struct {
	Action
	place string
}

func (this *Go) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.place = setting.GetPropertyAsString("place")
}

func (this *Go) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest: go ", tick.GetLastSubTree(), this.place)
	return b3.SUCCESS
}

//Give 自定义action节点
type Give struct {
	Action
	something string
}

func (this *Give) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.something = setting.GetPropertyAsString("something")
}

func (this *Give) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest: give ", tick.GetLastSubTree(), this.something)
	tick.Blackboard.Set("girl_"+this.something, true, "", "")
	tick.Blackboard.Remove("myself_" + this.something)
	return b3.SUCCESS
}

//Find 自定义action节点
type Find struct {
	Action
	someone string
}

func (this *Find) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.someone = setting.GetPropertyAsString("someone")
}

func (this *Find) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest: find ", tick.GetLastSubTree(), this.someone)
	return b3.SUCCESS
}

//PlayBall 自定义action节点
type PlayBall struct {
	Action
}

func (this *PlayBall) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *PlayBall) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest: playball ", tick.GetLastSubTree())
	return b3.SUCCESS
}

//Take 自定义action节点
type Take struct {
	Action
	something string
}

func (this *Take) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.something = setting.GetPropertyAsString("something")
}

func (this *Take) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest: take ", tick.GetLastSubTree(), this.something)
	tick.Blackboard.Set("myself_"+this.something, true, "", "")
	return b3.SUCCESS
}

//KFC 自定义action节点
type KFC struct {
	Action
}

func (this *KFC) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *KFC) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest: kfc ", tick.GetLastSubTree())
	return b3.SUCCESS
}
