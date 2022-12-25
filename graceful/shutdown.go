package graceful

func ShutDown(quietly bool) {
	logInfo(quietly, "Tmaic is shutting down")
	closeQueue(quietly)
	closeCache(quietly)
	closeDB(quietly)
	closeMonitor(quietly)
	logInfo(quietly, "Tmaic is shut down")
}
