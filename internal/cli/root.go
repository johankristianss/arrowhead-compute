package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const TimeLayout = "2006-01-02 15:04:05"

var Verbose bool
var ColoniesServerHost string
var ColoniesServerPort int
var ColoniesInsecure bool
var ColoniesSkipTLSVerify bool
var ColoniesUseTLS bool
var ColonyName string
var ColonyPrvKey string
var ExecutorName string
var ExecutorID string
var ExecutorType string
var ExecutorPrvKey string
var FsDir string
var SWName string
var SWType string
var SWVersion string
var HWCPU string
var HWModel string
var HWNodes int
var HWMem string
var HWStorage string
var HWGPUCount int
var HWGPUNodeCount int
var HWGPUName string
var HWGPUMem string
var LocDesc string
var Lat float64
var Long float64
var K8sName string
var K8sNamespace string
var K8sPVC string
var ProcessID string
var AddDebugLogs bool
var ParallelContainers bool
var GPU bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

var rootCmd = &cobra.Command{
	Use:   "arrowhead-compute",
	Short: "Arrowhead Compute System",
	Long:  "Arrowhead Compute System",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
