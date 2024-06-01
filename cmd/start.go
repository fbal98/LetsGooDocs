package cmd

import (
	"fmt"

	"LetsGooDocs/services"
	"LetsGooDocs/utils"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "This command starts the process of generating Docs",
	Long: `This command starts the process of generating Docs for
		You can use this command to generate Docs for your API or any other project that you are working on.`,
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

//API subCommand
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "This command starts the process of generating API Docs",
	Long: `This command starts the process of generating API Docs for
		You can use this command to generate API Docs for your API.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		if utils.IsPathNotFound(path) {
			fmt.Println("Path not found")
			return
		}
		fmt.Println("Generating API Docs for path: ", path)

		files, error := utils.ListDirNames(path)

		if error != nil {
			fmt.Println("Error: ", error)
			return
		}

		fileContents := utils.GetFilesContent(files)

		prompt := utils.GeneratePrompt(fileContents)

		document := services.Chat(prompt)

		utils.WriteToFile("Documentation.md", document)
		
		
	},


}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.AddCommand(apiCmd)
	
	// Here you will define your flags and configuration settings.

	const defaultPath = "."
	apiCmd.Flags().StringP("path", "p", defaultPath, "Path to the API's root directory")

}


