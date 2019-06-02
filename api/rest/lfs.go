package rest

import "lfserv/api/store"

type LFS interface {
	BatchHandler(cStore *store.ContentStore, mStore *store.MetaStore)
	PutHandler(cStore *store.ContentStore, mStore *store.MetaStore)
	PostHandler(cStore *store.ContentStore, mStore *store.MetaStore)
	LocksHandler(cStore *store.ContentStore, mStore *store.MetaStore)
	LocksVerifyHandler(cStore *store.ContentStore, mStore *store.MetaStore)
	CreateLockHandler(cStore *store.ContentStore, mStore *store.MetaStore)
	DeleteLockHandler(cStore *store.ContentStore, mStore *store.MetaStore)
	GetContentHandler(cStore *store.ContentStore, mStore *store.MetaStore)
	GetMetaHandler(cStore *store.ContentStore, mStore *store.MetaStore)
	VerifyHandler(cStore *store.ContentStore, mStore *store.MetaStore)
}
