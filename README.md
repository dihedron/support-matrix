# Products Support Matrix

This tool creates the Cartesian product of the domain values for:

- OpenStack distribution, 
- host operating systems, 
- hypervisor technology, 
- guest operating system, 
- Java Development Kit, and 
- application server

and then applies any filters in order to find out the viable combinations. Each piece of software (e.g. the OpenStack distribution, or a JDK) has its support matrix which specifies where it can be run safely and with full support from its vendor, expressed as a boolean expression of values from the aforementioned domains.

To make a simple example, if we wanted to preresents the support matrix for the OpenJDK, which only supports Red Hat Enterprise Linux running as a Virtual Machine inside a KVM , we would produce an expression like:
```golang
matrix := And(
	JDK(OpenJDK8),
	GuestOS(RHEL7x, RHEL6x),
)
```
These expressions can be created by combining simple Boolean operators: `And`, `Or` and `Not`.
