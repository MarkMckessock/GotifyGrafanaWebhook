package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotify/plugin-api"
)

// GetGotifyPluginInfo returns gotify plugin info
func GetGotifyPluginInfo() plugin.Info {
	return plugin.Info{
		Name:       "Grafana Webhook",
		ModulePath: "github.com/MarkMckessock/GotifyGrafanaWebhook",
	}
}

// Plugin is plugin instance
type Plugin struct {
	msgHandler plugin.MessageHandler
	basePath   string
}

func (c *Plugin) SetMessageHandler(h plugin.MessageHandler) {
	c.msgHandler = h
}

func (c *Plugin) RegisterWebhook(basePath string, mux *gin.RouterGroup) {
	c.basePath = basePath
	mux.POST("/webhook	", func(c *gin.Context) {
		// Processes webhook and take actions(sending messages, etc.)
	})
}

func (c *Plugin) GetDisplay(location *url.URL) string {
	// baseLocation := &url.URL{
	loc := &url.URL{
		Path: c.basePath,
	}
	if location != nil {
		// If the server location can be determined, make the URL absolute
		loc.Scheme = location.Scheme
		loc.Host = location.Host
	}
	loc = loc.ResolveReference(&url.URL{
		Path: "hook",
	})
	return fmt.Sprintf("Set your webhook URL to %s and you are all set", loc)
}

// Enable implements plugin.Plugin
func (c *Plugin) Enable() error {
	go func() {
		time.Sleep(5 * time.Second)
		c.msgHandler.SendMessage(plugin.Message{
			Message: "The plugin has been enabled for 5 seconds.",
		})
	}()
	return nil
}

// Disable implements plugin.Plugin
func (c *Plugin) Disable() error {
	return nil
}

// NewGotifyPluginInstance creates a plugin instance for a user context.
func NewGotifyPluginInstance(ctx plugin.UserContext) plugin.Plugin {
	return &Plugin{}
}

func main() {
	panic("this should be built as go plugin")
}
