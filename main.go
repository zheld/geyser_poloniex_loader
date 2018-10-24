//
// Document generated automatically by GEYSER v0.1.8
//
// Warning: any edits to this file will be lost
//

package main

import (
    "core"
    "pair"
    "manager"
)


func main() {
    api := map[string]core.HandlerFoo{
        "echo": echo,
        "get_pair_id": get_pair_id,
        "run": run,
    }
    core.ServiceStart(api)
}

func get_pair_id(args []interface{}) (interface{}) {
    if len(args) != 1 {
        return []interface{}{500, "no match count params"}
    }
    name := core.ToString(args[0])
    res, err := pair.GetPairID(name)
    if err != nil {
        return []interface{}{500, err.Error()}
    } else {
        return []interface{}{200, res}
    }
}

func run(args []interface{}) (interface{}) {
    if len(args) != 0 {
        return []interface{}{500, "no match count params"}
    }
    res, err := manager.RUN()
    if err != nil {
        return []interface{}{500, err.Error()}
    } else {
        return []interface{}{200, res}
    }
}

func echo(args []interface{}) (interface{}) {
    return []interface{}{200, args[0]}
}