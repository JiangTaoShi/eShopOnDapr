package env

import (
	"os"
)

var (
	active Environment
	dev    Environment = &environment{value: "dev"}
	sit    Environment = &environment{value: "sit"}
	uat    Environment = &environment{value: "uat"}
	pro    Environment = &environment{value: "pro"}
)

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsSIT() bool
	IsUat() bool
	IsPro() bool
}

type environment struct {
	Environment
	value string
}

func (e *environment) Value() string {
	return e.value
}

func (e *environment) IsDev() bool {
	return e.value == "dev"
}

func (e *environment) IsSIT() bool {
	return e.value == "sit"
}

func (e *environment) IsUat() bool {
	return e.value == "uat"
}

func (e *environment) IsPro() bool {
	return e.value == "pro"
}

func init() {
	envStr := os.Getenv("GOENV")
	switch envStr {
	case "dev":
		active = dev
	case "sit":
		active = sit
	case "uat":
		active = uat
	case "pro":
		active = pro
	default:
		active = dev
	}
}

// Active 当前配置的env
func Active() Environment {
	return active
}
