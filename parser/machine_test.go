package parser

import (
	"fmt"
	"testing"

	"github.com/reviewpad/go-conventionalcommits"
	"github.com/sirupsen/logrus"
	logrustest "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestMachineParse(t *testing.T) {
	runner(t, "minimaltypes", testCases, conventionalcommits.TypesMinimal, WithTypes(conventionalcommits.TypesMinimal))
}

func TestMachineParseWithConventionalTypes(t *testing.T) {
	runner(t, "conventionaltypes", testCasesForConventionalTypes, conventionalcommits.TypesConventional, WithTypes(conventionalcommits.TypesConventional))
}

func TestMachineParseWithFreeFormTypes(t *testing.T) {
	runner(t, "freeformtypes", testCasesForFreeFormTypes, conventionalcommits.TypesFreeForm, WithTypes(conventionalcommits.TypesFreeForm))
}

func setTypeConfig(msg conventionalcommits.Message, tc conventionalcommits.TypeConfig) {
	if cc, ok := msg.(*conventionalcommits.ConventionalCommit); ok {
		cc.TypeConfig = tc
	}
}

func runner(t *testing.T, label string, cases []testCase, typeConfig conventionalcommits.TypeConfig, machineOpts ...conventionalcommits.MachineOption) {
	t.Helper()

	for _, tc := range cases {
		tc := tc
		title := fmt.Sprintf("%s/%s", label, tc.title)

		setTypeConfig(tc.value, typeConfig)
		setTypeConfig(tc.partialValue, typeConfig)

		t.Run(title, func(t *testing.T) {
			message, messageErr := NewMachine(machineOpts...).Parse(tc.input)
			partial, partialErr := NewMachine(append(machineOpts, WithBestEffort())...).Parse(tc.input)

			if !tc.ok {
				assert.Nil(t, message)
				assert.Error(t, messageErr)
				assert.EqualError(t, messageErr, tc.errorString)

				if partial != nil {
					assert.True(t, partial.Ok())
				}
				assert.Equal(t, tc.partialValue, partial)
				assert.EqualError(t, partialErr, tc.errorString)
			} else {
				assert.Nil(t, messageErr)
				assert.NotEmpty(t, message)
				assert.True(t, message.Ok())
				assert.Equal(t, message, partial)
				assert.Equal(t, tc.partialValue, partial)
				assert.Equal(t, messageErr, partialErr)
			}

			assert.Equal(t, tc.value, message)
		})
	}
}

func TestMachineBestEffortOption(t *testing.T) {
	p1 := NewMachine().(conventionalcommits.BestEfforter)
	assert.False(t, p1.HasBestEffort())

	p2 := NewMachine(WithBestEffort()).(conventionalcommits.BestEfforter)
	assert.True(t, p2.HasBestEffort())
}

func TestParseLoggingErrorsOnly(t *testing.T) {
	l, hook := logrustest.NewNullLogger()
	l.SetLevel(logrus.ErrorLevel)

	p := NewMachine(WithLogger(l))
	p.Parse([]byte("fix: a wonderful logger\x0Aaaa"))

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "missing a blank line: col=24", hook.LastEntry().Message)

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}

func TestParseLoggingEverything(t *testing.T) {
	l := logrus.New()
	hook := logrustest.NewLocal(l)

	p := NewMachine(WithLogger(l))
	p.Parse([]byte("fix: a wonderful logger\x0Aaaa"))

	var logEntries = hook.AllEntries()
	assert.Equal(t, 3, len(logEntries))
	assert.Equal(t, logrus.InfoLevel, logEntries[0].Level)
	assert.Equal(t, logrus.InfoLevel, logEntries[1].Level)
	assert.Equal(t, logrus.ErrorLevel, logEntries[2].Level)
	assert.Equal(t, "fix", logEntries[0].Data["type"])
	assert.Equal(t, "valid commit message type", logEntries[0].Message)
	assert.Equal(t, "a wonderful logger", logEntries[1].Data["description"])
	assert.Equal(t, "valid commit message description", logEntries[1].Message)
	assert.Equal(t, "missing a blank line: col=24", logEntries[2].Message)

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}
