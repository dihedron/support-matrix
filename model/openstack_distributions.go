package model

import "github.com/dihedron/go-bool/logic"

// OpenStackDistributionsSupportFilter is the set of filters that select the combinations
// of stack elements supporting a set of feature-complete, recent and widely adopted
// OpenStack distributions; for the support matrix see
// https://www.openstack.org/marketplace/distros/
var OpenStackDistributionsSupportFilter = logic.Or(
	logic.And(
		Distribution(RedHat),
		HostOS(RHEL6x, RHEL7x),
		Hypervisor(KVM),
	),
	logic.And(
		Distribution(Canonical),
		logic.Or(
			logic.And(
				HostOS(Ubuntu1604LTS, Ubuntu1804LTS),
				Hypervisor(KVM, LXC),
			),
			logic.And(
				HostOS(BareMetal),
				Hypervisor(ESXi),
			),
		),
	),
	logic.And(
		Distribution(SuSE),
		logic.Or(
			logic.And(
				HostOS(SLES11, SLES12),
				Hypervisor(KVM, Xen),
			),
			logic.And(
				HostOS(BareMetal),
				Hypervisor(ESXi),
			),
		),
	),
	logic.And(
		Distribution(Oracle),
		logic.Or(
			logic.And(
				HostOS(OracleLinux6x, OracleLinux7x),
				Hypervisor(KVM),
			),
			logic.And(
				HostOS(BareMetal, Win2k12R2, Win2k16),
				Hypervisor(HyperV),
			),
		),
	),
	logic.And(
		Distribution(VMware),
		HostOS(BareMetal),
		Hypervisor(ESXi),
	),
	logic.And(
		Distribution(Huawei),
		logic.Or(
			logic.And(
				Hypervisor(KVM, Xen),
				HostOS(SLES11, SLES12, Ubuntu1204LTS, CentOS6x),
			),
			logic.And(
				HostOS(BareMetal),
				Hypervisor(ESXi),
			),
		),
	),
)
