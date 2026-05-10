package app

import "github.com/urfave/cli"

// Retrona aplicacao de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de linha de comando"
	app.Usage = "Buscar IPs e Nomes de Servidor na internet"
	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Buscar IPs de enderecos na internet",
			Flags: []cli.Flag{
				
			},
		}
	}

	return app
}
