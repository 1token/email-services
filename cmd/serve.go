package cmd

import (
	"fmt"
	"github.com/1token/email-services/pkg/config"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve --config [ config file ]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Example: "email-services serve --config examples/config-dev.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serve(cmd, args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func serve(cmd *cobra.Command, args []string) error {
	errc := make(chan error, 3)

	// unmarshal config into Struct
	var c config.Config
	err := viper.Unmarshal(&c)
	if err != nil {
		return fmt.Errorf("unmarshal failed: %v", err)
	}

	// fmt.Printf("Web: %s\n", c.Web.HTTP)

	return <-errc
}
