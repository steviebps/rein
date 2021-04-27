package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/steviebps/rein/internal/logger"
	rein "github.com/steviebps/rein/pkg"
)

var getCmdError = logger.ErrorWithPrefix("Error running get command: ")

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a value of a toggle",
	Long:  `we'll figure this out`,
	Run: func(cmd *cobra.Command, args []string) {
		var value interface{}
		version := viper.GetString("version")
		toggle, _ := cmd.Flags().GetString("toggle")
		chamberName, _ = cmd.Flags().GetString("chamber")

		globalChamber.TraverseAndBuild(func(c *rein.Chamber) bool {
			if c.Name == chamberName {
				value = c.GetToggleValue(toggle, version)
			}

			return value != nil
		})

		if value == nil {
			getCmdError(fmt.Sprintf("Could not find toggle value %q inside chamber %q", toggle, chamberName))
			os.Exit(1)
		}

		fmt.Println(value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("toggle", "t", "", "get the specified toggle value")
	getCmd.Flags().StringP("chamber", "c", "", "retrieves toggle from specified chamber")

	getCmd.MarkFlagRequired("toggle")
	getCmd.MarkFlagRequired("chamber")
}
