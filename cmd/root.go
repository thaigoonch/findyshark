/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/thaigoonch/findyshark/app"
)

var cfgFile string
var HashSpace string // hash that is used to replace spaces in user input
var HashTab string   // hash that is used to replace tabs in user input
var VERSION string

const (
	config_ignore = "ignore"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "findyshark",
	Short: "A recursive-searching CLI tool searching tool",
	Long:  "Runs out of pwd.",
	Run: func(cmd *cobra.Command, args []string) {
		fileCriteria := ":"
		fileExt, _ := cmd.Flags().GetString("extension")
		if fileExt != "" {                                   // if user specified a file extension,
			if app.ValidateFileExtension(fileExt) {          // validate it
				fileCriteria = fileCriteria + "." + fileExt  // use it
			} else {
				log.Fatal("Invalid file extension")
			}
		}

		fmt.Print(drawShark(cfgFile))
		inp := app.GetInput()
		HashSpace = app.RandomString(32)
		HashTab = app.RandomString(32)
		inp = replaceWhiteSpace(inp)
		ignoreFiles := getIgnoresFromConfig()
		bashCmd := os.Getenv("GOPATH") + "/bin/findysharksrch"
		istatus, _ := cmd.Flags().GetBool("insensitive")
		if istatus {                                             // if case-insensitive flag is true,
			bashCmd = os.Getenv("GOPATH") + "/bin/findysharkisrch" // use case-insensitive logic
		}

		output := doFind(inp, ignoreFiles, bashCmd, fileCriteria)
		app.TableContentResults(output)
	},
}

func Execute(version string) {
	VERSION = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "path to config file (default is $HOME/.findyshark.yaml)")
	rootCmd.PersistentFlags().BoolP("insensitive", "i", false, "search case-insensitive")
	rootCmd.PersistentFlags().StringP("extension", "e", "", "search in specified file extension; e.g. txt")
	flags := rootCmd.Flags()
	cobra.OnInitialize(initConfig)
	viper.BindPFlag(config_ignore, flags.Lookup(config_ignore))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		if app.ValidateConfigPath(cfgFile) {
			// Use config file from the flag.
			viper.SetConfigFile(cfgFile)
		} else {
			log.Fatal("Invalid config file")
		}
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".findyshark" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".findyshark")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		cfgFile = viper.ConfigFileUsed()
	}
}

func getIgnoresFromConfig() string {
	config_ignore := viper.GetString(config_ignore)
	ignore_items := strings.Split(config_ignore, ",")

	if strings.HasSuffix(config_ignore, ",") { // if trailing comma
		ignore_items = ignore_items[:len(ignore_items)-1] // ignore the last entry
	}
	finalCmd := ""
	i := 0
	for range ignore_items {
		ignore_items[i] = strings.Replace(ignore_items[i], " ", "", -1) // ignore spaces in config file entries
		ignore_items[i] = app.Sanitize_inputs(ignore_items[i])
		ignore_items[i] = "^" + ignore_items[i]
		if i < len(ignore_items)-1 {
			finalCmd = finalCmd + ignore_items[i] + ":[0-9]+:|"
		} else {
			finalCmd = finalCmd + ignore_items[i] + ":[0-9]+:"
		}
		i++
	}
	finalCmd = strings.Replace(finalCmd, "\\*", ".*", -1)
	return finalCmd
}

func replaceWhiteSpace(value string) string {
	value = strings.Replace(value, "\\ ", HashSpace, -1)
	value = strings.Replace(value, "\\	", HashTab, -1)
	return value
}

func doFind(term, ignore, bashCmd, fileCriteria string) string {
	cmd, err := exec.Command(bashCmd, fileCriteria, HashSpace, HashTab, term, ignore).Output()
	if err != nil {
		fmt.Printf("%s error: %s\n", bashCmd, err)
	}
	output := string(cmd)
	return output
}

func drawShark(configStr string) string {
	cmd, err := exec.Command(os.Getenv("GOPATH") + "/bin/findysharkbanner", configStr).Output()
	if err != nil {
		fmt.Printf("banner error: %s\n", err)
	}
	output := string(cmd)
	return output
}
