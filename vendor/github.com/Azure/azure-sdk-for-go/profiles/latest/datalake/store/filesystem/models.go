// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: 2014fbbf031942474ad27a5a66dffaed5347f3fb

package filesystem

import original "github.com/Azure/azure-sdk-for-go/services/datalake/store/2016-11-01/filesystem"

const (
	DefaultAdlsFileSystemDNSSuffix = original.DefaultAdlsFileSystemDNSSuffix
)

type BaseClient = original.BaseClient
type Client = original.Client
type AppendModeType = original.AppendModeType

const (
	Autocreate AppendModeType = original.Autocreate
)

type Exception = original.Exception

const (
	ExceptionAccessControlException        Exception = original.ExceptionAccessControlException
	ExceptionAdlsRemoteException           Exception = original.ExceptionAdlsRemoteException
	ExceptionBadOffsetException            Exception = original.ExceptionBadOffsetException
	ExceptionFileAlreadyExistsException    Exception = original.ExceptionFileAlreadyExistsException
	ExceptionFileNotFoundException         Exception = original.ExceptionFileNotFoundException
	ExceptionIllegalArgumentException      Exception = original.ExceptionIllegalArgumentException
	ExceptionIOException                   Exception = original.ExceptionIOException
	ExceptionRuntimeException              Exception = original.ExceptionRuntimeException
	ExceptionSecurityException             Exception = original.ExceptionSecurityException
	ExceptionThrottledException            Exception = original.ExceptionThrottledException
	ExceptionUnsupportedOperationException Exception = original.ExceptionUnsupportedOperationException
)

type ExpiryOptionType = original.ExpiryOptionType

const (
	Absolute               ExpiryOptionType = original.Absolute
	NeverExpire            ExpiryOptionType = original.NeverExpire
	RelativeToCreationDate ExpiryOptionType = original.RelativeToCreationDate
	RelativeToNow          ExpiryOptionType = original.RelativeToNow
)

type FileType = original.FileType

const (
	DIRECTORY FileType = original.DIRECTORY
	FILE      FileType = original.FILE
)

type SyncFlag = original.SyncFlag

const (
	CLOSE    SyncFlag = original.CLOSE
	DATA     SyncFlag = original.DATA
	METADATA SyncFlag = original.METADATA
)

type ACLStatus = original.ACLStatus
type ACLStatusResult = original.ACLStatusResult
type AdlsAccessControlException = original.AdlsAccessControlException
type AdlsBadOffsetException = original.AdlsBadOffsetException
type AdlsError = original.AdlsError
type AdlsFileAlreadyExistsException = original.AdlsFileAlreadyExistsException
type AdlsFileNotFoundException = original.AdlsFileNotFoundException
type AdlsIllegalArgumentException = original.AdlsIllegalArgumentException
type AdlsIOException = original.AdlsIOException
type BasicAdlsRemoteException = original.BasicAdlsRemoteException
type AdlsRemoteException = original.AdlsRemoteException
type AdlsRuntimeException = original.AdlsRuntimeException
type AdlsSecurityException = original.AdlsSecurityException
type AdlsThrottledException = original.AdlsThrottledException
type AdlsUnsupportedOperationException = original.AdlsUnsupportedOperationException
type ContentSummary = original.ContentSummary
type ContentSummaryResult = original.ContentSummaryResult
type FileOperationResult = original.FileOperationResult
type FileStatuses = original.FileStatuses
type FileStatusesResult = original.FileStatusesResult
type FileStatusProperties = original.FileStatusProperties
type FileStatusResult = original.FileStatusResult
type ReadCloser = original.ReadCloser

func New() BaseClient {
	return original.New()
}
func NewWithoutDefaults(adlsFileSystemDNSSuffix string) BaseClient {
	return original.NewWithoutDefaults(adlsFileSystemDNSSuffix)
}
func NewClient() Client {
	return original.NewClient()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
