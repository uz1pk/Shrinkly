package model

func GetAll() ([]Shrinkly, error) {
	var shrinks []Shrinkly

	res := db.Find(&shrinks)
	if res.Error != nil {
		return []Shrinkly{}, res.Error
	}

	return shrinks, nil
}

func GetOneShrink(id uint64) (Shrinkly, error) {
	var shrink Shrinkly

	res := db.Where("id = ?", id).First(shrink)
	if res.Error != nil {
		return Shrinkly{}, res.Error
	}

	return shrink, nil
}

func CreateShrink(shrink Shrinkly) error {
	res := db.Create(&shrink)
	return res.Error
}

func UpdateShrink(shrink Shrinkly) error {
	res := db.Save(&shrink)
	return res.Error
}

func DeleteShrink(id uint64) error {
	res := db.Unscoped().Delete(&Shrinkly{}, id)
	return res.Error
}

func FindShrinkFromUrl(url string) (Shrinkly, error) {
	var shrink Shrinkly
	res := db.Where("shrinkly = s", url).First(&shrink)
	return shrink, res.Error
}
