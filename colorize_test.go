// Implement tests for the `colorize` library
package colorize

import "testing"

// Validate `colorize.ColorString()`
func TestColorString(test *testing.T) {
    // Define our test cases for the colorize with string color cases
    cases := [] struct {
        input, color, expected string
    }{
        // Non-standard cases
        { "Test", "BLACK",   "\x1b[30mTest\x1b[0m" },
        { "Test", "Black",   "\x1b[30mTest\x1b[0m" },
        { "Test", "INVALID", "Test" },
        { "Test", "", "Test" },

        // Typical usage
        { "Test", "black",   "\x1b[30mTest\x1b[0m" },
        { "Test", "red",     "\x1b[31mTest\x1b[0m" },
        { "Test", "green",   "\x1b[32mTest\x1b[0m" },
        { "Test", "yellow",  "\x1b[33mTest\x1b[0m" },
        { "Test", "blue",    "\x1b[34mTest\x1b[0m" },
        { "Test", "magenta", "\x1b[35mTest\x1b[0m" },
        { "Test", "cyan",    "\x1b[36mTest\x1b[0m" },
        { "Test", "white",   "\x1b[37mTest\x1b[0m" },
    }

    // Run tests
    for _, tc := range cases {
        actual := ColorString(tc.input, tc.color)
        if actual != tc.expected {
            test.Errorf("ColorString(%q, %q) == %q, expected %q",
                        tc.input, tc.color, actual, tc.expected)
        }
    }
}

// Validate `colorize.Colorize()`
func TestColorize(test *testing.T) {
    // Define our test cases for the colorize with string color cases
    cases := [] struct {
        input, expected string
    }{
        { "Test <red>A</red> <blue>B</blue>", "Test \x1b[31mA\x1b[0m \x1b[34mB\x1b[0m" },
        { "Test <red>A</red> <red>B</red>",   "Test \x1b[31mA\x1b[0m \x1b[31mB\x1b[0m" },
    }

    // Run tests
    for _, tc := range cases {
        actual := Colorize(tc.input)
        if actual != tc.expected {
            test.Errorf("Colorize(%q) == %q, expected %q",
                        tc.input, actual, tc.expected)
        }
    }
}
