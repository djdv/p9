//go:build !race && linux
// +build !race,linux

package localfs

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"

	"github.com/djdv/p9/fsimpl/test/vmdriver"
	"github.com/djdv/p9/p9"
	"github.com/hugelgupf/vmtest"
	"github.com/hugelgupf/vmtest/qemu"
	"github.com/u-root/u-root/pkg/uroot"
	"github.com/u-root/uio/ulog/ulogtest"
)

func TestIntegration(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "localfs-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	serverSocket, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("err binding: %v", err)
	}
	defer serverSocket.Close()
	serverPort := serverSocket.Addr().(*net.TCPAddr).Port

	// Run the server.
	s := p9.NewServer(Attacher(tempDir), p9.WithServerLogger(ulogtest.Logger{TB: t}))
	go s.Serve(serverSocket)

	// Run the read-write tests from fsimpl/test/rwvm.
	vmtest.RunGoTestsInVM(t, []string{"github.com/djdv/p9/fsimpl/test/rwvmtests"}, &vmtest.UrootFSOptions{
		BuildOpts: uroot.Opts{
			Commands: uroot.BusyBoxCmds(
				"github.com/u-root/u-root/cmds/core/ls",
				"github.com/u-root/u-root/cmds/core/dhclient",
			),
			ExtraFiles: []string{
				"/usr/bin/dd:bin/dd",
			},
		},
		VMOptions: vmtest.VMOptions{
			QEMUOpts: qemu.Options{
				Devices: []qemu.Device{
					vmdriver.HostNetwork{
						Net: &net.IPNet{
							// 192.168.0.0/24
							IP:   net.IP{192, 168, 0, 0},
							Mask: net.CIDRMask(24, 32),
						},
					},
				},
				KernelArgs: fmt.Sprintf("P9_PORT=%d P9_TARGET=192.168.0.2", serverPort),
				Timeout:    30 * time.Second,
			},
		},
	})
}

func TestBenchmark(t *testing.T) {
	// Needs to definitely be in a tmpfs for performance testing.
	tempDir, err := ioutil.TempDir("/dev/shm", "localfs-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	serverSocket, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("err binding: %v", err)
	}
	defer serverSocket.Close()
	serverPort := serverSocket.Addr().(*net.TCPAddr).Port

	// Run the server. No logger -- slows down the benchmark.
	s := p9.NewServer(Attacher(tempDir)) //, p9.WithServerLogger(ulogtest.Logger{TB: t}))
	go s.Serve(serverSocket)

	// Run the read-write tests from fsimpl/test/rwvm.
	vmtest.RunGoTestsInVM(t, []string{"github.com/djdv/p9/fsimpl/test/benchmark"}, &vmtest.UrootFSOptions{
		BuildOpts: uroot.Opts{
			Commands: uroot.BusyBoxCmds(
				"github.com/u-root/u-root/cmds/core/ls",
				"github.com/u-root/u-root/cmds/core/dhclient",
			),
			ExtraFiles: []string{
				"/usr/bin/dd:bin/dd",
			},
		},
		VMOptions: vmtest.VMOptions{
			QEMUOpts: qemu.Options{
				Devices: []qemu.Device{
					vmdriver.HostNetwork{
						Net: &net.IPNet{
							// 192.168.0.0/24
							IP:   net.IP{192, 168, 0, 0},
							Mask: net.CIDRMask(24, 32),
						},
					},
				},
				KernelArgs: fmt.Sprintf("P9_PORT=%d P9_TARGET=192.168.0.2", serverPort),
				Timeout:    30 * time.Second,
			},
		},
	})
}
