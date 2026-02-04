package featurestore

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type FeatureStoreRepository struct {
	DB *gorm.DB
}

func NewFeatureStoreRepository() (*FeatureStoreRepository, error) {

	dsn := "root:rootpassword@tcp(127.0.0.1:3306)/feature_store?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
		return nil, err
	}
	return results, nil
}
