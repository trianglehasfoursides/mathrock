package main

import "github.com/spf13/cobra"

func init() {
	root.AddCommand(play)
}

var play = &cobra.Command{
	Use:    "play",
	Short:  "play random mathrock song",
	Long:   "",
	Hidden: true,
	RunE:   func(cmd *cobra.Command, args []string) error {},
}
