package updater

import (
    "core"
    "fmt"
    "time"
    "pair"
)

func normalize_num(num float64) (snum string) {
    sn := core.FloatToStr(num)
    sa := core.StrSplit(sn, ".")
    f := sa[0]
    if len(f) > 7 {
        f = f[:7]
    }
    s := sa[1]
    if len(s) > 8 {
        s = s[:8]
    }
    return f + "." + s
}

func updateTrades() {
    trades_publisher := core.CreateEventPublisher("main", "trades")
    pair_producer := pairProducer{}
    pair_producer.init()

    core.PublishInfo("Start to trades call.")

    for {
        for _, pr := range pair.PairList {
            n := time.Now()
            trade_history, err := pr.GetLastTradeHistory()

            if err != nil {
                err_count++
                if err_count > total_err_count {
                    panic("")
                }

                core.PublishError("updater.updateTrades: " + err.Error())
                time.Sleep(time.Second)
                continue
            }

            // check on empty history
            if len(trade_history) == 0 {
                core.PublishInfo("NO TRADES...")
            } else {
                // common blocks
                var values []string
                var all_publish_blocks []interface{}

                for _, trade := range trade_history {
                    // db
                    is_buy := "TRUE"
                    if trade.Type == "sell" {
                        is_buy = "FALSE"
                    }
                    item_value := fmt.Sprintf("('%v', %v, %v, %v, %v, %v)",
                        trade.Date,
                        pr.Id,
                        trade.GlobalTradeID,
                        normalize_num(trade.Rate),
                        normalize_num(trade.Amount),
                        is_buy)
                    values = append(values, item_value)

                    // publish
                    amount := trade.Amount
                    if trade.Type == "sell" {
                        amount *= -1
                    }
                    // unix date
                    udate, _ := core.TimeStrWithoutNanosecToTime(trade.Date)
                    unx := udate.Unix()
                    item_publish := []interface{}{
                        unx,
                        trade.Rate,
                        amount,
                    }
                    all_publish_blocks = append(all_publish_blocks, item_publish)

                }

                // db
                values_str := core.StrJoin(values, ",")
                sql := "INSERT INTO trades(datetime, pair, trade_id, price, amount, buy) VALUES " + values_str + " ON CONFLICT (trade_id) DO NOTHING;"
                core.PublishInfo(fmt.Sprintf("new trades block. %v", ""))

                // insert db
                _, err = core.Postgres.Exec(sql)
                if err != nil {
                    message := fmt.Sprintf("poloniex: updateTrades: Insert: %v", err.Error())
                    core.PublishError(message)
                }

                // publish
                err = trades_publisher.SendSequence(all_publish_blocks, pr.Name)
                if err != nil {
                    message := fmt.Sprintf("poloniex: updateTrades: Send event TRADES_PUBLISHER: %v", err.Error())
                    core.PublishError(message)
                }

            }


            // check time
            t := time.Now()
            duration := t.Sub(n)
            core.PublishInfo(fmt.Sprintf("%v", duration.Seconds()))
            s := time.Millisecond * 2000
            if duration < s {
                sleep := s - duration
                core.PublishInfo(fmt.Sprintf("SLEEP: %v", sleep))
                time.Sleep(sleep)
            }

        }
    }

}

