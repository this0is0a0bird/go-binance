package binance

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type currentGetAllServiceTestSuite struct {
	baseTestSuite
}

func TestGetAllCurrencyService(t *testing.T) {
	suite.Run(t, new(currentGetAllServiceTestSuite))
}

func (s *currentGetAllServiceTestSuite)TestGetAllCurrency()  {
	res, err := s.client.NewGetAllCurrencyService().Do(newContext())
	r := s.r()
	fmt.Println(res, err, r)
}