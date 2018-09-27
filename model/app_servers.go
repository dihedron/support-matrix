package model

import (
	"github.com/dihedron/go-bool/logic"
)

var JBossEAP7xSupportMatrix = logic.And(
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
		OracleJDK8SupportMatrix,
		//TODO: supportedByIBMJDK8,
	),
)

// see https://www.ibm.com/software/reports/compatibility/clarity-reports/report/html/softwareReqsForProduct?deliverableId=1337870535828&osPlatform=Linux
// and https://www.ibm.com/software/reports/compatibility/clarity-reports/report/html/softwareReqsForProduct?deliverableId=1337870535828&osPlatform=Windows
var WebSphere8xBaseSupportMatrix = logic.And(
	AppServer(WebSphere8xBase),
	// logic.Or(
	// 	logic.And(
	// 		AllLinuxGuests,
	// 		JDK(IBMJDK8, OracleJDK6, OracleJDK7, OracleJDK8, OpenJDK6, OpenJDK7, OpenJDK8),
	// 	),
	// 	logic.And(
	// 		GuestOS(Win2k16VM),
	// 		JDK(IBMJDK8),
	// 	),
	// ),
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

// see https://www.ibm.com/software/reports/compatibility/clarity-reports/report/html/softwareReqsForProduct?deliverableId=1345529544358&osPlatform=Linux
// and https://www.ibm.com/software/reports/compatibility/clarity-reports/report/html/softwareReqsForProduct?deliverableId=1345529544358&osPlatform=Windows
// TODO: check operating system levels
var WebSphere8xNetDeploySupportMatrix = logic.And(
	AppServer(WebSphere8xND),
	logic.Or(
		logic.And(
			AllLinuxGuests,
			JDK(IBMJDK8, OracleJDK6, OracleJDK7, OracleJDK8, OpenJDK6, OpenJDK7, OpenJDK8),
		),
		logic.And(
			GuestOS(Win2k16VM),
			JDK(IBMJDK8),
		),
	),
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

// see https://www.ibm.com/software/reports/compatibility/clarity-reports/report/html/osForProduct?deliverableId=1374008815765&osPlatforms=Linux&duComponentIds=S000
// and https://www.ibm.com/software/reports/compatibility/clarity-reports/report/html/osForProduct?deliverableId=1374008815765&osPlatforms=Windows&duComponentIds=S000
var WebSphere9xBaseSupportMatrix = logic.And(
	AppServer(WebSphere9xBase),
	JDK(IBMJDK8),
	//IBMJDK8SupportFilter, ??
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

// see https://www.ibm.com/software/reports/compatibility/clarity-reports/report/html/osForProduct?deliverableId=1380722427848&osPlatforms=Linux&duComponentIds=S000
// and https://www.ibm.com/software/reports/compatibility/clarity-reports/report/html/osForProduct?deliverableId=1380722427848&osPlatforms=Windows&duComponentIds=S000
// TODO: check operating system levels
var WebSphere9xNetDeploySupportMatrix = logic.And(
	AppServer(WebSphere9xND),
	JDK(IBMJDK8),
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

var WebLogic12cSupportMatrix = logic.And(
	AppServer(WebLogic12c),
	logic.Or(
		logic.And(
			HostOS(Win2k12, Win2k12R2, Win2k16),
			Hypervisor(HyperV),
		),
		logic.And(
			HostOS(OracleLinux5x, OracleLinux6x, OracleLinux7x),
			Hypervisor(KVM),
		),
		//HostOS(model.BareMetal),
	),
	logic.Or(
		GuestOS(OracleLinux6xVM, OracleLinux7xVM),
		GuestOS(SLES11VM, SLES12VM),
		GuestOS(RHEL6xVM, RHEL7xVM),
		GuestOS(Win2k12VM, Win2k12R2VM, Win2k16VM),
	),
	JDK(OracleJDK8),
)

var WebLogic11gSupportMatrix = logic.And(
	AppServer(WebLogic11g),
	logic.Or(
		logic.And(
			HostOS(Win2k12, Win2k12R2),
			Hypervisor(HyperV),
		),
		logic.And(
			HostOS(OracleLinux5x, OracleLinux6x, OracleLinux7x),
			Hypervisor(KVM),
		),
		//HostOS(model.BareMetal),
	),
	logic.Or(
		logic.And(
			GuestOS(RHEL5xVM, SLES11VM, Win2k8VM, Win2k8R2VM),
			JDK(OracleJDK6),
		),
		logic.And(
			GuestOS(RHEL6xVM, RHEL7xVM, OracleLinux6xVM, OracleLinux7xVM, SLES12VM, Win2k8R2VM),
			JDK(OracleJDK7),
		),
	),
)
