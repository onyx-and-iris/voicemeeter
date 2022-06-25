Function RunTests {
    $run_tests = "go clean -testcache; go test ../..."

	Invoke-Expression $run_tests
}

if ($MyInvocation.InvocationName -ne ".") {
    RunTests
}
