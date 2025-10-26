package gn

import (
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

var (
	titleRe   = regexp.MustCompile(`<title>(.*?)</title>`)
	warningRe = regexp.MustCompile(`<warn>(.*?)</warn>`)
	emRe      = regexp.MustCompile(`<em>(.*?)</em>`)
	errRe     = regexp.MustCompile(`<err>(.*?)</err>`)
)

func Message(msg string, vars ...any) {
	print(unknownMsgType, msg, vars)
}

func Warn(msg string, vars ...any) {
	print(warningMsgType, msg, vars)
}

func Info(msg string, vars ...any) {
	print(infoMsgType, msg, vars)
}

func Progress(msg string, vars ...any) {
	print(progressMsgType, msg, vars)
}

func Success(msg string, vars ...any) {
	print(successMsgType, msg, vars)
}

func print(mt msgType, msg string, vars []any) {
	um := userMsg{
		msgType: mt,
		msg:     msg,
		vars:    vars,
	}
	um.print()
}

type msgType int

const (
	unknownMsgType msgType = iota
	errorMsgType
	warningMsgType
	infoMsgType
	progressMsgType
	successMsgType
)

// userMsg produces a user-friendly progress message for applications.
type userMsg struct {
	msgType
	msg  string
	vars []any
}

func (u userMsg) print() {
	msg := u.msg
	if len(u.vars) > 0 {
		msg = fmt.Sprintf(u.msg, u.vars)
	}
	msg = u.colorize(msg)

	icon := ""
	switch u.msgType {
	case infoMsgType:
		icon = "ℹ️"
	case successMsgType:
		icon = "✅"
	case warningMsgType:
		icon = "⚠️"
	case errorMsgType:
		icon = "❌"
	case progressMsgType:
		icon = "⏳"
	}

	fmt.Println(icon + msg)
}

func (u userMsg) colorize(msg string) string {
	msg = titleRe.ReplaceAllStringFunc(msg, func(match string) string {
		res := color.GreenString(titleRe.FindStringSubmatch(match)[1])
		return "**" + res + "**"
	})
	msg = warningRe.ReplaceAllStringFunc(msg, func(match string) string {
		return color.YelloString(warningRe.FindStringSubmatch(match)[1])
	})
	msg = emRe.ReplaceAllStringFunc(msg, func(match string) string {
		return color.GreenString(emRe.FindStringSubmatch(match)[1])
	})
	msg = errRe.ReplaceAllStringFunc(msg, func(match string) string {
		return color.RedString(warningRe.FindStringSubmatch(match)[1])
	})

	return msg
}
