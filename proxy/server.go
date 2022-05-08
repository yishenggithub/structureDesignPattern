package main

type server interface {
	handleResquest(string, string) (int, string)
}
