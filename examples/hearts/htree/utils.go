package htree

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/magicsea/behavior3go/core"
)

var gRander = rand.New(rand.NewSource(time.Now().UnixNano()))
var rmux = new(sync.Mutex)

func random(weight int) int {
	rmux.Lock()
	val := gRander.Intn(weight)
	rmux.Unlock()
	return val
}

func curPlayerCards(blackbord *core.Blackboard) []int {
	val := blackbord.Get(KeyPlayerCards, "", "")
	playercards, ok := val.([][]int)
	if !ok {
		panic(fmt.Sprintf("key:%s bad val %v", KeyCards, val))
	}
	return playercards[curPlayer(blackbord)]
}
func curPlayer(blackbord *core.Blackboard) int {
	val := blackbord.Get(KeyCards, "", "")
	// if !ok {
	// 	panic(KeyCards + "isnt found")
	// }
	cards, ok := val.([]int)
	if !ok {
		panic(fmt.Sprintf("key:%s bad val %v", KeyCards, val))
	}
	return len(cards)
}

func masterColorCards(cards []int, masterColor int) []int {
	return []int{}
}

func sortCardsByPower(cards []int, masterColor int) []int {
	return []int{}
}
