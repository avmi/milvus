// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	querypb "github.com/milvus-io/milvus/pkg/v2/proto/querypb"
)

// QueryCoordCatalog is an autogenerated mock type for the QueryCoordCatalog type
type QueryCoordCatalog struct {
	mock.Mock
}

type QueryCoordCatalog_Expecter struct {
	mock *mock.Mock
}

func (_m *QueryCoordCatalog) EXPECT() *QueryCoordCatalog_Expecter {
	return &QueryCoordCatalog_Expecter{mock: &_m.Mock}
}

// GetCollectionTargets provides a mock function with given fields: ctx
func (_m *QueryCoordCatalog) GetCollectionTargets(ctx context.Context) (map[int64]*querypb.CollectionTarget, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetCollectionTargets")
	}

	var r0 map[int64]*querypb.CollectionTarget
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[int64]*querypb.CollectionTarget, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[int64]*querypb.CollectionTarget); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int64]*querypb.CollectionTarget)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryCoordCatalog_GetCollectionTargets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionTargets'
type QueryCoordCatalog_GetCollectionTargets_Call struct {
	*mock.Call
}

// GetCollectionTargets is a helper method to define mock.On call
//   - ctx context.Context
func (_e *QueryCoordCatalog_Expecter) GetCollectionTargets(ctx interface{}) *QueryCoordCatalog_GetCollectionTargets_Call {
	return &QueryCoordCatalog_GetCollectionTargets_Call{Call: _e.mock.On("GetCollectionTargets", ctx)}
}

func (_c *QueryCoordCatalog_GetCollectionTargets_Call) Run(run func(ctx context.Context)) *QueryCoordCatalog_GetCollectionTargets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *QueryCoordCatalog_GetCollectionTargets_Call) Return(_a0 map[int64]*querypb.CollectionTarget, _a1 error) *QueryCoordCatalog_GetCollectionTargets_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QueryCoordCatalog_GetCollectionTargets_Call) RunAndReturn(run func(context.Context) (map[int64]*querypb.CollectionTarget, error)) *QueryCoordCatalog_GetCollectionTargets_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollections provides a mock function with given fields: ctx
func (_m *QueryCoordCatalog) GetCollections(ctx context.Context) ([]*querypb.CollectionLoadInfo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetCollections")
	}

	var r0 []*querypb.CollectionLoadInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*querypb.CollectionLoadInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*querypb.CollectionLoadInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*querypb.CollectionLoadInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryCoordCatalog_GetCollections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollections'
type QueryCoordCatalog_GetCollections_Call struct {
	*mock.Call
}

// GetCollections is a helper method to define mock.On call
//   - ctx context.Context
func (_e *QueryCoordCatalog_Expecter) GetCollections(ctx interface{}) *QueryCoordCatalog_GetCollections_Call {
	return &QueryCoordCatalog_GetCollections_Call{Call: _e.mock.On("GetCollections", ctx)}
}

func (_c *QueryCoordCatalog_GetCollections_Call) Run(run func(ctx context.Context)) *QueryCoordCatalog_GetCollections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *QueryCoordCatalog_GetCollections_Call) Return(_a0 []*querypb.CollectionLoadInfo, _a1 error) *QueryCoordCatalog_GetCollections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QueryCoordCatalog_GetCollections_Call) RunAndReturn(run func(context.Context) ([]*querypb.CollectionLoadInfo, error)) *QueryCoordCatalog_GetCollections_Call {
	_c.Call.Return(run)
	return _c
}

// GetPartitions provides a mock function with given fields: ctx, collectionIDs
func (_m *QueryCoordCatalog) GetPartitions(ctx context.Context, collectionIDs []int64) (map[int64][]*querypb.PartitionLoadInfo, error) {
	ret := _m.Called(ctx, collectionIDs)

	if len(ret) == 0 {
		panic("no return value specified for GetPartitions")
	}

	var r0 map[int64][]*querypb.PartitionLoadInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64) (map[int64][]*querypb.PartitionLoadInfo, error)); ok {
		return rf(ctx, collectionIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []int64) map[int64][]*querypb.PartitionLoadInfo); ok {
		r0 = rf(ctx, collectionIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int64][]*querypb.PartitionLoadInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []int64) error); ok {
		r1 = rf(ctx, collectionIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryCoordCatalog_GetPartitions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitions'
type QueryCoordCatalog_GetPartitions_Call struct {
	*mock.Call
}

// GetPartitions is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionIDs []int64
func (_e *QueryCoordCatalog_Expecter) GetPartitions(ctx interface{}, collectionIDs interface{}) *QueryCoordCatalog_GetPartitions_Call {
	return &QueryCoordCatalog_GetPartitions_Call{Call: _e.mock.On("GetPartitions", ctx, collectionIDs)}
}

func (_c *QueryCoordCatalog_GetPartitions_Call) Run(run func(ctx context.Context, collectionIDs []int64)) *QueryCoordCatalog_GetPartitions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]int64))
	})
	return _c
}

