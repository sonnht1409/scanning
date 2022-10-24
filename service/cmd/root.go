package cmd

import (
	"github.com/sonnht1409/scanning/service/handlers"
	"github.com/sonnht1409/scanning/service/workers"
	"github.com/spf13/cobra"
)

// NewServiceCommand
func NewServiceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			h := handlers.NewServiceHandlers()
			h.ApiRegister()
			h.Start()
		},
	}
	return cmd
}

// NewWorkerCommand
func NewWorkerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "worker",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			w := workers.NewServiceWorker()
			w.Subscribe()
		},
	}
	return cmd
}

// Execute
func Execute() {
	rootCmd := &cobra.Command{
		Use:   "scanning-service",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.AddCommand(NewServiceCommand())
	rootCmd.AddCommand(NewWorkerCommand())
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
