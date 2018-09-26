package main

import (
	"fmt"

	"github.com/dihedron/go-bool/rules"
	"github.com/tealeg/xlsx"
)

// appserver > jdk > guest > hypervisor > host > distribution

//type Distribution string

func main() {

	supported := rules.And(
		supportedDistributions,
		//supportedByOracleJDK18,
		//supportedByWebSphere9x,
		supportedByJBossEAP7x,
	)

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, distribution := range AllDistributions {
		for _, host := range AllHostOSs {
			for _, hypervisor := range AllHypervisors {
				for _, guest := range AllGuestOSs {
					for _, jdk := range AllJDKs {
						for _, appserver := range AllAppServers {
							stack := Stack{
								distribution,
								host,
								hypervisor,
								guest,
								jdk,
								appserver,
							}
							if result, err := supported.Evaluate(stack); err == nil && result {
								fmt.Printf("ADDING %s > %s > %s > %s > %s > %s\n", distribution, host, hypervisor, guest, jdk, appserver)
								row := sheet.AddRow()
								cell := row.AddCell()
								cell.Value = string(distribution)
								cell = row.AddCell()
								cell.Value = string(host)
								cell = row.AddCell()
								cell.Value = string(hypervisor)
								cell = row.AddCell()
								cell.Value = string(guest)
								cell = row.AddCell()
								cell.Value = string(jdk)
								cell = row.AddCell()
								cell.Value = string(appserver)
								// } else {
								// 	fmt.Printf("SKIPPING %s > %s > %s > %s > %s > %s\n", distribution, host, hypervisor, guest, jdk, appserver)
							}
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
