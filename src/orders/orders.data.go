//
// Document generated automatically by GEYSER v0.1.8
//
// Warning: any edits to this file will be lost
//

package orders

import (
    "time"
    "fmt"
    "core"
    "strings"
)


type DBI_Orders struct {
    Id int
    Pair int
    Datetime time.Time
    Bid_price float64
    Ask_price float64
}

func (self *DBI_Orders) String() (str string) {
    str = fmt.Sprintf("type: DBI_Orders [ Id: %v, Pair: %v, Datetime: %v, Bid_price: %v, Ask_price: %v ]", self.Id, self.Pair, self.Datetime, self.Bid_price, self.Ask_price)
    return
}

func (self *DBI_Orders) Update() (error) {
    sql := "SELECT orders_update($1, $2, $3, $4, $5)"
    _, err := core.Postgres.Exec(sql, self.Pair, self.Datetime, self.Bid_price, self.Ask_price, self.Id)
    return err
}

func (self *DBI_Orders) SetPair(pair int) (error) {
    sql := "UPDATE orders SET pair = $1 WHERE id = $2"
    _, err := core.Postgres.Exec(sql, pair, self.Id)
    if err != nil {
        return err
    }
    self.Pair = pair
    return nil
}

func (self *DBI_Orders) SetDatetime(datetime time.Time) (error) {
    sql := "UPDATE orders SET datetime = $1 WHERE id = $2"
    _, err := core.Postgres.Exec(sql, datetime, self.Id)
    if err != nil {
        return err
    }
    self.Datetime = datetime
    return nil
}

func (self *DBI_Orders) SetBid_price(bid_price float64) (error) {
    sql := "UPDATE orders SET bid_price = $1 WHERE id = $2"
    _, err := core.Postgres.Exec(sql, bid_price, self.Id)
    if err != nil {
        return err
    }
    self.Bid_price = bid_price
    return nil
}

func (self *DBI_Orders) SetAsk_price(ask_price float64) (error) {
    sql := "UPDATE orders SET ask_price = $1 WHERE id = $2"
    _, err := core.Postgres.Exec(sql, ask_price, self.Id)
    if err != nil {
        return err
    }
    self.Ask_price = ask_price
    return nil
}

func (self *DBI_Orders) Fields() ([]interface{}) {
    res := []interface{}{}
    res = append(res, self.Id)
    res = append(res, self.Pair)
    if self.Datetime.IsZero() {
        res = append(res, "")
    } else {
        res = append(res, self.Datetime.String())
    }
    res = append(res, self.Bid_price)
    res = append(res, self.Ask_price)
    return res
}

func GetId(pair int, datetime time.Time) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT orders_get_id($1, $2)", pair, datetime)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("orders.GetId: " + err.Error())
        return _id, err
    }
    return _id, nil
}

func ReadByUnique(pair int, datetime time.Time) (t DBI_Orders, err error) {
    t = DBI_Orders{}
    row := core.Postgres.QueryRow("SELECT * FROM orders WHERE pair = $1 AND datetime = $2", pair, datetime)
    err = row.Scan(&t.Id, &t.Pair, &t.Datetime, &t.Bid_price, &t.Ask_price)
    if err != nil {
        core.PublishError("orders.ReadByUnique: " + err.Error())
        return
    }
    return
}

func NewDBI_Orders() (*DBI_Orders) {
    return &DBI_Orders{} 
}

func Read(id int) (t DBI_Orders, err error) {
    t = DBI_Orders{}
    row := core.Postgres.QueryRow("SELECT * FROM orders WHERE id = $1", id)
    err = row.Scan(&t.Id, &t.Pair, &t.Datetime, &t.Bid_price, &t.Ask_price)
    if err != nil {
        core.PublishError("orders.Read: " + err.Error())
        return
    }
    return
}

