package vcenter

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/venafi/vmware-vcenter-connector/internal/app/domain"
	"go.uber.org/zap"
)

// ConfigureInstallationEndpointRequest contains the request details for configuring usage of an installed certificate.
type ConfigureInstallationEndpointRequest struct {
	Connection *domain.Connection `json:"connection"`
	Keystore   domain.Keystore    `json:"keystore"`
	Binding    domain.Binding     `json:"binding"`
}

// GetTargetConfigurationRequest contains the request details for retrieving target configuration.
type GetTargetConfigurationRequest struct {
	Connection *domain.Connection `json:"connection"`
}

// GetTargetConfigurationResponse contains the response for the target configuration request.
type GetTargetConfigurationResponse struct {
	TargetConfiguration TargetConfiguration `json:"targetConfiguration"`
}

// HandleConfigureInstallationEndpoint configures a vCenter endpoint to use the installed certificate.
func (svc *WebhookServiceImpl) HandleConfigureInstallationEndpoint(c echo.Context) error {
	req := ConfigureInstallationEndpointRequest{}
	if err := c.Bind(&req); err != nil {
		zap.L().Error("invalid request, failed to unmarshall json", zap.Error(err))
		return c.String(http.StatusBadRequest, fmt.Sprintf("failed to unmarshall json: %s", err.Error()))
	}

	client := svc.ClientServices.NewClient(req.Connection, req.Keystore.Tenant)
	err := svc.ClientServices.Connect(client)
	defer svc.ClientServices.Close(client)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// TODO: configure the vCenter target to use the installed certificate.
	zap.L().Info("configureInstallationEndpoint received for vCenter", zap.String("virtualServiceName", req.Binding.VirtualServiceName), zap.String("certificateName", req.Keystore.CertificateName))
	return c.NoContent(http.StatusOK)
}

// HandleGetTargetConfiguration returns target configuration information.
func (svc *WebhookServiceImpl) HandleGetTargetConfiguration(c echo.Context) error {
	req := GetTargetConfigurationRequest{}
	if err := c.Bind(&req); err != nil {
		zap.L().Error("invalid request, failed to unmarshall json", zap.Error(err))
		return c.String(http.StatusBadRequest, fmt.Sprintf("failed to unmarshall json: %s", err.Error()))
	}

	// TODO: populate vCenter target configuration values.
	res := GetTargetConfigurationResponse{}
	return c.JSON(http.StatusOK, res)
}

// HandleDiscoverCertificates is a placeholder for discovery support.
func (svc *WebhookServiceImpl) HandleDiscoverCertificates(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}

