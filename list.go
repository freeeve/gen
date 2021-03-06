package main

import (
	"fmt"

	"github.com/clipperhouse/gen/typewriter"
)

func list() error {
	imports := []string{
		`"fmt"`,
		`"os"`,
		`"github.com/clipperhouse/gen/typewriter"`,
	}

	return execute(listStandard, imports, listBody)
}

func listStandard() error {
	app, err := typewriter.NewApp("+gen")

	if err != nil {
		return err
	}

	fmt.Fprintln(out, "Installed typewriters:")
	for _, tw := range app.TypeWriters {
		fmt.Fprintf(out, "  %s\n", tw.Name())
	}

	return nil
}

const listBody string = `
func main() {
	app, err := typewriter.NewApp("+gen")

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	fmt.Println("Installed typewriters (custom):")
	for _, tw := range app.TypeWriters {
		fmt.Println("  " + tw.Name())
	}
}
`
