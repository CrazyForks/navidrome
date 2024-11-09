package scanner2

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/Masterminds/squirrel"
	ppl "github.com/google/go-pipeline/pkg/pipeline"
	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/model"
)

type phaseRefreshAlbums struct {
	ds        model.DataStore
	ctx       context.Context
	libs      model.Libraries
	refreshed atomic.Uint32
	skipped   atomic.Uint32
}

func createPhaseRefreshAlbums(ctx context.Context, ds model.DataStore, libs model.Libraries) *phaseRefreshAlbums {
	return &phaseRefreshAlbums{ctx: ctx, ds: ds, libs: libs}
}

func (p *phaseRefreshAlbums) description() string {
	return "Refresh all new/changed albums"
}

func (p *phaseRefreshAlbums) producer() ppl.Producer[*model.Album] {
	return ppl.NewProducer(func(put func(album *model.Album)) error {
		for _, lib := range p.libs {
			// TODO Paginate
			albums, err := p.ds.Album(p.ctx).GetTouchedAlbums(lib.ID)
			if err != nil {
				return fmt.Errorf("error loading touched albums: %w", err)
			}
			if len(albums) == 0 {
				continue
			}
			log.Debug(p.ctx, "Scanner: Checking albums that may need refresh", "libraryId", lib.ID, "libraryName", lib.Name, "albumCount", len(albums))
			for _, album := range albums {
				put(&album)
			}
		}
		return nil
	}, ppl.Name("load albums from db"))
}

func (p *phaseRefreshAlbums) stages() []ppl.Stage[*model.Album] {
	return []ppl.Stage[*model.Album]{
		ppl.NewStage(p.filterUnmodified, ppl.Name("filter unmodified"), ppl.Concurrency(5)),
		ppl.NewStage(p.refreshAlbum, ppl.Name("refresh albums")),
	}
}

func (p *phaseRefreshAlbums) filterUnmodified(album *model.Album) (*model.Album, error) {
	mfs, err := p.ds.MediaFile(p.ctx).GetAll(model.QueryOptions{Filters: squirrel.Eq{"album_id": album.ID}})
	if err != nil {
		log.Error(p.ctx, "Error loading media files for album", "album_id", album.ID, err)
		return nil, err
	}
	if len(mfs) == 0 {
		log.Debug(p.ctx, "Scanner: album has no media files. Skipping", "album_id", album.ID,
			"name", album.Name, "songCount", album.SongCount, "updatedAt", album.UpdatedAt)
		p.skipped.Add(1)
		return nil, nil
	}

	newAlbum := mfs.ToAlbum()
	if album.Equals(newAlbum) {
		log.Trace("Scanner: album is up to date. Skipping", "album_id", album.ID,
			"name", album.Name, "songCount", album.SongCount, "updatedAt", album.UpdatedAt)
		p.skipped.Add(1)
		return nil, nil
	}
	return &newAlbum, nil
}

func (p *phaseRefreshAlbums) refreshAlbum(album *model.Album) (*model.Album, error) {
	if album == nil {
		return nil, nil
	}
	start := time.Now()
	err := p.ds.WithTx(func(tx model.DataStore) error {
		res := tx.Album(p.ctx).Put(album)
		p.refreshed.Add(1)
		log.Debug(p.ctx, "Scanner: refreshing album", "album_id", album.ID, "name", album.Name, "songCount", album.SongCount, "elapsed", time.Since(start))
		return res
	})
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (p *phaseRefreshAlbums) finalize(err error) error {
	refreshed := p.refreshed.Load()
	skipped := p.skipped.Load()
	if err != nil || refreshed+skipped == 0 {
		return err
	}
	logF := log.Info
	if refreshed == 0 {
		logF = log.Debug
	}
	logF(p.ctx, "Scanner: Finished checking for album updates", "refreshed", refreshed, "skipped", skipped, err)
	return p.ds.WithTx(func(tx model.DataStore) error {
		start := time.Now()
		stats, err := tx.Artist(p.ctx).RefreshStats()
		if err != nil {
			return fmt.Errorf("refreshing artist stats: %w", err)
		}
		log.Debug(p.ctx, "Scanner: Refreshed artist stats", "stats", stats, "elapsed", time.Since(start))
		return nil
	})
}