package rules

type Context struct {
	// Distribution   string
	// Virtualisation string
	Flavour   string
	GuestOS   string
	JDK       string
	AppServer string
}

type MatchOperand struct {
	Value string
}

func (o MatchOperand) Evaluate(ctx interface{}) (bool, error) {
	return false, nil
}

type Flavour struct {
	value string
}

func (f Flavour) Is(value string) MatchOperand {

}

type FlavourIsOperand struct {
	Value string
}

func (o FlavourIsOperand) Evaluate(ctx interface{}) (bool, error) {
	return o.Value == (ctx.(Context)).Flavour, nil
}

func FlavourIs(value string) FlavourIsOperand {
	return FlavourIsOperand{
		value,
	}
}

type FlavoursInOperand struct {
	Values []string
}

func (o FlavoursInOperand) Evaluate(ctx interface{}) (bool, error) {
	for _, flavour := range o.Values {
		if flavour == (ctx.(Context)).Flavour {
			return true, nil
		}
	}
	return false, nil
}

func FlavoursIn(values ...string) FlavoursInOperand {
	return FlavoursInOperand{
		values,
	}
}

type GuestOSIsOperand struct {
	Value string
}

func (o GuestOSIsOperand) Evaluate(ctx interface{}) (bool, error) {
	return o.Value == (ctx.(Context)).GuestOS, nil
}

func GuestOSIs(value string) GuestOSIsOperand {
	return GuestOSIsOperand{
		value,
	}
}

type GuestOSInOperand struct {
	Values []string
}

func (f GuestOSInOperand) Evaluate(ctx interface{}) (bool, error) {
	for _, guest := range f.Values {
		if guest == (ctx.(Context)).GuestOS {
			return true, nil
		}
	}
	return false, nil
}

var rule Operand = And(
	FlavoursIn("rhel", "ubuntu"),
	GuestOSIs("Ubuntu 16.04 LTS"),
)
