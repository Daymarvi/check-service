package main

import (
	"fmt"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	Example string
	Service string
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "check-service",
			Short:    "check for windows service aliveness via mgr",
			Keyspace: "sensu.io/plugins/check_name/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		{
			Path:      "service",
			Env:       "CHECK_SERVICE",
			Argument:  "service",
			Shorthand: "s",
			Default:   "",
			Usage:     "Expected service status",
			Value:     &plugin.Service,
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	if len(plugin.Service) == 0 {
		return sensu.CheckStateWarning, fmt.Errorf("--service environment variable is required")
	}
	return sensu.CheckStateOK, nil
}

func executeCheck(event *types.Event) (int, error) {
	m, err := mgr.Connect()
	if err != nil {
		return sensu.CheckStateUnknown, fmt.Errorf("failed to connect to service manager: %v", err)
	}
	s, err := m.OpenService(plugin.Service)
	if err != nil {
		return sensu.CheckStateUnknown, fmt.Errorf("could not access service: %v", err)
	}
	defer s.Close()
	statusCode, err := s.Query()
	if err != nil {
		return sensu.CheckStateUnknown, fmt.Errorf("failed to query to service manager: %v", err)
	}
	switch statusCode.State {
	case svc.Stopped:
		fmt.Printf("CRITICAL: %s stopped", plugin.Service)
		return sensu.CheckStateCritical, nil
	case svc.Running:
		fmt.Printf("OK: %s Running.", plugin.Service)
		return sensu.CheckStateOK, nil
	}
	return sensu.CheckStateUnknown, nil
}
