package license

import (
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

const (
	stateKey = "state"
)

func subscriptionAsTable(subscription *stackstate_api.SubscriptionState) printer.TableData {
	var header []string
	data := make([][]interface{}, 0)
	switch l := subscription.GetActualInstance().(type) {
	case *stackstate_api.ExpiredSubscription:
		header = []string{stateKey, "tenant", "expiration", "plan"}
		data = append(data, []interface{}{"expired", l.Subscription.Tenant, expiration(l.Subscription.ExpiryTimestampMs), l.Subscription.Plan})
	case *stackstate_api.LicensedSubscription:
		header = []string{stateKey, "tenant", "expiration", "plan"}
		data = append(data, []interface{}{"valid", l.Subscription.Tenant, expiration(l.Subscription.ExpiryTimestampMs), l.Subscription.Plan})
	case *stackstate_api.UnlicensedSubscription:
		header = []string{stateKey, "reason"}
		data = append(data, []interface{}{"unlicensed", l.Reason})
	}

	return printer.TableData{
		Header:              header,
		Data:                data,
		MissingTableDataMsg: printer.NotFoundMsg{Types: "license"},
	}
}
