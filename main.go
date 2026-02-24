package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// `StackOp` is a function that operates on the stack and returns a boolean
// error indicator.
type StackOp func(*Stack) bool

// `Env` holds the environment of a running program as a map from string keys
// to operations on the stack.
type Env map[string]StackOp

// `CompileBuffer` holds the state of a word being compiled before it is added
// to the environment.
type CompileBuffer struct {
	name string
	body []StackOp
}

// `State` holds the full execution state of a running program.
//
// `env`: program environment that holds the "words"
// `stack`: execution stack when in immediate mode
// `compileBuffer`: struct that captures intermediate compile state; `nil` means not compiling
type State struct {
	env           Env
	stack         Stack
	compileBuffer *CompileBuffer
}

// Read an input from the prompt and return a slice of tokens.
func read(prompt string) []string {
	fmt.Print(prompt + " ")
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	tokens := tokenize(command)
	return tokens
}

// Tokenize an input.
//
// This lowercases the input and splits by space-separated character groups
// with all extraneous whitespace removed.
func tokenize(command string) []string {
	return strings.Fields(strings.ToLower(command))
}

// Evaluate a slice of tokens in the context of the current program state.
//
// Note: This returns an error boolean, which can later be elevated to a
// complete error message.
func evaluate(tokens []string, state *State) bool {
	for _, token := range tokens {

		// ":" and ";" start and end word compilation.
		//
		// Note that ":" and ";" are treated as special characters, rather than
		// regular words.
		switch token {
		case ":":
			// Initiazlize the compile buffer.
			state.compileBuffer = &CompileBuffer{}
			continue
		case ";":
			// When compilation ends, translate the compile buffer into a word in the
			// environment, and then reset the buffer.
			compileWord(state.compileBuffer, state.env)
			state.compileBuffer = nil
			continue
		}

		// Compilation is stateful. First, the name is captured, and then
		// the implementation is collected into a slice that represents the sequence
		// of operations.
		//
		// If an existing word is encountered, it's stack operation is appended to
		// the buffer.
		//
		// If a literal (an integer) is encountered, it is captured in a closure that
		// pushes the value to the stack when the compiled word is executed.
		if state.compileBuffer != nil {
			if state.compileBuffer.name == "" {
				state.compileBuffer.name = token
				continue
			}
			if word, ok := state.env[token]; ok {
				state.compileBuffer.body = append(state.compileBuffer.body, word)
			} else if value, error := strconv.ParseInt(token, 10, 64); error == nil {
				state.compileBuffer.body = append(state.compileBuffer.body, makePushInt(value))
			} else {
				return true
			}
			continue
		}

		// When not in compilation mode, the token (a word or literal) is interpreted
		// immediately.
		if interpretToken(token, state) {
			return true
		}
	}
	return false
}

// Given a compile buffer, create the compiled word and add it to the
// evironment.
func compileWord(buffer *CompileBuffer, env Env) bool {
	f := func(stack *Stack) bool {
		// Apply each function in the compile buffer to the stack in sequence.
		for _, f := range buffer.body {
			if f(stack) {
				return true
			}
		}
		return false
	}

	// Add the compiled word to the environment.
	env[buffer.name] = f

	return false
}

// Return a closure that pushes a literal (integer) onto the stack.
func makePushInt(value int64) StackOp {
	return func(stack *Stack) bool {
		stack.Push(value)
		return false
	}
}

// Evaluate a single token against the current state.
//
// Note: This returns an error boolean, which can later be elevated to a
// complete error message.
func interpretToken(token string, state *State) bool {
	// If the token is in the environment, perform the corresponding operation
	// on the stack.
	if fun, ok := state.env[token]; ok {
		return fun(&state.stack)
	}

	// If the token can be parsed as an integer, push the parsed value onto the stack.
	if value, error := strconv.ParseInt(token, 10, 64); error == nil {
		state.stack.Push(value)
		return false
	}

	// Unrecognized token.
	return true
}

// Start the REPL.
func repl(prompt string, state *State) {
	command := read(prompt)
	if err := evaluate(command, state); err {
		fmt.Println("operation failed")
	}
	repl(prompt, state)
}

// Define the standard library.
//
// Note: `var` is needed for package-level variable declarations.
var std = Env{
	"+":    (*Stack).Add,
	"-":    (*Stack).Subtract,
	"*":    (*Stack).Multiply,
	"/":    (*Stack).Divide,
	">":    (*Stack).GreaterThan,
	"<":    (*Stack).LessThan,
	"=":    (*Stack).EqualTo,
	"and":  (*Stack).And,
	"or":   (*Stack).Or,
	".":    (*Stack).Dot,
	"dup":  (*Stack).Duplicate,
	"swap": (*Stack).Swap,
}

func main() {
	state := State{
		env:           std,
		stack:         NewStack(),
		compileBuffer: nil,
	}
	repl("goforth>", &state)
}
