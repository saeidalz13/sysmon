/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/spf13/cobra"
)

var (
	second int32
)

// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "A brief description of your command",
	Long: `Shows the information about your CPU, For example:
sysmon cpu -model`,
}

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Shows the model of CPU",
	Run:   subCmdCpuModel,
}

func subCmdCpuModel(cmd *cobra.Command, args []string) {
	info, err := cpu.Info()
	if err != nil {
		cmd.PrintErrln(err)
		os.Exit(1)
	}
	cmd.Println(info[0].ModelName)
}

var percentCmd = &cobra.Command{
	Use:   "usage",
	Short: "Shows the usage percentage of CPU",
	Run:   subCmdCpuUsage,
}

func subCmdCpuUsage(cmd *cobra.Command, args []string) {
	cmd.Printf("calculating cpu usage over %d second...\n\n", second)
	usage, err := cpu.Percent(time.Second*time.Duration(second), true)
	if err != nil {
		cmd.PrintErrln(err)
		os.Exit(ExitError)
	}
	for i, percent := range usage {
		cmd.Printf("cpu %d:\t%.2f%%\n", i, percent)
	}
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Shows the load averages from your CPU",
	Run:   subCmdCpuLoad,
}

func subCmdCpuLoad(cmd *cobra.Command, args []string) {
	avgStat, err := load.Avg()
	if err != nil {
		cmd.PrintErrln(err)
		os.Exit(ExitError)
	}

	cmd.Print("\nCPU load averages summary:\n\n")
	cmd.Printf("1-min load:\t%+.2f\n", avgStat.Load1)
	cmd.Printf("5-min load:\t%+.2f\n", avgStat.Load5)
	cmd.Printf("15-min load:\t%+.2f\n", avgStat.Load15)
}

var countCpu = &cobra.Command{
	Use:   "count",
	Short: "Shows the number of logical cores",
	Run:   subCmdCpuCount,
}

func subCmdCpuCount(cmd *cobra.Command, args []string) {
	cpuModes := []bool{false, true}

	for i, mode := range cpuModes {
		n, err := cpu.Counts(mode)
		if err != nil {
			cmd.Println(err)
			os.Exit(ExitError)
		}

		if i == 0 {
			cmd.Printf("Physical:\t%d\n", n)
		} else {
			cmd.Printf("Logical:\t%d\n", n)
		}
	}
}

func init() {
	rootCmd.AddCommand(cpuCmd)
	cpuCmd.AddCommand(modelCmd)
	cpuCmd.AddCommand(percentCmd)
	cpuCmd.AddCommand(loadCmd)
	cpuCmd.AddCommand(countCpu)

	percentCmd.Flags().Int32VarP(&second, "time", "t", 1, "set the seconds you want to calculate the cpu percentage over")
	percentCmd.MarkFlagRequired("time")
}
