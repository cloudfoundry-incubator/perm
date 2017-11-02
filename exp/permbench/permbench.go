// +build ignore

// permbench is a command which benchmarks the performance of a Perm service.
//
// Even though it does not use the "go test" program to run the benchmarks it
// emits the results in the same format so that tools like benchcmp can be used
// to analyse the results.
//
// The existing -test.* flags arw respected by this program. However, when you
// are writing benchmarks you should log errors rather than reporting them
// using the *testing.B reporter.
package main

import (
	"context"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	permpb "code.cloudfoundry.org/perm/protos"
)

var (
	addr        = flag.String("addr", "localhost:6283", "address of perm service")
	caPath      = flag.String("ca", "", "path to CA certificate (defaults to insecure)")
	serviceName = flag.String("serviceName", "", "server name override for TLS connection (defaults to no override)")

	ctx = context.Background()
)

func main() {
	flag.Parse()

	benchmark("RoleCreationDeletion", roleCreationDeletion)
}

func benchmark(name string, fn func(*testing.B, permpb.RoleServiceClient)) {
	client := dialClient()

	res := testing.Benchmark(func(b *testing.B) {
		fn(b, client)
	})

	fmt.Println("Benchmark"+name, res.String())
}

func dialClient() permpb.RoleServiceClient {
	opts := []grpc.DialOption{
		// Do not wait for connection on the first benchmark iteration.
		grpc.WithBlock(),
	}

	if *caPath != "" {
		pem, err := ioutil.ReadFile(*caPath)
		if err != nil {
			log.Fatalf("could not read certificate file %q: %q", *caPath, err)
		}

		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(pem)
		creds := credentials.NewClientTLSFromCert(pool, *serviceName)
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// We want to timeout the grpc.Dial if the server is not running. Due to the
	// grpc.WithBlock option above this call will block forever otherwise.
	dialCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, err := grpc.DialContext(dialCtx, *addr, opts...)
	if err != nil {
		log.Fatalf("failed to dial perm service at %q: %q", *addr, err)
	}

	return permpb.NewRoleServiceClient(conn)
}

func roleCreationDeletion(b *testing.B, client permpb.RoleServiceClient) {
	name := "creation-role"

	for i := 0; i < b.N; i++ {
		_, err := client.CreateRole(ctx, &permpb.CreateRoleRequest{
			Name: name,
		})
		if err != nil {
			log.Fatalf("failed to create role: %q", err)
		}

		_, err = client.DeleteRole(ctx, &permpb.DeleteRoleRequest{
			Name: name,
		})
		if err != nil {
			log.Fatalf("failed to delete role: %q", err)
		}
	}
}
