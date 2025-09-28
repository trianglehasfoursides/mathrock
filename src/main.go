package main

import (
	"context"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/wish"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

func main() {
	// setup
	if err := godotenv.Load(); err != nil {
		return
	}

	// index
	route := gin.Default()

	// handle subdomain
	route.Use(domain)
	subdomain(route, "blog", func(rg *gin.RouterGroup) {})

	//
	route.GET("/", proxy)

	group, _ := errgroup.WithContext(context.TODO())
	group.Go(func() error {
		server := &http.Server{
			Addr:    "",
			Handler: route,
		}

		if err := server.ListenAndServe(); err != nil {
			return err
		}

		return nil
	})

	group.Go(func() error {
		server, err := wish.NewServer()
		if err != nil {
			return err
		}

		if err := server.ListenAndServe(); err != nil {
			return err
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		log.Fatal(err.Error())
	}
}

func proxy(ctx *gin.Context) {
	target, err := url.Parse(os.Getenv(""))
	if err != nil {
		log.Fatal(err.Error())
	}

	prx := httputil.NewSingleHostReverseProxy(target)
	prx.ServeHTTP(ctx.Writer, ctx.Request)
}

func subdomain(r *gin.Engine, subdomain string, routes func(*gin.RouterGroup)) {
	g := r.Group("/")
	g.Use(func(c *gin.Context) {
		sd, _ := c.Get("subdomain")
		if sd == subdomain {
			c.Next()
		} else {
			c.Abort()
		}
	})
	routes(g)
}

func domain(ctx *gin.Context) {
	host := ctx.Request.Host
	honly, _, err := net.SplitHostPort(host)
	if err != nil {
		honly = host
	}

	parts := strings.Split(honly, ".")
	if len(parts) > 2 {
		ctx.Set("subdomain", parts[0])
	} else {
		ctx.Set("subdomain", "")
	}

	ctx.Next()
}
