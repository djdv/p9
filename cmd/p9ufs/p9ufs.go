// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Binary p9ufs provides a local 9P2000.L server for the p9 package.
//
// To use, first start the server:
//
//	p9ufs 127.0.0.1:3333
//
// Then, connect using the Linux 9P filesystem:
//
//	mount -t 9p -o trans=tcp,port=3333 127.0.0.1 /mnt
package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/hugelgupf/p9/fsimpl/localfs"
	"github.com/hugelgupf/p9/p9"
	"github.com/u-root/uio/ulog"
)

const (
	success int = iota
	failure
	misuse
)

func main() {
	var (
		flagSet = newFlagSet()
		verbose = flagSet.Bool("v", false, "verbose logging")
		root    = flagSet.String("root", "/", "root dir of file system to expose")
		unix    = flagSet.Bool("unix", false, "use unix domain socket instead of TCP")
		network string
	)
	parseFlags(flagSet, os.Args[1:])
	if *unix {
		network = "unix"
	} else {
		network = "tcp"
	}

	// Bind and listen on the socket.
	serverSocket, err := net.Listen(network, flagSet.Args()[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "err binding: %v\n", err)
		os.Exit(failure)
	}

	var opts []p9.ServerOpt
	if *verbose {
		opts = append(opts, p9.WithServerLogger(ulog.Log))
	}
	// Run the server.
	s := p9.NewServer(localfs.Attacher(*root), opts...)
	s.Serve(serverSocket)
}

func newFlagSet() *flag.FlagSet {
	var (
		execName = commandName(os.Args[0])
		flagSet  = flag.NewFlagSet(execName, flag.ContinueOnError)
	)
	flagSet.Usage = newUsage(flagSet)
	return flagSet
}

// commandName normalizes execPath, by
// removing path components, file extensions, etc.
func commandName(execPath string) string {
	baseName := filepath.Base(execPath)
	return strings.TrimSuffix(
		baseName,
		filepath.Ext(baseName),
	)
}

func newUsage(flagSet *flag.FlagSet) func() {
	return func() {
		var (
			output = flagSet.Output()
			cmd    = commandName(os.Args[0])
		)
		io.WriteString(output, fmt.Sprintf(
			"%s - local 9P2000.L server in userspace\n\n"+
				"usage: %s [flags] <bind-addr:port>\n\n"+
				"flags:\n",
			cmd, cmd,
		))
		flagSet.PrintDefaults()
	}
}

func parseFlags(flagSet *flag.FlagSet, arguments []string) {
	if err := flagSet.Parse(arguments); err != nil {
		if err == flag.ErrHelp {
			os.Exit(success)
		}
		os.Exit(misuse)
	}
	if len(flagSet.Args()) != 1 {
		flagSet.Usage()
		os.Exit(misuse)
	}
}
