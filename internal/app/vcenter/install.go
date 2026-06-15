package vcenter

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/venafi/vmware-vcenter-connector/internal/app/domain"
	"go.uber.org/zap"
)

// InstallCertificateBundleRequest contains the request details for installing a certificate bundle.
type InstallCertificateBundleRequest struct {
	Connection           *domain.Connection       `json:"connection"`
	CertificateBundle    domain.CertificateBundle `json:"certificateBundle"`
	InstallationKeystore domain.Keystore          `json:"keystore"`
}

// InstallCertificateBundleResponse contains the response for an install request.
type InstallCertificateBundleResponse struct {
	InstallationKeystore domain.Keystore `json:"keystore"`
}

// HandleInstallCertificateBundle attempts to install a certificate bundle.
func (svc *WebhookServiceImpl) HandleInstallCertificateBundle(c echo.Context) error {
	req := InstallCertificateBundleRequest{}
	if err := c.Bind(&req); err != nil {
		zap.L().Error("invalid request, failed to unmarshall json", zap.Error(err))
		return c.String(http.StatusBadRequest, fmt.Sprintf("failed to unmarshall json: %s", err.Error()))
	}

	client := svc.ClientServices.NewClient(req.Connection, req.InstallationKeystore.Tenant)
	err := svc.ClientServices.Connect(client)
	defer svc.ClientServices.Close(client)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// TODO: install the certificate, private key, and certificate chain into vCenter.
	zap.L().Info("installCertificateBundle received for vCenter", zap.String("address", req.Connection.HostnameOrAddress), zap.String("certificateName", req.InstallationKeystore.CertificateName))

	res := InstallCertificateBundleResponse{InstallationKeystore: req.InstallationKeystore}
	return c.JSON(http.StatusOK, &res)
}

