package main

import (
	"fmt"

	"github.com/dihedron/go-bool/logic"
	"github.com/dihedron/go-bool/model"
	"github.com/tealeg/xlsx"
)

// appserver > jdk > guest > hypervisor > host > distribution

//type Distribution string

func main() {

	supported := logic.And(
		// logic.Not(
		// 	model.Hypervisor(model.ESXi),
		// ),
		//model.Distribution(model.RedHat),
		model.OpenStackDistributionsSupportMatrix,
		model.Hypervisor(model.KVM),
		logic.Or(
			model.JBossEAP7xSupportMatrix,
			model.WebSphere8xBaseSupportMatrix,
			// model.WebSphere8xNetDeploySupportMatrix,
			model.WebSphere9xBaseSupportMatrix,
			// model.WebSphere9xNetDeploySupportMatrix,
			model.WebLogic11gSupportMatrix,
			model.WebLogic12cSupportMatrix,
		),
	)

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, distribution := range model.AllDistributions {
		for _, host := range model.AllHostOSs {
			for _, hypervisor := range model.AllHypervisors {
				for _, guest := range model.AllGuestOSs {
					for _, jdk := range model.AllJDKs {
						for _, appserver := range model.AllAppServers {
							stack := model.Stack{
								Distribution: distribution,
								HostOS:       host,
								Hypervisor:   hypervisor,
								GuestOS:      guest,
								JDK:          jdk,
								AppServer:    appserver,
							}
							if result, err := supported.Evaluate(stack); err == nil && result {
								// fmt.Printf("ADDING %s > %s > %s > %s > %s > %s\n", distribution, host, hypervisor, guest, jdk, appserver)
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

	err = file.Save("SupportMatrix.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
