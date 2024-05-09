/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
)

// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "A brief description of your command",
	Long: `Show the information about your CPU, For example:
sysmon cpu -model`,

	Run: cpuInfo,
}

func cpuInfo(cmd *cobra.Command, args []string) {
	flagModelCpu := cmd.Flags().Bool("model", false, "show the model of your cpu")
	cmd.Println(*flagModelCpu)

	if *flagModelCpu {
		info, err := cpu.Info()
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
		cmd.Println(info[0].ModelName)
		return
	}
}

func init() {
	rootCmd.AddCommand(cpuCmd)

	cpuCmd.Flags().BoolP("model", "m", false, "show the model of your cpu")
}