func InsertReadReturnInstance(pair int, datetime time.Time, bid_price float64, ask_price float64) (DBI_Orders, error) {
    res := DBI_Orders{
        Pair: pair,
        Datetime: datetime,
        Bid_price: bid_price,
        Ask_price: ask_price,
    }
    row := core.Postgres.QueryRow("SELECT orders_insert_or_read($1, $2, $3, $4)", pair, datetime, bid_price, ask_price)
    err := row.Scan(&res.Id)
    if err != nil {
        core.PublishError("orders.InsertReadReturnInstance: " + err.Error())
        return res, err
    }
    return res, nil
}

func InsertRead(pair int, datetime time.Time, bid_price float64, ask_price float64) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT orders_insert_or_read($1, $2, $3, $4)", pair, datetime, bid_price, ask_price)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("orders.InsertRead: " + err.Error())
        return _id, err
    }
    return _id, nil
}

func InsertReadLight(pair int, datetime time.Time) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT orders_insert_or_read_light($1, $2)", pair, datetime)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("orders.InsertReadLight: " + err.Error())
        return _id, err
    }
    return _id, nil
}

func InsertReturnInstance(pair int, datetime time.Time, bid_price float64, ask_price float64) (DBI_Orders, error) {
    newi := DBI_Orders{
        Pair: pair,
        Datetime: datetime,
        Bid_price: bid_price,
        Ask_price: ask_price,
    }
    row := core.Postgres.QueryRow("SELECT orders_insert($1, $2, $3, $4)", pair, datetime, bid_price, ask_price)
    err := row.Scan(&newi.Id)
    if err != nil {
        return newi, err
    }
    return newi, nil
}

func Insert(pair int, datetime time.Time, bid_price float64, ask_price float64) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT orders_insert($1, $2, $3, $4)", pair, datetime, bid_price, ask_price)
    err = row.Scan(&_id)
    if err != nil {
        return _id, err
    }
    return _id, nil
}

func List() (res []DBI_Orders, err error) {
    sql := "SELECT id, pair, datetime, bid_price, ask_price FROM orders"

    rows, err := core.Postgres.Query(sql)
    if err != nil {
        core.PublishError("orders.List: " + err.Error())
        return res, err
    }

    defer rows.Close()

    for rows.Next() {
        p := DBI_Orders{}
        err := rows.Scan(&p.Id, &p.Pair, &p.Datetime, &p.Bid_price, &p.Ask_price)
        res = append(res, p)
        if err != nil {
            core.PublishError("orders.List: " + err.Error())
            return res, err
        }
    }
    return res, err
}

func ListFilter(from *int, limit int, filter_list *[]string) (res []DBI_Orders, count int, err error) {
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
    sql := "SELECT id, pair, datetime, bid_price, ask_price FROM orders " + where_block + limit_block

    rows, err := core.Postgres.Query(sql)
    if err != nil {
        core.PublishError("orders.ListFilter: " + err.Error())
        return res, count, err
    }

    defer rows.Close()

    for rows.Next() {
        p := DBI_Orders{}
        err := rows.Scan(&p.Id, &p.Pair, &p.Datetime, &p.Bid_price, &p.Ask_price)
        res = append(res, p)
        if err != nil {
            core.PublishError("orders.List: " + err.Error())
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
    sql := "DELETE FROM orders WHERE id = $1"
    _, err := core.Postgres.Exec(sql, id)
    return err
}

func UpsertReturnInstance(pair int, datetime time.Time, bid_price float64, ask_price float64) (DBI_Orders, error) {
    res := DBI_Orders{
        Pair: pair,
        Datetime: datetime,
        Bid_price: bid_price,
        Ask_price: ask_price,
    }
    row := core.Postgres.QueryRow("SELECT orders_upsert($1, $2, $3, $4)", pair, datetime, bid_price, ask_price)
    err := row.Scan(&res.Id)
    if err != nil {
        core.PublishError("orders.InsertRead: " + err.Error())
        return res, err
    }
    return res, nil
}

func Upsert(pair int, datetime time.Time, bid_price float64, ask_price float64) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT orders_upsert($1, $2, $3, $4)", pair, datetime, bid_price, ask_price)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("orders.InsertRead: " + err.Error())
        return _id, err
    }
    return _id, nil
}