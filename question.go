package main

type Question struct {
	// Refers to the question number
	Num int
	// The question itself
	Body string
	// The answers
	Answers []Answer
}

type Answer struct {
	Body string
	// Indicates correctness of the answer
	Correct bool
}
