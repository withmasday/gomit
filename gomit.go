package main

import "flag"
import "fmt"
import "os/exec"
import "strings"

func commiter(action, message string) {
	action = strings.ToUpper(action)
	switch action {
	case "NEW":
		exec.Command("git", "commit", "-m", "'✨ NEW : "+ message +"'");
		fmt.Printf("[OKE] : git commit -m '✨ %s : %s' \n", action, message)
	case "FIX":
		exec.Command("git", "commit", "-m", "'🛠 FIX : "+ message +"'");
		fmt.Printf("[OKE] : git commit -m '🛠 %s : %s' \n", action, message)
	case "UPDATE":
		exec.Command("git", "commit", "-m", "'🔨 UPDATE : "+ message +"'");
		fmt.Printf("[OKE] : git commit -m '🔨 %s : %s' \n", action, message)
	case "DOC":
		exec.Command("git", "commit", "-m", "'📝 DOC : "+ message +"'");
		fmt.Printf("[OKE] : git commit -m '📝 %s : %s' \n", action, message)
	case "MERGE":
		exec.Command("git", "commit", "-m", "'🔀 MERGE : "+ message +"'");
		fmt.Printf("[OKE] : git commit -m '🔀 %s : %s' \n", action, message)
	case "DOWN":
		exec.Command("git", "commit", "-m", "'⏬ DOWN : "+ message +"'");
		fmt.Printf("[OKE] : git commit -m '⏬ %s : %s' \n", action, message)
	case "UP":
		exec.Command("git", "commit", "-m", "'⏫ UP : "+ message +"'");
		fmt.Printf("[OKE] : git commit -m '⏫ %s : %s' \n", action, message)
	case "PACKAGE":
		exec.Command("git", "commit", "-m", "'📦 PACKAGE : "+ message +"'");
		fmt.Printf("[OKE] : git commit -m '📦 %s : %s' \n", action, message)
	case "WORKING":
		exec.Command("git", "commit", "-m", "'🚧 WORKING : "+ message +"'");
		fmt.Printf("[OKE] : git commit -m '🚧 %s : %s' \n", action, message)
	default:
		fmt.Printf("[BAD] : git commit -m '🚧 %s : %s' \n", action, message)
	}
}

func main() {
    var action = flag.String("a", "", "Type your action.")
	var message = flag.String("m", "", "Type your commit message.")
    flag.Parse()

	fn := *action
	if fn == "" {
			fmt.Printf("\n--[ gomit ]-- \n\n")
			fmt.Printf("| ✨ : new \t| 📝 : doc \t| ⏫ : up \t| \n")
			fmt.Printf("| 🛠  : fix \t| 🔀 : merge \t| 📦 : package \t| \n")
			fmt.Printf("| 🔨 : update \t| ⏬ : down \t| 🚧 : working \t| \n")

			fmt.Printf("\nUsage : \n")
			fmt.Printf("./gomit -a=up -m='Create some function to file' \n")
	} else {
		commiter(*action, *message)
	}


}