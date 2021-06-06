package controller

import (
	"encoding/json"
	"fmt"

	"github.com/XrayR-project/XrayR/api"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/infra/conf"
)

//OutboundBuilder build freedom outbund config for addoutbound
func OutboundBuilder(config *Config, nodeInfo *api.NodeInfo) (*core.OutboundHandlerConfig, error) {
	outboundDetourConfig := &conf.OutboundDetourConfig{}
	outboundDetourConfig.Protocol = "freedom"
	outboundDetourConfig.Tag = fmt.Sprintf("%s_%d", nodeInfo.NodeType, nodeInfo.Port)

	// build send ip address
	if config.SendIP != "" {
		ipAddress := net.ParseAddress(config.ListenIP)
		outBoundConfig.SendThrough = &conf.Address{ipAddress}
	}

	// freedom protocol setting
	var DomainStrategy string = "Asis"
	if config.EnableDNS {
		if config.DNSType != "" {
			DomainStrategy = config.DNSType
		} else {
			DomainStrategy = "UseIP"
	    }
	}
	proxySetting := &conf.FreedomConfig{
		DomainStrategy: DomainStrategy,
	}
	var setting json.RawMessage
	setting, err := json.Marshal(proxySetting)
	if err != nil {
		return nil, fmt.Errorf("Marshal proxy %s config fialed: %s", nodeInfo.NodeType, err)
	}
	outboundDetourConfig.Settings = &setting
	return outboundDetourConfig.Build()
}
