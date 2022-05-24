// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://ac1d09f2c3b338c15e3b8a7e5bb1a4054831ac45d761cc488079760b7bde6693cf72b21c992165d0de256f706c1292e043758f5c334360c3216480a42c140abd@127.0.0.1:30303",
	"enode://b9dadb566d66dbe27023c2481385f57c4447b08b505208434c9aa3f5772fb04977c8eaa77d256e080f5a5ae5918c284a6eb0a55c64a683535634850d14a329ac@127.0.0.1:30304",
	"enode://2be9b659d28ef3cd4aa2866d77382c095d872655d2ec30c2cbb7cca0ddb3e32ac575d68b1fc56ee2b965d1a40b1c4adf81abea2dd4af8548d4bf80b4839a1c8d@127.0.0.1:30305",
	"enode://6db15c433f144ba9dcfa7a327c88f29b36a136465636ed7c2a74f020c55307d821343fabf28bd1b8ed49360748d7896ac62ec08db3c5cb451156aa5c834ea957@127.0.0.1:30306",
	"enode://c85269c2a8a734f91c84569cf7c298e88718c21bf69baf35dffce22ba069f84fb65425183f9f755af6571651803d31ec2494f1d35abe03a0f3b17f1a969064a9@127.0.0.1:30307",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// test network.
var TestnetBootnodes = []string{
	"enode://093dab36f391382510c7dd705dacb5d589f32da005ef39a8052f42a928581d676f5a04e4668c0d88ae721f7ccc018eb0af7fd54bfae7b70fb1d261890272884f@159.65.104.177:30301",
}
