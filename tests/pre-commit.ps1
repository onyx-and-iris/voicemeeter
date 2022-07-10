Function RunTests {
	$run_int_tests = "go clean -testcache; go test -v ."
	$run_ext_tests = "go clean -testcache; go test -v .\tests\"

	Invoke-Expression $run_ext_tests
	Invoke-Expression $run_int_tests
}

if ($MyInvocation.InvocationName -ne ".") {
    RunTests
}
