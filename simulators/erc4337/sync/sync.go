package main

import "github.com/ethereum/hive/hivesim"

func main() {
	suite := hivesim.Suite{
		Name:        "ERC4337-Suite",
		Description: "This test suite performs ERC4337 bundlers p2p tests.",
	}
	// add a plain test (does not run a client)
	suite.Add(hivesim.TestSpec{
		Name:        "p2p-sync-test",
		Description: "This test validates the syncing of bundler nodes",
		Run:         runMyTest,
	})
	// add a client test (starts the client)
	// suite.Add(hivesim.ClientTestSpec{
	//     Name:        "the-test-2",
	//     Description: "This is an example test case.",
	//     Files: map[string]string{"/genesis.json": "genesis.json"},
	//     Run: runMyClientTest,
	// })

	// Run the tests. This waits until all tests of the suite
	// have executed.
	hivesim.MustRunSuite(hivesim.New(), suite)
}

func runMyTest(t *hivesim.T) {
	// write your test code here
}

// func runMyClientTest(t *hivesim.T, c *hivesim.Client) {
//     // write your test code here
// }
