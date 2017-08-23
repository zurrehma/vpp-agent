package vpp

import (
	"github.com/ligato/cn-infra/core"
	"github.com/ligato/cn-infra/datasync/resync"
	"github.com/ligato/cn-infra/flavors/etcdkafka"
	"github.com/ligato/vpp-agent/plugins/defaultplugins"
	"github.com/ligato/vpp-agent/plugins/govppmux"
	"github.com/ligato/vpp-agent/plugins/linuxplugin"
)

// Flavor glues together multiple plugins to translate ETCD configuration into VPP.
type Flavor struct {
	Base   etcdkafka.Flavor
	Resync resync.Plugin
	GoVPP  govppmux.GOVPPPlugin
	VPP    defaultplugins.Plugin
	Linux  linuxplugin.Plugin

	injected bool
}

// Inject sets object references
func (f *Flavor) Inject() error {
	if f.injected {
		return nil
	}
	f.injected = true

	f.Base.Inject()

	f.GoVPP.Deps.PluginInfraDeps = *f.Base.FlavorLocal.InfraDeps("govpp")
	f.VPP.Deps.PluginInfraDeps = *f.Base.FlavorLocal.InfraDeps("default-plugins")
	f.VPP.Deps.Publish = &f.Base.ETCDDataSync
	f.VPP.Deps.Watch = &f.Base.ETCDDataSync
	f.VPP.Deps.Kafka = &f.Base.Kafka
	f.VPP.Deps.GoVppmux = &f.GoVPP
	f.VPP.Deps.Linux = &f.Linux
	f.VPP.Linux.Deps.Watcher = &f.Base.ETCDDataSync

	return nil
}

// Plugins combines Generic Plugins and Standard VPP Plugins + (their ETCD Connector/Adapter with RESYNC)
func (f *Flavor) Plugins() []*core.NamedPlugin {
	f.Inject()
	return core.ListPluginsInFlavor(f)
}