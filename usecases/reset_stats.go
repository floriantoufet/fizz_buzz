package usecases

// ResetStats implements Usecases interface
func (uc Vanilla) ResetStats() {
	logger := uc.logger.Named("ResetStats")
	logger.Debug("Success")

	uc.stats.ResetStats()
}
