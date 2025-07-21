/*
 * @Author: Justin
 * @Date: 2025-07-21 03:14:13
 * @filename:
 * @version:
 * @Description:
 * @LastEditTime: 2025-07-21 06:58:54
 */
package main

import (
	"test_server/config"
	"test_server/httpserver"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig("./config/config.toml")

	// 实例化自定义 HttpEngine
	engine := httpserver.NewHttpEngine(cfg.ServerPort)

	// 注册路由与对应处理函数
	engine.Handle("/config", httpserver.HandleConfig)
	engine.Handle("/startChain", httpserver.HandleStartChain)
	engine.Handle("/setupContracts", httpserver.HandleSetupContracts)
	engine.Handle("/startRelayer", httpserver.HandleStartRelayer)

	// 启动 HTTP 服务
	if err := engine.Start(); err != nil {
		panic(err)
	}

}
