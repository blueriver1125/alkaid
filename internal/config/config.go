/*
 * Copyright 2020. The Alkaid Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 *
 * Alkaid is a BaaS service based on Hyperledger Fabric.
 *
 */

package config

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/yakumioto/glog"

	"github.com/yakumioto/alkaid/internal/api/routers"
)

var (
	CfgPath  string
	LogLevel string
	DBPath   string
	Address  string
	Port     int
)

func InitConfig() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "alkaid",
		Short: "",
		Long:  "",
		Run:   run,
	}

	rootCmd.Flags().StringVar(&CfgPath, "config", "/var/alkaid", "config path.")
	rootCmd.Flags().StringVar(&LogLevel, "logLevel", "INFO", "log level.")
	rootCmd.Flags().StringVar(&DBPath, "dbPath", "/data/alkaid.db", "sqlite3 db path.")
	rootCmd.Flags().StringVar(&Address, "address", "0.0.0.0", "listening address.")
	rootCmd.Flags().IntVar(&Port, "port", 8080, "listening port.")

	return rootCmd
}

func run(_ *cobra.Command, _ []string) {
	switch LogLevel {
	case "INFO":
		gin.SetMode(gin.ReleaseMode)
		glog.SetLevel(glog.LevelInfo)
	case "WARN":
		glog.SetLevel(glog.LevelWarning)
	case "ERROR":
		glog.SetLevel(glog.LevelError)
	case "DUBE":
		gin.SetMode(gin.DebugMode)
		glog.SetLevel(glog.LevelDebug)
	}

	r := gin.Default()
	routers.AddRouters(r)

	if err := r.Run(Address + ":" + strconv.Itoa(Port)); err != nil {
		glog.Fatalf("service startup failed: %s", err)
	}
}
