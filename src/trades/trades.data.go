//
// Document generated automatically by GEYSER v0.1.8
//
// Warning: any edits to this file will be lost
//

package trades

import (
    "time"
    "fmt"
    "core"
    "strings"
)


type DBI_Trades struct {
    Id int
    Datetime time.Time
    Trade_id int64
    Pair int
    Price float64
    Amount float64
    Buy bool
}

func (self *DBI_Trades) String() (str string) {
    str = fmt.Sprintf("type: DBI_Trades [ Id: %v, Datetime: %v, Trade_id: %v, Pair: %v, Price: %v, Amount: %v, Buy: %v ]", self.Id, self.Datetime, self.Trade_id, self.Pair, self.Price, self.Amount, self.Buy)
    return
}

func (self *DBI_Trades) Update() (error) {
    sql := "SELECT trades_update($1, $2, $3, $4, $5, $6, $7)"
    _, err := core.Postgres.Exec(sql, self.Datetime, self.Trade_id, self.Pair, self.Price, self.Amount, self.Buy, self.Id)
    return err
}

func (self *DBI_Trades) SetDatetime(datetime time.Time) (error) {
    sql := "UPDATE trades SET datetime = $1 WHERE id = $2"
    _, err := core.Postgres.Exec(sql, datetime, self.Id)
    if err != nil {
        return err
    }
    self.Datetime = datetime
    return nil
}

func (self *DBI_Trades) SetTrade_id(trade_id int64) (error) {
    sql := "UPDATE trades SET trade_id = $1 WHERE id = $2"
    _, err := core.Postgres.Exec(sql, trade_id, self.Id)
    if err != nil {
        return err
    }
    self.Trade_id = trade_id
    return nil
}

func (self *DBI_Trades) SetPair(pair int) (error) {
    sql := "UPDATE trades SET pair = $1 WHERE id = $2"
    _, err := core.Postgres.Exec(sql, pair, self.Id)
    if err != nil {
        return err
    }
    self.Pair = pair
    return nil
}

func (self *DBI_Trades) SetPrice(price float64) (error) {
    sql := "UPDATE trades SET price = $1 WHERE id = $2"
    _, err := core.Postgres.Exec(sql, price, self.Id)
    if err != nil {
        return err
    }
    self.Price = price
    return nil
}

func (self *DBI_Trades) SetAmount(amount float64) (error) {
    sql := "UPDATE trades SET amount = $1 WHERE id = $2"
    _, err := core.Postgres.Exec(sql, amount, self.Id)
    if err != nil {
        return err
    }
    self.Amount = amount
    return nil
}

func (self *DBI_Trades) SetBuy(buy bool) (error) {
    sql := "UPDATE trades SET buy = $1 WHERE id = $2"
    _, err := core.Postgres.Exec(sql, buy, self.Id)
    if err != nil {
        return err
    }
    self.Buy = buy
    return nil
}

func (self *DBI_Trades) Fields() ([]interface{}) {
    res := []interface{}{}
    res = append(res, self.Id)
    if self.Datetime.IsZero() {
        res = append(res, "")
    } else {
        res = append(res, self.Datetime.String())
    }
    res = append(res, self.Trade_id)
    res = append(res, self.Pair)
    res = append(res, self.Price)
    res = append(res, self.Amount)
    res = append(res, self.Buy)
    return res
}

func GetId(id int) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT trades_get_id($1)", id)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("trades.GetId: " + err.Error())
        return _id, err
    }
    return _id, nil
}

func ReadByUnique(id int) (t DBI_Trades, err error) {
    t = DBI_Trades{}
    row := core.Postgres.QueryRow("SELECT * FROM trades WHERE id = $1", id)
    err = row.Scan(&t.Id, &t.Datetime, &t.Trade_id, &t.Pair, &t.Price, &t.Amount, &t.Buy)
    if err != nil {
        core.PublishError("trades.ReadByUnique: " + err.Error())
        return
    }
    return
}

func NewDBI_Trades() (*DBI_Trades) {
    return &DBI_Trades{} 
}

func Read(id int) (t DBI_Trades, err error) {
    t = DBI_Trades{}
    row := core.Postgres.QueryRow("SELECT * FROM trades WHERE id = $1", id)
    err = row.Scan(&t.Id, &t.Datetime, &t.Trade_id, &t.Pair, &t.Price, &t.Amount, &t.Buy)
    if err != nil {
        core.PublishError("trades.Read: " + err.Error())
        return
    }
    return
}

func InsertReadReturnInstance(datetime time.Time, trade_id int64, pair int, price float64, amount float64, buy bool) (DBI_Trades, error) {
    res := DBI_Trades{
        Datetime: datetime,
        Trade_id: trade_id,
        Pair: pair,
        Price: price,
        Amount: amount,
        Buy: buy,
    }
    row := core.Postgres.QueryRow("SELECT trades_insert_or_read($1, $2, $3, $4, $5, $6)", datetime, trade_id, pair, price, amount, buy)
    err := row.Scan(&res.Id)
    if err != nil {
        core.PublishError("trades.InsertReadReturnInstance: " + err.Error())
        return res, err
    }
    return res, nil
}

