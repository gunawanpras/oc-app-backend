package oracle

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"time"

	"github.com/gunawanpras/oc-app-backend/internal/core/profile/domain"
)

func (repo *ProfileRepository) GetProfile(ctx context.Context) (res domain.Profiles, err error) {
	var (
		profile  Profile
		profiles Profiles
	)

	repo.prepareGetProfile()

	rows, err := repo.statement.GetProfile.QueryxContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("data not found")
		}

		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		profile = Profile{}
		err = rows.StructScan(&profile)
		if err != nil {
			return res, err
		}

		profiles = append(profiles, profile)
	}

	if !profiles.Validate() {
		return res, errors.New("malformed data")
	}

	return profiles.ToDomain(), nil
}

func (repo *ProfileRepository) CreateProfile(ctx context.Context, newProfile domain.Profile) (res domain.Profile, err error) {
	repo.prepareCreateProfile()
	_, err = repo.statement.CreateProfile.ExecContext(ctx, newProfile.Name, newProfile.Detail, sql.Out{Dest: &res.ID})
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repo *ProfileRepository) UpdateProfile(ctx context.Context, updatedProfile domain.Profile) (res domain.Profile, err error) {
	repo.prepareUpdateProfile()
	_, err = repo.statement.UpdateProfile.ExecContext(ctx, updatedProfile.Name, updatedProfile.Detail, updatedProfile.ID, sql.Out{Dest: &res.ID})
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repo *ProfileRepository) DeleteProfile(ctx context.Context, profileID int64) (err error) {
	repo.prepareDeleteProfile()
	_, err = repo.statement.DeleteProfile.ExecContext(ctx, profileID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProfileRepository) GetProfileViaProcedure(ctx context.Context) (res domain.Profiles, err error) {
	var (
		profile  Profile
		profiles Profiles
		rows     driver.Rows
	)

	repo.prepareGetProfileViaProcedure()
	_, err = repo.statement.GetProfileViaProcedure.QueryContext(ctx, sql.Out{Dest: &rows})
	if err != nil {
		return res, err
	}
	defer rows.Close()

	cols := rows.Columns()
	values := make([]driver.Value, len(cols))

	for {
		if err = rows.Next(values); err != nil {
			if err == io.EOF {
				break
			}

			return res, err
		}

		profile = Profile{}
		profile.ID = values[0].(int64)
		profile.Name = values[1].(string)
		profile.Detail = values[2].([]byte)
		profile.CreatedAt = values[3].(time.Time)

		profiles = append(profiles, profile)
	}

	if !profiles.Validate() {
		return res, errors.New("malformed data")
	}

	return profiles.ToDomain(), nil
}
