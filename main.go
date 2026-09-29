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
			Usage:     "Name of the Windows service to check",
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
		return sensu.CheckStateWarning, errors.New("--service flag or CHECK_SERVICE environment variable is required")
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
	case svc.Running:
		fmt.Printf("OK: %s running\n", plugin.Service)
		return sensu.CheckStateOK, nil
	case svc.Stopped:
		fmt.Printf("CRITICAL: %s stopped\n", plugin.Service)
		return sensu.CheckStateCritical, nil
	case svc.StartPending, svc.StopPending, svc.ContinuePending, svc.PausePending, svc.Paused:
		fmt.Printf("WARNING: %s %s\n", plugin.Service, stateNames[statusCode.State])
		return sensu.CheckStateWarning, nil
	}
	fmt.Printf("UNKNOWN: %s in unexpected state %d\n", plugin.Service, statusCode.State)
	return sensu.CheckStateUnknown, nil
}

var stateNames = map[svc.State]string{
	svc.StartPending:    "start pending",
	svc.StopPending:     "stop pending",
	svc.ContinuePending: "continue pending",
	svc.PausePending:    "pause pending",
	svc.Paused:          "paused",
}
