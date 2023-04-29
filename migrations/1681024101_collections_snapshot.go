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
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2023-04-09 07:01:28.147Z",
				"updated": "2023-04-09 07:01:28.148Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
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
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": null
						}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": false
				}
			},
			{
				"id": "df654xxn69ysgkf",
				"created": "2023-04-09 07:03:01.166Z",
				"updated": "2023-04-09 07:07:38.056Z",
				"name": "bnet_users",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "2icp2fsw",
						"name": "username",
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
						"id": "7xrs8yrw",
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
						"id": "wtqmdnmr",
						"name": "membership_type",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
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
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_lGOxqrZ` + "`" + ` ON ` + "`" + `bnet_users` + "`" + ` (` + "`" + `bnet_id` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "n16hg12txqa3jw8",
				"created": "2023-04-09 07:04:02.571Z",
				"updated": "2023-04-09 07:07:47.622Z",
				"name": "destiny_characters",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "cf9gdive",
						"name": "class",
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
						"id": "cwlv8t6y",
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
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_QYrrBXj` + "`" + ` ON ` + "`" + `destiny_characters` + "`" + ` (` + "`" + `bnet_id` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "7hiyd0u6zfs8mx0",
				"created": "2023-04-09 07:06:14.610Z",
				"updated": "2023-04-09 07:07:57.745Z",
				"name": "events",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "zbl4sbq1",
						"name": "duration",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
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
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "kenxad2g",
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
						"id": "wf9iny4q",
						"name": "type",
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
						"id": "fmylnabq",
						"name": "start_time",
						"type": "date",
						"required": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "kc1ciz5e",
						"name": "end_time",
						"type": "date",
						"required": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_KhIhzhS` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `bnet_id` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
