package main

import (
	"fmt"

	"github.com/ethereum/hive/hivesim"
)

func main() {
	// suite := hivesim.Suite{
	// 	Name:        "ERC4337-Suite",
	// 	Description: "This test suite performs ERC4337 bundlers p2p tests.",
	// }
	// // add a plain test (does not run a client)
	// suite.Add(hivesim.TestSpec{
	// 	Name:        "p2p-sync-test",
	// 	Description: "This test validates the syncing of bundler nodes",
	// 	Run:         runSyncTest,
	// 	// Parameters: hivesim.Params{
	// 	// 	"HIVE_ERC4337_ETH1_RPC_ADDRS": "",
	// 	// },
	// })
	// // add a client test (starts the client)
	// // suite.Add(hivesim.ClientTestSpec{
	// //     Name:        "the-test-2",
	// //     Description: "This is an example test case.",
	// //     Files: map[string]string{"/genesis.json": "genesis.json"},
	// //     Run: runMyClientTest,
	// // })

	// // Run the tests. This waits until all tests of the suite
	// // have executed.
	// hivesim.MustRunSuite(hivesim.New(), suite)
	userOps, err := loadUserOps()
	if err != nil {
		panic("Failed to load userops from json")
	}
	fmt.Println("UserOps", userOps)
}

func runSyncTest(t *hivesim.T) {
	// load the userops from the json
	userOps, err := loadUserOps()
	if err != nil {
		panic("Failed to load userops from json")
	}
	fmt.Println("UserOps", userOps)
}

// func runMyClientTest(t *hivesim.T, c *hivesim.Client) {
//     // write your test code here
// }
