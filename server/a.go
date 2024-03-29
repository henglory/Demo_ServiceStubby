package server

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/henglory/Demo_ServiceStubby/handler"
	"github.com/henglory/Demo_ServiceStubby/service"
	"github.com/henglory/Demo_ServiceStubby/spec"
)

func doA(s service.Service, c *gin.Context) {
	var data spec.AReq
	raw, err := c.GetRawData()
	if err != nil {
		c.JSON(500, spec.ARes{
			StatusCode: "99",
		})
	}
	err = json.Unmarshal(raw, &data)
	if err != nil {
		c.JSON(500, spec.ARes{
			StatusCode: "99",
		})
	}
	res := handler.DoA(s, data)
	c.JSON(200, res)
}
