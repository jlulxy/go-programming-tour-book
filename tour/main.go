package main

import (
	"flag"
	"log"
	"fmt"
	"os"
)

func main(){
	var name string
	flag.Parse()
	goCmd := flag.NewFlagSet("go",flag.ExitOnError)
	goCmd.StringVar(&name,"name","Go 语言","帮助信息")
	phpCmd := flag.NewFlagSet("php",flag.ExitOnError)
	phpCmd.StringVar(&name,"n","PHP语言","帮助信息")
	args := flag.Args()
	fmt.Println(os.Args)
	fmt.Println(args)
	if len(args) == 0 {
		log.Fatalf("please do same canshu")
	}else{
		switch args[0]{
		case "go":
			_ = goCmd.Parse(args[1:])
			fmt.Println(goCmd.Usage)
		case "php":
			_ = phpCmd.Parse(args[1:])
			fmt.Println(phpCmd.Args())

		}
		log.Fatalf("name : %s",name)
	}

}
