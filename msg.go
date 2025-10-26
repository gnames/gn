// Package gn provides user-friendly messaging and error handling utilities
// for Go applications with colorized output and formatted messages.
package gn

import (
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

// Regular expressions for parsing inline formatting tags in messages.
// Supported tags:
//   - <title>...</title> - Green text with ** emphasis
//   - <warn>...</warn>   - Yellow text for warnings
//   - <em>...</em>       - Green text for emphasis
//   - <err>...</err>     - Red text for errors
var (
	titleRe   = regexp.MustCompile(`<title>(.*?)</title>`)
	warningRe = regexp.MustCompile(`<warn>(.*?)</warn>`)
	emRe      = regexp.MustCompile(`<em>(.*?)</em>`)
	errRe     = regexp.MustCompile(`<err>(.*?)</err>`)
)

// Message prints a general message without a specific type indicator.
// The msg parameter can include format specifiers and inline formatting tags.
// Optional vars are used to format the message using fmt.Sprintf.
func Message(msg string, vars ...any) {
	print(unknownMsgType, msg, vars)
}

// Warn prints a warning message with a warning icon (⚠️).
// The msg parameter can include format specifiers and inline formatting tags.
// Optional vars are used to format the message using fmt.Sprintf.
func Warn(msg string, vars ...any) {
	print(warningMsgType, msg, vars)
}

// Info prints an informational message with an info icon (ℹ️).
// The msg parameter can include format specifiers and inline formatting tags.
// Optional vars are used to format the message using fmt.Sprintf.
func Info(msg string, vars ...any) {
	print(infoMsgType, msg, vars)
}

// Progress prints a progress message with a progress icon (⏳).
// The msg parameter can include format specifiers and inline formatting tags.
// Optional vars are used to format the message using fmt.Sprintf.
func Progress(msg string, vars ...any) {
	print(progressMsgType, msg, vars)
}

// Success prints a success message with a success icon (✅).
// The msg parameter can include format specifiers and inline formatting tags.
// Optional vars are used to format the message using fmt.Sprintf.
func Success(msg string, vars ...any) {
	print(successMsgType, msg, vars)
}

// print is an internal function that creates a userMsg and prints it.
// It handles the creation and formatting of messages based on the message type.
func print(mt msgType, msg string, vars []any) {
	um := userMsg{
		msgType: mt,
		msg:     msg,
		vars:    vars,
	}
	um.print()
}

// msgType represents the type of message being displayed, which determines
// the icon and styling applied to the message.
type msgType int

// Message type constants define the different types of messages supported.
// Each type is associated with a specific icon and formatting style.
const (
	unknownMsgType  msgType = iota // General message without a specific icon
	errorMsgType                   // Error message with ❌ icon
	warningMsgType                 // Warning message with ⚠️ icon
	infoMsgType                    // Informational message with ℹ️ icon
	progressMsgType                // Progress message with ⏳ icon
	successMsgType                 // Success message with ✅ icon
)

// userMsg produces a user-friendly progress message for applications.
type userMsg struct {
	msgType
	msg  string
	vars []any
}

// print formats and displays the user message with appropriate icon and colors.
// It applies fmt.Sprintf formatting if variables are provided, then applies
// color formatting based on inline tags, and finally prints with the appropriate icon.
func (u userMsg) print() {
	msg := u.msg
	if len(u.vars) > 0 {
		msg = fmt.Sprintf(u.msg, u.vars...)
	}
	msg = u.colorize(msg)

	icon := ""
	switch u.msgType {
	case infoMsgType:
		icon = "ℹ️ "
	case successMsgType:
		icon = "✅ "
	case warningMsgType:
		icon = "⚠️ "
	case errorMsgType:
		icon = "❌ "
	case progressMsgType:
		icon = "⏳ "
	}

	fmt.Println(icon + msg)
}

// colorize processes inline formatting tags in the message and applies
// terminal colors using the fatih/color package. It supports:
//   - <title>text</title> - Green text wrapped with ** emphasis
//   - <warn>text</warn>   - Yellow text for warnings
//   - <em>text</em>       - Green text for emphasis
//   - <err>text</err>     - Red text for errors
func (u userMsg) colorize(msg string) string {
	msg = titleRe.ReplaceAllStringFunc(msg, func(match string) string {
		res := color.GreenString(titleRe.FindStringSubmatch(match)[1])
		return "**" + res + "**"
	})
	msg = warningRe.ReplaceAllStringFunc(msg, func(match string) string {
		return color.YellowString(warningRe.FindStringSubmatch(match)[1])
	})
	msg = emRe.ReplaceAllStringFunc(msg, func(match string) string {
		return color.GreenString(emRe.FindStringSubmatch(match)[1])
	})
	msg = errRe.ReplaceAllStringFunc(msg, func(match string) string {
		return color.RedString(errRe.FindStringSubmatch(match)[1])
	})

	return msg
}
