package nmap

import "github.com/arielril/attack-graph-flow-code-runner/internal/model"

func ApiExecuteToInternalStruct(apiEx model.ApiExecuteNmap) *model.NmapOptions {
	r := &model.NmapOptions{}

	r.Target = apiEx.Target

	if apiEx.Options != nil {
		r.ServiceVersion = apiEx.Options.ServiceVersion
		r.HostsOnline = apiEx.Options.HostsOnline
		r.PortRange = apiEx.Options.PortRange
		r.DefaultScripts = apiEx.Options.DefaultScripts
		r.SynScan = apiEx.Options.SynScan
		r.UdpScan = apiEx.Options.UdpScan
	}

	return r
}
