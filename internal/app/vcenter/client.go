package vcenter

import (
	"errors"
	"fmt"

	"github.com/venafi/vmware-vcenter-connector/internal/app/domain"
	"go.uber.org/zap"
)

// ClientServices defines the operations required by the vCenter connector.
type ClientServices interface {
	Close(client *domain.Client)
	Connect(client *domain.Client) error
	NewClient(connection *domain.Connection, tenant string) *domain.Client
}

// VcenterClientsImpl is the default implementation of ClientServices.
type VcenterClientsImpl struct {
}

// NewVcenterClients returns a new VcenterClientsImpl.
func NewVcenterClients() *VcenterClientsImpl {
	return &VcenterClientsImpl{}
}

// Close closes the client session.
func (c *VcenterClientsImpl) Close(client *domain.Client) {
	if client == nil || client.Session == nil {
		return
	}

	// TODO: close the vCenter session here.
	client.Session = nil
}

// Connect attempts to create a new client session.
func (c *VcenterClientsImpl) Connect(client *domain.Client) error {
	if client == nil {
		return errors.New("client is nil")
	}

	if err := validateHostnameOrAddress(client.Connection.HostnameOrAddress); err != nil {
		zap.L().Error("invalid hostname or address", zap.String("hostname", client.Connection.HostnameOrAddress), zap.Error(err))
		return fmt.Errorf("invalid hostname or address: %w", err)
	}

	// TODO: implement vCenter connection logic using the vCenter SDK or API.
	zap.L().Info("connected to vCenter host", zap.String("hostname", client.Connection.HostnameOrAddress), zap.Int("port", client.Connection.Port))

	client.Session = struct{}{}
	return nil
}

func validateHostnameOrAddress(hostnameOrAddress string) error {
	if hostnameOrAddress == "" {
		return errors.New("hostname or address cannot be empty")
	}

	return nil
}

// NewClient creates a new client instance.
func (c *VcenterClientsImpl) NewClient(connection *domain.Connection, tenant string) *domain.Client {
	if connection.Port == 0 {
		connection.Port = 443
	}

	if tenant == "" {
		tenant = "default"
	}

	return &domain.Client{
		Connection: connection,
		Session:    nil,
		Tenant:     tenant,
	}
}

