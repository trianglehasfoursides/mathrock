package storage

import (
	"errors"
	"github.com/spf13/cobra"
)

var Upload = &cobra.Command{
	Use:     "upload",
	Aliases: []string{"up"},
	Short:   "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if args == nil {
			return errors.New("")
		}

		if len(args) < 1 {
			return errors.New("")
		}

		if name, dir := args[0], args[1]; name == "" && dir == "" {
			return errors.New("")
		}
	},
}
