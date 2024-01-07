package models

import "bez/bez_server/internal/types"

type IWishes struct {
}

func (*IWishes) DefaultSelectFields() string {
	return "id, userId, wishes, createdAt, updatedAt"
}

func (*IWishes) Create(wish types.Wishes) (int64, error) {
	fields, vPlacholders := BuildFields([]string{"userId", "wishes"})
	res, err := DB.Exec("INSERT INTO users ("+fields+") VALUES ("+vPlacholders+")", wish.UserId, wish.Wishes)

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (w *IWishes) GetById(id int64) (types.Wishes, error) {
	wish := types.Wishes{}
	err := DB.Get(&wish, "SELECT "+w.DefaultSelectFields()+" FROM wishes WHERE id = ?", id)

	if err != nil {
		return types.Wishes{}, err
	}

	return wish, nil
}

func (w *IWishes) GetByUserId(id int64) (types.Wishes, error) {
	wish := types.Wishes{}
	err := DB.Get(&wish, "SELECT "+w.DefaultSelectFields()+" FROM wishes WHERE userId = ?", id)

	if err != nil {
		return types.Wishes{}, err
	}

	return wish, nil
}

func (w *IWishes) Update(id int64, wish types.Wishes) (types.Wishes, error) {
	wish.Id = id
	_, err := DB.NamedExec("UPDATE wishes SET wishes = :wishes WHERE id = :id", wish)

	if err != nil {
		return types.Wishes{}, err
	}

	return wish, nil
}

func (w *IWishes) Upsert(wish types.Wishes) (types.Wishes, error) {
	_, err := DB.NamedExec("INSERT INTO wishes (userId, wishes) VALUES (:userId, :wishes) ON CONFLICT(userId) DO UPDATE SET wishes = :wishes", wish)

	if err != nil {
		return types.Wishes{}, err
	}

	wish, err = w.GetByUserId(wish.UserId)

	if err != nil {
		return types.Wishes{}, err
	}

	return wish, nil
}
