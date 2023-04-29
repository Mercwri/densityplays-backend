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

		// update
		edit_bnet_id := &schema.SchemaField{}
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
		}`), edit_bnet_id)
		collection.Schema.AddField(edit_bnet_id)

		// update
		edit_field := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rzsta4pr",
			"name": "field",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "n16hg12txqa3jw8",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": [
					"bnet_id",
					"class"
				]
			}
		}`), edit_field)
		collection.Schema.AddField(edit_field)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("df654xxn69ysgkf")
		if err != nil {
			return err
		}

		// update
		edit_bnet_id := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "7xrs8yrw",
			"name": "bnet_id",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), edit_bnet_id)
		collection.Schema.AddField(edit_bnet_id)

		// update
		edit_field := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rzsta4pr",
			"name": "field",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "n16hg12txqa3jw8",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": []
			}
		}`), edit_field)
		collection.Schema.AddField(edit_field)

		return dao.SaveCollection(collection)
	})
}
