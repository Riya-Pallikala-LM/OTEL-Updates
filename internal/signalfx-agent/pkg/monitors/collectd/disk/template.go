//go:build linux
// +build linux

package disk

// AUTOGENERATED BY scripts/collectd-template-to-go.  DO NOT EDIT!!

import (
	"text/template"

	"github.com/signalfx/signalfx-agent/pkg/monitors/collectd"
)

// CollectdTemplate is a template for a disk collectd config file
var CollectdTemplate = template.Must(collectd.InjectTemplateFuncs(template.New("disk")).Parse(`
LoadPlugin disk
<Plugin "disk">
{{range .Disks -}}
  Disk "{{.}}"
{{end}}
  IgnoreSelected {{if .IgnoreSelected}}true{{else}}false{{end}}
</Plugin>


<Chain "PostCache"> 
  <Rule "set_disk_monitor_id"> 
    <Match "regex"> 
      Plugin "^disk$" 
    </Match> 
    <Target "set"> 
      MetaData "monitorID" "{{.MonitorID}}" 
    </Target> 
  </Rule> 
</Chain>
`)).Option("missingkey=error")
