package model

import "github.com/dihedron/go-bool/logic"

// DistributionType is the type used to represent Distributions.
type DistributionType string

// DistributionOperand is the internal representation of the Distribution.
type DistributionOperand struct {
	Values []DistributionType
}

// Evaluate checks if the current Distribution is one of those supported.
func (o DistributionOperand) Evaluate(ctx interface{}) (bool, error) {
	for _, value := range o.Values {
		if value == (ctx.(Stack)).Distribution {
			return true, nil
		}
	}
	return false, nil
}

// Distribution creates a new HypervisorOperand with the given set of acceptable values.
func Distribution(values ...DistributionType) DistributionOperand {
	return DistributionOperand{
		values,
	}
}

// OpenStack distributions.
const (
	RedHat    DistributionType = "Red Hat"
	Canonical                  = "Canonical"
	SuSE                       = "SuSE"
	Oracle                     = "Oracle"
	Huawei                     = "Huawei"
	VMware                     = "VMware"
)

// AllDistributions is the domain of openStack distributions.
var AllDistributions = []DistributionType{
	RedHat,
	Canonical,
	SuSE,
	Oracle,
	Huawei,
	VMware,
}

// HostOSType is the type used to represent host operating systems.
type HostOSType string

// HostOSOperand is the internal representation of the Host OS.
type HostOSOperand struct {
	Values []HostOSType
}

// Evaluate checks if the current HostOS is one of those supported.
func (o HostOSOperand) Evaluate(ctx interface{}) (bool, error) {
	for _, value := range o.Values {
		if value == (ctx.(Stack)).HostOS {
			return true, nil
		}
	}
	return false, nil
}

// HostOS creates a new HostOSOperand with the given set of acceptable values.
func HostOS(values ...HostOSType) HostOSOperand {
	return HostOSOperand{
		values,
	}
}

// Host OSs
const (
	BareMetal     HostOSType = "Bare Metal"
	Ubuntu1204LTS            = "Ubuntu 12.04 LTS"
	Ubuntu1404LTS            = "Ubuntu 14.04 LTS"
	Ubuntu1604LTS            = "Ubuntu 16.04 LTS"
	Ubuntu1804LTS            = "Ubuntu 18.04 LTS"
	RHEL5x                   = "Red Hat Enterprise Linux 5.x"
	RHEL6x                   = "Red Hat Enterprise Linux 6.x"
	RHEL7x                   = "Red Hat Enterprise Linux 7.x"
	OracleLinux5x            = "Oracle Linux 5.5"
	OracleLinux6x            = "Oracle Linux 6.x"
	OracleLinux7x            = "Oracle Linux 7.x"
	Win2k8                   = "Windows Server 2008"
	Win2k8R2                 = "Windows Server 2008 R2"
	Win2k12                  = "Windows Server 2012"
	Win2k12R2                = "Windows Server 2012 R2"
	Win2k16                  = "Windows Server 2016"
	SLES10                   = "SuSE Enterprise Linux 10"
	SLES11                   = "SuSE Enterprise Linux 11"
	SLES12                   = "SuSE Enterprise Linux 12"
	CentOS5x                 = "CentOS Linux 5.x"
	CentOS6x                 = "CentOS Linux 6.x"
	CentOS7x                 = "CentOS Linux 7.x"
)

// AllHostOSs represents the set of all supported host operating systems.
var AllHostOSs = []HostOSType{
	BareMetal,
	Ubuntu1204LTS,
	Ubuntu1404LTS,
	Ubuntu1604LTS,
	Ubuntu1804LTS,
	RHEL5x,
	RHEL6x,
	RHEL7x,
	OracleLinux5x,
	OracleLinux6x,
	OracleLinux7x,
	Win2k8,
	Win2k8R2,
	Win2k12,
	Win2k12R2,
	Win2k16,
	SLES10,
	SLES11,
	SLES12,
	CentOS5x,
	CentOS6x,
	CentOS7x,
}

// HypervisorType is the type used to represent Hypervisor technologies.
type HypervisorType string

// HypervisorOperand is the internal representation of the Hypervisor.
type HypervisorOperand struct {
	Values []HypervisorType
}

