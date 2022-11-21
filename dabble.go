package main

// ref: https://www.youtube.com/watch?v=C8LgvuEBraI

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
	syscmds_map["prime"] = ":Create a random prime number"
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
		fmt.Println("\t Both Numbers must be greater than 2.")
		return
	}
	if max <= min {
		fmt.Print("\n\t [-] Error of bounds values. ")
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
			fmt.Printf("\n %d ", min) //print out the primes on own line
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
		// Taking input from user
		fmt.Print("\n\t [+] Enter Your First Name: ")
		// declare variable name and type, then assign user input to variable
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

	// return command
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

	for i := 0; i < nNums; i++ {
		randomNum := rand.Intn(max-min) + min
		// 	number_slice[i] = randomNum
		number_slice = append(number_slice, randomNum)
		fmt.Print("\n", randomNum)
	}
	// fmt.Print("\n Random numbers : ", number_slice)
	// numAnalysis(number_slice) TODO
}

func numAnalysis(number_slice []int) {
	// fmt.Print("\n\t [+] numAnalysis says list is : ", number_slice)
	// measure entropy
	numProbs_slice := make(map[int]int) // we record each value and its number of occurrences
	// print out numbers
	for i := range number_slice {
		// fmt.Print("\n", i, " -maps-> ", number_slice[i]) // debugging
		fmt.Print("\n", number_slice[i])
		numProbs_slice[i] = number_slice[i]
	}

}

// Note:
// check https://zetcode.com/golang/word-frequency/ for finding frequencies of words

// Another option to check;
// https://www.tutorialspoint.com/write-a-golang-program-to-find-the-frequency-of-each-element-in-an-array#

func getPath() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user
	// return path
}

func drawSquare(count int) {
	/// draw a square in x's

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			fmt.Print(" X")
		}
		fmt.Print("\n")
	}
}

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

// junk bin. Can any of this code be helpful in someway here?!
// x := 5
// y := 7
// sum := x + y
// fmt.Print("\t [+] ", x, " + ", y, " = ", sum, "\n")
// // fmt.Print(" sum: %T\n", sum)
// fmt.Print("\t [+] The type of sum is :", reflect.TypeOf(sum), "\n")
// create an object called a 'struct'
// type studentThings struct {
// 	typeOfMachine     string
// 	writingInstrument string
// 	favouriteSong     string
// 	favouriteColour   string
// 	favouriteNumber   int
// }

// // create a map which are similar to a dictionary as in python

// vertices := make(map[string]int)
// vertices["triangle"] = 3
// vertices["square"] = 4
// vertices["dodecagon"] = 12
// vertices["myShapeOfMadness"] = 1010101

// fmt.Print("\n\t [+] My vertices are :", vertices)

// //https://freshman.tech/snippets/go/iterating-over-slices/
// for i := range vertices {
// 	fmt.Print(i, "\n")
// }

// delete(vertices, "myShapeOfMadness")
// fmt.Print("\n\t [+] Item Removed.  My vertices are now :", vertices)

// firstName, lastName := "Oliver", "Bonham-Carter"
// fmt.Print("\n\t [+] my first and last name is : ", firstName, " ", lastName, "\n")
// student := studentThings{
// 	typeOfMachine:     "Mac",
// 	writingInstrument: "pencil",
// 	favouriteSong:     "Born to be bad",
// 	favouriteColour:   "green",
// 	favouriteNumber:   13}
// fmt.Print("\t [+] A print out of the object student: ", student)

// fmt.Print("\n\t [+] Favourite colour of student is: ", student.favouriteColour)
