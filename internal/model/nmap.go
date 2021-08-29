package model

type ApiExecuteNmap struct {
	Target  string                 `json:"target"`
	Options *ApiExecuteNmapOptions `json:"options"`
}

type ApiExecuteNmapOptions struct {
	ServiceVersion bool   `json:"service_version" default:"true"`
	HostsOnline    bool   `json:"hosts_online" default:"true"`
	PortRange      string `json:"port_range" default:"1-65535"`
	DefaultScripts bool   `json:"default_scripts" default:"false"`
	SynScan        bool   `json:"syn_scan" default:"false"`
	UdpScan        bool   `json:"udp_scan" default:"false"`
}

type ApiExecuteNmap200Response struct {
	Target string                            `json:"target"`
	Ports  []*ApiExecuteNmap200PortsResponse `json:"ports"`
}

type ApiExecuteNmap200PortsResponse struct {
	Number   int    `json:"number,omitempty"`
	State    string `json:"state,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Service  string `json:"service,omitempty"`
	RpcInfo  string `json:"rpc_info,omitempty"`
	Version  string `json:"version,omitempty"`
}

type NmapOptions struct {
	Target         string
	ServiceVersion bool
	HostsOnline    bool
	PortRange      string
	DefaultScripts bool
	SynScan        bool
	UdpScan        bool
}
