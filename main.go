package main
import "fmt"
func main() {
	printWelcome()
	printGreeting(getResponseToPrompt())
	fmt.Println("Let's go on an adventure!")
	promptForRandomOrSpecificDestination()
}

func printWelcome() {
	fmt.Println("Welcome to the Solar System! \n There are 8 planets to explore. \n What is your name?")

}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}


func promptForRandomOrSpecificDestination() {
	var choice string
	fmt.Println("Shall I randomly choose a planet for you to visit? (Y or N)")
	for choice != "Y" && choice != "N" {
		choice = getResponseToPrompt()
		if choice == "Y" {
			travel()
		} else if choice == "N" {
			fmt.Println("Name the planet you would like to visit.")
			planetName := getResponseToPrompt()
			travel(planetName)
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}

}
func getResponseToPrompt() string {
	var response string	
	fmt.Scan(&response)
	return response
}

func travel(location ...string) {
	if len(location) == 1 {
		
		fmt.Printf("Traveling to %s...\n", location[0])
		fmt.Printf("Arrived at %s. A very cold planet, furthest from the sun.", location[0])
	} else {
		fmt.Println("Traveling to Jupiter...")
		fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
	}

}