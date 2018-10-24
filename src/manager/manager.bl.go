package manager

import (
    "core"
)

var conf = core.Config{
    Name:                   "poloniex_loader",
    Version:                "1.0",
    RabbitConn:             [][]string{{"main", "amqp://{your_amqp_connect_string}"}},
    RabbitPublogServerName: "main",
    PostgreSQLConn:         "host=localhost port=5432 dbname=poloniex_data user=postgres password=postgres",
}

func initCore() {
    core.Init(conf)
}

func init() {
    initCore()
    RUN()
}
