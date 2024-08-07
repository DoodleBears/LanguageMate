// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SentencesDao is the data access object for table sentences.
type SentencesDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SentencesColumns // columns contains all the column names of Table for convenient usage.
}

// SentencesColumns defines and stores column names for table sentences.
type SentencesColumns struct {
	UserUid   string // User Unique ID
	WordId    string // Word ID
	Content   string // Sentence content
	CreatedAt string // Created Time
	DeletedAt string // Soft deleted Time
}

// sentencesColumns holds the columns for table sentences.
var sentencesColumns = SentencesColumns{
	UserUid:   "user_uid",
	WordId:    "word_id",
	Content:   "content",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewSentencesDao creates and returns a new DAO object for table data access.
func NewSentencesDao() *SentencesDao {
	return &SentencesDao{
		group:   "default",
		table:   "sentences",
		columns: sentencesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SentencesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SentencesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SentencesDao) Columns() SentencesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SentencesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SentencesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SentencesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
