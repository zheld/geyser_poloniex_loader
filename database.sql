--
-- Document generated automatically by GEYSER v0.1.8
--
-- Warning: any edits to this file will be lost
--

-- TABLE CONVERTER  --------------------------------------------------------------------------------------------------------------

INSERT INTO history_conversions(structure) VALUES ('{"resources_sql": {"trades": {"indexes": [{"table": "trades", "unique": false, "name": "trades_pair_datetime_index", "fields": ["pair", "datetime"]}, {"table": "trades", "unique": false, "name": "trades_datetime_index", "fields": ["datetime"]}, {"table": "trades", "unique": true, "name": "trades_trade_id_index", "fields": ["trade_id"]}, {"table": "trades", "unique": true, "name": "trades_id_index", "fields": ["id"]}], "fields": [{"pkey": true, "_default": null, "name": "id", "foreign": null, "type": "INTEGER"}, {"pkey": false, "_default": null, "name": "datetime", "foreign": null, "type": "TIMESTAMP"}, {"pkey": false, "_default": null, "name": "trade_id", "foreign": null, "type": "BIGINT "}, {"pkey": false, "_default": null, "name": "pair", "foreign": {"table": "pair", "name": "id"}, "type": "INTEGER"}, {"pkey": false, "_default": null, "name": "price", "foreign": null, "type": "DECIMAL (15, 8)"}, {"pkey": false, "_default": null, "name": "amount", "foreign": null, "type": "DECIMAL (15, 8)"}, {"pkey": false, "_default": null, "name": "buy", "foreign": null, "type": "BOOL"}]}, "history_conversions": {"indexes": [{"table": "history_conversions", "unique": true, "name": "history_conversions_date_index", "fields": ["date"]}], "fields": [{"pkey": false, "_default": "now()", "name": "date", "foreign": null, "type": "TIMESTAMP"}, {"pkey": false, "_default": null, "name": "build", "foreign": null, "type": "TEXT"}, {"pkey": false, "_default": null, "name": "structure", "foreign": null, "type": "TEXT"}]}, "orders": {"indexes": [{"table": "orders", "unique": true, "name": "orders_pair_datetime_index", "fields": ["pair", "datetime"]}], "fields": [{"pkey": true, "_default": null, "name": "id", "foreign": null, "type": "INTEGER"}, {"pkey": false, "_default": null, "name": "pair", "foreign": {"table": "pair", "name": "id"}, "type": "INTEGER"}, {"pkey": false, "_default": null, "name": "datetime", "foreign": null, "type": "TIMESTAMP"}, {"pkey": false, "_default": null, "name": "bid_price", "foreign": null, "type": "DECIMAL (15, 8)"}, {"pkey": false, "_default": null, "name": "ask_price", "foreign": null, "type": "DECIMAL (15, 8)"}]}, "pair": {"indexes": [{"table": "pair", "unique": true, "name": "pair_name_index", "fields": ["name"]}], "fields": [{"pkey": true, "_default": null, "name": "id", "foreign": null, "type": "INTEGER"}, {"pkey": false, "_default": null, "name": "name", "foreign": null, "type": "TEXT"}]}}}');
        
-- CUSTOM functions --------------------------------------------------------------------------------------------------------------

-- DATAS functions --------------------------------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION public.pair_get_id( p_name TEXT )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  SELECT
    id
  FROM
    pair
  WHERE
    name = p_name INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.pair_read_by_unique( _name TEXT )
    RETURNS RECORD
    LANGUAGE plpgsql
AS $function$
DECLARE
    result RECORD;
BEGIN
    SELECT
        pair.id, pair.name
    FROM
        pair
    WHERE
        name = _name INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.pair_read( _id INTEGER )
    RETURNS RECORD
    LANGUAGE plpgsql
AS $function$
DECLARE
    result RECORD;
BEGIN
    SELECT
        pair.id, pair.name
    FROM
        pair
    WHERE
        pair.id = _id INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.pair_insert_or_read( p_name TEXT )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  SELECT
    id
  FROM
    pair
  WHERE
    name = p_name INTO result;
  IF result IS NULL THEN
    INSERT INTO pair (name) VALUES (p_name) RETURNING id INTO result;
  END IF;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.pair_insert_or_read_light( p_name TEXT )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  SELECT
    id
  FROM
    pair
  WHERE
    name = p_name INTO result;
  IF result IS NULL THEN
    INSERT INTO pair (name) VALUES (p_name) RETURNING id INTO result;
  END IF;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.pair_insert( p_name TEXT )
    RETURNS INTEGER
    LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  INSERT INTO pair (name)
      VALUES (p_name) RETURNING pair.id INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.pair_update( p_name TEXT, _id INTEGER )
    RETURNS VOID
