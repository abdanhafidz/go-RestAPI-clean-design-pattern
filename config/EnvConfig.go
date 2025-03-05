package config

import "os"

var TCP_ADDRESS string
var LOG_PATH string
var HOST_ADDRESS string
var HOST_PORT string

func init() {
	HOST_ADDRESS = os.Getenv("HOST_ADDRESS")
	HOST_PORT = os.Getenv("HOST_PORT")
	TCP_ADDRESS = HOST_ADDRESS + ":" + HOST_PORT
	LOG_PATH = os.Getenv("LOG_PATH")
	// Menampilkan nilai variabel lingkungan
}
