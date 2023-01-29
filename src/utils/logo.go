package utils

import (
	"fmt"

	"github.com/jwalton/gchalk"
)

func PrintLogo() {
	fmt.Println(gchalk.WithHex("#9066ea").Bold("\n\n+-----------------------------------------------------+\n|  8888888b.                           8888888b.      |\n|  888   Y88b                          888  \"Y88b     |\n|  888    888                          888    888     |\n|  888   d88P 888d888 .d88b.  888  888 888    888     |\n|  8888888P\"  888P\"  d88\"\"88b `Y8bd8P' 888    888     |\n|  888        888    888  888   X88K   888    888     |\n|  888        888    Y88..88P .d8\"\"8b. 888  .d88P     |\n|  888        888     \"Y88P\"  888  888 8888888P\"      |\n|                                                     |\n|                                                     |\n|  " +
		gchalk.WithAnsi256(47).Bold("\"A Blazingly Fast Proxy Utility Tool\"              ") + "|\n|                                                     |\n|  " +
		gchalk.WithAnsi256(51).Bold("INFO :-") +
		"                                            |\n|       " +
		gchalk.WithAnsi256(207).Bold("Discord -> GigaCat#1521") +
		"                       |\n|       " +
		gchalk.WithAnsi256(207).Bold("Invite  -> dsc.gg/TEAMXD") +
		"                      |\n|       " +
		gchalk.WithAnsi256(207).Bold("GitHub  -> https://github.com/AnAverageBeing") +
		"  |\n|                                                     |\n+-----------------------------------------------------+\n\n"))
}
