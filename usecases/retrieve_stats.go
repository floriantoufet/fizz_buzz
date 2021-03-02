package usecases

import (
	"fizzbuzz/domains"
)

// RetrieveStats implements Usecases interface
func (uc Vanilla) RetrieveStats() ([]domains.FizzBuzz, uint) {
	logger := uc.logger.Named("RetrieveStats")
	logger.Debug("Success")

	return uc.stats.RetrieveMostFrequentFizzBuzzRequest()
}
