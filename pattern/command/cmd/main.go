package main

import "L2/pattern/command/pkg"

func main() {
	tv := &pkg.Tv{}

	onCommand := &pkg.OnCommand{
		Device: tv,
	}

	offCommand := &pkg.OffCommand{
		Device: tv,
	}

	onButton := &pkg.Button{
		Command: onCommand,
	}

	onButton.Press()

	offButton := &pkg.Button{
		Command: offCommand,
	}
	offButton.Press()
}
