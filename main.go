package main

import (
	"Remote/config"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var port int                                                //Stores the port used to listen on
	var help bool                                               //Signals if the help should be printed or not
	flag.IntVar(&port, "p", 8080, "Listens to the given port.") //Reads the port flag from the arguments (-p)
	flag.BoolVar(&help, "h", false, "Prints the help message.") //Reads the help flag from the arguments (-h)

	flag.Parse() //Parse all the flags given

	if help { //if the help flag is set
		fmt.Print(generateHelp()) //print the help
		return                    //and quit the program
	}

	r := config.SetRoutes()
	r.Run(":" + strconv.Itoa(port))
}

// Generates a string containing the explanation for the flags of this program.
func generateHelp() string {
	helpMessage := "-p <portnumber>\tListen to the given port. Defaults to 8080.\n"
	helpMessage += "-h\tPrints this help message.\n"
	return helpMessage
}
