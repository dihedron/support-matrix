package main

import (
	"fmt"

	"github.com/dihedron/go-bool/rules"
	"github.com/tealeg/xlsx"
)

// appserver > jdk > guest > hypervisor > distribution

func main() {

	distributions := []string{
		"Red Hat",
		"Canonical",
		"SuSE",
		"Oracle",
		"Huawei",
	}

	hypervisors := []string{
		"KVM",
		"ESXi",
		"LXC",
		"Xen",
		"Hyper-V",
		"Oracle VM (KVM)",
	}

	guests := []string{
		"Ubuntu 16.04 LTS",
		"Ubuntu 18.04 LTS",
		"Red Hat Enterprise Linux 6",
		"Red Hat Enterprise Linux 7",
		"Windows Server 2008 R2",
		"Windows Server 2012 R2",
		"Windows Server 2016",
		"SuSE Enterprise Linux 11",
		"SuSE Enterprise Linux 12",
	}

	jdks := []string{
		"OpenJDK 1.8",
		"Oracle JDK 1.8",
		"IBM JDK 1.8",
	}

	appservers := []string{
		"JBoss EAP 7.1",
		"JBoss EAP 7.0",
		"IBM WebSphere 9",
		"Oracle WebLogic",
	}

	// supported configurations
	flavours := rules.Or(
		rules.And(
			Distribution("Red Hat"),
			Hypervisor("KVM"),
		),
		rules.And(
			Distribution("Canonical"),
			Hypervisor("KVM", "ESXi", "LXC"),
		),
		rules.And(
			Distribution("SuSE"),
			Hypervisor("KVM", "ESXi", "Xen"),
		),
		rules.And(
			Distribution("Oracle"),
			Hypervisor("Oracle VM (KVM)", "Hyper-V"),
		),
		rules.And(
			Distribution("VMware"),
			Hypervisor("ESXi"),
		),
		rules.And(
			Distribution("Huawei"),
			Hypervisor("KVM", "ESXi", "Xen"),
		),
	)

	supported := rules.And(
		flavours,
		rules.Or(
			JDK("OpenJDK 1.8"),
			JDK("IBM JDK 1.8"),
		),
	)

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, distribution := range distributions {
		for _, hypervisor := range hypervisors {
			for _, guest := range guests {
				for _, jdk := range jdks {
					for _, appserver := range appservers {
						stack := Stack{
							distribution,
							hypervisor,
							guest,
							jdk,
							appserver,
						}
						if result, err := supported.Evaluate(stack); err == nil && result {
							fmt.Printf("ADDING %s > %s > %s > %s > %s\n", distribution, hypervisor, guest, jdk, appserver)
							row := sheet.AddRow()
							cell := row.AddCell()
							cell.Value = distribution
							cell = row.AddCell()
							cell.Value = hypervisor
							cell = row.AddCell()
							cell.Value = guest
							cell = row.AddCell()
							cell.Value = jdk
							cell = row.AddCell()
							cell.Value = appserver
						} else {
							fmt.Printf("SKIPPING %s > %s > %s > %s > %s\n", distribution, hypervisor, guest, jdk, appserver)
						}
					}
				}
			}
		}
	}

	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