func InsertRead(datetime time.Time, trade_id int64, pair int, price float64, amount float64, buy bool) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT trades_insert_or_read($1, $2, $3, $4, $5, $6)", datetime, trade_id, pair, price, amount, buy)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("trades.InsertRead: " + err.Error())
        return _id, err
    }
    return _id, nil
}

func InsertReadLight(id int) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT trades_insert_or_read_light($1)", id)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("trades.InsertReadLight: " + err.Error())
        return _id, err
    }
    return _id, nil
}

func InsertReturnInstance(datetime time.Time, trade_id int64, pair int, price float64, amount float64, buy bool) (DBI_Trades, error) {
    newi := DBI_Trades{
        Datetime: datetime,
        Trade_id: trade_id,
        Pair: pair,
        Price: price,
        Amount: amount,
        Buy: buy,
    }
    row := core.Postgres.QueryRow("SELECT trades_insert($1, $2, $3, $4, $5, $6)", datetime, trade_id, pair, price, amount, buy)
    err := row.Scan(&newi.Id)
    if err != nil {
        return newi, err
    }
    return newi, nil
}

func Insert(datetime time.Time, trade_id int64, pair int, price float64, amount float64, buy bool) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT trades_insert($1, $2, $3, $4, $5, $6)", datetime, trade_id, pair, price, amount, buy)
    err = row.Scan(&_id)
    if err != nil {
        return _id, err
    }
    return _id, nil
}

func List() (res []DBI_Trades, err error) {
    sql := "SELECT id, datetime, trade_id, pair, price, amount, buy FROM trades"

    rows, err := core.Postgres.Query(sql)
    if err != nil {
        core.PublishError("trades.List: " + err.Error())
        return res, err
    }

    defer rows.Close()

    for rows.Next() {
        p := DBI_Trades{}
        err := rows.Scan(&p.Id, &p.Datetime, &p.Trade_id, &p.Pair, &p.Price, &p.Amount, &p.Buy)
        res = append(res, p)
        if err != nil {
            core.PublishError("trades.List: " + err.Error())
            return res, err
        }
    }
    return res, err
}

func ListFilter(from *int, limit int, filter_list *[]string) (res []DBI_Trades, count int, err error) {
    where_block_list := []string{}

    if filter_list != nil {
        for _, filter := range *filter_list {
            where_block_list = append(where_block_list, filter)
        }
    }

    if from != nil {
        where_block_list = append(where_block_list, fmt.Sprintf("id > %v", *from))
        //where_block_list = append(where_block_list, fmt.Sprintf("id <= %v", from + limit))
    }

    limit_block := ""
    if limit != 0 {
        limit_block = fmt.Sprintf(" LIMIT %v", limit)
    }

    where_block := ""
    if len(where_block_list) != 0 {
        where_block = strings.Join(where_block_list, " AND ")
        where_block = "WHERE " + where_block
    }
    sql := "SELECT id, datetime, trade_id, pair, price, amount, buy FROM trades " + where_block + limit_block

    rows, err := core.Postgres.Query(sql)
    if err != nil {
        core.PublishError("trades.ListFilter: " + err.Error())
        return res, count, err
    }

    defer rows.Close()

    for rows.Next() {
        p := DBI_Trades{}
        err := rows.Scan(&p.Id, &p.Datetime, &p.Trade_id, &p.Pair, &p.Price, &p.Amount, &p.Buy)
        res = append(res, p)
        if err != nil {
            core.PublishError("trades.List: " + err.Error())
            return res, count, err
        }
        count++
        if *from < p.Id {
            *from = p.Id
        }
    }
    return res, count, err
}

func Delete(id int) (error) {
    sql := "DELETE FROM trades WHERE id = $1"
    _, err := core.Postgres.Exec(sql, id)
    return err
}

func UpsertReturnInstance(datetime time.Time, trade_id int64, pair int, price float64, amount float64, buy bool) (DBI_Trades, error) {
    res := DBI_Trades{
        Datetime: datetime,
        Trade_id: trade_id,
        Pair: pair,
        Price: price,
        Amount: amount,
        Buy: buy,
    }
    row := core.Postgres.QueryRow("SELECT trades_upsert($1, $2, $3, $4, $5, $6)", datetime, trade_id, pair, price, amount, buy)
    err := row.Scan(&res.Id)
    if err != nil {
        core.PublishError("trades.InsertRead: " + err.Error())
        return res, err
    }
    return res, nil
}

func Upsert(datetime time.Time, trade_id int64, pair int, price float64, amount float64, buy bool) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT trades_upsert($1, $2, $3, $4, $5, $6)", datetime, trade_id, pair, price, amount, buy)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("trades.InsertRead: " + err.Error())
        return _id, err
    }
    return _id, nil
}