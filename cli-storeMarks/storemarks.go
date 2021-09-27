package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main(){
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "save",
			Value: "no",
			Usage: "should save to databse(yes/no)",
		},
	}
	app.Version = "1.0"
	//define action
	app.Action = func (c *cli.Context) error  {
		var args []string
		if c.NArg() > 0 {
			//fetch arguments in an array
			args = c.Args()
			personName := args[0]
			//fetch arguments in an array
			marks := args[1:]
			log.Println("Person:", personName)
			log.Println("marks:", marks)

		}

		if c.String("save") == "no"{
			log.Println("skipping saving to the database")
		} else {
			//add database logic here
			log.Println("saving to the database", args)
		}
		return nil
		
	}
	app.Run(os.Args)
}