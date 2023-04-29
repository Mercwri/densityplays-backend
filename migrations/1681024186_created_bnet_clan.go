package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "ir0h4c4zodi08cf",
			"created": "2023-04-09 07:09:46.375Z",
			"updated": "2023-04-09 07:09:46.375Z",
			"name": "bnet_clan",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "gyfhkdcs",
					"name": "clan_name",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "8s7qr9wq",
					"name": "bnet_id",
					"type": "number",
					"required": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null
					}
				},
				{
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
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_vy5QQBx` + "`" + ` ON ` + "`" + `bnet_clan` + "`" + ` (` + "`" + `clan_name` + "`" + `)"
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ir0h4c4zodi08cf")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
