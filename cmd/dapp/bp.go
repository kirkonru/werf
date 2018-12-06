package main

import (
	"fmt"

	"github.com/flant/dapp/cmd/dapp/docker_authorizer"
	"github.com/flant/dapp/pkg/build"
	"github.com/flant/dapp/pkg/dapp"
	"github.com/flant/dapp/pkg/docker"
	"github.com/flant/dapp/pkg/lock"
	"github.com/flant/dapp/pkg/ssh_agent"
	"github.com/flant/dapp/pkg/true_git"
	"github.com/spf13/cobra"
)

var bpCmdData struct {
	PullUsername     string
	PullPassword     string
	PushUsername     string
	PushPassword     string
	RegistryUsername string
	RegistryPassword string

	Tag        []string
	TagBranch  bool
	TagBuildID bool
	TagCI      bool
	TagCommit  bool
}

func newBPCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "bp",
		RunE: func(cmd *cobra.Command, args []string) error {
			if bpCmdData.PullUsername == "" {
				bpCmdData.PullUsername = bpCmdData.RegistryUsername
			}
			if bpCmdData.PullPassword == "" {
				bpCmdData.PullPassword = bpCmdData.RegistryPassword
			}
			if bpCmdData.PushUsername == "" {
				bpCmdData.PushUsername = bpCmdData.RegistryUsername
			}
			if bpCmdData.PushPassword == "" {
				bpCmdData.PushPassword = bpCmdData.RegistryPassword
			}

			return runBP()
		},
	}

	cmd.PersistentFlags().StringVarP(&pushCmdData.Repo, "repo", "", "", "Docker repository name to push images to. CI_REGISTRY_IMAGE will be used by default if available.")
	cmd.PersistentFlags().BoolVarP(&pushCmdData.WithStages, "with-stages", "", false, "Push images with stages cache")

	cmd.PersistentFlags().StringVarP(&bpCmdData.PullUsername, "pull-username", "", "", "Docker registry username to authorize pull of base images")
	cmd.PersistentFlags().StringVarP(&bpCmdData.PullPassword, "pull-password", "", "", "Docker registry password to authorize pull of base images")
	cmd.PersistentFlags().StringVarP(&bpCmdData.PushUsername, "push-username", "", "", "Docker registry username to authorize push to the docker repo")
	cmd.PersistentFlags().StringVarP(&bpCmdData.PushPassword, "push-password", "", "", "Docker registry password to authorize push to the docker repo")
	cmd.PersistentFlags().StringVarP(&bpCmdData.RegistryUsername, "registry-username", "", "", "Docker registry username to authorize pull of base images and push to the docker repo")
	cmd.PersistentFlags().StringVarP(&bpCmdData.RegistryUsername, "registry-password", "", "", "Docker registry password to authorize pull of base images and push to the docker repo")

	cmd.PersistentFlags().StringArrayVarP(&bpCmdData.Tag, "tag", "", []string{}, "Add tag (can be used one or more times)")
	cmd.PersistentFlags().BoolVarP(&bpCmdData.TagBranch, "tag-branch", "", false, "Tag by git branch")
	cmd.PersistentFlags().BoolVarP(&bpCmdData.TagBuildID, "tag-build-id", "", false, "Tag by CI build id")
	cmd.PersistentFlags().BoolVarP(&bpCmdData.TagCI, "tag-ci", "", false, "Tag by CI branch and tag")
	cmd.PersistentFlags().BoolVarP(&bpCmdData.TagCommit, "tag-commit", "", false, "Tag by git commit")

	return cmd
}

func runBP() error {
	if err := dapp.Init(rootCmdData.TmpDir, rootCmdData.HomeDir); err != nil {
		return fmt.Errorf("initialization error: %s", err)
	}

	if err := lock.Init(); err != nil {
		return err
	}

	if err := true_git.Init(); err != nil {
		return err
	}

	projectDir, err := getProjectDir()
	if err != nil {
		return fmt.Errorf("getting project dir failed: %s", err)
	}

	projectName, err := getProjectName(projectDir)
	if err != nil {
		return fmt.Errorf("getting project name failed: %s", err)
	}

	projectBuildDir, err := getProjectBuildDir(projectName)
	if err != nil {
		return fmt.Errorf("getting project build dir failed: %s", err)
	}

	projectTmpDir, err := getProjectTmpDir()
	if err != nil {
		return fmt.Errorf("getting project tmp dir failed: %s", err)
	}

	dappfile, err := parseDappfile(projectDir)
	if err != nil {
		return fmt.Errorf("dappfile parsing failed: %s", err)
	}

	repo, err := getRequiredRepoName(projectName, pushCmdData.Repo)
	if err != nil {
		return err
	}

	dockerAuthorizer, err := docker_authorizer.GetBPDockerAuthorizer(projectTmpDir, bpCmdData.PullUsername, bpCmdData.PullPassword, bpCmdData.PushUsername, bpCmdData.PushPassword, repo)
	if err != nil {
		return err
	}

	if err := docker.Init(dockerAuthorizer.HostDockerConfigDir); err != nil {
		return err
	}

	if err := ssh_agent.Init(rootCmdData.SSHKeys); err != nil {
		return fmt.Errorf("cannot initialize ssh-agent: %s", err)
	}

	opts, err := getPushOptions(projectDir, pushCmdData.Tag, pushCmdData.TagBranch, pushCmdData.TagCommit, pushCmdData.TagBuildID, pushCmdData.TagCI, pushCmdData.WithStages)
	if err != nil {
		return err
	}

	c := build.NewConveyor(dappfile, projectDir, projectName, projectBuildDir, projectTmpDir, ssh_agent.SSHAuthSock, dockerAuthorizer)
	if err = c.BP(repo, opts); err != nil {
		return err
	}

	return nil
}
