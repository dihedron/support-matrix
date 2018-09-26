package main

import "github.com/dihedron/go-bool/rules"

// Oracle JDK 1.8 is certified on certain Guest OSs, provided they all run on
// a very limited set of hypervisors; see https://www.oracle.com/technetwork/java/javase/certconfig-2095354.html
// Supported Hypervisors are:
// - Oracle VM
// - VirtualBox 3.x and 4.x
// - Solaris Containers
// - Solaris LDOMs
// - Microsoft Hyper-V runninng on Windows Server 2012

var supportedByOracleJDK18 = rules.And(
	JDK(OracleJDK8),
	rules.Or(
		rules.And(
			HostOS(OracleLinux5x, OracleLinux6x, OracleLinux7x),
			Hypervisor(KVM),
		),
		rules.And(
			HostOS(Win2k12, Win2k12R2),
			Hypervisor(HyperV),
		),
	),
	rules.Or(
		GuestOS(Win2k8R2VM, Win2k12VM, Win2k12R2VM, Win2k16VM),
		GuestOS(OracleLinux5xVM, OracleLinux6xVM, OracleLinux7xVM),
		GuestOS(RHEL5xVM, RHEL6xVM, RHEL7xVM),
		GuestOS(SLES10VM, SLES11VM, SLES12VM),
		GuestOS(Ubuntu1604VM),
	),
)

var supportedByOpenJDK = rules.And(
	GuestOS(RHEL5xVM, RHEL6xVM, RHEL7xVM),
)
