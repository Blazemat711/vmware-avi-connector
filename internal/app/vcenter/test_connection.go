package vcenter

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/venafi/vmware-vcenter-connector/internal/app/domain"
	"go.uber.org/zap"
)

// TestConnectionRequest contains the request details for testing connectivity.
type TestConnectionRequest struct {
	Connection *domain.Connection `json:"connection"`
}

// TestConnectionResponse contains the response for a TestConnectionRequest.
type TestConnectionResponse struct {
	Result bool `json:"result"`
}

// HandleTestConnection attempts to connect to a vCenter host.
func (svc *WebhookServiceImpl) HandleTestConnection(c echo.Context) error {
	req := TestConnectionRequest{}
	if err := c.Bind(&req); err != nil {
		zap.L().Error("invalid request, failed to unmarshall json", zap.Error(err))
		return c.String(http.StatusBadRequest, fmt.Sprintf("failed to unmarshall json: %s", err.Error()))
	}

	client := svc.ClientServices.NewClient(req.Connection, "")
	err := svc.ClientServices.Connect(client)
	defer svc.ClientServices.Close(client)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	res := TestConnectionResponse{Result: true}
	zap.L().Info("Success connecting to vCenter", zap.String("address", req.Connection.HostnameOrAddress), zap.Int("port", req.Connection.Port))
	return c.JSON(http.StatusOK, res)
}

