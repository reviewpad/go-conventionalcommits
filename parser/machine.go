package parser

import (
	"bytes"
	"fmt"

	"github.com/reviewpad/go-conventionalcommits"
	"github.com/sirupsen/logrus"
)

// ColumnPositionTemplate is the template used to communicate the column where errors occur.
var ColumnPositionTemplate = ": col=%02d"

const (
	// ErrType represents an error in the type part of the commit message.
	ErrType = "illegal '%s' character in commit message type"
	// ErrTypeIncomplete represents an error when the type part of the commit message is not complete.
	ErrTypeIncomplete = "incomplete commit message type after '%s' character"
	// ErrColon is the error message that communicate that the mandatory colon after the type part of the commit message is missing.
	ErrColon = "expecting colon (':') character, got '%s' character"
	// ErrScope represents an error about illegal characters into the the scope part of the commit message.
	ErrScope = "illegal '%s' character in scope"
	// ErrScopeIncomplete represents a specific early-exit error.
	ErrScopeIncomplete = "expecting closing parentheses (')') character, got early exit after '%s' character"
	// ErrEmpty represents an error when the input is empty.
	ErrEmpty = "empty input"
	// ErrEarly represents an error when the input makes the machine exit too early.
	ErrEarly = "early exit after '%s' character"
	// ErrDescriptionInit tells the user that before of the description part a whitespace is mandatory.
	ErrDescriptionInit = "expecting at least one white-space (' ') character, got '%s' character"
	// ErrDescription tells the user that after the whitespace is mandatory a description.
	ErrDescription = "expecting a description text (without newlines) after '%s' character"
	// ErrNewline communicates an illegal newline to the user.
	ErrNewline = "illegal newline"
	// ErrMissingBlankLineAtBeginning tells the user that the a blank line is missing after the description or after the body.
	ErrMissingBlankLineAtBeginning = "missing a blank line"
	// ErrTrailer represents an error due to an unexepected character while parsing a footer trailer.
	ErrTrailer = "illegal '%s' character in trailer"
	// ErrTrailerIncomplete represent an error when a trailer is not complete.
	ErrTrailerIncomplete = "incomplete footer trailer after '%s' character"
)

const start int = 1
const firstFinal int = 85

const enTrailerBeg int = 87
const enTrailerEnd int = 33
const enBody int = 34
const enMain int = 1
const enConventionalTypesMain int = 35
const enFreeFormTypesMain int = 76

type machine struct {
	data             []byte
	cs               int
	p, pe, eof       int
	pb               int
	err              error
	bestEffort       bool
	typeConfig       conventionalcommits.TypeConfig
	logger           *logrus.Logger
	currentFooterKey string
	countNewlines    int
	lastNewline      int
}

func (m *machine) text() []byte {
	return m.data[m.pb:m.p]
}

func (m *machine) emitInfo(s string, args ...interface{}) {
	if m.logger != nil {
		logEntry := logrus.NewEntry(m.logger)
		for i := 0; i < len(args); i = i + 2 {
			logEntry = m.logger.WithField(args[0].(string), args[1])
		}
		logEntry.Infoln(s)
	}
}

func (m *machine) emitDebug(s string, args ...interface{}) {
	if m.logger != nil {
		logEntry := logrus.NewEntry(m.logger)
		for i := 0; i < len(args); i = i + 2 {
			logEntry = m.logger.WithField(args[0].(string), args[1])
		}
		logEntry.Debugln(s)
	}
}

func (m *machine) emitError(s string, args ...interface{}) error {
	e := fmt.Errorf(s+ColumnPositionTemplate, args...)
	if m.logger != nil {
		m.logger.Errorln(e)
	}
	return e
}

func (m *machine) emitErrorWithoutCharacter(messageTemplate string) error {
	return m.emitError(messageTemplate, m.p)
}

func (m *machine) emitErrorOnCurrentCharacter(messageTemplate string) error {
	return m.emitError(messageTemplate, string(m.data[m.p]), m.p)
}

func (m *machine) emitErrorOnPreviousCharacter(messageTemplate string) error {
	return m.emitError(messageTemplate, string(m.data[m.p-1]), m.p)
}

// NewMachine creates a new FSM able to parse Conventional Commits.
func NewMachine(options ...conventionalcommits.MachineOption) conventionalcommits.Machine {
	m := &machine{}

	for _, opt := range options {
		opt(m)
	}

	return m
}

