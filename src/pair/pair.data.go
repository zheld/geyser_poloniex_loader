//
// Document generated automatically by GEYSER v0.1.8
//
// Warning: any edits to this file will be lost
//

package pair

import (
    "fmt"
    "core"
    "strings"
)


type DBI_Pair struct {
    Id int
    Name string
}

func (self *DBI_Pair) String() (str string) {
    str = fmt.Sprintf("type: DBI_Pair [ Id: %v, Name: %v ]", self.Id, self.Name)
    return
}

func (self *DBI_Pair) Update() (error) {
    sql := "SELECT pair_update($1, $2)"
    _, err := core.Postgres.Exec(sql, self.Name, self.Id)
    return err
}

func (self *DBI_Pair) SetName(name string) (error) {
    sql := "UPDATE pair SET name = $1 WHERE id = $2"
    _, err := core.Postgres.Exec(sql, name, self.Id)
    if err != nil {
        return err
    }
    self.Name = name
    return nil
}

func (self *DBI_Pair) Fields() ([]interface{}) {
    res := []interface{}{}
    res = append(res, self.Id)
    res = append(res, self.Name)
    return res
}

func GetId(name string) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT pair_get_id($1)", name)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("pair.GetId: " + err.Error())
        return _id, err
    }
    return _id, nil
}

func ReadByUnique(name string) (t DBI_Pair, err error) {
    t = DBI_Pair{}
    row := core.Postgres.QueryRow("SELECT * FROM pair WHERE name = $1", name)
    err = row.Scan(&t.Id, &t.Name)
    if err != nil {
        core.PublishError("pair.ReadByUnique: " + err.Error())
        return
    }
    return
}

func NewDBI_Pair() (*DBI_Pair) {
    return &DBI_Pair{} 
}

func Read(id int) (t DBI_Pair, err error) {
    t = DBI_Pair{}
    row := core.Postgres.QueryRow("SELECT * FROM pair WHERE id = $1", id)
    err = row.Scan(&t.Id, &t.Name)
    if err != nil {
        core.PublishError("pair.Read: " + err.Error())
        return
    }
    return
}

func InsertReadReturnInstance(name string) (DBI_Pair, error) {
    res := DBI_Pair{
        Name: name,
    }
    row := core.Postgres.QueryRow("SELECT pair_insert_or_read($1)", name)
    err := row.Scan(&res.Id)
    if err != nil {
        core.PublishError("pair.InsertReadReturnInstance: " + err.Error())
        return res, err
    }
    return res, nil
}

func InsertRead(name string) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT pair_insert_or_read($1)", name)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("pair.InsertRead: " + err.Error())
        return _id, err
    }
    return _id, nil
}

func InsertReadLight(name string) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT pair_insert_or_read_light($1)", name)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("pair.InsertReadLight: " + err.Error())
        return _id, err
    }
    return _id, nil
}

func InsertReturnInstance(name string) (DBI_Pair, error) {
    newi := DBI_Pair{
        Name: name,
    }
    row := core.Postgres.QueryRow("SELECT pair_insert($1)", name)
    err := row.Scan(&newi.Id)
    if err != nil {
        return newi, err
    }
    return newi, nil
}

func Insert(name string) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT pair_insert($1)", name)
    err = row.Scan(&_id)
    if err != nil {
        return _id, err
    }
    return _id, nil
}

func List() (res []DBI_Pair, err error) {
    sql := "SELECT id, name FROM pair"

    rows, err := core.Postgres.Query(sql)
    if err != nil {
        core.PublishError("pair.List: " + err.Error())
        return res, err
    }

    defer rows.Close()

    for rows.Next() {
        p := DBI_Pair{}
        err := rows.Scan(&p.Id, &p.Name)
        res = append(res, p)
        if err != nil {
            core.PublishError("pair.List: " + err.Error())
            return res, err
        }
    }
    return res, err
}

func ListFilter(from *int, limit int, filter_list *[]string) (res []DBI_Pair, count int, err error) {
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
    sql := "SELECT id, name FROM pair " + where_block + limit_block

    rows, err := core.Postgres.Query(sql)
    if err != nil {
        core.PublishError("pair.ListFilter: " + err.Error())
        return res, count, err
    }

    defer rows.Close()

    for rows.Next() {
        p := DBI_Pair{}
        err := rows.Scan(&p.Id, &p.Name)
        res = append(res, p)
        if err != nil {
            core.PublishError("pair.List: " + err.Error())
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
    sql := "DELETE FROM pair WHERE id = $1"
    _, err := core.Postgres.Exec(sql, id)
    return err
}

func UpsertReturnInstance(name string) (DBI_Pair, error) {
    res := DBI_Pair{
        Name: name,
    }
    row := core.Postgres.QueryRow("SELECT pair_upsert($1)", name)
    err := row.Scan(&res.Id)
    if err != nil {
        core.PublishError("pair.InsertRead: " + err.Error())
        return res, err
    }
    return res, nil
}

func Upsert(name string) (_id int, err error) {
    row := core.Postgres.QueryRow("SELECT pair_upsert($1)", name)
    err = row.Scan(&_id)
    if err != nil {
        core.PublishError("pair.InsertRead: " + err.Error())
        return _id, err
    }
    return _id, nil
}