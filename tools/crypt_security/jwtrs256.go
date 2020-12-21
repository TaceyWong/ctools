package crypt_security

import (
	"os"
	"log"
	"github.com/TaceyWong/ctools/utils"
	"github.com/urfave/cli"
)

// JWTRS256GeneratorCMD 生成jwt-rs256公钥私钥的命令界面
var JWTRS256GeneratorCMD = cli.Command{
	Name:     "jwt_rs256",
	Aliases:  []string{"jwt_256"},
	Usage:    "生成jwt-rs256公钥私钥",
	Category: "加密安全",
	Action: func(c *cli.Context) error {
		certsDir := "certs"
		os.RemoveAll(certsDir)
		os.MkdirAll(certsDir, 0755)
		genPrivateCMDStr := `ssh-keygen -t rsa -b 2048 -f certs/jwt-rs256.key`
		genPublicCMDStr := `openssl rsa  -in certs/jwt-rs256.key -pubout -outform PEM -out certs/jwt-rs256.key.pub `
		stdInfo, errInfo, err := utils.ExecShellCMD("sh", genPrivateCMDStr)
		if err != nil {
			log.Fatalln(stdInfo,errInfo)
			return err
		}
		stdInfo, errInfo, err = utils.ExecShellCMD("sh", genPublicCMDStr)
		if err != nil {
			log.Fatalln(stdInfo,errInfo)
			return err
		}
		println("公钥和私钥在", certsDir, "中")
		return nil
	},
}