// Parse parses the input byte array as a Conventional Commit message with no body neither footer.
//
// When a valid Conventional Commit message is given it outputs its structured representation.
// If the parsing detects an error it returns it with the position where the error occurred.
//
// It can also partially parse input messages returning a partially valid structured representation
// and the error that stopped the parsing.
func (m *machine) Parse(input []byte) (conventionalcommits.Message, error) {
	m.data = input
	m.p = 0
	m.pb = 0
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil
	m.currentFooterKey = ""
	m.countNewlines = 0
	output := &conventionalCommit{}
	output.footers = make(map[string][]string)

	switch m.typeConfig {
	case conventionalcommits.TypesFreeForm:
		m.cs = enFreeFormTypesMain
		break
	case conventionalcommits.TypesConventional:
		m.cs = enConventionalTypesMain
		break
	case conventionalcommits.TypesMinimal:
		fallthrough
	default:

		{
			m.cs = start
		}

		break
	}

	{
		var _widec int16
		if (m.p) == (m.pe) {
			goto _testEof
		}
		switch m.cs {
		case 1:
			goto stCase1
		case 0:
			goto stCase0
		case 2:
			goto stCase2
		case 3:
			goto stCase3
		case 4:
			goto stCase4
		case 5:
			goto stCase5
		case 6:
			goto stCase6
		case 7:
			goto stCase7
		case 8:
			goto stCase8
		case 85:
			goto stCase85
		case 9:
			goto stCase9
		case 86:
			goto stCase86
		case 10:
			goto stCase10
		case 11:
			goto stCase11
		case 12:
			goto stCase12
		case 13:
			goto stCase13
		case 33:
			goto stCase33
		case 90:
			goto stCase90
		case 91:
			goto stCase91
		case 34:
			goto stCase34
		case 92:
			goto stCase92
		case 35:
			goto stCase35
		case 36:
			goto stCase36
		case 37:
			goto stCase37
		case 38:
			goto stCase38
		case 39:
			goto stCase39
		case 40:
			goto stCase40
		case 41:
			goto stCase41
		case 42:
			goto stCase42
		case 43:
			goto stCase43
		case 93:
			goto stCase93
		case 44:
			goto stCase44
		case 94:
			goto stCase94
		case 45:
			goto stCase45
		case 46:
			goto stCase46
		case 47:
			goto stCase47
		case 48:
			goto stCase48
		case 49:
			goto stCase49
		case 50:
			goto stCase50
		case 51:
			goto stCase51
		case 52:
			goto stCase52
		case 53:
			goto stCase53
		case 54:
			goto stCase54
		case 55:
			goto stCase55
		case 56:
			goto stCase56
		case 57:
			goto stCase57
		case 58:
			goto stCase58
		case 59:
			goto stCase59
		case 60:
			goto stCase60
		case 61:
			goto stCase61
		case 62:
			goto stCase62
		case 63:
			goto stCase63
		case 64:
			goto stCase64
		case 65:
			goto stCase65
		case 66:
			goto stCase66
		case 67:
			goto stCase67
		case 68:
			goto stCase68
		case 69:
			goto stCase69
		case 70:
			goto stCase70
		case 71:
			goto stCase71
		case 72:
			goto stCase72
		case 73:
			goto stCase73
		case 74:
			goto stCase74
		case 75:
			goto stCase75
		case 76:
			goto stCase76
		case 77:
			goto stCase77
		case 78:
			goto stCase78
		case 79:
			goto stCase79
		case 80:
			goto stCase80
		case 95:
			goto stCase95
		case 81:
			goto stCase81
		case 96:
			goto stCase96
		case 82:
			goto stCase82
		case 83:
			goto stCase83
		case 84:
			goto stCase84
		case 87:
			goto stCase87
		case 14:
			goto stCase14
		case 15:
			goto stCase15
		case 88:
			goto stCase88
		case 16:
			goto stCase16
		case 17:
			goto stCase17
		case 89:
			goto stCase89
		case 18:
			goto stCase18
		case 19:
			goto stCase19
		case 20:
			goto stCase20
		case 21:
			goto stCase21
		case 22:
			goto stCase22
		case 23:
			goto stCase23
		case 24:
			goto stCase24
		case 25:
			goto stCase25
		case 26:
			goto stCase26
		case 27:
			goto stCase27
		case 28:
			goto stCase28
		case 29:
			goto stCase29
		case 30:
			goto stCase30
		case 31:
			goto stCase31
		case 32:
			goto stCase32
		}
		goto stOut
	stCase1:
		switch (m.data)[(m.p)] {
		case 70:
			goto tr1
		case 102:
			goto tr1
		}
		goto tr0
	tr0:

		if m.pe > 0 {
			if m.p != m.pe {
				m.err = m.emitErrorOnCurrentCharacter(ErrType)
			} else {
				// assert(m.p == m.pe)
				m.err = m.emitErrorOnPreviousCharacter(ErrTypeIncomplete)
			}
		}

		goto st0
	tr6:

		if m.err == nil {
			m.err = m.emitErrorOnCurrentCharacter(ErrColon)
		}

		goto st0
	tr10:

		if m.err == nil {
			m.err = m.emitErrorOnCurrentCharacter(ErrDescriptionInit)
		}

		goto st0
	tr13:

		if m.p < m.pe && m.data[m.p] == 10 {
			m.err = m.emitError(ErrNewline, m.p+1)
		} else {
			// assert(m.p == m.pe)
			m.err = m.emitErrorOnPreviousCharacter(ErrDescription)
		}

		goto st0
	tr14:

		m.err = m.emitErrorWithoutCharacter(ErrMissingBlankLineAtBeginning)

		goto st0
	tr16:

		if m.p < m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrScope)
		}

		goto st0
	tr21:

		if len(output.footers) == 0 {
			// Backtrack to the last marker
			// Ie., the text possibly a trailer token that is instead part of the body content
			if m.countNewlines > 0 {
				// In case new lines met while rewinding
				// advance the last marker by the number of the newlines so that they don't get parsed again
				// (they be added in the result by the body content appender)
				m.pb = m.lastNewline + 1
			}
			(m.p) = (m.pb) - 1

			m.emitDebug("try to parse body content", "pos", m.p)
			{
				goto st34
			}
		} else {
			// A rewind happens when an error while parsing a footer trailer is encountered
			// If this is not the first footer trailer the parser can't go back to parse body content again
			// Thus, emit an error
			if m.p != m.pe {
				m.err = m.emitErrorOnCurrentCharacter(ErrTrailer)
			} else {
				// assert(m.p == m.pe)
				m.err = m.emitErrorOnPreviousCharacter(ErrTrailerIncomplete)
			}
		}

		goto st0
	tr44:

		// Append newlines
		for m.countNewlines > 0 {
			output.body += "\n"
			m.countNewlines--
			m.emitInfo("valid commit message body content", "body", "\n")
		}
		// Append body content
		output.body += string(m.text())
		m.emitInfo("valid commit message body content", "body", string(m.text()))

		m.emitDebug("try to parse a footer trailer token", "pos", m.p)
		{
			goto st87
		}

		goto st0
	tr109:

		// Append newlines
		for m.countNewlines > 0 {
			output.body += "\n"
			m.countNewlines--
			m.emitInfo("valid commit message body content", "body", "\n")
		}
		// Append body content
		output.body += string(m.text())
		m.emitInfo("valid commit message body content", "body", string(m.text()))

		// Append content to body
		m.pb++
		m.p++
		output.body += string(m.text())
		m.emitInfo("valid commit message body content", "body", string(m.text()))
		// Do not advance over the current char
		(m.p)--

		m.emitDebug("try to parse a footer trailer token", "pos", m.p)
		{
			goto st87
		}

		goto st0
	stCase0:
	st0:
		m.cs = 0
		goto _out
	tr1:

		m.pb = m.p

		goto st2
	st2:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof2
		}
	stCase2:
		switch (m.data)[(m.p)] {
		case 69:
			goto st3
		case 73:
			goto st13
		case 101:
			goto st3
		case 105:
			goto st13
		}
		goto tr0
	st3:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof3
		}
	stCase3:
		switch (m.data)[(m.p)] {
		case 65:
			goto st4
		case 97:
			goto st4
		}
		goto tr0
	st4:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof4
		}
	stCase4:
		switch (m.data)[(m.p)] {
		case 84:
			goto st5
		case 116:
			goto st5
		}
		goto tr0
	st5:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof5
		}
	stCase5:

		output._type = string(m.text())
		m.emitInfo("valid commit message type", "type", output._type)

		switch (m.data)[(m.p)] {
		case 33:
			goto tr7
		case 40:
			goto st10
		case 58:
			goto st7
		}
		goto tr6
	tr7:

		output.exclamation = true
		m.emitInfo("commit message communicates a breaking change")

		goto st6
	st6:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof6
		}
	stCase6:
		if (m.data)[(m.p)] == 58 {
			goto st7
		}
		goto tr6
	st7:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof7
		}
	stCase7:
		if (m.data)[(m.p)] == 32 {
			goto st8
		}
		goto tr10
	st8:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof8
		}
	stCase8:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr13
		case 32:
			goto st8
		}
		goto tr12
	tr12:

		m.pb = m.p

		goto st85
	st85:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof85
		}
	stCase85:
		if (m.data)[(m.p)] == 10 {
			goto tr102
		}
		goto st85
	tr102:

		output.descr = string(m.text())
		m.emitInfo("valid commit message description", "description", output.descr)

		goto st9
	st9:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof9
		}
	stCase9:
		if (m.data)[(m.p)] == 10 {
			goto tr15
		}
		goto tr14
	tr15:

		m.emitDebug("found a blank line", "pos", m.p)

		m.emitDebug("try to parse a footer trailer token", "pos", m.p)
		{
			goto st87
		}

		goto st86
	st86:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof86
		}
	stCase86:
		goto st0
	st10:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof10
		}
	stCase10:
		if (m.data)[(m.p)] == 41 {
			goto tr18
		}
		switch {
		case (m.data)[(m.p)] > 39:
			if 42 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 126 {
				goto tr17
			}
		case (m.data)[(m.p)] >= 32:
			goto tr17
		}
		goto tr16
	tr17:

		m.pb = m.p

		goto st11
	st11:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof11
		}
	stCase11:
		if (m.data)[(m.p)] == 41 {
			goto tr20
		}
		switch {
		case (m.data)[(m.p)] > 39:
			if 42 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 126 {
				goto st11
			}
		case (m.data)[(m.p)] >= 32:
			goto st11
		}
		goto tr16
	tr18:

		m.pb = m.p

		output.scope = string(m.text())
		m.emitInfo("valid commit message scope", "scope", output.scope)

		goto st12
	tr20:

		output.scope = string(m.text())
		m.emitInfo("valid commit message scope", "scope", output.scope)

		goto st12
	st12:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof12
		}
	stCase12:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr7
		case 58:
			goto st7
		}
		goto tr6
	st13:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof13
		}
	stCase13:
		switch (m.data)[(m.p)] {
		case 88:
			goto st5
		case 120:
			goto st5
		}
		goto tr0
	st33:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof33
		}
	stCase33:
		if 32 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 126 {
			goto tr42
		}
		goto st0
	tr42:

		m.pb = m.p

		goto st90
	st90:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof90
		}
	stCase90:
		if (m.data)[(m.p)] == 10 {
			goto tr106
		}
		if 32 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 126 {
			goto st90
		}
		goto st0
	tr106:

		output.footers[m.currentFooterKey] = append(output.footers[m.currentFooterKey], string(m.text()))
		m.emitInfo("valid commit message footer trailer", m.currentFooterKey, string(m.text()))

		// Increment number of newlines to use in case we're still in the body
		m.countNewlines++
		m.lastNewline = m.p
		m.emitDebug("found a newline", "pos", m.p)

		m.emitDebug("try to parse a footer trailer token", "pos", m.p)
		{
			goto st87
		}

		goto st91
	tr108:

		// Increment number of newlines to use in case we're still in the body
		m.countNewlines++
		m.lastNewline = m.p
		m.emitDebug("found a newline", "pos", m.p)

		m.emitDebug("try to parse a footer trailer token", "pos", m.p)
		{
			goto st87
		}

		goto st91
	st91:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof91
		}
	stCase91:
		if (m.data)[(m.p)] == 10 {
			goto tr108
		}
		goto st0
	st34:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof34
		}
	stCase34:
		_widec = int16((m.data)[(m.p)])
		_widec = 256 + (int16((m.data)[(m.p)]) - 0)
		if m.p+2 < m.pe && m.data[m.p+1] == 10 && m.data[m.p+2] == 10 {
			_widec += 256
		}
		if 256 <= _widec && _widec <= 511 {
			goto tr45
		}
		goto tr44
	tr45:

		m.pb = m.p

		goto st92
	tr110:

		// Append newlines
		for m.countNewlines > 0 {
			output.body += "\n"
			m.countNewlines--
			m.emitInfo("valid commit message body content", "body", "\n")
		}
		// Append body content
		output.body += string(m.text())
		m.emitInfo("valid commit message body content", "body", string(m.text()))

		m.pb = m.p

		goto st92
	st92:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof92
		}
	stCase92:
		_widec = int16((m.data)[(m.p)])
		_widec = 256 + (int16((m.data)[(m.p)]) - 0)
		if m.p+2 < m.pe && m.data[m.p+1] == 10 && m.data[m.p+2] == 10 {
			_widec += 256
		}
		if 256 <= _widec && _widec <= 511 {
			goto tr110
		}
		goto tr109
	stCase35:
		switch (m.data)[(m.p)] {
		case 66:
			goto tr46
		case 67:
			goto tr47
		case 68:
			goto tr48
		case 70:
			goto tr49
		case 80:
			goto tr50
		case 82:
			goto tr51
		case 83:
			goto tr52
		case 84:
			goto tr53
		case 98:
			goto tr46
		case 99:
			goto tr47
		case 100:
			goto tr48
		case 102:
			goto tr49
		case 112:
			goto tr50
		case 114:
			goto tr51
		case 115:
			goto tr52
		case 116:
			goto tr53
		}
		goto tr0
	tr46:

		m.pb = m.p

		goto st36
	st36:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof36
		}
	stCase36:
		switch (m.data)[(m.p)] {
		case 85:
			goto st37
		case 117:
			goto st37
		}
		goto tr0
	st37:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof37
		}
	stCase37:
		switch (m.data)[(m.p)] {
		case 73:
			goto st38
		case 105:
			goto st38
		}
		goto tr0
	st38:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof38
		}
	stCase38:
		switch (m.data)[(m.p)] {
		case 76:
			goto st39
		case 108:
			goto st39
		}
		goto tr0
	st39:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof39
		}
	stCase39:
		switch (m.data)[(m.p)] {
		case 68:
			goto st40
		case 100:
			goto st40
		}
		goto tr0
	st40:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof40
		}
	stCase40:

		output._type = string(m.text())
		m.emitInfo("valid commit message type", "type", output._type)

		switch (m.data)[(m.p)] {
		case 33:
			goto tr58
		case 40:
			goto st45
		case 58:
			goto st42
		}
		goto tr6
	tr58:

		output.exclamation = true
		m.emitInfo("commit message communicates a breaking change")

		goto st41
	st41:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof41
		}
	stCase41:
		if (m.data)[(m.p)] == 58 {
			goto st42
		}
		goto tr6
	st42:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof42
		}
	stCase42:
		if (m.data)[(m.p)] == 32 {
			goto st43
		}
		goto tr10
	st43:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof43
		}
	stCase43:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr13
		case 32:
			goto st43
		}
		goto tr62
	tr62:

		m.pb = m.p

		goto st93
	st93:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof93
		}
	stCase93:
		if (m.data)[(m.p)] == 10 {
			goto tr112
		}
		goto st93
	tr112:

		output.descr = string(m.text())
		m.emitInfo("valid commit message description", "description", output.descr)

		goto st44
	st44:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof44
		}
	stCase44:
		if (m.data)[(m.p)] == 10 {
			goto tr63
		}
		goto tr14
	tr63:

		m.emitDebug("found a blank line", "pos", m.p)

		m.emitDebug("try to parse a footer trailer token", "pos", m.p)
		{
			goto st87
		}

		goto st94
	st94:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof94
		}
	stCase94:
		goto st0
	st45:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof45
		}
	stCase45:
		if (m.data)[(m.p)] == 41 {
			goto tr65
		}
		switch {
		case (m.data)[(m.p)] > 39:
			if 42 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 126 {
				goto tr64
			}
		case (m.data)[(m.p)] >= 32:
			goto tr64
		}
		goto tr16
	tr64:

		m.pb = m.p

		goto st46
	st46:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof46
		}
	stCase46:
		if (m.data)[(m.p)] == 41 {
			goto tr67
		}
		switch {
		case (m.data)[(m.p)] > 39:
			if 42 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 126 {
				goto st46
			}
		case (m.data)[(m.p)] >= 32:
			goto st46
		}
		goto tr16
	tr65:

		m.pb = m.p

		output.scope = string(m.text())
		m.emitInfo("valid commit message scope", "scope", output.scope)

		goto st47
	tr67:

		output.scope = string(m.text())
		m.emitInfo("valid commit message scope", "scope", output.scope)

		goto st47
	st47:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof47
		}
	stCase47:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr58
		case 58:
			goto st42
		}
		goto tr6
	tr47:

		m.pb = m.p

		goto st48
	st48:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof48
		}
	stCase48:
		switch (m.data)[(m.p)] {
		case 72:
			goto st49
		case 73:
			goto st40
		case 104:
			goto st49
		case 105:
			goto st40
		}
		goto tr0
	st49:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof49
		}
	stCase49:
		switch (m.data)[(m.p)] {
		case 79:
			goto st50
		case 111:
			goto st50
		}
		goto tr0
	st50:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof50
		}
	stCase50:
		switch (m.data)[(m.p)] {
		case 82:
			goto st51
		case 114:
			goto st51
		}
		goto tr0
	st51:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof51
		}
	stCase51:
		switch (m.data)[(m.p)] {
		case 69:
			goto st40
		case 101:
			goto st40
		}
		goto tr0
	tr48:

		m.pb = m.p

		goto st52
	st52:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof52
		}
	stCase52:
		switch (m.data)[(m.p)] {
		case 79:
			goto st53
		case 111:
			goto st53
		}
		goto tr0
	st53:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof53
		}
	stCase53:
		switch (m.data)[(m.p)] {
		case 67:
			goto st54
		case 99:
			goto st54
		}
		goto tr0
	st54:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof54
		}
	stCase54:
		switch (m.data)[(m.p)] {
		case 83:
			goto st40
		case 115:
			goto st40
		}
		goto tr0
	tr49:

		m.pb = m.p

		goto st55
	st55:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof55
		}
	stCase55:
		switch (m.data)[(m.p)] {
		case 69:
			goto st56
		case 73:
			goto st58
		case 101:
			goto st56
		case 105:
			goto st58
		}
		goto tr0
	st56:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof56
		}
	stCase56:
		switch (m.data)[(m.p)] {
		case 65:
			goto st57
		case 97:
			goto st57
		}
		goto tr0
	st57:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof57
		}
	stCase57:
		switch (m.data)[(m.p)] {
		case 84:
			goto st40
		case 116:
			goto st40
		}
		goto tr0
	st58:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof58
		}
	stCase58:
		switch (m.data)[(m.p)] {
		case 88:
			goto st40
		case 120:
			goto st40
		}
		goto tr0
	tr50:

		m.pb = m.p

		goto st59
	st59:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof59
		}
	stCase59:
		switch (m.data)[(m.p)] {
		case 69:
			goto st60
		case 101:
			goto st60
		}
		goto tr0
	st60:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof60
		}
	stCase60:
		switch (m.data)[(m.p)] {
		case 82:
			goto st61
		case 114:
			goto st61
		}
		goto tr0
	st61:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof61
		}
	stCase61:
		switch (m.data)[(m.p)] {
		case 70:
			goto st40
		case 102:
			goto st40
		}
		goto tr0
	tr51:

		m.pb = m.p

		goto st62
	st62:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof62
		}
	stCase62:
		switch (m.data)[(m.p)] {
		case 69:
			goto st63
		case 101:
			goto st63
		}
		goto tr0
	st63:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof63
		}
	stCase63:
		switch (m.data)[(m.p)] {
		case 70:
			goto st64
		case 86:
			goto st69
		case 102:
			goto st64
		case 118:
			goto st69
		}
		goto tr0
	st64:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof64
		}
	stCase64:
		switch (m.data)[(m.p)] {
		case 65:
			goto st65
		case 97:
			goto st65
		}
		goto tr0
	st65:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof65
		}
	stCase65:
		switch (m.data)[(m.p)] {
		case 67:
			goto st66
		case 99:
			goto st66
		}
		goto tr0
	st66:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof66
		}
	stCase66:
		switch (m.data)[(m.p)] {
		case 84:
			goto st67
		case 116:
			goto st67
		}
		goto tr0
	st67:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof67
		}
	stCase67:
		switch (m.data)[(m.p)] {
		case 79:
			goto st68
		case 111:
			goto st68
		}
		goto tr0
	st68:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof68
		}
	stCase68:
		switch (m.data)[(m.p)] {
		case 82:
			goto st40
		case 114:
			goto st40
		}
		goto tr0
	st69:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof69
		}
	stCase69:
		switch (m.data)[(m.p)] {
		case 69:
			goto st70
		case 101:
			goto st70
		}
		goto tr0
	st70:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof70
		}
	stCase70:
		switch (m.data)[(m.p)] {
		case 82:
			goto st57
		case 114:
			goto st57
		}
		goto tr0
	tr52:

		m.pb = m.p

		goto st71
	st71:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof71
		}
	stCase71:
		switch (m.data)[(m.p)] {
		case 84:
			goto st72
		case 116:
			goto st72
		}
		goto tr0
	st72:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof72
		}
	stCase72:
		switch (m.data)[(m.p)] {
		case 89:
			goto st73
		case 121:
			goto st73
		}
		goto tr0
	st73:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof73
		}
	stCase73:
		switch (m.data)[(m.p)] {
		case 76:
			goto st51
		case 108:
			goto st51
		}
		goto tr0
	tr53:

		m.pb = m.p

		goto st74
	st74:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof74
		}
	stCase74:
		switch (m.data)[(m.p)] {
		case 69:
			goto st75
		case 101:
			goto st75
		}
		goto tr0
	st75:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof75
		}
	stCase75:
		switch (m.data)[(m.p)] {
		case 83:
			goto st57
		case 115:
			goto st57
		}
		goto tr0
	stCase76:
		if 32 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 126 {
			goto tr89
		}
		goto tr0
	tr89:

		m.pb = m.p

		goto st77
	st77:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof77
		}
	stCase77:

		output._type = string(m.text())
		m.emitInfo("valid commit message type", "type", output._type)

		switch (m.data)[(m.p)] {
		case 33:
			goto tr91
		case 40:
			goto st82
		case 58:
			goto st79
		}
		if 32 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 126 {
			goto st77
		}
		goto tr6
	tr91:

		output.exclamation = true
		m.emitInfo("commit message communicates a breaking change")

		goto st78
	st78:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof78
		}
	stCase78:
		if (m.data)[(m.p)] == 58 {
			goto st79
		}
		goto tr6
	st79:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof79
		}
	stCase79:
		if (m.data)[(m.p)] == 32 {
			goto st80
		}
		goto tr10
	st80:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof80
		}
	stCase80:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr13
		case 32:
			goto st80
		}
		goto tr95
	tr95:

		m.pb = m.p

		goto st95
	st95:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof95
		}
	stCase95:
		if (m.data)[(m.p)] == 10 {
			goto tr114
		}
		goto st95
	tr114:

		output.descr = string(m.text())
		m.emitInfo("valid commit message description", "description", output.descr)

		goto st81
	st81:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof81
		}
	stCase81:
		if (m.data)[(m.p)] == 10 {
			goto tr96
		}
		goto tr14
	tr96:

		m.emitDebug("found a blank line", "pos", m.p)

		m.emitDebug("try to parse a footer trailer token", "pos", m.p)
		{
			goto st87
		}

		goto st96
	st96:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof96
		}
	stCase96:
		goto st0
	st82:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof82
		}
	stCase82:
		if (m.data)[(m.p)] == 41 {
			goto tr98
		}
		switch {
		case (m.data)[(m.p)] > 39:
			if 42 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 126 {
				goto tr97
			}
		case (m.data)[(m.p)] >= 32:
			goto tr97
		}
		goto tr16
	tr97:

		m.pb = m.p

		goto st83
	st83:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof83
		}
	stCase83:
		if (m.data)[(m.p)] == 41 {
			goto tr100
		}
		switch {
		case (m.data)[(m.p)] > 39:
			if 42 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 126 {
				goto st83
			}
		case (m.data)[(m.p)] >= 32:
			goto st83
		}
		goto tr16
	tr98:

		m.pb = m.p

		output.scope = string(m.text())
		m.emitInfo("valid commit message scope", "scope", output.scope)

		goto st84
	tr100:

		output.scope = string(m.text())
		m.emitInfo("valid commit message scope", "scope", output.scope)

		goto st84
	st84:

		if (m.p + 1) == m.pe {
			m.err = m.emitErrorOnCurrentCharacter(ErrEarly)
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof84
		}
	stCase84:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr91
		case 58:
			goto st79
		}
		goto tr6
	tr103:

		// Increment number of newlines to use in case we're still in the body
		m.countNewlines++
		m.lastNewline = m.p
		m.emitDebug("found a newline", "pos", m.p)

		goto st87
	st87:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof87
		}
	stCase87:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr103
		case 66:
			goto tr105
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr104
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr104
			}
		default:
			goto tr104
		}
		goto tr21
	tr104:

		m.pb = m.p

		goto st14
	st14:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof14
		}
	stCase14:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 45:
			goto st16
		case 58:
			goto tr25
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st14
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr21
	tr22:

		// todo > alnum[[- ]alnum] string to lower can be more performant?
		m.currentFooterKey = string(bytes.ToLower(m.text()))
		if m.currentFooterKey == "breaking change" {
			m.currentFooterKey = "breaking-change"
		}
		m.emitDebug("possibly valid footer token", "token", m.currentFooterKey, "pos", m.p)

		goto st15
	st15:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof15
		}
	stCase15:
		if (m.data)[(m.p)] == 35 {
			goto tr26
		}
		goto tr21
	tr26:

		m.emitDebug("try to parse a footer trailer value", "pos", m.p)
		{
			goto st33
		}

		goto st88
	st88:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof88
		}
	stCase88:
		goto st0
	st16:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof16
		}
	stCase16:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st14
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr21
	tr25:

		// todo > alnum[[- ]alnum] string to lower can be more performant?
		m.currentFooterKey = string(bytes.ToLower(m.text()))
		if m.currentFooterKey == "breaking change" {
			m.currentFooterKey = "breaking-change"
		}
		m.emitDebug("possibly valid footer token", "token", m.currentFooterKey, "pos", m.p)

		goto st17
	st17:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof17
		}
	stCase17:
		if (m.data)[(m.p)] == 32 {
			goto tr27
		}
		goto tr21
	tr27:

		m.emitDebug("try to parse a footer trailer value", "pos", m.p)
		{
			goto st33
		}

		goto st89
	st89:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof89
		}
	stCase89:
		if (m.data)[(m.p)] == 32 {
			goto tr27
		}
		goto st0
	tr105:

		m.pb = m.p

		goto st18
	st18:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof18
		}
	stCase18:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 45:
			goto st16
		case 58:
			goto tr25
		case 82:
			goto st19
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st14
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr21
	st19:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof19
		}
	stCase19:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 45:
			goto st16
		case 58:
			goto tr25
		case 69:
			goto st20
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st14
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr21
	st20:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof20
		}
	stCase20:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 45:
			goto st16
		case 58:
			goto tr25
		case 65:
			goto st21
		}
		switch {
		case (m.data)[(m.p)] < 66:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st14
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr21
	st21:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof21
		}
	stCase21:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 45:
			goto st16
		case 58:
			goto tr25
		case 75:
			goto st22
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st14
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr21
	st22:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof22
		}
	stCase22:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 45:
			goto st16
		case 58:
			goto tr25
		case 73:
			goto st23
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st14
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr21
	st23:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof23
		}
	stCase23:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 45:
			goto st16
		case 58:
			goto tr25
		case 78:
			goto st24
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st14
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr21
	st24:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof24
		}
	stCase24:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 45:
			goto st16
		case 58:
			goto tr25
		case 71:
			goto st25
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st14
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr21
	st25:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof25
		}
	stCase25:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr35
		case 45:
			goto st16
		case 58:
			goto tr25
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st14
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr21
	tr35:

		// todo > alnum[[- ]alnum] string to lower can be more performant?
		m.currentFooterKey = string(bytes.ToLower(m.text()))
		if m.currentFooterKey == "breaking change" {
			m.currentFooterKey = "breaking-change"
		}
		m.emitDebug("possibly valid footer token", "token", m.currentFooterKey, "pos", m.p)

		goto st26
	st26:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof26
		}
	stCase26:
		switch (m.data)[(m.p)] {
		case 35:
			goto tr26
		case 67:
			goto st27
		}
		goto tr21
	st27:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof27
		}
	stCase27:
		if (m.data)[(m.p)] == 72 {
			goto st28
		}
		goto tr21
	st28:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof28
		}
	stCase28:
		if (m.data)[(m.p)] == 65 {
			goto st29
		}
		goto tr21
	st29:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof29
		}
	stCase29:
		if (m.data)[(m.p)] == 78 {
			goto st30
		}
		goto tr21
	st30:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof30
		}
	stCase30:
		if (m.data)[(m.p)] == 71 {
			goto st31
		}
		goto tr21
	st31:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof31
		}
	stCase31:
		if (m.data)[(m.p)] == 69 {
			goto st32
		}
		goto tr21
	st32:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof32
		}
	stCase32:
		if (m.data)[(m.p)] == 58 {
			goto tr25
		}
		goto tr21
	stOut:
	_testEof2:
		m.cs = 2
		goto _testEof
	_testEof3:
		m.cs = 3
		goto _testEof
	_testEof4:
		m.cs = 4
		goto _testEof
	_testEof5:
		m.cs = 5
		goto _testEof
	_testEof6:
		m.cs = 6
		goto _testEof
	_testEof7:
		m.cs = 7
		goto _testEof
	_testEof8:
		m.cs = 8
		goto _testEof
	_testEof85:
		m.cs = 85
		goto _testEof
	_testEof9:
		m.cs = 9
		goto _testEof
	_testEof86:
		m.cs = 86
		goto _testEof
	_testEof10:
		m.cs = 10
		goto _testEof
	_testEof11:
		m.cs = 11
		goto _testEof
	_testEof12:
		m.cs = 12
		goto _testEof
	_testEof13:
		m.cs = 13
		goto _testEof
	_testEof33:
		m.cs = 33
		goto _testEof
	_testEof90:
		m.cs = 90
		goto _testEof
	_testEof91:
		m.cs = 91
		goto _testEof
	_testEof34:
		m.cs = 34
		goto _testEof
	_testEof92:
		m.cs = 92
		goto _testEof
	_testEof36:
		m.cs = 36
		goto _testEof
	_testEof37:
		m.cs = 37
		goto _testEof
	_testEof38:
		m.cs = 38
		goto _testEof
	_testEof39:
		m.cs = 39
		goto _testEof
	_testEof40:
		m.cs = 40
		goto _testEof
	_testEof41:
		m.cs = 41
		goto _testEof
	_testEof42:
		m.cs = 42
		goto _testEof
	_testEof43:
		m.cs = 43
		goto _testEof
	_testEof93:
		m.cs = 93
		goto _testEof
	_testEof44:
		m.cs = 44
		goto _testEof
	_testEof94:
		m.cs = 94
		goto _testEof
	_testEof45:
		m.cs = 45
		goto _testEof
	_testEof46:
		m.cs = 46
		goto _testEof
	_testEof47:
		m.cs = 47
		goto _testEof
	_testEof48:
		m.cs = 48
		goto _testEof
	_testEof49:
		m.cs = 49
		goto _testEof
	_testEof50:
		m.cs = 50
		goto _testEof
	_testEof51:
		m.cs = 51
		goto _testEof
	_testEof52:
		m.cs = 52
		goto _testEof
	_testEof53:
		m.cs = 53
		goto _testEof
	_testEof54:
		m.cs = 54
		goto _testEof
	_testEof55:
		m.cs = 55
		goto _testEof
	_testEof56:
		m.cs = 56
		goto _testEof
	_testEof57:
		m.cs = 57
		goto _testEof
	_testEof58:
		m.cs = 58
		goto _testEof
	_testEof59:
		m.cs = 59
		goto _testEof
	_testEof60:
		m.cs = 60
		goto _testEof
	_testEof61:
		m.cs = 61
		goto _testEof
	_testEof62:
		m.cs = 62
		goto _testEof
	_testEof63:
		m.cs = 63
		goto _testEof
	_testEof64:
		m.cs = 64
		goto _testEof
	_testEof65:
		m.cs = 65
		goto _testEof
	_testEof66:
		m.cs = 66
		goto _testEof
	_testEof67:
		m.cs = 67
		goto _testEof
	_testEof68:
		m.cs = 68
		goto _testEof
	_testEof69:
		m.cs = 69
		goto _testEof
	_testEof70:
		m.cs = 70
		goto _testEof
	_testEof71:
		m.cs = 71
		goto _testEof
	_testEof72:
		m.cs = 72
		goto _testEof
	_testEof73:
		m.cs = 73
		goto _testEof
	_testEof74:
		m.cs = 74
		goto _testEof
	_testEof75:
		m.cs = 75
		goto _testEof
	_testEof77:
		m.cs = 77
		goto _testEof
	_testEof78:
		m.cs = 78
		goto _testEof
	_testEof79:
		m.cs = 79
		goto _testEof
	_testEof80:
		m.cs = 80
		goto _testEof
	_testEof95:
		m.cs = 95
		goto _testEof
	_testEof81:
		m.cs = 81
		goto _testEof
	_testEof96:
		m.cs = 96
		goto _testEof
	_testEof82:
		m.cs = 82
		goto _testEof
	_testEof83:
		m.cs = 83
		goto _testEof
	_testEof84:
		m.cs = 84
		goto _testEof
	_testEof87:
		m.cs = 87
		goto _testEof
	_testEof14:
		m.cs = 14
		goto _testEof
	_testEof15:
		m.cs = 15
		goto _testEof
	_testEof88:
		m.cs = 88
		goto _testEof
	_testEof16:
		m.cs = 16
		goto _testEof
	_testEof17:
		m.cs = 17
		goto _testEof
	_testEof89:
		m.cs = 89
		goto _testEof
	_testEof18:
		m.cs = 18
		goto _testEof
	_testEof19:
		m.cs = 19
		goto _testEof
	_testEof20:
		m.cs = 20
		goto _testEof
	_testEof21:
		m.cs = 21
		goto _testEof
	_testEof22:
		m.cs = 22
		goto _testEof
	_testEof23:
		m.cs = 23
		goto _testEof
	_testEof24:
		m.cs = 24
		goto _testEof
	_testEof25:
		m.cs = 25
		goto _testEof
	_testEof26:
		m.cs = 26
		goto _testEof
	_testEof27:
		m.cs = 27
		goto _testEof
	_testEof28:
		m.cs = 28
		goto _testEof
	_testEof29:
		m.cs = 29
		goto _testEof
	_testEof30:
		m.cs = 30
		goto _testEof
	_testEof31:
		m.cs = 31
		goto _testEof
	_testEof32:
		m.cs = 32
		goto _testEof

	_testEof:
		{
		}
		if (m.p) == (m.eof) {
			switch m.cs {
			case 2, 3, 4, 13, 36, 37, 38, 39, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75:

				if m.pe > 0 {
					if m.p != m.pe {
						m.err = m.emitErrorOnCurrentCharacter(ErrType)
					} else {
						// assert(m.p == m.pe)
						m.err = m.emitErrorOnPreviousCharacter(ErrTypeIncomplete)
					}
				}

			case 5, 6, 12, 40, 41, 47, 77, 78, 84:

				if m.err == nil {
					m.err = m.emitErrorOnCurrentCharacter(ErrColon)
				}

			case 7, 42, 79:

				if m.err == nil {
					m.err = m.emitErrorOnCurrentCharacter(ErrDescriptionInit)
				}

			case 8, 43, 80:

				if m.p < m.pe && m.data[m.p] == 10 {
					m.err = m.emitError(ErrNewline, m.p+1)
				} else {
					// assert(m.p == m.pe)
					m.err = m.emitErrorOnPreviousCharacter(ErrDescription)
				}

			case 9, 44, 81:

				m.err = m.emitErrorWithoutCharacter(ErrMissingBlankLineAtBeginning)

			case 85, 93, 95:

				output.descr = string(m.text())
				m.emitInfo("valid commit message description", "description", output.descr)

			case 90:

				output.footers[m.currentFooterKey] = append(output.footers[m.currentFooterKey], string(m.text()))
				m.emitInfo("valid commit message footer trailer", m.currentFooterKey, string(m.text()))

			case 92:

				// Append newlines
				for m.countNewlines > 0 {
					output.body += "\n"
					m.countNewlines--
					m.emitInfo("valid commit message body content", "body", "\n")
				}
				// Append body content
				output.body += string(m.text())
				m.emitInfo("valid commit message body content", "body", string(m.text()))

			case 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32:

				if len(output.footers) == 0 {
					// Backtrack to the last marker
					// Ie., the text possibly a trailer token that is instead part of the body content
					if m.countNewlines > 0 {
						// In case new lines met while rewinding
						// advance the last marker by the number of the newlines so that they don't get parsed again
						// (they be added in the result by the body content appender)
						m.pb = m.lastNewline + 1
					}
					(m.p) = (m.pb) - 1

					m.emitDebug("try to parse body content", "pos", m.p)
					{
						goto st34
					}
				} else {
					// A rewind happens when an error while parsing a footer trailer is encountered
					// If this is not the first footer trailer the parser can't go back to parse body content again
					// Thus, emit an error
					if m.p != m.pe {
						m.err = m.emitErrorOnCurrentCharacter(ErrTrailer)
					} else {
						// assert(m.p == m.pe)
						m.err = m.emitErrorOnPreviousCharacter(ErrTrailerIncomplete)
					}
				}

			case 1, 35, 76:

				m.err = m.emitErrorWithoutCharacter(ErrEmpty)

				if m.pe > 0 {
					if m.p != m.pe {
						m.err = m.emitErrorOnCurrentCharacter(ErrType)
					} else {
						// assert(m.p == m.pe)
						m.err = m.emitErrorOnPreviousCharacter(ErrTypeIncomplete)
					}
				}

			case 10, 11, 45, 46, 82, 83:

				if m.p < m.pe {
					m.err = m.emitErrorOnCurrentCharacter(ErrScope)
				}

				// assert(m.p == m.pe)
				m.err = m.emitErrorOnPreviousCharacter(ErrScopeIncomplete)

			case 34:

				// Append newlines
				for m.countNewlines > 0 {
					output.body += "\n"
					m.countNewlines--
					m.emitInfo("valid commit message body content", "body", "\n")
				}
				// Append body content
				output.body += string(m.text())
				m.emitInfo("valid commit message body content", "body", string(m.text()))

				m.emitDebug("try to parse a footer trailer token", "pos", m.p)
				{
					goto st87
				}

			}
		}

	_out:
		{
		}
	}

	if m.cs < firstFinal {
		if m.bestEffort && output.minimal() {
			// An error occurred but partial parsing is on and partial message is minimally valid
			return output.export(), m.err
		}
		return nil, m.err
	}

	return output.export(), nil
}

// WithBestEffort enables best effort mode.
func (m *machine) WithBestEffort() {
	m.bestEffort = true
}

// HasBestEffort tells whether the receiving machine has best effort mode on or off.
func (m *machine) HasBestEffort() bool {
	return m.bestEffort
}

// WithTypes tells the parser which commit message types to consider.
func (m *machine) WithTypes(t conventionalcommits.TypeConfig) {
	m.typeConfig = t
}

// WithLogger tells the parser which logger to use.
func (m *machine) WithLogger(l *logrus.Logger) {
	m.logger = l
}
