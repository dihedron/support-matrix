package main

// Stack represents a combination of Distribution, Hypervisor, GuestOS, JDK
// and AppServer to be evaluated against a rule.
type Stack struct {
	Distribution string
	Hypervisor   string
	GuestOS      string
	JDK          string
	AppServer    string
}

// DistributionOperand is the internal representation of the Distribution.
type DistributionOperand struct {
	Values []string
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
func Distribution(values ...string) DistributionOperand {
	return DistributionOperand{
		values,
	}
}

// HypervisorOperand is the internal representation of the Hypervisor.
type HypervisorOperand struct {
	Values []string
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
func Hypervisor(values ...string) HypervisorOperand {
	return HypervisorOperand{
		values,
	}
}

// GuestOSOperand is the internal representation of the Guest OS.
type GuestOSOperand struct {
	Values []string
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
func GuestOS(values ...string) GuestOSOperand {
	return GuestOSOperand{
		values,
	}
}

// JDKOperand is the internal representation of the JDK.
type JDKOperand struct {
	Values []string
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
func JDK(values ...string) JDKOperand {
	return JDKOperand{
		values,
	}
}

// AppServerOperand is the internal representation of the AppServer.
type AppServerOperand struct {
	Values []string
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
func AppServer(values ...string) AppServerOperand {
	return AppServerOperand{
		values,
	}
}
