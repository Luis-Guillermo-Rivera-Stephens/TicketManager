package datatypes

var StatusHash map[string]int = map[string]int{
	"New Unassigned":   1,
	"New Assigned":     2,
	"In Progress":      3,
	"On Hold":          4,
	"Testing":          5,
	"Closed":           6,
	"Canceled":         7,
	"Testing Declined": 8,
}

var StatusSlice = []string{
	"None",
	"New Unassigned",
	"New Assigned",
	"In Progress",
	"On Hold",
	"Testing",
	"Closed",
	"Canceled",
	"Testing Declined",
}
