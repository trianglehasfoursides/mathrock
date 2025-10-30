package main

var conf *configuration

type configuration struct {
	Storage string `json:"storage"`
}

func setupconf() (err error) {
}
