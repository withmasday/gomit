package main

import "flag"
import "fmt"
import "os/exec"
import "strings"

func commiter(action, message string) {
	action = strings.ToUpper(action)
	switch action {
	case "NEW":
		fmt.Printf("[OKE] : git commit -m '✨ %s : %s' \n", action, message)
		cmd := exec.Command("git", "commit", "-m", "'✨ NEW : "+ message +"'");
		out, err := cmd.CombinedOutput()

		if err != nil && err.Error() != "exit status 1" {
		    fmt.Printf("issue failed with error: %s\n", err)
		}

		fmt.Printf("combined out:\n%s\n", string(out))
	default:
		fmt.Printf("\n ❌ : Error, commit action format failed.\n")
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