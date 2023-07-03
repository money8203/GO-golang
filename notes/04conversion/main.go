package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome pizza lovers!")
	fmt.Println("Please rate our services between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating: ", input)

	// But what if we want to add +1 to whatever rating our customer gives to us
	//We can't do that normally because input is a string and we can't add 1 to a string, So we need to convert the string to an integer

	//We can do that by using strconv package, and fow now we will use ParseFloat (convert string into a floating point number)
	// ParseFloat requires 2 values to be pass on i.e, 1st the string to be pass on and 2nd the bit size of it
	// bit size is the number of bits used to represent the value. 32 for float32 and 64 for float64 but usually we use float64
	//The first value is the converted value and the second value is the error, So we need to store both the values in two different variables

	// strconv.ParseFloat(strings.TrimSpace(input), 64) usage
	// The strings.TrimSpace() function removes all leading and trailing white space from a string like (space, tab, no-break space, etc.)
	// and all the line terminator characters (LF, CR, etc.).

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added +1 to the ratings: ", numRating+1)
	}
}
