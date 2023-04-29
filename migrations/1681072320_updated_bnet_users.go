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

		collection, err := dao.FindCollectionByNameOrId("df654xxn69ysgkf")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("7xrs8yrw")

		// add
		new_bnet_id := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "kfcjpsua",
			"name": "bnet_id",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_bnet_id)
		collection.Schema.AddField(new_bnet_id)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("df654xxn69ysgkf")
		if err != nil {
			return err
		}

		// add
		del_bnet_id := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "7xrs8yrw",
			"name": "bnet_id",
			"type": "number",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), del_bnet_id)
		collection.Schema.AddField(del_bnet_id)

		// remove
		collection.Schema.RemoveField("kfcjpsua")

		return dao.SaveCollection(collection)
	})
}
