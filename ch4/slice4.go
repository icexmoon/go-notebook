package main

type slice struct{
	pointer *int[]
	length int
	cap int
}