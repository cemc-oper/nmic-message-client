package app

import (
	"encoding/json"
	"fmt"
	nmic_message_client "github.com/nwpc-oper/nmic-message-client"
	"github.com/nwpc-oper/nmic-message-client/sender"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"log"
	"os"
	"time"
)

func init() {
	rootCmd.AddCommand(sendCmd)
}

var (
	apiKey = ""

	modelName = ""

	messageType = nmic_message_client.FlowMessageType
	target      = ""
	status      = "1"

	prodType         = nmic_message_client.Grib2ProductionType
	absoluteDataName = ""
	startTime        = ""
	forecastTime     = ""

	debug       = false
	disableSend = false
	ignoreError = false
	help        = false

	postUrl = "http://smart-view.nmic.cma/store/openapi/v2/logs/push_batch"
)

func checkFlag(name string, value string) {
	if value == "" {
		log.Fatalf("%s option is required", name)
	}
}

var sendCmd = &cobra.Command{
	Use:                "send",
	Short:              "Send message to NMIC",
	Long:               "Send message to NMIC",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		var sendFlagSet = pflag.NewFlagSet("send", pflag.ContinueOnError)
		sendFlagSet.ParseErrorsWhitelist = pflag.ParseErrorsWhitelist{UnknownFlags: true}
		sendFlagSet.SortFlags = false

		sendFlagSet.StringVar(&apiKey, "api-key", "", "API key")

		sendFlagSet.StringVar(&modelName, "model-name", "GRAPES_GFS", "model name")

		sendFlagSet.StringVar(&messageType, "type", "", "message type")
		sendFlagSet.StringVar(&target, "target", "", "message target, Send")
		sendFlagSet.StringVar(&status, "status", "0", "status")

		sendFlagSet.StringVar(&prodType, "prod-type", nmic_message_client.Grib2ProductionType,
			"production type")
		sendFlagSet.StringVar(&absoluteDataName, "absolute-data-name", "", "absolute data name")
		sendFlagSet.StringVar(&startTime, "start-time", "", "start time, YYYYMMDDHH")
		sendFlagSet.StringVar(&forecastTime, "forecast-time", "", "forecast time, 003h")

		sendFlagSet.BoolVar(&debug, "debug", false, "show debug information")
		sendFlagSet.BoolVar(&disableSend, "disable-send", false, "disable message send.")
		sendFlagSet.BoolVar(&ignoreError, "ignore-error", false,
			"ignore error. Should be open in operation systems.")
		sendFlagSet.BoolVar(&help, "help", false,
			"show help information.")

		if err := sendFlagSet.Parse(args); err != nil {
			cmd.Usage()
			log.Fatal(err)
		}

		// check if there are non-flag arguments in the command line
		cmds := sendFlagSet.Args()
		if len(cmds) > 0 {
			cmd.Usage()
			log.Fatalf("unknown command: %s", cmds[0])
		}

		// short-circuit on help
		help, err := sendFlagSet.GetBool("help")
		if err != nil {
			log.Fatal(`"help" flag is non-bool, programmer error, please correct`)
		}

		if help {
			cmd.Help()
			fmt.Printf("%s\n", sendFlagSet.FlagUsages())
			return
		}

		// check required flags
		requiredOptions := []struct {
			Name  string
			Value string
		}{
			{"api-key", apiKey},
			{"model-name", modelName},
			{"target", target},
			{"absolute-data-ame", absoluteDataName},
			{"start-time", startTime},
			{"forecast-time", forecastTime},
		}

		for _, requiredOption := range requiredOptions {
			checkFlag(requiredOption.Name, requiredOption.Value)
		}

		if debug {
			fmt.Printf("Version %s (%s)\n", Version, GitCommit)
			fmt.Printf("Build at %s\n", BuildTime)
		}

		modelInfo := nmic_message_client.CreateGrapesGfsGmfModelInfo()

		startTimeObject, _ := time.Parse("2006010215", startTime)
		forecastTimeObject, _ := time.ParseDuration(forecastTime)
		productionInfo := nmic_message_client.CreateProductionFileInfo(
			nmic_message_client.Grib2ProductionType,
			absoluteDataName,
			startTimeObject,
			forecastTimeObject)

		flowInfo := nmic_message_client.CreateFlow(
			nmic_message_client.FlowMessageType,
			fmt.Sprintf("%s message", modelName),
			target,
			status)

		message := nmic_message_client.CreateProdFileMessage(
			modelInfo,
			productionInfo,
			flowInfo)

		blob, err := json.MarshalIndent(message, "", "  ")

		if err != nil {
			f := os.Stderr
			returnCode := 2
			if ignoreError {
				f = os.Stdout
				returnCode = 0
			}
			fmt.Fprintf(f, "create message error: %s\n", err)
			os.Exit(returnCode)
		}

		if debug {
			fmt.Printf("message:\n")
			fmt.Printf("%s\n", blob)
		}

		if disableSend {
			if debug {
				fmt.Printf("disable send.\n")
				fmt.Printf("Bye.\n")
			}
			return
		}

		var s sender.Sender
		s = &sender.HttpSender{
			PostUrl:        fmt.Sprintf("%sapikey=%s", postUrl, apiKey),
			RequestTimeout: 10 * time.Second,
		}

		err = s.SendMessage(blob)

		if err != nil {
			f := os.Stderr
			returnCode := 4
			if ignoreError {
				f = os.Stdout
				returnCode = 0
			}
			fmt.Fprintf(f, "send message failed: %s\n", err)
			os.Exit(returnCode)
		}
	},
}