LANGUAGE plpgsql
AS $function$
BEGIN
  UPDATE
    pair
  SET
    name = p_name
  WHERE id = _id;
END;
$function$;

CREATE OR REPLACE FUNCTION public.pair_upsert( p_name TEXT )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$

DECLARE
  result INTEGER;

BEGIN
  
  SELECT
    id
  FROM
    pair
  WHERE
    name = p_name INTO result;
  
  IF result IS NULL THEN
    INSERT INTO pair (name) VALUES (p_name) RETURNING id INTO result;
  ELSE 
    UPDATE pair SET name = p_name WHERE id = result;
  END IF;
  
  RETURN result;

END;
$function$;

CREATE OR REPLACE FUNCTION public.trades_get_id( p_id INTEGER )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  SELECT
    id
  FROM
    trades
  WHERE
    id = p_id INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.trades_read_by_unique( _id INTEGER )
    RETURNS RECORD
    LANGUAGE plpgsql
AS $function$
DECLARE
    result RECORD;
BEGIN
    SELECT
        trades.id, trades.datetime, trades.trade_id, trades.pair, trades.price, trades.amount, trades.buy
    FROM
        trades
    WHERE
        id = _id INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.trades_read( _id INTEGER )
    RETURNS RECORD
    LANGUAGE plpgsql
AS $function$
DECLARE
    result RECORD;
BEGIN
    SELECT
        trades.id, trades.datetime, trades.trade_id, trades.pair, trades.price, trades.amount, trades.buy
    FROM
        trades
    WHERE
        trades.id = _id INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.trades_insert_or_read( p_datetime TIMESTAMP, p_trade_id BIGINT , p_pair INTEGER, p_price DECIMAL (15, 8), p_amount DECIMAL (15, 8), p_buy BOOL )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  SELECT
    id
  FROM
    trades
  WHERE
    id = p_id INTO result;
  IF result IS NULL THEN
    INSERT INTO trades (datetime, trade_id, pair, price, amount, buy) VALUES (p_datetime, p_trade_id, p_pair, p_price, p_amount, p_buy) RETURNING id INTO result;
  END IF;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.trades_insert_or_read_light( p_id INTEGER )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  SELECT
    id
  FROM
    trades
  WHERE
    id = p_id INTO result;
  IF result IS NULL THEN
    INSERT INTO trades (id) VALUES (p_id) RETURNING id INTO result;
  END IF;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.trades_insert( p_datetime TIMESTAMP, p_trade_id BIGINT , p_pair INTEGER, p_price DECIMAL (15, 8), p_amount DECIMAL (15, 8), p_buy BOOL )
    RETURNS INTEGER
    LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  INSERT INTO trades (datetime, trade_id, pair, price, amount, buy)
      VALUES (p_datetime, p_trade_id, p_pair, p_price, p_amount, p_buy) RETURNING trades.id INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.trades_update( p_datetime TIMESTAMP, p_trade_id BIGINT , p_pair INTEGER, p_price DECIMAL (15, 8), p_amount DECIMAL (15, 8), p_buy BOOL, _id INTEGER )
    RETURNS VOID
LANGUAGE plpgsql
AS $function$
BEGIN
  UPDATE
    trades
  SET
    datetime = p_datetime,
    trade_id = p_trade_id,
    pair = p_pair,
    price = p_price,
    amount = p_amount,
    buy = p_buy
  WHERE id = _id;
END;
$function$;

CREATE OR REPLACE FUNCTION public.trades_upsert( p_datetime TIMESTAMP, p_trade_id BIGINT , p_pair INTEGER, p_price DECIMAL (15, 8), p_amount DECIMAL (15, 8), p_buy BOOL )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$

DECLARE
  result INTEGER;

