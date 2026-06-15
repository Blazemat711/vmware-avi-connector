package vcenter

import "github.com/labstack/echo/v4"

// WebhookServiceImpl implements the connector route handlers.
type WebhookServiceImpl struct {
	ClientServices ClientServices
}

// NewWebhookService returns a new WebhookServiceImpl.
func NewWebhookService(clientServices ClientServices) *WebhookServiceImpl {
	return &WebhookServiceImpl{
		ClientServices: clientServices,
	}
}

// Ensure the WebhookServiceImpl matches the web handler interface.
var _ interface {
	HandleConfigureInstallationEndpoint(c echo.Context) error
	HandleDiscoverCertificates(c echo.Context) error
	HandleGetTargetConfiguration(c echo.Context) error
	HandleInstallCertificateBundle(c echo.Context) error
	HandleTestConnection(c echo.Context) error
} = (*WebhookServiceImpl)(nil)
