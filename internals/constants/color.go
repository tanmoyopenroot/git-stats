package constants

import "fmt"

const (
	// GreyText ... Grey color for text
	GreyText = "\x1b[30;1m"
	// RedText ... Red color for text
	RedText = "\x1b[31;1m"
	// GreenText ... Green color for text
	GreenText = "\x1b[32;1m"
	// YellowText ... Yellow color for text
	YellowText = "\x1b[33;1m"
	// DeepBlueText ... Deep Blue color for text
	DeepBlueText = "\x1b[34;1m"
	// VioletText ... Violet color for text
	VioletText = "\x1b[35;1m"
	// LightBlueText ... Light Blue color for text
	LightBlueText = "\x1b[36;1m"
	// WhiteText ... White color for text
	WhiteText = "\x1b[37;1m"
	// EndText ...
	EndText = "\x1b[0m"

	// UncommittedColor ... Uncommitted cell background color
	UncommittedColor = "\033[1;37;47m"
	// Commited0To5Color ... Commited cell background color range 0 to 5
	Commited0To5Color = "\033[1;30;47m"
	// Commited5To10Color ... Commited cell background color range 5 to 10
	Commited5To10Color = "\033[1;30;43m"
	// CommitedMoreThan10Color ... Commited cell background color more than 10
	CommitedMoreThan10Color = "\033[1;30;42m"
	// CommitedMoreThan100Color ... Commited cell background color more than 100
	CommitedMoreThan100Color = "\033[1;30;42m"
	// TodaysCellColor ... Today's cell color
	TodaysCellColor = "\033[1;37;45m"
	// EndColor ...
	EndColor = "\033[0m"

	// BoundaryColor ... Set boundary color
	BoundaryColor = WhiteText + "%s" + EndText
)

// PrintAvailableColors ... Print available colors
func PrintAvailableColors() {
	fmt.Println(GreyText, "Grey Text", EndText)
	fmt.Println(RedText, "Red Text", EndText)
	fmt.Println(GreenText, "Green Text", EndText)
	fmt.Println(YellowText, "Yellow Text", EndText)
	fmt.Println(DeepBlueText, "DeepBlue Text", EndText)
	fmt.Println(VioletText, "Violet Text", EndText)
	fmt.Println(LightBlueText, "LightBlue Text", EndText)
	fmt.Println(WhiteText, "White Text ||| ----", EndText)

	fmt.Println("--------------")

	for i := 0; i <= 9; i++ {
		fmt.Printf("\033[1;30;4%dm Text %d \033[0m\n", i, i)
	}

	fmt.Println("--------------")

	for i := 0; i <= 9; i++ {
		fmt.Printf("\033[1;37;4%dm Text %d \033[0m\n", i, i)
	}
}
