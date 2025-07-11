package persistence

import (
	"context"
	"sync"
	"time"

	. "github.com/Masterminds/squirrel"
	"github.com/navidrome/navidrome/conf"
	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/model"
	"github.com/navidrome/navidrome/utils/run"
	"github.com/pocketbase/dbx"
)

type libraryRepository struct {
	sqlRepository
}

var (
	libCache = map[int]string{}
	libLock  sync.RWMutex
)

func NewLibraryRepository(ctx context.Context, db dbx.Builder) model.LibraryRepository {
	r := &libraryRepository{}
	r.ctx = ctx
	r.db = db
	r.registerModel(&model.Library{}, nil)
	return r
}

func (r *libraryRepository) Get(id int) (*model.Library, error) {
	sq := r.newSelect().Columns("*").Where(Eq{"id": id})
	var res model.Library
	err := r.queryOne(sq, &res)
	return &res, err
}

func (r *libraryRepository) GetPath(id int) (string, error) {
	l := func() string {
		libLock.RLock()
		defer libLock.RUnlock()
		if l, ok := libCache[id]; ok {
			return l
		}
		return ""
	}()
	if l != "" {
		return l, nil
	}

	libLock.Lock()
	defer libLock.Unlock()
	libs, err := r.GetAll()
	if err != nil {
		log.Error(r.ctx, "Error loading libraries from DB", err)
		return "", err
	}
	for _, l := range libs {
		libCache[l.ID] = l.Path
	}
	if l, ok := libCache[id]; ok {
		return l, nil
	} else {
		return "", model.ErrNotFound
	}
}

func (r *libraryRepository) Put(l *model.Library) error {
	cols := map[string]any{
		"name":        l.Name,
		"path":        l.Path,
		"remote_path": l.RemotePath,
		"updated_at":  time.Now(),
	}
	if l.ID != 0 {
		cols["id"] = l.ID
	}

	sq := Insert(r.tableName).SetMap(cols).
		Suffix(`on conflict(id) do update set name = excluded.name, path = excluded.path, 
					remote_path = excluded.remote_path, updated_at = excluded.updated_at`)
	_, err := r.executeSQL(sq)
	if err != nil {
		libLock.Lock()
		defer libLock.Unlock()
		libCache[l.ID] = l.Path
	}
	return err
}

const hardCodedMusicFolderID = 1

// TODO Remove this method when we have a proper UI to add libraries
// This is a temporary method to store the music folder path from the config in the DB
func (r *libraryRepository) StoreMusicFolder() error {
	sq := Update(r.tableName).Set("path", conf.Server.MusicFolder).
		Set("updated_at", time.Now()).
		Where(Eq{"id": hardCodedMusicFolderID})
	_, err := r.executeSQL(sq)
	if err != nil {
		libLock.Lock()
		defer libLock.Unlock()
		libCache[hardCodedMusicFolderID] = conf.Server.MusicFolder
	}
	return err
}

func (r *libraryRepository) AddArtist(id int, artistID string) error {
	sq := Insert("library_artist").Columns("library_id", "artist_id").Values(id, artistID).
		Suffix(`on conflict(library_id, artist_id) do nothing`)
	_, err := r.executeSQL(sq)
	if err != nil {
		return err
	}
	return nil
}

func (r *libraryRepository) ScanBegin(id int, fullScan bool) error {
	sq := Update(r.tableName).
		Set("last_scan_started_at", time.Now()).
		Set("full_scan_in_progress", fullScan).
		Where(Eq{"id": id})
	_, err := r.executeSQL(sq)
	return err
}

func (r *libraryRepository) ScanEnd(id int) error {
	sq := Update(r.tableName).
		Set("last_scan_at", time.Now()).
		Set("full_scan_in_progress", false).
		Set("last_scan_started_at", time.Time{}).
		Where(Eq{"id": id})
	_, err := r.executeSQL(sq)
	if err != nil {
		return err
	}
	// https://www.sqlite.org/pragma.html#pragma_optimize
	_, err = r.executeSQL(Expr("PRAGMA optimize=0x10012;"))
	return err
}

func (r *libraryRepository) ScanInProgress() (bool, error) {
	query := r.newSelect().Where(NotEq{"last_scan_started_at": time.Time{}})
	count, err := r.count(query)
	return count > 0, err
}

func (r *libraryRepository) RefreshStats(id int) error {
	var songsRes, albumsRes, artistsRes, foldersRes, filesRes, missingRes struct{ Count int64 }
	var sizeRes struct{ Sum int64 }

	err := run.Parallel(
		func() error {
			return r.queryOne(Select("count(*) as count").From("media_file").Where(Eq{"library_id": id, "missing": false}), &songsRes)
		},
		func() error {
			return r.queryOne(Select("count(*) as count").From("album").Where(Eq{"library_id": id, "missing": false}), &albumsRes)
		},
		func() error {
			return r.queryOne(Select("count(*) as count").From("library_artist la").
				Join("artist a on la.artist_id = a.id").
				Where(Eq{"la.library_id": id, "a.missing": false}), &artistsRes)
		},
		func() error {
			return r.queryOne(Select("count(*) as count").From("folder").
				Where(And{
					Eq{"library_id": id, "missing": false},
					Gt{"num_audio_files": 0},
				}), &foldersRes)
		},
		func() error {
			return r.queryOne(Select("ifnull(sum(num_audio_files + num_playlists + json_array_length(image_files)),0) as count").
				From("folder").Where(Eq{"library_id": id, "missing": false}), &filesRes)
		},
		func() error {
			return r.queryOne(Select("count(*) as count").From("media_file").Where(Eq{"library_id": id, "missing": true}), &missingRes)
		},
		func() error {
			return r.queryOne(Select("ifnull(sum(size),0) as sum").From("album").Where(Eq{"library_id": id, "missing": false}), &sizeRes)
		},
	)()
	if err != nil {
		return err
	}

	sq := Update(r.tableName).
		Set("total_songs", songsRes.Count).
		Set("total_albums", albumsRes.Count).
		Set("total_artists", artistsRes.Count).
		Set("total_folders", foldersRes.Count).
		Set("total_files", filesRes.Count).
		Set("total_missing_files", missingRes.Count).
		Set("total_size", sizeRes.Sum).
		Set("updated_at", time.Now()).
		Where(Eq{"id": id})
	_, err = r.executeSQL(sq)
	return err
}

func (r *libraryRepository) GetAll(ops ...model.QueryOptions) (model.Libraries, error) {
	sq := r.newSelect(ops...).Columns("*")
	res := model.Libraries{}
	err := r.queryAll(sq, &res)
	return res, err
}

var _ model.LibraryRepository = (*libraryRepository)(nil)
