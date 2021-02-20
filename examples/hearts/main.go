/*
从导出的工程文件加载
*/
package main

import (
	"fmt"
	"os"
	"sync"

	"./htree"
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	core "github.com/magicsea/behavior3go/core"
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
	maps.Register("cur_player_idx", new(htree.CurPlayerIdx))
	maps.Register("same_card_bigger", new(htree.SameCardBigger))
	maps.Register("same_card_smallest", new(htree.SameCardSmallest))
	maps.Register("is_partner_biggest", new(htree.IsPartnerBiggest))
	maps.Register("other_card_nth", new(htree.OtherCardNth))
	maps.Register("other_card_small_nth", new(htree.OtherCardSmallNth))
	maps.Register("master_card_nth", new(htree.MasterCardNth))
	maps.Register("randN", new(htree.RandN))
	maps.Register("master_card_count", new(htree.MasterCardNth))
	maps.Register("same_color_count", new(htree.SamelColorCount))
	maps.Register("same_color_nth", new(htree.SamelColorNth))

	var mainTree *core.BehaviorTree
	//载入
	for _, v := range projectConfig.Trees {
		tree := CreateBevTreeFromConfig(&v, maps)
		fmt.Println(v.ID, v.Title)
		tree.Print()
		mapTreesByID.Store(v.ID, tree)
		if v.Title == "hearts-main" && mainTree == nil {
			mainTree = tree
		}
	}

	//输入板
	board := core.NewBlackboard()
	playerCards := [][]int{
		[]int{},
		[]int{},
		[]int{},
		[]int{},
	}
	playerIDs := []int64{1, 2, 3, 4}
	cards := []int{}
	curPlayer := int64(1)
	board.Set(htree.KeyMasterColor, 0x01, "", "")
	board.Set(htree.KeyPlayerCards, playerCards, "", "")
	board.Set(htree.KeyPlayerIDs, playerIDs, "", "")
	board.Set(htree.KeyCards, cards, "", "")
	board.Set(htree.KeyCurPlayer, curPlayer, "", "")
	//循环每一帧
	for i := 0; i < 1; i++ {
		mainTree.Tick(i, board)
	}
}

//所有的树管理
var mapTreesByID = sync.Map{}

func init() {
	//获取子树的方法
	core.SetSubTreeLoadFunc(func(id string) *core.BehaviorTree {
		println("==>load subtree:", id)
		t, ok := mapTreesByID.Load(id)
		if ok {
			return t.(*core.BehaviorTree)
		}
		return nil
	})
}
