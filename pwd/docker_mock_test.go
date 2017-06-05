package pwd

import (
	"io"
	"net"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/play-with-docker/play-with-docker/docker"
)

type mockDocker struct {
	createNetwork   func(string) error
	connectNetwork  func(container, network, ip string) (string, error)
	containerResize func(string, uint, uint) error
	createContainer func(opts docker.CreateContainerOpts) (string, error)
}

func (m *mockDocker) CreateNetwork(id string) error {
	if m.createNetwork == nil {
		return nil
	}
	return m.createNetwork(id)
}
func (m *mockDocker) ConnectNetwork(container, network, ip string) (string, error) {
	if m.connectNetwork == nil {
		return "10.0.0.1", nil
	}
	return m.connectNetwork(container, network, ip)
}

func (m *mockDocker) GetDaemonInfo() (types.Info, error) {
	return types.Info{}, nil
}

func (m *mockDocker) GetSwarmPorts() ([]string, []uint16, error) {
	return []string{}, []uint16{}, nil
}
func (m *mockDocker) GetPorts() ([]uint16, error) {
	return []uint16{}, nil
}
func (m *mockDocker) GetContainerStats(name string) (io.ReadCloser, error) {
	return nil, nil
}
func (m *mockDocker) ContainerResize(name string, rows, cols uint) error {
	if m.containerResize != nil {
		return m.containerResize(name, rows, cols)
	}
	return nil
}
func (m *mockDocker) CreateAttachConnection(name string) (net.Conn, error) {
	return &mockConn{}, nil
}
func (m *mockDocker) CopyToContainer(containerName, destination, fileName string, content io.Reader) error {
	return nil
}
func (m *mockDocker) DeleteContainer(id string) error {
	return nil
}
func (m *mockDocker) CreateContainer(opts docker.CreateContainerOpts) (string, error) {
	if m.createContainer != nil {
		return m.createContainer(opts)
	}
	return "10.0.0.1", nil
}
func (m *mockDocker) ExecAttach(instanceName string, command []string, out io.Writer) (int, error) {
	return 0, nil
}
func (m *mockDocker) DisconnectNetwork(containerId, networkId string) error {
	return nil
}
func (m *mockDocker) DeleteNetwork(id string) error {
	return nil
}
func (m *mockDocker) Exec(instanceName string, command []string) (int, error) {
	return 0, nil
}

type mockConn struct {
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	return len(b), nil
}

func (m *mockConn) Write(b []byte) (n int, err error) {
	return len(b), nil
}

func (m *mockConn) Close() error {
	return nil
}

func (m *mockConn) LocalAddr() net.Addr {
	return &net.IPAddr{}
}

func (m *mockConn) RemoteAddr() net.Addr {
	return &net.IPAddr{}
}

func (m *mockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}