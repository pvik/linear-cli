package git

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/rs/zerolog/log"
)

func getCWDRepo() *git.Repository {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to get CWD")
	}

	found := false
	path := cwd

	if !found {
		for _ = range 4 {

			r, err := git.PlainOpen(path)
			if err != nil {
				log.Debug().Err(err).Msg("Unable to open git repo")

				path = filepath.Join(path, "..")
				continue
			}

			return r

		}
	}

	log.Fatal().Str("path", cwd).Msg("Unable to open git repo")

	return nil

}

func getBranchRefName(branchName string) plumbing.ReferenceName {
	branchRefStr := fmt.Sprintf("refs/heads/%s", branchName)
	return plumbing.ReferenceName(branchRefStr)
}

func CheckBranchExists(branchName string) bool {
	r := getCWDRepo()

	branchRef := getBranchRefName(branchName)

	ref, _ := r.Reference(branchRef, true)
	if ref != nil {
		return true
	}

	return false
}

// An example of how to create and remove branches or any other kind of reference.
func CreateBranch(newBranchName string, fromBranch string, switchAfter bool) {
	r := getCWDRepo()

	headRef := &plumbing.Reference{}
	err := fmt.Errorf("")

	if fromBranch == "" {
		headRef, err = r.Head()
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to get HEAD for git repo")
		}
	} else {
		fromBranchRef := getBranchRefName(fromBranch)
		headRef, err = r.Reference(fromBranchRef, true)
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to get referench for source branch for git repo")
		}
	}

	// Create a new plumbing.HashReference object with the name of the branch
	// and the hash from the HEAD. The reference name should be a full reference
	// name and not an abbreviated one, as is used on the git cli.
	//
	// For tags we should use `refs/tags/%s` instead of `refs/heads/%s` used
	// for branches.
	newBranchRefStr := fmt.Sprintf("refs/heads/%s", newBranchName)
	newBranchRef := plumbing.ReferenceName(newBranchRefStr)
	ref := plumbing.NewHashReference(newBranchRef, headRef.Hash())

	// The created reference is saved in the storage,.
	err = r.Storer.SetReference(ref)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to create branch")
	}

	log.Info().Msg(fmt.Sprintf("Created branch: %s -> %s", fromBranch, newBranchName))

	if switchAfter {
		log.Info().Msg(fmt.Sprintf("Checking out branch: %s", newBranchName))

		w, err := r.Worktree()
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to get git worktree")
		}

		err = w.Checkout(&git.CheckoutOptions{
			Branch: newBranchRef,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to checkout branch")
		}
	}
}

func SwitchBranch(branchName string) {
	r := getCWDRepo()

	branchRef := getBranchRefName(branchName)

	log.Info().Msg(fmt.Sprintf("Checking out branch: %s", branchName))

	w, err := r.Worktree()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to get git worktree")
	}

	err = w.Checkout(&git.CheckoutOptions{
		Branch: branchRef,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to checkout branch")
	}

}
