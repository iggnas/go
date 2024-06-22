package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func prompt() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nprompt: ")
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		str = strings.TrimSpace(str)
		if str == "" {
			fmt.Println("Input cannot be empty. Try 'h' for help.")
			continue
		}
		return str
	}
}

func ArrtoStrr(strarr []string) string {
    strc := strings.Join(strarr, " ")
    strc = strings.Trim(strc, "[]")
    return strc
}

func checkEmpty(arrs string) bool {
    if(arrs == "_" || arrs == " " ){
		return true
	} else{
		return false
	}
}

func changeturn(who bool) bool {
    return !who
}

func change(who bool, place string) (string, bool) {
	if checkEmpty(place) {
		if who {
			return "x", true
		} else {
			return "o", true
		}
	} else {
		fmt.Println("This place isn't empty")
		return place, false
	}
}

func winning(row1, row2, row3 []string) (bool, string) {
	if (row1[1] == row1[4] && row1[4] == row1[7] && row1[1] == "x") ||
	   (row2[1] == row2[4] && row2[4] == row2[7] && row2[1] == "x") ||
	   (row3[1] == row3[4] && row3[4] == row3[7] && row3[1] == "x") ||
	   (row1[1] == row2[1] && row2[1] == row3[1] && row1[1] == "x") ||
	   (row1[4] == row2[4] && row2[4] == row3[4] && row1[4] == "x") ||
	   (row1[7] == row2[7] && row2[7] == row3[7] && row1[7] == "x") ||
	   (row1[1] == row2[4] && row2[4] == row3[7] && row1[1] == "x") ||
	   (row1[7] == row2[4] && row2[4] == row3[1] && row1[7] == "x") {
		return true, "x"
	}else{
		if (row1[1] == row1[4] && row1[4] == row1[7] && row1[1] == "o") ||
	   	(row2[1] == row2[4] && row2[4] == row2[7] && row2[1] == "o") ||
	   	(row3[1] == row3[4] && row3[4] == row3[7] && row3[1] == "o") ||
	   	(row1[1] == row2[1] && row2[1] == row3[1] && row1[1] == "o") ||
	   	(row1[4] == row2[4] && row2[4] == row3[4] && row1[4] == "o") ||
	   	(row1[7] == row2[7] && row2[7] == row3[7] && row1[7] == "o") ||
	   	(row1[1] == row2[4] && row2[4] == row3[7] && row1[1] == "o") ||
	   	(row1[7] == row2[4] && row2[4] == row3[1] && row1[7] == "o"){
			return true, "o"
		}
	}
	return false, ""
}

func main() {
	var who bool = true
	row1 := []string{"_","_", "|","_","_", "|", "_","_"}
	row2 := []string{"_","_", "|","_","_", "|", "_","_"}
	row3 := []string{" "," ", "|"," "," ", "|", " "," "}
	fmt.Println("Tic tac toe game")
	fmt.Println("     A     B     C")
	fmt.Println("1  ","_","_", "|","_","_", "|", "_","_")
	fmt.Println("2  ","_","_", "|","_","_", "|", "_","_")
	fmt.Println("3  "," "," ", "|"," "," ", "|", " "," ")

	for{
        prompt := prompt()
        switch prompt {
        case "A1", "a1":
            value, changed := change(who, row1[1])
			if changed {
				row1[1] = value
				who = changeturn(who)
			}
			
		case "B1", "b1":
			value, changed := change(who, row1[4])
			if changed {
				row1[4] = value
				who = changeturn(who)
			}
		case "C1", "c1":
			value, changed := change(who, row1[7])
			if changed {
				row1[7] = value
				who = changeturn(who)
			}
		case "A2", "a2":
			value, changed := change(who, row2[1])
			if changed {
				row2[1] = value
				who = changeturn(who)
			}
		case "B2", "b2":
			value, changed := change(who, row2[4])
			if changed {
				row2[4] = value
				who = changeturn(who)
			}
		case "C2", "c2":
			value, changed := change(who, row2[7])
			if changed {
				row2[7] = value
				who = changeturn(who)
			}
		case "A3", "a3":
			value, changed := change(who, row3[1])
			if changed {
				row3[1] = value
				who = changeturn(who)
			}
		case "B3", "b3":
			value, changed := change(who, row3[4])
			if changed {
				row3[4] = value
				who = changeturn(who)
			}
		case "C3", "c3":
			value, changed := change(who, row3[7])
			if changed {
				row3[7] = value
				who = changeturn(who)
			}
		case "q", "Q":
			os.Exit(0)
		case "h", "H", "help":
			fmt.Print("\nQ to quit.")
		default:
			fmt.Print("Type h for help.")
		}
		fmt.Println("     A     B     C")
		srow1 := ArrtoStrr(row1)
		srow2 := ArrtoStrr(row2)
		srow3 := ArrtoStrr(row3)
		fmt.Println("1", srow1)
		fmt.Println("2", srow2)
		fmt.Println("3", srow3)
		won, whowon := winning(row1, row2, row3)
		if(won == true){
			fmt.Println("\n\n", whowon, " won congratulations!!")
			os.Exit(0)
		} else{
			continue
		}
	}

}
