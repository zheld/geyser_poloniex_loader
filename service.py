import geyser as gs

class Config:
    lng = "golang" # or "python"
    db_host = "localhost"
    db_port = 5432
    db_name = "poloniex_loader"
    db_user = "postgres"
    db_pswd = "postgres"


class SERVICE:
    service = gs.ServiceBase("coin_polo", version = "1.0", config = Config)
    if service:
        # PAIR
        pair = service.addData("pair")
        if pair:
            # fields
            pair_id = pair.addIdentifier()
            pair_name = pair.addField("name", gs.go_types.text)
            # table
            pair_table = pair.addTable()
            if pair_table:
                pass
            # properties
            pair.addSignatureRowIndex(pair_name)

        # TRADES
        trade = service.addData("trades")
        if trade:
            # fields
            trade_id = trade.addIdentifier()
            trade_date = trade.addField("datetime", gs.go_types.timestamp)
            trade_trade_id = trade.addField("trade_id", gs.go_types.int64)
            trade_pair = trade.addForeignOneToOne(pair_id, "pair")
            trade_price = trade.addField("price", gs.go_types.crypto_money)
            trade_amount = trade.addField("amount", gs.go_types.crypto_money)
            trade_type = trade.addField("buy", gs.go_types.boolean)
            # table
            trade_table = trade.addTable()
            if trade_table:
                trade_table.addIndex(trade_pair, trade_date)
                trade_table.addIndex(trade_date)
                trade_table.addUniqueIndex(trade_trade_id)
            # properties
            trade.addSignatureRowIndex(trade_id)

        # ORDERS
        orders = service.addData("orders")
        if orders:
            # fields
            orders_id = orders.addIdentifier()
            orders_pair = orders.addForeignOneToOne(pair_id, "pair")
            orders_datetime = orders.addField("datetime", gs.go_types.timestamp)
            orders_bids_prices = orders.addField("bid_price", gs.go_types.crypto_money)
            orders_asks_prices = orders.addField("ask_price", gs.go_types.crypto_money)
            # table
            orders_table = orders.addTable()
            if orders_table:
                pass
            # properties
            orders.addSignatureRowIndex(orders_pair, orders_datetime)

if __name__ == '__main__':
    gs.BuildService(SERVICE.service)