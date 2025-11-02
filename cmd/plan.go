package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/at-plan/usecase"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "A CLI tool for managing your at-plan",
	Long:  `A CLI tool for managing your at-plan`,
	Run: func(cmd *cobra.Command, args []string) {
		atRequest := cmd.Context().Value("atlantis").(usecase.AtlantisRequest)
		fmt.Println(atRequest)
	},
}