BEGIN
  
  SELECT
    id
  FROM
    trades
  WHERE
    id = p_id INTO result;
  
  IF result IS NULL THEN
    INSERT INTO trades (datetime, trade_id, pair, price, amount, buy) VALUES (p_datetime, p_trade_id, p_pair, p_price, p_amount, p_buy) RETURNING id INTO result;
  ELSE 
    UPDATE trades SET datetime = p_datetime, trade_id = p_trade_id, pair = p_pair, price = p_price, amount = p_amount, buy = p_buy WHERE id = result;
  END IF;
  
  RETURN result;

END;
$function$;

CREATE OR REPLACE FUNCTION public.orders_get_id( p_pair INTEGER, p_datetime TIMESTAMP )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  SELECT
    id
  FROM
    orders
  WHERE
    pair = p_pair AND datetime = p_datetime INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.orders_read_by_unique( _pair INTEGER, _datetime TIMESTAMP )
    RETURNS RECORD
    LANGUAGE plpgsql
AS $function$
DECLARE
    result RECORD;
BEGIN
    SELECT
        orders.id, orders.pair, orders.datetime, orders.bid_price, orders.ask_price
    FROM
        orders
    WHERE
        pair = _pair AND datetime = _datetime INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.orders_read( _id INTEGER )
    RETURNS RECORD
    LANGUAGE plpgsql
AS $function$
DECLARE
    result RECORD;
BEGIN
    SELECT
        orders.id, orders.pair, orders.datetime, orders.bid_price, orders.ask_price
    FROM
        orders
    WHERE
        orders.id = _id INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.orders_insert_or_read( p_pair INTEGER, p_datetime TIMESTAMP, p_bid_price DECIMAL (15, 8), p_ask_price DECIMAL (15, 8) )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  SELECT
    id
  FROM
    orders
  WHERE
    pair = p_pair AND datetime = p_datetime INTO result;
  IF result IS NULL THEN
    INSERT INTO orders (pair, datetime, bid_price, ask_price) VALUES (p_pair, p_datetime, p_bid_price, p_ask_price) RETURNING id INTO result;
  END IF;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.orders_insert_or_read_light( p_pair INTEGER, p_datetime TIMESTAMP )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  SELECT
    id
  FROM
    orders
  WHERE
    pair = p_pair AND datetime = p_datetime INTO result;
  IF result IS NULL THEN
    INSERT INTO orders (pair, datetime) VALUES (p_pair, p_datetime) RETURNING id INTO result;
  END IF;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.orders_insert( p_pair INTEGER, p_datetime TIMESTAMP, p_bid_price DECIMAL (15, 8), p_ask_price DECIMAL (15, 8) )
    RETURNS INTEGER
    LANGUAGE plpgsql
AS $function$
DECLARE
  result INTEGER;
BEGIN
  INSERT INTO orders (pair, datetime, bid_price, ask_price)
      VALUES (p_pair, p_datetime, p_bid_price, p_ask_price) RETURNING orders.id INTO result;
  RETURN result;
END;
$function$;

CREATE OR REPLACE FUNCTION public.orders_update( p_pair INTEGER, p_datetime TIMESTAMP, p_bid_price DECIMAL (15, 8), p_ask_price DECIMAL (15, 8), _id INTEGER )
    RETURNS VOID
LANGUAGE plpgsql
AS $function$
BEGIN
  UPDATE
    orders
  SET
    pair = p_pair,
    datetime = p_datetime,
    bid_price = p_bid_price,
    ask_price = p_ask_price
  WHERE id = _id;
END;
$function$;

CREATE OR REPLACE FUNCTION public.orders_upsert( p_pair INTEGER, p_datetime TIMESTAMP, p_bid_price DECIMAL (15, 8), p_ask_price DECIMAL (15, 8) )
 RETURNS integer
 LANGUAGE plpgsql
AS $function$

DECLARE
  result INTEGER;

BEGIN
  
  SELECT
    id
  FROM
    orders
  WHERE
    pair = p_pair AND datetime = p_datetime INTO result;
  
  IF result IS NULL THEN
    INSERT INTO orders (pair, datetime, bid_price, ask_price) VALUES (p_pair, p_datetime, p_bid_price, p_ask_price) RETURNING id INTO result;
  ELSE 
    UPDATE orders SET pair = p_pair, datetime = p_datetime, bid_price = p_bid_price, ask_price = p_ask_price WHERE id = result;
  END IF;
  
  RETURN result;

END;
$function$;