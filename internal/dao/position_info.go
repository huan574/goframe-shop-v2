// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-shop-v2/internal/dao/internal"
)

// internalPositionInfoDao is internal type for wrapping internal DAO implements.
type internalPositionInfoDao = *internal.PositionInfoDao

// positionInfoDao is the data access object for table position_info.
// You can define custom methods on it to extend its functionality as you wish.
type positionInfoDao struct {
	internalPositionInfoDao
}

var (
	// PositionInfo is globally public accessible object for table position_info operations.
	PositionInfo = positionInfoDao{
		internal.NewPositionInfoDao(),
	}
)

// Fill with you ideas below.