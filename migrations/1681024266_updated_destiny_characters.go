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

		collection, err := dao.FindCollectionByNameOrId("n16hg12txqa3jw8")
		if err != nil {
			return err
		}

		// update
		edit_bnet_user := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "k6gegay0",
			"name": "bnet_user",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "df654xxn69ysgkf",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": [
					"username",
					"bnet_id"
				]
			}
		}`), edit_bnet_user)
		collection.Schema.AddField(edit_bnet_user)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("n16hg12txqa3jw8")
		if err != nil {
			return err
		}

		// update
		edit_bnet_user := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "k6gegay0",
			"name": "field",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "df654xxn69ysgkf",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": [
					"username",
					"bnet_id"
				]
			}
		}`), edit_bnet_user)
		collection.Schema.AddField(edit_bnet_user)

		return dao.SaveCollection(collection)
	})
}
