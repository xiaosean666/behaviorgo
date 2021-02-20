package m

import (
	"math/rand"
	"time"

	b3 "github.com/magicsea/behavior3go"
	//. "github.com/magicsea/behavior3go/actions"
	//. "github.com/magicsea/behavior3go/composites"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
	//. "github.com/magicsea/behavior3go/decorators"
)

// //自定义action节点
// type SetValue struct {
// 	Action
// 	value int
// 	key   string
// }

// func (this *SetValue) Initialize(setting *BTNodeCfg) {
// 	this.Action.Initialize(setting)
// 	this.value = setting.GetPropertyAsInt("value")
// 	this.key = setting.GetPropertyAsString("key")
// }

// func (this *SetValue) OnTick(tick *Tick) b3.Status {
// 	tick.Blackboard.SetMem(this.key, this.value)
// 	return b3.SUCCESS
// }

//SomeoneHasSomethine 自定义condition节点
type SomeoneHasSomething struct {
	Condition
	someone   string
	something string
}

func (this *SomeoneHasSomething) Initialize(setting *BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.someone = setting.GetPropertyAsString("someone")
	this.something = setting.GetPropertyAsString("something")
}

func (this *SomeoneHasSomething) OnTick(tick *Tick) b3.Status {
	v := tick.Blackboard.GetBool(this.someone+"_"+this.something, "", "")
	if v {
		return b3.SUCCESS
	}
	return b3.FAILURE
}

//IsDate 自定义condition节点
type IsDate struct {
	Condition
	month int
	day   int
}

func (this *IsDate) Initialize(setting *BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.month = setting.GetPropertyAsInt("month")
	this.day = setting.GetPropertyAsInt("day")
}

func (this *IsDate) OnTick(tick *Tick) b3.Status {
	m := tick.Blackboard.GetInt("month", "", "")
	d := tick.Blackboard.GetInt("day", "", "")
	if m == this.month && d == this.day {
		return b3.SUCCESS
	}
	return b3.FAILURE
}

//RandN 自定义condition节点
type RandN struct {
	Condition
	weight int
	val    int
}

func (this *RandN) Initialize(setting *BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.weight = setting.GetPropertyAsInt("weight")
	this.val = setting.GetPropertyAsInt("val")
}

var gRander = rand.New(rand.NewSource(time.Now().UnixNano()))

func (this *RandN) OnTick(tick *Tick) b3.Status {
	if gRander.Intn(this.weight) < this.val {
		return b3.SUCCESS
	}
	return b3.FAILURE
}
