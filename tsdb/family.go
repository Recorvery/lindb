package tsdb

//go:generate mockgen -source=./family.go -destination=./family_mock.go -package=tsdb -self_package=github.com/eleme/lindb/tsdb

type DataFamily interface {
	Scan(scanContext *ScanContext) Scanner
}
