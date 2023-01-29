// Package logger provides methods for logging various events in the application.
package utils

import (
	"GigaCat/ProxD/proxy"
	"fmt"
	"strings"

	"github.com/jwalton/gchalk"
)

// Constants representing the different log types.
const (
	Working = "WORKING"
	Info    = "INFO"
	Error   = "ERROR"
)

// Variables to store the log type colors.
var (
	workingColor = gchalk.Ansi256(46)(Working)
	infoColor    = gchalk.WithAnsi256(146).Bold(Info)
	errorColor   = gchalk.WithAnsi256(196).Bold(Error)
	lb           = gchalk.WithBold().Ansi256(241)("[")
	rb           = gchalk.WithBold().Ansi256(241)("]")
	proxyPrefix  = gchalk.WithAnsi256(171).Bold("PROXY ")
)

// LogProxy logs information about a proxy.
func LogProxy(proxy *proxy.Proxy) {
	str := new(strings.Builder)
	str.WriteString(encloseInBrackets(workingColor))
	str.WriteString(" ")
	str.WriteString(formatProxyString(proxy))
	fmt.Println(str.String())
}

// LogInfo logs an info message.
func LogInfo(info string) {
	// Add logic to handle the info message here
}

// LogErr logs an error message.
func LogErr(err error) {
	if err != nil {
		str := new(strings.Builder)
		str.WriteString(encloseInBrackets(errorColor))
		str.WriteString(" ")
		str.WriteString(err.Error())
		fmt.Println(str.String())
	}
}

// encloseInBrackets returns a string enclosed in brackets.
func encloseInBrackets(str string) string {
	return lb + str + rb
}

// formatProxyString returns a string representation of a proxy.
func formatProxyString(proxy *proxy.Proxy) string {
	str := new(strings.Builder)
	str.WriteString(proxyPrefix)
	str.WriteString(gchalk.WithWhite().Bold(" -> "))
	str.WriteString(gchalk.WithAnsi256(206).Bold(" { "))
	str.WriteString(gchalk.Ansi256(195)(fmt.Sprintf("%s:%d", proxy.IP, proxy.Port)))
	str.WriteString(gchalk.WithWhite().Bold(" , "))
	str.WriteString(gchalk.Ansi256(123)(fmt.Sprintf("%s", proxy.Protocol)))
	str.WriteString(gchalk.WithAnsi256(206).Bold(" } "))
	return str.String()
}