// Evaluate checks if the current Hypervisor is one of those supported.
func (o HypervisorOperand) Evaluate(ctx interface{}) (bool, error) {
	for _, value := range o.Values {
		if value == (ctx.(Stack)).Hypervisor {
			return true, nil
		}
	}
	return false, nil
}

// Hypervisor creates a new HypervisorOperand with the given set of acceptable values.
func Hypervisor(values ...HypervisorType) HypervisorOperand {
	return HypervisorOperand{
		values,
	}
}

// Hypervisor technologies
const (
	KVM    HypervisorType = "KVM"
	ESXi                  = "ESXi"
	LXC                   = "LXC"
	Xen                   = "Xen"
	HyperV                = "Hyper-V"
)

// AllHypervisors is the domain of all supported hypervisor technologies.
var AllHypervisors = []HypervisorType{
	KVM,
	ESXi,
	LXC,
	Xen,
	HyperV,
}

// AllHyperV is the set of all Hyper-V hypervisor technologies.
var AllHyperV = logic.And(
	Hypervisor(HyperV),
	HostOS(BareMetal, Win2k8, Win2k8R2, Win2k12, Win2k12R2, Win2k16),
)

// GuestOSType is the type used to represent guest operating systems.
type GuestOSType string

// GuestOSOperand is the internal representation of the Guest OS.
type GuestOSOperand struct {
	Values []GuestOSType
}

// Evaluate checks if the current GuestOS is one of those supported.
func (o GuestOSOperand) Evaluate(ctx interface{}) (bool, error) {
	for _, value := range o.Values {
		if value == (ctx.(Stack)).GuestOS {
			return true, nil
		}
	}
	return false, nil
}

// GuestOS creates a new GuestOSOperand with the given set of acceptable values.
func GuestOS(values ...GuestOSType) GuestOSOperand {
	return GuestOSOperand{
		values,
	}
}

// Guest OSs
const (
	Ubuntu1204VM    GuestOSType = "Ubuntu 12.04 LTS (VM)"
	Ubuntu1404VM                = "Ubuntu 14.04 LTS (VM)"
	Ubuntu1604VM                = "Ubuntu 16.04 LTS (VM)"
	Ubuntu1804VM                = "Ubuntu 18.04 LTS (VM)"
	RHEL5xVM                    = "Red Hat Enterprise Linux 5.x (VM)"
	RHEL6xVM                    = "Red Hat Enterprise Linux 6.x (VM)"
	RHEL7xVM                    = "Red Hat Enterprise Linux 7.x (VM)"
	OracleLinux5xVM             = "Oracle Linux 5.x (VM)"
	OracleLinux6xVM             = "Oracle Linux 6.x (VM)"
	OracleLinux7xVM             = "Oracle Linux 7.x (VM)"
	Win2k8VM                    = "Windows Server 2008 (VM)"
	Win2k8R2VM                  = "Windows Server 2008 R2 (VM)"
	Win2k12VM                   = "Windows Server 2012 (VM)"
	Win2k12R2VM                 = "Windows Server 2012 R2 (VM)"
	Win2k16VM                   = "Windows Server 2016 (VM)"
	SLES10VM                    = "SuSE Enterprise Linux 10 (VM)"
	SLES11VM                    = "SuSE Enterprise Linux 11 (VM)"
	SLES12VM                    = "SuSE Enterprise Linux 12 (VM)"
	CentOS5xVM                  = "CentOS Linux 5.x (VM)"
	CentOS6xVM                  = "CentOS Linux 6.x (VM)"
	CentOS7xVM                  = "CentOS Linux 7.x (VM)"
)

// AllGuestOSs represents the set of all supported guest operating systems.
var AllGuestOSs = []GuestOSType{
	Ubuntu1204VM,
	Ubuntu1404VM,
	Ubuntu1604VM,
	Ubuntu1804VM,
	RHEL5xVM,
	RHEL6xVM,
	RHEL7xVM,
	OracleLinux5xVM,
	OracleLinux6xVM,
	OracleLinux7xVM,
	Win2k8VM,
	Win2k8R2VM,
	Win2k12VM,
	Win2k12R2VM,
	Win2k16VM,
	SLES10VM,
	SLES11VM,
	SLES12VM,
	CentOS5xVM,
	CentOS6xVM,
	CentOS7xVM,
}

