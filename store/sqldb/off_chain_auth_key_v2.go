package sqldb

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	corespdb "github.com/zkMeLabs/mechain-storage-provider/core/spdb"
	"github.com/zkMeLabs/mechain-storage-provider/pkg/log"
)

// InsertAuthKeyV2 insert a new record into OffChainAuthKeyV2
func (s *SpDBImpl) InsertAuthKeyV2(newRecord *corespdb.OffChainAuthKeyV2) error {
	result := &OffChainAuthKeyV2Table{
		UserAddress:  newRecord.UserAddress,
		Domain:       newRecord.Domain,
		PublicKey:    newRecord.PublicKey,
		ExpiryDate:   newRecord.ExpiryDate,
		CreatedTime:  newRecord.CreatedTime,
		ModifiedTime: newRecord.ModifiedTime,
	}

	err := s.db.Table((&OffChainAuthKeyV2Table{}).TableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_address"}, {Name: "domain"}, {Name: "public_key"}},
		DoUpdates: clause.AssignmentColumns([]string{"expiry_date", "modified_time"}),
	}).Create(result).Error
	if err != nil {
		return fmt.Errorf("failed to insert record in OffChainAuthKeyV2Table: %s", err)
	}
	return nil
}

// GetAuthKeyV2 get OffChainAuthKeyV2 from OffChainAuthKeyV2Table
func (s *SpDBImpl) GetAuthKeyV2(userAddress string, domain string, publicKey string) (*corespdb.OffChainAuthKeyV2, error) {
	queryKeyReturn := &OffChainAuthKeyV2Table{}
	result := s.db.First(queryKeyReturn, "user_address = ? and domain =? and public_key=?", userAddress, domain, publicKey)

	if result.Error != nil {
		if errIsNotFound(result.Error) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query OffChainAuthKey table: %s", result.Error)
	}
	return &corespdb.OffChainAuthKeyV2{
		UserAddress:  queryKeyReturn.UserAddress,
		Domain:       queryKeyReturn.Domain,
		PublicKey:    queryKeyReturn.PublicKey,
		ExpiryDate:   queryKeyReturn.ExpiryDate,
		CreatedTime:  queryKeyReturn.CreatedTime,
		ModifiedTime: queryKeyReturn.ModifiedTime,
	}, nil
}

// ClearExpiredOffChainAuthKeys will clear those expired off chain auth keys from OffChainAuthKeyV2Table
func (s *SpDBImpl) ClearExpiredOffChainAuthKeys() error {
	result := s.db.Table(OffChainAuthKeyV2TableName).Where("expiry_date < ? ", time.Now()).Delete(&OffChainAuthKeyV2Table{})
	if result.Error != nil {
		return result.Error
	}
	log.Infow("ClearExpiredOffChainAuthKeys successfully.", "removed rows are ", result.RowsAffected)

	return nil
}

// ListAuthKeysV2 list user public keys
func (s *SpDBImpl) ListAuthKeysV2(userAddress string, domain string) ([]string, error) {
	var (
		result     *gorm.DB
		publicKeys []string
	)
	var results []*OffChainAuthKeyV2Table

	result = s.db.Table(OffChainAuthKeyV2TableName).Where("user_address = ? and domain =? ", userAddress, domain).Find(&results)

	if result.Error != nil {
		if errIsNotFound(result.Error) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query ListAuthKeysV2 table: %s", result.Error)
	}

	for _, res := range results {
		publicKeys = append(publicKeys, res.PublicKey)
	}

	return publicKeys, nil
}

// DeleteAuthKeysV2 delete user public keys
func (s *SpDBImpl) DeleteAuthKeysV2(userAddress string, domain string, publicKeys []string) (bool, error) {
	result := s.db.Table(OffChainAuthKeyV2TableName).Where("user_address = ? and domain =? and public_key in (?)", userAddress, domain, publicKeys).Delete(&OffChainAuthKeyV2Table{})
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
