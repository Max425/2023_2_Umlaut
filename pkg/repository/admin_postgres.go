package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/jackc/pgx/v5"
)

type AdminPostgres struct {
	db *pgxpool.Pool
}

func NewAdminPostgres(db *pgxpool.Pool) *AdminPostgres {
	return &AdminPostgres{db: db}
}

func (r *AdminPostgres) GetAdmin(ctx context.Context, mail string) (model.Admin, error) {
	var admin model.Admin

	query, args, err := psql.Select(static.AdminDbField).From(adminTable).Where(sq.Eq{"mail": mail}).ToSql()

	if err != nil {
		return admin, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanAdmin(row, &admin)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.Admin{}, fmt.Errorf("admin with mail: %s not found", mail)
	}

	return admin, err
}

func (r *AdminPostgres) CreateFeedback(ctx context.Context, stat model.Feedback) (int, error) {
	var id int
	query, args, err := psql.Insert(feedbackTable).
		Columns("user_id", "rating", "liked", "need_fix", "comment_fix", "comment", "show").
		Values(stat.UserId, stat.Rating, stat.Liked, stat.NeedFix, stat.CommentFix, stat.Comment, stat.Show).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to create statistic. err: %w", err)
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	return id, err
}

func (r *AdminPostgres) CreateRecommendation(ctx context.Context, rec model.Recommendation) (int, error) {
	var id int
	query, args, err := psql.Insert(recommendationTable).
		Columns("user_id", "recommend", "show").
		Values(rec.UserId, rec.Recommend, rec.Show).
		ToSql()

	if err != nil {
		return 0, err
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	return id, err
}

func (r *AdminPostgres) GetFeedbacks(ctx context.Context) ([]model.Feedback, error) {
	query, args, err := psql.
		Select(static.FeedbackDbField).
		From(feedbackTable).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to get feedback. err: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get feedback. err: %w", err)
	}
	defer rows.Close()

	feedbacks, err := scanFeedbacks(rows)
	if err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func scanAdmin(row pgx.Row, admin *model.Admin) error {
	err := row.Scan(
		&admin.Id,
		&admin.Mail,
		&admin.PasswordHash,
		&admin.Salt,
	)

	return err
}

func scanFeedbacks(rows pgx.Rows) ([]model.Feedback, error) {
	var feedbacks []model.Feedback
	var err error
	for rows.Next() {
		var feedback model.Feedback
		err = rows.Scan(
			&feedback.Id,
			&feedback.UserId,
			&feedback.Rating,
			&feedback.Liked,
			&feedback.NeedFix,
			&feedback.CommentFix,
			&feedback.Comment,
			&feedback.Show,
			&feedback.CreatedAt,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		feedbacks = append(feedbacks, feedback)
	}
	if err != nil {
		return nil, fmt.Errorf("scan error in GetFeedbacks: %v", err)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error in GetFeedbacks: %v", rows.Err())
	}

	return feedbacks, nil
}
