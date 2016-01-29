package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "apex",
	PersistentPreRun: pv.preRun,
	SilenceErrors:    true,
}

func init() {
	pf := rootCmd.PersistentFlags()

	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(invokeCmd)
	rootCmd.AddCommand(rollbackCmd)
	rootCmd.AddCommand(logsCmd)
	rootCmd.AddCommand(metricsCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(docsCmd)
	rootCmd.AddCommand(upgradeCmd)
	rootCmd.AddCommand(versionCmd)

	pf.StringVarP(&pv.Chdir, "chdir", "C", "", "Working directory")
	pf.BoolVarP(&pv.DryRun, "dry-run", "D", false, "Perform a dry-run")
	pf.StringSliceVarP(&pv.Env, "env", "e", nil, "Environment variable")
	pf.StringVarP(&pv.LogLevel, "log-level", "l", "info", "Log severity level")
	pf.BoolVarP(&pv.Yes, "yes", "y", false, "Automatic yes to prompts")
	pf.StringVarP(&pv.Profile, "profile", "p", "", "AWS profile to use")

	// Add a newline at the top of each help message
	rootCmd.SetHelpTemplate("\n" + rootCmd.HelpTemplate())
}
