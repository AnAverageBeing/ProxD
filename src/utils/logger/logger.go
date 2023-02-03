package logger

/*
*
*   Author       ->  AnAverageBeing
*   License      ->  GNU
*   Discord      ->  GigaCat#1521
*   Description  ->  Blazingly fast proxy utility tool
*
 */

import (
	"GigaCat/ProxD/proxy"
	"fmt"
	"strings"

	"github.com/jwalton/gchalk"
)

const (
	Working = "WORKING"
	Info    = "INFO"
	Error   = "ERROR"
)

var (
	workingColor = gchalk.Ansi256(46)(Working)
	infoColor    = gchalk.WithAnsi256(146).Bold(Info)
	errorColor   = gchalk.WithAnsi256(196).Bold(Error)
	lb           = gchalk.WithBold().Ansi256(241)("[")
	rb           = gchalk.WithBold().Ansi256(241)("]")
	proxyPrefix  = gchalk.WithAnsi256(171).Bold("PROXY ")
)

func LogProxy(proxy *proxy.Proxy, tries int) {
	str := new(strings.Builder)
	str.WriteString("\r")
	str.WriteString(encloseInBrackets(workingColor))
	str.WriteString(" ")
	str.WriteString(formatProxyString(proxy, tries))
	fmt.Println(str.String())
}

func LogInfo(info string) {
	str := new(strings.Builder)
	fmt.Print(encloseInBrackets(infoColor))
	str.WriteString(" ")
	str.WriteString(info)
	fmt.Println(str.String())
}

func LogErr(err error) {
	if err != nil {
		str := new(strings.Builder)
		str.WriteString(encloseInBrackets(errorColor))
		str.WriteString(" ")
		str.WriteString(err.Error())
		fmt.Println(str.String())
	}
}

func encloseInBrackets(str string) string {
	return lb + str + rb
}

func formatProxyString(proxy *proxy.Proxy, tries int) string {
	str := new(strings.Builder)
	str.WriteString(proxyPrefix)
	str.WriteString(gchalk.WithWhite().Bold(" -> "))
	str.WriteString(encloseInBrackets(fmt.Sprintf("%d", tries)))
	str.WriteString(gchalk.WithAnsi256(206).Bold(" { "))
	str.WriteString(gchalk.Ansi256(195)(fmt.Sprintf("%s:%d", proxy.IP, proxy.Port)))
	str.WriteString(gchalk.WithWhite().Bold(" , "))
	str.WriteString(gchalk.Ansi256(123)(fmt.Sprintf("%s", proxy.Protocol)))
	str.WriteString(gchalk.WithAnsi256(206).Bold(" } "))
	return str.String()
}
