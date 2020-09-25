// Code generated by mockery v2.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	bucket "github.com/FleekHQ/space-daemon/core/textile/bucket"

	io "io"

	mock "github.com/stretchr/testify/mock"

	path "github.com/ipfs/interface-go-ipfs-core/path"

	thread "github.com/textileio/go-threads/core/thread"
)

// Bucket is an autogenerated mock type for the Bucket type
type Bucket struct {
	mock.Mock
}

// CreateDirectory provides a mock function with given fields: ctx, _a1
func (_m *Bucket) CreateDirectory(ctx context.Context, _a1 string) (path.Resolved, path.Path, error) {
	ret := _m.Called(ctx, _a1)

	var r0 path.Resolved
	if rf, ok := ret.Get(0).(func(context.Context, string) path.Resolved); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(path.Resolved)
		}
	}

	var r1 path.Path
	if rf, ok := ret.Get(1).(func(context.Context, string) path.Path); ok {
		r1 = rf(ctx, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(path.Path)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteDirOrFile provides a mock function with given fields: ctx, _a1
func (_m *Bucket) DeleteDirOrFile(ctx context.Context, _a1 string) (path.Resolved, error) {
	ret := _m.Called(ctx, _a1)

	var r0 path.Resolved
	if rf, ok := ret.Get(0).(func(context.Context, string) path.Resolved); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(path.Resolved)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DirExists provides a mock function with given fields: ctx, _a1
func (_m *Bucket) DirExists(ctx context.Context, _a1 string) (bool, error) {
	ret := _m.Called(ctx, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileExists provides a mock function with given fields: ctx, _a1
func (_m *Bucket) FileExists(ctx context.Context, _a1 string) (bool, error) {
	ret := _m.Called(ctx, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClient provides a mock function with given fields:
func (_m *Bucket) GetClient() bucket.BucketsClient {
	ret := _m.Called()

	var r0 bucket.BucketsClient
	if rf, ok := ret.Get(0).(func() bucket.BucketsClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bucket.BucketsClient)
		}
	}

	return r0
}

// GetContext provides a mock function with given fields: ctx
func (_m *Bucket) GetContext(ctx context.Context) (context.Context, *thread.ID, error) {
	ret := _m.Called(ctx)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 *thread.ID
	if rf, ok := ret.Get(1).(func(context.Context) *thread.ID); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*thread.ID)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetData provides a mock function with given fields:
func (_m *Bucket) GetData() bucket.BucketData {
	ret := _m.Called()

	var r0 bucket.BucketData
	if rf, ok := ret.Get(0).(func() bucket.BucketData); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bucket.BucketData)
	}

	return r0
}

// GetFile provides a mock function with given fields: ctx, _a1, w
func (_m *Bucket) GetFile(ctx context.Context, _a1 string, w io.Writer) error {
	ret := _m.Called(ctx, _a1, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Writer) error); ok {
		r0 = rf(ctx, _a1, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetThreadID provides a mock function with given fields: ctx
func (_m *Bucket) GetThreadID(ctx context.Context) (*thread.ID, error) {
	ret := _m.Called(ctx)

	var r0 *thread.ID
	if rf, ok := ret.Get(0).(func(context.Context) *thread.ID); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*thread.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Key provides a mock function with given fields:
func (_m *Bucket) Key() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ListDirectory provides a mock function with given fields: ctx, _a1
func (_m *Bucket) ListDirectory(ctx context.Context, _a1 string) (*bucket.DirEntries, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *bucket.DirEntries
	if rf, ok := ret.Get(0).(func(context.Context, string) *bucket.DirEntries); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bucket.DirEntries)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Slug provides a mock function with given fields:
func (_m *Bucket) Slug() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UploadFile provides a mock function with given fields: ctx, _a1, reader
func (_m *Bucket) UploadFile(ctx context.Context, _a1 string, reader io.Reader) (path.Resolved, path.Path, error) {
	ret := _m.Called(ctx, _a1, reader)

	var r0 path.Resolved
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader) path.Resolved); ok {
		r0 = rf(ctx, _a1, reader)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(path.Resolved)
		}
	}

	var r1 path.Path
	if rf, ok := ret.Get(1).(func(context.Context, string, io.Reader) path.Path); ok {
		r1 = rf(ctx, _a1, reader)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(path.Path)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, io.Reader) error); ok {
		r2 = rf(ctx, _a1, reader)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
