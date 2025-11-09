package cmd

import (
	"github.com/jmilne22/distrohop/internal/flows"
	"github.com/spf13/cobra"
)

// stateCmd represents the state command
var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Compares the state of packages on the system vs the config.yaml",
	Long: `Compares packages installed on the system vs what's specified in the config.yaml:

Gives a list of the state of each package and whether it's installed or not`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return flows.State()
	},
}

func init() {
	rootCmd.AddCommand(stateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
