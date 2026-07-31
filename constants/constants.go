package main

const PI = 3.14
const GRAVITY = 9.81

func main() {

	const days int = 7

	const (
		monday       = 1
		tuesday      = 2
		wenesday     = 3
		thursday int = 4
	)

}

// // ---- CONSTANTS ----

// // "const" declares a value that is LOCKED once set
// // Attempting to reassign it later = compile error
// const Pi = 3.14159

// // Constants must be assigned a value immediately — no zero-value default
// // This would NOT compile:
// // const Radius int   ❌ missing value

// // Multiple constants can be grouped together
// const (
//     StatusActive   = "active"
//     StatusInactive = "inactive"
//     MaxRetries     = 3
// )

// // Constants are typically used for values that should NEVER
// // change during program execution — config values, fixed limits,
// // mathematical constants, status labels, etc.

// // Naming convention note:
// // Capitalized constants (like MaxRetries) are "exported" —
// // accessible from other files/packages, same rule as PascalCase structs
