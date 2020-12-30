package brokerhdl

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/na7r1x/acectl/internal/core/domain"
	"github.com/na7r1x/acectl/internal/core/ports"
)

type HTTPHandler struct {
	brokerService ports.BrokerService
}

func NewHttpHandler(brokerService ports.BrokerService) *HTTPHandler {
	return &HTTPHandler{
		brokerService: brokerService,
	}
}

func (hdl *HTTPHandler) Register(c *gin.Context) {
	var body HttpBodyRegister
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	broker := domain.NewBroker(body.Id, time.Now(), body.Host, body.Port, body.Username, body.Password)
	err := hdl.brokerService.Register(broker)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.Status(201)
}

func (hdl *HTTPHandler) Unregister(c *gin.Context) {
	err := hdl.brokerService.Unregister(c.Param("brokerId"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.Status(200)
}

func (hdl *HTTPHandler) List(c *gin.Context) {
	brokers, err := hdl.brokerService.List()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, BuildHttpResponseList(brokers))
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	broker, err := hdl.brokerService.Get(c.Param("brokerId"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, broker)
}

func (hdl *HTTPHandler) Status(c *gin.Context) {
	statusResponse, exitState, err := hdl.brokerService.Status(c.Param("brokerId"), 60)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	status := -1                                     // indeterminate
	if strings.Contains(statusResponse, "BIP1284") { // running
		status = 1
	}
	if strings.Contains(statusResponse, "BIP1285") { // stopped
		status = 0
	}
	c.JSON(200, BuildHttpResponseStatus(status, statusResponse, exitState))
}

func (hdl *HTTPHandler) Stop(c *gin.Context) {
	stopResponse, exitState, err := hdl.brokerService.Stop(c.Param("brokerId"), 60)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	var success bool
	if exitState == "" {
		success = true
	} else {
		success = false
	}
	c.JSON(200, BuildHttpResponseCommand(success, stopResponse, exitState))
}

func (hdl *HTTPHandler) Start(c *gin.Context) {
	startResponse, exitState, err := hdl.brokerService.Start(c.Param("brokerId"), 60)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	if exitState == "" {
		c.JSON(200, BuildHttpResponseCommand(true, startResponse, exitState))
	} else {
		c.JSON(500, BuildHttpResponseCommand(false, startResponse, exitState))
	}
}
