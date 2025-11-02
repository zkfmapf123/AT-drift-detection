package cmd

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/at-plan/usecase"
)

var rootCmd = &cobra.Command{
	Use:   "at-root",
	Short: "A CLI tool for managing your at-plan",
	Long:  `A CLI tool for managing your at-plan`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		atRequest := usecase.AtlantisRequest{
			cmd.Flag("github-token").Value.String(),
			cmd.Flag("github-repo-ref").Value.String(),
			cmd.Flag("atlantis-url").Value.String(),
			cmd.Flag("atlantis-token").Value.String(),
			cmd.Flag("atlantis-repository").Value.String(),
			cmd.Flag("atlantis-config").Value.String(),
		}

		ctx := context.WithValue(context.Background(), "atlantis", atRequest)
		cmd.SetContext(ctx)
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("github-token", "g", "", "The Github token")
	rootCmd.PersistentFlags().StringP("github-repo-ref", "r", "", "The Github repository reference")
	rootCmd.PersistentFlags().StringP("atlantis-url", "u", "", "The Atlantis URL")
	rootCmd.PersistentFlags().StringP("atlantis-token", "t", "", "The Atlantis token")
	rootCmd.PersistentFlags().StringP("atlantis-repository", "r", "", "Atlantis Repository")
	rootCmd.PersistentFlags().StringP("atlantis-config", "c", "", "Atlantis Config File")
}

func Execute() {
	rootCmd.AddCommand(planCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Println("Error executing command:", err)
		os.Exit(1)
	}
}
