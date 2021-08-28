package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Posts возвращает список всех статей из БД.
func (s *Storage) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			posts.id,
			authors.name,
			posts.title,
			posts.author_id,
			posts.content,
			posts.created_at
		FROM posts, authors
		WHERE
			posts.author_id = authors.id 
		ORDER BY id;
	`,
	)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.AuthorName,
			&t.Title,
			&t.AuthorID,
			&t.Content,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		posts = append(posts, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return posts, rows.Err()
}

// AddPost создает статью.
func (s *Storage) AddPost(p storage.Post) error {
	rows, err := s.db.Query(context.Background(), `
		INSERT INTO posts (author_id, title, content, created_at)
		VALUES ($1, $2, $3, $4);
	`,
		p.AuthorID, p.Title, p.Content, p.CreatedAt,
	)
	if err != nil {
		return err
	}
	// ВАЖНО не забыть проверить rows.Err()
	return rows.Err()
}

// UpdatePost обновляет статью.
func (s *Storage) UpdatePost(p storage.Post) error {
	rows, err := s.db.Query(context.Background(), `
		UPDATE posts
		SET created_at = $5,
			content = $4,
			title = $3,
			author_id = $2
		WHERE posts.id = $1;
	`,
		p.ID, p.AuthorID, p.Title, p.Content, p.CreatedAt,
	)
	if err != nil {
		return err
	}
	// ВАЖНО не забыть проверить rows.Err()
	return rows.Err()
}

// DeletePost удаляет статью по id.
func (s *Storage) DeletePost(p storage.Post) error {
	rows, err := s.db.Query(context.Background(), `
		DELETE FROM posts
		WHERE posts.id = $1;
	`,
		p.ID,
	)
	if err != nil {
		return err
	}
	// ВАЖНО не забыть проверить rows.Err()
	return rows.Err()
}

//// PostById возвращает статью по id.
//func (s *Storage) PostById(p storage.Post) ([]storage.Post, error) {
//	id := p.ID
//	rows, err := s.db.Query(context.Background(), `
//		SELECT
//			posts.id,
//			authors.name,
//			posts.title,
//			posts.author_id,
//			posts.content,
//			posts.created_at
//		FROM posts, authors
//		WHERE
//			posts.author_id = authors.id AND posts.id = $1
//		ORDER BY id;
//	`, id,
//	)
//	if err != nil {
//		return nil, err
//	}
//	var posts []storage.Post
//	// итерирование по результату выполнения запроса
//	// и сканирование каждой строки в переменную
//	for rows.Next() {
//		var t storage.Post
//		err = rows.Scan(
//			&t.ID,
//			&t.AuthorName,
//			&t.Title,
//			&t.AuthorID,
//			&t.Content,
//			&t.CreatedAt,
//		)
//		if err != nil {
//			return nil, err
//		}
//		// добавление переменной в массив результатов
//		posts = append(posts, t)
//
//	}
//	// ВАЖНО не забыть проверить rows.Err()
//	return posts, rows.Err()
//}
