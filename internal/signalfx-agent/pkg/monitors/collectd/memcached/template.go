//go:build linux
// +build linux

package memcached

// AUTOGENERATED BY scripts/collectd-template-to-go.  DO NOT EDIT!!

import (
	"text/template"

	"github.com/signalfx/signalfx-agent/pkg/monitors/collectd"
)

// CollectdTemplate is a template for a memcached collectd config file
var CollectdTemplate = template.Must(collectd.InjectTemplateFuncs(template.New("memcached")).Parse(`
<LoadPlugin "memcached">
  Interval {{.IntervalSeconds}}
</LoadPlugin>
<Plugin "memcached">
  <Instance "{{.Name}}[monitorID={{.MonitorID}}]">
    ReportHost {{toBool .ReportHost}}
    Host "{{.Host}}"
    Port "{{.Port}}"
  </Instance>
</Plugin>
`)).Option("missingkey=error")
