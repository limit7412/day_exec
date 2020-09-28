package repository

import (
	"fmt"
	"os/exec"
)

func Exec(cmd string) error {

	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to run cmd. output %s", string(out)))
		return err
	}

	fmt.Println(string(out))

	return nil
}
