package main

import (
	"errors"
	"fmt"

	corev2 "github.com/sensu/core/v2"
	"github.com/sensu/sensu-plugin-sdk/sensu"
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

	options = []sensu.ConfigOption{
		&sensu.PluginConfigOption[string]{
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
	check := sensu.NewCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *corev2.Event) (int, error) {
	if plugin.Service == "" {
		return sensu.CheckStateWarning, errors.New("--service environment variable is required")
	}
	return sensu.CheckStateOK, nil
}

func executeCheck(event *corev2.Event) (int, error) {
	m, err := mgr.Connect()
	if err != nil {
		return sensu.CheckStateUnknown, fmt.Errorf("failed to connect to service manager: %w", err)
	}
	defer m.Disconnect()
	s, err := m.OpenService(plugin.Service)
	if err != nil {
		return sensu.CheckStateUnknown, fmt.Errorf("could not access service: %w", err)
	}
	defer s.Close()
	statusCode, err := s.Query()
	if err != nil {
		return sensu.CheckStateUnknown, fmt.Errorf("failed to query to service manager: %w", err)
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
