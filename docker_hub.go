package which_registry

func isDockerHub(domain string) bool {
	return domain == "docker.io"
}
