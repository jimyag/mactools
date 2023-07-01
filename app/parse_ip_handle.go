/*
 * Copyright (c) 2023 by jimyag, All Rights Reserved.
 * Licensed under the MIT License. See LICENSE file in the project root for license information.
 */

package app

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/jimyag/mactools/log"
	"github.com/jimyag/mactools/notification"
	"github.com/jimyag/mactools/pasteboard"
)

var (
	templateUrl             = "http://ip-api.com/json/%s?fields=status,message,country,city,isp,reverse,query&lang=zh-CN"
	_, private24BitBlock, _ = net.ParseCIDR("10.0.0.0/8")
	_, private20BitBlock, _ = net.ParseCIDR("172.16.0.0/12")
	_, private16BitBlock, _ = net.ParseCIDR("192.168.0.0/16")
)

func init() {
	handle := &ParseIpHandle{
		client: http.DefaultClient,
	}
	pasteboard.PB.Register(handle)
}

type ParseIpHandle struct {
	ip     net.IP
	client *http.Client
}

func (h *ParseIpHandle) OnCopy(pb *pasteboard.Pasteboard, content string) {
}

func (h *ParseIpHandle) AfterHandle(pb *pasteboard.Pasteboard, res any) {
	if res == nil {
		return
	}
	h.show(res.(*ResponseData))
}

func (h *ParseIpHandle) Handle(pb *pasteboard.Pasteboard, content string) any {
	log.Debug("handle: ", content)
	if content == "" {
		return nil
	}
	ip := net.ParseIP(content)
	if ip == nil {
		return nil
	}
	h.ip = ip

	if h.isPrivateIP() {
		return nil
	}
	respData, err := h.getInfo()
	if err != nil || respData == nil {
		log.Error("get ip error: %v ,%v", err, respData)
		return nil
	}
	if respData.Status != "success" {
		log.Error("get ip error: %v", respData.Message)
		return nil
	}
	return respData
}

func (h *ParseIpHandle) show(content *ResponseData) {
	notification.
		New().
		SetTitle(content.Query).
		SetInformativeText(content.format()).
		Show()
}

func (h *ParseIpHandle) isPrivateIP() bool {
	return private24BitBlock.Contains(h.ip) || private20BitBlock.Contains(h.ip) || private16BitBlock.Contains(h.ip)
}

func (h *ParseIpHandle) getInfo() (*ResponseData, error) {
	u := fmt.Sprintf(templateUrl, h.ip.String())
	resp, err := h.client.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respData := ResponseData{}
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

// ResponseData https://ip-api.com/docs/api:json
type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Country string `json:"country"`
	City    string `json:"city"`
	ISP     string `json:"isp"`
	Reverse string `json:"reverse"`
	Query   string `json:"query"`
}

func (data *ResponseData) format() string {
	if data.Message != "" {
		return data.Message
	}

	str := ""
	if data.Country != "" {
		str = str + "country: " + data.Country + "\n"
	}
	if data.City != "" {
		str = str + "city: " + data.City + "\n"
	}
	if data.ISP != "" {
		str = str + "isp: " + data.ISP + "\n"
	}
	if data.Reverse != "" {
		str = str + "reverse: " + data.Reverse + "\n"
	}
	return str
}
