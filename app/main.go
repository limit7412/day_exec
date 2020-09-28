package main

import (
	"fmt"
	"os"

	"github.com/limit7412/day_exec/app/usecase"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "dayexec"
	app.Usage = "日付を指定してコマンドを連続で実行"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "cmd",
		},
		cli.IntFlag{
			Name: "y",
		},
		cli.IntFlag{
			Name: "m",
		},
		cli.IntFlag{
			Name: "d",
		},
		cli.IntFlag{
			Name: "period",
		},
		cli.IntFlag{
			Name: "times",
		},
	}

	defer func() {
		if p := recover(); p != nil {
			fmt.Sprintf("task finished because of panic : ", p)
			panic(p)
		}
	}()

	app.Action = run

	err := app.Run(os.Args)
	if err != nil {
		fmt.Sprintln(err.Error())
	}
}

func run(context *cli.Context) {
	err := usecase.Run(
		context.String("cmd"),
		context.Int("y"),
		context.Int("m"),
		context.Int("d"),
		context.Int("period"),
		context.Int("times"),
	)
	if err != nil {
		fmt.Println("exec failed…")
		fmt.Sprintln(err)
	}
}
