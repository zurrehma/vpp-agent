// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linuxcalls

import (
	"net"

	"github.com/ligato/cn-infra/logging/measure"
	"github.com/vishvananda/netlink"
)

// NetlinkAPI interface covers all methods inside linux calls package needed to manage linux interfaces.
type NetlinkAPI interface {
	// AddVethInterfacePair configures two connected VETH interfaces
	AddVethInterfacePair(ifName, peerIfName string) error
	// DelVethInterfacePair removes VETH pair
	DelVethInterfacePair(ifName, peerIfName string) error
	// SetInterfaceUp sets interface state to 'up'
	SetInterfaceUp(ifName string) error
	// SetInterfaceDown sets interface state to 'down'
	SetInterfaceDown(ifName string) error
	// AddInterfaceIP adds new IP address
	AddInterfaceIP(ifName string, addr *net.IPNet) error
	// DelInterfaceIP removes IP address from linux interface
	DelInterfaceIP(ifName string, addr *net.IPNet) error
	// SetInterfaceMac sets MAC address
	SetInterfaceMac(ifName string, macAddress string) error
	// SetInterfaceMTU set maximum transmission unit for interface
	SetInterfaceMTU(ifName string, mtu int) error
	// RenameInterface changes interface host name
	RenameInterface(ifName string, newName string) error
	// GetLinkByName returns netlink interface type
	GetLinkByName(ifName string) (netlink.Link, error)
	// GetLinkList return all links from namespace
	GetLinkList() ([]netlink.Link, error)
	// GetAddressList reads all IP addresses
	GetAddressList(ifName string) ([]netlink.Addr, error)
	// InterfaceExists verifies interface existence
	InterfaceExists(ifName string) (bool, error)
	// GetInterfaceType returns linux interface type
	GetInterfaceType(ifName string) (string, error)
	// GetVethPeerName returns VETH's peer name
	GetVethPeerName(ifName string) (string, error)
	// GetInterfaceByName returns *net.Interface type from name
	GetInterfaceByName(ifName string) (*net.Interface, error)
}

// NetLinkHandler is accessor for netlink methods
type NetLinkHandler struct {
	stopwatch *measure.Stopwatch
}

// NewNetLinkHandler creates new instance of netlink handler
func NewNetLinkHandler(stopwatch *measure.Stopwatch) *NetLinkHandler {
	return &NetLinkHandler{
		stopwatch: stopwatch,
	}
}
