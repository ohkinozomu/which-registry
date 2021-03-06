# which-registry

Returns which registry from the container image name.

# Example

```go
import(
  wr "github.com/ohkinozomu/which-registry"
)

image := "nginx:alpine"
registry, _ := wr.Which(image)

// Docker Hub
fmt.Println(registry)
```

# Supported registries

- Docker Hub
- GitHub Container Registry
- Amazon Elastic Container Registry Public
- Amazon Elastic Container Registry(private)
- Google Artifact Registry
- Google Container Registry
- Quay.io

# License

Apache-2.0

`parsers.go` was copied from https://github.com/kubernetes/kubernetes.
