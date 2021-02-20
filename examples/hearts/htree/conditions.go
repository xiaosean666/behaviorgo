package htree

import (
	"fmt"

	b3 "github.com/huilong-cn/behavior3go"
	//. "github.com/huilong-cn/behavior3go/actions"
	//. "github.com/huilong-cn/behavior3go/composites"
	"github.com/huilong-cn/behavior3go/config"
	"github.com/huilong-cn/behavior3go/core"
	//. "github.com/huilong-cn/behavior3go/decorators"
)

//SamelColorCount same_color_count 自定义condition节点
type SamelColorCount struct {
	core.Condition
	count int
}

func (this *SamelColorCount) Initialize(setting *config.BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.count = setting.GetPropertyAsInt("count")
}

func (this *SamelColorCount) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("node.name:", this.GetName(), "title:", this.GetTitle())
	// v := tick.Blackboard.GetBool(this.someone+"_"+this.something, "", "")
	// if v {
	// 	return b3.SUCCESS
	// }
	return b3.FAILURE
}

//MasterCardCount 自定义action节点
type MasterCardCount struct {
	core.Condition
	count int
}

func (this *MasterCardCount) Initialize(setting *config.BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.count = setting.GetPropertyAsInt("count")
}

func (this *MasterCardCount) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("node.name:", this.GetName(), "title:", this.GetTitle())
	// v := tick.Blackboard.GetBool(this.someone+"_"+this.something, "", "")
	// if v {
	// 	return b3.SUCCESS
	// }
	return b3.FAILURE
}

//RandN 自定义action节点
type RandN struct {
	core.Condition
	weight int
	val    int
}

func (this *RandN) Initialize(setting *config.BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.weight = setting.GetPropertyAsInt("weight")
	this.val = setting.GetPropertyAsInt("val")
}

func (this *RandN) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("node.name:", this.GetName(), "title:", this.GetTitle())
	if random(this.weight) < this.val {
		return b3.SUCCESS
	}
	return b3.FAILURE
}

//IsPartnerBiggest is_partner_biggest
type IsPartnerBiggest struct {
	core.Condition
}

func (this *IsPartnerBiggest) Initialize(setting *config.BTNodeCfg) {
	this.Condition.Initialize(setting)
}

func (this *IsPartnerBiggest) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("node.name:", this.GetName(), "title:", this.GetTitle())
	// v := tick.Blackboard.GetBool(this.someone+"_"+this.something, "", "")
	// if v {
	// 	return b3.SUCCESS
	// }
	return b3.FAILURE
}

//SameCardBigger same_card_bigger 是否有同花色比牌堆大的牌
type SameCardBigger struct {
	core.Condition
}

func (this *SameCardBigger) Initialize(setting *config.BTNodeCfg) {
	this.Condition.Initialize(setting)
}

func (this *SameCardBigger) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("node.name:", this.GetName(), "title:", this.GetTitle())
	// v := tick.Blackboard.GetBool(this.someone+"_"+this.something, "", "")
	// if v {
	// 	return b3.SUCCESS
	// }
	return b3.FAILURE
}

//CurPlayerIdx cur_player_idx 是否有同花色比牌堆大的牌
type CurPlayerIdx struct {
	core.Condition
	cardCount int
}

func (this *CurPlayerIdx) Initialize(setting *config.BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.cardCount = setting.GetPropertyAsInt("card_count")
}

func (this *CurPlayerIdx) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("node.name:", this.GetName(), "title:", this.GetTitle())
	v := tick.Blackboard.Get(KeyCards, "", "")
	if cards, ok := v.([]int); ok {
		if len(cards) == this.cardCount {
			return b3.SUCCESS
		}
		return b3.FAILURE
	}
	return b3.FAILURE
}
