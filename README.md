# which-registry

Returns which registry from the container image name.

# Example

```
import(
  wr "github.com/ohkinozomu/which-registry"
)

image := "nginx:alpine"
registry, _ := wr.Which(image)

// Docker Hub
fmt.Println(registry)
```

# Supported registry

- Docker Hub
- Amazon Elastic Container Registry Public
- Amazon Elastic Container Registry(private)

# License

Apache-2.0

`parsers.go` was copied from https://github.com/kubernetes/kubernetes.
