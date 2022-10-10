package git

//var mockedExitStatus = 0
//var mockedStdout string
//var execCommand = exec.Command
//
//func fakeExecCommand(command string, args ...string) *exec.Cmd {
//	cs := []string{"-test.run=TestExecCommandHelper", "--", command}
//	cs = append(cs, args...)
//	cmd := exec.Command(os.Args[0], cs...)
//	es := strconv.Itoa(mockedExitStatus)
//	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1",
//		"STDOUT=" + mockedStdout,
//		"EXIT_STATUS=" + es}
//	return cmd
//}
//
//func TestPrintDate(t *testing.T) {
//	mockedExitStatus = 1
//	mockedStdout = "feature/TAS-1212-some-feature"
//	execCommand = fakeExecCommand
//	defer func() { execCommand = exec.Command }()
//
//	out, err := GetCurrentBranch()
//	assert.Nil(t, err)
//	assert.Equal(t, mockedStdout, out)
//}
