package model

import "github.com/dihedron/go-bool/logic"

var JBossEAP7xSupportFilter = logic.And(
	AppServer(JBossEAP70, JBossEAP71),
	logic.Or(
		logic.And(
			GuestOS(RHEL7xVM),
			JDK(OpenJDK8),
		),
		logic.And(
			GuestOS(RHEL6xVM),
			JDK(OpenJDK8),
		),
		logic.And(
			GuestOS(RHEL5xVM),
			JDK(OpenJDK8),
		),
		OracleJDK8SupportFilter,
		//TODO: supportedByIBMJDK8,
	),
)

var WebSphere8xSupportFilter = logic.And(
	AppServer(WebSphere8x),
	logic.Or(
		logic.And(
			HostOS(Win2k8, Win2k12, Win2k12R2, Win2k16),
			Hypervisor(HyperV),
			GuestOS(RHEL6xVM, SLES11VM, SLES12VM, Ubuntu1404VM, Win2k8VM),
		),
		logic.And(
			HostOS(RHEL6x, RHEL7x),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k8VM, Win2k8R2VM),
		),
		logic.And(
			HostOS(SLES11),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k8VM, Win2k8R2VM),
		),
		logic.And(
			HostOS(BareMetal),
			Hypervisor(ESXi),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k8VM, Win2k8R2VM),
		),
	),
)

var WebSphere9xSupportFilter = logic.And(
	AppServer(WebSphere9x),
	logic.Or(
		logic.And(
			HostOS(SLES11),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Ubuntu1804VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		logic.And(
			HostOS(SLES12),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		logic.And(
			HostOS(Ubuntu1604LTS),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		logic.And(
			HostOS(Win2k12, Win2k12R2, Win2k16),
			Hypervisor(HyperV),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		logic.And(
			HostOS(OracleLinux5x, OracleLinux6x, OracleLinux7x),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		logic.And(
			HostOS(RHEL6x),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Ubuntu1804VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		logic.And(
			HostOS(RHEL7x),
			Hypervisor(KVM),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
		logic.And(
			HostOS(BareMetal),
			Hypervisor(ESXi),
			GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		),
	),
)

var WebLogic11gSupportFilter = logic.Or(
	logic.And(
		HostOS(SLES11),
		Hypervisor(KVM),
		GuestOS(RHEL6xVM, SLES11VM, Ubuntu1404VM, SLES12VM, RHEL7xVM, Ubuntu1604VM, Ubuntu1804VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
	),
)
