package main

import (
	"fmt"
	"math"

	"github.com/mehmetserdar/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Sqrt(81))
	fmt.Println(strutil.Reverse("olleh"))
}
