package rmByPattern

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	var (
		showVersion bool
		dirPath     string
		rxPattern   string
		remove      bool
		testMode    bool
	)

	var rootCmd = &cobra.Command{
		Use:   "rm-by-pattern [path]",
		Short: "Remove files by pattern",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pwd := os.Getenv("PWD")
			if len(args) > 0 {
				dirPath = args[0]

			}

			if showVersion {
				printVersion()
				return
			}

			if dirPath == "" {
				log.Fatal("Please provide folder path: rm-by-pattern [path] [flags]")
				return
			}

			if remove {
				confPath := filepath.Join(pwd, "rm-by-pattern.yaml")
				conf, err := GetYamlConfig(confPath)
				ErrChk(err)
				RmFiles(dirPath, conf.Patterns, testMode)
			} else {
				re := regexp.MustCompile(rxPattern)
				patterns, ext, err := IdentifyPatterns(dirPath, re)

				if err != nil {
					log.Fatalf("Error identifying patterns: %v", err)
				}

				err = CreateYamlConfig(pwd, patterns, ext)
				if err != nil {
					log.Fatalf("Error creating yaml config: %v", err)
				}

			}

		},
	}

	rootCmd.Flags().StringVarP(&dirPath, "path", "p", "", "Golder path")
	rootCmd.Flags().StringVarP(&rxPattern, "re", "r", "\\d+x\\d+", "Regex Pattern")
	rootCmd.Flags().BoolVarP(&remove, "start", "s", false, "Start Removing Files")
	rootCmd.Flags().BoolVarP(&testMode, "test", "t", false, "Run in test mode")
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Show Version")
	rootCmd.AddCommand(newCompletionCmd())

	return rootCmd
}

func newCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate fish completion script",
		Run:   generateFishCompletion,
	}
}

func generateFishCompletion(cmd *cobra.Command, args []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get user home directory: %v", err)
	}

	fishCompletionDir := filepath.Join(homeDir, ".config", "fish", "completions")
	if err := os.MkdirAll(fishCompletionDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create fish completions directory: %v", err)
	}

	fishCompletionFile := filepath.Join(fishCompletionDir, "rm-by-pattern.fish")
	f, err := os.Create(fishCompletionFile)
	if err != nil {
		log.Fatalf("failed to create fish completion file: %v", err)
	}
	defer f.Close()

	if err := cmd.Root().GenFishCompletion(f, true); err != nil {
		log.Fatalf("failed to generate fish completion script: %v", err)
	}

	fmt.Printf("Fish completion script generated at: %s\n", fishCompletionFile)
}
