//go:build !ignore_autogenerated

// Copyright (c) 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package cloudinit

import (
	"encoding/json"
	"github.com/vmware-tanzu/vm-operator/api/v1alpha2/common"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudConfig) DeepCopyInto(out *CloudConfig) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RunCmd != nil {
		in, out := &in.RunCmd, &out.RunCmd
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.WriteFiles != nil {
		in, out := &in.WriteFiles, &out.WriteFiles
		*out = make([]WriteFile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SSHPwdAuth != nil {
		in, out := &in.SSHPwdAuth, &out.SSHPwdAuth
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudConfig.
func (in *CloudConfig) DeepCopy() *CloudConfig {
	if in == nil {
		return nil
	}
	out := new(CloudConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	if in.CreateGroups != nil {
		in, out := &in.CreateGroups, &out.CreateGroups
		*out = new(bool)
		**out = **in
	}
	if in.ExpireDate != nil {
		in, out := &in.ExpireDate, &out.ExpireDate
		*out = new(string)
		**out = **in
	}
	if in.Gecos != nil {
		in, out := &in.Gecos, &out.Gecos
		*out = new(string)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HashedPasswd != nil {
		in, out := &in.HashedPasswd, &out.HashedPasswd
		*out = new(common.SecretKeySelector)
		**out = **in
	}
	if in.Homedir != nil {
		in, out := &in.Homedir, &out.Homedir
		*out = new(string)
		**out = **in
	}
	if in.Inactive != nil {
		in, out := &in.Inactive, &out.Inactive
		*out = new(int32)
		**out = **in
	}
	if in.LockPasswd != nil {
		in, out := &in.LockPasswd, &out.LockPasswd
		*out = new(bool)
		**out = **in
	}
	if in.NoCreateHome != nil {
		in, out := &in.NoCreateHome, &out.NoCreateHome
		*out = new(bool)
		**out = **in
	}
	if in.NoLogInit != nil {
		in, out := &in.NoLogInit, &out.NoLogInit
		*out = new(bool)
		**out = **in
	}
	if in.NoUserGroup != nil {
		in, out := &in.NoUserGroup, &out.NoUserGroup
		*out = new(bool)
		**out = **in
	}
	if in.Passwd != nil {
		in, out := &in.Passwd, &out.Passwd
		*out = new(common.SecretKeySelector)
		**out = **in
	}
	if in.PrimaryGroup != nil {
		in, out := &in.PrimaryGroup, &out.PrimaryGroup
		*out = new(string)
		**out = **in
	}
	if in.SELinuxUser != nil {
		in, out := &in.SELinuxUser, &out.SELinuxUser
		*out = new(string)
		**out = **in
	}
	if in.Shell != nil {
		in, out := &in.Shell, &out.Shell
		*out = new(string)
		**out = **in
	}
	if in.SnapUser != nil {
		in, out := &in.SnapUser, &out.SnapUser
		*out = new(string)
		**out = **in
	}
	if in.SSHAuthorizedKeys != nil {
		in, out := &in.SSHAuthorizedKeys, &out.SSHAuthorizedKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SSHImportID != nil {
		in, out := &in.SSHImportID, &out.SSHImportID
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SSHRedirectUser != nil {
		in, out := &in.SSHRedirectUser, &out.SSHRedirectUser
		*out = new(bool)
		**out = **in
	}
	if in.Sudo != nil {
		in, out := &in.Sudo, &out.Sudo
		*out = new(string)
		**out = **in
	}
	if in.System != nil {
		in, out := &in.System, &out.System
		*out = new(bool)
		**out = **in
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WriteFile) DeepCopyInto(out *WriteFile) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WriteFile.
func (in *WriteFile) DeepCopy() *WriteFile {
	if in == nil {
		return nil
	}
	out := new(WriteFile)
	in.DeepCopyInto(out)
	return out
}
