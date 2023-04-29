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
		edit_characters := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rzsta4pr",
			"name": "characters",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "n16hg12txqa3jw8",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": [
					"class"
				]
			}
		}`), edit_characters)
		collection.Schema.AddField(edit_characters)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("df654xxn69ysgkf")
		if err != nil {
			return err
		}

		// update
		edit_characters := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rzsta4pr",
			"name": "characters",
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
		}`), edit_characters)
		collection.Schema.AddField(edit_characters)

		return dao.SaveCollection(collection)
	})
}
