package exec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestSimpleCommand(t *testing.T) {
	testcases := []struct {
		Command        string
		ExpectedOutput string
		Error          bool
	}{
		{
			Command:        "echo hallo",
			ExpectedOutput: "hallo\n",
			Error:          false,
		},
		{
			Command:        "for i in 1 2 3 4 5; do echo $i; done",
			ExpectedOutput: "1\n2\n3\n4\n5\n",
			Error:          false,
		},
		{
			Command:        `[[ "ok" == "ok" ]] && echo "match"`,
			ExpectedOutput: "match\n",
			Error:          false,
		},
		{
			Command:        `[[ "ok" != "ok" ]] && echo "match"`,
			ExpectedOutput: "",
			Error:          false,
		},
		{
			Command:        `cat nofile`,
			ExpectedOutput: "",
			Error:          true,
		},
		{
			Command:        `for i in 1 2 3; do echo $(hostname); done`,
			ExpectedOutput: "sva-tm\nsva-tm\nsva-tm\n",
			Error:          false,
		},
		{
			Command:        `echo "many words dont matter" | cut -d " " -f4-`,
			ExpectedOutput: "matter\n",
			Error:          false,
		},
		{
			Command:        `TEST=var; echo $TEST`,
			ExpectedOutput: "var\n",
			Error:          false,
		},
		{
			Command:        `TEST=var; echo $TEST`,
			ExpectedOutput: "var\n",
			Error:          false,
		},
	}

	for _, tc := range testcases {
		out, err := RunCommand(tc.Command)

		if tc.Error {
			assert.Error(t, err)
		} else {
			assert.Equal(t, tc.ExpectedOutput, out)
		}
	}
}