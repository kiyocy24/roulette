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
	"errors"
	"fmt"
	"github.com/kiyocy24/roulette/helper"

	"github.com/spf13/cobra"
)

// lotCmd represents the lot command
var lotCmd = &cobra.Command{
	Use:   "lot",
	Short: "draw lots based on weight",
	Long: `Draw lots based on weight.
The name of the item you hit is displayed..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		items, err := helper.Load(filepath)
		if err != nil {
			return err
		}

		totalWeight := helper.Total(items)
		if totalWeight <= 0 {
			return errors.New("total weight is 0")
		}
		hit := helper.Lot(items)
		percentage := float64(items[hit]*100) / float64(totalWeight)
		fmt.Printf("%s %.2f%%\n", hit, percentage)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(lotCmd)
}
