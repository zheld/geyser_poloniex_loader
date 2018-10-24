package trades

import (
    "core"
    "time"
)

func GetLastDateByPair(pair_id int) (time.Time, error) {
    var tm time.Time
    row := core.Postgres.QueryRow("SELECT max(datetime) FROM trades WHERE pair = $1", pair_id)
    err := row.Scan(&tm)
    return tm, err

}
