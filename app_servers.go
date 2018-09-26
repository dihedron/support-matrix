package main

import "github.com/dihedron/go-bool/rules"

var supportedByJBossEAP7x = rules.And(
	AppServer(JBossEAP70, JBossEAP71),
	rules.Or(
		rules.And(
			GuestOS(RHEL7xVM),
			JDK(OpenJDK8),
		),
		rules.And(
			GuestOS(RHEL6xVM),
			JDK(OpenJDK8),
		),
		rules.And(
			GuestOS(RHEL5xVM),
			JDK(OpenJDK8),
		),
		supportedByOracleJDK18,
		//TODO: supportedByIBMJDK8,
	),
)

var supportedByWebSphere8x = rules.And(
	AppServer(WebSphere8x),
	rules.Or(
		rules.And(
			HostOS(Win2k8, Win2k12, Win2k12R2, Win2k16),
			Hypervisor(HyperV),
			GuestOS(RHEL6xVM, SLES11VM, SLES12VM, Ubuntu1404VM, Win2k8VM),
		),
		rules.And(
			HostOS(RHEL6x, RHEL7x),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k8VM, Win2k8R2VM),
		),
		rules.And(
			HostOS(SLES11),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k8VM, Win2k8R2VM),
		),
		rules.And(
			HostOS(BareMetal),
			Hypervisor(ESXi),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k8VM, Win2k8R2VM),
		),
	),
)

var supportedByWebSphere9x = rules.And(
	AppServer(WebSphere9x),
	rules.Or(
		rules.And(
			HostOS(SLES11),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Ubuntu1804VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		rules.And(
			HostOS(SLES12),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		rules.And(
			HostOS(Ubuntu1604LTS),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		rules.And(
			HostOS(Win2k12, Win2k12R2, Win2k16),
			Hypervisor(HyperV),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		rules.And(
			HostOS(OracleLinux5x, OracleLinux6x, OracleLinux7x),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		rules.And(
			HostOS(RHEL6x),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Ubuntu1804VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		rules.And(
			HostOS(RHEL7x),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		rules.And(
			HostOS(BareMetal),
			Hypervisor(ESXi),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
	),
)

var supportedByWebLogic11g = rules.Or(
	rules.And(
		HostOS(SLES11),
		Hypervisor(KVM),
		GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Ubuntu1804VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
	),
)
