package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

// appserver > jdk > guest > flavour

func main() {
	flavours := []string{
		"RHEL/KVM",
		"Canonical/KVM",
		"Canonical/ESXi",
		"Canonical/LXC",
		"SuSE/KVM",
		"SuSE/ESXi",
		"SuSE/Xen",
		"Oracle/KVM (OracleVM)",
		"Oracle/Hyper-V",
		"VMware/ESXi",
		"Huawei/KVM",
		"Huawei/ESXi",
		"Huawei/Xen",
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

	/*
		// contraints
		appserver2jdk := map[string][]string{
			"IBM WebSphere 9": []string{
				"",
			},
		}

		appserver2guest := map[string][]string{
			"IBM WebSphere 9": []string{
				"",
			},
		}

		appserver2flavour := map[string][]string{
			"IBM WebSphere 9": []string{
				"",
			},
		}
	*/

	//jdk2guest = map[string][]string{}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, appserver := range appservers {
		for _, jdk := range jdks {
			for _, guest := range guests {
				for _, flavour := range flavours {
					row := sheet.AddRow()
					cell := row.AddCell()
					cell.Value = flavour
					cell = row.AddCell()
					cell.Value = guest
					cell = row.AddCell()
					cell.Value = jdk
					cell = row.AddCell()
					cell.Value = appserver
					fmt.Sprintf("%s > %s > %s > %s\n", flavour, guest, jdk, appserver)
				}
			}
		}
	}

	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
