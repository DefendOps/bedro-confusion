package main

import (
	"context"
	"fmt"
	"os"

	"github.com/defendops/bedro-confuser/pkg/cmd/root"
	"github.com/defendops/bedro-confuser/pkg/utils"
)

type exitCode int

const (
	exitOK      exitCode = 0
	exitError   exitCode = 1
	exitCancel  exitCode = 2
	exitAuth    exitCode = 4
	exitPending exitCode = 8
)

func main() {
	code := mainCLI()
	os.Exit(int(code))
}

func mainCLI() exitCode {
	ctx := context.Background()
	
	utils.PrintBanner()
	
	rootCmd, err := root.NewCmdRoot()
	if err != nil {
		fmt.Printf("failed to create root command: %s\n", err)
		return exitError
	}

	expandedArgs := []string{}
	if len(os.Args) > 0 {
		expandedArgs = os.Args[1:]
	}

	if len(expandedArgs) >= 1 && expandedArgs[0] == "help" {
		expandedArgs = expandedArgs[1:]
		expandedArgs = append(expandedArgs, "--help")
	}

	rootCmd.SetArgs(expandedArgs)

	if _, err := rootCmd.ExecuteContextC(ctx); err != nil {
	}
	
	return exitOK
}