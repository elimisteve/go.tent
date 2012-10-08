// Steve Phillips / elimisteve
// 2012.09.26

package tent

type Profile map[string]map[string]interface{}

// // The idea here is to have BasicInfo and CoreInfo implement the
// // `Marshal() ([]byte, error)` method, but it doesn't work:

// type Profile map[ProfileInfo]json.Marshaler

//
// Unmarshal'd response
//

// map[string]interface {}{
// 	"https://tent.io/types/info/basic/v0.1.0": map[string]interface {}{
// 		"location":"Northeast Ohio",
// 		"name":"Phil Jaenke",
// 		"permissions":map[string]interface {}{"public":true},
// 		"gender":"DB9",
// 		"avatar_url":"",
// 		"birthdate":"",
// 		"bio":"AKA @rootwyrm. You know. Writer, storage, unix, backup, virtualization, PgSQL, performance monitoring. Personal account. (Resisting urge to make joke about pitching tent == personal. CRAP!)"
// 	},
// 	"https://tent.io/types/info/core/v0.1.0": map[string]interface {}{
// 		"permissions": map[string]interface {}{"public":true},
// 		"licenses":[]interface {}(nil),
// 		"servers":[]interface {}{
// 			"https://rootwyrm.tent.is/tent"},
// 		"entity":"https://rootwyrm.tent.is"
// 	}
// }

//
// JSON version (actual response)
//

// "profile": {
//         "https://tent.io/types/info/basic/v0.1.0": {
//             "avatar_url": "https://si0.twimg.com/profile_images/2609036383/h3jgsk4sxozfp5ihop0d_reasonably_small.jpeg",
//             "bio": "Freelance web and mobile developer. Tent evangelist.",
//             "birthdate": "",
//             "gender": "",
//             "location": "The Netherlands",
//             "name": "Dave Clayton",
//             "permissions": {
//                 "public": true
//             }
//         },
//         "https://tent.io/types/info/core/v0.1.0": {
//             "entity": "https://redskyforge.tent.is",
//             "licenses": [],
//             "permissions": {
//                 "public": true
//             },
//             "servers": [
//                 "https://redskyforge.tent.is/tent"
//             ]
//         }
//     }
