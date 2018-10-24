package updater

import (
    "pair"
    "core"
)

var activeTradesUpdater = false

func UpdateOrders() {
    go updateOrders()
}

func UpdateTrades() {
    go updateTrades()
}

type updateList struct {
    ls      []*pair.Pair
    current int
}

func (this *updateList) Next() (res *pair.Pair) {
    if len(this.ls) == 0 {
        return nil
    }
    res = this.ls[this.current]
    this.current++
    if this.current >= len(this.ls) {
        this.current = 0
    }
    return res
}

func (this *updateList) Add(pr *pair.Pair) {
    this.ls = append(this.ls, pr)
}

var usdPairList = updateList{}
var otherPairList = updateList{}

type pairProducer struct {
    usdPairList   updateList
    otherPairList updateList
    current       int
}

func (this *pairProducer) init() {
    this.usdPairList = updateList{}
    this.otherPairList = updateList{}

    for _, pr := range pair.PairList {
        items := core.StrSplit(pr.Name, "_")
        basename := items[0]
        if basename == "USDT" {
            this.usdPairList.Add(pr)
        } else {
            this.otherPairList.Add(pr)
        }
    }
}

func (this *pairProducer) GetUSDpair() *pair.Pair {
    return this.usdPairList.Next()
}

func (this *pairProducer) Get() *pair.Pair {
    if this.current == 0 {
        this.current = 1
        return this.usdPairList.Next()
    } else {
        this.current = 0
        return this.otherPairList.Next()
    }
}
