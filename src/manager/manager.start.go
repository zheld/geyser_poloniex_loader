package manager

import (
    "updater"
)

//!api run
func RUN() (res bool, err error) {
    // observe orders
    updater.UpdateOrders()

    res = true
    return res, nil
}
