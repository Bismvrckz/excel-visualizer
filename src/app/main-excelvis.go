package main

import "excelvis/app/routes"

func main() {

	s := routes.ServiceEXCELVIS()

	s.Start(":9070")
}
