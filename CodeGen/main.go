package main

import (
	"fmt"
	"os"
	"flag"
	"strconv"
)

var header string = "package main\n\nimport \"fmt\"\n\n"
var t string = "    "

func doubleNumString(num int) string {
	return fmt.Sprintf("fmt.Println(\"Your Number %d times 2 is %d\")\n", num, num * 2)
}

func createFunc(maxNum int, tab string) string {
	return createFuncRecurse("", 0, maxNum, tab)
}

func createFuncRecurse(out string, currNum int, maxNum int, tab string) string {
	if currNum >= maxNum {
		return fmt.Sprintf("%sfmt.Println(\"Your number is too big :(\")\n", tab)
	}
	recurse := createFuncRecurse(out, currNum + 1, maxNum, tab + t)
	message := doubleNumString(currNum)
	bs := fmt.Sprintf("%sif num == %d {\n%s    %s%s} else {\n%s%s}\n", tab, currNum, tab, message, tab, recurse, tab)
	return bs
}

func createMain(maxNum int) string {
	nested := createFunc(maxNum, "    ")
	mainstr := fmt.Sprintf("//this program can multiply numbers up to %d\nfunc main() {\n%svar num int\n%sfmt.Println(\"Input a number: \")\n%sfmt.Scanf(\"%%d\", &num)\n\n%s\n}", maxNum - 1, t, t, t, nested)
	return mainstr
}

func main() {
	dest := flag.String("dest", "./timesTwo.go", "Destination file path")
	flag.Parse()

	args := flag.Args()
	
	var maxNum int

	if len(args) == 0 {
		fmt.Println("Empty non-positional argument. Applying default config")
		maxNum = 5
	} else {
		temp, err := strconv.Atoi(args[0])
		if err != nil {
			panic("Please provide a positive integer only.")
		}
		maxNum = temp
	}

	if maxNum < 0 {
		panic("Please provide a positive integer only.")
	}

	output := []byte(header + createMain(maxNum))
  if err := os.WriteFile(*dest, output, 0644); err != nil {
		panic(err)
	}

	fmt.Printf("Successful write to %s", *dest)
}
