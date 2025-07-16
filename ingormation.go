package main

import (
	"fmt"

	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func main(){


	fontStyles := map[string]string{
		"bold" : "\033[1m",
		"italic" : "\033[3m",
		"reset" : "\033[0m",
	}
    
	greeting := `
 _                                        _   _                       _ 
(_)_ __   __ _  ___  _ __ _ __ ___   __ _| |_(_) ___  _ __   __   __ / |
| | '_ \ / _` + "`" + ` |/ _ \| '__| '_ ` + "`" + ` _ \ / _` + "`" + ` | __| |/ _ \| '_ \  \ \ / / | |
| | | | | (_| | (_) | |  | | | | | | (_| | |_| | (_) | | | |  \ V /  | |
|_|_| |_|\__, |\___/|_|  |_| |_| |_|\__,_|\__|_|\___/|_| |_|   \_/   |_|
         |___/                                                          
	
`
    host_info, err := host.Info()
	boot_time := host_info.BootTime / 360
	uptime := host_info.Uptime / 60

	cpu_info, err := cpu.Info()
	cpu_percentage, err := cpu.Percent(time.Second, false)
	cpu_count, _ := cpu.Counts(true)
	vm_info, err := mem.VirtualMemory()
	partitions, err := disk.Partitions(false)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(greeting)
    fmt.Println("_____________________________________________________________________")
	fmt.Println("\nWelcome ingormation | developed by " + fontStyles["italic"] + fontStyles["bold"] + "lexionq" + fontStyles["reset"] + " (github/lexionq)")
	fmt.Println("_____________________________________________________________________")
	fmt.Println(fontStyles["bold"] + "\nHost Informations")
	fmt.Println("_____________________________________________________________________")
	fmt.Print("\nHostname: "+fontStyles["reset"], host_info.Hostname)
	fmt.Print(fontStyles["bold"] + "\nHost ID: "+fontStyles["reset"], host_info.HostID)
	fmt.Print(fontStyles["bold"] + "\nOS: "+fontStyles["reset"], host_info.OS)
	fmt.Print(fontStyles["bold"] + "\nKernel Version: "+fontStyles["reset"], host_info.KernelVersion)
	fmt.Print(fontStyles["bold"] + "\nKernel Arch: "+fontStyles["reset"], host_info.KernelArch)
    fmt.Print(fontStyles["bold"] + "\nPlatform: "+fontStyles["reset"], host_info.Platform)
	fmt.Print(fontStyles["bold"] + "\nPlatform Version: "+fontStyles["reset"], host_info.PlatformVersion)
	fmt.Print(fontStyles["bold"] + "\nPlatform Family: "+fontStyles["reset"], host_info.PlatformFamily)
	fmt.Print(fontStyles["bold"] + "\nVirtualization Role: "+fontStyles["reset"], host_info.VirtualizationRole)
	fmt.Print(fontStyles["bold"] + "\nVirtualization System: "+fontStyles["reset"], host_info.VirtualizationSystem)
	fmt.Print(fontStyles["bold"] + "\nActive Processes: "+fontStyles["reset"], host_info.Procs)
	fmt.Println(fontStyles["bold"] + "\nBoot Time & Uptime: "+fontStyles["reset"], boot_time, " hours & " ,uptime, " mins")
    fmt.Println("_____________________________________________________________________")
	fmt.Println(fontStyles["bold"] + "\nHardware Informations")
	fmt.Println("_____________________________________________________________________")
	for i, ci := range cpu_info {
		fmt.Println("\n_____________________________________________________________________")
		fmt.Println(fontStyles["bold"] + "\nCPU", i)
		fmt.Println("_____________________________________________________________________")
		fmt.Print(fontStyles["bold"] + "\nCores: "+fontStyles["reset"], ci.Cores)
		fmt.Print(fontStyles["bold"] + "\nCore ID: "+fontStyles["reset"], ci.CoreID)
		fmt.Print(fontStyles["bold"] + "\nMicrocode: "+fontStyles["reset"], ci.Microcode)
		fmt.Print(fontStyles["bold"] + "\nFamily: "+fontStyles["reset"], ci.Family)
		fmt.Print(fontStyles["bold"] + "\nModel: "+fontStyles["reset"], ci.Model)
		fmt.Print(fontStyles["bold"] + "\nModel Name: "+fontStyles["reset"], ci.ModelName)
		fmt.Print(fontStyles["bold"] + "\nMhz: "+fontStyles["reset"], ci.Mhz)
		fmt.Print(fontStyles["bold"] + "\nPhysical ID & Vendor ID: "+fontStyles["reset"], ci.PhysicalID, " & ", ci.VendorID)
		fmt.Print(fontStyles["bold"] + "\nCacheSize: "+fontStyles["reset"], ci.CacheSize)
		fmt.Print(fontStyles["bold"] + "\nFlags: "+fontStyles["reset"], ci.Flags)
	}
	fmt.Println("\n_____________________________________________________________________")
	fmt.Println(fontStyles["bold"] + "\nGeneral CPU Information",)
	fmt.Println("_____________________________________________________________________")
	fmt.Print(fontStyles["bold"] + "\nCPU Count: " + fontStyles["reset"], cpu_count)
	for i, percent := range cpu_percentage {
		fmt.Print(fontStyles["bold"] + "\nCPU Usage: CPU " + fontStyles["reset"], i , percent)
	}
	fmt.Println("\n_____________________________________________________________________")
	fmt.Println(fontStyles["bold"] + "\nMemory Stats",)
	fmt.Println("_____________________________________________________________________")
	fmt.Print(fontStyles["bold"] + "\nTotal Memory: "+fontStyles["reset"], vm_info.Total/1024/1024, "MB")
	fmt.Print(fontStyles["bold"] + "\nUsed Memory (MB): "+fontStyles["reset"], vm_info.Used/1024/1024, "MB")
	fmt.Print(fontStyles["bold"] + "\nUsed Memory (Percent): "+fontStyles["reset"], uint64(vm_info.UsedPercent),"%")
	fmt.Print(fontStyles["bold"] + "\nFree Memory: "+fontStyles["reset"], vm_info.Available/1024/1024, "MB")
	fmt.Print(fontStyles["bold"] + "\nTotal Swap: "+fontStyles["reset"], vm_info.SwapTotal/1024/1024, "MB")
	fmt.Print(fontStyles["bold"] + "\nUsed Swap (MB): "+fontStyles["reset"], vm_info.SwapTotal/1024/1024 - vm_info.SwapFree/1024/1024, "MB")	
	fmt.Print(fontStyles["bold"] + "\nFree Swap: "+fontStyles["reset"], vm_info.SwapFree/1024/1024, "MB")
	fmt.Println("\n_____________________________________________________________________")
	fmt.Println(fontStyles["bold"] + "\nDisk Stats",)
	fmt.Println("_____________________________________________________________________")
	for _, part := range partitions{
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
           fmt.Println("Error: ",err)
		   return
		}

		fmt.Println("\n_____________________________________________________________________")
		fmt.Println(fontStyles["bold"] + "\nDisk: ", part.Device)
		fmt.Println("_____________________________________________________________________")
		fmt.Print(fontStyles["bold"] + "\nMount Point:  "+fontStyles["reset"], part.Mountpoint)
		fmt.Print(fontStyles["bold"] + "\nFilesystem:  "+fontStyles["reset"], part.Fstype)
		fmt.Print(fontStyles["bold"] + "\nTotal Disk (GB):  "+fontStyles["reset"], float32(usage.Total)/1e9)
	    fmt.Print(fontStyles["bold"] + "\nUsed Disk (GB):  "+fontStyles["reset"], float32(usage.Used)/1e9, usage.UsedPercent,"%")
		fmt.Print(fontStyles["bold"] + "\nFree Disk (GB):  "+fontStyles["reset"], float32(usage.Free)/1e9)
	}
	
}