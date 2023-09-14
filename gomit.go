package main

import "flag"
import "fmt"
import "os/exec"
import "strings"

func commiter(action, message string) {
	action = strings.ToUpper(action)
	switch action {
	case "NEW":
		new := exec.Command("git", "commit", "-m", "✨ NEW : "+ message);
		new.CombinedOutput()
		fmt.Printf("[OKE] : git commit -m '✨ %s : %s' \n", action, message)
	case "FIX":
		fix := exec.Command("git", "commit", "-m", "'🛠 FIX : "+ message +"'");
		fix.CombinedOutput()
		fmt.Printf("[OKE] : git commit -m '🛠 %s : %s' \n", action, message)
	case "UPDATE":
		update := exec.Command("git", "commit", "-m", "'🔨 UPDATE : "+ message +"'");
		update.CombinedOutput()
		fmt.Printf("[OKE] : git commit -m '🔨 %s : %s' \n", action, message)
	case "DOC":
		doc := exec.Command("git", "commit", "-m", "'📝 DOC : "+ message +"'");
		doc.CombinedOutput()
		fmt.Printf("[OKE] : git commit -m '📝 %s : %s' \n", action, message)
	case "MERGE":
		merge := exec.Command("git", "commit", "-m", "'🔀 MERGE : "+ message +"'");
		merge.CombinedOutput()
		fmt.Printf("[OKE] : git commit -m '🔀 %s : %s' \n", action, message)
	case "DOWN":
		down := exec.Command("git", "commit", "-m", "'⏬ DOWN : "+ message +"'");
		down.CombinedOutput()
		fmt.Printf("[OKE] : git commit -m '⏬ %s : %s' \n", action, message)
	case "UP":
		up := exec.Command("git", "commit", "-m", "'⏫ UP : "+ message +"'");
		up.CombinedOutput()
		fmt.Printf("[OKE] : git commit -m '⏫ %s : %s' \n", action, message)
	case "PACKAGE":
		pkg := exec.Command("git", "commit", "-m", "'📦 PACKAGE : "+ message +"'");
		pkg.CombinedOutput()
		fmt.Printf("[OKE] : git commit -m '📦 %s : %s' \n", action, message)
	case "WORKING":
		working := exec.Command("git", "commit", "-m", "'🚧 WORKING : "+ message +"'");
		working.CombinedOutput()
		fmt.Printf("[OKE] : git commit -m '🚧 %s : %s' \n", action, message)
	default:
		fmt.Printf("[BAD] : git commit -m '%s : %s' \n", action, message)
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