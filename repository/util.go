package repository

import (
	"auth/util"
	"context"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/samber/lo"

	"auth/ent"
	"auth/proto/gcommon"
)

type SortDirection int

const Ascending SortDirection = 0
const Descending SortDirection = 1

type Sortable struct {
	Field     string
	Direction SortDirection
}

type Pagination struct {
	Offset        int
	Limit         int
	Sort          *Sortable
	PagingIgnored bool
}

func ToSortable(sort string) *Sortable {
	if sort == "" {
		return nil
	}

	sorts := util.StringToList(sort, ",")

	sortField := sorts[0]
	sortDirection := Ascending

	if len(sorts) == 2 && strings.EqualFold(sorts[1], "desc") {
		sortDirection = Descending
	}

	return &Sortable{
		Field:     sortField,
		Direction: sortDirection,
	}
}

func ToPagination(pageable *gcommon.Pageable) Pagination {
	page := pageable.Page
	size := pageable.Size

	offset := int(page * size)
	sort := ToSortable(pageable.Sort)

	return Pagination{
		Offset:        offset,
		Limit:         int(size),
		Sort:          sort,
		PagingIgnored: pageable.PagingIgnored,
	}
}

func ToOrderSpecifier(sort *Sortable, defaultSortField string, defaultSortDirection SortDirection) func(selector *sql.Selector) {
	if sort == nil {
		sort = &Sortable{
			Field:     defaultSortField,
			Direction: defaultSortDirection,
		}
	}

	ent.Asc()

	return lo.Ternary(sort.Direction == Ascending, ent.Asc(sort.Field), ent.Desc(sort.Field))
}

func rollback(tx *ent.Tx, err error) error {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		err = fmt.Errorf("%w: error when rolling back. %v", err, rollbackErr)
	}

	return err
}

func commit(tx *ent.Tx) error {
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func WithTransaction(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	var err error

	tx, err := client.Tx(ctx)

	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if v := recover(); v != nil {
			if rollbackErr := rollback(tx, fmt.Errorf("panic: %v", v)); rollbackErr != nil {
				err = rollbackErr
			}
		}
	}()

	if execErr := fn(tx); execErr != nil {
		err = rollback(tx, execErr)
	} else {
		err = commit(tx)
	}

	return err
}
