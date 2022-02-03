package which_registry

type Registry int

const (
	UNKNOWN Registry = iota
	DOCKER_HUB
	ECR_PUBLIC
	ECR_PRIVATE
	GOOGLE_ARTIFACT_REGISTRY
	GOOGLE_CONTAINER_REGISTRY
)

func (r Registry) String() string {
	switch r {
	case UNKNOWN:
		return "Unknown"
	case DOCKER_HUB:
		return "Docker Hub"
	case ECR_PUBLIC:
		return "Amazon Elastic Container Registry Public"
	case ECR_PRIVATE:
		return "Amazon Elastic Container Registry(private)"
	case GOOGLE_ARTIFACT_REGISTRY:
		return "Google Artifact Registry"
	case GOOGLE_CONTAINER_REGISTRY:
		return "Google Container Registry"
	default:
		return "Unknown"
	}
}
