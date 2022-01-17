package search

import (
	"log"
	"sync"
)

// A map of registered matchers for searching.
var matchers = make(map[string]Matcher)

// Run performs the search logic. - 执行搜索逻辑
func Run(searchTerm string) {
	// Retrieve the list of feeds to search through. - 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err) // Fatal() 接受错误并终止程序
	}

	// Create an unbuffered channel to receive match results to display.
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	for _, feed := range feeds {
		// Retrieve a matcher for the search.
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 闭包 - 每个协程都变化的变量使用传参的方式入参，直接使用闭包会使所有协程都用同一个变量值（可能是最后一个变量值）
		// 闭包可以使得匿名函数直接访问到哪些没有作为参数传入的变量；匿名函数并没有拿到这些变量的副本，
		//而是直接访问外层函数作用域中声明的这些变量本身
		// 因为matcher和feed变量每次调用的值不相同，所以并没有使用闭包的方式访问
		// Launch the goroutine to perform the search.
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed) // 匿名函数，func(形参){}(实参) 定义后直接调用
	}
	// 闭包
	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program.
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	Display(results)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
