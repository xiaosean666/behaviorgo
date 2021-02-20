package share

import (
	"fmt"
	b3 "github.com/huilong-cn/behavior3go"
	//. "github.com/huilong-cn/behavior3go/actions"
	//. "github.com/huilong-cn/behavior3go/composites"
	. "github.com/huilong-cn/behavior3go/config"
	. "github.com/huilong-cn/behavior3go/core"
	//. "github.com/huilong-cn/behavior3go/decorators"
)

//自定义action节点
type LogTest struct {
	Action
	info string
}

func (this *LogTest) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *LogTest) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest:",tick.GetLastSubTree(), this.info)
	return b3.SUCCESS
}
