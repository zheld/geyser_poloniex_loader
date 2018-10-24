package pair

import (
    "trades"
    "net/url"
    "fmt"
    "time"
    "core"
)

var PairList = map[string]*Pair{}
var start_service = time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC).Unix()

type TradeHistory struct {
    GlobalTradeID int64   `json:"globalTradeID"`
    TradeID       int64   `json:"tradeID"`
    Date          string  `json:"date"`
    Type          string  `json:"type"`
    Rate          float64 `json:"rate,string"`
    Amount        float64 `json:"amount,string"`
    Total         float64 `json:"total,string"`
}

type Pair struct {
    DBI_Pair
    last int64
}

func (this *Pair) GetTradeHistory(start int64, end int64) ([]TradeHistory, error) {
    vals := url.Values{}
    vals.Set("currencyPair", this.Name)

    vals.Set("start", core.IToStr(start))
    vals.Set("end", core.IToStr(end))

    resp := []TradeHistory{}
    path := fmt.Sprintf("%s/public?command=returnTradeHistory&%s", "https://poloniex.com", vals.Encode())

    err := core.HTTPSendGetRequestTimeout(path, true, &resp, time.Second*10)

    if err != nil {
        return nil, err
    }
    return resp, nil
}

func (this *Pair) GetLastTradeHistory() ([]TradeHistory, error) {
    start := this.GetLast()
    end := start + 1200
    history, err := this.GetTradeHistory(start, end)
    if err != nil {
        return nil, err
    }
    if len(history) == 0 {
        if !time.Now().Before(time.Unix(end, 0)) {
            this.last = end
        }
    } else {
        for _, trade := range history {
            tm, err := core.TimeStrWithoutNanosecToTime(trade.Date)
            if err != nil {
                continue
            }
            utm := tm.Unix()
            if this.last < utm {
                this.last = utm
            }
        }
        this.last++
    }
    return history, nil
}

func (this *Pair) GetLast() int64 {
    if this.last > 0 {
        return this.last
    }
    last, err := trades.GetLastDateByPair(this.Id)
    if err != nil || last.IsZero() {
        core.PublishInfo(fmt.Sprintf("No data by current pair, start by: %v, name: %v", start_service, this.Name))
        this.last = start_service
    } else {
        this.last = last.Unix() + 1
    }
    return this.last
}

func NewPair(name string) (*Pair, error) {
    pr := &Pair{}
    ins, err := InsertReadReturnInstance(name)
    pr.DBI_Pair = ins
    return pr, err
}

//!api get_pair_id
func GetPairID(name string) (id int, err error) {
    pr, err := GetPairByName(name)
    return pr.Id, err
}

func GetPairByName(name string) (*Pair, error) {
    if pr, ok := PairList[name]; ok {
        return pr, nil
    }
    pr, err := NewPair(name)
    PairList[name] = pr
    return pr, err

}
