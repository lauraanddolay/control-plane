// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal/model"

// PackageRepository is an autogenerated mock type for the PackageRepository type
type PackageRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, item
func (_m *PackageRepository) Create(ctx context.Context, item *model.Package) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Package) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, tenant, id
func (_m *PackageRepository) Delete(ctx context.Context, tenant string, id string) error {
	ret := _m.Called(ctx, tenant, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenant, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: ctx, tenant, id
func (_m *PackageRepository) Exists(ctx context.Context, tenant string, id string) (bool, error) {
	ret := _m.Called(ctx, tenant, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, tenant, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenant, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, tenant, id
func (_m *PackageRepository) GetByID(ctx context.Context, tenant string, id string) (*model.Package, error) {
	ret := _m.Called(ctx, tenant, id)

	var r0 *model.Package
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.Package); ok {
		r0 = rf(ctx, tenant, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Package)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenant, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByInstanceAuthID provides a mock function with given fields: ctx, tenant, instanceAuthID
func (_m *PackageRepository) GetByInstanceAuthID(ctx context.Context, tenant string, instanceAuthID string) (*model.Package, error) {
	ret := _m.Called(ctx, tenant, instanceAuthID)

	var r0 *model.Package
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.Package); ok {
		r0 = rf(ctx, tenant, instanceAuthID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Package)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenant, instanceAuthID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForApplication provides a mock function with given fields: ctx, tenant, id, applicationID
func (_m *PackageRepository) GetForApplication(ctx context.Context, tenant string, id string, applicationID string) (*model.Package, error) {
	ret := _m.Called(ctx, tenant, id, applicationID)

	var r0 *model.Package
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *model.Package); ok {
		r0 = rf(ctx, tenant, id, applicationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Package)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, tenant, id, applicationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListByApplicationID provides a mock function with given fields: ctx, tenantID, applicationID, pageSize, cursor
func (_m *PackageRepository) ListByApplicationID(ctx context.Context, tenantID string, applicationID string, pageSize int, cursor string) (*model.PackagePage, error) {
	ret := _m.Called(ctx, tenantID, applicationID, pageSize, cursor)

	var r0 *model.PackagePage
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, string) *model.PackagePage); ok {
		r0 = rf(ctx, tenantID, applicationID, pageSize, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PackagePage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, string) error); ok {
		r1 = rf(ctx, tenantID, applicationID, pageSize, cursor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, item
func (_m *PackageRepository) Update(ctx context.Context, item *model.Package) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Package) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
