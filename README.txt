
To run the Load Balancer, you need to first run the two servers (go run server.go and go run server.go), then the Load Balancer (go run loadBalancer.go), and finally the client (go run client) by passing two words through the command line.
-----------------------------
This is a simple implementation of a LoadBalancer using RPC in Go.
The procedure receives two words and returns their lengths.