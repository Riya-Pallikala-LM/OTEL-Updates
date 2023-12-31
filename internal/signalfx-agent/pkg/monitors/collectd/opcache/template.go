//go:build linux
// +build linux

package opcache

// AUTOGENERATED BY scripts/collectd-template-to-go.  DO NOT EDIT!!

import (
	"text/template"

	"github.com/signalfx/signalfx-agent/pkg/monitors/collectd"
)

// CollectdTemplate is a template for a opcache collectd config file
var CollectdTemplate = template.Must(collectd.InjectTemplateFuncs(template.New("opcache")).Parse(`
<LoadPlugin "curl_json">
  Interval {{.IntervalSeconds}}
</LoadPlugin>
<Plugin "curl_json">
  <URL "{{renderValue .URL (toMap $) }}">
    Instance "{{.Name}}[monitorID={{.MonitorID}}]"
    <Key "memory_usage/used_memory">
      Type "cache_size"
    </Key>
    <Key "memory_usage/free_memory">
      Type "cache_size"
    </Key>
    <Key "opcache_statistics/opcache_hit_rate">
      Type "cache_ratio"
    </Key>
    <Key "opcache_statistics/misses">
      Type "cache_result"
    </Key>
    <Key "opcache_statistics/hits">
      Type "cache_result"
    </Key>
    <Key "opcache_statistics/num_cached_keys">
      Type "files"
    </Key>
    <Key "opcache_statistics/max_cached_keys">
      Type "files"
    </Key>
    <Key "opcache_statistics/num_cached_scripts">
      Type "files"
    </Key>
  </URL>
</Plugin>
`)).Option("missingkey=error")