func (_c *QueryCoordCatalog_GetPartitions_Call) Return(_a0 map[int64][]*querypb.PartitionLoadInfo, _a1 error) *QueryCoordCatalog_GetPartitions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QueryCoordCatalog_GetPartitions_Call) RunAndReturn(run func(context.Context, []int64) (map[int64][]*querypb.PartitionLoadInfo, error)) *QueryCoordCatalog_GetPartitions_Call {
	_c.Call.Return(run)
	return _c
}

// GetReplicas provides a mock function with given fields: ctx
func (_m *QueryCoordCatalog) GetReplicas(ctx context.Context) ([]*querypb.Replica, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetReplicas")
	}

	var r0 []*querypb.Replica
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*querypb.Replica, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*querypb.Replica); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*querypb.Replica)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryCoordCatalog_GetReplicas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReplicas'
type QueryCoordCatalog_GetReplicas_Call struct {
	*mock.Call
}

// GetReplicas is a helper method to define mock.On call
//   - ctx context.Context
func (_e *QueryCoordCatalog_Expecter) GetReplicas(ctx interface{}) *QueryCoordCatalog_GetReplicas_Call {
	return &QueryCoordCatalog_GetReplicas_Call{Call: _e.mock.On("GetReplicas", ctx)}
}

func (_c *QueryCoordCatalog_GetReplicas_Call) Run(run func(ctx context.Context)) *QueryCoordCatalog_GetReplicas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *QueryCoordCatalog_GetReplicas_Call) Return(_a0 []*querypb.Replica, _a1 error) *QueryCoordCatalog_GetReplicas_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QueryCoordCatalog_GetReplicas_Call) RunAndReturn(run func(context.Context) ([]*querypb.Replica, error)) *QueryCoordCatalog_GetReplicas_Call {
	_c.Call.Return(run)
	return _c
}

// GetResourceGroups provides a mock function with given fields: ctx
func (_m *QueryCoordCatalog) GetResourceGroups(ctx context.Context) ([]*querypb.ResourceGroup, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetResourceGroups")
	}

	var r0 []*querypb.ResourceGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*querypb.ResourceGroup, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*querypb.ResourceGroup); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*querypb.ResourceGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryCoordCatalog_GetResourceGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResourceGroups'
type QueryCoordCatalog_GetResourceGroups_Call struct {
	*mock.Call
}

// GetResourceGroups is a helper method to define mock.On call
//   - ctx context.Context
func (_e *QueryCoordCatalog_Expecter) GetResourceGroups(ctx interface{}) *QueryCoordCatalog_GetResourceGroups_Call {
	return &QueryCoordCatalog_GetResourceGroups_Call{Call: _e.mock.On("GetResourceGroups", ctx)}
}

func (_c *QueryCoordCatalog_GetResourceGroups_Call) Run(run func(ctx context.Context)) *QueryCoordCatalog_GetResourceGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *QueryCoordCatalog_GetResourceGroups_Call) Return(_a0 []*querypb.ResourceGroup, _a1 error) *QueryCoordCatalog_GetResourceGroups_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QueryCoordCatalog_GetResourceGroups_Call) RunAndReturn(run func(context.Context) ([]*querypb.ResourceGroup, error)) *QueryCoordCatalog_GetResourceGroups_Call {
	_c.Call.Return(run)
	return _c
}

