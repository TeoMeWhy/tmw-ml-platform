package featurestore

import (
	"palantir/configs"
	"palantir/errors"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type FeatureStoreRepository struct {
	DB *gorm.DB
}

func NewFeatureStoreRepository(cfg *configs.Config) (*FeatureStoreRepository, error) {

	db, err := gorm.Open(mysql.Open(cfg.MysqlDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fsRepo := &FeatureStoreRepository{
		DB: db,
	}

	return fsRepo, nil
}

func (r *FeatureStoreRepository) GetFeatures(table string, ids []string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	if err := r.DB.Table(table).Where("id IN ?", ids).Find(&results).Error; err != nil {

		if strings.Contains(err.Error(), "Table") && strings.Contains(err.Error(), "doesn't exist") {
			return nil, errors.ErrFeatureStoreDoesNotExist
		}

		return nil, err
	}
	return results, nil
}
