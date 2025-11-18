package firebird

import (
	"math/rand"
	"strings"

	"gorm.io/gorm/schema"
)

type NamingStrategy struct {
	schema.NamingStrategy
}

func (ns NamingStrategy) TableName(str string) string {
	return strings.ToUpper(ns.NamingStrategy.TableName(str))
}

func (ns NamingStrategy) ColumnName(table, column string) string {
	return strings.ToUpper(ns.NamingStrategy.ColumnName(table, column))
}

func (ns NamingStrategy) JoinTableName(table string) (name string) {
	return strings.ToUpper(ns.NamingStrategy.JoinTableName(table))
}

func (ns NamingStrategy) RelationshipFKName(relationship schema.Relationship) (name string) {
	name = strings.ToUpper(ns.NamingStrategy.RelationshipFKName(relationship))
	return safeFirebirdName(name)
}

func (ns NamingStrategy) CheckerName(table, column string) (name string) {
	name = strings.ToUpper(ns.NamingStrategy.CheckerName(table, column))
	return safeFirebirdName(name)
}

func (ns NamingStrategy) IndexName(table, column string) (name string) {
	name = strings.ToUpper(ns.NamingStrategy.IndexName(table, column))
	return safeFirebirdName(name)
}

func safeFirebirdName(name string) string {
	maxLen := 27
	suffixLen := 4

	if len(name) <= maxLen {
		return name
	}

	// Shorten and append 4-char random suffix
	prefixLen := maxLen - suffixLen
	if prefixLen < 1 {
		prefixLen = 1
	}

	short := name[:prefixLen] + randSuffix(suffixLen)
	return short
}

func randSuffix(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
