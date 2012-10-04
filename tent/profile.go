// Steve Phillips / elimisteve
// 2012.09.26

package tent

type Profile map[ProfileInfo]interface{}

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