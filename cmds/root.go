package cmds

import (
	"flag"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "hugo-checker [command]",
		Short:             `Check various aspects of a hugo site`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			c.Flags().VisitAll(func(flag *pflag.Flag) {
				log.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
			})
		},
	}

	// flags := rootCmd.PersistentFlags()
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	flag.CommandLine.Parse([]string{})
	flag.Set("stderrthreshold", "ERROR")

	rootCmd.AddCommand(NewCmdAddFrontMatter())
	rootCmd.AddCommand(NewCmdAddIntro())
	rootCmd.AddCommand(NewCmdCheckFrontMatter())
	return rootCmd
}
