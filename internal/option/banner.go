package option

import (
	"fmt"
	"os"

	"github.com/trhacknon/aurora"
)

func showBanner() {
	fmt.Fprintf(os.Stderr, "%s\n\n", aurora.Green(banner))
}
