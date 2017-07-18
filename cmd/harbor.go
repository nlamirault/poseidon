// Copyright (C) 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"io"

	"github.com/golang/glog"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/nlamirault/poseidon/tides"
)

var (
	name string
)

type harborCmd struct {
	out io.Writer
}

func newHarborCmd(out io.Writer) *cobra.Command {
	harborCmd := &harborCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "harbor",
		Short: "Manage harbors.",
		Long:  "Manage environments. See subcommands.",
		RunE:  nil,
	}

	listHarborsCmd := &cobra.Command{
		Use:   "list",
		Short: "List all harbors",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := harborCmd.listHarbors(out); err != nil {
				return err
			}
			return nil
		},
	}
	getHarborCmd := &cobra.Command{
		Use:   "get",
		Short: "Retreive a harbor",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(name) == 0 {
				return errors.New("missing harbor name")
			}
			if err := harborCmd.getHarbor(name, out); err != nil {
				return err
			}
			return nil
		},
	}

	getHarborCmd.PersistentFlags().StringVar(&name, "name", "", "Harbor's name")
	cmd.AddCommand(listHarborsCmd)
	cmd.AddCommand(getHarborCmd)
	return cmd
}

func (cmd harborCmd) listHarbors(out io.Writer) error {
	glog.V(1).Infof("List all harbors")
	harbors, err := tides.ExtractHarbors()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"ID", "Name"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	for id, name := range harbors {
		table.Append([]string{
			id,
			name,
		})
	}
	table.Render()
	return nil
}

func (cmd harborCmd) getHarbor(name string, out io.Writer) error {
	glog.V(1).Infof("Retrieve harbor: %s", name)
	harbor, err := tides.DescribeHarbor(name)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"Information", "Value"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	for id, name := range harbor {
		table.Append([]string{
			id,
			name,
		})
	}
	table.Render()
	return nil
}
