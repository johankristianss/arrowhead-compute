package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/johankristianss/arrowhead-compute/pkg/build"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start executor",
	Long:  "Start executor",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()

		if Verbose {
			log.SetLevel(log.DebugLevel)
		}

		fmt.Println("Starting Arrowhead Compute Service")
	},
}

func parseEnv() {
	var err error
	ColoniesServerHostEnv := os.Getenv("COLONIES_SERVER_HOST")
	if ColoniesServerHostEnv != "" {
		ColoniesServerHost = ColoniesServerHostEnv
	}

	ColoniesServerPortEnvStr := os.Getenv("COLONIES_SERVER_PORT")
	if ColoniesServerPortEnvStr != "" {
		ColoniesServerPort, err = strconv.Atoi(ColoniesServerPortEnvStr)
		CheckError(err)
	}

	ColoniesTLSEnv := os.Getenv("COLONIES_TLS")
	if ColoniesTLSEnv == "true" {
		ColoniesUseTLS = true
		ColoniesInsecure = false
	} else if ColoniesTLSEnv == "false" {
		ColoniesUseTLS = false
		ColoniesInsecure = true
	}

	VerboseEnv := os.Getenv("COLONIES_VERBOSE")
	if VerboseEnv == "true" {
		Verbose = true
	} else if VerboseEnv == "false" {
		Verbose = false
	}

	if ColonyName == "" {
		ColonyName = os.Getenv("COLONIES_COLONY_NAME")
	}
	if ColonyName == "" {
		CheckError(errors.New("Unknown Colony name"))
	}

	if ColonyPrvKey == "" {
		ColonyPrvKey = os.Getenv("COLONIES_COLONY_PRVKEY")
	}

	if ExecutorID == "" {
		ExecutorID = os.Getenv("COLONIES_EXECUTOR_ID")
	}

	if ExecutorName == "" {
		ExecutorName = os.Getenv("COLONIES_EXECUTOR_NAME")
	}
	if ExecutorName == "" {
		CheckError(errors.New("Unknown Executor name"))
	}

	ExecutorType = os.Getenv("EXECUTOR_TYPE")
	if ExecutorType == "" {
		CheckError(errors.New("Executor type not specifed"))
	}

	FsDir = os.Getenv("EXECUTOR_FS_DIR")
	SWName = os.Getenv("EXECUTOR_SW_NAME")
	SWType = os.Getenv("EXECUTOR_SW_TYPE")
	SWVersion = os.Getenv("EXECUTOR_SW_VERSION")
	HWCPU = os.Getenv("EXECUTOR_HW_CPU")
	HWModel = os.Getenv("EXECUTOR_HW_MODEL")
	hwNodesStr := os.Getenv("EXECUTOR_HW_NODES")
	HWNodes, err = strconv.Atoi(hwNodesStr)
	CheckError(err)

	HWMem = os.Getenv("EXECUTOR_HW_MEM")
	HWStorage = os.Getenv("EXECUTOR_HW_STORAGE")
	hwGPUCountStr := os.Getenv("EXECUTOR_HW_GPU_COUNT")
	HWGPUCount, err = strconv.Atoi(hwGPUCountStr)
	CheckError(err)
	hwGPUNodeCountStr := os.Getenv("EXECUTOR_HW_GPU_NODES_COUNT")
	HWGPUNodeCount, err = strconv.Atoi(hwGPUNodeCountStr)
	CheckError(err)

	HWGPUName = os.Getenv("EXECUTOR_HW_GPU_NAME")
	HWGPUMem = os.Getenv("EXECUTOR_HW_GPU_MEM")
	LocDesc = os.Getenv("EXECUTOR_LOCATION_DESC")

	longStr := os.Getenv("EXECUTOR_LOCATION_LONG")
	Long, err = strconv.ParseFloat(longStr, 64)
	if err != nil {
		log.Error("Failed to set location long")
	}

	latStr := os.Getenv("EXECUTOR_LOCATION_LAT")
	Lat, err = strconv.ParseFloat(latStr, 64)
	if err != nil {
		log.Error("Failed to set location long")
	}

	if ExecutorPrvKey == "" {
		ExecutorPrvKey = os.Getenv("COLONIES_PRVKEY")
	}

	addDebugLogsStr := os.Getenv("EXECUTOR_ADD_DEBUG_LOGS")
	AddDebugLogs = false
	if addDebugLogsStr == "true" {
		AddDebugLogs = true
	}

	parallelContainersStr := os.Getenv("EXECUTOR_PARALLEL_CONTAINERS")
	ParallelContainers = false
	if parallelContainersStr == "true" {
		ParallelContainers = true
	}

	gpuStr := os.Getenv("EXECUTOR_GPU")
	GPU = false
	if gpuStr == "true" {
		GPU = true
	}
}

func CheckError(err error) {
	if err != nil {
		log.WithFields(log.Fields{"Error": err, "BuildVersion": build.BuildVersion, "BuildTime": build.BuildTime}).Error(err.Error())
		os.Exit(-1)
	}
}
