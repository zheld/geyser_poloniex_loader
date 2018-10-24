package updater

import (
    "core"
    "fmt"
    "net/url"
    "pair"
    "strconv"
    "strings"
    "time"
)

var order_list_len = "1"
var polo_api_url = "https://poloniex.com"
var err_count int
var total_err_count = 20

type PoloniexOrderbookResponse struct {
    Asks     [][]interface{} `json:"asks"`
    Bids     [][]interface{} `json:"bids"`
    IsFrozen string          `json:"isFrozen"`
    seq      int             `json:"seq"`
}

func getAllOrders() (map[string]PoloniexOrderbookResponse, error) {
    vals := url.Values{}
    vals.Set("currencyPair", "all")
    vals.Set("depth", order_list_len)

    resp := map[string]PoloniexOrderbookResponse{}
    path := fmt.Sprintf("%s/public?command=returnOrderBook&%s", polo_api_url, vals.Encode())

    err := core.HTTPSendGetRequestTimeout(path, true, &resp, time.Second*10)

    if err != nil {
        err_count++
        if err_count > total_err_count {
            panic("")
        }
        return resp, err
    }

    return resp, nil
}

func insert_blocks(data [][]interface{}, db_block *[]string, publish_block *[]interface{}) (first string) {
    block := make([][]float64, len(data))

    if len(data) == 0 {
        return first
    }

    for i, data := range data {
        rate, _ := strconv.ParseFloat(data[0].(string), 64)
        if i == 0 {
            first = strconv.FormatFloat(rate, 'f', 8, 64)
        }
        amount := data[1].(float64)

        // publish
        block[i] = []float64{rate, amount}
    }

    // db
    *db_block = append(*db_block, first)

    // publish
    *publish_block = append(*publish_block, block)

    return first
}

func isOnUSD(pairname string) bool {
    items := core.StrSplit(pairname, "_")
    basename := items[0]
    return basename == "USDT"
}

func isOnBTC(pairname string) bool {
    items := core.StrSplit(pairname, "_")
    basename := items[0]
    return basename == "BTC"
}

func updateOrders() {
    glasses_publisher := core.CreateEventPublisher("main", "orders")

    core.PublishInfo("Start to orders call.")

    for {
        dt := time.Now()
        unix := dt.Unix()
        datetime_str := core.TimeToString(dt)
        resp, err := getAllOrders()
        if err != nil {
            message := fmt.Sprintf("poloniex: updateOrders: %v", err.Error())
            core.PublishError(message)
            continue
        }

        // common blocks
        var values []string
        var all_publish_blocks []interface{}

        for pairname, orders := range resp {
            // check by USD
            if isOnUSD(pairname) || isOnBTC(pairname) {
                var publish_block []interface{}
                var db_block = make([]string, 0, 8)

                // PAIR
                //   db block
                //    pair id
                pair_id, err := pair.GetPairID(pairname)
                if err != nil {
                    message := fmt.Sprintf("event: POLONIEX_ORDER_SLICE: get_pair_id: err: %v", err.Error())
                    core.PublishError(message)
                }

                // db
                db_block = append(db_block, strconv.Itoa(pair_id))

                //   publish block
                publish_block = append(publish_block, pairname)

                // DATETIME
                //   db_block
                db_block = append(db_block, "'"+datetime_str+"'")

                // RATE, AMOUNTS
                //   bids
                insert_blocks(orders.Bids, &db_block, &publish_block)

                //   asks
                insert_blocks(orders.Asks, &db_block, &publish_block)

                // ADD INTO COMMON BLOCKS
                // db
                values = append(values, fmt.Sprintf("(%v)", strings.Join(db_block, ",")))
                // publish
                all_publish_blocks = append(all_publish_blocks, publish_block)
            }

        }

        values_str := core.StrJoin(values, ",")
        sql_order := "INSERT INTO orders(pair, datetime, bid_price, ask_price) VALUES " + values_str + ";"

        //values_str_last_price := core.StrJoin(values_last_price, ",")

        //sql_last_price := "INSERT INTO last_prices(pair, datetime, bid_price, ask_price) VALUES " + values_str_last_price + ";"
        //sql_delete_last_price := "DELETE FROM last_prices WHERE datetime < ('" + datetime_str + "'::TIMESTAMP - interval '1 day'); "
        //sql := sql_order + sql_last_price + sql_delete_last_price
        sql := sql_order
        // fmt.Println(sql)

        // insert db
        _, err = core.Postgres.Exec(sql)
        if err != nil {
            message := fmt.Sprintf("poloniex: updateOrders: Insert: %v", err.Error())
            core.PublishError(message)
        }

        // publish
        err = glasses_publisher.SendSequence(all_publish_blocks, unix)
        if err != nil {
            message := fmt.Sprintf("poloniex: updateOrders: Send event GLASSES_PUBLISHER: %v", err.Error())
            core.PublishError(message)
        }

        if !activeTradesUpdater {
        	core.PublishInfo("start update trades...")
        	go UpdateTrades()
        	activeTradesUpdater = true
        }

        // check time
        t := time.Now()
        duration := t.Sub(dt)
        //core.PublishInfo(fmt.Sprintf("%v", duration.Seconds()))
        s := time.Millisecond * 2000
        if duration < s {
            sleep := s - duration
            core.PublishInfo(fmt.Sprintf("SLEEP: %v", sleep))
            time.Sleep(sleep)
        }

    }

}
