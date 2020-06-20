package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	zshCompdef    = "\ncompdef _ugc ugc\n"
	completionCmd = &cobra.Command{
		Use:   "completion [bash|zsh|powershell|fish]",
		Short: "Generate completion script",
		// DisableFlagsInUseLine: true,
		SuggestFor: []string{"bash", "zsh", "powershell", "fish"},
		ValidArgs:  []string{"bash", "zsh", "powershell", "fish"},
		Args:       cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				if err := cmd.Root().GenBashCompletion(os.Stdout); err != nil {
					log.Fatalf("GenBashCompletion failed with %v\n", err)
				}
			case "zsh":
				runCompletionZsh(cmd, os.Stdout)
			case "powershell":
				if err := cmd.Root().GenPowerShellCompletion(os.Stdout); err != nil {
					log.Fatalf("GenPowerShellCompletion failed with %v\n", err)
				}
			case "fish":
				if err := cmd.Root().GenFishCompletion(os.Stdout, true); err != nil {
					log.Fatalf("GenFishCompletion failed with %v\n", err)
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

func runCompletionZsh(cmd *cobra.Command, out io.Writer) {
	if err := cmd.Root().GenZshCompletion(out); err != nil {
		log.Fatalf("GenZshCompletion failed with %v\n", err)
	}

	if _, err := io.WriteString(out, zshCompdef); err != nil {
		log.Fatalf("zshCompdef failed with %v\n", err)
	}
}
