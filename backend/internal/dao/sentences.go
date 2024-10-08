// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"backend/internal/dao/internal"
)

// internalSentencesDao is internal type for wrapping internal DAO implements.
type internalSentencesDao = *internal.SentencesDao

// sentencesDao is the data access object for table sentences.
// You can define custom methods on it to extend its functionality as you wish.
type sentencesDao struct {
	internalSentencesDao
}

var (
	// Sentences is globally public accessible object for table sentences operations.
	Sentences = sentencesDao{
		internal.NewSentencesDao(),
	}
)

// Fill with you ideas below.
