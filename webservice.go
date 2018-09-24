package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type webService struct {
	c *echo.Echo
}

func newWebService() *webService {
	ec := echo.New()
	ec.Use(middleware.Logger())
	ec.Use(middleware.Recover())
	ec.Use(middleware.CORS())
	ec.GET("/custom/host", func(c echo.Context) error {
		return c.JSON(http.StatusOK, getCustomHosts())
	})
	ec.PUT("/custom/host/:domain", func(c echo.Context) error {
		err := addCustomHost(c.Param("domain"))
		if err != nil {
			return err
		}
		err = reloadDnsmasq()
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "OK")
	})
	ec.DELETE("/custom/host/:domain", func(c echo.Context) error {
		err := removeCustomHost(c.Param("domain"))
		if err != nil {
			return err
		}
		err = reloadDnsmasq()
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "OK")
	})
	ec.GET("/exclude/host", func(c echo.Context) error {
		return c.JSON(http.StatusOK, getCustomExcludeHosts())
	})
	ec.PUT("/exclude/host/:domain", func(c echo.Context) error {
		err := addCustomExcludeHost(c.Param("domain"))
		if err != nil {
			return err
		}
		err = reloadDnsmasq()
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "OK")
	})
	ec.DELETE("/exclude/host/:domain", func(c echo.Context) error {
		err := removeCustomExcludeHost(c.Param("domain"))
		if err != nil {
			return err
		}
		err = reloadDnsmasq()
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "OK")
	})
	ec.GET("/standard/update/date", func(c echo.Context) error {
		return c.String(http.StatusOK, strconv.FormatUint(appSetting.standardHostsUpdateDate, 10))
	})
	ec.PUT("/standard/update/action", func(c echo.Context) error {
		err := updateStandardHosts()
		if err != nil {
			return err
		}
		err = reloadDnsmasq()
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "OK")
	})
	return &webService{
		c: ec,
	}
}

func (w *webService) start() {
	addr := fmt.Sprintf("%s:%d", appConfig.ServiceListenHost, appConfig.ServiceListenPort)
	go func() {
		infoLogger.Println("server listens at http://", addr)
		if err := w.c.Start(addr); err != nil {
			infoLogger.Println("shutting down the server")
		}
	}()
}

func (w *webService) stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := w.c.Shutdown(ctx); err != nil {
		errorLogger.Fatal(err)
	}
}
