/*
Copyright © 2025 maxvolume <ben@schonbeck.io>

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
	"mimeutils/benmimer/parsers"

	"github.com/spf13/cobra"
)

// set variable for Peek messages flag
var PeekNumber int32

// set variable for Filename flag
var Filename string

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse a mbox file. Specify target file with -f flag. Currently benmimer only runs in Peek mode.",
	Long: `Eventually config file settings will dictate which portions of mail messages
	should be parsed. For now we're just defaulting to mime types of text/plain. File attachements 
	will get extracted/parsed at a later date.`,
	Run: func(cmd *cobra.Command, args []string) {
		parsers.ParseMbox(Filename, int(PeekNumber))
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
	parseCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Path to mbox file (required)")
	parseCmd.Flags().Int32VarP(&PeekNumber, "peekNumber", "p", 10, "Set number of messages to parse. Default is 10.")
	parseCmd.MarkFlagRequired("filename")
}
