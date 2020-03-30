package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func pushToGithub(data, token string) error {
	if data == "" {
		return errors.New("params error")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	c := "feat: add gocn news, date : " + time.Now().Format("2006-01-02")
	sha := ""
	content := &github.RepositoryContentFileOptions{
		Message: &c,
		SHA:     &sha,
		Committer: &github.CommitAuthor{
			Name:  github.String("lubanproj"),
			Email: github.String("1811704358@qq.com"),
			Login: github.String("lubanproj"),
		},
		Author: &github.CommitAuthor{
			Name:  github.String("lubanproj"),
			Email: github.String("1811704358@qq.com"),
			Login: github.String("lubanproj"),
		},
		Branch: github.String("master"),
	}
	op := &github.RepositoryContentGetOptions{}

	repo, _, _, er := client.Repositories.GetContents(ctx, "lubanproj", "go_read", "README.md", op)
	if er != nil || repo == nil {
		fmt.Println("get github repositories error, date: ", time.Now())
		return er
	}

	content.SHA = repo.SHA
	decodeBytes, err := base64.StdEncoding.DecodeString(*repo.Content)
	if err != nil {
		fmt.Println("decode repo error, ",err)
		return err
	}


	oldContentList := strings.Split(string(decodeBytes), "<br>")

	content.Content = []byte(oldContentList[0] + data + "<br>")

	_, _, err = client.Repositories.UpdateFile(ctx, "lubanproj", "go_read", "README.md", content)

	if err != nil {
		println(err)
		return err
	}

	return nil
}