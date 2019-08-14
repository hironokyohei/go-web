package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatarURL AvatarインスタンスがアバターのURLを返すことができない
// 場合に発生するエラーです。
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatar ユーザのプロフィール画像を表す型です。
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返します。
	// 問題が発生した場合にはエラーを返します。特に、URLを取得できなかった。
	// 場合にはErrorNoAvatarURLを返します。
	GetAvatarURL(ChatUser) (string, error)
}

// TryAvatars slice
type TryAvatars []Avatar

// GetAvatarURL torima
func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// AuthAvatar struct
type AuthAvatar struct{}

// UseAuthAvatar AuthAvatar
var UseAuthAvatar AuthAvatar

// GetAvatarURL comment
func (nouse AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar struct
type GravatarAvatar struct{}

// UseGravatar torima
var UseGravatar GravatarAvatar

// GetAvatarURL torima
func (nouse GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

// FileSystemAvatar stuct
type FileSystemAvatar struct{}

// UseFileSystemAvatar torima
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL torima
func (nouse FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}
