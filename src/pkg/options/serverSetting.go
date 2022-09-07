package options

import "time"

type Server struct {
	RunMode      string        `ini:"RunMode"`
	HttpPort     int           `ini:"HttpPort"`
	ReadTimeout  time.Duration `ini:"ReadTimeout"`
	WriteTimeout time.Duration `ini:"WriteTimeout"`
}
