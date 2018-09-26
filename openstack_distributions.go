package main

import "github.com/dihedron/go-bool/rules"

// OpenStack distributions and the hypervisor technologies they support, according
// to the support matrix at https://www.openstack.org/marketplace/distros/
var supportedDistributions = rules.Or(
	rules.And(
		Distribution(RedHat),
		HostOS(RHEL6x, RHEL7x),
		Hypervisor(KVM),
	),
	rules.And(
		Distribution(Canonical),
		rules.Or(
			rules.And(
				HostOS(Ubuntu1604LTS, Ubuntu1804LTS),
				Hypervisor(KVM, LXC),
			),
			rules.And(
				HostOS(BareMetal),
				Hypervisor(ESXi),
			),
		),
	),
	rules.And(
		Distribution(SuSE),
		rules.Or(
			rules.And(
				HostOS(SLES11, SLES12),
				Hypervisor(KVM, Xen),
			),
			rules.And(
				HostOS(BareMetal),
				Hypervisor(ESXi),
			),
		),
	),
	rules.And(
		Distribution(Oracle),
		rules.Or(
			rules.And(
				HostOS(OracleLinux6x, OracleLinux7x),
				Hypervisor(KVM),
			),
			rules.And(
				HostOS(BareMetal, Win2k12R2, Win2k16),
				Hypervisor(HyperV),
			),
		),
	),
	rules.And(
		Distribution(VMware),
		HostOS(BareMetal),
		Hypervisor(ESXi),
	),
	rules.And(
		Distribution(Huawei),
		rules.Or(
			rules.And(
				Hypervisor(KVM, Xen),
				HostOS(SLES11, SLES12, Ubuntu1204LTS, CentOS6x),
			),
			rules.And(
				HostOS(BareMetal),
				Hypervisor(ESXi),
			),
		),
	),
)
