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

		collection, err := dao.FindCollectionByNameOrId("7hiyd0u6zfs8mx0")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("5bhhnwg0")

		// add
		new_characters := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "fjrzarts",
			"name": "characters",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "df654xxn69ysgkf",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": [
					"username"
				]
			}
		}`), new_characters)
		collection.Schema.AddField(new_characters)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("7hiyd0u6zfs8mx0")
		if err != nil {
			return err
		}

		// add
		del_characters := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "5bhhnwg0",
			"name": "characters",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "n16hg12txqa3jw8",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": [
					"bnet_user",
					"bnet_id"
				]
			}
		}`), del_characters)
		collection.Schema.AddField(del_characters)

		// remove
		collection.Schema.RemoveField("fjrzarts")

		return dao.SaveCollection(collection)
	})
}
