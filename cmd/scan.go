/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans object apiVersion skew in a Kubernetes Cluster",
	Long: `This utility connects to your Kubernetes Cluster using your default kubeconfig file or any
	kubeconfig file you specify, gets the apiVersion of all objects in a particular namespace  or all namespaces,
	and compares these values with the apiVersions in the corresponding Kubernetes version you specify to scheck for skews
	before upgrading your cluster.`,
	//Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scan called")
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")
	scanCmd.PersistentFlags().StringP("kubeconfig", "k", "~/.kube/config", "Specify the path to your kubeconfig file. If not specified, default kubeconfig file found in ~/kube/config will be used.")
	scanCmd.PersistentFlags().StringP("upgrade-version", "U", "", "Specify the Kubernetes version you want to upgrade to.")
	scanCmd.PersistentFlags().StringP("namespace", "n", "default", "Specify the pnamespace you want to scan. If not specified, kskew will scan the default namespace.")
	scanCmd.PersistentFlags().StringP("all-namespaces", "A", "", "Scan all namespaces.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
