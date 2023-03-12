// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"shop/internal/dao/internal"
)

// internalAddressInfoDao is internal type for wrapping internal DAO implements.
type internalAddressInfoDao = *internal.AddressInfoDao

// addressInfoDao is the data access object for table address_info.
// You can define custom methods on it to extend its functionality as you wish.
type addressInfoDao struct {
	internalAddressInfoDao
}

var (
	// AddressInfo is globally public accessible object for table address_info operations.
	AddressInfo = addressInfoDao{
		internal.NewAddressInfoDao(),
	}
)

// Fill with you ideas below.
