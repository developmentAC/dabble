package main

// author: Oliver Bonham-Carter
// mail: obonhamcarter@allegheny.edu
// date: 21 November 2022
// comment: A basic shell to learn how a shell might work and to spend some quality time with Golang. Uh-huh

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
)

func printMenu() {
	// prints the system commands
	syscmds_map := make(map[string]string) // define a map (i.e., a dictionary)
	syscmds_map["hello"] = ":Say hello"
	syscmds_map["pwd"] = ":Show current path"
	syscmds_map["exit"] = ":End program"
	syscmds_map["help"] = ":Show the menu of commands"
	syscmds_map["square"] = ":Draw a square"
	syscmds_map["ffile"] = ":Create a file with a fortune in it"
	syscmds_map["ran"] = ":Create n random numbers between min and max bounds"
	syscmds_map["prime"] = ":Create n random prime numbers between min and max bounds"
	fmt.Print(("\n [+] System commands:"))
	for i := range syscmds_map {
		print("\n\t [+=+] ", i, " ", syscmds_map[i])
	}
	fmt.Print("\n")
}

func main() {
	fmt.Print("\t :: OBC's Dabble Terminal ::\n")

	fmt.Print("\n [+] Enter a system command (help) :")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {

		// fmt.Print("\n\t [+] type 'help' for list of commands ")
		if input.Text() == "exit" {
			break
		}
		getArguments(input.Text()) // determine what to do with this command

		if input.Text() == "exit" {
			fmt.Print(" [+] Exiting...\n")
			break
		}
		if input.Text() == "help" {
			printMenu()
		}
		if input.Text() == "square" {
			drawSquare(10)
		}
		if input.Text() == "ffile" {
			fortuneFile()
		}
		if input.Text() == "prime" {
			getPrime()
		}

		fmt.Print("\n [+] Enter a system command (help) :")
	}

	fmt.Print("\n") // Drop a line on exit
}

func getPrime() {
	fmt.Print("\t [+] Finding all prime numbers between two values")

	fmt.Print("\n\t Enter a lower bounds :")
	var min int
	fmt.Scan(&min)

	fmt.Print("\n\t Enter an upper bounds for a number :")
	var max int
	fmt.Scan(&max)

	if min < 2 || max < 2 {
		fmt.Println("\t Both Numbers must be greater than 2")
		return
	}
	fmt.Print("\nPrimes")

	if max <= min {
		fmt.Print("\n\t [-] Error of bounds values")
		return
	}
	for min <= max {

		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(min))); i++ {
			if min%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Printf("\n%d ", min) //print out the primes on own line
		}
		min++
	}
	fmt.Println()
}

func fortuneFile() {
	fmt.Print("\t [+] Enter the filename to create :")
	var fname string
	fmt.Scan((&fname))
	var message = "You are special and unique just like everyone else in the herd."
	// make a file
	f, err := os.Create(fname + "_fortune.md")
	if err != nil {
		log.Fatal(err)
	}
	f.Write([]byte(message))
	defer f.Close()
	fmt.Println("\n\t [+] Saving file :", f.Name())

}

func getArguments(command string) {
	if command == "hello" {
		fmt.Print("\t Why, Hello to you too!\n")
	}

	if command == "details" {
		fmt.Print("\n\t [+] Enter Your First Name: ")
		var first string
		fmt.Scan(&first)

		fmt.Print("\n\t [+] Enter Last Name: ")
		var last string
		fmt.Scan(&last)
		fmt.Println(" Nice to meet you:", first, last, "\n")
	}

	if command == "pwd" {
		/// print out the current directory
		getPath()
	}

	if command == "ran" {
		getRandomNumber()
	}

}

func getRandomNumber() {
	/// ask user for min and max values; choose random numbers in this window
	// var min int = 0    // Debugging Min value for random numbers
	// var max int = 100  // Debugging Max value for random numbers
	// var nNums int = 10 // Debugging number of random numbers to produce

	fmt.Print("\t We will find n random numbers between lower a and upper bounds")
	fmt.Print("\n\t Enter a lower bounds :")
	var min int
	fmt.Scan(&min)

	fmt.Print("\n\t Enter an upper bounds for a number :")
	var max int
	fmt.Scan(&max)

	/////////////////////////////////////////////////////////////
	// The Intn() function of the rand package can be used to generate an integer in the interval of 0 and n. It takes only one argument, the n or the upper bound. It throws an error if the given argument is less than 0.

	// v := rand.Int()  // generates a random integer

	// We can improve it so that we can define lower and upper bounds and the function will generate random within that specified range. Here is how to do that.

	// v := rand.Intn(max-min) + min    // range is min to max

	// ref: https://golangdocs.com/generate-random-numbers-in-golang
	/////////////////////////////////////////////////////////////

	fmt.Print("\n\t Number of random numbers to create :")
	var nNums int
	fmt.Scan(&nNums)

	number_slice := []int{} // define a slice (i.e., a list)
	fmt.Print("\nData")

	for i := 0; i < nNums; i++ {
		randomNum := rand.Intn(max-min) + min
		// 	number_slice[i] = randomNum
		number_slice = append(number_slice, randomNum)
		fmt.Print("\n", randomNum)
	}
	// check the randomness of the numbers
	for i := range number_slice{
		// fmt.Print("\nTesting: ",number_slice[i])
		findFrequency(number_slice, number_slice[i])
	}

}


func findFrequency(arr []int, num int){
	count := 0
	// fmt.Print("length of array: ",len(arr))
	for _, item := range arr{
	   if item == num{
		  count++
	   }
	}
	var freq float64
	freq = (float64(count)) / (float64(len(arr)))
	fmt.Printf("\n\t    Frequency( %d ) = %f", num, freq)
 }


// Note: Frequency of number value analysis:
// ideas: https://zetcode.com/golang/word-frequency/ for finding frequencies of words

// ideas: https://www.tutorialspoint.com/write-a-golang-program-to-find-the-frequency-of-each-element-in-an-array#

func getPath() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	// return path
}

func drawSquare(count int) {
	/// draw a square in X's

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			fmt.Print(" X")
		}
		fmt.Print("\n")
	}
}
