package rpc

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

var rpcMode = debugCode
var modeName = DebugMode
var DefaultWriter io.Writer = os.Stdout
var DefaultErrorWriter io.Writer = os.Stderr


const (
	debugCode = iota
	releaseCode
	testCode
)

func SetMode(value string) {
	if value == "" {
		value = DebugMode
	}

	switch value {
	case DebugMode:
		rpcMode = debugCode
	case ReleaseMode:
		rpcMode = releaseCode
	case TestMode:
		rpcMode = testCode
	default:
		panic("micro mode unknown: " + value + " (available mode: debug release test)")
	}

	modeName = value
}

func debugPrint(format string, values ...interface{}) {
	if IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(DefaultWriter, "[Micro-debug] "+format, values...)
	}
}

func debugPrintError(err error) {
	if err != nil {
		if IsDebugging() {
			fmt.Fprintf(DefaultErrorWriter, "[Micro-debug] [ERROR] %v\n", err)
		}
	}
}


func IsDebugging() bool {
	return rpcMode == debugCode
}



func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			debugPrint("Environment variable PORT=\"%s\"", port)
			return ":" + port
		}
		debugPrint("Environment variable PORT is undefined. Using port :8080 by default")
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}
