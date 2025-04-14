/*
Copyright © 2025 Naufal Khairil Imami
*/
package cmd

import (
	"github.com/naufalkhairil/Filewatcher/modules/handler"
	"github.com/naufalkhairil/Filewatcher/modules/watcher"
	"github.com/spf13/cobra"
)

var watcherCmd = &cobra.Command{
	Use:   "watcher",
	Short: "Watch changes in your directory",
	PreRun: func(cmd *cobra.Command, args []string) {
		handler.InitHandler()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return watcher.Start()
	},
}

func init() {
	rootCmd.AddCommand(watcherCmd)
}
