package lib

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type Day interface {
	Part1() (string, error)
	Part2() (string, error)
}

func runPart1WithTimings(d Day) {
	start := time.Now()

	result, err := d.Part1()

	fmt.Println("======")
	fmt.Printf("Part 1 completed in %dms\n", time.Since(start).Milliseconds())

	if err != nil {
		fmt.Printf("Error:\n%v", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}
}

func runPart2WithTimings(d Day) {
	start := time.Now()

	result, err := d.Part2()

	fmt.Println("======")
	fmt.Printf("Part 2 completed in %dms\n", time.Since(start).Milliseconds())

	if err != nil {
		fmt.Printf("Error:\n%v", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}
}

func Run(day Day) {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "part1",
				Usage: "run part 1",
				Action: func(c *cli.Context) error {
					runPart1WithTimings(day)
					return nil
				},
			},
			{
				Name:  "part2",
				Usage: "run part 2",
				Action: func(c *cli.Context) error {
					runPart1WithTimings(day)
					return nil
				},
			},
			{
				Name:  "all",
				Usage: "run both parts",
				Action: func(c *cli.Context) error {
					runPart1WithTimings(day)
					fmt.Println()
					runPart2WithTimings(day)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
