package tea

import (
	"github.com/charmbracelet/x/ansi"
)

// setModifyOtherKeysMsg is a message to set XTerm modifyOtherKeys mode.
type setModifyOtherKeysMsg int

// EnableXtermModifyOtherKeys is a command to enable XTerm modifyOtherKeys mode.
//
// The mode can be on of the following:
//
//	1: Report ambiguous keys as escape codes
//	2: Report ambiguous keys as escape codes including modified keys like Alt-<key>
//	   and Meta-<key>
//
// See https://invisible-island.net/xterm/manpage/xterm.html#VT100-Widget-Resources:modifyOtherKeys
func EnableModifyOtherKeys(mode int) Cmd { //nolint:unused
	return func() Msg {
		return setModifyOtherKeysMsg(mode)
	}
}

// DisableModifyOtherKeys is a command to disable XTerm modifyOtherKeys mode.
func DisableModifyOtherKeys() Msg { //nolint:unused
	return setModifyOtherKeysMsg(0)
}

func parseXTermModifyOtherKeys(csi *ansi.CsiSequence) Msg {
	// XTerm modify other keys starts with ESC [ 27 ; <modifier> ; <code> ~
	mod := KeyMod(csi.Param(1) - 1)
	r := rune(csi.Param(2))

	switch r {
	case ansi.BS:
		return KeyPressMsg{Mod: mod, Code: KeyBackspace}
	case ansi.HT:
		return KeyPressMsg{Mod: mod, Code: KeyTab}
	case ansi.CR:
		return KeyPressMsg{Mod: mod, Code: KeyEnter}
	case ansi.ESC:
		return KeyPressMsg{Mod: mod, Code: KeyEscape}
	case ansi.DEL:
		return KeyPressMsg{Mod: mod, Code: KeyBackspace}
	}

	// CSI 27 ; <modifier> ; <code> ~ keys defined in XTerm modifyOtherKeys
	k := KeyPressMsg{Code: r, Mod: mod}
	if k.Mod <= ModShift {
		k.Text = string(r)
	}

	return k
}

// modifyOtherKeys is an internal message that queries the terminal for its
// modifyOtherKeys mode.
type modifyOtherKeys struct{}

// ModifyOtherKeys is a command that queries the terminal for its
// modifyOtherKeys mode.
func ModifyOtherKeys() Msg { //nolint:unused
	return modifyOtherKeys{}
}

// ModifyOtherKeysMsg is a message that represents XTerm modifyOtherKeys
// report. Querying the terminal for the modifyOtherKeys mode will return a
// ModifyOtherKeysMsg message with the current mode set.
//
//	0: disable
//	1: enable mode 1
//	2: enable mode 2
//
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h3-Functions-using-CSI-_-ordered-by-the-final-character_s_
// See: https://invisible-island.net/xterm/manpage/xterm.html#VT100-Widget-Resources:modifyOtherKeys
type ModifyOtherKeysMsg int

// TerminalVersionMsg is a message that represents the terminal version.
type TerminalVersionMsg string

// terminalVersion is an internal message that queries the terminal for its
// version using XTVERSION.
type terminalVersion struct{}

// TerminalVersion is a command that queries the terminal for its version using
// XTVERSION. Note that some terminals may not support this command.
func TerminalVersion() Msg {
	return terminalVersion{}
}