// ReleaseCollection provides a mock function with given fields: ctx, collection
func (_m *QueryCoordCatalog) ReleaseCollection(ctx context.Context, collection int64) error {
	ret := _m.Called(ctx, collection)

	if len(ret) == 0 {
		panic("no return value specified for ReleaseCollection")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, collection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_ReleaseCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseCollection'
type QueryCoordCatalog_ReleaseCollection_Call struct {
	*mock.Call
}

// ReleaseCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - collection int64
func (_e *QueryCoordCatalog_Expecter) ReleaseCollection(ctx interface{}, collection interface{}) *QueryCoordCatalog_ReleaseCollection_Call {
	return &QueryCoordCatalog_ReleaseCollection_Call{Call: _e.mock.On("ReleaseCollection", ctx, collection)}
}

func (_c *QueryCoordCatalog_ReleaseCollection_Call) Run(run func(ctx context.Context, collection int64)) *QueryCoordCatalog_ReleaseCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *QueryCoordCatalog_ReleaseCollection_Call) Return(_a0 error) *QueryCoordCatalog_ReleaseCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_ReleaseCollection_Call) RunAndReturn(run func(context.Context, int64) error) *QueryCoordCatalog_ReleaseCollection_Call {
	_c.Call.Return(run)
	return _c
}

// ReleasePartition provides a mock function with given fields: ctx, collection, partitions
func (_m *QueryCoordCatalog) ReleasePartition(ctx context.Context, collection int64, partitions ...int64) error {
	_va := make([]interface{}, len(partitions))
	for _i := range partitions {
		_va[_i] = partitions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, collection)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ReleasePartition")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...int64) error); ok {
		r0 = rf(ctx, collection, partitions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_ReleasePartition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleasePartition'
type QueryCoordCatalog_ReleasePartition_Call struct {
	*mock.Call
}

// ReleasePartition is a helper method to define mock.On call
//   - ctx context.Context
//   - collection int64
//   - partitions ...int64
func (_e *QueryCoordCatalog_Expecter) ReleasePartition(ctx interface{}, collection interface{}, partitions ...interface{}) *QueryCoordCatalog_ReleasePartition_Call {
	return &QueryCoordCatalog_ReleasePartition_Call{Call: _e.mock.On("ReleasePartition",
		append([]interface{}{ctx, collection}, partitions...)...)}
}

func (_c *QueryCoordCatalog_ReleasePartition_Call) Run(run func(ctx context.Context, collection int64, partitions ...int64)) *QueryCoordCatalog_ReleasePartition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int64, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(int64)
			}
		}
		run(args[0].(context.Context), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *QueryCoordCatalog_ReleasePartition_Call) Return(_a0 error) *QueryCoordCatalog_ReleasePartition_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_ReleasePartition_Call) RunAndReturn(run func(context.Context, int64, ...int64) error) *QueryCoordCatalog_ReleasePartition_Call {
	_c.Call.Return(run)
	return _c
}

// ReleaseReplica provides a mock function with given fields: ctx, collection, replicas
func (_m *QueryCoordCatalog) ReleaseReplica(ctx context.Context, collection int64, replicas ...int64) error {
	_va := make([]interface{}, len(replicas))
	for _i := range replicas {
		_va[_i] = replicas[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, collection)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ReleaseReplica")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...int64) error); ok {
		r0 = rf(ctx, collection, replicas...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_ReleaseReplica_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseReplica'
type QueryCoordCatalog_ReleaseReplica_Call struct {
	*mock.Call
}

// ReleaseReplica is a helper method to define mock.On call
//   - ctx context.Context
//   - collection int64
//   - replicas ...int64
func (_e *QueryCoordCatalog_Expecter) ReleaseReplica(ctx interface{}, collection interface{}, replicas ...interface{}) *QueryCoordCatalog_ReleaseReplica_Call {
	return &QueryCoordCatalog_ReleaseReplica_Call{Call: _e.mock.On("ReleaseReplica",
		append([]interface{}{ctx, collection}, replicas...)...)}
}

func (_c *QueryCoordCatalog_ReleaseReplica_Call) Run(run func(ctx context.Context, collection int64, replicas ...int64)) *QueryCoordCatalog_ReleaseReplica_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int64, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(int64)
			}
		}
		run(args[0].(context.Context), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *QueryCoordCatalog_ReleaseReplica_Call) Return(_a0 error) *QueryCoordCatalog_ReleaseReplica_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_ReleaseReplica_Call) RunAndReturn(run func(context.Context, int64, ...int64) error) *QueryCoordCatalog_ReleaseReplica_Call {
	_c.Call.Return(run)
	return _c
}

