package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ir0h4c4zodi08cf")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("5ig7akcs")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ir0h4c4zodi08cf")
		if err != nil {
			return err
		}

		// add
		del_members := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "5ig7akcs",
			"name": "members",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "df654xxn69ysgkf",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": []
			}
		}`), del_members)
		collection.Schema.AddField(del_members)

		return dao.SaveCollection(collection)
	})
}
