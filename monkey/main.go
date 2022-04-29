package main 

import(
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
)
func main (){
	user,err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey lang ! \n",
	user.Username)
	fmt.Printf("Fell free to type in commands\n")
	repl.Start(os.Stdin,os.Stdout)
}