package main

//
// func BenchmarkSetupCoinTossersConc1MaxTosses10000000(b *testing.B) {
// 	doBenchmarkCoinTossers(b, Config{concurrentThreads: 1, maxTosses: 10000000})
// }
// func BenchmarkSetupCoinTossersConc5MaxTosses10000000(b *testing.B) {
// 	doBenchmarkCoinTossers(b, Config{concurrentThreads: 5, maxTosses: 10000000})
// }
//
// func doBenchmarkCoinTossers(b *testing.B, conf Config) {
// 	for n := 0; n < b.N; n++ {
// 		tF := createTestFixture(conf)
// 		tF.cleanup()
// 		stop = false
// 		cT.inARowCounter = make(map[int]int)
// 		cT.count = 0
// 		tF.app.start()
// 	}
// }
