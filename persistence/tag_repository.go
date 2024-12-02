package persistence

import (
	"context"
	"fmt"
	"slices"

	. "github.com/Masterminds/squirrel"
	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/model"
	"github.com/pocketbase/dbx"
)

type tagRepository struct {
	sqlRepository
}

func NewTagRepository(ctx context.Context, db dbx.Builder) model.TagRepository {
	r := &tagRepository{}
	r.ctx = ctx
	r.db = db
	r.tableName = "tag"
	r.registerModel(&model.Tag{}, nil)
	return r
}

func (r *tagRepository) Add(tags ...model.Tag) error {
	for chunk := range slices.Chunk(tags, 200) {
		sq := Insert(r.tableName).Columns("id", "tag_name", "tag_value").
			Suffix("on conflict (id) do nothing")
		for _, t := range chunk {
			sq = sq.Values(t.ID, t.TagName, t.TagValue)
		}
		_, err := r.executeSQL(sq)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *tagRepository) purgeNonUsed() error {
	del := Delete(r.tableName).Where(`
		id not in 
(select atom from media_file left join json_tree(media_file.tags, '$') where atom is not null and key = 'id')
`)
	c, err := r.executeSQL(del)
	if err != nil {
		return fmt.Errorf("error purging unused tags: %w", err)
	}
	if c > 0 {
		log.Debug(r.ctx, "Purged unused tags", "totalDeleted", c)
	}
	return err
}