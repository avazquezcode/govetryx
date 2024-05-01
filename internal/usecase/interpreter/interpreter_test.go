package interpreter_test

import (
	"bytes"
	"testing"

	interpreter_pkg "github.com/avazquezcode/govetryx/internal/usecase/interpreter"
	"github.com/avazquezcode/govetryx/internal/usecase/parser"
	"github.com/avazquezcode/govetryx/internal/usecase/scanner"
	"github.com/stretchr/testify/assert"
)

func TestInterpret(t *testing.T) {
	// I use heavily the print function to check the interpreter results of different operations
	tests := map[string]struct {
		src            string
		expectedStdout string
		expectedErr    bool
	}{
		// variable declaration and assignment
		"variable declaration with assignment": {
			src:            "dec a = 1\n print a\n",
			expectedStdout: "1\n",
			expectedErr:    false,
		},
		"variable declaration without asignment": {
			src:            "dec a\n print a\n",
			expectedStdout: "null\n",
			expectedErr:    false,
		},
		"variable declaration with assignment based on another variable": {
			src:            "dec a = 1\n dec b = a\n print a\n print b\n",
			expectedStdout: "1\n1\n",
			expectedErr:    false,
		},
		// binary operations
		"sum of 2 values": {
			src:            "print 1 + 1\n",
			expectedStdout: "2\n",
			expectedErr:    false,
		},
		"multiplication of 2 values": {
			src:            "print 2 * 2\n",
			expectedStdout: "4\n",
			expectedErr:    false,
		},
		"division of 2 values": {
			src:            "print 2 / 2\n",
			expectedStdout: "1\n",
			expectedErr:    false,
		},
		"substraction of 2 values": {
			src:            "print 4 - 2\n",
			expectedStdout: "2\n",
			expectedErr:    false,
		},
		"modulus between 2 values": {
			src:            "print 4 % 2\n",
			expectedStdout: "0\n",
			expectedErr:    false,
		},
		// unary operations
		"negation": {
			src:            "dec a = 1\n print -a\n",
			expectedStdout: "-1\n",
			expectedErr:    false,
		},
		"bang operation": {
			src:            "dec a = true\n print !a\n",
			expectedStdout: "false\n",
			expectedErr:    false,
		},
		// grouping
		"operation with grouping": {
			src:            "print (1+1) * 2\n",
			expectedStdout: "4\n",
			expectedErr:    false,
		},
		// functions
		"function declaration (not being called)": {
			src:            "fn a(){}",
			expectedStdout: "",
			expectedErr:    false,
		},
		"function declaration (being called - returns nil)": {
			src:            "fn a(){} print a()\n",
			expectedStdout: "null\n",
			expectedErr:    false,
		},
		"function declaration (being called - returns a value)": {
			src:            "fn a(){return 1\n} print a()\n",
			expectedStdout: "1\n",
			expectedErr:    false,
		},
		// if
		"simple if condition": {
			src:            "if 1 == 1 {print true\n}",
			expectedStdout: "true\n",
			expectedErr:    false,
		},
		"if condition with else": {
			src:            "if 1 == 1 {print true\n} else {print false\n}",
			expectedStdout: "true\n",
			expectedErr:    false,
		},
		"if condition evaluates to false, with else": {
			src:            "if 1 == 2 {print true\n} else {print false\n}",
			expectedStdout: "false\n",
			expectedErr:    false,
		},
		"if condition with AND operation (both evaluate to true)": {
			src:            "if 1 == 1 && 2 == 2 {print true\n} else {print false\n}",
			expectedStdout: "true\n",
			expectedErr:    false,
		},
		"if condition with AND operation (one evaluate to false)": {
			src:            "if 1 == 2 && 2 == 2 {print true\n} else {print false\n}",
			expectedStdout: "false\n",
			expectedErr:    false,
		},
		"if condition with AND operation (both evaluate to false)": {
			src:            "if 1 == 2 && 2 == 3 {print true\n} else {print false\n}",
			expectedStdout: "false\n",
			expectedErr:    false,
		},
		"if condition with OR operation (both evaluate to true)": {
			src:            "if 1 == 1 || 2 == 2 {print true\n} else {print false\n}",
			expectedStdout: "true\n",
			expectedErr:    false,
		},
		"if condition with OR operation (one evaluate to false)": {
			src:            "if 1 == 2 || 2 == 2 {print true\n} else {print false\n}",
			expectedStdout: "true\n",
			expectedErr:    false,
		},
		"if condition with OR operation (both conditions evaluate to false)": {
			src:            "if 1 == 2 || 2 == 3 {print true\n} else {print false\n}",
			expectedStdout: "false\n",
			expectedErr:    false,
		},
		// while
		"while simple loop": {
			src:            "dec a = 0\n while (a < 3) {print a\n a = a + 1\n}",
			expectedStdout: "0\n1\n2\n",
			expectedErr:    false,
		},
		// closures
		"closures": {
			src:            "fn a() { fn b() { return 1\n } return b()\n } print a()\n",
			expectedStdout: "1\n",
			expectedErr:    false,
		},
		// scoping
		"scoping": {
			src:            "dec a = 1\n {dec a = 2\n print a\n}",
			expectedStdout: "2\n",
			expectedErr:    false,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			lexer := scanner.NewScanner(bytes.Runes(strToBytes(test.src)))
			tokens, _ := lexer.Scan() // We are generating a src that we know is valid, so no need for handling error here
			parser := parser.NewParser(tokens)
			statements, err := parser.Parse()
			if err != nil {
				t.Fail()
			}

			var testStdOut bytes.Buffer
			interpreter := interpreter_pkg.NewInterpreter(&testStdOut)
			resolver := interpreter_pkg.NewResolver(interpreter)
			err = resolver.Resolve(statements)
			if err != nil {
				t.Fail()
			}

			err = interpreter.Interpret(statements)
			assert.Equal(t, test.expectedErr, err != nil)
			assert.Equal(t, test.expectedStdout, testStdOut.String())
		})
	}
}

func strToBytes(str string) []byte {
	return []byte(str)
}
