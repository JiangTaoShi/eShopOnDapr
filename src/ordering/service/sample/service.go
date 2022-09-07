package sample

import (
	"github.com/JiangTaoShi/eShopOnDapr/ordering/infrastructure/repositories/sample"
)

type SampleService struct {
	reposity sample.ISampleRepository
}

func New() (SampleService, error) {
	return SampleService{}, nil
}