// ReleaseReplicas provides a mock function with given fields: ctx, collectionID
func (_m *QueryCoordCatalog) ReleaseReplicas(ctx context.Context, collectionID int64) error {
	ret := _m.Called(ctx, collectionID)

	if len(ret) == 0 {
		panic("no return value specified for ReleaseReplicas")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, collectionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_ReleaseReplicas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseReplicas'
type QueryCoordCatalog_ReleaseReplicas_Call struct {
	*mock.Call
}

// ReleaseReplicas is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
func (_e *QueryCoordCatalog_Expecter) ReleaseReplicas(ctx interface{}, collectionID interface{}) *QueryCoordCatalog_ReleaseReplicas_Call {
	return &QueryCoordCatalog_ReleaseReplicas_Call{Call: _e.mock.On("ReleaseReplicas", ctx, collectionID)}
}

func (_c *QueryCoordCatalog_ReleaseReplicas_Call) Run(run func(ctx context.Context, collectionID int64)) *QueryCoordCatalog_ReleaseReplicas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *QueryCoordCatalog_ReleaseReplicas_Call) Return(_a0 error) *QueryCoordCatalog_ReleaseReplicas_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_ReleaseReplicas_Call) RunAndReturn(run func(context.Context, int64) error) *QueryCoordCatalog_ReleaseReplicas_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveCollectionTarget provides a mock function with given fields: ctx, collectionID
func (_m *QueryCoordCatalog) RemoveCollectionTarget(ctx context.Context, collectionID int64) error {
	ret := _m.Called(ctx, collectionID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveCollectionTarget")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, collectionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_RemoveCollectionTarget_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveCollectionTarget'
type QueryCoordCatalog_RemoveCollectionTarget_Call struct {
	*mock.Call
}

// RemoveCollectionTarget is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
func (_e *QueryCoordCatalog_Expecter) RemoveCollectionTarget(ctx interface{}, collectionID interface{}) *QueryCoordCatalog_RemoveCollectionTarget_Call {
	return &QueryCoordCatalog_RemoveCollectionTarget_Call{Call: _e.mock.On("RemoveCollectionTarget", ctx, collectionID)}
}

func (_c *QueryCoordCatalog_RemoveCollectionTarget_Call) Run(run func(ctx context.Context, collectionID int64)) *QueryCoordCatalog_RemoveCollectionTarget_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *QueryCoordCatalog_RemoveCollectionTarget_Call) Return(_a0 error) *QueryCoordCatalog_RemoveCollectionTarget_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_RemoveCollectionTarget_Call) RunAndReturn(run func(context.Context, int64) error) *QueryCoordCatalog_RemoveCollectionTarget_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveResourceGroup provides a mock function with given fields: ctx, rgName
func (_m *QueryCoordCatalog) RemoveResourceGroup(ctx context.Context, rgName string) error {
	ret := _m.Called(ctx, rgName)

	if len(ret) == 0 {
		panic("no return value specified for RemoveResourceGroup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, rgName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_RemoveResourceGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveResourceGroup'
type QueryCoordCatalog_RemoveResourceGroup_Call struct {
	*mock.Call
}

// RemoveResourceGroup is a helper method to define mock.On call
//   - ctx context.Context
//   - rgName string
func (_e *QueryCoordCatalog_Expecter) RemoveResourceGroup(ctx interface{}, rgName interface{}) *QueryCoordCatalog_RemoveResourceGroup_Call {
	return &QueryCoordCatalog_RemoveResourceGroup_Call{Call: _e.mock.On("RemoveResourceGroup", ctx, rgName)}
}

func (_c *QueryCoordCatalog_RemoveResourceGroup_Call) Run(run func(ctx context.Context, rgName string)) *QueryCoordCatalog_RemoveResourceGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *QueryCoordCatalog_RemoveResourceGroup_Call) Return(_a0 error) *QueryCoordCatalog_RemoveResourceGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_RemoveResourceGroup_Call) RunAndReturn(run func(context.Context, string) error) *QueryCoordCatalog_RemoveResourceGroup_Call {
	_c.Call.Return(run)
	return _c
}

// SaveCollection provides a mock function with given fields: ctx, collection, partitions
func (_m *QueryCoordCatalog) SaveCollection(ctx context.Context, collection *querypb.CollectionLoadInfo, partitions ...*querypb.PartitionLoadInfo) error {
	_va := make([]interface{}, len(partitions))
	for _i := range partitions {
		_va[_i] = partitions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, collection)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SaveCollection")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.CollectionLoadInfo, ...*querypb.PartitionLoadInfo) error); ok {
		r0 = rf(ctx, collection, partitions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_SaveCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveCollection'
type QueryCoordCatalog_SaveCollection_Call struct {
	*mock.Call
}

// SaveCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - collection *querypb.CollectionLoadInfo
//   - partitions ...*querypb.PartitionLoadInfo
func (_e *QueryCoordCatalog_Expecter) SaveCollection(ctx interface{}, collection interface{}, partitions ...interface{}) *QueryCoordCatalog_SaveCollection_Call {
	return &QueryCoordCatalog_SaveCollection_Call{Call: _e.mock.On("SaveCollection",
		append([]interface{}{ctx, collection}, partitions...)...)}
}

func (_c *QueryCoordCatalog_SaveCollection_Call) Run(run func(ctx context.Context, collection *querypb.CollectionLoadInfo, partitions ...*querypb.PartitionLoadInfo)) *QueryCoordCatalog_SaveCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*querypb.PartitionLoadInfo, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*querypb.PartitionLoadInfo)
			}
		}
		run(args[0].(context.Context), args[1].(*querypb.CollectionLoadInfo), variadicArgs...)
	})
	return _c
}

