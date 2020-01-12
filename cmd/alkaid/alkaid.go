/*
 * Copyright 2020. The Alkaid Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 *
 * Alkaid is a BaaS service based on Hyperledger Fabric.
 *
 */

package main

import (
	"github.com/yakumioto/glog"

	"github.com/yakumioto/alkaid/internal/config"
)

func main() {
	cmd := config.InitConfig()
	if err := cmd.Execute(); err != nil {
		glog.Fatalln(err)
	}
}
