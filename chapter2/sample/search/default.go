package search

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// init registers the default matcher with the program. - init 将默认匹配器注册到程序里
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
