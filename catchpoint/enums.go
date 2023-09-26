package catchpoint

type TestType int

const (
	Web         TestType = 0
	Api         TestType = 9
	Transaction TestType = 1
	Traceroute  TestType = 12
	Dns         TestType = 5
	Ping        TestType = 6
	Bgp         TestType = 20
	Ssl         TestType = 18
)
