module test-iris

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200323165209-0ec3e9974c59
	golang.org/x/mod => github.com/golang/mod v0.2.0
	golang.org/x/net => github.com/golang/net v0.0.0-20200324143707-d3edc9973b7e
	golang.org/x/sync => github.com/golang/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200323222414-85ca7c5b95cd
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20200325010219-a49f79bcc224
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191204190536-9bdfabe68543
)

require (
	github.com/kataras/iris/v12 v12.1.8
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
)
