package examples

// {
//  "self": {
//   "namespace": "com.whatwapp.euchre",
//   "category": "analytics",
//   "name": "gameplay_picktrump",
//   "version": "1-0-0"
//  },
//  "type": "object",
//  "description": "the user accepts the trump",
//  "properties": {
//   "type": {
//    "type": "string",
//    "const": "gameplay"
//   },
//   "action": {
//    "type": "string",
//    "const": "pick trump"
//   },
//   "trump_suit": {
//    "type": "string",
//    "enum": [
//     "diamonds",
//     "hearts",
//     "clubs",
//     "spades"
//    ]
//   },
//   "is_alone": {
//    "type": "boolean"
//   }
//  },
//  "required": [
//   "type",
//   "action",
//   "trump_suit",
//   "is_alone"
//  ],
//  "$schema": "https://schema.whatwapp.io/api/v1/schemas/com.whatwapp/self-schema/schema/1-0-0",
//  "$id": "depot:com.whatwapp.euchre/analytics/gameplay_picktrump/1-0-0"
// }

// {
// 	"self": {
// 		"namespace": "com.whatwapp.euchre",
// 		"category": "analytics",
// 		"name": "transaction_ctx",
// 		"version": "1-0-0"
// 	},
// 	"type": "object",
// 	"properties": {
// 		"id": {
// 			"type": "string"
// 		},
// 		"coins": {
// 			"type": "integer"
// 		},
// 		"league.points": {
// 			"type": "integer"
// 		},
// 		"xp.points": {
// 			"type": "integer"
// 		}
// 	},
// 	"required": [],
// 	"meta": {
// 		"vertical": true,
// 		"dynamic": "integer"
// 	},
// 	"$schema": "https://schema.whatwapp.io/api/v1/schemas/com.whatwapp/self-schema/schema/1-0-0",
// 	"$id": "depot:com.whatwapp.euchre/analytics/transaction_ctx/1-0-0"
// }
