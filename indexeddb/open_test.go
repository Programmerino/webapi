package indexeddb_test

import (
	"github.com/Programmerino/webapi"
	"github.com/Programmerino/webapi/dom/domcore"
	"github.com/Programmerino/webapi/indexeddb"
)

func Example_open() {
	request := webapi.GetWindow().IndexedDB().Open("test", nil)
	request.AddEventSuccess(func(event *domcore.Event, target *indexeddb.IDBRequest) {})
	request.AddEventUpgradeNeeded(func(event *domcore.Event, currentTarget *indexeddb.IDBOpenDBRequest) {
		db := indexeddb.IDBDatabaseFromJS(request.Result())
		db.CreateObjectStore("names", &indexeddb.IDBObjectStoreParameters{
			AutoIncrement: true,
		})
		/* javascript version is:
		   var db = event.target.result;
		   var objStore = db.createObjectStore("names", { autoIncrement : true });
		*/
	})
}
