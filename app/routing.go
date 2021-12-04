package app

import (
	r "tubespbbo/modules/router"
)

func mapURLs() {
	r.RouteUsers(router)
	r.RoutePaymentMethods(router)
	r.RoutePayments(router)
	r.RouteExpenses(router)
	r.RouteProducts(router)
	r.RouteTransactionMethods(router)
	r.RouteTransactions(router)
	r.RouterReport(router)
}
