# <span style="font-size: 2em;">**LoadBalancer**⚖️💻</span>

This is a simple implementation of a LoadBalancer that balance the workload between 2 different server using RPC in Go.

**How to Use**

The procedure receives two words and returns their lengths.
To run the Load Balancer, you need to first run the two servers (```go run server.go``` and ```go run secondServer.go```), then the Load Balancer (```go run loadBalancer.go```), and finally the client (```go run client.go```) by passing two words through the command line.

