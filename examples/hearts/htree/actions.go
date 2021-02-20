package htree

import (
	"fmt"

	b3 "github.com/huilong-cn/behavior3go"
	"github.com/huilong-cn/behavior3go/config"
	"github.com/huilong-cn/behavior3go/core"
)

//Buy 自定义action节点
type SamelColorNth struct {
	core.Action
	// something string
}

func (this *SamelColorNth) Initialize(setting *config.BTNodeCfg) {
	this.Action.Initialize(setting)
	// this.something = setting.GetPropertyAsString("something")
}

func (this *SamelColorNth) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("node.name:", this.GetName(), "title:", this.GetTitle())
	// tick.Blackboard.Set("myself_"+this.something, true, "", "")
	// tick.Blackboard.Remove("myself_money")
	return b3.SUCCESS
}

//MasterCardNth master_card_nth
type MasterCardNth struct {
	core.Action
	rank int
}

func (this *MasterCardNth) Initialize(setting *config.BTNodeCfg) {
	this.Action.Initialize(setting)
	// this.rank = setting.GetPropertyAsInt("rank")
}

func (this *MasterCardNth) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("node.name:", this.GetName(), "title:", this.GetTitle())
	// tick.Blackboard.Set("myself_"+this.something, true, "", "")
	// tick.Blackboard.Remove("myself_money")
	return b3.SUCCESS
}

//OtherCardSmallNth other_card_small_nth
type OtherCardSmallNth struct {
	core.Action
	rank int
}

func (this *OtherCardSmallNth) Initialize(setting *config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.rank = setting.GetPropertyAsInt("rank")
}

func (this *OtherCardSmallNth) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("node.name:", this.GetName(), "title:", this.GetTitle())
	// tick.Blackboard.Set("myself_"+this.something, true, "", "")
	// tick.Blackboard.Remove("myself_money")
	return b3.SUCCESS
}

//OtherCardNth other_card_nth
type OtherCardNth struct {
	core.Action
	rank int
}

func (this *OtherCardNth) Initialize(setting *config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.rank = setting.GetPropertyAsInt("rank")
}

func (this *OtherCardNth) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("node.name:", this.GetName(), "title:", this.GetTitle())
	// tick.Blackboard.Set("myself_"+this.something, true, "", "")
	// tick.Blackboard.Remove("myself_money")
	return b3.SUCCESS
}

//SameCardSmallest same_card_smallest
type SameCardSmallest struct {
	core.Action
}

func (this *SameCardSmallest) Initialize(setting *config.BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *SameCardSmallest) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("node.name:", this.GetName(), "title:", this.GetTitle())
	// tick.Blackboard.Set("myself_"+this.something, true, "", "")
	// tick.Blackboard.Remove("myself_money")
	return b3.SUCCESS
}
