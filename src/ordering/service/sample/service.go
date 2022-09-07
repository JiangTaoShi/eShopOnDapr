package sample

import (
	"github.com/JiangTaoShi/eShopOnDapr/ordering/infrastructure/repositories/sample"
)

type SampleService struct {
	Reposity sample.ISampleRepository
}

func New() SampleService {
	return SampleService{Reposity: sample.New()}
}
