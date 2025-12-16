package main

import (
	"assembly_line/service"
)

func main() {
	// 建立流水線，5個員工
	assembly_line := service.NewAssemblyLine(5)

	// 新增物品：每種物品各10件
	assembly_line.AddItems(10, 10, 10)

	// 啟動流水線
	assembly_line.Start()
}
