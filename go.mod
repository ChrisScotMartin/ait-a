module github.com/ChrisScotMartin/ait-a

go 1.16

require (
	github.com/ChrisScotMartin/ait-b v0.0.6
	github.com/ChrisScotMartin/ait-c v0.0.1 // indirect
	github.com/ChrisScotMartin/ait-c/pkgmodc v0.0.0-00010101000000-000000000000
)

replace github.com/ChrisScotMartin/ait-c/pkgmodc => github.com/ChrisScotMartin/ait-c/pkgmodc v0.0.3

replace github.com/ChrisScotMartin/ait-c => github.com/ChrisScotMartin/ait-c v0.0.3
