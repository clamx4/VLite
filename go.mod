module github.com/xiaokangwang/VLite

go 1.13

require (
	github.com/FlowerWrong/netstack v0.0.0-20191009041010-964b57cc9e6e
	github.com/FlowerWrong/water v0.0.0-20180301012659-01a4eaa1f6f2
	github.com/davecgh/go-spew v1.1.1
	github.com/google/gopacket v1.1.17
	github.com/juju/ratelimit v1.0.1
	github.com/lunixbochs/struc v0.0.0-20190916212049-a5c72983bc42
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pion/dtls v1.5.4 // indirect
	github.com/pion/dtls/v2 v2.0.0-rc.7
	github.com/pion/logging v0.2.2
	github.com/pion/sctp v1.7.6
	github.com/stretchr/testify v1.5.1
	github.com/txthinking/runnergroup v0.0.0-20200327135940-540a793bb997 // indirect
	github.com/txthinking/socks5 v0.0.0-20200327133705-caf148ab5e9d
	github.com/txthinking/x v0.0.0-20200330144832-5ad2416896a9 // indirect
	github.com/xiaokangwang/water v0.0.0-20180524022535-3e9597577724
	github.com/xtaci/smux v1.5.12
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
)

replace github.com/FlowerWrong/netstack => ./vendor/github.com/FlowerWrong/netstack

replace github.com/FlowerWrong/water => ./vendor/github.com/FlowerWrong/water
