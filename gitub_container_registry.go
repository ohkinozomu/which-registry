package which_registry

func isGitHubContainerRegistry(domain string) bool {
	return domain == "ghcr.io"
}