func updateTradesWithPairProducer() {
    glasses_publisher := core.CreateEventPublisher("main", "trades")
    pair_producer := pairProducer{}
    pair_producer.init()

    core.PublishInfo("Start to trades call.")

    for {
        n := time.Now()
        pr := pair_producer.Get()
        trade_history, err := pr.GetLastTradeHistory()

        if err != nil {
            err_count++
            if err_count > 30 {
                panic("")
            }

            core.PublishError("updater.updateTrades: " + err.Error())
            time.Sleep(time.Second)
            continue
        }

        // check on empty history
        if len(trade_history) == 0 {
            core.PublishInfo("NO TRADES...")
            continue
        }

        // common blocks
        var values []string
        var all_publish_blocks []interface{}

        for _, trade := range trade_history {
            // db
            is_buy := "TRUE"
            if trade.Type == "sell" {
                is_buy = "FALSE"
            }
            item_value := fmt.Sprintf("('%v', %v, %v, %v, %v, %v)",
                trade.Date,
                pr.Id,
                trade.GlobalTradeID,
                normalize_num(trade.Rate),
                normalize_num(trade.Amount),
                is_buy)
            values = append(values, item_value)

            // publish
            amount := trade.Amount
            if trade.Type == "sell" {
                amount *= -1
            }
            // unix date
            udate, _ := core.TimeStrWithoutNanosecToTime(trade.Date)
            unx := udate.Unix()
            item_publish := []interface{}{
                unx,
                trade.Rate,
                amount,
            }
            all_publish_blocks = append(all_publish_blocks, item_publish)

        }

        // db
        values_str := core.StrJoin(values, ",")
        sql := "INSERT INTO trades(datetime, pair, trade_id, price, amount, buy) VALUES " + values_str + " ON CONFLICT (trade_id) DO NOTHING;"
        core.PublishInfo(fmt.Sprintf("new trades block. %v", ""))

        // insert db
        _, err = core.Postgres.Exec(sql)
        if err != nil {
            message := fmt.Sprintf("poloniex: updateTrades: Insert: %v", err.Error())
            core.PublishError(message)
        }

        // publish
        err = glasses_publisher.SendSequence(all_publish_blocks, pr.Id)
        if err != nil {
            message := fmt.Sprintf("poloniex: updateTrades: Send event TRADES_PUBLISHER: %v", err.Error())
            core.PublishError(message)
        }

        // check time
        t := time.Now()
        duration := t.Sub(n)
        core.PublishInfo(fmt.Sprintf("%v", duration.Seconds()))
        s := time.Millisecond * 400
        if duration < s {
            sleep := s - duration
            core.PublishInfo(fmt.Sprintf("SLEEP: %v", sleep))
            time.Sleep(sleep)
        }

    }

}

func updateUSDTradesWithPairProducer() {
    glasses_publisher := core.CreateEventPublisher("main", "trades")
    pair_producer := pairProducer{}
    pair_producer.init()

    core.PublishInfo("Start to trades call.")

    for {
        n := time.Now()
        pr := pair_producer.Get()
        trade_history, err := pr.GetLastTradeHistory()

        if err != nil {
            err_count++
            if err_count > 30 {
                panic("")
            }

            core.PublishError("updater.updateTrades: " + err.Error())
            time.Sleep(time.Second)
            continue
        }

        // check on empty history
        if len(trade_history) == 0 {
            core.PublishInfo("NO TRADES...")
            continue
        }

        // common blocks
        var values []string
        var all_publish_blocks []interface{}

        for _, trade := range trade_history {
            // db
            is_buy := "TRUE"
            if trade.Type == "sell" {
                is_buy = "FALSE"
            }
            item_value := fmt.Sprintf("('%v', %v, %v, %v, %v, %v)",
                trade.Date,
                pr.Id,
                trade.GlobalTradeID,
                normalize_num(trade.Rate),
                normalize_num(trade.Amount),
                is_buy)
            values = append(values, item_value)

            // publish
            amount := trade.Amount
            if trade.Type == "sell" {
                amount *= -1
            }
            // unix date
            udate, _ := core.TimeStrWithoutNanosecToTime(trade.Date)
            unx := udate.Unix()
            item_publish := []interface{}{
                unx,
                trade.Rate,
                amount,
            }
            all_publish_blocks = append(all_publish_blocks, item_publish)

        }

        // db
        values_str := core.StrJoin(values, ",")
        sql := "INSERT INTO trades(datetime, pair, trade_id, price, amount, buy) VALUES " + values_str + " ON CONFLICT (trade_id) DO NOTHING;"
        core.PublishInfo(fmt.Sprintf("new trades block. %v", ""))

        // insert db
        _, err = core.Postgres.Exec(sql)
        if err != nil {
            message := fmt.Sprintf("poloniex: updateTrades: Insert: %v", err.Error())
            core.PublishError(message)
        }

        // publish
        err = glasses_publisher.SendSequence(all_publish_blocks, pr.Id)
        if err != nil {
            message := fmt.Sprintf("poloniex: updateTrades: Send event TRADES_PUBLISHER: %v", err.Error())
            core.PublishError(message)
        }

        // check time
        t := time.Now()
        duration := t.Sub(n)
        core.PublishInfo(fmt.Sprintf("%v", duration.Seconds()))
        s := time.Millisecond * 400
        if duration < s {
            sleep := s - duration
            core.PublishInfo(fmt.Sprintf("SLEEP: %v", sleep))
            time.Sleep(sleep)
        }

    }

}
