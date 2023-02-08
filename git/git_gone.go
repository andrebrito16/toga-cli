package git

import (
	"fmt"

	"github.com/andrebrito16/toga-cli/evaluator"
	"github.com/urfave/cli/v2"
)

func GitGone(c *cli.Context) error {

	err := evaluator.Run("git fetch -p && git branch -vv | grep 'origin/.*: gone]' | awk '{print \\$1}' | xargs git branch -D", []string{""})

	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	fmt.Println("All local branches that not exist on remote were removed!")

	return nil
}
