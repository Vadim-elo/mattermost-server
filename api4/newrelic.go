package api4

import (
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (api *API) InitNewrelic() {
	appNewrelic, err := newrelic.NewApplication(
		newrelic.ConfigAppName("TestingMM"),
		newrelic.ConfigLicense("eu01xxc1419f805697352db19f20d69ffa63NRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		mlog.Warn("Something went wrong on appNewrelic", mlog.Err(err))
	}
	api.newrelic = appNewrelic
}