func (_c *QueryCoordCatalog_SaveCollection_Call) Return(_a0 error) *QueryCoordCatalog_SaveCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_SaveCollection_Call) RunAndReturn(run func(context.Context, *querypb.CollectionLoadInfo, ...*querypb.PartitionLoadInfo) error) *QueryCoordCatalog_SaveCollection_Call {
	_c.Call.Return(run)
	return _c
}

// SaveCollectionTargets provides a mock function with given fields: ctx, target
func (_m *QueryCoordCatalog) SaveCollectionTargets(ctx context.Context, target ...*querypb.CollectionTarget) error {
	_va := make([]interface{}, len(target))
	for _i := range target {
		_va[_i] = target[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SaveCollectionTargets")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*querypb.CollectionTarget) error); ok {
		r0 = rf(ctx, target...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_SaveCollectionTargets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveCollectionTargets'
type QueryCoordCatalog_SaveCollectionTargets_Call struct {
	*mock.Call
}

// SaveCollectionTargets is a helper method to define mock.On call
//   - ctx context.Context
//   - target ...*querypb.CollectionTarget
func (_e *QueryCoordCatalog_Expecter) SaveCollectionTargets(ctx interface{}, target ...interface{}) *QueryCoordCatalog_SaveCollectionTargets_Call {
	return &QueryCoordCatalog_SaveCollectionTargets_Call{Call: _e.mock.On("SaveCollectionTargets",
		append([]interface{}{ctx}, target...)...)}
}

func (_c *QueryCoordCatalog_SaveCollectionTargets_Call) Run(run func(ctx context.Context, target ...*querypb.CollectionTarget)) *QueryCoordCatalog_SaveCollectionTargets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*querypb.CollectionTarget, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*querypb.CollectionTarget)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *QueryCoordCatalog_SaveCollectionTargets_Call) Return(_a0 error) *QueryCoordCatalog_SaveCollectionTargets_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_SaveCollectionTargets_Call) RunAndReturn(run func(context.Context, ...*querypb.CollectionTarget) error) *QueryCoordCatalog_SaveCollectionTargets_Call {
	_c.Call.Return(run)
	return _c
}

// SavePartition provides a mock function with given fields: ctx, info
func (_m *QueryCoordCatalog) SavePartition(ctx context.Context, info ...*querypb.PartitionLoadInfo) error {
	_va := make([]interface{}, len(info))
	for _i := range info {
		_va[_i] = info[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SavePartition")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*querypb.PartitionLoadInfo) error); ok {
		r0 = rf(ctx, info...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_SavePartition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SavePartition'
type QueryCoordCatalog_SavePartition_Call struct {
	*mock.Call
}

// SavePartition is a helper method to define mock.On call
//   - ctx context.Context
//   - info ...*querypb.PartitionLoadInfo
func (_e *QueryCoordCatalog_Expecter) SavePartition(ctx interface{}, info ...interface{}) *QueryCoordCatalog_SavePartition_Call {
	return &QueryCoordCatalog_SavePartition_Call{Call: _e.mock.On("SavePartition",
		append([]interface{}{ctx}, info...)...)}
}

func (_c *QueryCoordCatalog_SavePartition_Call) Run(run func(ctx context.Context, info ...*querypb.PartitionLoadInfo)) *QueryCoordCatalog_SavePartition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*querypb.PartitionLoadInfo, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*querypb.PartitionLoadInfo)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *QueryCoordCatalog_SavePartition_Call) Return(_a0 error) *QueryCoordCatalog_SavePartition_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_SavePartition_Call) RunAndReturn(run func(context.Context, ...*querypb.PartitionLoadInfo) error) *QueryCoordCatalog_SavePartition_Call {
	_c.Call.Return(run)
	return _c
}

// SaveReplica provides a mock function with given fields: ctx, replicas
func (_m *QueryCoordCatalog) SaveReplica(ctx context.Context, replicas ...*querypb.Replica) error {
	_va := make([]interface{}, len(replicas))
	for _i := range replicas {
		_va[_i] = replicas[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SaveReplica")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*querypb.Replica) error); ok {
		r0 = rf(ctx, replicas...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_SaveReplica_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveReplica'
type QueryCoordCatalog_SaveReplica_Call struct {
	*mock.Call
}

// SaveReplica is a helper method to define mock.On call
//   - ctx context.Context
//   - replicas ...*querypb.Replica
func (_e *QueryCoordCatalog_Expecter) SaveReplica(ctx interface{}, replicas ...interface{}) *QueryCoordCatalog_SaveReplica_Call {
	return &QueryCoordCatalog_SaveReplica_Call{Call: _e.mock.On("SaveReplica",
		append([]interface{}{ctx}, replicas...)...)}
}

func (_c *QueryCoordCatalog_SaveReplica_Call) Run(run func(ctx context.Context, replicas ...*querypb.Replica)) *QueryCoordCatalog_SaveReplica_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*querypb.Replica, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*querypb.Replica)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *QueryCoordCatalog_SaveReplica_Call) Return(_a0 error) *QueryCoordCatalog_SaveReplica_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_SaveReplica_Call) RunAndReturn(run func(context.Context, ...*querypb.Replica) error) *QueryCoordCatalog_SaveReplica_Call {
	_c.Call.Return(run)
	return _c
}

// SaveResourceGroup provides a mock function with given fields: ctx, rgs
func (_m *QueryCoordCatalog) SaveResourceGroup(ctx context.Context, rgs ...*querypb.ResourceGroup) error {
	_va := make([]interface{}, len(rgs))
	for _i := range rgs {
		_va[_i] = rgs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SaveResourceGroup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*querypb.ResourceGroup) error); ok {
		r0 = rf(ctx, rgs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_SaveResourceGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveResourceGroup'
type QueryCoordCatalog_SaveResourceGroup_Call struct {
	*mock.Call
}

// SaveResourceGroup is a helper method to define mock.On call
//   - ctx context.Context
//   - rgs ...*querypb.ResourceGroup
func (_e *QueryCoordCatalog_Expecter) SaveResourceGroup(ctx interface{}, rgs ...interface{}) *QueryCoordCatalog_SaveResourceGroup_Call {
	return &QueryCoordCatalog_SaveResourceGroup_Call{Call: _e.mock.On("SaveResourceGroup",
		append([]interface{}{ctx}, rgs...)...)}
}

func (_c *QueryCoordCatalog_SaveResourceGroup_Call) Run(run func(ctx context.Context, rgs ...*querypb.ResourceGroup)) *QueryCoordCatalog_SaveResourceGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*querypb.ResourceGroup, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*querypb.ResourceGroup)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *QueryCoordCatalog_SaveResourceGroup_Call) Return(_a0 error) *QueryCoordCatalog_SaveResourceGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_SaveResourceGroup_Call) RunAndReturn(run func(context.Context, ...*querypb.ResourceGroup) error) *QueryCoordCatalog_SaveResourceGroup_Call {
	_c.Call.Return(run)
	return _c
}

// NewQueryCoordCatalog creates a new instance of QueryCoordCatalog. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueryCoordCatalog(t interface {
	mock.TestingT
	Cleanup(func())
}) *QueryCoordCatalog {
	mock := &QueryCoordCatalog{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
