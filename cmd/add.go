/*
Copyright © 2022 Kiyoshi Kanazawa <kiyocy24@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kiyocy24/roulette/helper"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add an item and its weight",
	Long: `Add an item and its weight.
If you specify an already registered item name, it will be updated

example)
  add --name banana --weight 100
  add -n apple --weight 50`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if name == "" {
			return errors.New("--name is required")
		}
		weight, err := cmd.Flags().GetInt("weight")
		if err != nil {
			return err
		}
		if weight <= 0 {
			return errors.New("--weight is required and set one or more")
		}

		items, err := helper.Load(filepath)
		if err != nil {
			return err
		}
		_, exist := items[name]
		items[name] = weight
		b2, err := json.MarshalIndent(items, "", "\t")
		if err != nil {
			return err
		}
		err = helper.WriteFile(filepath, b2)
		if err != nil {
			return err
		}

		if exist {
			fmt.Printf("%s is updated", name)
		} else {
			fmt.Printf("%s is added", name)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("name", "n", "", "item name (required)")
	addCmd.Flags().IntP("weight", "w", 0, "weight (required)")
}
