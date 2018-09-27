package model

import "github.com/dihedron/go-bool/logic"

// OracleJDK8SupportFilter represents the Oracle JDK 8 support matrix; it is
// certified on certain Guest OSs, provided they all run on a very limited set
// of hypervisors; see https://www.oracle.com/technetwork/java/javase/certconfig-2095354.html
// Supported Hypervisors are:
// - Oracle VM
// - VirtualBox 3.x and 4.x
// - Solaris Containers
// - Solaris LDOMs
// - Microsoft Hyper-V runninng on Windows Server 2012
var OracleJDK8SupportMatrix = logic.And(
	JDK(OracleJDK8),
	logic.Or(
		logic.And(
			HostOS(OracleLinux5x, OracleLinux6x, OracleLinux7x),
			Hypervisor(KVM),
		),
		logic.And(
			HostOS(Win2k12, Win2k12R2),
			Hypervisor(HyperV),
		),
	),
	logic.Or(
		GuestOS(Win2k8R2VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		GuestOS(OracleLinux5xVM, OracleLinux6xVM, OracleLinux7xVM),
		GuestOS(RHEL5xVM, RHEL6xVM, RHEL7xVM),
		GuestOS(SLES10VM, SLES11VM, SLES12VM),
		GuestOS(Ubuntu1604VM),
	),
)

var OpenJDK8SupportMatrix = logic.And(
	JDK(OpenJDK8),
	GuestOS(RHEL5xVM, RHEL6xVM, RHEL7xVM),
)

// see https://www.ibm.com/support/knowledgecenter/SSYKE2_8.0.0/com.ibm.java.80.doc/user/supported_env_80.html
var IBMJDK8SupportMatrix = logic.And(
	JDK(IBMJDK8),
	logic.Or(
		GuestOS(SLES11VM, SLES12VM, RHEL6xVM, RHEL7xVM, Ubuntu1204VM, Ubuntu1404VM, Ubuntu1604VM, Ubuntu1804VM),
		GuestOS(Win2k12VM, Win2k12R2VM, Win2k16VM),
	),
)