// JDKType is the type used to represent the Java Developer Kit.
type JDKType string

// JDKOperand is the internal representation of the JDK.
type JDKOperand struct {
	Values []JDKType
}

// Evaluate checks if the current JDK is one of those supported.
func (o JDKOperand) Evaluate(ctx interface{}) (bool, error) {
	for _, value := range o.Values {
		if value == (ctx.(Stack)).JDK {
			return true, nil
		}
	}
	return false, nil
}

// JDK creates a new JDKOperand with the given set of acceptable values.
func JDK(values ...JDKType) JDKOperand {
	return JDKOperand{
		values,
	}
}

// JDKs
const (
	OpenJDK6    = "OpenJDK 1.6"
	OpenJDK7    = "OpenJDK 1.7"
	OpenJDK8    = "OpenJDK 1.8"
	OpenJDK9    = "OpenJDK 1.9"
	OpenJDK10   = "OpenJDK 1.10"
	OpenJDK11   = "OpenJDK 1.11"
	OracleJDK6  = "Oracle JDK 1.6"
	OracleJDK7  = "Oracle JDK 1.7"
	OracleJDK8  = "Oracle JDK 1.8"
	OracleJDK9  = "Oracle JDK 1.9"
	OracleJDK10 = "Oracle JDK 1.10"
	OracleJDK11 = "Oracle JDK 1.11"
	IBMJDK6     = "IBM JDK 1.6"
	IBMJDK7     = "IBM JDK 1.7"
	IBMJDK8     = "IBM JDK 1.8"
	IBMJDK9     = "IBM JDK 1.9"
	IBMJDK10    = "IBM JDK 1.10"
	IBMJDK11    = "IBM JDK 1.11"
)

// AllJDKs represents the domain of all existing JDKs.
var AllJDKs = []JDKType{
	OpenJDK6,
	OpenJDK7,
	OpenJDK8,
	OpenJDK9,
	OpenJDK10,
	OpenJDK11,
	OracleJDK6,
	OracleJDK7,
	OracleJDK8,
	OracleJDK9,
	OracleJDK10,
	OracleJDK11,
	IBMJDK6,
	IBMJDK7,
	IBMJDK8,
	IBMJDK9,
	IBMJDK10,
	IBMJDK11,
}

// AppServerType is the type used to represent application servers.
type AppServerType string

// AppServerOperand is the internal representation of the AppServer.
type AppServerOperand struct {
	Values []AppServerType
}

// Evaluate checks if the current AppServer is one of those supported.
func (o AppServerOperand) Evaluate(ctx interface{}) (bool, error) {
	for _, value := range o.Values {
		if value == (ctx.(Stack)).AppServer {
			return true, nil
		}
	}
	return false, nil
}

// AppServer creates a new AppServerOperand with the given set of acceptable values.
func AppServer(values ...AppServerType) AppServerOperand {
	return AppServerOperand{
		values,
	}
}

// Application Server
const (
	JBossEAP70  AppServerType = "JBoss EAP 7.0"
	JBossEAP71                = "JBoss EAP 7.1"
	WebSphere8x               = "IBM WebSphere 8"
	WebSphere9x               = "IBM WebSphere 9"
	WebLogic11g               = "Oracle WebLogic 11g"
	WebLogic12c               = "Oracle WebLogic 12c"
)

// AllAppServers represents the domain of all application servers.
var AllAppServers = []AppServerType{
	JBossEAP71,
	JBossEAP70,
	WebSphere8x,
	WebSphere9x,
	WebLogic11g,
	WebLogic12c,
}

// Stack represents a combination of Distribution, HostOS, Hypervisor, GuestOS,
// JDK and AppServer to be evaluated against a filter.
type Stack struct {
	Distribution DistributionType
	HostOS       HostOSType
	Hypervisor   HypervisorType
	GuestOS      GuestOSType
	JDK          JDKType
	AppServer    AppServerType
}
