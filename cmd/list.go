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
	"fmt"
	"github.com/kiyocy24/roulette/helper"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a list of items.",
	Long: `Name, weight and percentage are displayed.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		items, err := helper.Load(filepath)
		if err != nil {
			return err
		}
		totalWeight := helper.Total(items)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Weight", "Percent"})
		var totalPercent float32
		for k, v := range items {
			percent := float32(v*100) / float32(totalWeight)
			table.Append([]string{k, strconv.Itoa(v), fmt.Sprintf("%.2f%%", percent)})
			totalPercent += percent
		}
		table.SetFooter([]string{"Total", strconv.Itoa(totalWeight), fmt.Sprintf("%.2f%%", totalPercent)})
		table.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
