/*
从导出的工程文件加载
*/
package main

import (
	"fmt"
	"os"

	"./m"
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
	. "github.com/magicsea/behavior3go/loader"
)

func main() {
	fmt.Println(os.Args[1])
	projectConfig, ok := LoadProjectCfg(os.Args[1])
	if !ok {
		fmt.Println("LoadTreeCfg err")
		return
	}

	//自定义节点注册
	maps := b3.NewRegisterStructMaps()
	// maps.Register("Log", new(LogTest))
	maps.Register("shs", new(m.SomeoneHasSomething))
	maps.Register("isdate", new(m.IsDate))
	maps.Register("buy", new(m.Buy))
	maps.Register("find", new(m.Find))
	maps.Register("go", new(m.Go))
	maps.Register("give", new(m.Give))
	maps.Register("playball", new(m.PlayBall))
	maps.Register("take", new(m.Take))
	maps.Register("kfc", new(m.KFC))
	maps.Register("randN", new(m.RandN))

	var firstTree *BehaviorTree
	//载入
	for _, v := range projectConfig.Trees {
		tree := CreateBevTreeFromConfig(&v, maps)
		tree.Print()
		if firstTree == nil {
			firstTree = tree
		}
	}

	//输入板
	board := NewBlackboard()
	// board.Set("myself_flower", true, "", "")
	board.Set("month", 2, "", "")
	board.Set("day", 14, "", "")
	//循环每一帧
	for i := 0; i < 5; i++ {
		firstTree.Tick(i, board)
	}
}
