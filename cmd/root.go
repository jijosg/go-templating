/*
Copyright © 2022 Jijo Sunny

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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"text/template"
)

var cfgFile, valuesFile string
var temp *template.Template

type Count struct {
	Count int
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-templating",
	Short: "Go templating example using basic template",
	Long:  `Go templating example using basic template.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("template called")
		// Reads a go template and executes it
		dat, err := ioutil.ReadFile(valuesFile)
		if err != nil {
			fmt.Errorf("Error reading values file: %v", err)
		}
		result, err := unmarshalYamlFile(dat)
		err = temp.Execute(os.Stdout, result)
		if err != nil {
			panic(err)
		}

	},
}

func unmarshalYamlFile(content []byte) (map[string]interface{}, error) {
	var vars map[string]interface{}
	err := yaml.Unmarshal(content, &vars)
	if err != nil {
		return nil, err
	}
	return vars, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.template.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&valuesFile, "values", "f", "resources/values.yaml", "Values file Path")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	temp = template.Must(template.ParseFiles("resources/example.gotmpl"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".template" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".template")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
